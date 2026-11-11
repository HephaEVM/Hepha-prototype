package templateminer

// ABI regions retain the complete data/control/effect program. Unlike scalar
// canonicality templates, this IR supports multiple returned values and ordered
// memory effects. A successful replacement is a translation, not an assumption
// that compiler/library checks always pass.
import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"
	"sort"
	"strings"

	"github.com/ethereum/go-ethereum/core/vm"
)

type ABIValue struct {
	Op     string  `json:"op"`
	Args   []int   `json:"args,omitempty"`
	Value  string  `json:"value,omitempty"`
	Input  int     `json:"input,omitempty"`
	Origin *uint64 `json:"origin,omitempty"`
}
type ABIAction struct {
	Op      string `json:"op"`
	Args    []int  `json:"args,omitempty"`
	Result  int    `json:"result,omitempty"`
	Nonzero bool   `json:"nonzero,omitempty"`
}
type ABIRegionPath struct {
	Nodes   []ABIValue  `json:"nodes"`
	Actions []ABIAction `json:"actions"`
	Outputs []int       `json:"outputs"`
	PCs     []uint64    `json:"pcs"`
	Peak    int         `json:"peak"`
	Outcome string      `json:"outcome"`
}
type ABIRegion struct {
	Entry    uint64           `json:"entry"`
	Inputs   int              `json:"inputs"`
	ID       string           `json:"id"`
	Category string           `json:"category"`
	Paths    []ABIRegionPath  `json:"paths"`
	Graph    []ABIRegionBlock `json:"graph,omitempty"`
}
type ABIRegionLibrary struct {
	Version   int               `json:"version"`
	Templates map[string]string `json:"templates"` // ID -> category
	Prefixes  map[string][]int  `json:"prefixes"`  // instruction prefix -> input slots
}
type abiRegionState struct {
	pc    uint64
	stack []int
	path  ABIRegionPath
	seen  map[string]bool
}

func abiPureArity(op vm.OpCode) int {
	switch op {
	case vm.ISZERO, vm.NOT:
		return 1
	case vm.MULMOD, vm.ADDMOD:
		return 3
	case vm.ADD, vm.SUB, vm.MUL, vm.DIV, vm.MOD, vm.SDIV, vm.SMOD, vm.LT, vm.GT, vm.SLT, vm.SGT, vm.EQ, vm.AND, vm.OR, vm.XOR, vm.SIGNEXTEND, vm.SHL, vm.SHR, vm.SAR, vm.BYTE:
		return 2
	}
	return 0
}
func (s *abiRegionState) clone() abiRegionState {
	c := *s
	c.stack = append([]int(nil), s.stack...)
	c.path.Nodes = append([]ABIValue(nil), s.path.Nodes...)
	c.path.Actions = append([]ABIAction(nil), s.path.Actions...)
	c.path.PCs = append([]uint64(nil), s.path.PCs...)
	c.seen = make(map[string]bool, len(s.seen))
	for k, v := range s.seen {
		c.seen[k] = v
	}
	return c
}
func (s *abiRegionState) node(v ABIValue) int {
	// Fold only literal EVM word expressions. Preserve PUSH origins for relocation.
	if v.Op != "const" && v.Op != "input" && len(v.Args) > 0 {
		args := make([]*Expr, 0, len(v.Args))
		all := true
		for _, i := range v.Args {
			n := s.path.Nodes[i]
			if n.Op != "const" {
				all = false
				break
			}
			args = append(args, &Expr{Op: "const", Value: n.Value})
		}
		if all {
			e := makeExpr(v.Op, args...)
			if e.Op == "const" {
				v = ABIValue{Op: "const", Value: e.Value}
			}
		}
	}
	s.path.Nodes = append(s.path.Nodes, v)
	return len(s.path.Nodes) - 1
}
func (s *abiRegionState) push(i int) {
	s.stack = append(s.stack, i)
	if len(s.stack) > s.path.Peak {
		s.path.Peak = len(s.stack)
	}
}
func (s *abiRegionState) pop() int {
	n := len(s.stack) - 1
	i := s.stack[n]
	s.stack = s.stack[:n]
	return i
}

