package vm

import (
	"go-eth1.16.5-evm/regionexec/prepare"
	"sort"
)

// Publication must include ABI-only plans: arithmeticBinding can legitimately
// return nil while ABI binding still contains executable regions. This scan is
// outside execution, when an immutable plan is published or captured.
func (p *abiExecutionPlan) hasReadyRegion() bool {
	if p == nil {
		return false
	}
	if p.arithmetic != nil && !p.arithmetic.nativeOnly {
		return true
	}
	if p.abi != nil {
		for _, entry := range p.abi.entries {
			if entry != nil {
				return true
			}
		}
	}
	return false
}

func backgroundDemands(p *abiExecutionPlan) []*demandPlan {
	var out []*demandPlan
	seen := map[*demandPlan]bool{}
	var add func(*demandPlan)
	add = func(d *demandPlan) {
		if d == nil || seen[d] {
			return
		}
		seen[d] = true
		if !experimentWorkEligible(d) {
			add(d.fallback)
			return
		}
		out = append(out, d)
		add(d.fallback)
	}
	if p.arithmetic != nil {
		for _, c := range p.arithmetic.computeEntries {
			if c != nil && c.extension != nil {
				add(c.extension.demand)
			}
		}
	}
	// Prefer substantial already-proved regions; no contract/address whitelist.
	sort.SliceStable(out, func(i, j int) bool { return out[i].regions[0].Instructions > out[j].regions[0].Instructions })
	return out
}

func (p *BackgroundPreparation) publishReady(req *backgroundRequest) bool {
	if p.manager != nil {
		req.revision = p.manager.Revision()
	}
	ready := 0
	settled := true
	for _, d := range req.demands {
		if d.slot.Loaded() != nil {
			ready++
		}
		if !d.slot.Settled() {
			settled = false
		}
	}
	wasWaiting := req.waitingNative
	req.waitingNative = !settled
	complete := settled && req.analyzed+req.deliveryDeferred == len(req.entries)
	if ready != req.published || wasWaiting != req.waitingNative || complete {
		raw := req.full
		if raw == nil {
			raw = req.base
		}
		if valuePolicy {
			cached := readyBackgroundView(raw, req.code)
			cached.preparing = nil
			cached.tradeAllowed = tradeSelectors(req)
			cached = regionPreserveReady(cached, req.slot.cachedOnly.Load(), req.code)
			req.slot.cachedOnly.Store(cached)
			if cached.hasReadyRegion() && (cached.tradeAllowed == nil || len(cached.tradeAllowed.keys) > 0) {
				cell := &p.readyPresence[req.key.hash[0]&63]
				bit := uint64(1) << (req.key.hash[1] & 63)
				cell.Or(bit)
			}
		}
		view := readyBackgroundView(raw, req.code)
		view = regionPreserveReady(view, req.slot.cachedOnly.Load(), req.code)
		observeBackgroundView(view, req.entries)
		if complete {
			view.tradeAllowed = tradeSelectors(req)
		}
		if !complete {
			view.preparing = req.slot
		}
		req.slot.plan.Store(view)
		req.published = ready
	}
	return complete
}

// Copy-on-publication: active frames retain their old opcode view. Pending
// native regions do not get entry markers, while ready shorter helpers remain.
func readyBackgroundView(full *abiExecutionPlan, code []byte) *abiExecutionPlan {
	view := *full
	if full.arithmetic == nil {
		return &view
	}
	a := *full.arithmetic
	view.arithmetic = &a
	a.dispatchCode = append([]byte(nil), a.dispatchCode...)
	a.entryIndex = append([]uint16(nil), a.entryIndex...)
	a.computeEntries = append([]*computeHelperRegion(nil), a.computeEntries...)
	if full.abi != nil {
		abi := *full.abi
		abi.view = append([]byte(nil), abi.view...)
		view.abi = &abi
	}
	resolved := map[*demandPlan]*demandPlan{}
	var resolve func(*demandPlan) *demandPlan
	resolve = func(d *demandPlan) *demandPlan {
		if d == nil {
			return nil
		}
		if r, ok := resolved[d]; ok {
			return r
		}
		fallback := resolve(d.fallback)
		if !experimentWorkEligible(d) {
			resolved[d] = fallback
			return fallback
		}
		fn := d.slot.Loaded()
		if fn == nil {
			resolved[d] = fallback
			return fallback
		}
		r := *d
		r.slot = prepare.ReadySlot(fn)
		r.fallback = fallback
		resolved[d] = &r
		return &r
	}
	for id, c := range a.computeEntries {
		if c == nil || c.extension == nil || c.extension.demand == nil {
			continue
		}
		r := *c
		x := *c.extension
		x.demand = resolve(x.demand)
		r.extension = &x
		if x.demand == nil && x.loop == nil && x.fn == nil {
			r.extension = nil
		}
		if r.extension == nil && r.address == nil && r.ordinary == nil && r.preparation == nil && r.muldiv == nil {
			a.computeEntries[id] = nil
		} else {
			a.computeEntries[id] = &r
		}
	}
	a.nativeOnly = true
	for pc, id := range a.entryIndex {
		if id == 0 {
			continue
		}
		if a.entryRegions[id] == nil && (int(id) >= len(a.computeEntries) || a.computeEntries[id] == nil) {
			a.entryIndex[pc] = 0
			a.dispatchCode[pc] = code[pc]
			if view.abi != nil {
				view.abi.view[pc] = code[pc]
				if view.abi.entries[uint64(pc)] != nil {
					view.abi.view[pc] = byte(abiScalarMarker)
				}
			}
		} else {
			a.nativeOnly = false
		}
	}
	if a.nativeOnly && view.abi != nil {
		for _, r := range view.abi.entries {
			if r != nil {
				a.nativeOnly = false
				break
			}
		}
	}
	prepareCompactEntries(&view)
	return &view
}
