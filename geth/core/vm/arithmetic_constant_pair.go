package vm

// Binding proves two full-width operands, no preserved prefix, and a valid
// constant continuation. Original stack/gas admission and exact native checks
// remain; exceptional paths use the existing recovery executor.
func (f *arithmeticFrame) trySyncConstantPair(pc *uint64, scope *ScopeContext, r *arithmeticRegion) (bool, *arithmeticResult) {
	stack, contract := scope.Stack, scope.Contract
	n := len(stack.data)
	if n < 2 || n > r.syncHot.maxStack || contract.Gas < r.syncHot.admissionGas {
		f.fallbacks++
		return false, nil
	}
	a, b := &stack.data[n-2], &stack.data[n-1]
	if r.syncHot.constantPairLayout == 2 {
		a, b = b, a
	}
	index, result := r.syncHot.native(*a, *b)
	cost := r.syncHot.costs[index]
	if cost == syncHotFallback || contract.Gas < cost {
		return f.trySynchronous(pc, scope, r)
	}
	stack.data[n-2] = result
	stack.data = stack.data[:n-1]
	contract.Gas -= cost
	*pc = r.k.constant.Uint64()
	f.sequence++
	f.checked++
	f.nativeChecked++
	f.constantPairs++
	f.recordPrune(r)
	return true, nil
}

// The operands are the first and last consumed slots (in either order); all
// slots except the last are preserved in place. Only the top result is written.
func (f *arithmeticFrame) trySyncPreservedPair(pc *uint64, scope *ScopeContext, r *arithmeticRegion) (bool, *arithmeticResult) {
	stack, contract := scope.Stack, scope.Contract
	n := len(stack.data)
	if n < r.inputCount || n > r.syncHot.maxStack || contract.Gas < r.syncHot.admissionGas {
		f.fallbacks++
		return false, nil
	}
	a, b := &stack.data[n-r.inputCount], &stack.data[n-1]
	if r.syncHot.constantPairLayout == 4 {
		a, b = b, a
	}
	index, result := r.syncHot.native(*a, *b)
	cost := r.syncHot.costs[index]
	if cost == syncHotFallback || contract.Gas < cost {
		return f.trySynchronous(pc, scope, r)
	}
	stack.data[n-1] = result
	contract.Gas -= cost
	*pc = r.k.constant.Uint64()
	f.sequence++
	f.checked++
	f.nativeChecked++
	f.constantPairs++
	f.recordPrune(r)
	return true, nil
}
