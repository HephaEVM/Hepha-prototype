package vm

import (
	"github.com/ethereum/go-ethereum/common"
	"os"
	"sort"
	"sync/atomic"
	"time"
)

// The control arm retains its original scheduler and observation path.
var deliveryEnabled = os.Getenv("COLD_SELECTIVE_DELIVERY") == "1"
var deliveryBases, deliverySlices, deliveryFiltered, deliveryDeferred atomic.Uint64
var deliveryLegacyNS, deliveryExtraNS atomic.Int64
var deliverySealed atomic.Uint64
var deliveryCompileProtected, deliveryCompileExploratory atomic.Uint64
var deliveryCompileRetries atomic.Uint64
var deliveryDistinctAdmissions = tradeInt("COMPILER_DISTINCT_ADMISSIONS", 0, 0, 1) == 1

const deliveryProbeLimit = 4096

var deliveryEntryLimit = tradeEntries
var deliveryObservationFrames = uint64(tradeFrames)
var deliveryProtectedLimit = uint64(tradeInt("COMPILER_PROTECTED_LIMIT", 8, 1, 16384))
var deliveryMinimumAdditional = tradeInt("COMPILER_MIN_ADDITIONAL", 64, 1, 4096)
var deliveryMinimumHits = uint64(tradeInt("COMPILER_MIN_HITS", 16, 1, 64))
var deliveryMinimumScore = int64(tradeInt("COMPILER_MIN_SCORE", 16384, 1, 1<<24))
var deliverySubmissionsPerSlice = tradeInt("COMPILER_PER_SLICE", 1, 1, 16)

// A hash hint may only choose native/cached-only execution. It never authorizes
// an executor; lookupReady still validates the complete execution-policy key.
type deliveryHint struct {
	hash common.Hash
	slot *backgroundPlan
}

func (p *BackgroundPreparation) deliveryReadyOnly(hash common.Hash, legacy, primary bool) bool {
	if !deliveryEnabled || legacy {
		return false
	}
	if tradeProgressive {
		// Frozen snapshots have no discovery state and must never admit work.
		if p.deliveryProbation == nil {
			return true
		}
		if h := p.deliveryHints[hash[0]].Load(); h != nil && h.hash == hash {
			return false
		}
		return !p.deliveryAdmit(hash, legacy, primary)
	}
	if h := p.deliveryHints[hash[0]].Load(); h != nil && h.hash == hash {
		s := h.slot
		if s.deliveryLegacy.Load() {
			return false
		}
		if !s.tracking.Load() || s.deliveryFrames.Load() >= deliveryObservationFrames {
			s.observe()
			return true
		}
		return false
	}
	return !p.deliveryAdmit(hash, legacy, primary)
}
func (p *BackgroundPreparation) deliveryRemember(hash common.Hash, s *backgroundPlan) {
	if deliveryEnabled && !s.deliveryLegacy.Load() {
		p.deliveryHints[hash[0]].Store(&deliveryHint{hash, s})
	}
}
func (s *backgroundPlan) deliveryObserve() {
	if deliveryEnabled && !s.deliveryLegacy.Load() && s.deliveryFrames.Load() < deliveryObservationFrames {
		if s.deliveryFrames.Add(1) == deliveryObservationFrames && s.owner != nil {
			s.owner.signal()
		}
	}
}

func deliveryLegacyAddress(a common.Address) bool {
	if tradeProgressive && hotspotRanks[a] > 0 {
		return true
	}
	if valuePolicy && a == valueFocus {
		return true
	}
	return hotspotRanks[a] > 0 && !(valuePolicy && valueGate && valueRejected[a])
}

// The admission hint is bounded and stores no bytecode. Internal code has to
// recur more often than a selected root before allocating an analysis request.
func (p *BackgroundPreparation) deliveryAdmit(hash common.Hash, legacy, primary bool) bool {
	if regionPriority && regionProfile[hash] != nil {
		return true
	}
	if !deliveryEnabled || legacy || primary {
		return true
	}
	return p.deliveryProbation.observe(hash)
}

