package vm

import (
	"encoding/json"
	"github.com/ethereum/go-ethereum/common"
	"math/bits"
	"os"
	"strconv"
	"time"
)

var valuePolicy = os.Getenv("COLD_VALUE_POLICY") == "1"
var valueGate = os.Getenv("COLD_VALUE_GATE") == "1"
var valueRejected = map[common.Address]bool{}
var valueSynchronous = os.Getenv("COLD_VALUE_PREPARATION") == "boundary"
var valueBudget = 2 * time.Millisecond
var valueFocus = common.HexToAddress("0xc6fd8084fb9b6a0768cf943c341049edd1085b82")
var valueProfiles = map[common.Address]uint64{}

func init() {
	if !valuePolicy {
		return
	}
	if raw := os.Getenv("COLD_VALUE_BUDGET_US"); raw != "" {
		n, e := strconv.Atoi(raw)
		if e != nil || n < 1 || n > 100000 {
			panic("invalid boundary preparation budget")
		}
		valueBudget = time.Duration(n) * time.Microsecond
	}
	if path := os.Getenv("COLD_VALUE_PROFILE"); path != "" {
		raw, e := os.ReadFile(path)
		if e != nil {
			panic(e)
		}
		var file struct {
			BenefitNS map[common.Address]uint64 `json:"benefitNS"`
			Rejected  map[common.Address]bool   `json:"rejected"`
		}
		if e = json.Unmarshal(raw, &file); e != nil {
			panic(e)
		}
		valueProfiles = file.BenefitNS
		valueRejected = file.Rejected
	}
}
func valueTier(a common.Address) uint32 {
	if tier, ok := hotcallRootTier(a); ok {
		return tier
	}
	if valuePolicy && a == valueFocus {
		return 2
	}
	if valuePolicy && valueGate && valueRejected[a] {
		return 0
	}
	return hotspotRanks[a]
}
func valueBenefit(a common.Address) uint64 {
	if !valuePolicy {
		return 0
	}
	return valueProfiles[a]
}
func (s *backgroundPlan) promoteBenefit(n uint64) {
	for old := s.benefit.Load(); n > old; old = s.benefit.Load() {
		if s.benefit.CompareAndSwap(old, n) {
			return
		}
	}
}

// Retrospective measured time saving is a scheduling prior, not a semantic
// guard or a claim of per-region profitability. Live observations scale it.
// A logarithmic cap retains the base/composition and secondary-tier reservations.
func valuePriority(s *backgroundPlan) uint64 {
	benefit := min(s.benefit.Load(), uint64(time.Second))
	visits := min(s.hits.Load(), 64)
	if benefit == 0 {
		return visits
	}
	return uint64(bits.Len64(benefit*max(1, visits)))*100 + visits
}
func ValuePolicyReport() any {
	return map[string]any{"enabled": valuePolicy, "rejectPreviouslyUnprofitable": valueGate, "manualFocus": valueFocus, "profile": os.Getenv("COLD_VALUE_PROFILE"), "preparation": os.Getenv("COLD_VALUE_PREPARATION"), "compiler": os.Getenv("COLD_VALUE_COMPILER"), "budgetNS": valueBudget.Nanoseconds(), "profileScope": "Previous paired development measurements; not per-region or held-out evidence. Safety proofs are unchanged."}
}

// This path can only borrow a published, probe-free snapshot. A missing plan
// cannot allocate a work slot, promote priority, observe heat, or compile code.
func (p *BackgroundPreparation) lookupReady(e *EVM, c *Contract) *abiExecutionPlan {
	if !c.arithmeticHashTrusted {
		return nil
	}
	hash := c.CodeHash
	if p.readyPresence[hash[0]&63].Load()&(uint64(1)<<(hash[1]&63)) == 0 {
		return nil
	}
	cfg := e.Config.Arithmetic
	if cfg == nil || cfg.Engine != p.engine || e.Config.ABIScalar == nil || e.Config.ABIScalar.Engine != p.abi {
		return nil
	}
	session := &e.abiBinding
	if e.depth == 0 || session.cache == nil || session.config != cfg || session.engine != p.abi {
		var ok bool
		hash, ok = e.prepareABIBindingPolicy(c)
		if !ok {
			return nil
		}
	}
	cache := session.cache
	direct := &cache.direct[hash[0]&63]
	unwrap := func(plan *abiExecutionPlan) *abiExecutionPlan {
		if plan == nil {
			return nil
		}
		if slot := plan.preparing; slot != nil {
			return slot.cachedOnly.Load()
		}
		return plan
	}
	if direct.plan != nil && direct.hash == hash {
		return unwrap(direct.plan)
	}
	if plan := cache.plans[hash]; plan != nil {
		*direct = abiPolicySlot{hash, plan}
		return unwrap(plan)
	}
	key := session.key.arithmetic
	key.hash = hash
	p.mu.RLock()
	slot := p.slots[key]
	p.mu.RUnlock()
	if slot == nil {
		return nil
	}
	if plan := slot.plan.Load(); plan != nil {
		cache.plans[hash] = plan
		*direct = abiPolicySlot{hash, plan}
	}
	return slot.cachedOnly.Load()
}

