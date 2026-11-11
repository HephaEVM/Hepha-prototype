package vm

import (
	"math/big"
	"math/rand"
	"testing"

	"github.com/holiman/uint256"
)

func TestMulDivFloorPrototype(t *testing.T) {
	var maximum, half, limb, almostMaximum uint256.Int
	maximum.SetAllOne()
	half.Lsh(uint256.NewInt(1), 255)
	limb.Lsh(uint256.NewInt(1), 128)
	almostMaximum.Sub(&maximum, uint256.NewInt(1))
	values := []uint256.Int{{}, *uint256.NewInt(1), *uint256.NewInt(2), *uint256.NewInt(3), limb, half, almostMaximum, maximum}
	cases, wide, rejected := 0, 0, 0
	check := func(a, b, d uint256.Int) {
		t.Helper()
		cases++
		product := new(big.Int).Mul(a.ToBig(), b.ToBig())
		if product.BitLen() > 256 {
			wide++
		}
		valid := !d.IsZero()
		var expected uint256.Int
		if valid {
			q := new(big.Int).Quo(product, d.ToBig())
			valid = q.BitLen() <= 256
			if valid {
				expected.SetFromBig(q)
			}
		}
		if !valid {
			rejected++
		}
		// Check independent output and aliasing each operand. Rejection must
		// preserve the output and every input, including the aliased input.
		for alias := -1; alias < 3; alias++ {
			inputs := [3]uint256.Int{a, b, d}
			before := inputs
			result := *uint256.NewInt(0xabcdef)
			out := &result
			if alias >= 0 {
				out = &inputs[alias]
			}
			old := *out
			ok := tryMulDivFloor(out, &inputs[0], &inputs[1], &inputs[2])
			if ok != valid || (ok && *out != expected) || (!ok && *out != old) {
				t.Fatalf("a=%x b=%x d=%x alias=%d accepted=%v expected=%v out=%x", a, b, d, alias, ok, valid, *out)
			}
			for i := range inputs {
				if i != alias && inputs[i] != before[i] {
					t.Fatal("input mutated")
				}
			}
		}
	}
	for _, a := range values {
		for _, b := range values {
			for _, d := range values {
				check(a, b, d)
			}
		}
	}
	rng := rand.New(rand.NewSource(20261020))
	for i := 0; i < 10000; i++ {
		var x [3]uint256.Int
		for j := range x {
			x[j] = uint256.Int{rng.Uint64(), rng.Uint64(), rng.Uint64(), rng.Uint64()}
			// Mix ordinary, limb-boundary and full-width operands.
			x[j].Rsh(&x[j], uint(rng.Intn(257)))
		}
		check(x[0], x[1], x[2])
	}
	t.Logf("cases=%d alias_variants=%d wide_products=%d rejected=%d", cases, cases*4, wide, rejected)
}

func TestMulDivFloorPrototypeNoAllocations(t *testing.T) {
	var a, b, d, out uint256.Int
	a.SetAllOne()
	b.SetAllOne()
	d.SetAllOne()
	if n := testing.AllocsPerRun(1000, func() {
		if !tryMulDivFloor(&out, &a, &b, &d) || out != a {
			panic("unexpected result")
		}
	}); n != 0 {
		t.Fatalf("allocations: %v", n)
	}
}
