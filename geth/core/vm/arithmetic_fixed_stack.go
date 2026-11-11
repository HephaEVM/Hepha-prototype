package vm

// Selected only after binding proves [continuation,a,b] or [continuation,b,a]
// -> [result], full-width words, and an exact native plan. Runtime checks still
// enforce stack bounds, the continuation and original gas. Other layouts retain
// the general path; no assumption is inferred from a few observed inputs.
func (f *arithmeticFrame) trySyncFixed(pc *uint64, scope *ScopeContext, r *arithmeticRegion) (bool, *arithmeticResult) {
	stack, contract := scope.Stack, scope.Contract
	n := len(stack.data)
	if n < 3 || n > r.syncHot.maxStack {
		f.fallbacks++
		return false, nil
	}
	k := &stack.data[n-3]
	if !contract.validJumpdest(k) || contract.Gas < r.syncHot.admissionGas {
		f.fallbacks++
		return false, nil
	}
	a, b := &stack.data[n-2], &stack.data[n-1]
	if r.syncHot.fixedLayout == 2 {
		a, b = b, a
	}
	index, result := r.syncHot.native(*a, *b)
	cost := r.syncHot.costs[index]
	if cost == syncHotFallback || contract.Gas < cost {
		return f.trySynchronous(pc, scope, r)
	}
	continuation := k.Uint64()
	stack.data[n-3] = result
	stack.data = stack.data[:n-2]
	contract.Gas -= cost
	*pc = continuation
	f.sequence++
	f.checked++
	f.nativeChecked++
	f.fixedPrunes++
	f.recordPrune(r)
	return true, nil
}
