package templateminer

// The extra string is a memory input, never an extra arithmetic operand.
// Normal paths are analyzed symbolically; failure encoders (including copy
// loops) remain executable continuations. Bounded, varied concrete checks
// corroborate that they return Error(string) with that argument. This is not
// a proof for every possible string or memory layout, nor pruning permission.
import (
	"bytes"
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/core/vm"
)

type localValue struct {
	value    *big.Int
	producer *uint64
}
type localExecution struct {
	PCs     []uint64
	Before  [][]*big.Int
	Output  []*big.Int
	Outcome string
	Data    []byte
	Jumps   map[uint64]uint64
	Writes  bool
}

// Deliberately excludes all storage/environment/call opcodes. Memory writes
// occur on a private copy. The original contract/state is never executed here.
func executeLocal(p *Program, entry uint64, initial []*big.Int, original []byte, continuation uint64) (*localExecution, error) {
	r := &localExecution{Jumps: map[uint64]uint64{}}
	st := make([]localValue, len(initial))
	for i, v := range initial {
		st[i] = localValue{value: new(big.Int).Mod(new(big.Int).Set(v), modulus)}
	}
	mem := map[uint64]byte{}
	pop := func() localValue { v := st[len(st)-1]; st = st[:len(st)-1]; return v }
	read := func(at uint64) byte {
		if v, ok := mem[at]; ok {
			return v
		}
		if at < uint64(len(original)) {
			return original[at]
		}
		return 0
	}
	need := func(n int) error {
		if len(st) < n {
			return fmt.Errorf("local stack underflow")
		}
		return nil
	}
	bounded := func(v *big.Int) (uint64, error) {
		if !v.IsUint64() || v.Uint64() > 1<<20 {
			return 0, fmt.Errorf("local memory bound")
		}
		return v.Uint64(), nil
	}
	pc := entry
	for count := 0; count < 8192; count++ {
		ins, ok := p.ByPC[pc]
		if !ok {
			return nil, fmt.Errorf("local instruction boundary %d", pc)
		}
		r.PCs = append(r.PCs, pc)
		before := make([]*big.Int, len(st))
		for i, v := range st {
			before[i] = v.value
		}
		r.Before = append(r.Before, before)
		pc += uint64(1 + len(ins.Data))
		push := func(v *big.Int) { st = append(st, localValue{value: new(big.Int).Mod(v, modulus)}) }
		switch {
		case ins.Op.IsPush():
			at := ins.PC
			st = append(st, localValue{new(big.Int).SetBytes(ins.Data), &at})
		case ins.Op >= vm.DUP1 && ins.Op <= vm.DUP16:
			k := int(ins.Op-vm.DUP1) + 1
			if err := need(k); err != nil {
				return nil, err
			}
			st = append(st, st[len(st)-k])
		case ins.Op >= vm.SWAP1 && ins.Op <= vm.SWAP16:
			k := int(ins.Op-vm.SWAP1) + 1
			if err := need(k + 1); err != nil {
				return nil, err
			}
			n := len(st)
			st[n-1], st[n-1-k] = st[n-1-k], st[n-1]
		case ins.Op == vm.JUMPDEST:
		case ins.Op == vm.POP:
			if err := need(1); err != nil {
				return nil, err
			}
			pop()
		case ins.Op == vm.JUMP || ins.Op == vm.JUMPI:
			n := 1
			if ins.Op == vm.JUMPI {
				n = 2
			}
			if err := need(n); err != nil {
				return nil, err
			}
			dest := pop()
			take := true
			if n == 2 {
				take = pop().value.Sign() != 0
			}
			if !dest.value.IsUint64() {
				return nil, fmt.Errorf("local jump is not uint64")
			}
			target := dest.value.Uint64()
			if ins.Op == vm.JUMP && target == continuation && dest.producer == nil {
				r.Outcome = "return"
				for _, v := range st {
					r.Output = append(r.Output, v.value)
				}
				return r, nil
			}
			if dest.producer == nil {
				return nil, fmt.Errorf("computed local jump")
			}
			if old, ok := r.Jumps[*dest.producer]; ok && old != target {
				return nil, fmt.Errorf("inconsistent jump producer")
			}
			r.Jumps[*dest.producer] = target
			if take {
				if p.ByPC[target].Op != vm.JUMPDEST {
					return nil, fmt.Errorf("invalid local jump")
				}
				pc = target
			}
		case ins.Op == vm.MLOAD:
			if err := need(1); err != nil {
				return nil, err
			}
			at, err := bounded(pop().value)
			if err != nil {
				return nil, err
			}
			data := make([]byte, 32)
			for i := range data {
				data[i] = read(at + uint64(i))
			}
			push(new(big.Int).SetBytes(data))
		case ins.Op == vm.MSTORE || ins.Op == vm.MSTORE8:
			if err := need(2); err != nil {
				return nil, err
			}
			at, err := bounded(pop().value)
			if err != nil {
				return nil, err
			}
			v := pop().value
			if ins.Op == vm.MSTORE8 {
				mem[at] = byte(v.Uint64())
			} else {
				for i, b := range v.FillBytes(make([]byte, 32)) {
					mem[at+uint64(i)] = b
				}
			}
			r.Writes = true
		case ins.Op == vm.CODECOPY:
			if err := need(3); err != nil {
				return nil, err
			}
			at, e := bounded(pop().value)
			src, e2 := bounded(pop().value)
			size, e3 := bounded(pop().value)
			if e != nil || e2 != nil || e3 != nil || size > 4096 {
				return nil, fmt.Errorf("unsupported local CODECOPY")
			}
			for k := uint64(0); k < size; k++ {
				v := byte(0)
				if src+k < uint64(len(p.Code)) {
					v = p.Code[src+k]
				}
				mem[at+k] = v
			}
			r.Writes = true
		case ins.Op == vm.REVERT:
			if err := need(2); err != nil {
				return nil, err
			}
			at, e := bounded(pop().value)
			size, e2 := bounded(pop().value)
			if e != nil || e2 != nil || size > 8192 {
				return nil, fmt.Errorf("local revert bound")
			}
			r.Data = make([]byte, size)
			for i := range r.Data {
				r.Data[i] = read(at + uint64(i))
			}
			r.Outcome = "revert"
			return r, nil
		case ins.Op == vm.INVALID:
			r.Outcome = "invalid"
			return r, nil
		default:
			n := 2
			switch ins.Op {
			case vm.ADD, vm.SUB, vm.MUL, vm.DIV, vm.MOD, vm.SDIV, vm.SMOD, vm.SLT, vm.SGT, vm.LT, vm.GT, vm.EQ, vm.AND, vm.OR, vm.XOR, vm.SHL, vm.SHR, vm.SAR, vm.SIGNEXTEND, vm.EXP, vm.BYTE:
			case vm.ISZERO, vm.NOT:
				n = 1
			default:
				return nil, fmt.Errorf("unsupported local opcode %s", ins.Op)
			}
			if err := need(n); err != nil {
				return nil, err
			}
			args := make([]*Expr, n)
			for i := range args {
				args[i] = constant(pop().value, nil)
			}
			v, err := (&Expr{Op: ins.Op.String(), Args: args}).Evaluate(nil)
			if err != nil {
				return nil, err
			}
			push(v)
		}
	}
	return nil, fmt.Errorf("local execution budget exhausted")
}

