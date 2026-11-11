package vm

import (
	"encoding/json"
	"github.com/ethereum/go-ethereum/common"
	"os"
	"time"
)

// A deployment policy, not a semantic whitelist: every root can still execute
// natively. The benchmark supplies ranks learned strictly before its test window.
var hotspotEnabled = os.Getenv("COLD_HOTSET") == "1"
var hotspotGuard = os.Getenv("COLD_HOTSET_GUARD") != "0"
var hotspotRanks map[common.Address]uint32
var hotspotAccounts = map[common.Address]*hotspotAccount{}

type hotspotAccount struct {
	Calls, Samples, DisabledCalls uint64
	Disabled                      bool
	cells                         map[hotspotSampleKey]*hotspotCell
}
type hotspotSampleKey struct {
	Selector    [4]byte
	Length      int
	Gas, Status uint64
}
type hotspotCell struct {
	N        [2]int
	Min, Max [2]time.Duration
	Bad      int
}
type HotspotSample struct {
	account      *hotspotAccount
	selector     [4]byte
	length, mode int
	start        time.Time
}

func init() {
	if !hotspotEnabled {
		return
	}
	raw, err := os.ReadFile(os.Getenv("COLD_HOTSET_FILE"))
	if err != nil {
		panic(err)
	}
	var addresses []common.Address
	if err = json.Unmarshal(raw, &addresses); err != nil {
		panic(err)
	}
	if len(addresses) != 100 {
		panic("hotset requires 100 historical addresses")
	}
	hotspotRanks = map[common.Address]uint32{}
	for i, a := range addresses {
		if _, ok := hotspotRanks[a]; ok {
			panic("duplicate hotset address")
		}
		tier := uint32(1)
		if i < 20 {
			tier = 2
		}
		hotspotRanks[a] = tier
	}
}
func hotspotNative(cfg Config) Config {
	cfg.Pruning = nil
	cfg.Arithmetic = nil
	cfg.ABIScalar = nil
	cfg.HotRegions = nil
	cfg.RegionExecutionStats = nil
	cfg.PreparationTier = 0
	return cfg
}
func ExperimentRootConfig(cfg Config, to *common.Address) Config {
	if !hotspotEnabled || cfg.Pruning == nil {
		return cfg
	}
	if to == nil {
		return hotspotNative(cfg)
	}
	tier := valueTier(*to)
	if tradeProgressive {
		tier = max(uint32(1), hotspotRanks[*to])
	}
	if tier == 0 && !valuePolicy {
		return hotspotNative(cfg)
	}
	cfg.PreparationTier = tier
	if tradeEnabled {
		cfg.tradeRoot = *to
		cfg.tradeLegacyRoot = deliveryLegacyAddress(*to)
	}
	return cfg
}

// The context address remains the proxy's address through DELEGATECALL. Its
// implementation is still bound by actual code hash/rules, never proxy address.
// A CALL to an unrelated token/pool does not inherit the router's preparation
// entitlement; otherwise a small router list admits thousands of cold codes.
func hotspotFrameTier(c *Contract) uint32 { return valueTier(c.Address()) }

// Single block-executor owner. No per-op timing, speculative state mutation or
// replay. A control call is executed exactly once with the native interpreter.
// One in 32 calls is native; another one in 32 samples optimized execution.
func BeginHotspotSample(cfg Config, to *common.Address, input []byte) (Config, *HotspotSample) {
	cfg = ExperimentRootConfig(cfg, to)
	if !hotspotEnabled || cfg.Pruning == nil || !hotspotGuard || cfg.PreparationTier == 0 {
		return cfg, nil
	}
	a := hotspotAccounts[*to]
	if a == nil {
		a = &hotspotAccount{cells: map[hotspotSampleKey]*hotspotCell{}}
		hotspotAccounts[*to] = a
	}
	a.Calls++
	if a.Disabled {
		a.DisabledCalls++
		return hotspotNative(cfg), nil
	}
	if a.Calls%16 != 0 {
		return cfg, nil
	}
	mode := int((a.Calls / 16) % 2)
	var selector [4]byte
	copy(selector[:], input)
	s := &HotspotSample{account: a, selector: selector, length: len(input), mode: mode, start: time.Now()}
	if mode == 0 {
		cfg = hotspotNative(cfg)
	}
	return cfg, s
}
func EndHotspotSample(s *HotspotSample, gas, status uint64) {
	if s == nil {
		return
	}
	elapsed := time.Since(s.start)
	a := s.account
	a.Samples++
	key := hotspotSampleKey{s.selector, s.length, gas, status}
	c := a.cells[key]
	if c == nil {
		if len(a.cells) >= 64 {
			return
		}
		c = new(hotspotCell)
		a.cells[key] = c
	}
	// Exact gas/status and selector buckets reduce, but cannot eliminate, input
	// variation. Require two windows of disjoint latency ranges with a 15% gap.
	if c.record(s.mode, elapsed) {
		a.Disabled = true
	}
}
func (c *hotspotCell) record(mode int, d time.Duration) bool {
	if c.N[mode] >= 4 {
		return false
	}
	if c.N[mode] == 0 || d < c.Min[mode] {
		c.Min[mode] = d
	}
	if d > c.Max[mode] {
		c.Max[mode] = d
	}
	c.N[mode]++
	if c.N[0] < 4 || c.N[1] < 4 {
		return false
	}
	if c.Max[0] >= 10*time.Microsecond && c.Min[1] > c.Max[0]+c.Max[0]*15/100 {
		c.Bad++
	} else {
		c.Bad = 0
	}
	c.N = [2]int{}
	c.Min = [2]time.Duration{}
	c.Max = [2]time.Duration{}
	return c.Bad >= 2
}
func HotspotPolicyReport() any { return hotspotAccounts }

func (s *backgroundPlan) promote(tier uint32) {
	for old := s.tier.Load(); tier > old; old = s.tier.Load() {
		if s.tier.CompareAndSwap(old, tier) {
			return
		}
	}
}
