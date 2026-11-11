package vm

import (
	"math/bits"

	"github.com/holiman/uint256"
)

// Runtime operand widths can be smaller than their Solidity type. These two
// exact cases avoid computing high product limbs known to be zero; all other
// inputs retain the library's complete 512-bit multiplication. No overflow
// predicate or type admission is weakened, and z may alias either input.
func arithmeticMulOverflow(z, a, b *uint256.Int) bool {
	if a[2]|a[3]|b[2]|b[3] == 0 {
		if a[1]|b[1] == 0 {
			hi, lo := bits.Mul64(a[0], b[0])
			*z = uint256.Int{lo, hi}
			return false
		}
		z.Mul(a, b) // 128 x 128 fits 256 bits exactly.
		return false
	}
	_, overflow := z.MulOverflow(a, b)
	return overflow
}
