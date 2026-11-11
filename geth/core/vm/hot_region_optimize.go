package vm

// Offline only: intervals describe ALL possible words, never sampled operands.
// Overflow, signed operations and unknown effects widen to the full word range.
import (
	"math/big"

	"github.com/holiman/uint256"
)

type hotInterval struct{ lo, hi *big.Int }

var hotWordMax = new(big.Int).Sub(new(big.Int).Lsh(big.NewInt(1), 256), big.NewInt(1))

func hotFull() hotInterval               { return hotInterval{new(big.Int), new(big.Int).Set(hotWordMax)} }
func hotExact(w uint256.Int) hotInterval { return hotInterval{w.ToBig(), w.ToBig()} }
func (x hotInterval) exact() bool        { return x.lo.Cmp(x.hi) == 0 }
func hotBool(v bool) hotInterval {
	if v {
		return hotExact(*uint256.NewInt(1))
	}
	return hotExact(uint256.Int{})
}
func hotSmallBool() hotInterval { return hotInterval{new(big.Int), big.NewInt(1)} }

func hotTransfer(op OpCode, a []hotInterval) hotInterval {
	constant := len(a) > 0
	words := make([]uint256.Int, len(a))
	for i, x := range a {
		constant = constant && x.exact()
		words[i].SetFromBig(x.lo)
	}
	// Only pure deterministic word operations. Calldata and memory are unknown.
	if constant && hotRegionPure(op) && op != CALLDATALOAD && op != CALLDATASIZE {
		return hotExact(hotRegionValue(op, words, nil))
	}
	if len(a) == 0 {
		return hotFull()
	}
	x := a[0]
	if op == ISZERO {
		if x.hi.Sign() == 0 {
			return hotBool(true)
		}
		if x.lo.Sign() > 0 {
			return hotBool(false)
		}
		return hotSmallBool()
	}
	if len(a) < 2 {
		return hotFull()
	}
	y := a[1]
	z := hotFull()
	switch op {
	case ADD, MUL:
		if op == ADD {
			z.lo.Add(x.lo, y.lo)
			z.hi.Add(x.hi, y.hi)
		} else {
			z.lo.Mul(x.lo, y.lo)
			z.hi.Mul(x.hi, y.hi)
		}
		if z.hi.BitLen() <= 256 {
			return z
		}
	case SUB:
		if x.lo.Cmp(y.hi) >= 0 {
			return hotInterval{new(big.Int).Sub(x.lo, y.hi), new(big.Int).Sub(x.hi, y.lo)}
		}
	case AND:
		z.hi.Set(x.hi)
		if y.hi.Cmp(z.hi) < 0 {
			z.hi.Set(y.hi)
		}
		return z
	case DIV:
		if y.lo.Sign() > 0 {
			return hotInterval{new(big.Int).Div(x.lo, y.hi), new(big.Int).Div(x.hi, y.lo)}
		}
	case MOD:
		if y.hi.Sign() > 0 {
			z.hi.Sub(y.hi, big.NewInt(1))
			if x.hi.Cmp(z.hi) < 0 {
				z.hi.Set(x.hi)
			}
			return z
		}
	case SHL, SHR:
		if x.exact() {
			if x.lo.Cmp(big.NewInt(256)) >= 0 {
				return hotBool(false)
			}
			shift := uint(x.lo.Uint64())
			if op == SHR {
				return hotInterval{new(big.Int).Rsh(y.lo, shift), new(big.Int).Rsh(y.hi, shift)}
			}
			z.lo.Lsh(y.lo, shift)
			z.hi.Lsh(y.hi, shift)
			if z.hi.BitLen() <= 256 {
				return z
			}
		}
	case LT, GT:
		if op == GT {
			x, y = y, x
		}
		if x.hi.Cmp(y.lo) < 0 {
			return hotBool(true)
		}
		if x.lo.Cmp(y.hi) >= 0 {
			return hotBool(false)
		}
		return hotSmallBool()
	case EQ:
		if x.hi.Cmp(y.lo) < 0 || y.hi.Cmp(x.lo) < 0 {
			return hotBool(false)
		}
		return hotSmallBool()
	case SLT, SGT:
		return hotSmallBool()
	}
	return hotFull()
}

