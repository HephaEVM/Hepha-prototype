package vm

// A binding-time decision tree and hash-consed expression DAG. Both are rebuilt
// from the deployed bytecode analysis, not trusted from a template bundle. This
// shares arithmetic between guards and results without removing any predicate.
import (
	"fmt"
	"strings"

	"github.com/holiman/uint256"
	miner "go-eth1.16.5-evm/template_miner"
)

type arithmeticNode struct {
	op       OpCode
	x, y     uint8
	constant uint256.Int
}

// A worker needs the exact decision-tree leaf, not a second returned value.
type arithmeticGuardFunc func(a, b uint256.Int) int
type arithmeticDecision struct {
	value               uint8
	zero, nonzero, path int
}
type arithmeticPlan struct {
	native        arithmeticNativeFunc
	nativeHot     arithmeticNativeFunc
	nativeReduced arithmeticNativeFunc
	guardHot      arithmeticGuardFunc
	guardReduced  arithmeticGuardFunc
	nodes         []arithmeticNode
	decisions     []arithmeticDecision
	results       []uint8
	root          int
}
type arithmeticScratch struct {
	words [64]uint256.Int
	valid uint64
}

func buildArithmeticPlan(paths []arithmeticPath) (*arithmeticPlan, error) {
	p := &arithmeticPlan{nodes: make([]arithmeticNode, 2), root: -1}
	ids := map[string]uint8{"a": 0, "b": 1}
	var add func(*miner.Expr) (uint8, string, error)
	add = func(e *miner.Expr) (uint8, string, error) {
		if e == nil {
			return 0, "", fmt.Errorf("missing expression")
		}
		if e.Op == "input" {
			id, ok := ids[e.Input]
			if !ok || (e.Input != "a" && e.Input != "b") {
				return 0, "", fmt.Errorf("unsupported input %s", e.Input)
			}
			return id, e.Input, nil
		}
		n := arithmeticNode{}
		key := e.Op
		if e.Op == "const" {
			v, err := uint256.FromHex(e.Value)
			if err != nil {
				return 0, "", err
			}
			n.op = PUSH32
			n.constant = *v
			key = v.Hex()
		} else {
			var op OpCode
			switch e.Op {
			case "ADD":
				op = ADD
			case "SUB":
				op = SUB
			case "MUL":
				op = MUL
			case "DIV":
				op = DIV
			case "MOD":
				op = MOD
			case "SDIV":
				op = SDIV
			case "SMOD":
				op = SMOD
			case "LT":
				op = LT
			case "GT":
				op = GT
			case "SLT":
				op = SLT
			case "SGT":
				op = SGT
			case "EQ":
				op = EQ
			case "ISZERO":
				op = ISZERO
			case "NOT":
				op = NOT
			case "AND":
				op = AND
			case "OR":
				op = OR
			case "XOR":
				op = XOR
			case "SHL":
				op = SHL
			case "SHR":
				op = SHR
			case "SAR":
				op = SAR
			case "SIGNEXTEND":
				op = SIGNEXTEND
			case "BYTE":
				op = BYTE
			case "EXP":
				op = EXP
			default:
				return 0, "", fmt.Errorf("unsupported arithmetic expression %s", e.Op)
			}
			n.op = op
			arity := 2
			if op == ISZERO || op == NOT {
				arity = 1
			}
			if len(e.Args) != arity {
				return 0, "", fmt.Errorf("expression arity")
			}
			keys := []string{}
			for i, a := range e.Args {
				id, k, err := add(a)
				if err != nil {
					return 0, "", err
				}
				if i == 0 {
					n.x = id
				} else {
					n.y = id
				}
				keys = append(keys, k)
			}
			key += "(" + strings.Join(keys, ",") + ")"
		}
		if id, ok := ids[key]; ok {
			return id, key, nil
		}
		if len(p.nodes) >= 64 {
			return 0, "", fmt.Errorf("expression DAG budget exceeded")
		}
		id := uint8(len(p.nodes))
		ids[key] = id
		p.nodes = append(p.nodes, n)
		return id, key, nil
	}
	var insert func(int, []miner.Condition, int) (int, error)
	insert = func(at int, cs []miner.Condition, path int) (int, error) {
		if len(cs) == 0 {
			if at >= 0 {
				return 0, fmt.Errorf("overlapping paths")
			}
			p.decisions = append(p.decisions, arithmeticDecision{path: path})
			return len(p.decisions) - 1, nil
		}
		id, _, err := add(cs[0].Expression)
		if err != nil {
			return 0, err
		}
		if at < 0 {
			at = len(p.decisions)
			p.decisions = append(p.decisions, arithmeticDecision{value: id, zero: -1, nonzero: -1, path: -1})
		}
		if p.decisions[at].path >= 0 || p.decisions[at].value != id {
			return 0, fmt.Errorf("inconsistent path prefix")
		}
		child := p.decisions[at].zero
		if cs[0].Nonzero {
			child = p.decisions[at].nonzero
		}
		next, err := insert(child, cs[1:], path)
		if err != nil {
			return 0, err
		}
		if cs[0].Nonzero {
			p.decisions[at].nonzero = next
		} else {
			p.decisions[at].zero = next
		}
		return at, nil
	}
	for i, path := range paths {
		id := uint8(0)
		if path.path.Outcome == "return" {
			var err error
			id, _, err = add(path.path.Result)
			if err != nil {
				return nil, err
			}
		}
		p.results = append(p.results, id)
		var err error
		p.root, err = insert(p.root, path.path.Conditions, i)
		if err != nil {
			return nil, err
		}
	}
	for _, d := range p.decisions {
		if d.path < 0 && (d.zero < 0 || d.nonzero < 0) {
			return nil, fmt.Errorf("incomplete decision tree")
		}
	}
	signature := p.signature()
	p.native = arithmeticNativePlans[signature]
	p.nativeHot = arithmeticNativeBooleanPlans[signature]
	p.nativeReduced = arithmeticNativeReducedPlans[signature]
	p.guardHot = arithmeticNativeBooleanPlansGuard[signature]
	p.guardReduced = arithmeticNativeReducedPlansGuard[signature]
	return p, nil
}

