package vm

import "github.com/holiman/uint256"

// Compact runtime data computed from the full bytecode proof at binding time.
// The original paths stay available for diagnostics and exact exception handling.
type arithmeticSyncHot struct {
	native             arithmeticNativeFunc
	guard              arithmeticGuardFunc
	costs              []uint64
	admissionGas       uint64
	maxStack, drop     int
	fixedLayout        uint8 // 1: [continuation,a,b], 2: [continuation,b,a]; both return [result]
	constantPairLayout uint8 // 1/2: consume pair; 3/4: preserve identity prefix, replace top
}

const syncHotFallback = ^uint64(0)

func (r *arithmeticRegion) prepareSyncHot() {
	h := arithmeticSyncHot{
		native:       r.plan.native,
		guard:        r.plan.guardHot,
		admissionGas: r.successGas + r.extraGas,
		maxStack:     1024 + r.inputCount - r.peak,
		drop:         r.inputCount - len(r.prefix),
	}
	if r.plan.nativeHot != nil {
		h.native = r.plan.nativeHot
	}
	if h.native != nil {
		if r.inputCount == 2 && len(r.prefix) == 0 && r.domain.Bits == 256 && r.constantContinuation && r.k.index < 0 {
			if r.a.index == 0 && r.b.index == 1 {
				h.constantPairLayout = 1
			} else if r.a.index == 1 && r.b.index == 0 {
				h.constantPairLayout = 2
			}
		}
		if r.inputCount >= 2 && len(r.prefix) == r.inputCount-1 && r.prefixIdentity && r.domain.Bits == 256 && r.constantContinuation && r.k.index < 0 {
			if r.a.index == 0 && r.b.index == r.inputCount-1 {
				h.constantPairLayout = 3
			} else if r.b.index == 0 && r.a.index == r.inputCount-1 {
				h.constantPairLayout = 4
			}
		}
		if r.inputCount == 3 && len(r.prefix) == 0 && r.domain.Bits == 256 && !r.constantContinuation && r.k.index == 0 {
			if r.a.index == 1 && r.b.index == 2 {
				h.fixedLayout = 1
			} else if r.a.index == 2 && r.b.index == 1 {
				h.fixedLayout = 2
			}
		}
		if h.guard == nil {
			// A future native plan can be registered without a guard-only form.
			// Retain its complete exact implementation in that case.
			native := h.native
			h.guard = func(a, b uint256.Int) int { index, _ := native(a, b); return index }
		}
		h.costs = make([]uint64, len(r.paths))
		for i, p := range r.paths {
			h.costs[i] = syncHotFallback
			if !p.originalFailure && p.path.Outcome == "return" && p.path.MemoryBytes == 0 {
				h.costs[i] = p.gas + r.extraGas
			}
		}
	}
	r.syncHot = h
}

// Called only for synchronous, original-gas execution without the optional
// diagnostic oracle. Stack, type, continuation and gas checks remain dynamic.
func (f *arithmeticFrame) trySyncHot(pc *uint64, scope *ScopeContext, r *arithmeticRegion) (bool, *arithmeticResult) {
	if r.syncHot.native == nil {
		return f.trySynchronous(pc, scope, r)
	}
	stack, contract := scope.Stack, scope.Contract
	n := len(stack.data)
	if n < r.inputCount || n > r.syncHot.maxStack {
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
	if contract.Gas < r.syncHot.admissionGas {
		f.fallbacks++
		return false, nil
	}
	index, result := r.syncHot.native(*a, *b)
	cost := r.syncHot.costs[index]
	if cost == syncHotFallback || contract.Gas < cost {
		return f.trySynchronous(pc, scope, r)
	}
	// Save the continuation before any in-place write can overwrite its source.
	continuation := k.Uint64()
	base := n - r.inputCount
	switch r.syncLayout.output {
	case syncOutputIdentity:
		end := n - r.syncHot.drop
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
