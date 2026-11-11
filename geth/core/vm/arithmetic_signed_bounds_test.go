package vm

import (
	"math/big"
	"math/rand"
	"testing"

	"github.com/holiman/uint256"
)

func signedQuadrantPlan(kind int) *arithmeticPlan {
	limit := uint256.Int{^uint64(0), ^uint64(0), ^uint64(0), uint64(1)<<63 - 1}
	division, comparison, xSign, ySign := DIV, GT, SGT, SGT
	if kind == 1 {
		limit = uint256.Int{0, 0, 0, uint64(1) << 63}
		division, comparison, ySign = SDIV, SLT, SLT
	} else if kind == 2 {
		division, comparison, xSign, ySign = SDIV, SLT, SLT, SLT
	}
	return &arithmeticPlan{nodes: []arithmeticNode{
		{}, {}, {op: PUSH32}, {op: PUSH32, constant: limit},
		{op: MUL, x: 0, y: 1}, {op: division, x: 3, y: 0},
		{op: comparison, x: 1, y: 5}, {op: xSign, x: 0, y: 2}, {op: ySign, x: 1, y: 2},
		{op: AND, x: 7, y: 8}, {op: AND, x: 9, y: 6},
	}}
}

func TestArithmeticSignedQuadrantIdentity(t *testing.T) {
	modulus := new(big.Int).Lsh(big.NewInt(1), 256)
	min := new(big.Int).Neg(new(big.Int).Rsh(new(big.Int).Set(modulus), 1))
	max := new(big.Int).Sub(new(big.Int).Neg(new(big.Int).Set(min)), big.NewInt(1))
	signed := func(v uint256.Int) *big.Int {
		n := v.ToBig()
		if v[3]>>63 != 0 {
			n.Sub(n, modulus)
		}
		return n
	}
	edges := []uint256.Int{{}, {1}, {2}, {3}, {7}, {127}}
	for _, bit := range []uint{63, 64, 127, 128, 191, 192, 254, 255} {
		var v uint256.Int
		v.SetOne().Lsh(&v, bit)
		edges = append(edges, v, *new(uint256.Int).SubUint64(&v, 1), *new(uint256.Int).AddUint64(&v, 1))
	}
	for _, v := range append([]uint256.Int(nil), edges...) {
		edges = append(edges, *new(uint256.Int).Neg(&v))
	}
	cases := 0
	for kind := 0; kind < 3; kind++ {
		p := signedQuadrantPlan(kind)
		_, _, _, px, py, ok := matchSignedMulQuadrant(p, 10)
		if !ok {
			t.Fatal("missing quadrant", kind)
		}
		check := func(a, b uint256.Int) {
			t.Helper()
			var scratch arithmeticScratch
			scratch.words[0], scratch.words[1], scratch.valid = a, b, 3
			original := !p.value(10, &scratch).IsZero()
			_, overflow := arithmeticSignedProduct(a, b)
			x, y := signed(a), signed(b)
			signs := (px && x.Sign() > 0 || !px && x.Sign() < 0) && (py && y.Sign() > 0 || !py && y.Sign() < 0)
			product := new(big.Int).Mul(x, y)
			oracle := signs && (product.Cmp(min) < 0 || product.Cmp(max) > 0)
			if original != oracle || (overflow && signs) != oracle {
				t.Fatalf("kind=%d a=%s b=%s original=%t oracle=%t", kind, a.Hex(), b.Hex(), original, oracle)
			}
			cases++
		}
		for _, a := range edges {
			for _, b := range edges {
				check(a, b)
			}
		}
		rng := rand.New(rand.NewSource(20261019016))
		for i := 0; i < 20000; i++ {
			a := uint256.Int{rng.Uint64(), rng.Uint64(), rng.Uint64(), rng.Uint64()}
			b := uint256.Int{rng.Uint64(), rng.Uint64(), rng.Uint64(), rng.Uint64()}
			check(a, b)
			a.SRsh(&a, uint(rng.Intn(256)))
			b.SRsh(&b, uint(rng.Intn(256)))
			check(a, b)
		}
	}
	t.Logf("%d independent sign-qualified bounds comparisons", cases)
}

func TestArithmeticSignedQuadrantRecognition(t *testing.T) {
	for kind := 0; kind < 3; kind++ {
		for _, order := range [][3]uint8{{6, 7, 8}, {6, 8, 7}, {7, 6, 8}, {7, 8, 6}, {8, 6, 7}, {8, 7, 6}} {
			p := signedQuadrantPlan(kind)
			p.nodes[9].x, p.nodes[9].y, p.nodes[10].x, p.nodes[10].y = order[0], order[1], 9, order[2]
			if product, _, _, _, _, ok := matchSignedMulQuadrant(p, 10); !ok || product != 4 {
				t.Fatal("commuted conjunction not matched")
			}
			p.nodes[7].op = EQ
			if _, _, _, _, _, ok := matchSignedMulQuadrant(p, 10); ok {
				t.Fatal("missing strict sign guard accepted")
			}
		}
		p := signedQuadrantPlan(kind)
		p.nodes[3].constant.SubUint64(&p.nodes[3].constant, 1)
		if _, _, _, _, _, ok := matchSignedMulQuadrant(p, 10); ok {
			t.Fatal("different numerical bound accepted")
		}
		p = signedQuadrantPlan(kind)
		p.nodes[4].op = ADD
		if _, _, _, _, _, ok := matchSignedMulQuadrant(p, 10); ok {
			t.Fatal("missing matching multiplication accepted")
		}
	}
}
