package vm

import (
	"github.com/ethereum/go-ethereum/params"
	"github.com/holiman/uint256"
	"reflect"
)

const addressPredicateKind uint8 = 100

type addressPredicate struct {
	a, b   addressValue
	op     OpCode
	invert bool
}
type addressGuardPlan struct {
	summary *addressSlicePlan
	guards  []addressPredicate
}

func (g *addressGuardPlan) accept(in []uint256.Int, s *ScopeContext, loaded uint256.Int) bool {
	for _, p := range g.guards {
		a, b := p.a.eval(in, s, loaded, uint256.Int{}, uint256.Int{}), p.b.eval(in, s, loaded, uint256.Int{}, uint256.Int{})
		v := false
		switch p.op {
		case LT:
			v = a.Lt(&b)
		case GT:
			v = a.Gt(&b)
		case EQ:
			v = a == b
		}
		if p.invert {
			v = !v
		}
		if !v {
			return false
		}
	}
	return true
}
func bindAddressGuard(code []byte, entry uint64, b *boundArithmetic, table *JumpTable) *addressGuardPlan {
	if b == nil || entry+2 >= uint64(len(code)) || OpCode(code[entry]) != KECCAK256 || OpCode(code[entry+1]) != SLOAD {
		return nil
	}
	if table[SLOAD].constantGas != 0 || table[SLOAD].dynamicGas == nil || reflect.ValueOf(table[SLOAD].dynamicGas).Pointer() != reflect.ValueOf(gasSLoadEIP2929).Pointer() {
		return nil
	}
	var tests []addressPredicate
	var guards []addressPredicate
	p := &addressSlicePlan{entry: entry, minStack: 2, fixedGas: 42}
	const base = 32
	st := make([]addressValue, base)
	for i := range st {
		st[i] = addressLeaf(addressInput, base-i)
	}
	// These are dynamic entry guards, not assumptions learned from a sample.
	st[base-2] = addressConst(64)
	st[base-1] = addressConst(0)
	st = st[:base-1]
	st[base-2] = addressLeaf(addressLoaded, 0)
	var wrote [2]bool
	hashed := false
	pc := entry + 2
	need := func(n int) bool {
		if len(st) < n {
			return false
		}
		d := base - len(st) + n
		if d > p.minStack {
			p.minStack = d
		}
		return true
	}
	pop := func() addressValue { x := st[len(st)-1]; st = st[:len(st)-1]; return x }
	push := func(v addressValue) {
		st = append(st, v)
		if len(st)-base > p.extraPeak {
			p.extraPeak = len(st) - base
		}
	}
	for count := 0; count < 72 && pc < uint64(len(code)); count++ {
		p.steps++
		op := OpCode(code[pc])
		o := table[op]
		if op == SSTORE {
			if p.helper == nil || !hashed || !wrote[0] || !wrote[1] || !need(2) {
				return nil
			}
			if st[len(st)-1].kind != addressKey {
				return nil
			}
			p.exit = pc
			p.delta = len(st) - base
			for i, v := range st {
				if v.kind == addressInput && v.depth == base-i && v.mask == addressLeaf(addressInput, 0).mask {
					continue
				}
				if v.kind == addressInput && v.depth > p.minStack {
					p.minStack = v.depth
				}
				if i-base < -p.minStack {
					p.minStack = base - i
				}
				p.writes = append(p.writes, addressWrite{i - base, v})
			}
			if len(p.writes) > 32 || p.minStack > 32 {
				return nil
			}
			if len(guards) != 1 {
				return nil
			}
			return &addressGuardPlan{summary: p, guards: guards}
		}
		if o == nil || o.undefined || !need(o.minStack) || o.dynamicGas != nil && op != MSTORE && op != KECCAK256 && op != EXP {
			return nil
		}
		p.fixedGas += o.constantGas
		switch {
		case op >= PUSH0 && op <= PUSH32:
			n := uint64(op - PUSH0)
			if pc+n >= uint64(len(code)) {
				return nil
			}
			var v uint256.Int
			v.SetBytes(code[pc+1 : pc+1+n])
			push(addressValue{value: v})
			pc += n
		case op >= DUP1 && op <= DUP16:
			push(st[len(st)-int(op-DUP1)-1])
		case op >= SWAP1 && op <= SWAP16:
			j := len(st) - int(op-SWAP1) - 2
			st[len(st)-1], st[j] = st[j], st[len(st)-1]
		case op == POP:
			pop()
		case op == JUMPDEST:
		case op == CALLER:
			push(addressLeaf(addressCaller, 0))
		case op == AND || op == ADD || op == SUB || op == SHL || op == EXP:
			a, c := pop(), pop()
			var v uint256.Int
			if a.kind == addressConstant && c.kind == addressConstant {
				switch op {
				case AND:
					v.And(&a.value, &c.value)
				case ADD:
					v.Add(&a.value, &c.value)
				case SUB:
					v.Sub(&a.value, &c.value)
				case SHL:
					if a.value.LtUint64(256) {
						v.Lsh(&c.value, uint(a.value.Uint64()))
					}
				case EXP:
					if reflect.ValueOf(o.dynamicGas).Pointer() != reflect.ValueOf(gasExpEIP158).Pointer() {
						return nil
					}
					p.fixedGas += params.ExpGas + uint64((c.value.BitLen()+7)/8)*params.ExpByteEIP158
					v.Exp(&a.value, &c.value)
				}
				push(addressValue{value: v})
			} else if op == AND && (a.kind == addressConstant || c.kind == addressConstant) {
				if a.kind == addressConstant {
					a, c = c, a
				}
				a.mask.And(&a.mask, &c.value)
				push(a)
			} else {
				return nil
			}
		case op == LT || op == GT || op == EQ:
			a, c := pop(), pop()
			if a.kind == addressPredicateKind || c.kind == addressPredicateKind {
				return nil
			}
			tests = append(tests, addressPredicate{a: a, b: c, op: op})
			push(addressValue{kind: addressPredicateKind, depth: len(tests) - 1})
		case op == ISZERO:
			v := pop()
			if v.kind != addressPredicateKind {
				return nil
			}
			v.value[0] ^= 1
			push(v)
		case op == JUMPI:
			target, v := pop(), pop()
			if p.helper != nil || len(guards) != 0 || v.kind != addressPredicateKind || target.kind != addressConstant || !target.value.IsUint64() {
				return nil
			}
			dest := target.value.Uint64()
			if dest <= pc || dest > pc+256 || dest >= uint64(len(code)) || OpCode(code[dest]) != JUMPDEST {
				return nil
			}
			g := tests[v.depth]
			g.invert = v.value[0] != 0
			guards = append(guards, g)
			pc = dest - 1
		case op == JUMP:
			dest := pop()
			if dest.kind != addressConstant || !dest.value.IsUint64() {
				return nil
			}
			h := b.regions[dest.value.Uint64()]
			if p.helper != nil || h == nil || h.syncHot.native == nil || h.syncHot.fixedLayout == 0 || h.extraGas != 0 || !h.constantSuccessGas || h.pathDependentResult || !need(3) {
				return nil
			}
			p.continuation = pc + 1
			if p.continuation >= uint64(len(code)) || OpCode(code[p.continuation]) != JUMPDEST {
				return nil
			}
			p.k = st[len(st)-3]
			p.a = st[len(st)-2]
			p.b = st[len(st)-1]
			if p.k.kind == addressConstant {
				if !p.k.value.IsUint64() || p.k.value.Uint64() != p.continuation {
					return nil
				}
			} else if p.k.kind != addressInput {
				return nil
			}
			if h.syncHot.fixedLayout == 2 {
				p.a, p.b = p.b, p.a
			}
			if p.a.kind != addressLoaded && p.b.kind != addressLoaded {
				return nil
			}
			peak := len(st) - base + h.peak - 3
			if peak > p.extraPeak {
				p.extraPeak = peak
			}
			p.helper = h
			p.fixedGas += h.syncHot.admissionGas
			st = st[:len(st)-3]
			push(addressLeaf(addressResult, 0))
		case op == MSTORE:
			if p.helper == nil || hashed {
				return nil
			}
			off, v := pop(), pop()
			if off.kind != addressConstant || !off.value.IsUint64() || (off.value.Uint64() != 0 && off.value.Uint64() != 32) || v.kind == addressLoaded || v.kind == addressResult || v.kind == addressKey {
				return nil
			}
			i := int(off.value.Uint64() / 32)
			if wrote[i] && p.memory[i] != v {
				return nil
			}
			p.memory[i] = v
			wrote[i] = true
		case op == KECCAK256:
			off, size := pop(), pop()
			if p.helper == nil || hashed || !wrote[0] || !wrote[1] || off != addressConst(0) || size != addressConst(64) {
				return nil
			}
			p.fixedGas += 2 * params.Keccak256WordGas
			push(addressLeaf(addressKey, 0))
			hashed = true
		default:
			return nil
		}
		pc++
	}
	return nil
}
func (g *addressGuardPlan) try(pc *uint64, s *ScopeContext, e *EVM) bool {
	p := g.summary
	in := s.Stack.data
	n := len(in)
	if n < p.minStack || n+p.extraPeak > 1024 || s.Memory.Len() < 64 || e.abort.Load() || e.Config.Tracer != nil || e.Config.EnablePreimageRecording || s.Contract.Gas < p.fixedGas+params.ColdSloadCostEIP2929 {
		return false
	}
	if !in[n-1].IsZero() || !in[n-2].IsUint64() || in[n-2].Uint64() != 64 {
		return false
	}
	zero := uint256.Int{}
	k := p.k.eval(in, s, zero, zero, zero)
	if !k.IsUint64() || k.Uint64() != p.continuation {
		return false
	}
	for i := 0; i < 2; i++ {
		v := p.memory[i].eval(in, s, zero, zero, zero)
		var got uint256.Int
		got.SetBytes(s.Memory.store[i*32 : (i+1)*32])
		if got != v {
			return false
		}
	}
	s.Contract.Gas -= 42
	opKeccak256(pc, e, s)
	key := *s.Stack.peek()
	cost, _ := gasSLoadEIP2929(e, s.Contract, s.Stack, s.Memory, 0)
	s.Contract.Gas -= cost
	opSload(pc, e, s)
	loaded := *s.Stack.peek()
	if !g.accept(in, s, loaded) {
		*pc = p.entry + 2
		return true
	}
	a, b := p.a.eval(in, s, loaded, zero, key), p.b.eval(in, s, loaded, zero, key)
	ix, result := p.helper.syncHot.native(a, b)
	if p.helper.syncHot.costs[ix] != p.helper.syncHot.admissionGas || e.abort.Load() {
		*pc = p.entry + 2
		return true
	}
	if paperWithoutReuse {
		key = paperRehash64(e, s)
	}
	var values [32]uint256.Int
	for i, w := range p.writes {
		values[i] = w.value.eval(in, s, loaded, result, key)
	}
	if cap(s.Stack.data) < n+p.delta {
		s.Stack.data = append(s.Stack.data, make([]uint256.Int, n+p.delta-len(s.Stack.data))...)
	} else {
		s.Stack.data = s.Stack.data[:n+p.delta]
	}
	for i, w := range p.writes {
		s.Stack.data[n+w.offset] = values[i]
	}
	s.Contract.Gas -= p.fixedGas - 42
	*pc = p.exit
	return true
}
