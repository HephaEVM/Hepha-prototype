package vm

import (
	"crypto/sha256"
	_ "embed"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"sort"
	"sync"
	"sync/atomic"

	"github.com/holiman/uint256"
	"go-eth1.16.5-evm/regionexec"
	"go-eth1.16.5-evm/regionexec/prepare"
)

// A catalog entry covers the complete verified region, including guards, gas,
// staged memory writes and helper layouts. Terminal PCs and unconsumed output
// constants are late-bound; every computation bound remains part of the key.
// Unknown shapes continue through the unchanged dynamic compiler.
type precompiledFlowEntry struct {
	bind func(*precompiledFlowArguments) regionexec.Executor
}
type precompiledFlowArguments struct {
	exits  [4]uint64
	values []uint256.Int
}

type precompiledFlowCatalog struct {
	version    string
	entries    map[string]*precompiledFlowEntry
	buckets    map[uint64]bool
	once       sync.Once
	valid      bool
	memoPrefix []byte
}

var precompiledFlows = &precompiledFlowCatalog{}
var precompiledFlowHits, precompiledFlowProbes, precompiledFlowMemoHits atomic.Uint64

// Matching results are compiler/catalog-specific; this does not invalidate
// the existing dynamic lowering cache when the static catalog changes.
//
//go:embed flow_precompiled.go
var precompiledFlowImplementation string

type precompiledFlowMemo struct {
	Shape  string
	Values []uint256.Int
}

// Used only by the offline collector overlay; normal clients leave it nil.
var precompiledFlowObserver func([]*preparedFlow, bool, bool, *demandPlan)

func PrecompiledFlowStats() map[string]uint64 {
	valid := uint64(0)
	if len(precompiledFlows.entries) > 0 && precompiledFlows.version == flowPreparationVersion() {
		valid = 1
	}
	return map[string]uint64{"hits": precompiledFlowHits.Load(), "probes": precompiledFlowProbes.Load(), "memoHits": precompiledFlowMemoHits.Load(), "catalogEntries": uint64(len(precompiledFlows.entries)), "catalogValid": valid}
}
func precompiledFlowBucket(ps []*preparedFlow, proof, shared bool) uint64 {
	h := uint64(14695981039346656037)
	add := func(v int) { h ^= uint64(v); h *= 1099511628211 }
	add(len(ps))
	if proof {
		add(1)
	} else {
		add(0)
	}
	if shared {
		add(1)
	} else {
		add(0)
	}
	for _, p := range ps {
		r := p.original
		add(r.Inputs)
		add(r.Drop)
		add(r.MaxStack)
		add(len(r.nodes))
		add(len(r.outputs))
	}
	return h
}

