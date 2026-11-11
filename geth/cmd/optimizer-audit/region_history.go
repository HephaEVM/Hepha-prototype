package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"go-eth1.16.5-evm/core"
)

// Filled only from blocks that passed the online parent-continuity check.
// The ordinary header cache need not retain the beginning of a 1000-block run.
var regionCostHeaders = map[uint64]common.Hash{}

func regionCostObserveHash(n uint64, h common.Hash) {
	if h == (common.Hash{}) {
		return
	}
	if old, ok := regionCostHeaders[n]; ok && old != h {
		panic("diagnostic observed conflicting historical hashes")
	}
	regionCostHeaders[n] = h
}
func regionCostFreezeHistory() (func(), error) {
	core.ExperimentRegionHistoricalHashObserver = nil
	old := core.ReplayHistoricalHashes
	frozen := make(map[uint64]common.Hash, len(old)+len(regionCostHeaders))
	for n, h := range old {
		frozen[n] = h
	}
	for n, h := range regionCostHeaders {
		if h == (common.Hash{}) {
			return nil, fmt.Errorf("empty diagnostic header hash at %d", n)
		}
		if prior, ok := frozen[n]; ok && prior != h {
			return nil, fmt.Errorf("diagnostic header conflict at %d", n)
		}
		frozen[n] = h
	}
	core.ReplayHistoricalHashes = frozen
	return func() { core.ReplayHistoricalHashes = old }, nil
}