// Charge actual analysis activity, excluding foreground-priority waits. A
// non-preemptible slice can overrun once; its debt blocks subsequent exploration.
// Legacy work is protected. No unused wall-clock idle time mints extra credits.
func (p *BackgroundPreparation) deliveryCredit() int64 {
	if tradeProgressive && p.synchronous {
		// AdvanceValuePreparation already charges every analysis and compiler
		// batch to its amortized window. Do not add a permanent cohort quota.
		return 1
	}
	if tradeEnabled && p.deliveryLegacySpent < int64(tradeHeadstartMS)*int64(time.Millisecond) {
		return -p.deliveryExtraSpent
	}
	return int64(tradeBootstrapMS)*int64(time.Millisecond) + max(0, p.deliveryLegacySpent-int64(tradeHeadstartMS)*int64(time.Millisecond))*int64(tradeBudget)/100 - p.deliveryExtraSpent
}
func (p *BackgroundPreparation) deliveryCharge(r *backgroundRequest, ns int64) {
	if !deliveryEnabled {
		return
	}
	if r.slot.deliveryLegacy.Load() {
		p.deliveryLegacySpent += ns
		deliveryLegacyNS.Add(ns)
	} else {
		p.deliveryExtraSpent += ns
		deliveryExtraNS.Add(ns)
		if r.deliverySlices == 0 {
			deliveryBases.Add(1)
		}
		r.deliverySlices++
		deliverySlices.Add(1)
	}
}

// Called with p.mu held, by the single analysis worker. New work cannot displace
// every legacy slice. Within exploration, finish observed work before opening
// another code. Unobserved candidates do not occupy an active-work permit.
func (p *BackgroundPreparation) deliveryTake(revision uint64) *backgroundRequest {
	best := [2]int{-1, -1}
	score := [2]uint64{}
	for i, r := range p.queue {
		publish := r.waitingNative && r.revision != revision
		retry := tradeProgressive && r.waitingNative && r.slot.hits.Load() >= max(deliveryMinimumHits, r.compileHeat*2) && r.slot.hits.Load() > r.compileHeat
		seal := !tradeProgressive && r.base != nil && !r.deliverySealed && r.slot.deliveryFrames.Load() >= deliveryObservationFrames && !r.slot.deliveryLegacy.Load()
		if p.draining.Load() == 0 {
			if r.base == nil {
				if r.slot.hits.Load() < p.limits.MinimumHits {
					continue
				}
			} else if r.slot.observed.Load() <= uint64(r.analyzed) && !publish && !seal && !retry {
				continue
			}
		} else if r.base != nil && r.analyzed+r.deliveryDeferred == len(r.entries) && r.waitingNative && !publish {
			continue
		}
		group := 0
		if !r.slot.deliveryLegacy.Load() {
			group = 1
		}
		if group == 1 && p.draining.Load() == 0 && p.deliveryCredit() <= 0 && !publish && !seal {
			continue
		}
		if r.base != nil && r.analyzed+r.deliveryDeferred == len(r.entries) && !publish && !retry {
			continue
		}
		n := valuePriority(r.slot)
		if group == 0 {
			// Preserve the retained stage/tier preference inside protected work.
			tier := r.slot.tier.Load()
			preferred := uint32(2)
			if p.tradeLegacyPicks%4 == 3 {
				preferred = 1
			}
			if (r.base == nil) == (p.hotsetStage[tier]%4 == 0) {
				n += 4096
			}
			if tier == preferred {
				n += 16384
			}
		} else {
			n = min(r.slot.hits.Load(), 64) * 65536 / uint64(256+len(r.code))
			if regionPriority {
				n = regionExtraPriority(r)
			}
			if r.slot.deliveryPrimary.Load() {
				n += uint64(tradeRootBonus)
			}
			if r.base != nil && !regionPriority {
				n += 1 << 40
			}
		}
		if publish {
			n += 1 << 48
		}
		if tradeProgressive && tradeObservedFirst && p.synchronous && r.base != nil && r.slot.observed.Load() > uint64(r.analyzed) {
			// Use each cohort's own turn count: a global fourth turn would
			// coincide with every reserved exploratory slice.
			picks := p.tradeLegacyPicks
			if group == 1 {
				picks = p.deliveryPicks - p.tradeLegacyPicks
			}
			if picks%4 != 3 {
				n += 1 << 49
			}
		}
		if best[group] < 0 || n > score[group] {
			best[group] = i
			score[group] = n
		}
	}
	group := 0
	period := uint64(4)
	if regionPriority {
		period = 8
	}
	if best[0] < 0 || (best[1] >= 0 && p.deliveryPicks%period == period-1) {
		group = 1
	}
	if best[group] < 0 {
		return nil
	}
	i := best[group]
	r := p.queue[i]
	p.deliveryPicks++
	if group == 0 {
		p.tradeLegacyPicks++
		p.hotsetStage[r.slot.tier.Load()]++
	}
	copy(p.queue[i:], p.queue[i+1:])
	p.queue[len(p.queue)-1] = nil
	p.queue = p.queue[:len(p.queue)-1]
	return r
}

