package vm

import (
	"fmt"
	"go-eth1.16.5-evm/regionexec/prepare"
	"os"
	"sync"
	"sync/atomic"
	"time"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/params"
)

// BackgroundPreparation never receives StateDB, stack, memory or transaction
// values. Each frame chooses one immutable snapshot. Already compiled helpers
// are published before expensive composition; later frames can upgrade safely.
type BackgroundPreparation struct {
	windowCodes                                      map[arithmeticCacheKey][]byte
	tradeLegacyPicks                                 uint64
	deliveryLegacySpent, deliveryExtraSpent          int64
	deliveryPicks                                    uint64
	deliveryProbation                                *preparationProbation
	deliveryHints                                    [256]atomic.Pointer[deliveryHint]
	synchronous                                      bool
	budgetCredit                                     time.Duration
	readyPresence                                    [64]atomic.Uint64
	hotsetPicks                                      uint64 // worker only, one secondary slice per four choices
	hotsetStage                                      [3]uint64
	probation                                        *preparationProbation
	engine                                           *ArithmeticEngine
	abi                                              *ABIScalarDiagnostic
	analyzer                                         *ArithmeticEngine
	mu                                               sync.RWMutex
	slots                                            map[arithmeticCacheKey]*backgroundPlan
	queue                                            []*backgroundRequest
	candidateAudit                                   []map[string]any // optional immutable shutdown diagnostic
	wake, stop, done                                 chan struct{}
	closed                                           bool
	frozenReadyOnly                                  bool // immutable diagnostic snapshot; never admits work
	frozenLeanFallback                               bool // diagnostic: match warm loop for a prepared empty plan
	limit                                            int
	limits                                           BackgroundPreparationLimits
	draining                                         atomic.Int64
	manager                                          *prepare.Manager
	activity                                         *prepare.Activity
	priorityIdle                                     time.Duration // worker only; excluded from active preparation time
	pending                                          sync.WaitGroup
	enqueued, earlyReady, ready, fallback, dropped   atomic.Uint64
	preparationNS, earlyNS, fullNS                   atomic.Int64
	candidateCount, observedEntries, analyzedEntries atomic.Uint64
}
type backgroundPlan struct {
	deliveryLegacy, deliveryPrimary atomic.Bool
	deliveryFrames                  atomic.Uint64
	tier                            atomic.Uint32
	benefit                         atomic.Uint64
	cachedOnly                      atomic.Pointer[abiExecutionPlan]
	plan                            atomic.Pointer[abiExecutionPlan]
	owner                           *BackgroundPreparation
	hits                            atomic.Uint64
	observed                        atomic.Uint64
	analyzed                        atomic.Uint64
	tracking                        atomic.Bool
}
type backgroundRequest struct {
	deliverySlices, deliveryDeferred int
	deliverySealed                   bool
	slot                             *backgroundPlan
	key                              arithmeticCacheKey
	code                             []byte
	rules                            params.Rules
	// Kept only between the two stages; no analysis is repeated.
	evm           *EVM
	base          *abiExecutionPlan
	full          *abiExecutionPlan
	demands       []*demandPlan
	published     int
	entries       []*backgroundEntry
	analyzed      int
	waitingNative bool
	revision      uint64
	compileHeat   uint64 // worker-owned; retry admission after observed heat grows
}
type BackgroundPreparationStats struct {
	Enqueued, EarlyReady, Ready, NativeFrames, Dropped   uint64
	PreparationNS, EarlyPreparationNS, FullPreparationNS int64
	Pending                                              int
	CandidateEntries, ObservedEntries, AnalyzedEntries   uint64
	PriorityWaits                                        uint64
	PriorityWaitNS                                       int64
}

// Limits affect preparation only. MinimumHits counts code-level frame visits;
// EntryMinimumHits counts arrivals at an individual candidate. An entry threshold
// of zero means one arrival. WorkPercent is a cooperative wall-time budget; an
// offline Wait lifts observation gates and the budget.
type BackgroundPreparationLimits struct {
	MinimumHits, EntryMinimumHits uint64
	WorkPercent                   int
}

