package templateminer

import (
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/core/vm"
)

// Expr is a 256-bit EVM expression, not a host-language mathematical integer.
// Guards are retained in the output; finite differential tests are not proofs.
type Expr struct {
	Op              string  `json:"op"`
	Value           string  `json:"value,omitempty"`
	Input           string  `json:"input,omitempty"`
	Args            []*Expr `json:"args,omitempty"`
	OriginPC        *uint64 `json:"originPc,omitempty"`
	constantOrigins []uint64
}
type Condition struct {
	Expression *Expr  `json:"expression"`
	Nonzero    bool   `json:"nonzero"`
	GuardPC    uint64 `json:"-"`
}
type SymbolicPath struct {
	// Absolute memory high-water mark for constant modern panic encoders.
	MemoryBytes    uint64      `json:"memoryBytes,omitempty"`
	FailureEntry   uint64      `json:"failureEntry,omitempty"`
	FailureProgram []uint64    `json:"failureProgramPcs,omitempty"`
	Conditions     []Condition `json:"conditions"`
	PCs            []uint64    `json:"pcs"`
	Outcome        string      `json:"outcome"`
	Result         *Expr       `json:"result,omitempty"`
	PeakStack      int         `json:"peakStack"`
	FailureData    string      `json:"failureData,omitempty"`
	FailureMask    string      `json:"failureDataKnownMask,omitempty"`
	FailureMemory  *Expr       `json:"failureMemoryOffset,omitempty"`
}
type SymbolicAnalysis struct {
	InputStack   []string       `json:"inputStackBottomToTop,omitempty"`
	Semantics    string         `json:"semantics,omitempty"`
	Operation    string         `json:"operation,omitempty"`
	IntegerType  string         `json:"integerType,omitempty"`
	Paths        []SymbolicPath `json:"paths"`
	OperandOrder []string       `json:"operandOrder"`
	// These are instruction boundaries, not a contiguous byte range: cleanup
	// functions are often elsewhere in the contract and called more than once.
	ReachablePCs      []uint64          `json:"reachablePcs"`
	JumpProducers     map[uint64]uint64 `json:"-"`
	DeadPushes        map[uint64]bool   `json:"-"`
	CodeDataProducers map[uint64]string `json:"-"`
}
type symbolicState struct {
	pc             uint64
	stack          []*Expr
	memory         map[uint64]byte
	relativeMemory map[uint64]byte
	freeMemory     *Expr
	writes         bool
	path           SymbolicPath
}

var modulus = new(big.Int).Lsh(big.NewInt(1), 256)
var mask256 = new(big.Int).Sub(new(big.Int).Set(modulus), big.NewInt(1))