func errorString(data []byte) ([]byte, bool) {
	if len(data) < 68 || !bytes.Equal(data[:4], []byte{8, 195, 121, 160}) || new(big.Int).SetBytes(data[4:36]).Cmp(big.NewInt(32)) != 0 {
		return nil, false
	}
	n := new(big.Int).SetBytes(data[36:68])
	if !n.IsUint64() || n.Uint64() > 4096 || uint64(len(data)) != 68+(n.Uint64()+31)/32*32 {
		return nil, false
	}
	return data[68 : 68+n.Uint64()], true
}
func stringArgumentMemory(text []byte, ptr, free uint64) []byte {
	size := free + 512
	if ptr+32+uint64(len(text))+32 > size {
		size = ptr + 32 + uint64(len(text)) + 32
	}
	m := make([]byte, size)
	// Nonzero padding catches an accidental assumption that unused memory is zero.
	for i := range m {
		m[i] = 0xa5
	}
	new(big.Int).SetUint64(free).FillBytes(m[64:96])
	new(big.Int).SetUint64(uint64(len(text))).FillBytes(m[ptr : ptr+32])
	copy(m[ptr+32:], text)
	return m
}
func matchingPath(a *SymbolicAnalysis, inputs map[string]*big.Int) (*SymbolicPath, error) {
	var found *SymbolicPath
	for i := range a.Paths {
		q := &a.Paths[i]
		holds := true
		for _, c := range q.Conditions {
			v, e := c.Expression.Evaluate(inputs)
			if e != nil {
				return nil, e
			}
			if (v.Sign() != 0) != c.Nonzero {
				holds = false
				break
			}
		}
		if holds {
			if found != nil {
				return nil, fmt.Errorf("ambiguous symbolic path")
			}
			found = q
		}
	}
	if found == nil {
		return nil, fmt.Errorf("no symbolic path")
	}
	return found, nil
}