type ValuePreparationStep struct {
	Nanoseconds              int64
	Analyses, CompileBatches int
	CreditNS                 int64
}

// Invoke between blocks, with no EVM frame active. Budget is amortized: a
// non-preemptible analysis/compile may overrun once; subsequent blocks repay
// this debt. Every actual stall, including an overrun, is reported and charged.
func (p *BackgroundPreparation) AdvanceValuePreparation() ValuePreparationStep {
	var compile func() bool
	if p.manager != nil {
		compile = p.manager.AdvanceBoundary
	}
	return p.advanceValuePreparation(compile)
}

func (p *BackgroundPreparation) advanceValuePreparation(compile func() bool) ValuePreparationStep {
	var out ValuePreparationStep
	if !p.synchronous {
		return out
	}
	budget := regionPreparationBudget()
	p.budgetCredit = min(p.budgetCredit+budget, 4*budget)
	started := time.Now()
	if tradeProgressive && p.budgetCredit > 0 && compile != nil {
		// A continuing analysis stream must not starve already proved work.
		// Publish one pending compiler batch before spending the remaining
		// window on new analysis. Its full cost is charged to the same budget.
		t := time.Now()
		if compile() {
			out.CompileBatches++
			p.budgetCredit -= time.Since(t)
		}
	}
	for out.Analyses < 4096 && p.budgetCredit > 0 {
		req := p.take()
		if req == nil {
			// Analysis may have filled the bounded compiler queue. Spend only
			// existing credit, then revisit publication after the new revision.
			// Never drain the manager or lift admission/observation gates.
			if tradeProgressive && tradeFillWindow && compile != nil {
				t := time.Now()
				if compile() {
					out.CompileBatches++
					p.budgetCredit -= time.Since(t)
					continue
				}
			}
			break
		}
		t := time.Now()
		p.prepareOne(req)
		p.budgetCredit -= time.Since(t)
		out.Analyses++
	}
	if p.budgetCredit > 0 && compile != nil && out.CompileBatches == 0 {
		t := time.Now()
		if compile() {
			out.CompileBatches++
			p.budgetCredit -= time.Since(t)
		}
	}
	out.Nanoseconds = time.Since(started).Nanoseconds()
	out.CreditNS = p.budgetCredit.Nanoseconds()
	return out
}
func (p *BackgroundPreparation) CancelValueCompilation() {
	if p.synchronous && p.manager != nil {
		p.manager.CancelBoundary()
	}
}
func (p *BackgroundPreparation) retirePending() {
	p.mu.Lock()
	defer p.mu.Unlock()
	if os.Getenv("PAPER_DEMAND_AUDIT") == "1" {
		p.candidateAudit = p.capturePendingCandidates()
	}
	for _, r := range p.queue {
		if raw := r.full; raw != nil {
			r.slot.plan.Store(readyBackgroundView(raw, r.code))
		} else if r.base != nil {
			r.slot.plan.Store(readyBackgroundView(r.base, r.code))
		} else if old := r.slot.plan.Load(); old != nil {
			v := *old
			v.preparing = nil
			r.slot.plan.Store(&v)
		}
		p.pending.Done()
	}
	p.queue = nil
}

func valueFrameTier(e *EVM, c *Contract) uint32 {
	if tier, ok := tradeFrameTier(e, c); ok {
		return tier
	}
	if valuePolicy && e.Config.PreparationTier == 0 {
		return 0
	}
	return hotcallFrameTier(e.Config.PreparationTier, hotspotFrameTier(c))
}
