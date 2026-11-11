package vm

import "github.com/holiman/uint256"

// Immutable after binding. Source kinds and common output permutations are
// selected once, not interpreted separately for every source on every hit.
type arithmeticSyncLayout struct {
	ready   bool
	sources uint8 // bits 0/1/2: a/b/continuation is a constant
	output  uint8
}

const (
	syncOutputGeneric uint8 = iota
	syncOutputIdentity
	syncOutputOneSlot
	syncOutputOneConstant
	syncOutputTwoSlots
)

func (r *arithmeticRegion) prepareSyncLayout() {
	l := arithmeticSyncLayout{ready: true}
	if r.a.index < 0 {
		l.sources |= 1
	}
	if r.b.index < 0 {
		l.sources |= 2
	}
	if r.k.index < 0 {
		l.sources |= 4
	}
	switch {
	case r.prefixIdentity:
		l.output = syncOutputIdentity
	case len(r.prefix) == 1 && r.prefix[0].index >= 0:
		l.output = syncOutputOneSlot
	case len(r.prefix) == 1:
		l.output = syncOutputOneConstant
	case len(r.prefix) == 2 && r.prefix[0].index >= 0 && r.prefix[1].index >= 0:
		l.output = syncOutputTwoSlots
	}
	r.syncLayout = l
}

// trySync fuses native evaluation and the successful commit. It does not queue
// work or create task/result envelopes. Before the commit, the live stack, gas
// and PC are untouched, so uncommon cases can safely use the reference path.
func (f *arithmeticFrame) trySync(pc *uint64, scope *ScopeContext, r *arithmeticRegion) (bool, *arithmeticResult) {
	if !r.syncLayout.ready || r.plan.native == nil || f.config.GasPolicy != OriginalArithmeticGas {
		return f.try(pc, scope, r)
	}
	stack, contract := scope.Stack, scope.Contract
	n := len(stack.data)
	if n < r.inputCount || n-r.inputCount+r.peak > 1024 {
		f.fallbacks++
		return false, nil
	}
	input := stack.data[n-r.inputCount:]
	var a, b, k *uint256.Int
	switch r.syncLayout.sources {
	case 0:
		a, b, k = &input[r.a.index], &input[r.b.index], &input[r.k.index]
	case 1:
		a, b, k = &r.a.constant, &input[r.b.index], &input[r.k.index]
	case 2:
		a, b, k = &input[r.a.index], &r.b.constant, &input[r.k.index]
	case 3:
		a, b, k = &r.a.constant, &r.b.constant, &input[r.k.index]
	case 4:
		a, b, k = &input[r.a.index], &input[r.b.index], &r.k.constant
	case 5:
		a, b, k = &r.a.constant, &input[r.b.index], &r.k.constant
	case 6:
		a, b, k = &input[r.a.index], &r.b.constant, &r.k.constant
	case 7:
		a, b, k = &r.a.constant, &r.b.constant, &r.k.constant
	}
	if (r.domain.Bits != 256 && (!r.domain.canonical(*a) || !r.domain.canonical(*b))) || (!r.constantContinuation && !contract.validJumpdest(k)) {
		f.fallbacks++
		return false, nil
	}
	// Preserve the reference admission rule. Low gas must execute the original
	// path: its failure branch can cost less than the successful branch.
	if contract.Gas < r.successGas+r.extraGas {
		f.fallbacks++
		return false, nil
	}
	index, result := r.plan.native(*a, *b)
	p := &r.paths[index]
	if p.originalFailure || p.path.Outcome != "return" || p.path.MemoryBytes != 0 {
		return f.try(pc, scope, r)
	}
	cost := p.gas + r.extraGas
	if contract.Gas < cost {
		return f.try(pc, scope, r)
	}
	// Save the continuation before any in-place write can overwrite its source.
	continuation := k.Uint64()
	base := n - r.inputCount
	switch r.syncLayout.output {
	case syncOutputIdentity:
		end := base + len(r.prefix)
		if end < n {
			stack.data[end] = result
			stack.data = stack.data[:end+1]
		} else {
			stack.data = append(stack.data[:end], result)
		}
	case syncOutputOneSlot:
		value := input[r.prefix[0].index]
		stack.data = append(stack.data[:base], value, result)
	case syncOutputOneConstant:
		stack.data = append(stack.data[:base], r.prefix[0].constant, result)
	case syncOutputTwoSlots:
		// Snapshot both reads before writing, including swaps and repeated slots.
		x, y := input[r.prefix[0].index], input[r.prefix[1].index]
		stack.data = append(stack.data[:base], x, y, result)
	default:
		restoreArithmeticPrefix(stack, input, r)
		stack.push(&result)
	}
	contract.Gas -= cost
	*pc = continuation
	f.sequence++
	f.checked++
	f.nativeChecked++
	f.recordPrune(r)
	return true, nil
}
