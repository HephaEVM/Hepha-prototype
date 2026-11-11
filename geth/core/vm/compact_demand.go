package vm

import "go-eth1.16.5-evm/regionexec"

// The executable header is private to a sequential ABI engine. Warm bindings
// may fill fn once when their lazy slot becomes ready; background publication
// builds a separate header with fn already loaded. Neither path changes the
// analysis plan, opcode view, gas rules or validated fallbacks.
type compactDemandEntry struct {
	fn      regionexec.Executor
	exit    uint64
	helpers []regionexec.Helper
	plan    *demandPlan
}

func prepareCompactEntries(v *abiExecutionPlan) {
	if v == nil || v.arithmetic == nil {
		return
	}
	a := v.arithmetic
	a.compactEntries = nil
	for id, c := range a.computeEntries {
		if c == nil || c.extension == nil {
			continue
		}
		p := c.extension.demand
		if p == nil || p.slot == nil || len(p.counts) == 0 {
			continue
		}
		fn := p.slot.Loaded()
		if a.compactEntries == nil {
			a.compactEntries = make([]compactDemandEntry, len(a.computeEntries))
		}
		a.compactEntries[id] = compactDemandEntry{fn, p.exit, p.entryHelpers, p}
	}
}

func (p *compactDemandEntry) try(pc *uint64, s *ScopeContext, e *EVM) bool {
	stats := e.Config.RegionExecutionStats
	if stats != nil {
		stats.DemandAttempts++
		stats.recordDemand(p.plan, s.Contract, -1)
	}
	if p.fn == nil {
		p.fn = p.plan.slot.Ready()
		if p.fn == nil {
			return p.plan.tryOriginal(pc, s, e)
		}
	}
	if e.demandFrame == nil {
		e.demandFrame = new(regionexec.Frame)
	}
	f := e.demandFrame
	f.Stack = s.Stack.data
	f.Input = s.Contract.Input
	f.Memory = s.Memory.store
	f.Gas = s.Contract.Gas
	f.MemoryGas = s.Memory.lastGasCost
	f.Exit = p.exit
	f.Helpers = p.helpers
	f.Path = 0
	if p.plan.branch && !p.plan.shared {
		f.Exits = [2]uint64{p.plan.regions[0].Exit, p.plan.regions[1].Exit}
		f.BranchHelpers = [2][]regionexec.Helper{p.plan.helpers[0], p.plan.helpers[1]}
	}
	if !p.fn(f) {
		return p.plan.tryOriginal(pc, s, e)
	}
	s.Stack.data = f.Stack
	s.Memory.store = f.Memory
	s.Memory.lastGasCost = f.MemoryGas
	s.Contract.Gas = f.Gas
	*pc = f.PC
	if stats != nil {
		plan := p.plan
		q := &plan.counts[f.Path]
		stats.DemandCommits++
		stats.recordDemand(plan, s.Contract, int(f.Path))
		stats.DemandInstructions += q.instructions
		stats.ForwardedLoads += q.forwarded
		stats.SkippedStores += q.stores
		if len(plan.flow) > 0 {
			stats.FlowHelpers += q.helpers
			stats.FlowGuards += q.guards
			if plan.shared {
				stats.FlowShared++
				if len(plan.regions) > 2 {
					stats.FlowMulti++
				}
			}
		}
		if plan.branch {
			stats.BranchCommits++
		}
	}
	return true
}