// AnalyzeABIRegion exhaustively explores a bounded, closed helper. Unresolved
// jumps, storage/calls, gas observation, loops beyond the bound and unsupported
// effects reject the entire candidate. Failure paths are left to the real EVM.
func AnalyzeABIRegion(p *Program, entry uint64, inputs int) (*ABIRegion, error) {
	return analyzeABIRegion(p, entry, inputs, false)
}

// The arithmetic composition recognizer may retain INVALID as an exceptional
// leaf. Existing ABI mining keeps its original, more restrictive acceptance.
func analyzeABIRegion(p *Program, entry uint64, inputs int, allowInvalid bool) (*ABIRegion, error) {
	if inputs < 1 || inputs > 18 || p.ByPC[entry].Op != vm.JUMPDEST {
		return nil, fmt.Errorf("entry/arity")
	}
	s := abiRegionState{pc: entry, seen: map[string]bool{}}
	for i := 0; i < inputs; i++ {
		s.push(s.node(ABIValue{Op: "input", Input: i}))
	}
	queue := []abiRegionState{s}
	var paths []ABIRegionPath
	jumps := map[uint64]uint64{}
	for len(queue) > 0 {
		s = queue[len(queue)-1]
		queue = queue[:len(queue)-1]
		for {
			if len(s.path.PCs) >= 1024 || len(s.path.Nodes) >= 512 || len(paths)+len(queue) > 32 {
				return nil, fmt.Errorf("region bound/loop")
			}
			ins, ok := p.ByPC[s.pc]
			if !ok {
				return nil, fmt.Errorf("instruction boundary")
			}
			// Identical symbolic stack at the same PC is a cycle. Repeated calls to
			// cleanup with different continuations/arguments remain valid.
			key := fmt.Sprintf("%d:%v", s.pc, s.stack)
			if s.seen[key] {
				return nil, fmt.Errorf("cycle")
			}
			s.seen[key] = true
			s.path.PCs = append(s.path.PCs, s.pc)
			s.pc += uint64(1 + len(ins.Data))
			need := func(n int) bool { return len(s.stack) >= n }
			dest := func(i int) (uint64, error) {
				v := s.path.Nodes[i]
				if v.Op != "const" || v.Origin == nil {
					return 0, fmt.Errorf("unresolved jump")
				}
				n, ok := new(big.Int).SetString(v.Value[2:], 16)
				if !ok || !n.IsUint64() || p.ByPC[n.Uint64()].Op != vm.JUMPDEST {
					return 0, fmt.Errorf("jump destination")
				}
				jumps[*v.Origin] = n.Uint64()
				return n.Uint64(), nil
			}
			switch {
			case ins.Op >= vm.PUSH0 && ins.Op <= vm.PUSH32:
				pc := ins.PC
				s.push(s.node(ABIValue{Op: "const", Value: "0x" + new(big.Int).SetBytes(ins.Data).Text(16), Origin: &pc}))
			case ins.Op >= vm.DUP1 && ins.Op <= vm.DUP16:
				n := int(ins.Op-vm.DUP1) + 1
				if !need(n) {
					return nil, fmt.Errorf("stack underflow")
				}
				s.push(s.stack[len(s.stack)-n])
			case ins.Op >= vm.SWAP1 && ins.Op <= vm.SWAP16:
				n := int(ins.Op-vm.SWAP1) + 1
				if !need(n + 1) {
					return nil, fmt.Errorf("stack underflow")
				}
				last := len(s.stack) - 1
				s.stack[last], s.stack[last-n] = s.stack[last-n], s.stack[last]
			case ins.Op == vm.POP:
				if !need(1) {
					return nil, fmt.Errorf("stack underflow")
				}
				s.pop()
			case ins.Op == vm.JUMPDEST:
			case abiPureArity(ins.Op) > 0:
				n := abiPureArity(ins.Op)
				if !need(n) {
					return nil, fmt.Errorf("stack underflow")
				}
				args := make([]int, n)
				for i := range args {
					args[i] = s.pop()
				}
				s.push(s.node(ABIValue{Op: ins.Op.String(), Args: args}))
			case ins.Op == vm.CALLDATALOAD || ins.Op == vm.MLOAD:
				if !need(1) {
					return nil, fmt.Errorf("stack underflow")
				}
				a := s.pop()
				r := s.node(ABIValue{Op: "read"})
				s.path.Actions = append(s.path.Actions, ABIAction{Op: ins.Op.String(), Args: []int{a}, Result: r})
				s.push(r)
			case ins.Op == vm.CALLDATASIZE || ins.Op == vm.MSIZE:
				r := s.node(ABIValue{Op: "read"})
				s.path.Actions = append(s.path.Actions, ABIAction{Op: ins.Op.String(), Result: r})
				s.push(r)
			case ins.Op == vm.MSTORE || ins.Op == vm.MSTORE8 || ins.Op == vm.CALLDATACOPY:
				n := 2
				if ins.Op == vm.CALLDATACOPY {
					n = 3
				}
				if !need(n) {
					return nil, fmt.Errorf("stack underflow")
				}
				args := make([]int, n)
				for i := range args {
					args[i] = s.pop()
				}
				s.path.Actions = append(s.path.Actions, ABIAction{Op: ins.Op.String(), Args: args})
			case ins.Op == vm.JUMP:
				if !need(1) {
					return nil, fmt.Errorf("stack underflow")
				}
				a := s.pop()
				v := s.path.Nodes[a]
				if v.Op == "input" && v.Input == 0 {
					s.path.Outcome = "return"
					s.path.Outputs = append([]int(nil), s.stack...)
					paths = append(paths, s.path)
					goto next
				}
				n, err := dest(a)
				if err != nil {
					return nil, err
				}
				s.pc = n
			case ins.Op == vm.JUMPI:
				if !need(2) {
					return nil, fmt.Errorf("stack underflow")
				}
				a, c := s.pop(), s.pop()
				n, err := dest(a)
				if err != nil {
					return nil, err
				}
				v := s.path.Nodes[c]
				if v.Op == "const" {
					if v.Value != "0x0" {
						s.pc = n
					}
					continue
				}
				branch := s.clone()
				branch.pc = n
				branch.path.Actions = append(branch.path.Actions, ABIAction{Op: "guard", Args: []int{c}, Nonzero: true})
				queue = append(queue, branch)
				s.path.Actions = append(s.path.Actions, ABIAction{Op: "guard", Args: []int{c}})
			case ins.Op == vm.REVERT:
				if !need(2) {
					return nil, fmt.Errorf("stack underflow")
				}
				s.path.Outcome = "revert"
				paths = append(paths, s.path)
				goto next
			case ins.Op == vm.INVALID && allowInvalid:
				s.path.Outcome = "invalid"
				paths = append(paths, s.path)
				goto next
			default:
				return nil, fmt.Errorf("unsupported %s", ins.Op)
			}
		}
	next:
	}
	success := false
	for _, path := range paths {
		if path.Outcome == "return" {
			success = true
			if len(path.Outputs) > 18 {
				return nil, fmt.Errorf("output arity")
			}
		}
	}
	if !success {
		return nil, fmt.Errorf("no successful return")
	}
	labels := map[uint64]int{}
	for _, path := range paths {
		for _, pc := range path.PCs {
			if _, ok := labels[pc]; !ok {
				labels[pc] = len(labels)
			}
		}
	}
	var normalized []string
	for _, path := range paths {
		normalized = append(normalized, path.Outcome)
		for _, pc := range path.PCs {
			i := p.ByPC[pc]
			v := i.Op.String()
			if i.Op >= vm.PUSH0 && i.Op <= vm.PUSH32 {
				if target, ok := jumps[pc]; ok {
					v = fmt.Sprintf("PUSH L%d", labels[target])
				} else {
					v = "PUSH 0x" + new(big.Int).SetBytes(i.Data).Text(16)
				}
			}
			normalized = append(normalized, fmt.Sprintf("L%d %s", labels[pc], v))
		}
	}
	b, _ := json.Marshal(struct {
		Inputs       int
		Instructions []string
	}{inputs, normalized})
	h := sha256.Sum256(b)
	return &ABIRegion{Entry: entry, Inputs: inputs, ID: hex.EncodeToString(h[:]), Paths: paths}, nil
}

