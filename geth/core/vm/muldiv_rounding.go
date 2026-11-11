package vm

import "github.com/holiman/uint256"

// The floor quotient has already been proved to fit a word. Since
// 0 <= a*b-q*d < d <= 2^256-1, the low-word subtraction is the EXACT remainder.
// This reuses the quotient and low product instead of doing another full-width
// MULMOD division. It is only used for a complete checked-ceil success proof.
func mulDivRoundingResult(a, b, d *uint256.Int) (result uint256.Int, wide, round, ok bool) {
	if d.IsZero() {
		return
	}
	var product uint256.Int
	_, wide = product.MulOverflow(a, b)
	if wide {
		if !tryMulDivFloor(&result, a, b, d) {
			return
		}
	} else {
		result.Div(&product, d)
	}
	var rem uint256.Int
	rem.Mul(&result, d)
	rem.Sub(&product, &rem)
	round = !rem.IsZero()
	if round {
		result.AddUint64(&result, 1)
		if result.IsZero() {
			return result, wide, round, false
		}
	}
	return result, wide, round, true
}

func (r *computeHelperRegion) tryMulDivRounding(pc *uint64, s *ScopeContext, stats *ComputeHelperStats) bool {
	n := len(s.Stack.data)
	if n < 4 {
		return false
	}
	in := s.Stack.data[n-4:]
	m := r.rounding
	if !s.Contract.validJumpdest(&in[0]) {
		return false
	}
	result, wide, round, ok := mulDivRoundingResult(&in[m.A], &in[m.B], &in[m.D])
	if !ok {
		return false
	}
	index := 0
	floor := &r.short
	if wide {
		index, floor = 2, &r.wide
	}
	if round {
		index++
	}
	path := &r.rounded[index]
	if n-4+path.peak > 1024 || s.Contract.Gas < path.gas {
		return false
	}
	dest := in[0].Uint64()
	s.Stack.data[n-4] = result
	s.Stack.data = s.Stack.data[:n-3]
	s.Contract.Gas -= path.gas
	*pc = dest
	if stats != nil {
		stats.MulDiv++
		stats.MulDivInstructions += floor.instructions
		stats.MulDivFusions++
		stats.MulDivFusionInstructions += path.instructions - floor.instructions
		if wide {
			stats.MulDivWide++
		}
		if round {
			stats.MulDivRoundUps++
		}
	}
	return true
}
