package templateminer

// Integer domains are inferred from expression/guard semantics, not from the
// manual cleanup opcode arrays. Finite checks corroborate, but do not prove,
// the inferred specification. Narrow inputs must be canonically represented.
import (
	"crypto/sha256"
	"fmt"
	"math/big"
	"sort"
	"strconv"
	"strings"
)

type integerType struct {
	bits   int
	signed bool
}

func parseIntegerType(name string) (integerType, error) {
	if name == "" {
		name = "uint256"
	}
	t := integerType{signed: strings.HasPrefix(name, "int")}
	prefix := "uint"
	if t.signed {
		prefix = "int"
	}
	n, err := strconv.Atoi(strings.TrimPrefix(name, prefix))
	if err != nil || !strings.HasPrefix(name, prefix) || n < 8 || n > 256 || n%8 != 0 {
		return t, fmt.Errorf("invalid integer type %q", name)
	}
	t.bits = n
	return t, nil
}
func (t integerType) name() string {
	prefix := "uint"
	if t.signed {
		prefix = "int"
	}
	return fmt.Sprintf("%s%d", prefix, t.bits)
}
func (t integerType) bounds() (*big.Int, *big.Int) {
	min := new(big.Int)
	bits := t.bits
	if t.signed {
		bits--
		min.Neg(new(big.Int).Lsh(big.NewInt(1), uint(bits)))
	}
	max := new(big.Int).Sub(new(big.Int).Lsh(big.NewInt(1), uint(bits)), big.NewInt(1))
	return min, max
}
func signedWord(n *big.Int) *big.Int {
	x := new(big.Int).Mod(new(big.Int).Set(n), modulus)
	if x.Bit(255) != 0 {
		x.Sub(x, modulus)
	}
	return x
}
func (t integerType) decode(n *big.Int) (*big.Int, error) {
	x := new(big.Int).Set(n)
	if t.signed {
		x = signedWord(n)
	}
	lo, hi := t.bounds()
	if x.Cmp(lo) < 0 || x.Cmp(hi) > 0 {
		return nil, fmt.Errorf("input not canonical %s", t.name())
	}
	return x, nil
}
func binaryOperation(op string) bool {
	switch op {
	case "ADD", "SUB", "MUL", "DIV", "MOD", "SDIV", "SMOD":
		return true
	}
	return false
}
func panicOutcome(data string) string {
	data = strings.TrimPrefix(data, "0x")
	prefix := "4e487b71" + strings.Repeat("0", 62)
	if data == prefix+"11" {
		return "panic_0x11"
	}
	if data == prefix+"12" {
		return "panic_0x12"
	}
	return ""
}
func lowMaskBits(n *big.Int) int {
	if n.Sign() <= 0 {
		return 0
	}
	p := new(big.Int).Add(n, big.NewInt(1))
	if new(big.Int).And(p, new(big.Int).Sub(p, big.NewInt(1))).Sign() != 0 {
		return 0
	}
	bits := n.BitLen()
	if bits < 8 || bits > 256 || bits%8 != 0 {
		return 0
	}
	return bits
}

// Strip only well-defined integer cleanup operators; arbitrary arithmetic
// around the returned operation is not an acceptable binary helper shape.
func unwrapCleanup(e *Expr) *Expr {
	for e != nil && len(e.Args) == 2 {
		if e.Op == "SIGNEXTEND" {
			n, ok := constValue(e.Args[0])
			if ok && n.IsUint64() && n.Uint64() <= 31 {
				e = e.Args[1]
				continue
			}
		}
		if e.Op == "AND" {
			if n, ok := constValue(e.Args[0]); ok && lowMaskBits(n) != 0 {
				e = e.Args[1]
				continue
			}
			if n, ok := constValue(e.Args[1]); ok && lowMaskBits(n) != 0 {
				e = e.Args[0]
				continue
			}
		}
		break
	}
	return e
}
func resultOperation(e *Expr) (string, []string, bool) {
	e = unwrapCleanup(e)
	if e == nil || !binaryOperation(e.Op) || len(e.Args) != 2 {
		return "", nil, false
	}
	l, r := unwrapCleanup(e.Args[0]), unwrapCleanup(e.Args[1])
	if l.Op != "input" || r.Op != "input" || l.Input == r.Input || (l.Input != "a" && l.Input != "b") || (r.Input != "a" && r.Input != "b") {
		return "", nil, false
	}
	return e.Op, []string{l.Input, r.Input}, true
}

