package vm

import (
	"crypto/sha256"
	"fmt"
	"os"
	"sort"
	"strings"
	"time"
)

func WindowBudgetNS() int64 { return regionPreparationBudget().Nanoseconds() }

// Settle only work admitted by the normal scope/heat/coverage/budget policy.
// A zero-work step means no presently runnable work, not that all code is covered.
func (p *BackgroundPreparation) WindowSettle() error {
	if !p.synchronous {
		return fmt.Errorf("ceiling needs boundary preparation")
	}
	start := time.Now()
	for i := 0; i < 10000; i++ {
		p.budgetCredit = 0
		q := p.AdvanceValuePreparation()
		if q.Analyses == 0 && q.CompileBatches == 0 {
			// Offline async-policy reference: keep the same admitted requests and
			// compiler policy, merely drain already submitted compiler tasks.
			// Unlike Wait(), no heat threshold or publication gate is lifted.
			if os.Getenv("PAPER_BACKGROUND_REFERENCE") == "1" && FlowPreparationResources()["pending"] > 0 {
				if _, err := WaitDemand(); err != nil {
					return err
				}
				continue
			}
			return nil
		}
		if time.Since(start) > 10*time.Minute {
			return fmt.Errorf("ceiling settle timeout")
		}
	}
	return fmt.Errorf("ceiling settle iteration limit")
}
func (p *BackgroundPreparation) WindowFingerprint() string {
	p.mu.RLock()
	defer p.mu.RUnlock()
	var rows []string
	for k, s := range p.slots {
		v := s.plan.Load()
		if v == nil {
			continue
		}
		v = experimentFreezeView(v, p.windowCodes[k])
		row := fmt.Sprintf("%x:", k.hash)
		if a := v.arithmetic; a != nil {
			for pc, id := range a.entryIndex {
				if id == 0 {
					continue
				}
				if int(id) < len(a.entryRegions) && a.entryRegions[id] != nil {
					row += fmt.Sprintf("a%d,", pc)
				}
				if int(id) < len(a.computeEntries) && a.computeEntries[id] != nil {
					row += fmt.Sprintf("c%d,", pc)
				}
			}
		}
		if v.abi != nil {
			for pc, e := range v.abi.entries {
				if e != nil {
					row += fmt.Sprintf("b%d,", pc)
				}
			}
		}
		rows = append(rows, row)
	}
	sort.Strings(rows)
	q := FlowPreparationResources()
	return fmt.Sprintf("%x/%d/%d/%d", sha256.Sum256([]byte(strings.Join(rows, ";"))), q["completed"], q["loadedFamilies"], q["pending"])
}
func (p *BackgroundPreparation) WindowFrozenConfig(c Config) Config {
	old := c.Pruning
	abi := NewABIScalarDiagnostic(old.ABI.Known)
	abi.Regions = old.ABI.Regions
	abi.Optimize = true
	frozen := &BackgroundPreparation{windowCodes: p.windowCodes, engine: old.Engine, abi: abi, closed: true, slots: map[arithmeticCacheKey]*backgroundPlan{}}
	p.mu.RLock()
	defer p.mu.RUnlock()
	for k, s := range p.slots {
		if v := s.plan.Load(); v != nil {
			v = experimentFreezeView(v, p.windowCodes[k])
			slot := &backgroundPlan{}
			slot.tracking.Store(true)
			slot.plan.Store(v)
			slot.cachedOnly.Store(v)
			frozen.slots[k] = slot
			if v.hasReadyRegion() {
				frozen.readyPresence[k.hash[0]&63].Or(uint64(1) << (k.hash[1] & 63))
			}
		}
	}
	pc := *old
	pc.ABI = abi
	pc.Preparation = frozen
	c = c.WithPruning(pc)
	c.LoopTiming = new(InterpreterLoopTiming)
	return c
}