func NewBackgroundPreparation(engine *ArithmeticEngine, abi *ABIScalarDiagnostic, capacity int) (*BackgroundPreparation, error) {
	return NewBackgroundPreparationWithLimits(engine, abi, capacity, BackgroundPreparationLimits{4, 2, 100})
}
func NewBackgroundPreparationWithLimits(engine *ArithmeticEngine, abi *ABIScalarDiagnostic, capacity int, limits BackgroundPreparationLimits) (*BackgroundPreparation, error) {
	if engine == nil || abi == nil || !abi.Optimize || capacity < 1 || limits.WorkPercent < 0 || limits.WorkPercent > 100 {
		return nil, fmt.Errorf("background preparation needs arithmetic and optimized ABI libraries and positive capacity")
	}
	analyzer, err := NewArithmeticEngine(engine.templates, ArithmeticEngineOptions{CacheEntries: 64, FeedbackEntries: 128})
	if err != nil {
		return nil, err
	}
	p := &BackgroundPreparation{engine: engine, abi: abi, analyzer: analyzer, slots: make(map[arithmeticCacheKey]*backgroundPlan), wake: make(chan struct{}, 1), stop: make(chan struct{}), done: make(chan struct{}), limit: capacity, limits: limits, manager: getDemandManager()}
	if p.manager != nil {
		p.manager.SetDutyPercent(limits.WorkPercent)
		p.activity = p.manager.Activity()
	}
	if p.activity == nil {
		p.activity = prepare.NewActivity()
	}
	if cooperativePreparation {
		p.analyzer.preparationYield = p.yield
	}
	if deliveryEnabled {
		p.deliveryProbation = &preparationProbation{threshold: uint64(tradeDependencyRepeats)}
	}
	p.synchronous = valueSynchronous || RegionWindowEnabled() || os.Getenv("PAPER_BACKGROUND_REFERENCE") == "1"
	go p.worker()
	return p, nil
}

// These optional host boundaries also protect work outside EVM.Run, such as
// transaction processing or a native prefetch pass sharing the same process.
func (p *BackgroundPreparation) EnterForeground() {
	if p.activity != nil {
		p.activity.Begin()
	}
}
func (p *BackgroundPreparation) LeaveForeground() {
	if p.activity != nil {
		p.activity.End()
	}
}

// Count only observations made while preparation is outstanding. A fully warm
// lookup has no queue lock or counter update. Never predict heat from addresses.
func latestBackgroundPlan(plan *abiExecutionPlan, tiers ...uint32) *abiExecutionPlan {
	if slot := plan.preparing; slot != nil {
		if len(tiers) > 0 {
			slot.promote(tiers[0])
		}
		slot.observe()
		if latest := slot.plan.Load(); latest != nil {
			if slot.owner != nil && latest.arithmetic == nil && latest.abi == nil {
				slot.owner.fallback.Add(1)
			}
			return latest
		}
	}
	return plan
}

