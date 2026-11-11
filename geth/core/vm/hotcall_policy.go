package vm

import (
	"github.com/ethereum/go-ethereum/common"
	"os"
	"sort"
	"sync/atomic"
)

// Root cohorts stay historical. A selected transaction may prepare reused code
// in its callees, including proxy implementations. Ordinary roots only borrow.
var hotcallPolicy = os.Getenv("COLD_HOTCALL_POLICY") == "1"
var hotcallSelective = os.Getenv("COLD_VALUE_COMPILER") == "selective"
var hotcallCompileEligible, hotcallCompileDeferred atomic.Uint64
var hotcallReused atomic.Uint64
var progressiveJoinRevisits atomic.Uint64
var compilerTierWeight = tradeInt("COMPILER_TIER_WEIGHT", 0, 0, 8)

func hotcallRootTier(a common.Address) (uint32, bool) {
	if hotcallPolicy && hotspotRanks[a] == 2 {
		return 2, true
	}
	return 0, false
}
func hotcallFrameTier(root, own uint32) uint32 {
	if root == 0 {
		return 0
	}
	if hotcallPolicy {
		return max(uint32(1), own)
	}
	return own
}

// Selection affects delivery only, never the semantic proof. A loaded fallback
// remains executable. Retain at most the best unknown upgrade per entry, and
// give each code at most four submissions per analysis/publication slice.
func hotcallActivate(req *backgroundRequest) {
	req.compileHeat = req.slot.hits.Load()
	if tradeProgressive {
		for _, d := range req.demands {
			if d.slot.Loaded() == nil && d.slot.Reuse() != nil {
				hotcallReused.Add(1)
			}
		}
	}
	if !hotcallSelective {
		for _, d := range req.demands {
			hotcallPrioritize(req, d)
			d.slot.Ready()
		}
		return
	}
	coverage := func(d *demandPlan) int {
		if len(d.regions) == 0 {
			return 0
		}
		n := d.regions[0].Instructions
		for _, r := range d.regions[1:] {
			n = min(n, r.Instructions)
		}
		return n
	}
	covered := map[uint64]int{}
	for _, d := range req.demands {
		if len(d.regions) > 0 && d.slot.Loaded() != nil {
			entry := d.regions[0].Entry
			covered[entry] = max(covered[entry], coverage(d))
		}
	}
	best := map[uint64]*demandPlan{}
	for _, d := range req.demands {
		if d.slot.Loaded() != nil || len(d.regions) == 0 {
			continue
		}
		if deliveryEnabled && req.slot.hits.Load() < deliveryMinimumHits {
			continue
		}
		if coverage(d) <= covered[d.regions[0].Entry] {
			continue
		}
		score := backgroundPriority(d, req.slot.hits.Load())
		threshold := int64(4096)
		if deliveryEnabled {
			threshold = deliveryMinimumScore
		}
		if score < threshold {
			hotcallCompileDeferred.Add(1)
			continue
		}
		entry := d.regions[0].Entry
		if prev := best[entry]; prev == nil || score > backgroundPriority(prev, req.slot.hits.Load()) {
			best[entry] = d
		}
	}
	ds := make([]*demandPlan, 0, len(best))
	for _, d := range best {
		ds = append(ds, d)
	}
	sort.Slice(ds, func(i, j int) bool {
		a, b := backgroundPriority(ds[i], req.slot.hits.Load()), backgroundPriority(ds[j], req.slot.hits.Load())
		if a == b {
			return ds[i].regions[0].Entry < ds[j].regions[0].Entry
		}
		return a > b
	})
	submitted := 0
	for _, d := range ds {
		// Existing shared work does not spend another code's new-work allowance.
		pending := d.slot.PreparationStarted()
		if !deliveryCompileAdmit(req, d, pending) {
			hotcallCompileDeferred.Add(1)
			continue
		}
		limit := 4
		if deliveryEnabled {
			limit = deliverySubmissionsPerSlice
		}
		if !pending && submitted == limit {
			hotcallCompileDeferred.Add(1)
			continue
		}
		hotcallPrioritize(req, d)
		d.slot.Ready()
		if !pending && d.slot.PreparationStarted() {
			submitted++
			hotcallCompileEligible.Add(1)
			deliveryCompileRecord(req, d)
		}
	}
}
func hotcallPrioritize(req *backgroundRequest, d *demandPlan) {
	score := backgroundPriority(d, req.slot.hits.Load())
	if hotspotEnabled {
		score = hotcallTierPriority(score, req.slot.tier.Load())
	}
	if valuePolicy && !deliveryEnabled {
		score += int64(valuePriority(req.slot)) << 28
	}
	d.slot.Prioritize(score)
}
func hotcallTierPriority(score int64, tier uint32) int64 {
	if tradeProgressive && compilerTierWeight > 0 {
		// Cohort is a bounded preference, not a lexicographic barrier: a
		// substantial callee can outrank a small root-code upgrade.
		if tier >= 2 {
			return score * int64(compilerTierWeight)
		}
		return score
	}
	return score + int64(tier)<<40
}
func HotcallPolicyReport() any {
	capacity := ExperimentPreparationCapacity()
	return map[string]any{
		"enabled": hotcallPolicy, "restoreHistoricalTop20": hotcallPolicy, "inheritPreparationInSelectedCalls": hotcallPolicy, "codeCapacity": capacity, "dependencyCapacity": capacity / 2, "rootTier": 2, "dependencyTier": 1,
		"selectiveCompiler": hotcallSelective, "newSubmissions": hotcallCompileEligible.Load(), "deferredDecisions": hotcallCompileDeferred.Load(), "reusedWithoutNewCompilation": hotcallReused.Load(),
		"incrementalJoinUpdates": tradeProgressive && tradeJoinUpdates, "joinedPredecessorUpdates": progressiveJoinRevisits.Load(),
		"boundedCompilerTierWeight": compilerTierWeight, "scoreThreshold": 4096, "perCodeSliceSubmissions": 4, "scoreMeaning": "bounded observed reuse times minimum additional instruction coverage / graph size; scheduling proxy, not measured CPU saving",
	}
}
