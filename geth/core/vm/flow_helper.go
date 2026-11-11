package vm

import "math/big"

func flowProveHelper(h *arithmeticRegion, inputs []hotInterval, facts *flowFacts, external []int) (uint8, bool) {
	if h.plan == nil || (h.domain.Signed && (!paperCEExtendedProof || h.domain.Bits != 256)) || len(inputs) != 2 {
		return 0, false
	}
	max := new(big.Int).Sub(new(big.Int).Lsh(big.NewInt(1), uint(h.domain.Bits)), big.NewInt(1))
	for _, x := range inputs {
		if x.hi.Cmp(max) > 0 {
			return 0, false
		}
	}
	parts := [][]hotInterval{inputs}
	for i := 0; i < 2; i++ {
		var next [][]hotInterval
		for _, p := range parts {
			if p[i].lo.Sign() == 0 && p[i].hi.Sign() > 0 {
				zero := append([]hotInterval(nil), p...)
				zero[i] = hotBool(false)
				positive := append([]hotInterval(nil), p...)
				positive[i] = hotInterval{big.NewInt(1), p[i].hi}
				next = append(next, zero, positive)
			} else {
				next = append(next, p)
			}
		}
		parts = next
	}
	result := -1
	for _, part := range parts {
		p := h.plan
		bounds := append([]hotInterval(nil), part...)
		alias := []int{0, 1}
		for i := 2; i < len(p.nodes); i++ {
			n := p.nodes[i]
			x, y := alias[n.x], alias[n.y]
			v := hotFull()
			same := i
			if n.op == PUSH32 {
				v = hotExact(n.constant)
			} else {
				a := []hotInterval{bounds[x], bounds[y]}
				if n.op == ISZERO || n.op == NOT {
					a = a[:1]
				}
				v = hotTransfer(n.op, a)
				if x < 2 && y < 2 {
					if q, ok := facts.compare(n.op, external[x], external[y]); ok {
						v = hotBool(q)
					}
				}
				if n.op == EQ && x == y {
					v = hotBool(true)
				}
				// Correlation missed by intervals: a non-wrapping sum is >=
				// either summand; exact quotient of a non-wrapping product.
				if n.op == DIV && x >= 2 && bounds[y].lo.Sign() > 0 {
					m := p.nodes[x]
					if m.op == MUL && new(big.Int).Mul(bounds[alias[m.x]].hi, bounds[alias[m.y]].hi).BitLen() <= 256 {
						if alias[m.x] == y {
							same = alias[m.y]
						} else if alias[m.y] == y {
							same = alias[m.x]
						}
						if same != i {
							v = bounds[same]
						}
					}
				}
				if n.op == LT || n.op == GT {
					sum, operand := x, y
					if n.op == GT {
						sum, operand = y, x
					}
					if sum >= 2 {
						m := p.nodes[sum]
						if m.op == ADD && (alias[m.x] == operand || alias[m.y] == operand) && new(big.Int).Add(bounds[alias[m.x]].hi, bounds[alias[m.y]].hi).BitLen() <= 256 {
							v = hotBool(false)
						}
					}
				}
			}
			bounds = append(bounds, v)
			alias = append(alias, same)
		}
		var visit func(int) bool
		visit = func(at int) bool {
			if at < 0 || at >= len(p.decisions) {
				return false
			}
			d := p.decisions[at]
			if d.path >= 0 {
				j := d.path
				if j >= len(h.syncHot.costs) || h.syncHot.costs[j] == syncHotFallback || h.syncHot.costs[j] != h.syncHot.admissionGas {
					return false
				}
				if result < 0 {
					result = int(p.results[j])
				}
				return result == int(p.results[j])
			}
			v := bounds[d.value]
			if v.hi.Sign() == 0 {
				return visit(d.zero)
			}
			if v.lo.Sign() > 0 {
				return visit(d.nonzero)
			}
			return visit(d.zero) && visit(d.nonzero)
		}
		if !visit(p.root) {
			return 0, false
		}
	}
	return uint8(result), result >= 0
}
