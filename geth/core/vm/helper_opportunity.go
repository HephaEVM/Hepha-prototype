package vm

// Explicitly enabled original-execution census. No interpreter hot path calls
// this code. Dynamic witnesses are not static proofs or permission to optimize.
import (
	"fmt"
	"math/big"

	"github.com/holiman/uint256"
	miner "go-eth1.16.5-evm/template_miner"
)

type HelperHint struct {
	Kind            string
	Entry           uint64
	Inputs, Outputs int
	ID, Name        string           `json:",omitempty"`
	Region          *miner.ABIRegion `json:",omitempty"`
}
type HelperVisit struct {
	Kind                                 string
	Entry, Start, End, Instructions, Gas uint64
	Success, Wide, ResultMatches         bool
	Inputs, Outputs                      []string `json:",omitempty"`
}
type HelperPacked struct {
	Kind       string
	PC, LoadPC uint64
	// Bits and Shift describe extraction only. A merge can update disjoint
	// bits rather than one contiguous field, so both are unspecified (zero).
	Bits, Shift         int
	ComputeInstructions uint64
}
type HelperCensusCounts struct {
	Instructions              uint64
	Opcodes                   [256]uint64
	Visits                    []HelperVisit
	Packed                    []HelperPacked
	PackedComputeInstructions uint64
	MulmodMax                 uint64
}
type helperPending struct {
	h                HelperHint
	start, gas, exit uint64
	prefix           int
	input            []uint256.Int
	unsafe           bool
}
type helperValue struct {
	op              OpCode
	word            uint256.Int
	known, constant bool
	a, b            *helperValue
	ordinal, pc     uint64
	depth           int
	ones            uint256.Int
	load            uint64
	slot            uint256.Int
}
type HelperCensusProbe struct {
	table      *JumpTable
	hints      map[uint64][]HelperHint
	pending    []helperPending
	counts     HelperCensusCounts
	stack      []*helperValue
	seenFields map[uint64]bool
	compute    map[uint64]bool
}

func NewHelperCensusProbe(table *JumpTable, hints []HelperHint) *HelperCensusProbe {
	p := &HelperCensusProbe{table: table, hints: map[uint64][]HelperHint{}, seenFields: map[uint64]bool{}, compute: map[uint64]bool{}}
	for _, h := range hints {
		p.hints[h.Entry] = append(p.hints[h.Entry], h)
	}
	return p
}
func helperAll() uint256.Int { var v uint256.Int; v.SetAllOne(); return v }
func helperLowMask(v uint256.Int) (int, bool) {
	if v.IsZero() {
		return 0, false
	}
	var w uint256.Int
	w.AddUint64(&v, 1)
	if w.IsZero() {
		return 256, true
	}
	if w.BitLen() > 1 {
		var one uint256.Int
		one.Lsh(uint256.NewInt(1), uint(w.BitLen()-1))
		if w == one {
			return w.BitLen() - 1, true
		}
	}
	if w.Uint64() == 2 && w.IsUint64() {
		return 1, true
	}
	return 0, false
}