func variable(name string) *Expr { return &Expr{Op: "input", Input: name} }
func constant(n *big.Int, pc *uint64) *Expr {
	n = new(big.Int).Mod(new(big.Int).Set(n), modulus)
	return &Expr{Op: "const", Value: "0x" + n.Text(16), OriginPC: pc}
}
func uintConst(n uint64) *Expr { return constant(new(big.Int).SetUint64(n), nil) }
func constValue(e *Expr) (*big.Int, bool) {
	if e.Op != "const" {
		return nil, false
	}
	n, ok := new(big.Int).SetString(e.Value[2:], 16)
	return n, ok
}
func makeExpr(op string, args ...*Expr) *Expr {
	e := &Expr{Op: op, Args: args}
	all := true
	for _, a := range args {
		if a.Op != "const" {
			all = false
		}
	}
	if all {
		n, err := e.Evaluate(nil)
		if err == nil {
			folded := constant(n, nil)
			for _, arg := range args {
				folded.constantOrigins = append(folded.constantOrigins, expressionOrigins(arg)...)
			}
			return folded
		}
	}
	return e
}
func (e *Expr) Evaluate(inputs map[string]*big.Int) (*big.Int, error) {
	if e.Op == "const" {
		n, _ := constValue(e)
		return n, nil
	}
	if e.Op == "input" {
		n, ok := inputs[e.Input]
		if !ok {
			return nil, fmt.Errorf("missing input %s", e.Input)
		}
		return new(big.Int).Mod(new(big.Int).Set(n), modulus), nil
	}
	a := make([]*big.Int, len(e.Args))
	for i, arg := range e.Args {
		n, err := arg.Evaluate(inputs)
		if err != nil {
			return nil, err
		}
		a[i] = n
	}
	z := new(big.Int)
	boolean := func(v bool) *big.Int {
		if v {
			return big.NewInt(1)
		}
		return new(big.Int)
	}
	switch e.Op {
	case "ADD":
		z.Add(a[0], a[1])
	case "SUB":
		z.Sub(a[0], a[1])
	case "MUL":
		z.Mul(a[0], a[1])
	case "EXP":
		z.Exp(a[0], a[1], modulus)
	case "DIV":
		if a[1].Sign() != 0 {
			z.Div(a[0], a[1])
		}
	case "MOD":
		if a[1].Sign() != 0 {
			z.Mod(a[0], a[1])
		}
	case "SDIV":
		if a[1].Sign() != 0 {
			z.Quo(signedWord(a[0]), signedWord(a[1]))
		}
	case "SMOD":
		if a[1].Sign() != 0 {
			z.Rem(signedWord(a[0]), signedWord(a[1]))
		}
	case "SLT":
		return boolean(signedWord(a[0]).Cmp(signedWord(a[1])) < 0), nil
	case "SGT":
		return boolean(signedWord(a[0]).Cmp(signedWord(a[1])) > 0), nil
	case "SIGNEXTEND":
		z.Set(a[1])
		if a[0].IsUint64() && a[0].Uint64() < 31 {
			bits := uint(8 * (a[0].Uint64() + 1))
			limit := new(big.Int).Lsh(big.NewInt(1), bits)
			z.Mod(z, limit)
			if z.Bit(int(bits-1)) != 0 {
				z.Sub(z, limit)
			}
		}
	case "LT":
		return boolean(a[0].Cmp(a[1]) < 0), nil
	case "GT":
		return boolean(a[0].Cmp(a[1]) > 0), nil
	case "EQ":
		return boolean(a[0].Cmp(a[1]) == 0), nil
	case "ISZERO":
		return boolean(a[0].Sign() == 0), nil
	case "AND":
		z.And(a[0], a[1])
	case "OR":
		z.Or(a[0], a[1])
	case "XOR":
		z.Xor(a[0], a[1])
	case "NOT":
		z.Xor(a[0], mask256)
	case "SHL":
		if a[0].IsUint64() && a[0].Uint64() < 256 {
			z.Lsh(a[1], uint(a[0].Uint64()))
		}
	case "SHR":
		if a[0].IsUint64() && a[0].Uint64() < 256 {
			z.Rsh(a[1], uint(a[0].Uint64()))
		}
	case "BYTE":
		if a[0].IsUint64() && a[0].Uint64() < 32 {
			z.Rsh(a[1], uint(8*(31-a[0].Uint64())))
			z.And(z, big.NewInt(255))
		}
	case "SAR":
		x := signedWord(a[1])
		if a[0].IsUint64() && a[0].Uint64() < 256 {
			z.Rsh(x, uint(a[0].Uint64()))
		} else if x.Sign() < 0 {
			z.SetInt64(-1)
		}
	default:
		return nil, fmt.Errorf("unsupported expression %s", e.Op)
	}
	return z.Mod(z, modulus), nil
}

func Analyze(p *Program, h Helper) (*SymbolicAnalysis, error) {
	return analyze(p, h, false)
}

func analyze(p *Program, h Helper, legacy bool) (*SymbolicAnalysis, error) {
	return analyzeLayout(p, h, legacy, nil)
}

