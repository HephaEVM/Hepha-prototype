package vm

import (
	"github.com/ethereum/go-ethereum/params"
	"github.com/holiman/uint256"
	"reflect"
)

// A bounded scalar form: (source & mask) >> shift. Comparisons are separate
// nodes, evaluated once at a native-region entry, not per EVM instruction.
type summaryWord struct {
	kind       uint8
	depth      int
	word, mask uint256.Int
	shift      uint
	invert     bool
}

const (
	sumConstant uint8 = iota
	sumInput
	sumCaller
	sumKey
	sumLoaded
	sumResult
	sumPredicate
)

type summaryCompare struct {
	a, b summaryWord
	op   OpCode
}
type summaryWrite struct {
	offset int
	v      summaryWord
}
type summaryRegion struct {
	entry, exit, prefixExit, gas, prefixGas uint64
	minStack, peak, delta, prefixDelta      int
	words                                   [2]summaryWord
	equalWords                              [2]summaryWord
	compareWords                            bool
	comparisons                             []summaryCompare
	guards                                  []summaryWord
	math                                    *arithmeticRegion
	a, b                                    summaryWord
	writes, prefixWrites                    []summaryWrite
	pcs                                     []uint64
}

func sumConst(v uint64) summaryWord { return summaryWord{word: *uint256.NewInt(v)} }
func sumLeaf(k uint8, d int) summaryWord {
	v := summaryWord{kind: k, depth: d}
	v.mask.SetAllOne()
	return v
}
func (v summaryWord) eval(in []uint256.Int, s *ScopeContext, key, loaded, result uint256.Int, pred *[8]bool) uint256.Int {
	w := v.word
	switch v.kind {
	case sumInput:
		w = in[len(in)-v.depth]
	case sumCaller:
		w.SetBytes(effectiveCaller(s))
	case sumKey:
		w = key
	case sumLoaded:
		w = loaded
	case sumResult:
		w = result
	case sumPredicate:
		w.Clear()
		if pred[v.depth] != v.invert {
			w.SetOne()
		}
		return w
	}
	if v.kind != sumConstant {
		w.And(&w, &v.mask)
		if v.shift != 0 {
			w.Rsh(&w, v.shift)
		}
	}
	return w
}
func sumWrites(st []summaryWord) []summaryWrite {
	var ws []summaryWrite
	for i, v := range st {
		if v == sumLeaf(sumInput, 32-i) {
			continue
		}
		ws = append(ws, summaryWrite{i - 32, v})
	}
	return ws
}
func summaryCandidateEntries(code []byte) []uint64 {
	var out []uint64
	for pc := uint64(0); pc < uint64(len(code)); pc++ {
		op := OpCode(code[pc])
		if op >= PUSH1 && op <= PUSH4 {
			n := uint64(op - PUSH0)
			if pc+n >= uint64(len(code)) {
				break
			}
			var w uint256.Int
			w.SetBytes(code[pc+1 : pc+1+n])
			k := w.Uint64()
			if k > pc && k <= pc+96 && k < uint64(len(code)) && OpCode(code[k]) == JUMPDEST {
				out = append(out, pc)
			}
			pc += n
		} else if op >= PUSH1 && op <= PUSH32 {
			pc += uint64(op - PUSH0)
		}
	}
	return out
}
func bindSummaryRegion(code []byte, entry uint64, b *boundArithmetic, table *JumpTable) *summaryRegion {
	if b == nil || table[SLOAD].dynamicGas == nil || table[SLOAD].constantGas != 0 || reflect.ValueOf(table[SLOAD].dynamicGas).Pointer() != reflect.ValueOf(gasSLoadEIP2929).Pointer() {
		return nil
	}
	bitmap := codeBitmap(code)
	p := &summaryRegion{entry: entry}
	st := make([]summaryWord, 32)
	for i := range st {
		st[i] = sumLeaf(sumInput, 32-i)
	}
	need := func(n int) bool {
		if len(st) < n {
			return false
		}
		if d := 32 - len(st) + n; d > p.minStack {
			p.minStack = d
		}
		return p.minStack <= 32
	}
	pop := func() summaryWord { v := st[len(st)-1]; st = st[:len(st)-1]; return v }
	push := func(v summaryWord) {
		st = append(st, v)
		if len(st)-32 > p.peak {
			p.peak = len(st) - 32
		}
	}
	var mem [2]summaryWord
	var known [2]bool
	hashes, loads := 0, 0
	pc := entry
	visits := map[uint64]int{}
	compare := func(a, b summaryWord, op OpCode) (summaryWord, bool) {
		if len(p.comparisons) >= 8 || a.kind == sumPredicate || b.kind == sumPredicate {
			return summaryWord{}, false
		}
		p.comparisons = append(p.comparisons, summaryCompare{a, b, op})
		return summaryWord{kind: sumPredicate, depth: len(p.comparisons) - 1}, true
	}
	for step := 0; step < 192 && pc < uint64(len(code)); step++ {
		visits[pc]++
		if visits[pc] > 2 {
			return nil
		}
		p.pcs = append(p.pcs, pc)
		op := OpCode(code[pc])
		o := table[op]
		if o == nil || o.undefined || !need(o.minStack) || o.dynamicGas != nil && op != KECCAK256 && op != MSTORE && op != SLOAD {
			return nil
		}
		p.gas += o.constantGas
		switch {
		case op >= PUSH0 && op <= PUSH32:
			n := uint64(op - PUSH0)
			if pc+n >= uint64(len(code)) {
				return nil
			}
			var w uint256.Int
			w.SetBytes(code[pc+1 : pc+1+n])
			push(summaryWord{word: w})
			pc += n
		case op >= DUP1 && op <= DUP16:
			push(st[len(st)-int(op-DUP1)-1])
		case op >= SWAP1 && op <= SWAP16:
			j := len(st) - int(op-SWAP1) - 2
			st[j], st[len(st)-1] = st[len(st)-1], st[j]
		case op == POP:
			pop()
		case op == JUMPDEST:
		case op == CALLER:
			push(sumLeaf(sumCaller, 0))
		case op == AND:
			a, c := pop(), pop()
			if a.kind == sumConstant {
				a, c = c, a
			}
			if c.kind != sumConstant || a.kind == sumPredicate {
				return nil
			}
			if a.kind == sumConstant {
				a.word.And(&a.word, &c.word)
			} else {
				mask := c.word
				mask.Lsh(&mask, a.shift)
				a.mask.And(&a.mask, &mask)
			}
			push(a)
		case op == SHR:
			amount, v := pop(), pop()
			if amount.kind != sumConstant || !amount.word.LtUint64(256) || v.kind == sumPredicate {
				return nil
			}
			if v.kind == sumConstant {
				v.word.Rsh(&v.word, uint(amount.word.Uint64()))
			} else {
				v.shift += uint(amount.word.Uint64())
				if v.shift >= 256 {
					v = sumConst(0)
				}
			}
			push(v)
		case op == ADD || op == SUB || op == SHL:
			a, c := pop(), pop()
			if a.kind != sumConstant || c.kind != sumConstant {
				return nil
			}
			var w uint256.Int
			switch op {
			case ADD:
				w.Add(&a.word, &c.word)
			case SUB:
				w.Sub(&a.word, &c.word)
			case SHL:
				if a.word.LtUint64(256) {
					w.Lsh(&c.word, uint(a.word.Uint64()))
				}
			}
			push(summaryWord{word: w})
		case op == EQ || op == LT || op == GT:
			if loads != 1 {
				return nil
			}
			a, c := pop(), pop()
			v, ok := compare(a, c, op)
			if !ok {
				return nil
			}
			push(v)
		case op == ISZERO:
			if loads != 1 {
				return nil
			}
			v := pop()
			if v.kind == sumPredicate {
				v.invert = !v.invert
			} else {
				var ok bool
				v, ok = compare(v, sumConst(0), EQ)
				if !ok {
					return nil
				}
			}
			push(v)
		case op == MSTORE:
			if hashes == 1 && loads == 0 {
				return nil
			}
			off, v := pop(), pop()
			if off.kind != sumConstant || !off.word.IsUint64() || (off.word.Uint64() != 0 && off.word.Uint64() != 32) || v.kind == sumPredicate {
				return nil
			}
			i := off.word.Uint64() / 32
			mem[i] = v
			known[i] = true
		case op == KECCAK256:
			off, size := pop(), pop()
			if off != sumConst(0) || size != sumConst(64) || !known[0] || !known[1] {
				return nil
			}
			hashes++
			p.gas += 12
			if hashes == 1 {
				if loads != 0 {
					return nil
				}
				for _, v := range mem {
					if v.kind != sumConstant && v.kind != sumInput && v.kind != sumCaller {
						return nil
					}
				}
				p.words = mem
				push(sumLeaf(sumKey, 0))
			} else if hashes == 2 {
				if loads != 1 {
					return nil
				}
				for _, v := range mem {
					if v.kind != sumConstant && v.kind != sumInput && v.kind != sumCaller {
						return nil
					}
				}
				if p.words != mem {
					return nil
				}
				p.equalWords = mem
				push(sumLeaf(sumKey, 0))
				p.exit = pc + 1
				p.delta = len(st) - 32
				p.writes = sumWrites(st)
				if len(p.writes) > 32 || len(p.prefixWrites) > 32 || len(p.guards) == 0 {
					return nil
				}
				return p
			} else {
				return nil
			}
		case op == SLOAD:
			key := pop()
			if hashes != 1 || loads != 0 || key != sumLeaf(sumKey, 0) {
				return nil
			}
			loads++
			push(sumLeaf(sumLoaded, 0))
			p.prefixExit = pc + 1
			p.prefixGas = p.gas
			p.prefixDelta = len(st) - 32
			p.prefixWrites = sumWrites(st)
		case op == JUMP:
			target := pop()
			if target.kind != sumConstant || !target.word.IsUint64() {
				return nil
			}
			dest := target.word.Uint64()
			if dest >= uint64(len(code)) || (OpCode(code[dest]) != JUMPDEST || !bitmap.codeSegment(dest)) {
				return nil
			}
			if h := b.regions[dest]; h != nil && h.syncHot.native != nil && h.syncHot.fixedLayout != 0 && h.extraGas == 0 && h.constantSuccessGas && !h.pathDependentResult {
				if loads != 1 || p.math != nil || !need(3) {
					return nil
				}
				k := st[len(st)-3]
				if k.kind != sumConstant || !k.word.IsUint64() {
					return nil
				}
				ret := k.word.Uint64()
				if ret >= uint64(len(code)) || (OpCode(code[ret]) != JUMPDEST || !bitmap.codeSegment(ret)) {
					return nil
				}
				p.a, p.b = st[len(st)-2], st[len(st)-1]
				if p.a.kind == sumPredicate || p.b.kind == sumPredicate {
					return nil
				}
				if h.syncHot.fixedLayout == 2 {
					p.a, p.b = p.b, p.a
				}
				p.math = h
				p.gas += h.syncHot.admissionGas
				if peak := len(st) - 32 + h.peak - 3; peak > p.peak {
					p.peak = peak
				}
				st = st[:len(st)-3]
				push(sumLeaf(sumResult, 0))
				pc = ret - 1
			} else {
				pc = dest - 1
			}
		case op == JUMPI:
			target, v := pop(), pop()
			if loads != 1 || target.kind != sumConstant || !target.word.IsUint64() || v.kind != sumPredicate || len(p.guards) >= 4 {
				return nil
			}
			dest := target.word.Uint64()
			if dest >= uint64(len(code)) || (OpCode(code[dest]) != JUMPDEST || !bitmap.codeSegment(dest)) || dest <= pc {
				return nil
			}
			p.guards = append(p.guards, v)
			pc = dest - 1
		default:
			return nil
		}
		pc++
	}
	return nil
}
func applySummaryWrites(s *ScopeContext, in []uint256.Int, ws []summaryWrite, delta int, key, loaded, result uint256.Int, pred *[8]bool) {
	var vals [32]uint256.Int
	for i, w := range ws {
		vals[i] = w.v.eval(in, s, key, loaded, result, pred)
	}
	n := len(in)
	end := n + delta
	if cap(s.Stack.data) < end {
		s.Stack.data = append(s.Stack.data, make([]uint256.Int, end-len(s.Stack.data))...)
	} else {
		s.Stack.data = s.Stack.data[:end]
	}
	for i, w := range ws {
		s.Stack.data[n+w.offset] = vals[i]
	}
}
func (p *summaryRegion) try(pc *uint64, s *ScopeContext, e *EVM) bool {
	in := s.Stack.data
	n := len(in)
	if n < p.minStack || n+p.peak > 1024 || e.abort.Load() || e.Config.Tracer != nil || e.Config.EnablePreimageRecording {
		return false
	}
	extra, fee := uint64(0), s.Memory.lastGasCost
	if s.Memory.Len() < 64 {
		fee = 6
		if fee < s.Memory.lastGasCost {
			return false
		}
		extra = fee - s.Memory.lastGasCost
	}
	if s.Contract.Gas < p.gas+extra+params.ColdSloadCostEIP2929 {
		return false
	}
	var zero uint256.Int
	var pred [8]bool
	a, b := p.words[0].eval(in, s, zero, zero, zero, &pred), p.words[1].eval(in, s, zero, zero, zero, &pred)
	s.Memory.Resize(64)
	s.Memory.lastGasCost = fee
	s.Memory.Set32(0, &a)
	s.Memory.Set32(32, &b)
	e.hasher.Reset()
	e.hasher.Write(s.Memory.store[:64])
	e.hasher.Read(e.hasherBuf[:])
	var key uint256.Int
	key.SetBytes(e.hasherBuf[:])
	s.Contract.Gas -= p.prefixGas + extra
	s.Stack.data = append(s.Stack.data, key)
	cost, _ := gasSLoadEIP2929(e, s.Contract, s.Stack, s.Memory, 0)
	s.Contract.Gas -= cost
	opSload(pc, e, s)
	loaded := *s.Stack.peek()
	s.Stack.data = s.Stack.data[:n]
	result := zero
	ok := true
	if p.math != nil {
		a, b := p.a.eval(in, s, key, loaded, zero, &pred), p.b.eval(in, s, key, loaded, zero, &pred)
		ix, v := p.math.syncHot.native(a, b)
		result = v
		ok = p.math.syncHot.costs[ix] == p.math.syncHot.admissionGas
	}
	if ok {
		for i, c := range p.comparisons {
			a, b := c.a.eval(in, s, key, loaded, result, &pred), c.b.eval(in, s, key, loaded, result, &pred)
			switch c.op {
			case EQ:
				pred[i] = a == b
			case LT:
				pred[i] = a.Lt(&b)
			case GT:
				pred[i] = a.Gt(&b)
			}
		}
		for _, g := range p.guards {
			if pred[g.depth] == g.invert {
				ok = false
				break
			}
		}
	}
	if !ok || e.abort.Load() {
		applySummaryWrites(s, in, p.prefixWrites, p.prefixDelta, key, loaded, zero, &pred)
		*pc = p.prefixExit
		return true
	}
	if paperWithoutReuse {
		key = paperRehash64(e, s)
	}
	applySummaryWrites(s, in, p.writes, p.delta, key, loaded, result, &pred)
	s.Contract.Gas -= p.gas - p.prefixGas
	*pc = p.exit
	return true
}
