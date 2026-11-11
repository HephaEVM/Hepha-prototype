package vm

import (
	"fmt"
	"go-eth1.16.5-evm/regionexec"
	"go-eth1.16.5-evm/regionexec/prepare"
	"os"
	"strings"
	"sync"
)

type demandPlan struct {
	compileBudgetCounted bool // preparation worker only; retries are not new candidates
	slot                 *prepare.Slot
	exit                 uint64
	entryHelpers         []regionexec.Helper
	counts               []demandPathCounts
	fallback             *demandPlan // original finite region, used only if a joined plan refuses admission
	joinSource           *demandPlan // preparation-only, unjoined source for incremental pair composition
	regions              []*hotRegion
	helpers              [][]regionexec.Helper
	forward, dead        []int
	branch               bool
	shared               bool
	flow                 []flowReport
	joinParents          []*hotRegion // preparation-only provenance, one original path per joined leaf
}

// The executable view is immutable once published. Execution need not visit
// the analysis graph and four separate slices to recover per-path constants.
type demandPathCounts struct {
	instructions, helpers, guards, forwarded, stores uint64
}

func (p *demandPlan) finishBinding() *demandPlan {
	p.exit = p.regions[0].Exit
	p.entryHelpers = p.helpers[0]
	p.counts = make([]demandPathCounts, len(p.regions))
	for i, r := range p.regions {
		q := &p.counts[i]
		q.instructions = uint64(r.Instructions)
		q.forwarded, q.stores = uint64(p.forward[i]), uint64(p.dead[i])
		if i < len(p.flow) {
			q.helpers = uint64(max(0, p.flow[i].HelpersRemoved))
			q.guards = uint64(max(0, p.flow[i].GuardsRemoved))
		}
	}
	return p
}

var demandOnce sync.Once
var demandManager *prepare.Manager
var demandError error

func getDemandManager() *prepare.Manager {
	root, dir := os.Getenv("DEMAND_ROOT"), os.Getenv("DEMAND_DIR")
	// An optional preparation service may be disabled initially. A missing
	// configuration must not consume the one-time initialization forever.
	if root == "" || dir == "" {
		return nil
	}
	demandOnce.Do(func() {
		demandManager, demandError = prepare.NewManager(root, dir)
	})
	return demandManager
}
func WaitDemand() ([]prepare.Event, error) {
	m := getDemandManager()
	if m == nil {
		if demandError == nil {
			return nil, fmt.Errorf("DEMAND_ROOT and DEMAND_DIR required")
		}
		return nil, demandError
	}
	m.Drain()
	return m.Events(), nil
}
func makeDemandPlan(rs []*hotRegion, memory bool) *demandPlan {
	return makePreparedDemandPlan(prepareFlows(rs), memory)
}
func makePreparedDemandPlan(prepared []*preparedFlow, memory bool) *demandPlan {
	rs := make([]*hotRegion, len(prepared))
	for i, p := range prepared {
		rs[i] = p.original
	}
	m := getDemandManager()
	if m == nil {
		return nil
	}
	p := &demandPlan{regions: rs, branch: len(rs) == 2}
	var src strings.Builder
	key := fmt.Sprintf("demand-v1-mem%t", memory)
	if paperCEExtendedProof {
		key += "-extended-ce-v1"
	}
	if paperWithoutCheckElimination {
		key += "-checks-retained"
	}
	for i, r := range rs {
		body, helpers, err := emitPreparedDemand(prepared[i], memory)
		if err != nil {
			panic(err)
		}
		p.helpers = append(p.helpers, helpers)
		rr, _ := prepared[i].optimized(false)
		f, d := map[int]int{}, map[int]bool{}
		if memory {
			f, d = memoryDemand(rr)
		}
		p.forward = append(p.forward, len(f))
		p.dead = append(p.dead, len(d))
		key += "/" + r.key
		if len(rs) == 2 {
			body = strings.ReplaceAll(body, "Exec", fmt.Sprintf("Exec_%d", i))
		}
		src.WriteString(body)
	}
	if len(rs) == 2 {
		src.WriteString("\nfunc Exec(s *regionexec.Frame)bool{old:=s.Exit;s.Exit=s.Exits[0];s.Helpers=s.BranchHelpers[0];if Exec_0(s){s.Path=0;return true};s.Exit=s.Exits[1];s.Helpers=s.BranchHelpers[1];if Exec_1(s){s.Path=1;return true};s.Exit=old;return false}\n")
	}
	text := src.String()
	p.slot = m.Slot(key, func() (string, error) { return text, nil })
	if p.slot == nil {
		return nil
	}
	return p.finishBinding()
}
func (evm *EVM) bindDemand(b *boundArithmetic, code []byte, pc uint64, x *regionExtension, bitmaps ...BitVec) *demandPlan {
	if evm.preparationYield != nil {
		evm.preparationYield()
	}
	if experimentSkipRejectedAnalysis() {
		return nil
	}
	return evm.bindFlow(b, code, pc, x, bitmaps...)
}
func (p *demandPlan) try(pc *uint64, s *ScopeContext, e *EVM) bool {
	stats := e.Config.RegionExecutionStats
	if stats != nil {
		stats.DemandAttempts++
		stats.recordDemand(p, s.Contract, -1)
	}
	fn := p.slot.Published()
	if fn == nil {
		fn = p.slot.Ready()
		if fn == nil {
			return p.tryOriginal(pc, s, e)
		}
	}
	if e.demandFrame == nil {
		e.demandFrame = &regionexec.Frame{}
	}
	f := e.demandFrame
	f.Stack = s.Stack.data
	f.Input = s.Contract.Input
	f.Memory = s.Memory.store
	f.Gas = s.Contract.Gas
	f.MemoryGas = s.Memory.lastGasCost
	f.Exit = p.exit
	f.Helpers = p.entryHelpers
	f.Path = 0
	if p.branch && !p.shared {
		f.Exits = [2]uint64{p.regions[0].Exit, p.regions[1].Exit}
		f.BranchHelpers = [2][]regionexec.Helper{p.helpers[0], p.helpers[1]}
	}
	if !fn(f) {
		return p.tryOriginal(pc, s, e)
	}
	s.Stack.data = f.Stack
	s.Memory.store = f.Memory
	s.Memory.lastGasCost = f.MemoryGas
	s.Contract.Gas = f.Gas
	*pc = f.PC
	if stats != nil {
		stats.DemandCommits++
		stats.recordDemand(p, s.Contract, int(f.Path))
		q := &p.counts[f.Path]
		if len(p.flow) > 0 {
			stats.FlowHelpers += q.helpers
			stats.FlowGuards += q.guards
			if p.shared {
				stats.FlowShared++
				if len(p.regions) > 2 {
					stats.FlowMulti++
				}
			}
		}
		stats.DemandInstructions += q.instructions
		if p.branch {
			stats.BranchCommits++
		}
		stats.ForwardedLoads += q.forwarded
		stats.SkippedStores += q.stores
	}
	return true
}

func (p *demandPlan) tryOriginal(pc *uint64, s *ScopeContext, e *EVM) bool {
	if p.fallback != nil {
		return p.fallback.try(pc, s, e)
	}
	return false
}

func FlowPreparationResources() map[string]int {
	if m := getDemandManager(); m != nil {
		return m.Resources()
	}
	return nil
}
