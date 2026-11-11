package vm

import "github.com/holiman/uint256"

// Compute the same low 256 product bits as MUL and whether the mathematical
// signed product lies outside [-2^255, 2^255-1]. Magnitudes require at most
// 256 bits, including abs(MIN); MulOverflow also detects discarded high bits.
func arithmeticSignedProduct(a, b uint256.Int) (product uint256.Int, overflow bool) {
	negativeA, negativeB := a[3]>>63 != 0, b[3]>>63 != 0
	if negativeA {
		a.Neg(&a)
	}
	if negativeB {
		b.Neg(&b)
	}
	wide := arithmeticMulOverflow(&product, &a, &b)
	if negativeA != negativeB {
		// Magnitude 2^255 is representable only for a negative product.
		overflow = wide || product[3] > uint64(1)<<63 || (product[3] == uint64(1)<<63 && product[0]|product[1]|product[2] != 0)
		product.Neg(&product)
	} else {
		overflow = wide || product[3]>>63 != 0
	}
	return
}

// Same signed range test without converting the magnitude back to a wrapped
// EVM result. The producer already computed that result; no predicate is lost.
func arithmeticSignedMulOverflows(a, b uint256.Int) bool {
	negativeA, negativeB := a[3]>>63 != 0, b[3]>>63 != 0
	if negativeA {
		a.Neg(&a)
	}
	if negativeB {
		b.Neg(&b)
	}
	var magnitude uint256.Int
	wide := arithmeticMulOverflow(&magnitude, &a, &b)
	if negativeA != negativeB {
		return wide || magnitude[3] > uint64(1)<<63 || (magnitude[3] == uint64(1)<<63 && magnitude[0]|magnitude[1]|magnitude[2] != 0)
	}
	return wide || magnitude[3]>>63 != 0
}

// EVM SDIV has a single wrapping quotient, MIN / -1 == MIN. It makes
// SDIV(MUL(x,y),x)==y true for x=-1,y=MIN even though the product overflows.
// Keep this exception in addition to the division-by-zero case. Otherwise,
// wrapping the product changes it by >=2^256, greater than any signed divisor,
// so truncating division cannot recover y after signed multiplication overflow.
func arithmeticSignedDivisionWrap(divisor, other uint256.Int) bool {
	return divisor[0]&divisor[1]&divisor[2]&divisor[3] == ^uint64(0) && other[3] == uint64(1)<<63 && other[0]|other[1]|other[2] == 0
}
