package vm

import "github.com/holiman/uint256"

// Match a COMPLETE sign-qualified overflow test, not its division comparison
// alone. The conjunction must contain precisely the two operand sign guards
// and one of these comparisons (operand order may be commuted):
//
//	x>0, y>0, y>DIV(MAX,x)
//	x>0, y<0, y<SDIV(MIN,x)
//	x<0, y<0, y<SDIV(MAX,x)
//
// Each conjunction equals signedOverflow(x*y) AND the same sign guards.
// This is an integer identity: signed division truncates towards zero, hence
// the strict comparison is exactly the first integer beyond the allowed bound.
// Extra predicates outside the matched conjunction remain in the decision tree.
func matchSignedMulQuadrant(p *arithmeticPlan, id uint8) (product, divisor, other uint8, positiveDivisor, positiveOther, ok bool) {
	if id < 2 || p.nodes[id].op != AND {
		return
	}
	var leaves [3]uint8
	n := 0
	var flatten func(uint8) bool
	flatten = func(at uint8) bool {
		if at >= 2 && p.nodes[at].op == AND {
			x := p.nodes[at]
			return flatten(x.x) && flatten(x.y)
		}
		if n == len(leaves) {
			return false
		}
		leaves[n] = at
		n++
		return true
	}
	if !flatten(id) || n != 3 {
		return
	}
	maximum := uint256.Int{^uint64(0), ^uint64(0), ^uint64(0), uint64(1)<<63 - 1}
	minimum := uint256.Int{0, 0, 0, uint64(1) << 63}
	for i, leaf := range leaves {
		c := p.nodes[leaf]
		var quotient uint8
		signed := true
		switch c.op {
		case SLT:
			other, quotient = c.x, c.y
		case SGT:
			other, quotient = c.y, c.x
		case GT:
			other, quotient, signed = c.x, c.y, false
		case LT:
			other, quotient, signed = c.y, c.x, false
		default:
			continue
		}
		d := p.nodes[quotient]
		if p.nodes[d.x].op != PUSH32 || (signed && d.op != SDIV) || (!signed && d.op != DIV) {
			continue
		}
		limit := p.nodes[d.x].constant
		if signed && limit == minimum {
			positiveDivisor, positiveOther = true, false
		} else if signed && limit == maximum {
			positiveDivisor, positiveOther = false, false
		} else if !signed && limit == maximum {
			positiveDivisor, positiveOther = true, true
		} else {
			continue
		}
		divisor = d.y
		left, right := leaves[(i+1)%3], leaves[(i+2)%3]
		if !(p.signGuard(left, divisor, positiveDivisor) && p.signGuard(right, other, positiveOther)) && !(p.signGuard(right, divisor, positiveDivisor) && p.signGuard(left, other, positiveOther)) {
			continue
		}
		for j, m := range p.nodes {
			if m.op == MUL && ((m.x == divisor && m.y == other) || (m.y == divisor && m.x == other)) {
				return uint8(j), divisor, other, positiveDivisor, positiveOther, true
			}
		}
	}
	return
}

func (p *arithmeticPlan) signGuard(id, operand uint8, positive bool) bool {
	if id < 2 {
		return false
	}
	n := p.nodes[id]
	op := SGT
	if !positive {
		op = SLT
	}
	if n.op == op && n.x == operand && p.nodes[n.y].op == PUSH32 && p.nodes[n.y].constant.IsZero() {
		return true
	}
	if positive {
		op = SLT
	} else {
		op = SGT
	}
	return n.op == op && n.y == operand && p.nodes[n.x].op == PUSH32 && p.nodes[n.x].constant.IsZero()
}
