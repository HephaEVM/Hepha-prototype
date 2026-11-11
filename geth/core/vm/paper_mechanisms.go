package vm

import (
	"github.com/holiman/uint256"
	"sync/atomic"
)

var paperNoCheckStaticKnown = map[string]bool{}
var paperCheckStaticFallbacks atomic.Uint64

func paperStaticChecksAvailable(r *hotRegion, key string) bool {
	if !paperWithoutCheckElimination {
		return true
	}
	_, _, _, report := optimizeHotRegionChecks(r, true)
	// The earlier check-retaining catalog only certified helper checks. Do not
	// reuse a function that may already have folded an inline guard predicate.
	if report.GuardConstants != 0 {
		paperCheckStaticFallbacks.Add(1)
		return false
	}
	if paperNoCheckStaticKnown[key] {
		return true
	}
	if report.FusedHelpers != 0 {
		paperCheckStaticFallbacks.Add(1)
	}
	return report.FusedHelpers == 0
}

// The experiment disables transforms after structural recognition. In
// particular, disabling reuse must not invalidate the chain's profitability
// threshold or its preselected direct executor. Normal builds fold the policy
// constants away; no experimental mode is exposed by the client configuration.
func paperDisableChainReuse(p *effectChain) {
	p.repeatAt = nil
	if len(p.repeats) != 0 {
		p.repeatAt = make([][]int, len(p.actions)+1)
	}
	for _, repeat := range p.repeats {
		p.repeatAt[repeat.action] = append(p.repeatAt[repeat.action], repeat.node)
	}
	for i := range p.actions {
		a := &p.actions[i]
		if a.op == KECCAK256 && a.reuse >= 0 {
			a.result, a.reuse = a.reuse, -1
		}
	}
}

func paperRepeatChainValues(p *effectChain, action int, values *[256]uint256.Int) {
	// Binding groups occurrences once. Execution visits only computations at
	// this boundary, never rescans all occurrences in the region.
	if len(p.repeatAt) == 0 {
		return
	}
	for _, id := range p.repeatAt[action] {
		n := p.nodes[id]
		values[id] = chainPure(n.op, values[n.a], values[n.b])
	}
}

func paperRehash64(e *EVM, s *ScopeContext) (key uint256.Int) {
	e.hasher.Reset()
	e.hasher.Write(s.Memory.store[:64])
	e.hasher.Read(e.hasherBuf[:])
	key.SetBytes(e.hasherBuf[:])
	return
}

// Metadata describes system ablations, including interactions. These variants
// must not be described as an additive attribution with identical coverage.
func PaperMechanismConfiguration() map[string]any {
	return map[string]any{
		"level":                         paperAblationLevel,
		"valuePassing":                  paperAblationLevel != 2,
		"cePredicateIsolationVersion":   2,
		"extendedCEProof":               paperCEExtendedProof,
		"branchPrefixReuse":             !paperWithoutReuse,
		"computationReuse":              !paperWithoutReuse,
		"checkElimination":              !paperWithoutCheckElimination,
		"checkRetainingStaticExecutors": len(paperNoCheckStaticKnown),
		"staticProofFallbackBindings":   paperCheckStaticFallbacks.Load(),
		"interpretation":                "progressive fusion: VP, VP+CR, VP+CR+CE; shared regional execution; report coverage alongside latency",
	}
}
