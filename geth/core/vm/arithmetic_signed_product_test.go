package vm

import (
	"math/big"
	"math/rand"
	"testing"

	"github.com/holiman/uint256"
)

// Independent unbounded-integer oracle, plus a direct comparison with the
// original EVM operations. Include overflowing and non-overflowing products;
// uniformly random full-width factors alone almost always overflow.
func TestArithmeticSignedProductIdentity(t *testing.T) {
	modulus := new(big.Int).Lsh(big.NewInt(1), 256)
	maximum := new(big.Int).Sub(new(big.Int).Rsh(new(big.Int).Set(modulus), 1), big.NewInt(1))
	minimum := new(big.Int).Neg(new(big.Int).Add(new(big.Int).Set(maximum), big.NewInt(1)))
	signed := func(v uint256.Int) *big.Int {
		n := v.ToBig()
		if v[3]>>63 != 0 {
			n.Sub(n, modulus)
		}
		return n
	}
	cases := 0
	check := func(a, b uint256.Int) {
		t.Helper()
		product, overflow := arithmeticSignedProduct(a, b)
		mathematical := new(big.Int).Mul(signed(a), signed(b))
		wantOverflow := mathematical.Cmp(minimum) < 0 || mathematical.Cmp(maximum) > 0
		if arithmeticSignedMulOverflows(a, b) != wantOverflow {
			t.Fatalf("guard-only signed range mismatch: a=%s b=%s want=%t", a.Hex(), b.Hex(), wantOverflow)
		}
		low, _ := uint256.FromBig(new(big.Int).Mod(mathematical, modulus))
		if product != *low || overflow != wantOverflow {
			t.Fatalf("a=%s b=%s product=%s/%s overflow=%t/%t", a.Hex(), b.Hex(), product.Hex(), low.Hex(), overflow, wantOverflow)
		}
		var original, quotient uint256.Int
		original.Mul(&a, &b)
		quotient.SDiv(&original, &a)
		predicate := (!overflow && (!a.IsZero() || b.IsZero())) || arithmeticSignedDivisionWrap(a, b)
		if product != original || predicate != quotient.Eq(&b) {
			t.Fatalf("inverse predicate mismatch: a=%s b=%s product=%s quotient=%s predicate=%t", a.Hex(), b.Hex(), product.Hex(), quotient.Hex(), predicate)
		}
		cases++
	}
	edges := []uint256.Int{{}, {1}, {2}, {3}}
	for _, bit := range []uint{7, 8, 31, 32, 63, 64, 65, 95, 96, 127, 128, 129, 191, 192, 253, 254, 255} {
		var v, one uint256.Int
		one.SetOne()
		v.Lsh(&one, bit)
		edges = append(edges, v)
		v.Sub(&v, &one)
		edges = append(edges, v)
		v.AddUint64(&v, 2)
		edges = append(edges, v)
	}
	for _, v := range append([]uint256.Int(nil), edges...) {
		v.Neg(&v)
		edges = append(edges, v)
	}
	for _, a := range edges {
		for _, b := range edges {
			check(a, b)
		}
	}
	rng := rand.New(rand.NewSource(20261019012))
	for i := 0; i < 20000; i++ {
		a := uint256.Int{rng.Uint64(), rng.Uint64(), rng.Uint64(), rng.Uint64()}
		b := uint256.Int{rng.Uint64(), rng.Uint64(), rng.Uint64(), rng.Uint64()}
		check(a, b)
		a.Rsh(&a, uint(rng.Intn(256)))
		b.Rsh(&b, uint(rng.Intn(256)))
		for sign := 0; sign < 4; sign++ {
			x, y := a, b
			if sign&1 != 0 {
				x.Neg(&x)
			}
			if sign&2 != 0 {
				y.Neg(&y)
			}
			check(x, y)
		}
	}
	// MIN/-1 is the one EVM SDIV wrapping exception. Cancellation is
	// asymmetric here, even though the product itself is commutative.
	minusOne := uint256.Int{^uint64(0), ^uint64(0), ^uint64(0), ^uint64(0)}
	minWord := uint256.Int{0, 0, 0, uint64(1) << 63}
	if !arithmeticSignedDivisionWrap(minusOne, minWord) || arithmeticSignedDivisionWrap(minWord, minusOne) {
		t.Fatal("signed division wrap exception is not directional")
	}
	t.Logf("%d independent signed-product and exact cancellation comparisons", cases)
}
