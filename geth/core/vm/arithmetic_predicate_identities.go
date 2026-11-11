package vm

// Structural identity only: EQ(DIV(MUL(x,y),x),y), allowing commuted EQ/MUL.
// No signed division, intervening masks, truncation or type inference is used.
// For nonzero x, writing x*y=q*2^256+r proves r/x=y iff q=0.
func matchMulCancellation(p *arithmeticPlan, id uint8) (product, divisor, other uint8, ok bool) {
	return matchInverseMultiplication(p, id, DIV)
}

func matchSignedMulCancellation(p *arithmeticPlan, id uint8) (product, divisor, other uint8, ok bool) {
	return matchInverseMultiplication(p, id, SDIV)
}

func matchInverseMultiplication(p *arithmeticPlan, id uint8, division OpCode) (product, divisor, other uint8, ok bool) {
	n := p.nodes[id]
	if id < 2 || n.op != EQ {
		return
	}
	for _, operands := range [][2]uint8{{n.x, n.y}, {n.y, n.x}} {
		d := p.nodes[operands[0]]
		if d.op != division {
			continue
		}
		m := p.nodes[d.x]
		if m.op == MUL && ((m.x == d.y && m.y == operands[1]) || (m.y == d.y && m.x == operands[1])) {
			return d.x, d.y, operands[1], true
		}
	}
	return
}

// Exact unsigned identity with literal bit cleanup on the product:
// DIV(MUL(x,y)&mask,x)==y iff the full product fits 256 bits, the mask
// removes no product bits, and (x!=0 || y==0). For x>0 a changed product
// is strictly below mathematical x*y, so its quotient must be below y.
// No assumption about source types or mask contiguity is needed.
func matchMaskedMulCancellation(p *arithmeticPlan, id uint8) (product, masked, divisor, other uint8, ok bool) {
	if id < 2 || p.nodes[id].op != EQ {
		return
	}
	n := p.nodes[id]
	for _, eq := range [][2]uint8{{n.x, n.y}, {n.y, n.x}} {
		d := p.nodes[eq[0]]
		if d.op != DIV || p.nodes[d.x].op != AND {
			continue
		}
		cleanup := p.nodes[d.x]
		for _, and := range [][2]uint8{{cleanup.x, cleanup.y}, {cleanup.y, cleanup.x}} {
			m := p.nodes[and[0]]
			if p.nodes[and[1]].op == PUSH32 && m.op == MUL && ((m.x == d.y && m.y == eq[1]) || (m.y == d.y && m.x == eq[1])) {
				return and[0], d.x, d.y, eq[1], true
			}
		}
	}
	return
}

// y > floor(C/x) iff the full mathematical product x*y > C for nonzero
// unsigned x. For x=0 the EVM quotient is 0, so the predicate is y!=0.
// C must be a literal. Reuse only an existing exact MUL of these same nodes.
func matchMulBound(p *arithmeticPlan, id uint8) (product, divisor, other, limit uint8, ok bool) {
	if id < 2 {
		return
	}
	n := p.nodes[id]
	var quotient uint8
	if n.op == GT {
		other, quotient = n.x, n.y
	} else if n.op == LT {
		other, quotient = n.y, n.x
	} else {
		return
	}
	d := p.nodes[quotient]
	if d.op != DIV || p.nodes[d.x].op != PUSH32 {
		return
	}
	for i, m := range p.nodes {
		if m.op == MUL && ((m.x == d.y && m.y == other) || (m.y == d.y && m.x == other)) {
			return uint8(i), d.y, other, d.x, true
		}
	}
	return
}

// Expensive work must survive strength reduction and belong to a guard, not
// merely the mandatory result computation. This is a structural eligibility
// rule for an experimental policy, not a measured universal speedup threshold.
func (p *arithmeticPlan) hasExpensivePredicate(reduced bool) bool {
	var visit func(uint8) bool
	visit = func(id uint8) bool {
		if id < 2 {
			return false
		}
		if reduced {
			if product, _, _, _, _, ok := matchSignedMulQuadrant(p, id); ok {
				n := p.nodes[product]
				return visit(n.x) || visit(n.y)
			}
			if product, _, _, _, ok := matchMaskedMulCancellation(p, id); ok {
				n := p.nodes[product]
				return visit(n.x) || visit(n.y)
			}
			if product, _, _, ok := matchSignedMulCancellation(p, id); ok {
				n := p.nodes[product]
				return visit(n.x) || visit(n.y)
			}
			if product, _, _, ok := matchMulCancellation(p, id); ok {
				n := p.nodes[product]
				return visit(n.x) || visit(n.y)
			}
			if product, _, _, _, ok := matchMulBound(p, id); ok {
				n := p.nodes[product]
				return visit(n.x) || visit(n.y)
			}
		}
		n := p.nodes[id]
		switch n.op {
		case PUSH32:
			return false
		case DIV, SDIV, MOD, SMOD, EXP:
			return true
		case NOT, ISZERO:
			return visit(n.x)
		default:
			return visit(n.x) || visit(n.y)
		}
	}
	for _, d := range p.decisions {
		if d.path < 0 && visit(d.value) {
			return true
		}
	}
	return false
}
