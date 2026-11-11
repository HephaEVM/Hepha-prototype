package vm

import (
	"crypto/sha256"
	_ "embed"
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"

	"github.com/holiman/uint256"
	"go-eth1.16.5-evm/regionexec/prepare"
)

// Enabled by the region_families build tag; no strategy branch is added to the
// instruction loop. Every family still covers the entire original region.
var regionFamiliesEnabled bool

//go:embed flow_family.go
var flowFamilyImplementation string
var flowFamilyOnce sync.Once
var flowFamilyPrefix, flowFamilyCompilerVersion string
var flowFamilyBound, flowFamilyMemo, flowFamilyStatic atomic.Uint64
var flowFamilySourceBuilt, flowFamilySourceReused atomic.Uint64
var flowFamilyExitRE = regexp.MustCompile(`(s\.Path\s*=\s*(\d+)\s*\n\s*)s\.PC\s*=\s*\d+`)
var flowFamilyLiteralRE = regexp.MustCompile(`uint256.Int\{([0-9]+),\s*([0-9]+),\s*([0-9]+),\s*([0-9]+)\}`)
var flowFamilySignatureRE = regexp.MustCompile(`func Exec\(s \*regionexec.Frame\)\s*bool\s*\{`)

type flowFamilyIndex struct {
	Shape  string
	Values []uint256.Int
}

func FlowFamilyStats() map[string]uint64 {
	return map[string]uint64{"bindings": flowFamilyBound.Load(), "memoHits": flowFamilyMemo.Load(), "staticBindings": flowFamilyStatic.Load(), "sourceBuilds": flowFamilySourceBuilt.Load(), "sourceReuses": flowFamilySourceReused.Load()}
}
func familyPrefix() string {
	flowFamilyOnce.Do(func() {
		flowFamilyCompilerVersion = flowPreparationVersion()
		flowFamilyPrefix = fmt.Sprintf("family-v1-%x", sha256.Sum256([]byte(flowFamilyCompilerVersion+precompiledFlowImplementation+flowFamilyImplementation)))
	})
	return flowFamilyPrefix
}

func bindReusableFlow(ps []*preparedFlow, proof, shared bool) *demandPlan {
	if !shared || len(ps) < 2 || len(ps) > 4 {
		return nil
	}
	manager := getDemandManager()
	if manager == nil {
		return nil
	}
	cache := getFlowPreparationCache()
	raw, refs := flowPreparationKey(ps, proof, shared)
	rawKey := append([]byte(familyPrefix()+"/binding/"), raw...)
	if rec, ok := cache.get(rawKey, refs, len(ps), true); ok {
		if rec.Error != "" {
			return nil
		}
		var idx flowFamilyIndex
		if json.Unmarshal([]byte(rec.Source), &idx) == nil {
			static := precompiledFlows.entries[idx.Shape]
			sourceKey := []byte(familyPrefix() + "/source/" + idx.Shape)
			present, _ := cache.db.Has(sourceKey, nil)
			if static != nil && precompiledFlows.version == flowFamilyCompilerVersion || present {
				flowFamilyMemo.Add(1)
				return publishFlowFamily(ps, refs, idx, rec, manager, cache.sourceBuilder(sourceKey))
			}
		}
	}
	shape, _, values := precompiledFlowKey(ps, proof, shared)
	idx := flowFamilyIndex{shape, values}
	meta, _ := json.Marshal(idx)
	rec := flowPreparationRecord{Source: string(meta), Helpers: flowPreparationHelpers(ps, refs, proof, shared), Reports: precompiledFlowReports(ps)}
	if precompiledFlows.entries[shape] != nil && precompiledFlows.version == flowFamilyCompilerVersion {
		cache.put(rawKey, rec)
		return publishFlowFamily(ps, refs, idx, rec, manager, nil)
	}
	// The canonical key includes optimized nodes, proof bounds, gas, stack
	// layout and helper aliasing. Only terminal PCs and unused output constants
	// differ. Lower a shape once; every binding retains its own reports/helpers.
	sourceKey := []byte(familyPrefix() + "/source/" + shape)
	if cache != nil {
		if _, err := cache.sourceBuilder(sourceKey)(); err == nil {
			flowFamilySourceReused.Add(1)
			cache.put(rawKey, rec)
			return publishFlowFamily(ps, refs, idx, rec, manager, cache.sourceBuilder(sourceKey))
		}
	}
	if !lazyFlowShareable(ps, proof) {
		// Preserve historical rejection/fallback for shapes whose generated
		// statement tree cannot be proved shareable without lowering it.
		flowFamilySourceBuilt.Add(1)
		body, _, err := emitPreparedFlowSharedRaw(ps, proof)
		if err == nil {
			body, err = parameterizeFlowSource(body, values, len(ps))
		}
		if err != nil {
			cache.put(rawKey, flowPreparationRecord{Error: err.Error()})
			return nil
		}
		builder := prepare.Builder(func() (string, error) { return body, nil })
		if cache.put(sourceKey, flowPreparationRecord{Source: body}) {
			cache.put(rawKey, rec)
			builder = cache.sourceBuilder(sourceKey)
		}
		return publishFlowFamily(ps, refs, idx, rec, manager, builder)
	}
	// Lower only a selected family. Candidate discovery can produce overlapping
	// alternatives which never become useful. The manager deduplicates the
	// shared family before invoking this builder, so they retain graphs, not
	// generated source strings, until admitted.
	// Retain only original paths already owned by the plan, not both optimized
	// graphs and their big-integer interval objects while waiting in the queue.
	originals := make([]*hotRegion, len(ps))
	for i, p := range ps {
		originals[i] = p.original
	}
	builder := prepare.Builder(func() (string, error) {
		if cache != nil {
			if body, err := cache.sourceBuilder(sourceKey)(); err == nil {
				flowFamilySourceReused.Add(1)
				return body, nil
			}
		}
		flowFamilySourceBuilt.Add(1)
		body, _, err := emitPreparedFlowSharedRaw(prepareFlows(originals), proof)
		if err == nil {
			body, err = parameterizeFlowSource(body, values, len(originals))
		}
		if err != nil {
			return "", err
		}
		cache.put(sourceKey, flowPreparationRecord{Source: body})
		return body, nil
	})
	cache.put(rawKey, rec)
	return publishFlowFamily(ps, refs, idx, rec, manager, builder)
}

