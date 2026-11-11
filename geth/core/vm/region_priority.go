package vm

import (
	"encoding/json"
	"github.com/ethereum/go-ethereum/common"
	"os"
	"sort"
	"time"
)

// A small, measured development profile is a performance prior only. All
// selected entries are reconstructed and proved against the full code/policy key.
var regionPriority = os.Getenv("COLD_REGION_PRIORITY") == "1"
var regionProfile = map[common.Hash]*regionCodeValue{}
var regionProfileHint [64]uint64

type regionEntryValue struct {
	Entry            uint64
	SavingNS, Weight uint64
}
type regionCodeValue struct {
	Score   uint64
	Entries []regionEntryValue
}

var regionProfileLimit = 8

func init() {
	if !regionPriority {
		return
	}
	raw, err := os.ReadFile(os.Getenv("COLD_REGION_PROFILE"))
	if err != nil {
		panic(err)
	}
	var f struct {
		Codes map[common.Hash]*regionCodeValue
	}
	if err = json.Unmarshal(raw, &f); err != nil {
		panic(err)
	}
	regionProfile = f.Codes
	for hash := range regionProfile {
		regionProfileHint[hash[0]&63] |= uint64(1) << (hash[1] & 63)
	}
}
func regionKnownCode(c *Contract) bool {
	h := c.CodeHash
	return c.arithmeticHashTrusted && regionProfileHint[h[0]&63]&(uint64(1)<<(h[1]&63)) != 0 && regionProfile[h] != nil
}
func regionFrameTier(e *EVM, c *Contract) (uint32, bool) {
	if !regionPriority {
		return 0, false
	}
	if e.Config.PreparationTier == 0 {
		return 0, true
	}
	// Keep every previously eligible frame. New callees must have measured code
	// evidence; address or membership in a router's call tree is insufficient.
	if e.Config.tradeLegacyRoot {
		if c.Address() == e.Config.tradeRoot {
			return e.Config.PreparationTier, true
		}
		if deliveryLegacyAddress(c.Address()) {
			if c.Address() == valueFocus {
				return 2, true
			}
			return hotspotRanks[c.Address()], true
		}
	}
	// Most calls beneath a restored router are outside this bounded profile.
	// Reject those with a read-only hint before consulting any address/profile map.
	// A positive hint must still match the complete hash; it cannot authorize code.
	if regionKnownCode(c) {
		if c.Address() == e.Config.tradeRoot {
			return e.Config.PreparationTier, true
		}
		return 1, true
	}
	return 0, true
}
func regionCandidates(r *backgroundRequest, entries []*backgroundEntry) []*backgroundEntry {
	if !regionPriority {
		return entries
	}
	v := regionProfile[r.key.hash]
	if r.slot.deliveryLegacy.Load() || tradeProgressive {
		// Keep the old candidate set. Move measured high-value entries first without
		// adding probes or lowering the existing observation threshold.
		if v != nil {
			weight := map[uint64]uint64{}
			for _, q := range v.Entries {
				weight[q.Entry] = q.Weight
			}
			sort.SliceStable(entries, func(i, j int) bool { return weight[entries[i].pc] > weight[entries[j].pc] })
		}
		return entries
	}
	if v == nil {
		return nil
	}
	wanted := map[uint64]regionEntryValue{}
	for _, q := range v.Entries {
		wanted[q.Entry] = q
	}
	out := make([]*backgroundEntry, 0, min(len(entries), regionProfileLimit))
	for _, q := range entries {
		if x, ok := wanted[q.pc]; ok {
			q.deliveryWeight = x.Weight
			q.threshold = 1
			out = append(out, q)
		}
	}
	sort.SliceStable(out, func(i, j int) bool {
		if out[i].deliveryWeight == out[j].deliveryWeight {
			return out[i].pc < out[j].pc
		}
		return out[i].deliveryWeight > out[j].deliveryWeight
	})
	if len(out) > regionProfileLimit {
		out = out[:regionProfileLimit]
	}
	// These entries were observed in the separately declared profiling sample.
	// Removing live discovery probes also avoids arbitrary instruction thresholds.
	for _, q := range out {
		q.hits.Store(q.threshold)
	}
	r.slot.observed.Store(uint64(len(out)))
	return out
}
func regionExtraPriority(r *backgroundRequest) uint64 {
	v := regionProfile[r.key.hash]
	if v == nil {
		return 0
	}
	return min(v.Score, uint64(1<<32)) * max(uint64(1), min(r.slot.hits.Load(), 64)) / uint64(256+len(r.code))
}

