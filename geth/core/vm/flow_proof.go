package vm

import "math/big"

// Facts are scoped to one bounded path and refer to immutable SSA values.
// Definition-time intervals are kept separate from later guard refinements:
// a later check must never justify truncating a computation before that check.
type flowFacts struct {
	less, le, equal map[[2]int]bool
	env             []hotInterval
	nodes           []hotRegionNode
}
type flowReport struct{ GuardsRemoved, HelpersRemoved int }

func newFlowFacts() *flowFacts {
	return &flowFacts{less: map[[2]int]bool{}, le: map[[2]int]bool{}, equal: map[[2]int]bool{}}
}
func (f *flowFacts) compare(op OpCode, a, b int) (bool, bool) {
	if op == GT {
		return f.compare(LT, b, a)
	}
	if op == LT {
		if a == b || f.le[[2]int{b, a}] {
			return false, true
		}
		if f.less[[2]int{a, b}] {
			return true, true
		}
	}
	if op == EQ {
		if a == b || f.equal[[2]int{a, b}] {
			return true, true
		}
		if f.less[[2]int{a, b}] || f.less[[2]int{b, a}] {
			return false, true
		}
	}
	return false, false
}
func (f *flowFacts) bound(id int, lo, hi *big.Int) {
	x := f.env[id]
	l, h := new(big.Int).Set(x.lo), new(big.Int).Set(x.hi)
	if lo != nil && lo.Cmp(l) > 0 {
		l.Set(lo)
	}
	if hi != nil && hi.Cmp(h) < 0 {
		h.Set(hi)
	}
	// An impossible path is retained conservatively; its actual guard rejects.
	if l.Cmp(h) <= 0 {
		f.env[id] = hotInterval{l, h}
	}
}
func (f *flowFacts) assume(id int, nonzero bool) {
	if id < hotRegionInputs {
		return
	}
	n := f.nodes[id-hotRegionInputs]
	if n.helper != nil {
		return
	}
	if n.op == ISZERO {
		f.assume(n.args[0], !nonzero)
		return
	}
	if len(n.args) != 2 {
		return
	}
	a, b := n.args[0], n.args[1]
	if n.op == GT {
		a, b = b, a
	}
	x, y := f.env[a], f.env[b]
	if n.op == LT || n.op == GT {
		if nonzero {
			f.less[[2]int{a, b}] = true
			f.le[[2]int{a, b}] = true
			if y.hi.Sign() > 0 {
				f.bound(a, nil, new(big.Int).Sub(y.hi, big.NewInt(1)))
			}
			if x.lo.Cmp(hotWordMax) < 0 {
				f.bound(b, new(big.Int).Add(x.lo, big.NewInt(1)), nil)
			}
		} else {
			f.le[[2]int{b, a}] = true
			f.bound(a, y.lo, nil)
			f.bound(b, nil, x.hi)
		}
	} else if n.op == EQ && nonzero {
		f.equal[[2]int{a, b}] = true
		f.equal[[2]int{b, a}] = true
		f.le[[2]int{a, b}] = true
		f.le[[2]int{b, a}] = true
		f.bound(a, y.lo, y.hi)
		f.bound(b, x.lo, x.hi)
	}
}
func optimizeFlowRegion(r *hotRegion) (*hotRegion, []hotInterval, flowReport) {
	return optimizeFlowRegionChecks(r, !paperWithoutCheckElimination)
}

// The control retains region membership, inferred facts and execution
// requirements. It disables check removal only, rather than switching off
// range inference (which also changes arithmetic lowering).
func optimizeFlowRegionChecks(r *hotRegion, eliminate bool) (*hotRegion, []hotInterval, flowReport) {
	out := *r
	out.nodes = nil
	out.outputs = nil
	f := newFlowFacts()
	var report flowReport
	refs := make([]int, hotRegionInputs+len(r.nodes))
	bounds := make([]hotInterval, hotRegionInputs)
	for i := 0; i < hotRegionInputs; i++ {
		refs[i] = i
		bounds[i] = hotFull()
	}
	f.env = append([]hotInterval(nil), bounds...)
	guardDeps := hotGuardDependencies(r)
	preserveCondition := false
	add := func(n hotRegionNode) int {
		v := hotFull()
		if n.op == PUSH32 {
			v = hotExact(n.value)
		} else if n.helper == nil && n.op != MSTORE && n.op != JUMPI {
			args := make([]hotInterval, len(n.args))
			for i, a := range n.args {
				args[i] = f.env[a]
			}
			v = hotTransfer(n.op, args)
			if len(n.args) == 2 {
				if q, ok := f.compare(n.op, n.args[0], n.args[1]); ok {
					v = hotBool(q)
				}
			}
			if v.exact() && !preserveCondition {
				n = hotRegionNode{op: PUSH32}
				n.value.SetFromBig(v.lo)
			}
		}
		id := hotRegionInputs + len(out.nodes)
		out.nodes = append(out.nodes, n)
		bounds = append(bounds, v)
		f.env = append(f.env, v)
		f.nodes = out.nodes
		return id
	}
	for i, n := range r.nodes {
		preserveCondition = !eliminate && guardDeps[hotRegionInputs+i]
		n.args = append([]int(nil), n.args...)
		for j, a := range n.args {
			n.args[j] = refs[a]
		}
		if n.op == JUMPI && n.helper == nil {
			v := f.env[n.args[0]]
			proved := v.hi.Sign() == 0 && !n.nonzero || v.lo.Sign() > 0 && n.nonzero
			if proved && eliminate {
				refs[hotRegionInputs+i] = add(hotRegionNode{op: PUSH32})
				report.GuardsRemoved++
				continue
			}
			refs[hotRegionInputs+i] = add(n)
			f.assume(n.args[0], n.nonzero)
			continue
		}
		if h := n.helper; h != nil {
			if result, ok := flowProveHelper(h, []hotInterval{f.env[n.args[0]], f.env[n.args[1]]}, f, n.args); ok {
				if !eliminate {
					id := add(n)
					refs[hotRegionInputs+i] = id
					v := hotHelperResultBounds(h, result, []hotInterval{f.env[n.args[0]], f.env[n.args[1]]})
					bounds[id], f.env[id] = v, v
					continue
				}
				ids := map[uint8]int{0: n.args[0], 1: n.args[1]}
				var emit func(uint8) int
				emit = func(id uint8) int {
					if v, ok := ids[id]; ok {
						return v
					}
					p := h.plan.nodes[id]
					z := hotRegionNode{op: p.op, value: p.constant}
					if p.op != PUSH32 {
						z.args = []int{emit(p.x)}
						if p.op != ISZERO && p.op != NOT {
							z.args = append(z.args, emit(p.y))
						}
					}
					ids[id] = add(z)
					return ids[id]
				}
				refs[hotRegionInputs+i] = emit(result)
				report.HelpersRemoved++
				continue
			}
		}
		refs[hotRegionInputs+i] = add(n)
	}
	for _, id := range r.outputs {
		out.outputs = append(out.outputs, refs[id])
	}
	return &out, bounds, report
}