func (p *BackgroundPreparation) lookup(e *EVM, c *Contract, tiers ...uint32) *abiExecutionPlan {
	tier := e.Config.PreparationTier
	if len(tiers) > 0 {
		tier = tiers[0]
	}
	benefit := valueBenefit(c.Address())
	legacy, primary := false, false
	if deliveryEnabled {
		if tradeEnabled && c.Address() == e.Config.tradeRoot {
			legacy = e.Config.tradeLegacyRoot
			primary = e.Config.PreparationTier == 2 || tradeProgressive
		} else {
			legacy = deliveryLegacyAddress(c.Address())
			primary = hotspotRanks[c.Address()] == 2 || c.Address() == valueFocus
		}
	}
	mark := func(s *backgroundPlan) {
		if s == nil {
			return
		}
		if legacy && !s.deliveryLegacy.Load() {
			s.deliveryLegacy.Store(true)
		}
		if primary && !s.deliveryPrimary.Load() {
			s.deliveryPrimary.Store(true)
		}
	}
	cfg := e.Config.Arithmetic
	if cfg == nil || cfg.Engine != p.engine || e.Config.ABIScalar == nil || e.Config.ABIScalar.Engine != p.abi {
		return nil
	}
	if c.arithmeticHashTrusted && p.deliveryReadyOnly(c.CodeHash, legacy, primary) {
		return p.lookupReady(e, c)
	}
	session := &e.abiBinding
	hash := c.CodeHash
	if e.depth == 0 || session.cache == nil || session.config != cfg || session.engine != p.abi {
		var ok bool
		hash, ok = e.prepareABIBindingPolicy(c)
		if !ok {
			return nil
		}
	} else if !c.arithmeticHashTrusted {
		hash = crypto.Keccak256Hash(c.Code)
	}
	policy := session.cache
	direct := &policy.direct[hash[0]&63]
	if direct.plan != nil && direct.hash == hash {
		if direct.plan.preparing != nil {
			mark(direct.plan.preparing)
			direct.plan.preparing.promoteBenefit(benefit)
		}
		ready := latestBackgroundPlan(direct.plan, tier)
		if ready != direct.plan {
			direct.plan = ready
			policy.plans[hash] = ready
		}
		cfg.Engine.stats.hits.Add(1)
		return ready
	}
	if ready := policy.plans[hash]; ready != nil {
		if ready.preparing != nil {
			mark(ready.preparing)
			ready.preparing.promoteBenefit(benefit)
		}
		ready = latestBackgroundPlan(ready, tier)
		policy.plans[hash] = ready
		*direct = abiPolicySlot{hash, ready}
		cfg.Engine.stats.hits.Add(1)
		return ready
	}
	key := session.key.arithmetic
	key.hash = hash
	publish := func(slot *backgroundPlan) *abiExecutionPlan {
		mark(slot)
		slot.promoteBenefit(benefit)
		slot.promote(tier)
		slot.observe()
		ready := slot.plan.Load()
		if ready != nil {
			policy.plans[hash] = ready
			*direct = abiPolicySlot{hash, ready}
		}
		if ready == nil || (ready.arithmetic == nil && ready.abi == nil) {
			p.fallback.Add(1)
		}
		return ready
	}
	p.mu.RLock()
	slot, closed := p.slots[key], p.closed
	p.mu.RUnlock()
	if slot != nil {
		return publish(slot)
	}
	if closed {
		return nil
	}
	if !c.arithmeticHashTrusted && !p.deliveryAdmit(hash, legacy, primary) {
		p.fallback.Add(1)
		return nil
	}
	if p.probation != nil && !(regionPriority && regionProfile[hash] != nil) && !p.probation.observe(hash) {
		p.fallback.Add(1)
		return nil
	}
	p.mu.Lock()
	defer p.mu.Unlock()
	if slot = p.slots[key]; slot != nil {
		return publish(slot)
	}
	// Optional progressive admission can use otherwise idle reserved slots.
	// The total capacity, observation gates and preparation budgets still apply.
	reserveBlocked := hotspotEnabled && tier == 1 && len(p.slots) >= p.limit/2 && (!deliveryEnabled || !legacy)
	if p.closed || len(p.slots) >= p.limit || reserveBlocked && !(tradeProgressive && tradeBorrowCodeReserve) {
		p.dropped.Add(1)
		p.fallback.Add(1)
		return nil
	}
	slot = &backgroundPlan{owner: p}
	mark(slot)
	p.deliveryRemember(hash, slot)
	slot.promoteBenefit(benefit)
	slot.promote(tier)
	initial := uint64(1)
	if p.probation != nil {
		initial = max(p.probation.threshold, p.limits.MinimumHits)
	}
	slot.hits.Store(initial)
	placeholder := &abiExecutionPlan{preparing: slot}
	slot.plan.Store(placeholder)
	policy.plans[hash] = placeholder
	*direct = abiPolicySlot{hash, placeholder}
	req := &backgroundRequest{slot: slot, key: key, code: append([]byte(nil), c.Code...), rules: e.chainRules}
	p.pending.Add(1)
	p.slots[key] = slot
	if p.windowCodes == nil {
		p.windowCodes = make(map[arithmeticCacheKey][]byte)
	}
	p.windowCodes[key] = req.code
	p.queue = append(p.queue, req)
	p.enqueued.Add(1)
	p.fallback.Add(1)
	p.signal()
	return nil
}
func (s *backgroundPlan) observe() {
	if s.tracking.Load() {
		s.deliveryObserve()
		if hotcallPolicy && s.hits.Load() < 64 {
			s.hits.Add(1)
		}
		return
	}
	n := s.hits.Add(1)
	if p := s.owner; p != nil && n == p.limits.MinimumHits {
		p.signal()
	}
}
func (p *BackgroundPreparation) signal() {
	select {
	case p.wake <- struct{}{}:
	default:
	}
}

