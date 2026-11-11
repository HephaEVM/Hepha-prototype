package templateminer

// Scalar ABI pilot. Compiler labels select training examples; deployment binding
// re-analyzes the whole closed region and matches its normalized structure.
import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/core/vm"
)

type ABIScalarRegion struct {
	Entry      uint64          `json:"entry"`
	Inputs     int             `json:"inputs"`
	Offset     int             `json:"offsetFromBottom"`
	Kind       string          `json:"kind"` // word, mask, sign, bool
	Mask       string          `json:"mask,omitempty"`
	SignByte   uint64          `json:"signByte,omitempty"`
	Peak       int             `json:"peak"`
	SuccessPCs []uint64        `json:"successPcs"`
	TemplateID string          `json:"templateId"`
	Paths      []ABIScalarPath `json:"paths"`
}
type ABIScalarPath struct {
	PCs        []uint64    `json:"pcs"`
	Conditions []Condition `json:"conditions,omitempty"`
	Outcome    string      `json:"outcome"`
	Output     *Expr       `json:"output,omitempty"`
	Peak       int         `json:"peak"`
}
type abiScalarState struct {
	pc    uint64
	stack []*Expr
	path  ABIScalarPath
	loads int
}

func abiExprKey(e *Expr) string {
	if e == nil {
		return "nil"
	}
	if e.Op == "const" {
		return e.Value
	}
	if e.Op == "input" {
		return e.Input
	}
	x := e.Op + "("
	for _, a := range e.Args {
		x += abiExprKey(a) + ","
	}
	return x + ")"
}
func abiSame(a, b *Expr) bool { return abiExprKey(a) == abiExprKey(b) }

