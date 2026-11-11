package vm

import "github.com/holiman/uint256"

// tryMulDivFloor computes floor(a*b/d) using the full 512-bit product.
// computeHelperRegion calls it only after a complete success-expression proof
// and stack/gas admission. This kernel alone is not authority to replace code.
//
// On division by zero or quotient overflow it leaves out unchanged. The caller
// falls back to the original bytecode, which determines the exact
// error data and gas; this helper does not invent a Solidity panic code.
// out may alias any input.
func tryMulDivFloor(out, a, b, d *uint256.Int) bool {
	if d.IsZero() {
		return false
	}
	var result uint256.Int
	if _, overflow := result.MulDivOverflow(a, b, d); overflow {
		return false
	}
	*out = result
	return true
}
