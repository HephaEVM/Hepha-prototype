package vm

// Immutable recipe derived from a proved native encode plan. These offsets
// share a base; offsets, cleaning and write order remain exactly as proved.
type abiFixedEncode struct {
	base         int
	offset, size uint64
}

func (n *abiNativeRegion) prepareFixedEncode() {
	if n.kind != 2 || len(n.writes) == 0 || len(n.writes) > 18 {
		return
	}
	w := n.writes[0]
	if !w.offset.add.IsUint64() {
		return
	}
	f := &abiFixedEncode{w.offset.input, w.offset.add.Uint64(), uint64(len(n.writes) * 32)}
	if f.offset > abiMaxMemory-f.size {
		return
	}
	for i, w := range n.writes {
		if w.offset.input != f.base || !w.offset.add.IsUint64() || w.offset.add.Uint64() != f.offset+uint64(i*32) {
			return
		}
	}
	n.encode = f
}

// A conservative no-wrap fast path. A wrapping base can be valid EVM input;
// reject without mutation so the old native executor handles that case.
func (e *ABIScalarDiagnostic) runABIFixedEncode(n *abiNativeRegion, prefix int, s *ScopeContext) (uint64, bool) {
	f := n.encode
	if prefix+n.peak > 1024 || s.Contract.Gas < n.cost {
		return 0, false
	}
	input := s.Stack.data[prefix:]
	offset := f.offset
	if f.base >= 0 {
		v, ov := input[f.base].Uint64WithOverflow()
		if ov || v > abiMaxMemory-f.size-offset {
			return 0, false
		}
		offset += v
	}
	size, fee, cost, ok := abiNativeMemory(s.Memory, offset+f.size, n.cost, s.Contract.Gas)
	if !ok {
		return 0, false
	}
	// Capture outputs before overwriting the input stack; writes read it first.
	values := e.scratch.values[:]
	for i := range n.outputs {
		values[i] = n.outputs[i].eval(input)
	}
	s.Memory.Resize(size)
	s.Memory.lastGasCost = fee
	for i := range n.writes {
		w := &n.writes[i]
		v := input[w.input]
		w.clean.apply(&v)
		s.Memory.Set32(offset+uint64(i*32), &v)
	}
	s.Stack.data = append(s.Stack.data[:prefix], values[:len(n.outputs)]...)
	s.Contract.Gas -= cost
	e.Stats.Gas += cost
	return n.instructions, true
}