// Recognize a storage word with zero or more fixed shifts/masks. SLOAD itself
// is an input boundary and is excluded from the computation opportunity.
func helperExtract(v *helperValue) (load *helperValue, shift, bits int, ok bool) {
	if v == nil {
		return
	}
	if v.op == SLOAD {
		return v, 0, 256, true
	}
	if v.op == AND {
		a, b := v.a, v.b
		if a != nil && a.constant {
			a, b = b, a
		}
		if b == nil || !b.constant {
			return
		}
		n, yes := helperLowMask(b.word)
		if !yes {
			return
		}
		l, s, k, yes := helperExtract(a)
		if yes {
			return l, s, min(k, n), true
		}
		return
	}
	if v.op == SHR || v.op == SAR || v.op == DIV {
		a, b := v.a, v.b
		var shiftBy uint64
		if v.op == DIV {
			if b == nil || !b.constant {
				return
			}
			n := b.word.BitLen()
			if n == 0 {
				return
			}
			var pow uint256.Int
			pow.Lsh(uint256.NewInt(1), uint(n-1))
			if pow != b.word {
				return
			}
			shiftBy = uint64(n - 1)
		} else {
			if a == nil || !a.constant || !a.word.IsUint64() {
				return
			}
			shiftBy = a.word.Uint64()
			a = b
		}
		if shiftBy >= 256 {
			return
		}
		l, s, k, yes := helperExtract(a)
		if yes && int(shiftBy) < k {
			return l, s + int(shiftBy), k - int(shiftBy), true
		}
	}
	if v.op == SIGNEXTEND && v.a != nil && v.a.constant && v.a.word.IsUint64() && v.a.word.Uint64() < 31 {
		l, s, k, yes := helperExtract(v.b)
		if yes {
			return l, s, min(k, 8*(int(v.a.word.Uint64())+1)), true
		}
	}
	return
}
func (p *HelperCensusProbe) addDeps(v *helperValue, seen map[uint64]bool) {
	if v == nil || v.op == SLOAD || v.depth == 0 {
		return
	}
	if seen[v.ordinal] {
		return
	}
	seen[v.ordinal] = true
	p.compute[v.ordinal] = true
	p.addDeps(v.a, seen)
	p.addDeps(v.b, seen)
}
func (p *HelperCensusProbe) packed(v *helperValue, pc uint64, kind string) {
	if v == nil || p.seenFields[v.ordinal] {
		return
	}
	l, s, b, ok := helperExtract(v)
	if !ok || b == 256 {
		return
	}
	p.seenFields[v.ordinal] = true
	deps := map[uint64]bool{}
	p.addDeps(v, deps)
	p.counts.Packed = append(p.counts.Packed, HelperPacked{kind, pc, l.pc, b, s, uint64(len(deps))})
}
func (p *HelperCensusProbe) Observe(pc, gas uint64, scope *ScopeContext, err error) error {
	words := scope.Stack.data
	if len(words) != len(p.stack) {
		return fmt.Errorf("helper value stack mismatch pc=%d actual=%d symbolic=%d", pc, len(words), len(p.stack))
	}
	for i, v := range p.stack {
		if !v.known {
			v.word, v.known = words[i], true
			if v.constant {
				v.ones = v.word
			}
		}
	}
	// A helper ends before executing its continuation. Same prefix depth and
	// exact mathematical output are checked for cast/mulDiv witnesses.
	active := p.pending[:0]
	for _, v := range p.pending {
		if pc == v.exit && p.counts.Instructions > v.start {
			r := HelperVisit{Kind: v.h.Kind, Entry: v.h.Entry, Start: v.start, End: p.counts.Instructions, Instructions: p.counts.Instructions - v.start, Gas: v.gas - gas, Success: !v.unsafe}
			for _, w := range v.input {
				r.Inputs = append(r.Inputs, w.Hex())
			}
			if len(words) >= v.prefix {
				for _, w := range words[v.prefix:] {
					r.Outputs = append(r.Outputs, w.Hex())
				}
			}
			if v.h.Kind == "muldiv_candidate" {
				r.ResultMatches = false
				if len(v.input) == 4 && len(words) == v.prefix+1 {
					a, b, d := v.input[1].ToBig(), v.input[2].ToBig(), v.input[3].ToBig()
					prod := new(big.Int).Mul(a, b)
					r.Wide = prod.BitLen() > 256
					if d.Sign() != 0 {
						prod.Div(prod, d)
						r.ResultMatches = prod.BitLen() <= 256 && prod.Cmp(words[v.prefix].ToBig()) == 0
					}
				}
			} else {
				r.ResultMatches = len(words) == v.prefix+v.h.Outputs || v.h.Outputs == 0
			}
			p.counts.Visits = append(p.counts.Visits, r)
		} else if p.counts.Instructions-v.start < 4096 {
			active = append(active, v)
		}
	}
	p.pending = active
	if err != nil {
		return nil
	}
	for _, h := range p.hints[pc] {
		if len(words) < h.Inputs {
			continue
		}
		prefix := len(words) - h.Inputs
		k := words[prefix]
		if !scope.Contract.validJumpdest(&k) {
			continue
		}
		p.pending = append(p.pending, helperPending{h: h, start: p.counts.Instructions, gas: gas, exit: k.Uint64(), prefix: prefix, input: append([]uint256.Int(nil), words[prefix:]...)})
	}
	op := OpCode(scope.Contract.GetOp(pc))
	operation := p.table[op]
	n := len(p.stack)
	if n < operation.minStack || n > operation.maxStack {
		return nil
	}
	p.counts.Instructions++
	p.counts.Opcodes[op]++
	if op == MULMOD && words[n-3] == helperAll() {
		p.counts.MulmodMax++
	}
	// Side effects invalidate purported pure math witnesses. ABI intentionally
	// retains memory effects; external/state operations invalidate both.
	for i := range p.pending {
		if op == SLOAD || op == SSTORE || op == TLOAD || op == TSTORE || op == CALL || op == STATICCALL || op == DELEGATECALL || op == CREATE || op == CREATE2 {
			p.pending[i].unsafe = true
		}
	}
	if op >= DUP1 && op <= DUP16 {
		p.stack = append(p.stack, p.stack[n-int(op-DUP1)-1])
		return nil
	}
	if op >= SWAP1 && op <= SWAP16 {
		i := n - int(op-SWAP1) - 2
		p.stack[n-1], p.stack[i] = p.stack[i], p.stack[n-1]
		return nil
	}
	pops := operation.minStack
	pushes := pops + 1024 - operation.maxStack
	var a, b *helperValue
	if pops > 0 {
		a = p.stack[n-1]
	}
	if pops > 1 {
		b = p.stack[n-2]
	}
	// Defer extraction reporting until a consumer: SHR then AND is one field,
	// not two. Pure DUP/SWAP preserve identity and cannot duplicate the count.
	if op != AND && op != SHR && op != SAR && op != DIV && op != SIGNEXTEND && op != POP {
		for _, v := range p.stack[n-pops:] {
			p.packed(v, pc, "extract")
		}
	}
	if op == SSTORE && b != nil && b.op == OR {
		x, y := b.a, b.b
		for turn := 0; turn < 2; turn++ {
			if x != nil && x.op == AND {
				old, mask := x.a, x.b
				if old != nil && old.constant {
					old, mask = mask, old
				}
				if old != nil && old.op == SLOAD && mask != nil && mask.constant && y != nil && old.slot == words[n-1] {
					var overlap uint256.Int
					overlap.And(&mask.word, &y.ones)
					if overlap.IsZero() && mask.word != helperAll() {
						deps := map[uint64]bool{}
						p.addDeps(b, deps)
						p.counts.Packed = append(p.counts.Packed, HelperPacked{"merge", pc, old.pc, 0, 0, uint64(len(deps))})
						break
					}
				}
			}
			x, y = y, x
		}
	}
	p.stack = p.stack[:n-pops]
	for i := 0; i < pushes; i++ {
		v := &helperValue{op: op, ordinal: p.counts.Instructions, pc: pc, ones: helperAll()}
		if op >= PUSH0 && op <= PUSH32 {
			end := pc + 1 + uint64(op-PUSH0)
			if end > uint64(len(scope.Contract.Code)) {
				end = uint64(len(scope.Contract.Code))
			}
			v.word.SetBytes(scope.Contract.Code[pc+1 : end])
			v.known = true
			v.constant = true
			v.depth = 1
			v.ones = v.word
		} else if op == SLOAD {
			v.load = v.ordinal
			v.slot = words[n-1]
		} else if (op == AND || op == OR || op == SHR || op == SHL || op == SAR || op == DIV || op == SIGNEXTEND || op == NOT || op == SUB || op == ADD || op == EXP) && a != nil {
			depth := a.depth
			if b != nil {
				depth = max(depth, b.depth)
			}
			if depth < 16 {
				v.a, v.b, v.depth = a, b, depth+1
				v.constant = a.constant && (op == NOT || b != nil && b.constant)
				switch op {
				case AND:
					if b != nil {
						v.ones.And(&a.ones, &b.ones)
					}
				case OR:
					if b != nil {
						v.ones.Or(&a.ones, &b.ones)
					}
				case SHL:
					if a.constant && a.word.IsUint64() && b != nil {
						v.ones.Lsh(&b.ones, uint(a.word.Uint64()))
					}
				case SHR:
					if a.constant && a.word.IsUint64() && b != nil {
						v.ones.Rsh(&b.ones, uint(a.word.Uint64()))
					}
				}
			}
		}
		p.stack = append(p.stack, v)
	}
	return nil
}
func (p *HelperCensusProbe) Finish() HelperCensusCounts {
	p.counts.PackedComputeInstructions = uint64(len(p.compute))
	for _, v := range p.pending {
		p.counts.Visits = append(p.counts.Visits, HelperVisit{Kind: v.h.Kind, Entry: v.h.Entry, Start: v.start, End: p.counts.Instructions})
	}
	return p.counts
}
