package vm

import "slices"

// A sufficient (not necessary) check for the existing shared-source emitter.
// Equal node definitions may be factored; the only admitted divergence is a
// complementary guard on the same value. Unknown shapes keep eager validation.
// This changes preparation timing, never the source or semantic cache key.
func lazyFlowShareable(ps []*preparedFlow, proof bool) bool {
	type path struct {
		nodes  []hotRegionNode
		bounds []hotInterval
		at     int
	}
	paths := make([]path, len(ps))
	for i, p := range ps {
		r, b := p.optimized(proof)
		paths[i] = path{nodes: r.nodes, bounds: b}
	}
	equal := func(a, b path) bool {
		if a.at >= len(a.nodes) || b.at >= len(b.nodes) || a.at != b.at {
			return false
		}
		x, y := a.nodes[a.at], b.nodes[b.at]
		if x.op != y.op || x.nonzero != y.nonzero || x.helper != y.helper || x.value != y.value || !slices.Equal(x.args, y.args) {
			return false
		}
		id := 32 + a.at
		if id >= len(a.bounds) || id >= len(b.bounds) || (a.bounds[id].lo.Cmp(b.bounds[id].lo) != 0 || a.bounds[id].hi.Cmp(b.bounds[id].hi) != 0) {
			return false
		}
		for _, id := range x.args {
			if id >= len(a.bounds) || id >= len(b.bounds) || (a.bounds[id].lo.Cmp(b.bounds[id].lo) != 0 || a.bounds[id].hi.Cmp(b.bounds[id].hi) != 0) {
				return false
			}
		}
		return true
	}
	var visit func([]path) bool
	visit = func(p []path) bool {
		if len(p) <= 1 {
			return true
		}
		for {
			same := true
			for _, q := range p[1:] {
				if !equal(p[0], q) {
					same = false
					break
				}
			}
			if !same {
				break
			}
			for i := range p {
				p[i].at++
			}
		}
		if p[0].at >= len(p[0].nodes) {
			return false
		}
		guard := p[0].nodes[p[0].at]
		if guard.op != JUMPI {
			return false
		}
		var zero, nonzero []path
		for _, q := range p {
			if q.at >= len(q.nodes) {
				return false
			}
			n := q.nodes[q.at]
			if n.op != JUMPI || !slices.Equal(n.args, guard.args) {
				return false
			}
			q.at++
			if n.nonzero {
				nonzero = append(nonzero, q)
			} else {
				zero = append(zero, q)
			}
		}
		if len(zero) == 0 || len(nonzero) == 0 {
			return false
		}
		return visit(zero) && visit(nonzero)
	}
	return visit(paths)
}