// Additive publication: retain executable old entries when a new revision has
// not produced a ready replacement. Active frames keep immutable old snapshots.
func regionPreserveReady(view, old *abiExecutionPlan, code []byte) *abiExecutionPlan {
	if !regionPriority || old == nil || old.arithmetic == nil || view.arithmetic == nil {
		return view
	}
	a, b := view.arithmetic, old.arithmetic
	a.entryRegions = append([]*arithmeticRegion(nil), a.entryRegions...)
	for pc, oid := range b.entryIndex {
		if oid == 0 {
			continue
		}
		has := b.entryRegions[oid] != nil || int(oid) < len(b.computeEntries) && b.computeEntries[oid] != nil
		if !has {
			continue
		}
		id := a.entryIndex[pc]
		if id != 0 && (a.entryRegions[id] != nil || int(id) < len(a.computeEntries) && a.computeEntries[id] != nil) {
			continue
		}
		if id == 0 {
			if len(a.entryRegions) >= 65535 {
				continue
			}
			id = uint16(len(a.entryRegions))
			a.entryIndex[pc] = id
			a.entryRegions = append(a.entryRegions, nil)
			for len(a.computeEntries) <= int(id) {
				a.computeEntries = append(a.computeEntries, nil)
			}
		}
		a.entryRegions[id] = b.entryRegions[oid]
		if int(oid) < len(b.computeEntries) {
			a.computeEntries[id] = b.computeEntries[oid]
		}
		a.dispatchCode[pc] = byte(arithmeticEntryMarker)
		if view.abi != nil {
			view.abi.view[pc] = byte(arithmeticEntryMarker)
		}
		a.nativeOnly = false
	}
	prepareCompactEntries(view)
	return view
}
func RegionPriorityReport() any {
	limit := regionProfileLimit
	if tradeProgressive {
		limit = 0
	}
	return map[string]any{"enabled": regionPriority, "profileCodes": len(regionProfile), "profilePath": os.Getenv("COLD_REGION_PROFILE"), "maxEntriesPerExtraCode": limit, "profileIsWhitelist": !tradeProgressive, "additivePublication": true, "profileScope": "Separate local region replay on development inputs. Original code/gas, no precompiled dynamic artifacts. Ranking evidence is not an independent validation set.", "windowNS": regionWindow.Nanoseconds(), "windowTimer": "before block execution, separate from lat and Process"}
}

var regionWindow = time.Duration(tradeInt("WINDOW_US", 0, 0, 10000000)) * time.Microsecond

func RegionWindowEnabled() bool { return regionPriority && regionWindow > 0 }
func regionPreparationBudget() time.Duration {
	if RegionWindowEnabled() {
		return regionWindow
	}
	return valueBudget
}

// Called by the sequential host using the parent state, before execution. Only
// immutable bytecode is copied to the worker; no StateDB is shared with it.
func (e *EVM) OfferRegionCode(address common.Address, hash common.Hash, code []byte) {
	if (!RegionWindowEnabled() && !finalBackgroundOffers) || e.Config.Pruning == nil || e.Config.Pruning.Preparation == nil || len(code) == 0 {
		return
	}
	c := &Contract{address: address}
	c.setResolvedCallCode(hash, code)
	tier := valueFrameTier(e, c)
	if tier > 0 {
		e.Config.Pruning.Preparation.lookup(e, c, tier)
	}
}

// Offering bytecode does not run analysis or compilation on the caller.
var finalBackgroundOffers = os.Getenv("FINAL_BACKGROUND") == "1"