// Prove the actual bound helper decision tree, not merely its operation name.
// Splitting each operand at zero proves Solidity's MUL/DIV identity including
// the zero case. At most four partitions; any unresolved failure leaf rejects
// fusion. Each accepted leaf must retain the same result and original gas.
func hotProveHelper(h *arithmeticRegion, inputs []hotInterval) (uint8, bool) {
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

type hotOptimizationReport struct{ Constants, GuardConstants, FusedHelpers, RetainedHelpers int }

// Keep the original proof/PC/gas/peak metadata. Only generated computations
// change. Unproved helpers remain exact calls and atomic fallback boundaries.
func optimizeHotRegion(r *hotRegion) (*hotRegion, map[int]int, []hotInterval, hotOptimizationReport) {
	return optimizeHotRegionChecks(r, !paperWithoutCheckElimination)
}

func optimizeHotRegionChecks(r *hotRegion, eliminate bool) (*hotRegion, map[int]int, []hotInterval, hotOptimizationReport) {
	out := *r
	out.nodes = nil
	out.outputs = nil
	var report hotOptimizationReport
	guardDeps := hotGuardDependencies(r)
	preserveCondition, guardCondition := false, false
	origins := map[int]int{}
	bounds := make([]hotInterval, hotRegionInputs)
	refs := make([]int, hotRegionInputs+len(r.nodes))
	for i := 0; i < hotRegionInputs; i++ {
		bounds[i] = hotFull()
		refs[i] = i
	}
	var add func(hotRegionNode) int
	add = func(n hotRegionNode) int {
		v := hotFull()
		if n.op == PUSH32 {
			v = hotExact(n.value)
		} else if n.helper == nil && n.op != MSTORE && n.op != JUMPI {
			a := make([]hotInterval, len(n.args))
			for i, id := range n.args {
				a[i] = bounds[id]
			}
			v = hotTransfer(n.op, a)
			if v.exact() && !preserveCondition {
				if guardCondition {
					report.GuardConstants++
				}
				n = hotRegionNode{op: PUSH32}
				n.value.SetFromBig(v.lo)
				report.Constants++
			}
		}
		id := hotRegionInputs + len(out.nodes)
		out.nodes = append(out.nodes, n)
		bounds = append(bounds, v)
		return id
	}
	for i, n := range r.nodes {
		guardCondition = guardDeps[hotRegionInputs+i]
		preserveCondition = !eliminate && guardCondition
		n.args = append([]int(nil), n.args...)
		for j, a := range n.args {
			n.args[j] = refs[a]
		}
		if h := n.helper; h != nil {
			if result, ok := hotProveHelper(h, []hotInterval{bounds[n.args[0]], bounds[n.args[1]]}); ok {
				if !eliminate {
					id := add(n)
					refs[hotRegionInputs+i] = id
					origins[id-hotRegionInputs] = i
					bounds[id] = hotHelperResultBounds(h, result, []hotInterval{bounds[n.args[0]], bounds[n.args[1]]})
					report.RetainedHelpers++
					continue
				}
				// Import only the result slice. Discharged guard DAGs never
				// become runtime instructions or separately computed results.
				ids := map[uint8]int{0: n.args[0], 1: n.args[1]}
				var importNode func(uint8) int
				importNode = func(id uint8) int {
					if x, ok := ids[id]; ok {
						return x
					}
					p := h.plan.nodes[id]
					x := hotRegionNode{op: p.op, value: p.constant}
					if p.op != PUSH32 {
						x.args = []int{importNode(p.x)}
						if p.op != ISZERO && p.op != NOT {
							x.args = append(x.args, importNode(p.y))
						}
					}
					ids[id] = add(x)
					return ids[id]
				}
				refs[hotRegionInputs+i] = importNode(result)
				report.FusedHelpers++
				continue
			}
			report.RetainedHelpers++
		}
		id := add(n)
		refs[hotRegionInputs+i] = id
		if n.helper != nil {
			origins[id-hotRegionInputs] = i
		}
	}
	for _, id := range r.outputs {
		out.outputs = append(out.outputs, refs[id])
	}
	return &out, origins, bounds, report
}

// Retaining a runtime check does not discard the already proved result range.
func hotHelperResultBounds(h *arithmeticRegion, result uint8, inputs []hotInterval) hotInterval {
	memo := map[uint8]hotInterval{0: inputs[0], 1: inputs[1]}
	var visit func(uint8) hotInterval
	visit = func(id uint8) hotInterval {
		if v, ok := memo[id]; ok {
			return v
		}
		n := h.plan.nodes[id]
		v := hotFull()
		if n.op == PUSH32 {
			v = hotExact(n.constant)
		} else {
			a := []hotInterval{visit(n.x)}
			if n.op != ISZERO && n.op != NOT {
				a = append(a, visit(n.y))
			}
			v = hotTransfer(n.op, a)
		}
		memo[id] = v
		return v
	}
	return visit(result)
}

// A no-elimination control must retain the predicate computation as well as
// JUMPI. Keeping a branch whose condition was folded to a constant would not
// restore the check that the experiment purports to measure.
func hotGuardDependencies(r *hotRegion) map[int]bool {
	out := map[int]bool{}
	var visit func(int)
	visit = func(id int) {
		if id < hotRegionInputs || out[id] {
			return
		}
		out[id] = true
		for _, a := range r.nodes[id-hotRegionInputs].args {
			visit(a)
		}
	}
	for _, n := range r.nodes {
		if n.helper == nil && n.op == JUMPI {
			visit(n.args[0])
		}
	}
	return out
}
