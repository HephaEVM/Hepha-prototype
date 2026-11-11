package vm

import "time"

// PruningConfig is the supported runtime interface. Plans and counters are
// resources, not strategy switches. Execution is synchronous and retains
// original gas and all admission/exception checks.
// An ABI engine (including its scratch space) belongs to one sequential runner.
type PruningConfig struct {
	// Optional cold-start preparation; nil preserves synchronous binding.
	Preparation  *BackgroundPreparation
	Engine       *ArithmeticEngine
	ABI          *ABIScalarDiagnostic
	ComputeStats *ComputeHelperStats
	HashStats    *HashPreparationStats
}

// WithPruning resolves the selected policy once, outside the instruction loop.
// Historical ArithmeticConfig ablation fields are intentionally not exposed.
func (c Config) WithPruning(p PruningConfig) Config {
	c.Pruning = &p
	c.Arithmetic = &ArithmeticConfig{
		Engine: p.Engine, Mode: ArithmeticSync, GasPolicy: OriginalArithmeticGas,

		HashPreparation: HashPreparationDirect,
		ComputeStats:    p.ComputeStats, HashPreparationStats: p.HashStats,
	}
	c.ABIScalar = nil
	if p.ABI != nil {
		c.ABIScalar = &ABIScalarConfig{Engine: p.ABI, ReuseBindingPolicy: true}
	}

	return c
}

func (evm *EVM) runSelected(contract *Contract, input []byte, readOnly bool) ([]byte, error, time.Duration) {
	if p := evm.Config.Pruning; p != nil && p.Preparation != nil && evm.depth == 0 {
		p.Preparation.EnterForeground()
		defer p.Preparation.LeaveForeground()
	}

	c := evm.Config.Arithmetic
	if c == nil && evm.Config.ABIScalar == nil {
		if cache := evm.Config.StreamRegions; cache != nil && evm.Config.Tracer == nil && !evm.chainRules.IsEIP4762 && len(evm.Config.ExtraEips) == 0 {
			contract.useStreamCache(cache, &evm.Config)
			return evm.runUnboundStack(contract, input, readOnly)
		}
		return evm.RunNative(contract, input, readOnly)
	}
	if evm.Config.Tracer != nil || evm.chainRules.IsEIP4762 || len(evm.Config.ExtraEips) != 0 || c != nil && (c.Mode != ArithmeticSync || c.GasPolicy != OriginalArithmeticGas) {
		return evm.RunNative(contract, input, readOnly)
	}
	if p := evm.Config.Pruning; p != nil && p.Preparation != nil {
		tier := evm.Config.PreparationTier
		var plan *abiExecutionPlan
		if p.Preparation.frozenReadyOnly {
			plan = p.Preparation.lookupReady(evm, contract)
		} else {
			if hotspotEnabled {
				tier = valueFrameTier(evm, contract)
			}
			if hotspotEnabled && tier == 0 {
				if valuePolicy {
					plan = p.Preparation.lookupReady(evm, contract)
				}
			} else {
				plan = p.Preparation.lookup(evm, contract, tier)
			}
		}
		if plan == nil || !plan.tradeAllowed.accept(input) || ((plan.arithmetic == nil && plan.abi == nil) || (plan.arithmetic != nil && plan.arithmetic.nativeOnly)) && !p.Preparation.frozenLeanFallback {
			if plan != nil && plan.arithmetic != nil {
				contract.analysis = plan.arithmetic.bitmap
			}
			if evm.Config.StreamRegions != nil {
				contract.useStreamCache(evm.Config.StreamRegions, &evm.Config)
			}
			return evm.runUnboundStack(contract, input, readOnly)
		}
		return evm.runPreparedRegions(contract, input, readOnly, plan)
	}
	if paperAblationLevel == 3 {
		evm.abiPolicyExecutionBinding(contract)
		return evm.RunNative(contract, input, readOnly)
	}
	return evm.runPrunedRegions(contract, input, readOnly)
}