// First publish completed native work, then rotate bounded observed-entry
// slices. Historical code heat cannot let one code monopolize the worker.
// Basic analysis retains a bounded code-heat preference.
func (p *BackgroundPreparation) take() *backgroundRequest {
	p.mu.Lock()
	defer p.mu.Unlock()
	if p.closed || len(p.queue) == 0 {
		return nil
	}
	revision := uint64(0)
	if p.manager != nil {
		revision = p.manager.Revision()
	}
	if deliveryEnabled {
		return p.deliveryTake(revision)
	}
	preferred := uint32(2)
	if hotspotEnabled {
		p.hotsetPicks++
		if p.hotsetPicks%4 == 0 {
			preferred = 1
		}
	}
	score := func(r *backgroundRequest) uint64 {
		if hotspotEnabled {
			tier := r.slot.tier.Load()
			score := min(r.slot.hits.Load(), 64)
			if valuePolicy {
				score = valuePriority(r.slot)
			}
			// Within each tier, alternate delivery stages: 1 base slice to 3
			// composition slices. Always sweeping bases can starve useful execution.
			baseTurn := p.hotsetStage[tier]%4 == 0
			if (r.base == nil) == baseTurn {
				score += 4096
			}
			if r.waitingNative && r.revision != revision {
				score += 8192
			}
			if tier == preferred {
				score += 16384
			}
			return score
		}
		if r.base == nil {
			return min(r.slot.hits.Load(), 64) * 4
		}
		if r.waitingNative && r.revision != revision {
			return 1024
		}
		return 512
	}
	best := -1
	priority := uint64(0)
	for i, r := range p.queue {
		if p.draining.Load() != 0 {
			// Offline drain admits unobserved entries, but still waits for the
			// compiler instead of spinning after all analysis has completed.
			if r.base != nil && r.analyzed == len(r.entries) && r.waitingNative && r.revision == revision {
				continue
			}
		} else {
			if r.base == nil {
				if r.slot.hits.Load() < p.limits.MinimumHits {
					continue
				}
			} else if r.slot.observed.Load() <= uint64(r.analyzed) && (!r.waitingNative || r.revision == revision) {
				continue
			}
		}
		if n := score(r); best < 0 || n > priority {
			best, priority = i, n
		}
	}
	if best < 0 {
		return nil
	}
	req := p.queue[best]
	if hotspotEnabled {
		p.hotsetStage[req.slot.tier.Load()]++
	}
	copy(p.queue[best:], p.queue[best+1:])
	p.queue[len(p.queue)-1] = nil
	p.queue = p.queue[:len(p.queue)-1]
	return req
}
func (p *BackgroundPreparation) prepareBase(req *backgroundRequest) {
	abi := NewABIScalarDiagnostic(p.abi.Known)
	abi.Optimize = true
	abi.Regions = p.abi.Regions
	abi.preparationYield = p.analyzer.preparationYield
	cfg := Config{}.WithPruning(PruningConfig{Engine: p.analyzer, ABI: abi})
	if lib := req.key.ordinary; lib != nil {
		cfg.HotRegions = &HotRegionRuntime{codes: lib.codes, engine: p.analyzer}
	}
	e := &EVM{Config: cfg, table: req.key.rules, chainRules: req.rules, preparationYield: p.yield}
	e.preparationCandidateYield = p.analyzer.preparationYield
	c := &Contract{}
	c.setResolvedCallCode(req.key.hash, req.code)
	a := e.arithmeticBinding(c)
	view := req.code
	if a != nil {
		view = a.dispatchCode
	}
	p.yield()
	req.evm = e
	req.base = &abiExecutionPlan{arithmetic: a, abi: abi.bind(c, e.table, a, view, false)}
	req.entries = p.candidateEntries(req)
	req.slot.tracking.Store(true)
	req.published = -1
	p.publishReady(req)
	p.earlyReady.Add(1)
}
func (p *BackgroundPreparation) yield() {
	if p.synchronous {
		return
	}
	if p.activity == nil {
		return
	}
	start := time.Now()
	p.activity.Wait(p.stop, &p.draining)
	p.priorityIdle += time.Since(start)
}
func (p *BackgroundPreparation) worker() {
	defer close(p.done)
	if p.synchronous {
		<-p.stop
		p.retirePending()
		return
	}
	tick := time.NewTicker(50 * time.Millisecond)
	defer tick.Stop()
	for {
		req := p.take()
		if req == nil {
			select {
			case <-p.stop:
				p.mu.Lock()
				for _, r := range p.queue {
					// Retire observation callbacks as well as the upgrade pointer.
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
				p.mu.Unlock()
				return
			case <-p.wake:
			case <-tick.C:
			}
			continue
		}
		if !p.activity.Wait(p.stop, &p.draining) {
			p.mu.Lock()
			p.queue = append(p.queue, req)
			p.mu.Unlock()
			continue
		}
		elapsed := p.prepareOne(req)
		p.rest(elapsed)
	}
}
func (p *BackgroundPreparation) rest(elapsed time.Duration) {
	percent := p.limits.WorkPercent
	if percent <= 0 || percent >= 100 || p.draining.Load() > 0 {
		return
	}
	timer := time.NewTimer(elapsed * time.Duration(100-percent) / time.Duration(percent))
	defer timer.Stop()
	for p.draining.Load() == 0 {
		select {
		case <-timer.C:
			return
		case <-p.stop:
			return
		case <-p.wake:
		}
	}
}

// Wait is an explicit offline drain after foreground admission has stopped.
// It lifts heat thresholds and pacing only for the caller's validation phase.
func (p *BackgroundPreparation) Wait() {
	p.draining.Add(1)
	p.signal()
	if p.manager != nil {
		done := p.manager.Unthrottle()
		defer done()
	}
	p.pending.Wait()
	p.draining.Add(-1)
}
func (p *BackgroundPreparation) Close() {
	p.mu.Lock()
	if !p.closed {
		p.closed = true
		close(p.stop)
	}
	p.mu.Unlock()
	<-p.done
}
func (p *BackgroundPreparation) Stats() BackgroundPreparationStats {
	p.mu.RLock()
	pending := len(p.queue)
	p.mu.RUnlock()
	var waits uint64
	var ns int64
	if p.activity != nil {
		waits, ns = p.activity.Stats()
	}
	return BackgroundPreparationStats{Enqueued: p.enqueued.Load(), EarlyReady: p.earlyReady.Load(), Ready: p.ready.Load(), NativeFrames: p.fallback.Load(), Dropped: p.dropped.Load(), PreparationNS: p.preparationNS.Load(), EarlyPreparationNS: p.earlyNS.Load(), FullPreparationNS: p.fullNS.Load(), Pending: pending, CandidateEntries: p.candidateCount.Load(), ObservedEntries: p.observedEntries.Load(), AnalyzedEntries: p.analyzedEntries.Load(), PriorityWaits: waits, PriorityWaitNS: ns}
}

func (p *BackgroundPreparation) prepareOne(req *backgroundRequest) time.Duration {
	started := time.Now()
	idleBefore := p.priorityIdle
	finished := false
	if req.base == nil {
		p.prepareBase(req)
		elapsed := (time.Since(started) - (p.priorityIdle - idleBefore)).Nanoseconds()
		p.earlyNS.Add(elapsed)
		p.preparationNS.Add(elapsed)
		finished = len(req.entries) == 0
	} else {
		var selected map[uint64]bool
		publishPending := req.waitingNative && p.manager != nil && req.revision != p.manager.Revision()
		if !publishPending || (tradeProgressive && tradeObservedFirst && p.synchronous && req.slot.observed.Load() > uint64(req.analyzed)) {
			selected = req.selectObserved(p.draining.Load() > 0)
		}
		if len(selected) > 0 {
			base := req.full
			if base == nil {
				base = req.base
			}
			req.full = req.evm.bindComposedEntries(base, req.code, selected)
			req.demands = backgroundDemands(req.full)
			req.published = -1 // also retire the probes just analyzed
			p.analyzedEntries.Add(uint64(len(selected)))
			req.slot.analyzed.Add(uint64(len(selected)))
			elapsed := (time.Since(started) - (p.priorityIdle - idleBefore)).Nanoseconds()
			p.fullNS.Add(elapsed)
			p.preparationNS.Add(elapsed)
		}
		// These are only observed entries and their validated fallbacks.
		// Admission remains on this worker; no foreground queue access.
		hotcallActivate(req)
		finished = p.publishReady(req)
	}
	p.deliveryCharge(req, (time.Since(started) - (p.priorityIdle - idleBefore)).Nanoseconds())
	if finished {
		p.ready.Add(1)
		p.pending.Done()
	} else {
		p.mu.Lock()
		p.queue = append(p.queue, req)
		p.mu.Unlock()
	}
	return time.Since(started) - (p.priorityIdle - idleBefore)
}