func validateArgumentFailure(p *Program, entry uint64, a *SymbolicAnalysis, t integerType) error {
	seen := map[*SymbolicPath]bool{}
	union := map[*SymbolicPath]map[uint64]bool{}
	const continuation = uint64(0x7fffffffffffffff)
	for _, pair := range typedValidationPairs(t) {
		inputs := map[string]*big.Int{a.OperandOrder[0]: pair[0], a.OperandOrder[1]: pair[1]}
		q, err := matchingPath(a, inputs)
		if err != nil {
			return err
		}
		if q.Outcome != "revert_argument" {
			continue
		}
		lengths := []int{33}
		if !seen[q] {
			lengths = []int{0, 1, 30, 31, 32, 33, 63, 64, 65, 257}
		}
		for _, length := range lengths {
			for _, ptr := range []uint64{256, 768} {
				text := make([]byte, length)
				for i := range text {
					text[i] = byte(i*17 + 3)
				}
				inputs["continuation"] = new(big.Int).SetUint64(continuation)
				inputs["errorMessage"] = new(big.Int).SetUint64(ptr)
				initial := []*big.Int{}
				for _, name := range analysisInputs(a) {
					initial = append(initial, inputs[name])
				}
				r, e := executeLocal(p, entry, initial, stringArgumentMemory(text, ptr, 4096), continuation)
				if e != nil {
					return fmt.Errorf("argument failure encoder: %w", e)
				}
				message, ok := errorString(r.Data)
				if r.Outcome != "revert" || !ok || !bytes.Equal(message, text) {
					return fmt.Errorf("failure does not return the errorMessage argument")
				}
				if len(r.PCs) <= len(q.PCs) {
					return fmt.Errorf("missing failure continuation")
				}
				for i, pc := range q.PCs {
					if r.PCs[i] != pc {
						return fmt.Errorf("failure prefix mismatch")
					}
				}
				if union[q] == nil {
					union[q] = map[uint64]bool{}
				}
				for _, pc := range r.PCs[len(q.PCs):] {
					if !union[q][pc] {
						union[q][pc] = true
						q.FailureProgram = append(q.FailureProgram, pc)
						a.ReachablePCs = append(a.ReachablePCs, pc)
					}
					delete(a.DeadPushes, pc)
				}
				for pc, target := range r.Jumps {
					if old, ok := a.JumpProducers[pc]; ok && old != target {
						return fmt.Errorf("inconsistent failure target")
					}
					a.JumpProducers[pc] = target
					delete(a.DeadPushes, pc)
				}
			}
		}
		seen[q] = true
	}
	for i := range a.Paths {
		if a.Paths[i].Outcome == "revert_argument" && !seen[&a.Paths[i]] {
			return fmt.Errorf("unvalidated error argument failure path")
		}
	}
	return nil
}

