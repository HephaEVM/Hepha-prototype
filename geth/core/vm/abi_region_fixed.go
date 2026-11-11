package vm

import "github.com/holiman/uint256"

// Only contiguous calldata words with identity output order qualify. Padded,
// wrapped or memory reads retain the original exact native/general executor.
type abiFixedDecode struct {
	base   int
	offset uint64
	count  int
	clean  [4]abiClean
}

func (n *abiNativeRegion) prepareFixed() {
	if n.kind != 1 || len(n.reads) == 0 || len(n.reads) > 4 || len(n.readOutputs) != len(n.reads) {
		return
	}
	r := n.reads[0]
	if r.memory || !r.offset.add.IsUint64() {
		return
	}
	f := &abiFixedDecode{base: r.offset.input, offset: r.offset.add.Uint64(), count: len(n.reads)}
	if f.offset > ^uint64(0)-128 {
		return
	}
	for i, r := range n.reads {
		if r.memory || r.offset.input != f.base || !r.offset.add.IsUint64() || r.offset.add.Uint64() != f.offset+uint64(i*32) || n.readOutputs[i] != i {
			return
		}
		f.clean[i] = r.clean
	}
	n.fixed = f
}

func (e *ABIScalarDiagnostic) runABIFixed(n *abiNativeRegion, prefix int, s *ScopeContext) (uint64, bool) {
	f := n.fixed
	if prefix+n.peak > 1024 || s.Contract.Gas < n.cost {
		return 0, false
	}
	input := s.Stack.data[prefix:]
	offset := f.offset
	if f.base >= 0 {
		v, ov := input[f.base].Uint64WithOverflow()
		if ov || v > ^uint64(0)-offset {
			return 0, false
		}
		offset += v
	}
	length := uint64(f.count * 32)
	// Preserve the repository-specific selector tracking through the old path.
	if offset == 0 || offset > uint64(len(s.Contract.Input)) || length > uint64(len(s.Contract.Input))-offset {
		return 0, false
	}
	for _, b := range n.bounds {
		var span uint256.Int
		span.Sub(&input[b.end], &input[b.start])
		if span.Slt(&b.size) {
			return 0, false
		}
	}
	var words [4]uint256.Int
	data := s.Contract.Input[offset : offset+length]
	for i := 0; i < f.count; i++ {
		words[i].SetBytes(data[i*32 : (i+1)*32])
		clean := words[i]
		f.clean[i].apply(&clean)
		if clean != words[i] {
			return 0, false
		}
	}
	// No memory operation, resizing or memory-gas computation is needed.
	s.Stack.data = append(s.Stack.data[:prefix], words[:f.count]...)
	s.Contract.Gas -= n.cost
	e.Stats.Gas += n.cost
	return n.instructions, true
}
