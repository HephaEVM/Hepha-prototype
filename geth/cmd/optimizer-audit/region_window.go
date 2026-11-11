package main

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"go-eth1.16.5-evm/config"
	"go-eth1.16.5-evm/core"
	"go-eth1.16.5-evm/core/state"
	"go-eth1.16.5-evm/core/types"
	"go-eth1.16.5-evm/core/vm"
	"os"
	"time"
)

// No execution/tracing/prefetch oracle: only root recipients known from this
// block's transaction list and code available in its parent state. Dependencies
// first discovered during execution join later windows through the normal queue.
func regionBlockWindow(block *types.Block, s *state.StateDB, processor *core.StateProcessor, cfg vm.Config, preparers []*vm.BackgroundPreparation) map[string]any {
	if !vm.RegionWindowEnabled() && os.Getenv("FINAL_BACKGROUND") != "1" {
		return nil
	}
	started := time.Now()
	offered := 0
	var addresses []common.Address
	counts := map[common.Address]int{}
	for _, tx := range block.Transactions() {
		if tx.To() != nil {
			addr := *tx.To()
			if counts[addr] == 0 {
				addresses = append(addresses, addr)
			}
			counts[addr]++
		}
	}
	for _, addr := range addresses {
		c := vm.ExperimentRootConfig(cfg, &addr)
		if c.PreparationTier == 0 {
			continue
		}
		code := s.GetCode(addr)
		if target, ok := types.ParseDelegation(code); ok {
			code = s.GetCode(target)
		}
		if len(code) == 0 {
			continue
		}
		e := vm.NewEVM(processor.ExperimentPreparationContext(block.Header()), s, config.MainnetChainConfig, c)
		hash := crypto.Keccak256Hash(code)
		for i := 0; i < min(counts[addr], 4); i++ {
			e.OfferRegionCode(addr, hash, code)
		}
		offered++
	}
	offeredNS := time.Since(started).Nanoseconds()
	var step vm.ValuePreparationStep
	for _, p := range preparers {
		q := p.AdvanceValuePreparation()
		step.Nanoseconds += q.Nanoseconds
		step.Analyses += q.Analyses
		step.CompileBatches += q.CompileBatches
		step.CreditNS += q.CreditNS
	}
	return map[string]any{"block": block.NumberU64(), "codeOffers": offered, "offerNS": offeredNS, "step": step, "totalNS": time.Since(started).Nanoseconds(), "beforeStatePrefetch": true}
}