// Independent mathematical oracle: Solidity division truncates toward zero;
// signed remainder follows the dividend; min/-1 overflows only for division.
func arithmeticSpec(op string, t integerType, l, r *big.Int) (*big.Int, string, error) {
	z := new(big.Int)
	if (op == "SDIV" || op == "SMOD") != t.signed && (op == "DIV" || op == "MOD" || op == "SDIV" || op == "SMOD") {
		return nil, "", fmt.Errorf("operation/type signedness mismatch")
	}
	switch op {
	case "ADD":
		z.Add(l, r)
	case "SUB":
		z.Sub(l, r)
	case "MUL":
		z.Mul(l, r)
	case "DIV", "SDIV":
		if r.Sign() == 0 {
			return z, "panic_0x12", nil
		}
		z.Quo(l, r)
	case "MOD", "SMOD":
		if r.Sign() == 0 {
			return z, "panic_0x12", nil
		}
		z.Rem(l, r)
	default:
		return nil, "", fmt.Errorf("unsupported arithmetic %s", op)
	}
	lo, hi := t.bounds()
	if z.Cmp(lo) < 0 || z.Cmp(hi) > 0 {
		return z, "panic_0x11", nil
	}
	return z, "return", nil
}

func typedValidationPairs(t integerType) [][2]*big.Int {
	if !t.signed && t.bits == 256 {
		return validationPairs()
	}
	lo, hi := t.bounds()
	var values []*big.Int
	if t.signed {
		values = []*big.Int{lo, new(big.Int).Add(lo, big.NewInt(1)), big.NewInt(-2), big.NewInt(-1), big.NewInt(0), big.NewInt(1), big.NewInt(2), big.NewInt(3), new(big.Int).Sub(hi, big.NewInt(1)), hi}
	} else {
		half := new(big.Int).Lsh(big.NewInt(1), uint(t.bits/2))
		sign := new(big.Int).Lsh(big.NewInt(1), uint(t.bits-1))
		values = []*big.Int{big.NewInt(0), big.NewInt(1), big.NewInt(2), big.NewInt(3), new(big.Int).Sub(half, big.NewInt(1)), half, new(big.Int).Sub(sign, big.NewInt(1)), sign, new(big.Int).Sub(hi, big.NewInt(1)), hi}
	}
	var pairs [][2]*big.Int
	for _, l := range values {
		for _, r := range values {
			pairs = append(pairs, [2]*big.Int{l, r})
		}
	}
	width := new(big.Int).Lsh(big.NewInt(1), uint(t.bits))
	for i := 0; i < 128; i++ {
		var pair [2]*big.Int
		for j := 0; j < 2; j++ {
			h := sha256.Sum256([]byte(fmt.Sprintf("integer-miner-%s-%d-%d", t.name(), i, j)))
			x := new(big.Int).Mod(new(big.Int).SetBytes(h[:]), width)
			if t.signed && x.Bit(t.bits-1) != 0 {
				x.Sub(x, width)
			}
			pair[j] = x
		}
		pairs = append(pairs, pair)
	}
	return pairs
}

// Constants propose domains; only the independent oracle decides acceptance.
// This does not claim recovery of the source type when optimizations erase it.
func inferredTypes(a *SymbolicAnalysis) []integerType {
	types := map[string]integerType{}
	add := func(bits int, signed bool) {
		if bits >= 8 && bits <= 256 && bits%8 == 0 {
			t := integerType{bits, signed}
			types[t.name()] = t
		}
	}
	add(256, false)
	add(256, true)
	var visit func(*Expr)
	visit = func(e *Expr) {
		if e == nil {
			return
		}
		if n, ok := constValue(e); ok {
			add(lowMaskBits(n), false)
			p := new(big.Int).Add(n, big.NewInt(1))
			if p.Sign() > 0 && new(big.Int).And(p, new(big.Int).Sub(p, big.NewInt(1))).Sign() == 0 {
				add(n.BitLen()+1, true)
			}
			x := signedWord(n)
			if x.Sign() < 0 {
				x.Neg(x)
				if new(big.Int).And(x, new(big.Int).Sub(x, big.NewInt(1))).Sign() == 0 {
					add(x.BitLen(), true)
				}
			}
		}
		if e.Op == "SIGNEXTEND" && len(e.Args) == 2 {
			if n, ok := constValue(e.Args[0]); ok && n.IsUint64() && n.Uint64() < 32 {
				add(int(n.Uint64()+1)*8, true)
			}
		}
		for _, arg := range e.Args {
			visit(arg)
		}
	}
	for _, p := range a.Paths {
		visit(p.Result)
		for _, c := range p.Conditions {
			visit(c.Expression)
		}
	}
	var out []integerType
	for _, t := range types {
		out = append(out, t)
	}
	sort.Slice(out, func(i, j int) bool { return out[i].name() < out[j].name() })
	return out
}

func validateIntegerDomain(a *SymbolicAnalysis) error {
	var accepted []string
	for _, t := range inferredTypes(a) {
		valid := true
		for _, pair := range typedValidationPairs(t) {
			if checkTypedInputs(a, a.Operation, t, pair[0], pair[1]) != nil {
				valid = false
				break
			}
		}
		if valid {
			accepted = append(accepted, t.name())
		}
	}
	if len(accepted) != 1 {
		return fmt.Errorf("arithmetic domain has %d validated interpretations: %v", len(accepted), accepted)
	}
	a.IntegerType = accepted[0]
	return nil
}
