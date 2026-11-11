package vm

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"

	"github.com/holiman/uint256"
)

// IntegerDomain describes mathematical operands, not already wrapped EVM results.
type IntegerDomain struct {
	Bits   uint
	Signed bool
}

func ParseIntegerDomain(name string) (IntegerDomain, error) {
	if name == "" {
		name = "uint256"
	}
	signed := strings.HasPrefix(name, "int")
	prefix := "uint"
	if signed {
		prefix = "int"
	}
	n, err := strconv.Atoi(strings.TrimPrefix(name, prefix))
	if err != nil || !strings.HasPrefix(name, prefix) || n < 8 || n > 256 || n%8 != 0 {
		return IntegerDomain{}, fmt.Errorf("invalid integer domain %q", name)
	}
	return IntegerDomain{uint(n), signed}, nil
}

func (d IntegerDomain) decode(w uint256.Int) *big.Int {
	x := w.ToBig()
	if d.Signed && w.Sign() < 0 {
		x.Sub(x, new(big.Int).Lsh(big.NewInt(1), 256))
	}
	return x
}

func (d IntegerDomain) bounds() (*big.Int, *big.Int) {
	n := d.Bits
	lo := new(big.Int)
	if d.Signed {
		n--
		lo.Neg(new(big.Int).Lsh(big.NewInt(1), n))
	}
	hi := new(big.Int).Sub(new(big.Int).Lsh(big.NewInt(1), n), big.NewInt(1))
	return lo, hi
}

func (d IntegerDomain) canonical(w uint256.Int) bool {
	if d.Bits == 256 {
		return true
	}
	if !d.Signed {
		return w.BitLen() <= int(d.Bits)
	}
	var normalized uint256.Int
	byteIndex := uint256.NewInt(uint64(d.Bits/8 - 1))
	normalized.ExtendSign(&w, byteIndex)
	return normalized == w
}

// CheckArithmeticRange is an optional diagnostic oracle shared by both modes.
// The default production checker evaluates exact bound bytecode guards instead.
// The exact result is calculated BEFORE reduction modulo 2^256. Division by zero
// is a separate exception; signed division truncates toward zero.
func CheckArithmeticRange(op string, d IntegerDomain, lhs, rhs uint256.Int) (result uint256.Int, panicCode byte, err error) {
	if d.Bits < 8 || d.Bits > 256 || d.Bits%8 != 0 || !d.canonical(lhs) || !d.canonical(rhs) {
		return result, 0, fmt.Errorf("noncanonical integer operands")
	}
	a, b := d.decode(lhs), d.decode(rhs)
	z := new(big.Int)
	switch op {
	case "ADD":
		z.Add(a, b)
	case "SUB":
		z.Sub(a, b)
	case "MUL":
		z.Mul(a, b)
	case "DIV", "SDIV", "MOD", "SMOD":
		if (op == "SDIV" || op == "SMOD") != d.Signed {
			return result, 0, fmt.Errorf("operation/domain mismatch")
		}
		if b.Sign() == 0 {
			return result, 0x12, nil
		}
		if op == "DIV" || op == "SDIV" {
			z.Quo(a, b)
		} else {
			z.Rem(a, b)
		}
	default:
		return result, 0, fmt.Errorf("unsupported arithmetic %s", op)
	}
	lo, hi := d.bounds()
	if z.Cmp(lo) < 0 || z.Cmp(hi) > 0 {
		panicCode = 0x11
	}
	z.Mod(z, new(big.Int).Lsh(big.NewInt(1), 256))
	result.SetFromBig(z)
	return result, panicCode, nil
}
