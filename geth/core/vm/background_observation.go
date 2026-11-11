package vm

import (
	"sort"
	"sync/atomic"
)

// A probe observes an entry, not stack values. Only the first few visits write
// counters. It cannot request compilation or change EVM state, gas or control.
type backgroundEntry struct {
	tradeSelector    atomic.Uint64
	deliveryWeight   uint64
	deliveryDeferred bool
	pc, threshold    uint64
	owner            *backgroundPlan
	hits             atomic.Uint64
	analyzed         bool // preparation worker only
}

func (q *backgroundEntry) record(_ *uint64, scope *ScopeContext, _ *hotRegion) bool {
	if scope != nil && scope.Contract != nil {
		q.tradeRecordInput(scope.Contract.Input)
	}
	q.saturateObservation()
	return false
}

// Use exactly the candidate entry families already consumed by composition.
// This is a cheap bytecode walk; symbolic analysis waits for observed reuse.
func (p *BackgroundPreparation) candidateEntries(req *backgroundRequest) []*backgroundEntry {
	if req.base.arithmetic == nil {
		return nil
	}
	starts := partialStarts(req.code, true)
	for _, pc := range summaryCandidateEntries(req.code) {
		starts[pc] = true
	}
	for pc := uint64(0); pc < uint64(len(req.code)); pc++ {
		op := OpCode(req.code[pc])
		if op >= PUSH1 && op <= PUSH32 {
			pc += uint64(op - PUSH0)
			continue
		}
		if op == KECCAK256 || req.base.arithmetic.entryIndex[pc] != 0 {
			starts[pc] = true
		}
	}
	entries := make([]*backgroundEntry, 0, len(starts))
	for pc := range starts {
		if req.base.abi != nil && req.base.abi.entries[pc] != nil {
			continue // composition also leaves existing ABI entries alone
		}
		entries = append(entries, &backgroundEntry{pc: pc, threshold: max(1, p.limits.EntryMinimumHits), owner: req.slot})
	}
	sort.Slice(entries, func(i, j int) bool { return entries[i].pc < entries[j].pc })
	entries = deliveryCandidates(req, entries)
	p.candidateCount.Add(uint64(len(entries)))
	return entries
}

func (r *backgroundRequest) selectObserved(drain bool) map[uint64]bool {
	if deliveryEnabled && !tradeProgressive && !drain && !r.slot.deliveryLegacy.Load() {
		return r.deliverySelect()
	}
	selected := make(map[uint64]bool)
	for _, q := range r.entries {
		if q.analyzed || q.deliveryDeferred || !drain && q.hits.Load() < q.threshold {
			continue
		}
		q.analyzed = true
		r.analyzed++
		selected[q.pc] = true
		if len(selected) == 16 {
			break // bounded analysis slice; publish before examining more entries
		}
	}
	return selected
}

// Reuse the existing outlined fallback callback, which always returns false.
// No new test is added to the interpreter loop or to warm region execution.
// Probe-only visits increment PartialAttempts, but never PartialCommits or any
// skipped-instruction counter. Snapshots own their markers and cloned tables.
func observeBackgroundView(view *abiExecutionPlan, entries []*backgroundEntry) {
	a := view.arithmetic
	if a == nil {
		return
	}
	a.observationOwned = true
	a.observationNative = a.nativeOnly
	a.observationRemaining = 0
	a.entryRegions = append([]*arithmeticRegion(nil), a.entryRegions...)
	for len(a.computeEntries) < len(a.entryRegions) {
		a.computeEntries = append(a.computeEntries, nil)
	}
	for _, q := range entries {
		if q.analyzed || q.deliveryDeferred || q.hits.Load() >= q.threshold {
			continue
		}
		id := a.entryIndex[q.pc]
		if id == 0 {
			if len(a.entryRegions) >= 65535 {
				continue
			}
			id = uint16(len(a.entryRegions))
			a.entryIndex[q.pc] = id
			a.entryRegions = append(a.entryRegions, nil)
			a.computeEntries = append(a.computeEntries, nil)
		}
		var c computeHelperRegion
		if a.computeEntries[id] != nil {
			c = *a.computeEntries[id]
		}
		// Unanalyzed entries have no composed extension yet. Their existing
		// arithmetic, address, mulDiv and ordinary helpers stay in the fallback.
		x := regionExtension{}
		if c.extension != nil {
			x = *c.extension
		}
		x.probe = q
		x.probeOnly = x.demand == nil && x.fn == nil && x.loop == nil && c.address == nil && c.ordinary == nil && c.muldiv == nil && c.preparation == nil && c.rounding == nil && a.entryRegions[id] == nil
		if x.probeOnly {
			a.observationRemaining++
		}
		c.extension = &x
		a.computeEntries[id] = &c
		a.dispatchCode[q.pc] = byte(arithmeticEntryMarker)
		if view.abi != nil {
			view.abi.view[q.pc] = byte(arithmeticEntryMarker)
		}
		a.nativeOnly = false
	}
}