func analyzeLayout(p *Program, h Helper, legacy bool, layout []string) (*SymbolicAnalysis, error) {
	if i, ok := p.ByPC[h.Entry]; !ok || i.Op != vm.JUMPDEST {
		return nil, fmt.Errorf("helper entry is not JUMPDEST")
	}
	analysis := &SymbolicAnalysis{JumpProducers: map[uint64]uint64{}, DeadPushes: map[uint64]bool{}, CodeDataProducers: map[uint64]string{}}
	if legacy {
		analysis.Semantics = legacySemantics
	}
	livePushes := map[uint64]bool{}
	markLive := func(e *Expr) {
		for _, pc := range expressionOrigins(e) {
			livePushes[pc] = true
		}
	}
	analysis.InputStack = layout
	var initial []*Expr
	for _, name := range analysisInputs(analysis) {
		initial = append(initial, variable(name))
	}
	pending := []symbolicState{{pc: h.Entry, stack: initial, memory: map[uint64]byte{}, relativeMemory: map[uint64]byte{}, freeMemory: variable("freeMemory"), path: SymbolicPath{PeakStack: len(initial)}}}
	seenPC := map[uint64]bool{}
	explored := 0
	for len(pending) > 0 {
		s := pending[len(pending)-1]
		pending = pending[:len(pending)-1]
		terminal := false
		for !terminal {
			explored++
			if explored > 4096 || len(s.path.PCs) >= 512 || len(pending)+len(analysis.Paths) > 32 {
				return nil, fmt.Errorf("symbolic exploration budget exceeded (loop or complex helper)")
			}
			ins, ok := p.ByPC[s.pc]
			if !ok {
				return nil, fmt.Errorf("invalid instruction boundary %d", s.pc)
			}
			if !seenPC[ins.PC] {
				seenPC[ins.PC] = true
				analysis.ReachablePCs = append(analysis.ReachablePCs, ins.PC)
			}
			// A four-slot helper hands failure handling to the original memory
			// encoder. Its tail is validated separately against variable strings;
			// successful paths must still remain memory-free.
			if len(layout) == 4 && (ins.Op == vm.MLOAD || ins.Op == vm.MSTORE || ins.Op == vm.MSTORE8 || ins.Op == vm.CODECOPY) {
				s.path.Outcome = "revert_argument"
				s.path.FailureEntry = ins.PC
				analysis.Paths = append(analysis.Paths, s.path)
				break
			}
			s.path.PCs = append(s.path.PCs, ins.PC)
			n := len(s.stack)
			next := ins.PC + uint64(1+len(ins.Data))
			pop := func() *Expr { e := s.stack[len(s.stack)-1]; s.stack = s.stack[:len(s.stack)-1]; return e }
			need := func(k int) error {
				if len(s.stack) < k {
					return fmt.Errorf("helper reads below its three-slot input at PC %d", ins.PC)
				}
				return nil
			}
			jump := func(dest *Expr) error {
				markLive(dest)
				if dest.Op == "input" && dest.Input == "continuation" {
					if s.writes || len(s.stack) != 1 {
						return fmt.Errorf("successful helper has side effects or unexpected stack layout")
					}
					result := s.stack[0]
					markLive(result)
					op, order, ok := resultOperation(result)
					// Older SafeMath.mul returns zero early when one operand is
					// zero. The other path identifies MUL; the oracle checks that
					// this constant return is restricted to the correct domain.
					if legacy && !ok {
						if v, constant := constValue(result); constant && v.Sign() == 0 {
							s.path.Outcome, s.path.Result = "return", result
							analysis.Paths = append(analysis.Paths, s.path)
							terminal = true
							return nil
						}
					}
					if !ok || (h.Operation != "" && op != h.Operation) || (analysis.Operation != "" && analysis.Operation != op) {
						return fmt.Errorf("returned value is not a binary %s of original inputs", h.Operation)
					}
					analysis.Operation = op
					if analysis.OperandOrder != nil && (analysis.OperandOrder[0] != order[0] || analysis.OperandOrder[1] != order[1]) {
						return fmt.Errorf("inconsistent operand order")
					}
					analysis.OperandOrder = order
					s.path.Outcome = "return"
					s.path.Result = result
					analysis.Paths = append(analysis.Paths, s.path)
					terminal = true
					return nil
				}
				value, ok := constValue(dest)
				if !ok || !value.IsUint64() {
					return fmt.Errorf("unresolved dynamic jump at PC %d", ins.PC)
				}
				target := value.Uint64()
				targetIns, exists := p.ByPC[target]
				if !exists || targetIns.Op != vm.JUMPDEST {
					return fmt.Errorf("invalid jump destination %d", target)
				}
				if dest.OriginPC == nil {
					return fmt.Errorf("computed jump address unsupported")
				}
				analysis.JumpProducers[*dest.OriginPC] = target
				next = target
				return nil
			}
			switch {
			case ins.Op == vm.PUSH0:
				s.stack = append(s.stack, constant(new(big.Int), &ins.PC))
			case ins.Op >= vm.PUSH1 && ins.Op <= vm.PUSH32:
				s.stack = append(s.stack, constant(new(big.Int).SetBytes(ins.Data), &ins.PC))
			case ins.Op >= vm.DUP1 && ins.Op <= vm.DUP16:
				depth := int(ins.Op-vm.DUP1) + 1
				if err := need(depth); err != nil {
					return nil, err
				}
				s.stack = append(s.stack, s.stack[n-depth])
			case ins.Op >= vm.SWAP1 && ins.Op <= vm.SWAP16:
				depth := int(ins.Op-vm.SWAP1) + 1
				if err := need(depth + 1); err != nil {
					return nil, err
				}
				s.stack[n-1], s.stack[n-1-depth] = s.stack[n-1-depth], s.stack[n-1]
			case ins.Op == vm.POP:
				if err := need(1); err != nil {
					return nil, err
				}
				pop()
			case ins.Op == vm.JUMPDEST:
			case ins.Op == vm.JUMP:
				if err := need(1); err != nil {
					return nil, err
				}
				if err := jump(pop()); err != nil {
					return nil, err
				}
			case ins.Op == vm.JUMPI:
				if err := need(2); err != nil {
					return nil, err
				}
				dest, condition := pop(), pop()
				markLive(condition)
				if value, ok := constValue(condition); ok {
					if value.Sign() != 0 {
						if err := jump(dest); err != nil {
							return nil, err
						}
					}
					break
				}
				branch := cloneState(s)
				branch.pc = next
				branch.path.Conditions = append(branch.path.Conditions, Condition{Expression: condition, Nonzero: false, GuardPC: ins.PC})
				pending = append(pending, branch)
				s.path.Conditions = append(s.path.Conditions, Condition{Expression: condition, Nonzero: true, GuardPC: ins.PC})
				if err := jump(dest); err != nil {
					return nil, err
				}
			case ins.Op == vm.MSTORE:
				if err := need(2); err != nil {
					return nil, err
				}
				offset, value := pop(), pop()
				markLive(offset)
				markLive(value)
				if legacy {
					if err := legacyStore(&s, offset, value); err != nil {
						return nil, err
					}
					s.writes = true
					break
				}
				o, ok := constValue(offset)
				v, ok2 := constValue(value)
				if !ok || !ok2 || !o.IsUint64() || o.Uint64() > 4096 {
					return nil, fmt.Errorf("non-constant panic memory write")
				}
				word := v.FillBytes(make([]byte, 32))
				s.path.MemoryBytes = max(s.path.MemoryBytes, o.Uint64()+32)
				for i, b := range word {
					s.memory[o.Uint64()+uint64(i)] = b
				}
				s.writes = true
			case legacy && ins.Op == vm.CODECOPY:
				if err := need(3); err != nil {
					return nil, err
				}
				dest, source, size := pop(), pop(), pop()
				markLive(dest)
				markLive(source)
				markLive(size)
				src, ok := constValue(source)
				length, ok2 := constValue(size)
				if !ok || !ok2 || source.OriginPC == nil || !src.IsUint64() || !length.IsUint64() || length.Uint64() > 4096 || src.Uint64() > uint64(len(p.Code)) || length.Uint64() > uint64(len(p.Code))-src.Uint64() {
					return nil, fmt.Errorf("unsupported failure CODECOPY range")
				}
				relative, at, err := memoryAddress(dest)
				if err != nil {
					return nil, err
				}
				data := p.Code[src.Uint64() : src.Uint64()+length.Uint64()]
				encoded := hex.EncodeToString(data)
				if old, exists := analysis.CodeDataProducers[*source.OriginPC]; exists && old != encoded {
					return nil, fmt.Errorf("inconsistent code-data reference")
				}
				analysis.CodeDataProducers[*source.OriginPC] = encoded
				mem := s.memory
				if relative {
					mem = s.relativeMemory
				}
				for i, b := range data {
					mem[at+uint64(i)] = b
				}
				s.writes = true
			case legacy && ins.Op == vm.MLOAD:
				if err := need(1); err != nil {
					return nil, err
				}
				offset := pop()
				markLive(offset)
				v, err := legacyLoad(&s, offset)
				if err != nil {
					return nil, err
				}
				s.stack = append(s.stack, v)
			case legacy && ins.Op == vm.INVALID:
				s.path.Outcome = "invalid"
				analysis.Paths = append(analysis.Paths, s.path)
				terminal = true
			case ins.Op == vm.REVERT:
				if err := need(2); err != nil {
					return nil, err
				}
				offset, size := pop(), pop()
				markLive(offset)
				markLive(size)
				if legacy {
					data, mask, err := legacyReturnData(&s, offset, size)
					if err != nil {
						return nil, err
					}
					s.path.Outcome, s.path.FailureData = "revert", hex.EncodeToString(data)
					if mask != nil {
						s.path.FailureMask = hex.EncodeToString(mask)
						s.path.FailureMemory = offset
					}
					analysis.Paths = append(analysis.Paths, s.path)
					terminal = true
					break
				}
				o, ok := constValue(offset)
				z, ok2 := constValue(size)
				if !ok || !ok2 || !o.IsUint64() || z.Cmp(big.NewInt(36)) != 0 || o.Uint64() > 4096 {
					return nil, fmt.Errorf("unsupported failure return")
				}
				data := make([]byte, 36)
				for i := range data {
					v, written := s.memory[o.Uint64()+uint64(i)]
					if !written {
						return nil, fmt.Errorf("panic bytes depend on incoming memory")
					}
					data[i] = v
				}
				outcome := panicOutcome(hex.EncodeToString(data))
				if outcome == "" {
					return nil, fmt.Errorf("failure path is not arithmetic Panic(0x11/0x12)")
				}
				s.path.Outcome = outcome
				s.path.MemoryBytes = max(s.path.MemoryBytes, o.Uint64()+36)
				analysis.Paths = append(analysis.Paths, s.path)
				terminal = true
			default:
				arity := 2
				switch ins.Op {
				case vm.ADD, vm.SUB, vm.MUL, vm.DIV, vm.MOD, vm.SDIV, vm.SMOD, vm.SLT, vm.SGT, vm.SIGNEXTEND, vm.SAR, vm.LT, vm.GT, vm.EQ, vm.AND, vm.OR, vm.XOR, vm.SHL, vm.SHR, vm.EXP:
				case vm.NOT, vm.ISZERO:
					arity = 1
				default:
					return nil, fmt.Errorf("unsupported or effectful opcode %s at PC %d", ins.Op, ins.PC)
				}
				if err := need(arity); err != nil {
					return nil, err
				}
				args := []*Expr{pop()}
				if arity == 2 {
					args = append(args, pop())
				}
				value := makeExpr(ins.Op.String(), args...)
				if legacy {
					if coefficient, offset, ok := memoryAffine(value); ok && coefficient == 0 && value.Op != "const" {
						folded := constant(offset, nil)
						folded.constantOrigins = expressionOrigins(value)
						value = folded
					}
				}
				s.stack = append(s.stack, value)
			}
			s.path.PeakStack = max(s.path.PeakStack, len(s.stack))
			if legacy {
				for _, e := range s.stack {
					if expressionSize(e, 2048) > 2048 {
						return nil, fmt.Errorf("legacy expression complexity budget exceeded")
					}
				}
			}
			s.pc = next
		}
	}
	successes, failures := 0, 0
	for _, path := range analysis.Paths {
		if path.Outcome == "return" {
			successes++
		} else {
			failures++
		}
	}
	if successes == 0 || failures == 0 || analysis.Operation == "" {
		return nil, fmt.Errorf("helper does not include both success and overflow failure paths")
	}
	for _, pc := range analysis.ReachablePCs {
		if p.ByPC[pc].Op.IsPush() && !livePushes[pc] {
			analysis.DeadPushes[pc] = true
		}
	}
	return analysis, nil
}