// Canonicalize only constants whose values are never consumed by a node.
// The optimized graph and all computation bounds remain exact. This also
// records value aliasing: two output constants share a parameter iff equal.
func precompiledFlowKey(ps []*preparedFlow, proof, shared bool) (string, []*arithmeticRegion, []uint256.Int) {
	eligible := map[uint256.Int]bool{}
	optimized := make([]*hotRegion, len(ps))
	bounds := make([][]hotInterval, len(ps))
	for i, p := range ps {
		optimized[i], bounds[i] = p.optimized(proof)
		for _, n := range optimized[i].nodes {
			if n.op == PUSH32 {
				eligible[n.value] = true
			}
		}
	}
	for _, r := range optimized {
		for _, n := range r.nodes {
			for _, a := range n.args {
				if a >= hotRegionInputs {
					source := r.nodes[a-hotRegionInputs]
					if source.op == PUSH32 {
						eligible[source.value] = false
					}
				}
			}
		}
	}
	ids := map[uint256.Int]int{}
	var values []uint256.Int
	normalized := make([]*preparedFlow, len(ps))
	h := sha256.New()
	for i, r := range optimized {
		clone := *r
		clone.Exit = 0
		clone.nodes = append([]hotRegionNode(nil), r.nodes...)
		for j := range clone.nodes {
			n := &clone.nodes[j]
			if n.op == PUSH32 && eligible[n.value] {
				id, ok := ids[n.value]
				if !ok {
					id = len(values)
					ids[n.value] = id
					values = append(values, n.value)
				}
				n.value = uint256.Int{uint64(id), 0, 0, 1 << 63}
				fmt.Fprintf(h, "parameter:%d/%d/%d;", i, j, id)
			}
		}
		// Bounds guide narrower native operations, so matching only node opcodes
		// would be insufficient. Parameter bounds cannot affect any computation.
		for j, v := range bounds[i] {
			if j >= hotRegionInputs && j-hotRegionInputs < len(r.nodes) {
				n := r.nodes[j-hotRegionInputs]
				if n.op == PUSH32 && eligible[n.value] {
					continue
				}
			}
			fmt.Fprintf(h, "%d/%d:%s:%s;", i, j, v.lo.Text(16), v.hi.Text(16))
		}
		clone.key = clone.fingerprint()
		normalized[i] = &preparedFlow{original: &clone}
	}
	key, refs := flowPreparationKey(normalized, proof, shared)
	h.Write(key)
	return hex.EncodeToString(h.Sum(nil)), refs, values
}
func bindPrecompiledFlow(ps []*preparedFlow, proof, shared bool) *demandPlan {
	// Existing catalog functions share prefixes and are not CR-off controls.
	if paperWithoutReuse && shared {
		return nil
	}
	if regionFamiliesEnabled {
		if p := bindReusableFlow(ps, proof, shared); p != nil {
			return p
		}
	}
	c := precompiledFlows
	if len(c.entries) == 0 || !shared || len(ps) < 2 || len(ps) > 4 {
		return nil
	}
	if !c.buckets[precompiledFlowBucket(ps, proof, shared)] {
		return nil
	}
	c.once.Do(func() {
		c.valid = c.version == flowPreparationVersion()
		h := sha256.New()
		h.Write([]byte(precompiledFlowImplementation))
		h.Write([]byte(c.version))
		keys := make([]string, 0, len(c.entries))
		for k := range c.entries {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		for _, k := range keys {
			h.Write([]byte(k))
		}
		c.memoPrefix = []byte(fmt.Sprintf("precompiled/%x/", h.Sum(nil)))
	})
	if !c.valid {
		return nil
	}
	// A cached structural decision avoids re-running optimization and proof on
	// every process restart. Rebind helper references from the current region;
	// never persist Go function pointers or transaction values.
	raw, refs := flowPreparationKey(ps, proof, shared)
	memoKey := append(append([]byte(nil), c.memoPrefix...), raw...)
	cache := getFlowPreparationCache()
	if record, ok := cache.get(memoKey, refs, len(ps), true); ok {
		if record.Error != "" {
			precompiledFlowMemoHits.Add(1)
			return nil
		}
		var memo precompiledFlowMemo
		if json.Unmarshal([]byte(record.Source), &memo) == nil {
			if entry := c.entries[memo.Shape]; entry != nil {
				precompiledFlowMemoHits.Add(1)
				return instantiatePrecompiledFlow(ps, refs, memo.Values, entry, record)
			}
		}
	}
	precompiledFlowProbes.Add(1)
	key, _, values := precompiledFlowKey(ps, proof, shared)
	entry := c.entries[key]
	if entry == nil {
		cache.put(memoKey, flowPreparationRecord{Error: "no matching precompiled structure"})
		return nil
	}
	source, _ := json.Marshal(precompiledFlowMemo{Shape: key, Values: values})
	record := flowPreparationRecord{Source: string(source), Helpers: flowPreparationHelpers(ps, refs, proof, shared), Reports: precompiledFlowReports(ps)}
	cache.put(memoKey, record)
	return instantiatePrecompiledFlow(ps, refs, values, entry, record)
}
func instantiatePrecompiledFlow(ps []*preparedFlow, refs []*arithmeticRegion, values []uint256.Int, entry *precompiledFlowEntry, record flowPreparationRecord) *demandPlan {
	p := &demandPlan{branch: true, shared: true, flow: record.Reports, forward: make([]int, len(ps)), dead: make([]int, len(ps))}
	args := precompiledFlowArguments{values: values}
	for i, pp := range ps {
		p.regions = append(p.regions, pp.original)
		args.exits[i] = pp.original.Exit
	}
	p.helpers = record.bindHelpers(refs)
	// The existing Ready path returns this function immediately. No new hot-loop
	// dispatch, graph evaluation, shorter region, or runtime shape test is added.
	p.slot = prepare.ReadySlot(entry.bind(&args))
	precompiledFlowHits.Add(1)
	return p.finishBinding()
}

func precompiledFlowReports(ps []*preparedFlow) []flowReport {
	out := make([]flowReport, len(ps))
	count := func(r *hotRegion, b []hotInterval) (int, int) {
		h, g := 0, 0
		for _, n := range r.nodes {
			if n.helper != nil {
				h++
			}
			if n.op == JUMPI && !b[n.args[0]].exact() {
				g++
			}
		}
		return h, g
	}
	for i, p := range ps {
		old, ob := p.optimized(false)
		next, nb := p.optimized(true)
		oh, og := count(old, ob)
		nh, ng := count(next, nb)
		out[i] = flowReport{HelpersRemoved: oh - nh, GuardsRemoved: og - ng}
	}
	return out
}
