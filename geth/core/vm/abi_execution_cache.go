package vm

import (
	"github.com/ethereum/go-ethereum/crypto"
	"time"
)

type abiExecutionKey struct {
	arithmetic arithmeticCacheKey
	engine     *ArithmeticEngine
}
type abiExecutionPlan struct {
	tradeAllowed *tradeSelectorSet
	arithmetic   *boundArithmetic
	abi          *abiScalarBound
	// Only an early snapshot carries an upgrade pointer. A complete plan has
	// no mutable preparation state; active frames retain their chosen snapshot.
	preparing *backgroundPlan
}
type abiExecutionSlot struct {
	key   abiExecutionKey
	value *abiExecutionPlan
}

// One code-level cache hit supplies both region kinds, including negative
// results. This engine (also its scratch) is sequential, not a
// concurrently shared executor. Engine identity isolates template versions.
func (evm *EVM) abiExecutionBinding(contract *Contract) (*boundArithmetic, *abiScalarBound) {
	ac := evm.Config.ABIScalar
	if ac == nil || ac.Engine == nil || evm.chainRules.IsEIP4762 || evm.Config.Tracer != nil || len(evm.Config.ExtraEips) > 0 {
		return evm.arithmeticBinding(contract), nil
	}
	e := ac.Engine
	if !e.Optimize {
		a := evm.arithmeticBinding(contract)
		view := contract.Code
		if a != nil {
			view = a.dispatchCode
		}
		return a, e.bind(contract, evm.table, a, view)
	}
	cfg, key, ok := evm.abiArithmeticKey(contract)
	if !ok || (cfg.Mode != ArithmeticSync) {
		a := evm.arithmeticBinding(contract)
		view := contract.Code
		if a != nil {
			view = a.dispatchCode
		}
		return a, e.bind(contract, evm.table, a, view)
	}
	k := abiExecutionKey{arithmetic: key, engine: cfg.Engine}
	slot := &e.direct[key.hash[0]&63]
	if slot.value != nil && slot.key == k {
		cfg.Engine.stats.hits.Add(1)
		return slot.value.arithmetic, slot.value.abi
	}
	if p := e.execution[k]; p != nil {
		*slot = abiExecutionSlot{k, p}
		cfg.Engine.stats.hits.Add(1)
		return p.arithmetic, p.abi
	}
	if stats := evm.Config.RegionExecutionStats; stats != nil {
		started := time.Now()
		defer func() { stats.BindingNS += time.Since(started).Nanoseconds(); stats.Bindings++ }()
	}
	a := evm.arithmeticBinding(contract)
	view := contract.Code
	if a != nil {
		view = a.dispatchCode
	}
	p := &abiExecutionPlan{arithmetic: a, abi: e.bind(contract, evm.table, a, view, false)}

	{
		p = evm.bindComposedRegions(p, contract.Code)
	}
	if e.execution == nil {
		e.execution = make(map[abiExecutionKey]*abiExecutionPlan)
	}
	e.execution[k] = p
	*slot = abiExecutionSlot{k, p}
	return p.arithmetic, p.abi
}

// Same policy identity as arithmeticBinding. Shared by this cache
// so ordinary arithmetic-only builds retain their existing hot path.
func (evm *EVM) abiArithmeticKey(contract *Contract) (*ArithmeticConfig, arithmeticCacheKey, bool) {
	if evm.Config.Arithmetic == nil || evm.Config.Arithmetic.Engine == nil {
		return nil, arithmeticCacheKey{}, false
	}
	cfg := evm.Config.Arithmetic
	if cfg.Mode > ArithmeticSync || cfg.GasPolicy != OriginalArithmeticGas {
		return nil, arithmeticCacheKey{}, false
	}
	// EIP-4762 code-chunk charging and opcode-level tracing require original steps.
	if evm.chainRules.IsEIP4762 || evm.Config.Tracer != nil || len(evm.Config.ExtraEips) > 0 {
		return nil, arithmeticCacheKey{}, false
	}
	codeHash := contract.CodeHash
	if !contract.arithmeticHashTrusted {
		codeHash = crypto.Keccak256Hash(contract.Code)
	}
	key := arithmeticCacheKey{hash: codeHash, rules: evm.table, helperOnly: false, reduced: true}
	key.ordinary = evm.ordinaryLibrary(cfg)

	key.compute = cfg.computeHelperPolicy()

	return cfg, key, true
}