func expressionOrigins(e *Expr) []uint64 {
	var origins []uint64
	seen := map[*Expr]bool{}
	var visit func(*Expr)
	visit = func(e *Expr) {
		if e == nil || seen[e] {
			return
		}
		seen[e] = true
		origins = append(origins, e.constantOrigins...)
		if e.OriginPC != nil {
			origins = append(origins, *e.OriginPC)
		}
		for _, arg := range e.Args {
			visit(arg)
		}
	}
	visit(e)
	return origins
}
func cloneState(s symbolicState) symbolicState {
	c := s
	c.stack = append([]*Expr(nil), s.stack...)
	c.memory = map[uint64]byte{}
	for k, v := range s.memory {
		c.memory[k] = v
	}
	c.relativeMemory = map[uint64]byte{}
	for k, v := range s.relativeMemory {
		c.relativeMemory[k] = v
	}
	c.path.PCs = append([]uint64(nil), s.path.PCs...)
	c.path.Conditions = append([]Condition(nil), s.path.Conditions...)
	return c
}

// CheckInputs compares the complete extracted path guards with an independent
// unbounded-integer specification. It is evidence over supplied inputs only.
func CheckInputs(a *SymbolicAnalysis, operation string, left, right *big.Int) error {
	typ, err := parseIntegerType(a.IntegerType)
	if err != nil {
		return err
	}
	return checkTypedInputs(a, operation, typ, left, right)
}

