package main

import "go-eth1.16.5-evm/core/vm"

func executionReportingConfig(cfg vm.Config, enabled bool) vm.Config {
	if enabled {
		return cfg
	}
	cfg.RegionExecutionStats = nil
	// NewEVM calls WithPruning again, so update the authoritative source too.
	// Copy both configurations to avoid modifying another arm's shared settings.
	if cfg.Pruning != nil {
		p := *cfg.Pruning
		p.ComputeStats = nil
		p.HashStats = nil
		cfg.Pruning = &p
	}
	if cfg.Arithmetic != nil {
		a := *cfg.Arithmetic
		a.ComputeStats = nil
		a.HashPreparationStats = nil
		cfg.Arithmetic = &a
	}
	return cfg
}