// Bounded, syntax-only work estimate before symbolic composition. This is a
// ranking hint, never a proof or a claim of measured CPU savings. PUSH data is
// skipped. Region formation still uses the unchanged validator and full bounds.
func deliveryPotential(code []byte, pc uint64) uint64 {
	var score uint64
	for n := 0; pc < uint64(len(code)) && n < 64; n++ {
		op := OpCode(code[pc])
		pc++
		switch {
		case op >= PUSH1 && op <= PUSH32:
			pc += uint64(op - PUSH0)
			score++
		case op == PUSH0 || op >= DUP1 && op <= SWAP16:
			score++
		case op == MUL || op == DIV || op == SDIV || op == MOD || op == SMOD:
			score += 8
		case op == KECCAK256:
			score += 24
		case op == MSTORE || op == MSTORE8 || op == MLOAD:
			score += 4
		case op == CALLDATACOPY || op == CODECOPY || op == RETURNDATACOPY || op == MCOPY:
			score += 16
		case op == STOP || op == RETURN || op == REVERT || op == INVALID || op == SSTORE || op == SLOAD || op == CALL || op == STATICCALL || op == DELEGATECALL || op == CREATE || op == CREATE2:
			return score
		case op == JUMP || op == JUMPI:
			return score + 4
		default:
			score += 2
		}
	}
	return score
}

// Cover the observed route before cutting the expensive analysis set. A larger
// observation envelope is short-lived; symbolic work still has the same cap.
func deliveryAdditional(d *demandPlan) int {
	coverage := func(x *demandPlan) int {
		if x == nil || len(x.regions) == 0 {
			return 0
		}
		n := x.regions[0].Instructions
		for _, r := range x.regions[1:] {
			n = min(n, r.Instructions)
		}
		return n
	}
	old := 0
	for f := d.fallback; f != nil; f = f.fallback {
		if f.slot.Loaded() != nil {
			old = max(old, coverage(f))
		}
	}
	delta := max(0, coverage(d)-old)
	if tradePathUpgrades {
		delta = max(delta, joinedPathAdditional(d))
	}
	return delta
}
func deliveryCompileAdmit(r *backgroundRequest, d *demandPlan, started bool) bool {
	if !deliveryEnabled || started {
		return true
	}
	if tradeEnabled && !tradeCompileLegacy && r.slot.deliveryLegacy.Load() {
		return false
	}
	return deliveryAdditional(d) >= deliveryMinimumAdditional && (!r.slot.deliveryLegacy.Load() || tradeProgressive && deliveryDistinctAdmissions && d.compileBudgetCounted || deliveryCompileProtected.Load() < deliveryProtectedLimit)
}
func deliveryCompileRecord(r *backgroundRequest, d *demandPlan) {
	if !deliveryEnabled {
		return
	}
	if tradeProgressive && deliveryDistinctAdmissions {
		if d.compileBudgetCounted {
			deliveryCompileRetries.Add(1)
			return
		}
		d.compileBudgetCounted = true
	}
	if r.slot.deliveryLegacy.Load() {
		deliveryCompileProtected.Add(1)
	} else {
		deliveryCompileExploratory.Add(1)
	}
}