func abiRegionPrefix(p *Program, pc uint64) string {
	var ops []string
	for i := 0; i < 8; i++ {
		in, ok := p.ByPC[pc]
		if !ok {
			break
		}
		op := in.Op.String()
		if in.Op >= vm.PUSH0 && in.Op <= vm.PUSH32 {
			op = "PUSH"
		}
		ops = append(ops, op)
		pc += uint64(1 + len(in.Data))
		if in.Op == vm.JUMP || in.Op == vm.JUMPI || in.Op == vm.REVERT {
			break
		}
	}
	return strings.Join(ops, " ")
}
func abiRegionCategory(name string) string {
	switch {
	case strings.HasPrefix(name, "abi_decode_"):
		return "decode"
	case strings.HasPrefix(name, "abi_encode_"):
		return "encode"
	case name == "allocate_memory" || name == "finalize_allocation":
		return "allocate"
	case strings.HasPrefix(name, "copy_") && strings.Contains(name, "_to_memory"):
		return "copy"
	}
	return ""
}
func ABIRegionTraining(paths []string) (*ABIRegionLibrary, []map[string]any, error) {
	l := &ABIRegionLibrary{Version: 1, Templates: map[string]string{}, Prefixes: map[string][]int{}}
	var labels []map[string]any
	for _, file := range paths {
		a, err := readArtifactForABI(file)
		if err != nil {
			return nil, nil, err
		}
		for _, src := range sortedKeys(a.Output.Contracts) {
			for _, name := range sortedKeys(a.Output.Contracts[src]) {
				p, err := Load(file, src, name)
				if err != nil {
					return nil, nil, err
				}
				for _, fn := range sortedKeys(p.Contract.EVM.DeployedBytecode.FunctionDebugData) {
					d := p.Contract.EVM.DeployedBytecode.FunctionDebugData[fn]
					kind := abiRegionCategory(fn)
					if kind == "" || d.EntryPoint == nil {
						continue
					}
					label := map[string]any{"artifact": file, "contract": name, "compiler": a.CompilerVersion, "function": fn, "entry": *d.EntryPoint, "category": kind}
					r, err := AnalyzeABIRegion(p, *d.EntryPoint, d.ParameterSlots+1)
					if err != nil {
						r, err = AnalyzeABIGraph(p, *d.EntryPoint, d.ParameterSlots+1)
					}
					if err != nil {
						label["rejected"] = err.Error()
					} else {
						valid := true
						for _, path := range r.Paths {
							if path.Outcome == "return" && len(path.Outputs) != d.ReturnSlots {
								valid = false
							}
						}
						for _, block := range r.Graph {
							if block.Path.Outcome == "return" && len(block.Path.Outputs) != d.ReturnSlots {
								valid = false
							}
						}
						if !valid {
							label["rejected"] = "compiler return arity mismatch"
						} else {
							r.Category = kind
							label["region"] = r
							l.Templates[r.ID] = kind
							prefix := abiRegionPrefix(p, r.Entry)
							found := false
							for _, n := range l.Prefixes[prefix] {
								if n == r.Inputs {
									found = true
								}
							}
							if !found {
								l.Prefixes[prefix] = append(l.Prefixes[prefix], r.Inputs)
								sort.Ints(l.Prefixes[prefix])
							}
						}
					}
					labels = append(labels, label)
				}
			}
		}
	}
	return l, labels, nil
}
func DiscoverABIRegions(code []byte, l *ABIRegionLibrary, yields ...func()) []*ABIRegion {
	var pause func()
	if len(yields) > 0 {
		pause = yields[0]
	}
	ins, err := Disassemble(code)
	if err != nil {
		return nil
	}
	p := &Program{Code: code, Instructions: ins, ByPC: map[uint64]Instruction{}}
	for _, i := range ins {
		p.ByPC[i.PC] = i
	}
	var out []*ABIRegion
	for _, i := range ins {
		if i.Op != vm.JUMPDEST {
			continue
		}
		for _, n := range l.Prefixes[abiRegionPrefix(p, i.PC)] {
			if pause != nil {
				pause()
			}
			r, err := AnalyzeABIRegion(p, i.PC, n)
			if err != nil {
				r, err = AnalyzeABIGraph(p, i.PC, n)
			}
			if err != nil {
				continue
			}
			if kind, ok := l.Templates[r.ID]; ok {
				r.Category = kind
				out = append(out, r)
				break
			}
		}
	}
	return out
}
