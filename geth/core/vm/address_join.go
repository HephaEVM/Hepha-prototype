package vm

import (
	"github.com/ethereum/go-ethereum/params"
	"github.com/holiman/uint256"
	miner "go-eth1.16.5-evm/template_miner"
)

type addressJoinPlan struct {
	summary            *addressSlicePlan
	preparation        *hashPreparationPlan
	fallback           []addressWrite
	fallbackDelta      int
	first              [2]addressValue
	guardMemory        bool
	prefixGas, tailGas uint64
}

func applyAddressMask(v addressValue, mask uint256.Int) addressValue {
	if v.kind == addressConstant {
		v.value.And(&v.value, &mask)
	} else {
		v.mask.And(&v.mask, &mask)
	}
	return v
}
func importedHashValue(v miner.HashSource, inputs int) (addressValue, bool) {
	var x addressValue
	switch v.Kind {
	case miner.HashInput:
		x = addressLeaf(addressInput, inputs-v.Slot)
	case miner.HashConstant:
		return addressValue{value: v.Word}, true
	case miner.HashCaller:
		x = addressLeaf(addressCaller, 0)
	case miner.HashResult:
		x = addressLeaf(addressKey, 0)
	default:
		return x, false
	}
	return applyAddressMask(x, v.Mask), true
}
func compactAddressWrites(st []addressValue, base int) []addressWrite {
	var out []addressWrite
	for i, v := range st {
		if v.kind == addressInput && v.depth == base-i && v.mask == addressLeaf(addressInput, 0).mask {
			continue
		}
		out = append(out, addressWrite{i - base, v})
	}
	return out
}
func bindAddressJoin(prep *hashPreparationPlan, code []byte, b *boundArithmetic, table *JumpTable) *addressJoinPlan {
	if prep == nil || prep.mode != HashPreparationDirect {
		return nil
	}
	r := prep.region
	if r.Offset != 0 || r.Size != 64 || r.MemoryEnd != 64 || len(r.Writes) != 2 || len(r.PCs) == 0 || len(r.Outputs) == 0 {
		return nil
	}
	hp := r.PCs[len(r.PCs)-1]
	if OpCode(code[hp]) != KECCAK256 || r.Exit != hp+1 {
		return nil
	}
	p := bindAddressSlice(code, hp, b, table)
	if p == nil {
		return nil
	}
	const base = 32
	if r.Inputs > base {
		return nil
	}
	out := make([]addressValue, base-r.Inputs)
	for i := range out {
		out[i] = addressLeaf(addressInput, base-i)
	}
	for _, v := range r.Outputs {
		x, ok := importedHashValue(v, r.Inputs)
		if !ok {
			return nil
		}
		out = append(out, x)
	}
	if out[len(out)-1] != addressLeaf(addressKey, 0) {
		return nil
	}
	beforeHash := append([]addressValue(nil), out[:len(out)-1]...)
	beforeHash = append(beforeHash, addressConst(64), addressConst(0))
	compose := func(v addressValue) (addressValue, bool) {
		if v.kind != addressInput {
			return v, true
		}
		if v.depth < 1 || v.depth > len(beforeHash) {
			return v, false
		}
		return applyAddressMask(beforeHash[len(beforeHash)-v.depth], v.mask), true
	}
	q := *p
	var ok bool
	if q.a, ok = compose(p.a); !ok {
		return nil
	}
	if q.b, ok = compose(p.b); !ok {
		return nil
	}
	if q.k, ok = compose(p.k); !ok {
		return nil
	}
	for i := range q.memory {
		if q.memory[i], ok = compose(p.memory[i]); !ok {
			return nil
		}
	}
	var first [2]addressValue
	var wrote [2]bool
	for _, w := range r.Writes {
		if w.Offset != 0 && w.Offset != 32 {
			return nil
		}
		i := w.Offset / 32
		x, valid := importedHashValue(w.Value, r.Inputs)
		if !valid || x.kind == addressKey {
			return nil
		}
		first[i] = x
		wrote[i] = true
	}
	// Equality is proved over symbolic origins, not guessed from concrete values.
	if !wrote[0] || !wrote[1] {
		return nil
	}
	q.entry = r.Entry
	q.delta = len(beforeHash) - base + p.delta
	q.minStack = r.Inputs
	if n := p.minStack - (len(beforeHash) - base); n > q.minStack {
		q.minStack = n
	}
	q.extraPeak = r.Peak - r.Inputs
	if n := len(beforeHash) - base + p.extraPeak; n > q.extraPeak {
		q.extraPeak = n
	}
	q.fixedGas = prep.gas + p.fixedGas - 42
	final := append([]addressValue(nil), beforeHash...)
	if len(final)+p.delta > len(final) {
		final = append(final, make([]addressValue, p.delta)...)
	} else {
		final = final[:len(final)+p.delta]
	}
	for _, w := range p.writes {
		v, ok := compose(w.value)
		if !ok {
			return nil
		}
		final[len(beforeHash)+w.offset] = v
	}
	q.writes = compactAddressWrites(final, base)
	if len(q.writes) > 32 {
		return nil
	}
	fall := append([]addressValue(nil), out...)
	fall[len(fall)-1] = addressLeaf(addressLoaded, 0)
	return &addressJoinPlan{first: first, guardMemory: first != q.memory, summary: &q, preparation: prep, fallback: compactAddressWrites(fall, base), fallbackDelta: len(fall) - base, prefixGas: prep.gas, tailGas: p.fixedGas - 42}
}
func writeAddressValues(s *ScopeContext, in []uint256.Int, writes []addressWrite, delta int, loaded, result, key uint256.Int) {
	var values [32]uint256.Int
	for i, w := range writes {
		values[i] = w.value.eval(in, s, loaded, result, key)
	}
	n := len(in)
	end := n + delta
	if cap(s.Stack.data) < end {
		s.Stack.data = append(s.Stack.data, make([]uint256.Int, end-len(s.Stack.data))...)
	} else {
		s.Stack.data = s.Stack.data[:end]
	}
	for i, w := range writes {
		s.Stack.data[n+w.offset] = values[i]
	}
}
func (j *addressJoinPlan) try(pc *uint64, s *ScopeContext, e *EVM) bool {
	p := j.summary
	in := s.Stack.data
	n := len(in)
	if n < p.minStack || n+p.extraPeak > 1024 || e.abort.Load() || e.Config.Tracer != nil || e.Config.EnablePreimageRecording {
		return false
	}
	fee := s.Memory.lastGasCost
	extra := uint64(0)
	if s.Memory.Len() < 64 {
		fee = 2 * params.MemoryGas
		if fee < s.Memory.lastGasCost {
			return false
		}
		extra = fee - s.Memory.lastGasCost
	}
	if s.Contract.Gas < p.fixedGas+extra+params.ColdSloadCostEIP2929 {
		return false
	}
	zero := uint256.Int{}
	k := p.k.eval(in, s, zero, zero, zero)
	if !k.IsUint64() || k.Uint64() != p.continuation {
		return false
	}
	a0, a1 := j.first[0].eval(in, s, zero, zero, zero), j.first[1].eval(in, s, zero, zero, zero)
	if j.guardMemory && (a0 != p.memory[0].eval(in, s, zero, zero, zero) || a1 != p.memory[1].eval(in, s, zero, zero, zero)) {
		return false
	}
	s.Memory.Resize(64)
	s.Memory.lastGasCost = fee
	s.Memory.Set32(0, &a0)
	s.Memory.Set32(32, &a1)
	e.hasher.Reset()
	e.hasher.Write(s.Memory.store[:64])
	e.hasher.Read(e.hasherBuf[:])
	var key uint256.Int
	key.SetBytes(e.hasherBuf[:])
	s.Contract.Gas -= j.prefixGas + extra
	// One temporary word lets the original EIP-2929 gas and SLOAD operations
	// run without materializing the entire intermediate preparation stack.
	s.Stack.data = append(s.Stack.data, key)
	cost, _ := gasSLoadEIP2929(e, s.Contract, s.Stack, s.Memory, 0)
	s.Contract.Gas -= cost
	opSload(pc, e, s)
	loaded := *s.Stack.peek()
	s.Stack.data = s.Stack.data[:n]
	a, b := p.a.eval(in, s, loaded, zero, key), p.b.eval(in, s, loaded, zero, key)
	ix, result := p.helper.syncHot.native(a, b)
	if p.helper.syncHot.costs[ix] != p.helper.syncHot.admissionGas || e.abort.Load() {
		writeAddressValues(s, in, j.fallback, j.fallbackDelta, loaded, zero, key)
		*pc = j.preparation.region.Exit + 1
		return true
	}
	if paperWithoutReuse {
		key = paperRehash64(e, s)
	}
	writeAddressValues(s, in, p.writes, p.delta, loaded, result, key)
	s.Contract.Gas -= j.tailGas
	*pc = p.exit
	return true
}
