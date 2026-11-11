package vm

func experimentFreezeView(src *abiExecutionPlan, code []byte) *abiExecutionPlan {
	v := *src
	v.preparing = nil
	if src.abi != nil {
		b := *src.abi
		b.view = append([]byte(nil), b.view...)
		v.abi = &b
	}
	if src.arithmetic == nil {
		return &v
	}
	a := *src.arithmetic
	v.arithmetic = &a
	a.observationOwned = false
	a.observationNative = false
	a.observationRemaining = 0
	a.entryRegions = append([]*arithmeticRegion(nil), a.entryRegions...)
	a.computeEntries = append([]*computeHelperRegion(nil), a.computeEntries...)
	for i, c := range a.computeEntries {
		if c == nil || c.extension == nil {
			continue
		}
		r := *c
		x := *c.extension
		r.extension = &x
		x.probe = nil
		x.probeOnly = false
		if x.demand == nil && x.fn == nil && x.loop == nil {
			r.extension = nil
		}
		if r.extension == nil && r.address == nil && r.ordinary == nil && r.preparation == nil && r.muldiv == nil && r.rounding == nil {
			a.computeEntries[i] = nil
		} else {
			a.computeEntries[i] = &r
		}
	}
	// Resolves only loaded executors to independent ReadySlots; removes pending
	// markers, restores original opcodes, and rebuilds compact dispatch headers.
	return readyBackgroundView(&v, code)
}
