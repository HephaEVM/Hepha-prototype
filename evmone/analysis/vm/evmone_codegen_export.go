//go:build hepha_export

package vm

// Export the existing, re-proved IR. No C++ backend-specific bytecode matcher.
import (
	"encoding/json"
	"fmt"
	"os"
	"sort"

	"github.com/ethereum/go-ethereum/crypto"
	miner "go-eth1.16.5-evm/template_miner"
)

func cppSource(s stackSource) any {
	return map[string]any{"index": s.index, "constant": s.constant.Hex()}
}

func cppArithmetic(h *arithmeticRegion) any {
	if h.plan == nil {
		return nil
	}
	nodes := []any{}
	for _, n := range h.plan.nodes {
		nodes = append(nodes, map[string]any{"op": n.op.String(), "args": []int{int(n.x), int(n.y)}, "value": n.constant.Hex()})
	}
	ds := [][4]int{}
	for _, d := range h.plan.decisions {
		ds = append(ds, [4]int{int(d.value), d.zero, d.nonzero, d.path})
	}
	costs, steps := []int64{}, []int{}
	success := false
	prefixSteps, err := hotRegionPrefixSteps(h)
	if err != nil {
		return nil
	}
	for _, p := range h.paths {
		cost := int64(-1)
		if !p.originalFailure && p.path.Outcome == "return" && p.path.MemoryBytes == 0 {
			cost, success = int64(p.gas+h.extraGas), true
		}
		costs = append(costs, cost)
		steps = append(steps, len(p.path.PCs)+prefixSteps)
	}
	if !success {
		return nil
	}
	results := []int{}
	for _, id := range h.plan.results {
		results = append(results, int(id))
	}
	return map[string]any{"bits": h.domain.Bits, "signed": h.domain.Signed, "nodes": nodes,
		"decisions": ds, "root": h.plan.root, "results": results, "costs": costs, "steps": steps}
}

func cppPath(orig *hotRegion) any {
	r, _, report := optimizeFlowRegion(orig)
	forward, dead := memoryDemand(r)
	nodes := []any{}
	for _, n := range r.nodes {
		x := map[string]any{"op": n.op.String(), "args": n.args, "value": n.value.Hex(), "nonzero": n.nonzero}
		if n.helper != nil {
			h := cppArithmetic(n.helper)
			if h == nil {
				return nil
			}
			x["helper"], x["cost"] = h, n.helper.syncHot.admissionGas
		}
		nodes = append(nodes, x)
	}
	return map[string]any{"inputs": r.Inputs, "max_stack": r.MaxStack, "drop": r.Drop,
		"gas": r.Gas, "exit": r.Exit, "nodes": nodes, "outputs": r.outputs,
		"instructions": r.Instructions, "forward": forward, "dead": dead,
		"decodes": r.decodes, "encodes": r.encodes, "proof": report}
}

// Stage 1 adds helper summaries. Stage 2 lowers bounded, effect-free (except
// staged word memory) computation paths. Storage, calls, logs and GAS stay out.
func ExportEvmonePrograms(code []byte, libraryPath string, engine *ArithmeticEngine, sites map[uint64]bool) ([]EvmoneRegion, error) {
	exporter, err := NewEvmoneProgramExporter(libraryPath, engine)
	if err != nil {
		return nil, err
	}
	return exporter.Regions(code, sites)
}

// Load and validate the ABI library once per export, not once per code.
type EvmoneProgramExporter struct {
	library miner.ABIRegionLibrary
	engine  *ArithmeticEngine
}

