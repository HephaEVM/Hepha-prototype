package vm

import (
	"github.com/holiman/uint256"
	miner "go-eth1.16.5-evm/template_miner"
)

// Selected at binding time, not by a runtime opcode/type switch. The general
// result compiler remains the fallback for more complicated expressions.
var arithmeticDirectResults = map[string]arithmeticExpression{
	"ADD(a,b)":  func(a, b uint256.Int) (z uint256.Int) { z.Add(&a, &b); return },
	"ADD(b,a)":  func(a, b uint256.Int) (z uint256.Int) { z.Add(&b, &a); return },
	"SUB(a,b)":  func(a, b uint256.Int) (z uint256.Int) { z.Sub(&a, &b); return },
	"SUB(b,a)":  func(a, b uint256.Int) (z uint256.Int) { z.Sub(&b, &a); return },
	"MUL(a,b)":  func(a, b uint256.Int) (z uint256.Int) { z.Mul(&a, &b); return },
	"MUL(b,a)":  func(a, b uint256.Int) (z uint256.Int) { z.Mul(&b, &a); return },
	"DIV(a,b)":  func(a, b uint256.Int) (z uint256.Int) { z.Div(&a, &b); return },
	"DIV(b,a)":  func(a, b uint256.Int) (z uint256.Int) { z.Div(&b, &a); return },
	"MOD(a,b)":  func(a, b uint256.Int) (z uint256.Int) { z.Mod(&a, &b); return },
	"MOD(b,a)":  func(a, b uint256.Int) (z uint256.Int) { z.Mod(&b, &a); return },
	"SDIV(a,b)": func(a, b uint256.Int) (z uint256.Int) { z.SDiv(&a, &b); return },
	"SDIV(b,a)": func(a, b uint256.Int) (z uint256.Int) { z.SDiv(&b, &a); return },
	"SMOD(a,b)": func(a, b uint256.Int) (z uint256.Int) { z.SMod(&a, &b); return },
	"SMOD(b,a)": func(a, b uint256.Int) (z uint256.Int) { z.SMod(&b, &a); return },
}

// The async producer has already checked that BOTH input words are canonical.
// Eliminate type-specific cleanup only on these inputs. On computed results,
// remove only universal word identities (AND all-ones, SIGNEXTEND >= 31).
// Narrow result masks/sign extensions remain. The guard/gas proof is unchanged.
func canonicalArithmeticInputs(e *miner.Expr, d IntegerDomain) *miner.Expr {
	if e == nil {
		return nil
	}
	out := *e
	out.Args = make([]*miner.Expr, len(e.Args))
	for i, a := range e.Args {
		out.Args[i] = canonicalArithmeticInputs(a, d)
	}
	isInput := func(x *miner.Expr) bool { return x != nil && x.Op == "input" && (x.Input == "a" || x.Input == "b") }
	if len(out.Args) != 2 {
		return &out
	}
	if out.Op == "AND" {
		for _, pair := range [][2]*miner.Expr{{out.Args[0], out.Args[1]}, {out.Args[1], out.Args[0]}} {
			if pair[1].Op != "const" {
				continue
			}
			v, err := uint256.FromHex(pair[1].Value)
			if err != nil {
				continue
			}
			var mask uint256.Int
			mask.SetAllOne()
			if *v == mask {
				return pair[0]
			}
			if isInput(pair[0]) && !d.Signed && d.Bits > 0 && d.Bits < 256 {
				mask.Rsh(&mask, uint(256-d.Bits))
				if *v == mask {
					return pair[0]
				}
			}
		}
	}
	if out.Op == "SIGNEXTEND" && out.Args[0].Op == "const" {
		v, err := uint256.FromHex(out.Args[0].Value)
		if err == nil && v.IsUint64() && (v.Uint64() >= 31 || (isInput(out.Args[1]) && d.Signed && d.Bits > 0 && v.Uint64() == uint64(d.Bits/8-1))) {
			return out.Args[1]
		}
	}
	return &out
}
