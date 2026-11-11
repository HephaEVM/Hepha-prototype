package vm

import (
	"fmt"
	"math/big"
	"math/rand"
	"testing"

	"github.com/holiman/uint256"
)

func TestArithmeticProductWidthSpecialization(t *testing.T) {
	rng := rand.New(rand.NewSource(20261019014))
	for i := 0; i < 20000; i++ {
		a := uint256.Int{rng.Uint64(), rng.Uint64(), rng.Uint64(), rng.Uint64()}
		b := uint256.Int{rng.Uint64(), rng.Uint64(), rng.Uint64(), rng.Uint64()}
		// Cross every limb boundary, including full-width fallbacks and zeros.
		shift := uint(i % 5 * 64)
		a.Rsh(&a, shift)
		b.Rsh(&b, uint((i/5)%5*64))
		full := new(big.Int).Mul(a.ToBig(), b.ToBig())
		want, overflow := uint256.FromBig(full)
		for alias := 0; alias < 3; alias++ {
			x, y := a, b
			var z uint256.Int
			dst := &z
			if alias == 1 {
				dst = &x
			} else if alias == 2 {
				dst = &y
			}
			got := arithmeticMulOverflow(dst, &x, &y)
			if got != overflow || *dst != *want {
				t.Fatalf("a=%s b=%s alias=%d got=%s overflow=%t want=%s overflow=%t", a.Hex(), b.Hex(), alias, dst.Hex(), got, want.Hex(), overflow)
			}
		}
	}
}

// Primitive cost only: both variants compute the identical low product and
// overflow flag. The cases include distributions that take the full fallback.
func BenchmarkArithmeticProductWidthSpecialization(b *testing.B) {
	for _, widths := range [][2]uint{{64, 64}, {128, 128}, {64, 192}, {128, 192}, {192, 192}, {256, 256}} {
		rng := rand.New(rand.NewSource(20261019014))
		var inputs [256][2]uint256.Int
		for i := range inputs {
			for j, width := range widths {
				v := uint256.Int{rng.Uint64(), rng.Uint64(), rng.Uint64(), rng.Uint64()}
				v.Rsh(&v, 256-width)
				inputs[i][j] = v
			}
		}
		for _, special := range []bool{false, true} {
			b.Run(fmt.Sprintf("%dx%d/special_%t", widths[0], widths[1], special), func(b *testing.B) {
				b.ReportAllocs()
				for i := 0; i < b.N; i++ {
					v := &inputs[i&255]
					if special {
						arithmeticDiagnosticFlag = arithmeticMulOverflow(&arithmeticDiagnosticWord, &v[0], &v[1])
					} else {
						_, arithmeticDiagnosticFlag = arithmeticDiagnosticWord.MulOverflow(&v[0], &v[1])
					}
				}
			})
		}
	}
}