// AnalyzeABIScalar admits only a single-word decoder, a complete single success
// path, and at most one complementary empty-revert path. No unknown memory,
// storage, calls, loops, or unconstrained jumps are allowed.
func AnalyzeABIScalar(p *Program, entry uint64, inputs int) (*ABIScalarRegion, error) {
	if inputs < 2 || inputs > 3 || p.ByPC[entry].Op != vm.JUMPDEST {
		return nil, fmt.Errorf("not scalar entry")
	}
	s := abiScalarState{pc: entry}
	for i := 0; i < inputs; i++ {
		s.stack = append(s.stack, variable(fmt.Sprintf("v%d", i)))
	}
	s.path.Peak = inputs
	queue := []abiScalarState{s}
	var paths []ABIScalarPath
	jumps := map[uint64]uint64{}
	for len(queue) > 0 {
		s = queue[len(queue)-1]
		queue = queue[:len(queue)-1]
		for {
			if len(s.path.PCs) >= 128 || len(paths)+len(queue) > 4 {
				return nil, fmt.Errorf("ABI region bound")
			}
			ins, ok := p.ByPC[s.pc]
			if !ok {
				return nil, fmt.Errorf("not instruction boundary")
			}
			for _, old := range s.path.PCs {
				if old == s.pc {
					return nil, fmt.Errorf("loop")
				}
			}
			s.path.PCs = append(s.path.PCs, s.pc)
			s.pc += uint64(1 + len(ins.Data))
			need := func(n int) bool { return len(s.stack) >= n }
			pop := func() *Expr { n := len(s.stack) - 1; v := s.stack[n]; s.stack = s.stack[:n]; return v }
			push := func(v *Expr) {
				s.stack = append(s.stack, v)
				if len(s.stack) > s.path.Peak {
					s.path.Peak = len(s.stack)
				}
			}
			dest := func(e *Expr) (uint64, error) {
				n, ok := constValue(e)
				if !ok || !n.IsUint64() || p.ByPC[n.Uint64()].Op != vm.JUMPDEST || e.OriginPC == nil {
					return 0, fmt.Errorf("unresolved jump")
				}
				jumps[*e.OriginPC] = n.Uint64()
				return n.Uint64(), nil
			}
			switch {
			case ins.Op >= vm.PUSH0 && ins.Op <= vm.PUSH32:
				pc := ins.PC
				push(constant(new(big.Int).SetBytes(ins.Data), &pc))
			case ins.Op >= vm.DUP1 && ins.Op <= vm.DUP16:
				n := int(ins.Op-vm.DUP1) + 1
				if !need(n) {
					return nil, fmt.Errorf("stack underflow")
				}
				push(s.stack[len(s.stack)-n])
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
				pop()
			case ins.Op == vm.JUMPDEST:
			case ins.Op == vm.CALLDATALOAD:
				if !need(1) || s.loads != 0 {
					return nil, fmt.Errorf("load count")
				}
				a := pop()
				push(makeExpr("CALLDATALOAD", a))
				s.loads++
			case ins.Op == vm.AND || ins.Op == vm.EQ || ins.Op == vm.SIGNEXTEND || ins.Op == vm.SHL || ins.Op == vm.SHR || ins.Op == vm.SUB || ins.Op == vm.ADD || ins.Op == vm.OR || ins.Op == vm.LT || ins.Op == vm.GT:
				if !need(2) {
					return nil, fmt.Errorf("stack underflow")
				}
				a, b := pop(), pop()
				push(makeExpr(ins.Op.String(), a, b))
			case ins.Op == vm.ISZERO || ins.Op == vm.NOT:
				if !need(1) {
					return nil, fmt.Errorf("stack underflow")
				}
				push(makeExpr(ins.Op.String(), pop()))
			case ins.Op == vm.JUMP:
				if !need(1) {
					return nil, fmt.Errorf("stack underflow")
				}
				a := pop()
				if a.Op == "input" && a.Input == "v0" {
					if len(s.stack) != 1 || s.loads != 1 {
						return nil, fmt.Errorf("not one-word return")
					}
					s.path.Outcome = "return"
					s.path.Output = s.stack[0]
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
				a, c := pop(), pop()
				n, err := dest(a)
				if err != nil {
					return nil, err
				}
				if value, ok := constValue(c); ok {
					if value.Sign() != 0 {
						s.pc = n
					}
					continue
				}
				if len(s.path.Conditions) != 0 {
					return nil, fmt.Errorf("multiple dynamic guards")
				}
				branch := s
				branch.stack = append([]*Expr(nil), s.stack...)
				branch.path.PCs = append([]uint64(nil), s.path.PCs...)
				branch.path.Conditions = []Condition{{Expression: c, Nonzero: true}}
				branch.pc = n
				queue = append(queue, branch)
				s.path.Conditions = []Condition{{Expression: c, Nonzero: false}}
			case ins.Op == vm.REVERT:
				if !need(2) {
					return nil, fmt.Errorf("stack underflow")
				}
				a, b := pop(), pop()
				x, ok := constValue(a)
				y, ok2 := constValue(b)
				if !ok || !ok2 || x.Sign() != 0 || y.Sign() != 0 {
					return nil, fmt.Errorf("nonempty revert")
				}
				s.path.Outcome = "revert"
				paths = append(paths, s.path)
				goto next
			default:
				return nil, fmt.Errorf("unsupported %s", ins.Op)
			}
		}
	next:
	}
	var success *ABIScalarPath
	for i := range paths {
		if paths[i].Outcome == "return" {
			if success != nil {
				return nil, fmt.Errorf("multiple success paths")
			}
			success = &paths[i]
		}
	}
	if success == nil {
		return nil, fmt.Errorf("no return")
	}
	var load *Expr
	var find func(*Expr)
	find = func(e *Expr) {
		if e == nil {
			return
		}
		if e.Op == "CALLDATALOAD" {
			load = e
		}
		for _, a := range e.Args {
			find(a)
		}
	}
	find(success.Output)
	for _, c := range success.Conditions {
		find(c.Expression)
	}
	if load == nil || len(load.Args) != 1 || load.Args[0].Op != "input" {
		return nil, fmt.Errorf("not input offset")
	}
	offset := -1
	for i := 1; i < inputs; i++ {
		if load.Args[0].Input == fmt.Sprintf("v%d", i) {
			offset = i
		}
	}
	if offset < 0 {
		return nil, fmt.Errorf("return address used as offset")
	}
	r := &ABIScalarRegion{Entry: entry, Inputs: inputs, Offset: offset, Kind: "word", Peak: success.Peak, SuccessPCs: success.PCs, Paths: paths}
	if len(success.Conditions) == 0 {
		if len(paths) != 1 || !abiSame(success.Output, load) {
			return nil, fmt.Errorf("unconstrained transformation")
		}
	} else {
		if len(paths) != 2 {
			return nil, fmt.Errorf("not complementary paths")
		}
		c := success.Conditions[0]
		for c.Expression.Op == "ISZERO" {
			c.Expression = c.Expression.Args[0]
			c.Nonzero = !c.Nonzero
		}
		if c.Expression.Op != "EQ" || !c.Nonzero {
			return nil, fmt.Errorf("not canonicality equality")
		}
		a, b := c.Expression.Args[0], c.Expression.Args[1]
		if abiSame(b, load) {
			a, b = b, a
		}
		if !abiSame(a, load) || (!abiSame(success.Output, load) && !abiSame(success.Output, b)) {
			return nil, fmt.Errorf("not canonical word")
		}
		cleanupOp := b.Op
		if abiSame(b, load) {
			cleanupOp = "identity"
		}
		switch cleanupOp {
		case "identity":
			r.Kind = "word"
		case "AND":
			x, y := b.Args[0], b.Args[1]
			if abiSame(x, load) {
				x, y = y, x
			}
			mask, ok := constValue(x)
			if !ok || !abiSame(y, load) {
				return nil, fmt.Errorf("not mask")
			}
			r.Kind = "mask"
			r.Mask = "0x" + mask.Text(16)
		case "SIGNEXTEND":
			width, ok := constValue(b.Args[0])
			if !ok || !width.IsUint64() || width.Uint64() > 31 || !abiSame(b.Args[1], load) {
				return nil, fmt.Errorf("not sign extend")
			}
			r.Kind = "sign"
			r.SignByte = width.Uint64()
		case "ISZERO":
			if b.Args[0].Op != "ISZERO" || !abiSame(b.Args[0].Args[0], load) {
				return nil, fmt.Errorf("not boolean")
			}
			r.Kind = "bool"
		default:
			return nil, fmt.Errorf("unknown cleanup")
		}
	}
	// First-visit PC labels and proven jump producers remove relocation only.
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
			ins := p.ByPC[pc]
			v := ins.Op.String()
			if ins.Op >= vm.PUSH0 && ins.Op <= vm.PUSH32 {
				if target, ok := jumps[pc]; ok {
					v = fmt.Sprintf("PUSH L%d", labels[target])
				} else {
					v = "PUSH 0x" + new(big.Int).SetBytes(ins.Data).Text(16)
				}
			}
			normalized = append(normalized, fmt.Sprintf("L%d %s", labels[pc], v))
		}
	}
	encoded, _ := json.Marshal(struct {
		Inputs       int
		Instructions []string
	}{inputs, normalized})
	hash := sha256.Sum256(encoded)
	r.TemplateID = hex.EncodeToString(hash[:])
	return r, nil
}