func argumentWitness(p *Program, f *chainFrame, end AuditRecord, start int, c *candidate) (*ChainWitness, error) {
	z := f.ops[start]
	arity := len(analysisInputs(c.analysis))
	n := len(z.Stack)
	if n < arity {
		return nil, fmt.Errorf("four-slot entry underflow")
	}
	values := make([]*big.Int, n)
	for i, text := range z.Stack {
		v, e := word(text)
		if e != nil {
			return nil, e
		}
		values[i] = v
	}
	ret := values[n-arity]
	if !ret.IsUint64() {
		return nil, fmt.Errorf("invalid continuation")
	}
	inputs := map[string]*big.Int{}
	for i, name := range analysisInputs(c.analysis) {
		inputs[name] = values[n-arity+i]
	}
	if e := CheckInputs(c.analysis, c.analysis.Operation, inputs[c.analysis.OperandOrder[0]], inputs[c.analysis.OperandOrder[1]]); e != nil {
		return nil, e
	}
	path, e := matchingPath(c.analysis, inputs)
	if e != nil {
		return nil, e
	}
	var mem []byte
	if path.Outcome == "revert_argument" {
		mem, e = hex.DecodeString(z.Memory)
		if e != nil || len(mem) == 0 {
			return nil, fmt.Errorf("error argument failure requires entry memory witness")
		}
	}
	run, e := executeLocal(p, c.helper.Entry, values, mem, ret.Uint64())
	if e != nil {
		return nil, e
	}
	if start+len(run.PCs) > len(f.ops) {
		return nil, fmt.Errorf("incomplete argument-helper trace")
	}
	for k, pc := range run.PCs {
		step := f.ops[start+k]
		if step.PC != pc || len(step.Stack) != len(run.Before[k]) {
			return nil, fmt.Errorf("argument-helper trace shape mismatch")
		}
		for j, txt := range step.Stack {
			v, e := word(txt)
			if e != nil || v.Cmp(run.Before[k][j]) != 0 {
				return nil, fmt.Errorf("argument-helper concrete stack mismatch")
			}
		}
	}
	w := &ChainWitness{Block: f.header.Block, Tx: f.header.Tx, CallID: f.header.CallID, EntryStep: start, InstructionCount: len(run.PCs), Continuation: ret.Uint64(), Operands: []string{inputs["a"].Text(16), inputs["b"].Text(16)}, Outcome: run.Outcome, Jumps: []JumpOrigin{}}
	if path.Outcome == "return" {
		at := start + len(run.PCs)
		if run.Outcome != "return" || run.Writes || at >= len(f.ops) || f.ops[at].PC != ret.Uint64() || len(run.Output) != n-arity+1 || len(f.ops[at].Stack) != len(run.Output) {
			return nil, fmt.Errorf("argument-helper return mismatch")
		}
		for i, v := range run.Output {
			observed, e := word(f.ops[at].Stack[i])
			if e != nil || v.Cmp(observed) != 0 {
				return nil, fmt.Errorf("argument-helper output stack mismatch")
			}
			if i < n-arity && v.Cmp(values[i]) != 0 {
				return nil, fmt.Errorf("changed caller prefix")
			}
		}
		result, e := path.Result.Evaluate(inputs)
		if e != nil || result.Cmp(run.Output[len(run.Output)-1]) != 0 {
			return nil, fmt.Errorf("argument-helper result mismatch")
		}
		w.Result = result.Text(16)
	} else {
		actual, e := hex.DecodeString(end.Output)
		if e != nil {
			actual, e = hex.DecodeString(trimHex(end.Output))
		}
		if e != nil || run.Outcome != "revert" || !end.Reverted || start+len(run.PCs) != len(f.ops) || !bytes.Equal(actual, run.Data) {
			return nil, fmt.Errorf("argument-helper actual failure mismatch")
		}
		w.FailureData = hex.EncodeToString(run.Data)
	}
	return w, nil
}
func trimHex(s string) string {
	if len(s) >= 2 && s[:2] == "0x" {
		return s[2:]
	}
	return s
}
