package vm

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

// The namespace fixes engine identities, rules and the complete binding policy.
// Only code identity changes within it. Positive and negative plans are equally
// cacheable; a previous path without hits is never evidence for a negative plan.
type abiPolicySlot struct {
	hash common.Hash
	plan *abiExecutionPlan
}
type abiPolicyCache struct {
	direct [64]abiPolicySlot
	plans  map[common.Hash]*abiExecutionPlan
}
type abiRunBinding struct {
	key    abiExecutionKey // hash is zero: code-independent policy identity
	engine *ABIScalarDiagnostic
	config *ArithmeticConfig
	cache  *abiPolicyCache
}

func (evm *EVM) abiPolicyExecutionBinding(contract *Contract) (*boundArithmetic, *abiScalarBound) {
	ac := evm.Config.ABIScalar
	if ac != nil && ac.ReuseBindingPolicy {
		return evm.abiPreparedBinding(contract)
	}
	return evm.abiExecutionBinding(contract)
}

func (evm *EVM) abiPreparedBinding(contract *Contract) (*boundArithmetic, *abiScalarBound) {
	ac := evm.Config.ABIScalar
	session := &evm.abiBinding
	hash := contract.CodeHash
	// Config is immutable during execution, as required by ArithmeticConfig.
	// Depth zero callers and each root invocation revalidate all policy fields.
	if evm.depth <= 1 || session.cache == nil || session.config != evm.Config.Arithmetic || session.engine != ac.Engine {
		var ok bool
		if hash, ok = evm.prepareABIBindingPolicy(contract); !ok {
			return evm.abiExecutionBinding(contract)
		}
	} else if !contract.arithmeticHashTrusted {
		hash = crypto.Keccak256Hash(contract.Code)
	}
	cache := session.cache
	slot := &cache.direct[hash[0]&63]
	if slot.plan != nil && slot.hash == hash {
		session.config.Engine.stats.hits.Add(1)
		return slot.plan.arithmetic, slot.plan.abi
	}
	return evm.abiPolicyMiss(contract, hash, cache, slot)
}

// Cold path kept out of the nested lookup. Do not store an EVM/state pointer
// in the shared engine; session lifetime follows the EVM itself.
func (evm *EVM) prepareABIBindingPolicy(contract *Contract) (common.Hash, bool) {
	ac := evm.Config.ABIScalar
	if ac.Engine == nil || !ac.Engine.Optimize {
		evm.abiBinding = abiRunBinding{}
		return common.Hash{}, false
	}
	cfg, key, ok := evm.abiArithmeticKey(contract)
	if !ok || (cfg.Mode != ArithmeticSync) {
		evm.abiBinding = abiRunBinding{}
		return common.Hash{}, false
	}
	hash := key.hash
	key.hash = common.Hash{}
	policy := abiExecutionKey{arithmetic: key, engine: cfg.Engine}
	old := &evm.abiBinding
	if old.engine == ac.Engine && old.key == policy && old.cache != nil {
		old.config = cfg
		return hash, true
	}
	e := ac.Engine
	cache := e.policies[policy]
	if cache == nil {
		cache = &abiPolicyCache{plans: make(map[common.Hash]*abiExecutionPlan)}
		if e.policies == nil {
			e.policies = make(map[abiExecutionKey]*abiPolicyCache)
		}
		e.policies[policy] = cache
	}
	evm.abiBinding = abiRunBinding{policy, e, cfg, cache}
	return hash, true
}

func (evm *EVM) abiPolicyMiss(contract *Contract, hash common.Hash, cache *abiPolicyCache, slot *abiPolicySlot) (*boundArithmetic, *abiScalarBound) {
	p := cache.plans[hash]
	if p == nil {
		a, b := evm.abiExecutionBinding(contract)
		// Share the immutable plan already inserted by the reference joint cache.
		key := evm.abiBinding.key
		key.arithmetic.hash = hash
		p = evm.Config.ABIScalar.Engine.execution[key]
		if p == nil {
			return a, b
		}
		cache.plans[hash] = p
	} else {
		evm.abiBinding.config.Engine.stats.hits.Add(1)
	}
	*slot = abiPolicySlot{hash, p}
	return p.arithmetic, p.abi
}
