package vm

import (
	"encoding/binary"
	"os"
	"strconv"
)

// Experiment-only delivery policy. All region proofs and original gas remain
// unchanged. Controls select when a proved plan is prepared or made available.
var tradeEnabled = os.Getenv("COLD_TRADE") == "1"

// Progressive preparation keeps the per-window work bound, but lets observed
// code outside historical cohorts earn preparation and later upgrades.
var tradeProgressive = tradeInt("PROGRESSIVE", 0, 0, 1) == 1
var tradeJoinUpdates = tradeInt("JOIN_UPDATES", 0, 0, 1) == 1
var tradePathUpgrades = tradeInt("PATH_UPGRADES", 0, 0, 1) == 1
var tradeBorrowCodeReserve = tradeInt("BORROW_CODE_RESERVE", 0, 0, 1) == 1
var tradeObservedFirst = tradeInt("OBSERVED_FIRST", 0, 0, 1) == 1
var tradeFillWindow = tradeInt("FILL_WINDOW", 0, 0, 1) == 1
var tradeBudget = tradeInt("BUDGET_PCT", 25, 0, 100)
var tradeBootstrapMS = tradeInt("BOOTSTRAP_MS", 20, 0, 100)
var tradeFrames = tradeInt("OBSERVATION_FRAMES", 8, 1, 64)
var tradeEntries = tradeInt("ENTRY_LIMIT", 64, 1, 256)
var tradeMinPublish = tradeInt("MIN_PUBLISH", 0, 0, 4096)
var tradeRootBonus = tradeInt("ROOT_BONUS", 1<<24, 0, 1<<24)
var tradeDependencyRepeats = tradeInt("DEPENDENCY_REPEATS", 32, 2, 64)
var tradeSlice = tradeInt("SLICE_ENTRIES", 16, 1, 16)
var tradeHeadstartMS = tradeInt("HEADSTART_MS", 0, 0, 5000)
var tradeCompileLegacy = tradeInt("COMPILE_LEGACY", 0, 0, 1) == 1
var tradeDependencies = tradeInt("DEPENDENCIES", 0, 0, 1) == 1

// ExperimentPreparationCapacity is shared by admission and reporting. This
// bounds retained code plans, independently of per-window work and compiler
// concurrency; increasing it does not admit unobserved entry analysis.
func ExperimentPreparationCapacity() int {
	if tradeProgressive {
		return tradeInt("CODE_CAPACITY", 1024, 256, 16384)
	}
	if hotspotEnabled {
		if hotcallPolicy && !deliveryEnabled {
			return 1024
		}
		return 256
	}
	return 4096
}

func tradeInt(name string, def, lo, hi int) int {
	s := os.Getenv("COLD_TRADE_" + name)
	if s == "" {
		return def
	}
	v, e := strconv.Atoi(s)
	if e != nil || v < lo || v > hi {
		panic("invalid COLD_TRADE_" + name)
	}
	return v
}
func tradeFrameTier(e *EVM, c *Contract) (uint32, bool) {
	if tradeProgressive {
		if e.Config.PreparationTier == 0 {
			return 0, true
		}
		if c.Address() == e.Config.tradeRoot {
			return e.Config.PreparationTier, true
		}
		return max(uint32(1), hotspotRanks[c.Address()]), true
	}
	if tier, ok := regionFrameTier(e, c); ok {
		return tier, true
	}
	if !tradeEnabled {
		return 0, false
	}
	if e.Config.PreparationTier == 0 {
		return 0, true
	}
	own := hotspotFrameTier(c)
	if !e.Config.tradeLegacyRoot && c.Address() != e.Config.tradeRoot {
		// A restored root formerly used cached-only lookup for ALL of its callees.
		// Do not silently promote an old-profile callee through this new root.
		if !tradeDependencies || deliveryLegacyAddress(c.Address()) {
			return 0, true
		}
		return 1, true
	}
	if own > 0 {
		return own, true
	}
	if tradeDependencies {
		return 1, true
	}
	return 0, true
}

// An immutable per-code selector set, published with a ready snapshot. Empty
// means native execution; nil means the retained policy with no extra gate.
// The selector only restricts performance admission, never authorizes a skip.
type tradeSelectorSet struct{ keys []uint64 }

func tradeInputKey(input []byte) uint64 {
	if len(input) < 4 {
		return 1<<33 | uint64(len(input))
	}
	return uint64(binary.BigEndian.Uint32(input)) + 1
}
func (s *tradeSelectorSet) accept(input []byte) bool {
	if s == nil {
		return true
	}
	key := tradeInputKey(input)
	for _, v := range s.keys {
		if v == key {
			return true
		}
	}
	return false
}
func (q *backgroundEntry) tradeRecordInput(input []byte) {
	if !tradeEnabled || tradeMinPublish == 0 || q.owner == nil || q.owner.deliveryLegacy.Load() {
		return
	}
	q.tradeSelector.CompareAndSwap(0, tradeInputKey(input))
}
func tradeSelectors(req *backgroundRequest) *tradeSelectorSet {
	if !tradeEnabled || tradeMinPublish == 0 || req.slot.deliveryLegacy.Load() {
		return nil
	}
	out := new(tradeSelectorSet)
	for _, d := range req.demands {
		if d.slot.Loaded() == nil || len(d.regions) == 0 {
			continue
		}
		n := d.regions[0].Instructions
		for _, r := range d.regions[1:] {
			n = min(n, r.Instructions)
		}
		if n < tradeMinPublish {
			continue
		}
		for _, q := range req.entries {
			if q.pc != d.regions[0].Entry || q.hits.Load() < q.threshold {
				continue
			}
			key := q.tradeSelector.Load()
			if key == 0 {
				continue
			}
			found := false
			for _, v := range out.keys {
				if key == v {
					found = true
					break
				}
			}
			if !found {
				out.keys = append(out.keys, key)
			}
		}
	}
	return out
}
func TradePolicyReport() any {
	return map[string]any{"enabled": tradeEnabled, "budgetPct": tradeBudget, "bootstrapMS": tradeBootstrapMS, "observationFrames": tradeFrames, "analysisEntriesPerNewCode": tradeEntries, "minPublishedRegionInstructions": tradeMinPublish, "rootBonus": tradeRootBonus, "dependencyRepeats": tradeDependencyRepeats, "sliceEntries": tradeSlice, "headstartMS": tradeHeadstartMS, "compileLegacy": tradeCompileLegacy, "newDependencies": tradeDependencies, "independentLegacyScheduling": true, "performanceGate": "observed selector and loaded region; instruction count is a coarse proxy, not measured CPU profit"}
}
