package vm

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"testing"

	"github.com/holiman/uint256"
	miner "go-eth1.16.5-evm/template_miner"
)

func TestArithmeticCanonicalResults(t *testing.T) {
	rng := rand.New(rand.NewSource(20261019))
	for _, signed := range []bool{false, true} {
		for _, bits := range []int{8, 64, 128, 248, 256} {
			name := "uint"
			if signed {
				name = "int"
			}
			d, err := ParseIntegerDomain(fmt.Sprintf("%s%d", name, bits))
			if err != nil {
				t.Fatal(err)
			}
			var mask, index uint256.Int
			mask.SetAllOne().Rsh(&mask, uint(256-bits))
			index.SetUint64(uint64(bits/8 - 1))
			clean := func(x *miner.Expr) *miner.Expr {
				if signed {
					return &miner.Expr{Op: "SIGNEXTEND", Args: []*miner.Expr{{Op: "const", Value: index.Hex()}, x}}
				}
				return &miner.Expr{Op: "AND", Args: []*miner.Expr{{Op: "const", Value: mask.Hex()}, x}}
			}
			for _, op := range []string{"ADD", "SUB", "MUL", "DIV", "MOD", "SDIV", "SMOD"} {
				e := clean(&miner.Expr{Op: op, Args: []*miner.Expr{clean(&miner.Expr{Op: "input", Input: "a"}), clean(&miner.Expr{Op: "input", Input: "b"})}})
				before, _ := json.Marshal(e)
				original, _, err := compileArithmeticResult(e)
				if err != nil {
					t.Fatal(err)
				}
				fast, _, err := compileArithmeticResultPolicy(canonicalArithmeticInputs(e, d), true)
				if err != nil {
					t.Fatal(err)
				}
				after, _ := json.Marshal(e)
				if string(before) != string(after) {
					t.Fatal("expression mutated")
				}
				for i := 0; i < 1000; i++ {
					a := uint256.Int{rng.Uint64(), rng.Uint64(), rng.Uint64(), rng.Uint64()}
					b := uint256.Int{rng.Uint64(), rng.Uint64(), rng.Uint64(), rng.Uint64()}
					if i < 4 {
						a.SetUint64(uint64(i))
						b.Clear()
					}
					if signed {
						a.ExtendSign(&a, &index)
						b.ExtendSign(&b, &index)
					} else {
						a.And(&a, &mask)
						b.And(&b, &mask)
					}
					if !d.canonical(a) || !d.canonical(b) {
						t.Fatal("invalid witness")
					}
					if got, want := fast(a, b), original(a, b); got != want {
						t.Fatalf("%s%d %s a=%s b=%s got=%s want=%s", name, bits, op, a.Hex(), b.Hex(), got.Hex(), want.Hex())
					}
				}
			}
		}
	}
}

func TestArithmeticCanonicalRetainsUnprovedCleanup(t *testing.T) {
	for _, c := range []struct{ domain, op, value string }{{"int8", "AND", "0xff"}, {"uint16", "AND", "0xff"}, {"uint8", "SIGNEXTEND", "0x0"}} {
		d, _ := ParseIntegerDomain(c.domain)
		e := &miner.Expr{Op: c.op, Args: []*miner.Expr{{Op: "const", Value: c.value}, {Op: "input", Input: "a"}}}
		if got := canonicalArithmeticInputs(e, d); got.Op != c.op {
			t.Fatal("unproved cleanup eliminated", c)
		}
	}
}