func checkTypedInputs(a *SymbolicAnalysis, operation string, typ integerType, left, right *big.Int) error {
	if len(a.OperandOrder) != 2 {
		return fmt.Errorf("missing operand order")
	}
	l, err := typ.decode(left)
	if err != nil {
		return err
	}
	r, err := typ.decode(right)
	if err != nil {
		return err
	}
	inputs := map[string]*big.Int{a.OperandOrder[0]: left, a.OperandOrder[1]: right}
	exact, outcome, err := arithmeticSpec(operation, typ, l, r)
	if err != nil {
		return err
	}
	matches := 0
	for _, path := range a.Paths {
		holds := true
		for _, c := range path.Conditions {
			v, err := c.Expression.Evaluate(inputs)
			if err != nil {
				return err
			}
			if (v.Sign() != 0) != c.Nonzero {
				holds = false
				break
			}
		}
		if !holds {
			continue
		}
		matches++
		if path.Outcome != outcome && !(isLegacySemantics(a.Semantics) && outcome != "return" && (path.Outcome == "revert" || path.Outcome == "invalid" || path.Outcome == "revert_argument")) {
			return fmt.Errorf("guard disagrees with arithmetic specification for %s,%s", left, right)
		}
		if outcome == "return" {
			result, err := path.Result.Evaluate(inputs)
			if err != nil {
				return err
			}
			if result.Cmp(new(big.Int).Mod(exact, modulus)) != 0 {
				return fmt.Errorf("result mismatch")
			}
		}
	}
	if matches != 1 {
		return fmt.Errorf("expected exactly one applicable path, got %d", matches)
	}
	return nil
}