func publishFlowFamily(ps []*preparedFlow, refs []*arithmeticRegion, idx flowFamilyIndex, rec flowPreparationRecord, m *prepare.Manager, build prepare.Builder) *demandPlan {
	if entry := precompiledFlows.entries[idx.Shape]; entry != nil && precompiledFlows.version == flowFamilyCompilerVersion {
		flowFamilyStatic.Add(1)
		return instantiatePrecompiledFlow(ps, refs, idx.Values, entry, rec)
	}
	var exits [4]uint64
	p := &demandPlan{branch: true, shared: true, flow: rec.Reports, helpers: rec.bindHelpers(refs), forward: make([]int, len(ps)), dead: make([]int, len(ps))}
	for i, x := range ps {
		p.regions = append(p.regions, x.original)
		exits[i] = x.original.Exit
	}
	p.slot = m.FactorySlot(familyPrefix()+"/native/"+idx.Shape, build, idx.Values, exits)
	flowFamilyBound.Add(1)
	return p.finishBinding()
}

// Constants used by arithmetic, guards or memory addresses never appear in
// values. Replacing only these output constants preserves all proved bounds.
func parameterizeFlowSource(source string, values []uint256.Int, paths int) (string, error) {
	if len(flowFamilyExitRE.FindAllStringSubmatch(source, -1)) != paths {
		return "", fmt.Errorf("family missing exits")
	}
	source = flowFamilyExitRE.ReplaceAllString(source, "${1}s.PC=exits[${2}]")
	ids := map[uint256.Int]int{}
	for i, v := range values {
		ids[v] = i
	}
	source = flowFamilyLiteralRE.ReplaceAllStringFunc(source, func(lit string) string {
		parts := flowFamilyLiteralRE.FindStringSubmatch(lit)
		var v uint256.Int
		for i := range v {
			v[i], _ = strconv.ParseUint(parts[i+1], 10, 64)
		}
		if id, ok := ids[v]; ok {
			return fmt.Sprintf("values[%d]", id)
		}
		return lit
	})
	if !flowFamilySignatureRE.MatchString(source) {
		return "", fmt.Errorf("family missing signature")
	}
	check := ""
	if len(values) > 0 {
		check = fmt.Sprintf("_ = values[%d]\n", len(values)-1)
	}
	source = flowFamilySignatureRE.ReplaceAllString(source, "func Exec(values []uint256.Int,exits [4]uint64) regionexec.Executor {return func(s *regionexec.Frame)bool {\n"+check)
	return strings.TrimSpace(source) + "\n}\n", nil
}
