package main

import (
	"go-eth1.16.5-evm/config"
	"go-eth1.16.5-evm/core"

	"go-eth1.16.5-evm/ethdb"
)

// Use the same consensus engine and HeaderChain setup as experiments.newProcessor.
func newReplayProcessor(db ethdb.Database) (*core.StateProcessor, error) {
	engine, err := config.CreateConsensusEngine(config.MainnetChainConfig)
	if err != nil {
		return nil, err
	}
	chain, err := core.NewHeaderChain(db, config.MainnetChainConfig, engine)
	if err != nil {
		return nil, err
	}
	return core.NewStateProcessor(chain), nil
}
