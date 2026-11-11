package templateminer

// FullMath recognition compares complete success expressions and guards with
// the 512-bit division algorithm. Dynamic samples and prefix matches are never
// certificates. Failure paths remain the original EVM's responsibility.
import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"sort"
	"strings"

	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/holiman/uint256"
)

type MulDivRegion struct {
	Entry        uint64
	A, B, D      int
	Short, Wide  ABIRegionPath
	StructuralID string
}

type wordForm struct {
	key   string
	op    string
	args  []*wordForm
	value uint256.Int
}

func wordConst(x uint256.Int) *wordForm {
	return &wordForm{key: x.Hex(), op: "const", value: x}
}
func wordSmall(x uint64) *wordForm { return wordConst(*uint256.NewInt(x)) }
func wordInput(i int) *wordForm    { return &wordForm{key: fmt.Sprintf("input:%d", i), op: "input"} }
func wordOp(op string, args ...*wordForm) *wordForm {
	if len(args) == 0 {
		return nil
	}
	all := true
	for _, a := range args {
		if a == nil {
			return nil
		}
		all = all && a.op == "const"
	}
	if all {
		var z uint256.Int
		a := &args[0].value
		var b, c *uint256.Int
		if len(args) > 1 {
			b = &args[1].value
		}
		if len(args) > 2 {
			c = &args[2].value
		}
		switch op {
		case "ADD":
			z.Add(a, b)
		case "SUB":
			z.Sub(a, b)
		case "MUL":
			z.Mul(a, b)
		case "DIV":
			z.Div(a, b)
		case "MOD":
			z.Mod(a, b)
		case "MULMOD":
			z.MulMod(a, b, c)
		case "ADDMOD":
			z.AddMod(a, b, c)
		case "AND":
			z.And(a, b)
		case "OR":
			z.Or(a, b)
		case "XOR":
			z.Xor(a, b)
		case "NOT":
			z.Not(a)
		case "ISZERO":
			if a.IsZero() {
				z.SetOne()
			}
		case "EQ":
			if a.Eq(b) {
				z.SetOne()
			}
		case "LT":
			if a.Lt(b) {
				z.SetOne()
			}
		case "GT":
			if a.Gt(b) {
				z.SetOne()
			}
		case "SHL":
			if a.IsUint64() && a.Uint64() < 256 {
				z.Lsh(b, uint(a.Uint64()))
			}
		case "SHR":
			if a.IsUint64() && a.Uint64() < 256 {
				z.Rsh(b, uint(a.Uint64()))
			}
		default:
			return nil
		}
		return wordConst(z)
	}
	if op == "GT" {
		return wordOp("LT", args[1], args[0])
	}
	if len(args) == 2 {
		a, b := args[0], args[1]
		if b.op == "const" && b.value.IsZero() && (op == "ADD" || op == "SUB" || op == "OR" || op == "XOR") {
			return a
		}
		if a.op == "const" && a.value.IsZero() && (op == "ADD" || op == "OR" || op == "XOR") {
			return b
		}
		if op == "MUL" {
			if a.op == "const" && a.value.IsZero() || b.op == "const" && b.value.IsZero() {
				return wordSmall(0)
			}
			if a.op == "const" && a.value.IsUint64() && a.value.Uint64() == 1 {
				return b
			}
			if b.op == "const" && b.value.IsUint64() && b.value.Uint64() == 1 {
				return a
			}
		}
		if op == "SUB" && a.key == b.key {
			return wordSmall(0)
		}
	}
	args = append([]*wordForm(nil), args...)
	switch op {
	case "ADD", "MUL", "AND", "OR", "XOR", "EQ", "MULMOD", "ADDMOD":
		if args[0].key > args[1].key {
			args[0], args[1] = args[1], args[0]
		}
	}
	var s strings.Builder
	s.WriteString(op)
	for _, a := range args {
		s.WriteByte('/')
		s.WriteString(a.key)
	}
	h := sha256.Sum256([]byte(s.String()))
	return &wordForm{key: hex.EncodeToString(h[:]), op: op, args: args}
}
func wordGuard(x *wordForm, nonzero bool) string {
	if x == nil {
		return "unsupported"
	}
	for x.op == "ISZERO" {
		x = x.args[0]
		nonzero = !nonzero
	}
	// For unsigned words, 0 < x and x != 0 are identical.
	if x.op == "LT" && x.args[0].op == "const" && x.args[0].value.IsZero() {
		return wordGuard(x.args[1], nonzero)
	}
	if x.op == "const" {
		if (!x.value.IsZero()) == nonzero {
			return "true"
		}
		return "false"
	}
	return fmt.Sprintf("%t:%s", nonzero, x.key)
}
func wordSpec(a, b, d int) (short, wide *wordForm, sg, wg []string) {
	x, y, z := wordInput(a), wordInput(b), wordInput(d)
	var all uint256.Int
	all.SetAllOne()
	low := wordOp("MUL", x, y)
	mm := wordOp("MULMOD", x, y, wordConst(all))
	high := wordOp("SUB", wordOp("SUB", mm, low), wordOp("LT", mm, low))
	rem := wordOp("MULMOD", x, y, z)
	twos := wordOp("AND", wordOp("SUB", wordSmall(0), z), z)
	odd := wordOp("DIV", z, twos)
	inv := wordOp("XOR", wordOp("MUL", wordSmall(3), odd), wordSmall(2))
	for i := 0; i < 6; i++ {
		inv = wordOp("MUL", inv, wordOp("SUB", wordSmall(2), wordOp("MUL", odd, inv)))
	}
	newHigh := wordOp("SUB", high, wordOp("GT", rem, low))
	factor := wordOp("ADD", wordOp("DIV", wordOp("SUB", wordSmall(0), twos), twos), wordSmall(1))
	numerator := wordOp("OR", wordOp("DIV", wordOp("SUB", low, rem), twos), wordOp("MUL", newHigh, factor))
	short = wordOp("DIV", low, z)
	wide = wordOp("MUL", numerator, inv)
	sg = []string{wordGuard(high, false), wordGuard(z, true)}
	wg = []string{wordGuard(high, true), wordGuard(wordOp("GT", z, high), true)}
	sort.Strings(sg)
	sort.Strings(wg)
	return
}
func mulDivPathShape(p ABIRegionPath) (*wordForm, []string, bool) {
	if p.Outcome != "return" || len(p.Outputs) != 1 {
		return nil, nil, false
	}
	memo := make(map[int]*wordForm)
	var form func(int) *wordForm
	form = func(id int) *wordForm {
		if x, ok := memo[id]; ok {
			return x
		}
		if id < 0 || id >= len(p.Nodes) {
			return nil
		}
		n := p.Nodes[id]
		var x *wordForm
		switch n.Op {
		case "input":
			if n.Input == 0 {
				return nil
			}
			x = wordInput(n.Input)
		case "const":
			var w uint256.Int
			if w.SetFromHex(n.Value) != nil {
				return nil
			}
			x = wordConst(w)
		default:
			a := make([]*wordForm, len(n.Args))
			for i, j := range n.Args {
				a[i] = form(j)
			}
			x = wordOp(n.Op, a...)
		}
		memo[id] = x
		return x
	}
	var guards []string
	for _, a := range p.Actions {
		if a.Op != "guard" || len(a.Args) != 1 {
			return nil, nil, false
		}
		g := wordGuard(form(a.Args[0]), a.Nonzero)
		if g == "unsupported" || g == "false" {
			return nil, nil, false
		}
		if g != "true" {
			guards = append(guards, g)
		}
	}
	sort.Strings(guards)
	return form(p.Outputs[0]), guards, true
}
func AnalyzeMulDiv(p *Program, entry uint64) (*MulDivRegion, error) {
	r, e := AnalyzeABIRegion(p, entry, 4)
	if e != nil {
		return nil, e
	}
	var paths []ABIRegionPath
	for _, path := range r.Paths {
		if path.Outcome == "return" {
			paths = append(paths, path)
		}
	}
	if len(paths) != 2 {
		return nil, fmt.Errorf("expected exactly two pure success paths")
	}
	for _, perm := range [][3]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}} {
		short, wide, sg, wg := wordSpec(perm[0], perm[1], perm[2])
		out := &MulDivRegion{Entry: entry, A: perm[0], B: perm[1], D: perm[2], StructuralID: r.ID}
		for _, path := range paths {
			x, g, ok := mulDivPathShape(path)
			if !ok || x == nil {
				continue
			}
			if x.key == short.key && strings.Join(g, ",") == strings.Join(sg, ",") {
				out.Short = path
			}
			if x.key == wide.key && strings.Join(g, ",") == strings.Join(wg, ",") {
				out.Wide = path
			}
		}
		if len(out.Short.PCs) > 0 && len(out.Wide.PCs) > 0 {
			return out, nil
		}
	}
	return nil, fmt.Errorf("success expressions/guards do not implement full-product floor division")
}

func MulDivCandidates(p *Program) []uint64 {
	var out []uint64
	for i, in := range p.Instructions {
		if in.Op != vm.JUMPDEST {
			continue
		}
		mod, mul, sub := 0, 0, 0
		for j := i; j < len(p.Instructions) && j < i+45; j++ {
			op := p.Instructions[j].Op
			if op == vm.MULMOD {
				mod++
			}
			if op == vm.MUL {
				mul++
			}
			if op == vm.SUB {
				sub++
			}
			if op == vm.JUMP || op == vm.JUMPI || op == vm.STOP || op == vm.RETURN || op == vm.REVERT {
				break
			}
		}
		if mod > 0 && mul > 0 && sub >= 2 {
			out = append(out, in.PC)
		}
	}
	return out
}
