package vm

import (
	"math/big"
	"math/rand"
	"testing"

	"github.com/holiman/uint256"
)

func TestArithmeticMaskedProductIdentity(t *testing.T) {
	rng := rand.New(rand.NewSource(20261019013))
	cases := 0
	check := func(a, b, mask uint256.Int) {
		t.Helper()
		var product, masked, quotient uint256.Int
		_, overflow := product.MulOverflow(&a, &b)
		masked.And(&product, &mask)
		quotient.Div(&masked, &a)
		got := !overflow && product.Eq(&masked) && (!a.IsZero() || b.IsZero())
		full := new(big.Int).Mul(a.ToBig(), b.ToBig())
		want := full.BitLen() <= 256 && new(big.Int).And(full, mask.ToBig()).Cmp(full) == 0 && (!a.IsZero() || b.IsZero())
		if got != want || got != quotient.Eq(&b) {
			t.Fatalf("a=%s b=%s mask=%s got=%t want=%t quotient=%s", a.Hex(), b.Hex(), mask.Hex(), got, want, quotient.Hex())
		}
		cases++
	}
	for width := uint(1); width <= 256; width++ {
		var mask uint256.Int
		mask.SetAllOne().Rsh(&mask, 256-width)
		for _, a := range []uint256.Int{{}, {1}, {2}, mask, *new(uint256.Int).SetAllOne()} {
			for _, b := range []uint256.Int{{}, {1}, {2}, mask, *new(uint256.Int).SetAllOne()} {
				check(a, b, mask)
			}
		}
		for i := 0; i < 100; i++ {
			a := uint256.Int{rng.Uint64(), rng.Uint64(), rng.Uint64(), rng.Uint64()}
			b := uint256.Int{rng.Uint64(), rng.Uint64(), rng.Uint64(), rng.Uint64()}
			check(a, b, mask)
			a.Rsh(&a, uint(rng.Intn(256)))
			b.Rsh(&b, uint(rng.Intn(256)))
			check(a, b, mask)
			check(a, b, uint256.Int{rng.Uint64(), rng.Uint64(), rng.Uint64(), rng.Uint64()})
		}
	}
	t.Logf("%d independent masked-product and DIV comparisons", cases)
}

func TestArithmeticMaskedProductRecognition(t *testing.T) {
	for order := 0; order < 8; order++ {
		p := &arithmeticPlan{nodes: []arithmeticNode{{}, {}, {op: MUL, x: 0, y: 1}, {op: PUSH32, constant: uint256.Int{255}}, {op: AND, x: 2, y: 3}, {op: DIV, x: 4, y: 0}, {op: EQ, x: 5, y: 1}}}
		for bit, id := range []int{2, 4, 6} {
			if order&(1<<bit) != 0 {
				p.nodes[id].x, p.nodes[id].y = p.nodes[id].y, p.nodes[id].x
			}
		}
		if product, masked, divisor, other, ok := matchMaskedMulCancellation(p, 6); !ok || product != 2 || masked != 4 || divisor != 0 || other != 1 {
			t.Fatal("commuted exact identity not recognized", order)
		}
		p.nodes[4].op = ADD
		if _, _, _, _, ok := matchMaskedMulCancellation(p, 6); ok {
			t.Fatal("arbitrary intervening operation accepted")
		}
		p.nodes[4].op, p.nodes[5].op = AND, SDIV
		if _, _, _, _, ok := matchMaskedMulCancellation(p, 6); ok {
			t.Fatal("signed division accepted by unsigned identity")
		}
	}
}
