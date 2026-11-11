package vm

// A preparation owns one immutable path. Proof, accounting and lowering share
// its optimized forms; no sampled runtime value enters this cache.
type preparedFlow struct {
	original                  *hotRegion
	plain, proved             *hotRegion
	plainBounds, provedBounds []hotInterval
}

func prepareFlows(rs []*hotRegion) []*preparedFlow {
	out := make([]*preparedFlow, len(rs))
	for i, r := range rs {
		out[i] = &preparedFlow{original: r}
	}
	return out
}
func (p *preparedFlow) optimized(proof bool) (*hotRegion, []hotInterval) {
	if proof {
		if p.proved == nil {
			p.proved, p.provedBounds, _ = optimizeFlowRegion(p.original)
		}
		return p.proved, p.provedBounds
	}
	if p.plain == nil {
		p.plain, _, p.plainBounds, _ = optimizeHotRegion(p.original)
	}
	return p.plain, p.plainBounds
}
func preparedFlowsChanged(ps []*preparedFlow) bool {
	for _, p := range ps {
		old, _ := p.optimized(false)
		next, _ := p.optimized(true)
		if old.fingerprint() != next.fingerprint() {
			return true
		}
	}
	return false
}
func flowBitmap(code []byte, shared []BitVec) BitVec {
	if len(shared) > 0 {
		return shared[0]
	}
	return codeBitmap(code)
}