func deliveryCandidates(r *backgroundRequest, entries []*backgroundEntry) []*backgroundEntry {
	if regionPriority {
		return regionCandidates(r, entries)
	}
	if !deliveryEnabled || r.slot.deliveryLegacy.Load() {
		return entries
	}
	for _, q := range entries {
		if tradeEnabled {
			q.threshold = 1
		}
		q.deliveryWeight = deliveryPotential(r.code, q.pc)
	}
	sort.SliceStable(entries, func(i, j int) bool {
		if entries[i].deliveryWeight == entries[j].deliveryWeight {
			return entries[i].pc < entries[j].pc
		}
		return entries[i].deliveryWeight > entries[j].deliveryWeight
	})
	if len(entries) > deliveryProbeLimit {
		deliveryFiltered.Add(uint64(len(entries) - deliveryProbeLimit))
		entries = entries[:deliveryProbeLimit]
	}
	return entries
}
func (r *backgroundRequest) deliverySelect() map[uint64]bool {
	selected := make(map[uint64]bool)
	if !r.deliverySealed && r.slot.deliveryFrames.Load() >= deliveryObservationFrames {
		for _, q := range r.entries {
			if !q.analyzed && !q.deliveryDeferred && q.hits.Load() < q.threshold {
				q.deliveryDeferred = true
				r.deliveryDeferred++
				deliveryDeferred.Add(1)
			}
		}
		r.deliverySealed = true
		r.published = -1
		deliverySealed.Add(1)
		if r.slot.owner != nil && r.slot.owner.deliveryCredit() <= 0 {
			return selected
		}
	}
	for _, q := range r.entries {
		if q.analyzed || q.deliveryDeferred || q.hits.Load() < q.threshold {
			continue
		}
		q.analyzed = true
		r.analyzed++
		selected[q.pc] = true
		if len(selected) == tradeSlice || r.analyzed == deliveryEntryLimit {
			break
		}
	}
	if r.analyzed >= deliveryEntryLimit {
		for _, q := range r.entries {
			if !q.analyzed && !q.deliveryDeferred {
				q.deliveryDeferred = true
				r.deliveryDeferred++
				deliveryDeferred.Add(1)
			}
		}
	}
	return selected
}
func DeliveryPolicyReport() any {
	capacity, entries := ExperimentPreparationCapacity(), deliveryEntryLimit
	frames := deliveryObservationFrames
	if tradeProgressive {
		entries, frames = 0, 0
	}
	return map[string]any{
		"enabled": deliveryEnabled, "progressiveObservedScope": tradeProgressive, "codeCapacity": capacity, "dependencyRepeatThreshold": tradeDependencyRepeats,
		"explorationBootstrapNS": int64(tradeBootstrapMS) * int64(time.Millisecond), "explorationShareOfLegacyActivity": float64(tradeBudget) / 100,
		"probeLimitPerNewCode": deliveryProbeLimit, "analysisLimitPerNewCode": entries,
		"observationFrames": frames, "zeroAnalysisOrObservationLimitMeans": "no lifetime cap; only observed entries, bounded analysis slices", "sealedCodeObservations": deliverySealed.Load(),
		"newCodeBases": deliveryBases.Load(), "newCodeSlices": deliverySlices.Load(),
		"filteredProbeCandidates": deliveryFiltered.Load(), "deferredAnalysisEntries": deliveryDeferred.Load(),
		"legacyActivityNS": deliveryLegacyNS.Load(), "explorationActivityNS": deliveryExtraNS.Load(),
		"compilerOutstandingLimit": tradeInt("COMPILER_OUTSTANDING_LIMIT", 8, 1, 128), "compilerLifetimeLimit": tradeInt("COMPILER_TOTAL_LIMIT", 32, 1, 16384),
		"compilerMinimumAdditionalInstructions": deliveryMinimumAdditional, "compilerProtectedLifetimeLimit": deliveryProtectedLimit,
		"pathSpecificJoinUpgradeProxy": tradePathUpgrades,
		"prioritizeObservedEntries":    tradeProgressive && tradeObservedFirst,
		"borrowUnusedCodeReserve":      tradeProgressive && tradeBorrowCodeReserve,
		"distinctCandidateAccounting":  tradeProgressive && deliveryDistinctAdmissions, "retrySubmissionsNotCharged": deliveryCompileRetries.Load(), "compilerMinimumHits": deliveryMinimumHits, "compilerMinimumScore": deliveryMinimumScore, "compilerSubmissionsPerSlice": deliverySubmissionsPerSlice,
		"protectedCompileSubmissions": deliveryCompileProtected.Load(), "exploratoryCompileSubmissions": deliveryCompileExploratory.Load(),
		"ranking": "bounded syntax cost and observed reuse; not measured per-region CPU savings",
	}
}