func NewEvmoneProgramExporter(path string, engine *ArithmeticEngine) (*EvmoneProgramExporter, error) {
	raw, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	e := &EvmoneProgramExporter{engine: engine}
	if err = json.Unmarshal(raw, &e.library); err != nil {
		return nil, err
	}
	if e.library.Version != 1 {
		return nil, fmt.Errorf("unsupported ABI library")
	}
	return e, nil
}
func (e *EvmoneProgramExporter) Regions(code []byte, sites map[uint64]bool) ([]EvmoneRegion, error) {
	abi := map[uint64]*abiRegionPrepared{}
	out := exportEvmoneRegions(code, &e.library, abi)
	if sites != nil && len(sites) == 0 {
		return out, nil
	}
	engine := e.engine
	b := engine.bind(code, arithmeticCacheKey{hash: crypto.Keccak256Hash(code), rules: &pragueInstructionSet, reduced: true, compute: 1 | 4 | uint8(HashPreparationDirect)<<3})
	if b.err != nil {
		return nil, b.err
	}
	add := func(pc uint64, stage int, native any) {
		out = append(out, EvmoneRegion{Kind: 5, Entry: pc, Stage: stage, Native: native, Evidence: "existing bytecode proof / offline native lowering"})
	}
	for pc, h := range b.regions {
		if sites != nil && !sites[pc] {
			continue
		}
		v := cppArithmetic(h)
		if v == nil {
			continue
		}
		prefix := []any{}
		for _, s := range h.prefix {
			prefix = append(prefix, cppSource(s))
		}
		add(pc, 1, map[string]any{"family": "arithmetic", "helper": v, "inputs": h.inputCount,
			"peak": h.peak, "a": cppSource(h.a), "b": cppSource(h.b), "k": cppSource(h.k), "prefix": prefix})
	}
	barriers, occupied := map[uint64]bool{}, map[uint64]bool{}
	for _, r := range out {
		if r.Stage == 0 {
			occupied[r.Entry] = true
		}
		if r.Kind == 3 || r.Kind == 4 {
			barriers[r.Entry] = true
		}
	}
	for pc, index := range b.entryIndex {
		if sites != nil && !sites[uint64(pc)] {
			continue
		}
		if index == 0 || int(index) >= len(b.computeEntries) {
			continue
		}
		r := b.computeEntries[index]
		if r == nil {
			continue
		}
		if r.preparation != nil {
			add(uint64(pc), 1, map[string]any{"family": "hash", "region": r.preparation.region, "gas": r.preparation.gas})
			barriers[uint64(pc)] = true
		} else if r.muldiv != nil {
			paths := []computeHelperPath{r.short, r.wide}
			if r.rounding != nil {
				paths = r.rounded[:]
			}
			costs, peaks, steps := []uint64{}, []int{}, []uint64{}
			for _, p := range paths {
				costs = append(costs, p.gas)
				peaks = append(peaks, p.peak)
				steps = append(steps, p.instructions)
			}
			a, c, d := r.muldiv.A, r.muldiv.B, r.muldiv.D
			if r.rounding != nil {
				a, c, d = r.rounding.A, r.rounding.B, r.rounding.D
			}
			add(uint64(pc), 1, map[string]any{"family": "muldiv", "a": a, "b": c, "d": d, "rounding": r.rounding != nil, "costs": costs, "peaks": peaks, "steps": steps})
			barriers[uint64(pc)] = true
		}
	}
	bitmap := codeBitmap(code)
	for pc := range partialStarts(code, true) {
		if sites != nil && !sites[pc] {
			continue
		}
		if occupied[pc] || barriers[pc] {
			continue
		}
		paths, seen := []any{}, map[string]bool{}
		for mask := uint(0); mask < 4; mask++ {
			r := buildFlowPathWithABI(b, code, &pragueInstructionSet, pc, mask, 2, abi, barriers, bitmap)
			if r == nil || len(r.nodes) > 160 || seen[r.key] {
				continue
			}
			seen[r.key] = true
			if p := cppPath(r); p != nil {
				paths = append(paths, p)
			}
		}
		family := "branch"
		if len(paths) == 0 {
			r := buildHotRegionWithABI(b, code, &pragueInstructionSet, pc, false, ^uint64(0), abi, barriers)
			if r == nil || len(r.nodes) > 160 {
				continue
			}
			if p := cppPath(r); p != nil {
				paths = append(paths, p)
			}
			family = "region"
			if r.decodes+r.encodes > 0 {
				family = "pipeline"
			}
		}
		if len(paths) > 0 {
			add(pc, 2, map[string]any{"family": family, "paths": paths})
		}
	}
	sort.SliceStable(out, func(i, j int) bool {
		if out[i].Entry != out[j].Entry {
			return out[i].Entry < out[j].Entry
		}
		return out[i].Stage > out[j].Stage
	})
	return out, nil
}