func DiscoverABIScalars(code []byte, known map[string]bool, yields ...func()) []*ABIScalarRegion {
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
	var out []*ABIScalarRegion
	for idx, i := range ins {
		if i.Op != vm.JUMPDEST {
			continue
		}
		possible := false
		for j := idx; j < len(ins) && j < idx+16; j++ {
			op := ins[j].Op
			if op == vm.CALLDATALOAD {
				possible = true
				break
			}
			if op == vm.JUMP || op == vm.JUMPI || op == vm.REVERT || op == vm.STOP {
				break
			}
		}
		if !possible {
			continue
		}
		for _, n := range []int{2, 3} {
			if pause != nil {
				pause()
			}
			r, err := AnalyzeABIScalar(p, i.PC, n)
			if err == nil && (known == nil || known[r.TemplateID]) {
				out = append(out, r)
				break
			}
		}
	}
	return out
}

func ABIScalarTraining(artifactPaths []string) (map[string]bool, []map[string]any, error) {
	known := map[string]bool{}
	var labels []map[string]any
	for _, path := range artifactPaths {
		data, err := readArtifactForABI(path)
		if err != nil {
			return nil, nil, err
		}
		for _, source := range sortedKeys(data.Output.Contracts) {
			for _, name := range sortedKeys(data.Output.Contracts[source]) {
				p, err := Load(path, source, name)
				if err != nil {
					return nil, nil, err
				}
				for _, fn := range sortedKeys(p.Contract.EVM.DeployedBytecode.FunctionDebugData) {
					d := p.Contract.EVM.DeployedBytecode.FunctionDebugData[fn]
					if !strings.HasPrefix(fn, "abi_decode_t_") || d.EntryPoint == nil || d.ReturnSlots != 1 {
						continue
					}
					r, err := AnalyzeABIScalar(p, *d.EntryPoint, d.ParameterSlots+1)
					label := map[string]any{"artifact": path, "contract": name, "compiler": data.CompilerVersion, "function": fn, "entry": *d.EntryPoint}
					if err != nil {
						label["rejected"] = err.Error()
					} else {
						known[r.TemplateID] = true
						label["region"] = r
					}
					labels = append(labels, label)
				}
			}
		}
	}
	return known, labels, nil
}

func readArtifactForABI(path string) (*Artifact, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var a Artifact
	err = json.Unmarshal(b, &a)
	return &a, err
}