// Full structural equality is required, including path ordering and every
// guard. A template name or inferred arithmetic domain is not sufficient.
func (p *arithmeticPlan) signature() string {
	var b strings.Builder
	fmt.Fprintf(&b, "v1;r%d;", p.root)
	for _, n := range p.nodes {
		fmt.Fprintf(&b, "%d,%d,%d,%x;", n.op, n.x, n.y, n.constant)
	}
	b.WriteByte('|')
	for _, d := range p.decisions {
		fmt.Fprintf(&b, "%d,%d,%d,%d;", d.value, d.zero, d.nonzero, d.path)
	}
	fmt.Fprintf(&b, "|%v", p.results)
	return b.String()
}

type arithmeticNativeFunc func(a, b uint256.Int) (int, uint256.Int)

func (p *arithmeticPlan) value(id uint8, s *arithmeticScratch) *uint256.Int {
	z := &s.words[id]
	mask := uint64(1) << id
	if s.valid&mask != 0 {
		return z
	}
	n := &p.nodes[id]
	if n.op == PUSH32 {
		*z = n.constant
		s.valid |= mask
		return z
	}
	x := p.value(n.x, s)
	if n.op == NOT {
		z.Not(x)
	} else if n.op == ISZERO {
		z.Clear()
		if x.IsZero() {
			z.SetOne()
		}
	} else {
		y := p.value(n.y, s)
		switch n.op {
		case ADD:
			z.Add(x, y)
		case SUB:
			z.Sub(x, y)
		case MUL:
			z.Mul(x, y)
		case DIV:
			z.Div(x, y)
		case MOD:
			z.Mod(x, y)
		case SDIV:
			z.SDiv(x, y)
		case SMOD:
			z.SMod(x, y)
		case AND:
			z.And(x, y)
		case OR:
			z.Or(x, y)
		case XOR:
			z.Xor(x, y)
		case SIGNEXTEND:
			z.ExtendSign(y, x)
		case LT, GT, SLT, SGT, EQ:
			ok := false
			switch n.op {
			case LT:
				ok = x.Lt(y)
			case GT:
				ok = x.Gt(y)
			case SLT:
				ok = x.Slt(y)
			case SGT:
				ok = x.Sgt(y)
			case EQ:
				ok = x.Eq(y)
			}
			z.Clear()
			if ok {
				z.SetOne()
			}
		case SHL, SHR, SAR:
			shift := uint(256)
			if x.IsUint64() && x.Uint64() < 256 {
				shift = uint(x.Uint64())
			}
			if n.op == SHL {
				z.Lsh(y, shift)
			} else if n.op == SHR {
				z.Rsh(y, shift)
			} else {
				z.SRsh(y, shift)
			}
		case BYTE:
			z.Set(y)
			z.Byte(x)
		case EXP:
			z.Exp(x, y)
		}
	}
	s.valid |= mask
	return z
}

func (p *arithmeticPlan) selectPath(a, b uint256.Int, s *arithmeticScratch) int {
	s.words[0], s.words[1], s.valid = a, b, 3
	at := p.root
	for p.decisions[at].path < 0 {
		d := &p.decisions[at]
		if p.value(d.value, s).IsZero() {
			at = d.zero
		} else {
			at = d.nonzero
		}
	}
	return p.decisions[at].path
}
