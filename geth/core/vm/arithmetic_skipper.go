package vm

import (
	"sync"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/holiman/uint256"
)

type arithmeticTask struct {
	seq              uint64
	region           *arithmeticRegion
	a, b             uint256.Int // copied words; never pointers into the live stack
	gas, memoryBytes uint64
}
type arithmeticResult struct {
	result    uint256.Int
	fallback  bool
	seq       uint64
	err       error
	data      []byte
	gas, cost uint64
}
type arithmeticFrame struct {
	scratch                                             *arithmeticScratch
	config                                              ArithmeticConfig
	bound                                               *boundArithmetic
	first                                               *arithmeticResult
	sequence                                            uint64
	pruned, extended, fallbacks, checked, checkFailures uint64
	nativeChecked, fixedPrunes, constantPairs           uint64
	finished                                            bool
}

var arithmeticFrames = sync.Pool{New: func() any { return new(arithmeticFrame) }}

func releaseArithmeticFrame(f *arithmeticFrame) {
	// finish has joined all work; no worker retains this frame or its operands.
	scratch := f.scratch
	*f = arithmeticFrame{scratch: scratch}
	arithmeticFrames.Put(f)
}

func (evm *EVM) arithmeticBinding(contract *Contract) *boundArithmetic {
	if evm.Config.Arithmetic == nil || evm.Config.Arithmetic.Engine == nil {
		return nil
	}
	cfg := evm.Config.Arithmetic
	if cfg.Mode > ArithmeticSync || cfg.GasPolicy != OriginalArithmeticGas {
		return nil
	}
	// EIP-4762 code-chunk charging and opcode-level tracing require original steps.
	if evm.chainRules.IsEIP4762 || evm.Config.Tracer != nil || len(evm.Config.ExtraEips) > 0 {
		return nil
	}
	codeHash := contract.CodeHash
	if !contract.arithmeticHashTrusted {
		codeHash = crypto.Keccak256Hash(contract.Code)
	}
	key := arithmeticCacheKey{hash: codeHash, rules: evm.table, helperOnly: false, reduced: true}
	key.ordinary = evm.ordinaryLibrary(cfg)

	key.compute = cfg.computeHelperPolicy()

	var bound *boundArithmetic
	{
		// A full key comparison protects against direct-map collisions, rule or
		// engine changes, and helper-only bindings. Negative results are cached.
		slot := &evm.arithmeticDirect[codeHash[0]&63]
		if slot.engine == cfg.Engine && slot.key == key {
			bound = slot.value
			cfg.Engine.stats.hits.Add(1)
		} else {
			bound = evm.arithmeticLookup(contract, key)
			if bound == nil {
				return nil
			}
			*slot = arithmeticTinyEntry{cfg.Engine, key, bound}
		}
		if cfg.Mode == ArithmeticObserve || bound.err != nil || (len(bound.entryRegions) <= 1 && paperAblationLevel != 1) {
			return nil
		}
		return bound
	}
}

func newPreparedArithmeticFrame(cfg ArithmeticConfig, bound *boundArithmetic) *arithmeticFrame {
	f := arithmeticFrames.Get().(*arithmeticFrame)
	f.config, f.bound = cfg, bound
	cfg.Engine.stats.frames.Add(1)
	return f
}

// Kept for direct tests; interpreter creates a context only at the first entry.
func (evm *EVM) newArithmeticFrame(contract *Contract) *arithmeticFrame {
	bound := evm.arithmeticBinding(contract)
	if bound == nil {
		return nil
	}
	return newPreparedArithmeticFrame(*evm.Config.Arithmetic, bound)
}

type arithmeticTinyEntry struct {
	engine *ArithmeticEngine
	key    arithmeticCacheKey
	value  *boundArithmetic
}

func memoryGas(bytes uint64) uint64 { w := (bytes + 31) / 32; return 3*w + w*w/512 }
func (f *arithmeticFrame) evaluate(t arithmeticTask) arithmeticResult {
	r := t.region
	f.checked++
	var index int
	var nativeResult uint256.Int
	useNative := r.plan.native != nil
	if useNative {
		fn := r.plan.native

		index, nativeResult = fn(t.a, t.b)
		f.nativeChecked++
	} else {
		if f.scratch == nil {
			f.scratch = new(arithmeticScratch)
		}
		index = r.plan.selectPath(t.a, t.b, f.scratch)
	}
	return f.evaluatedPath(t, index, nativeResult, useNative)
}

// Shared exceptional path after exact native selection, without evaluating the
// same predicate again. A batch uses this only when its compact success test fails.
func (f *arithmeticFrame) evaluatedPath(t arithmeticTask, index int, nativeResult uint256.Int, useNative bool) arithmeticResult {
	r := t.region
	p := &r.paths[index]
	if p.originalFailure {
		return arithmeticResult{seq: t.seq, fallback: true}
	}
	cost := uint64(0)
	if f.config.GasPolicy == OriginalArithmeticGas {
		cost = p.gas + r.extraGas
	}
	if p.path.MemoryBytes > t.memoryBytes {
		delta := memoryGas(p.path.MemoryBytes) - memoryGas(t.memoryBytes)
		if delta > ^uint64(0)-cost {
			return arithmeticResult{seq: t.seq, err: ErrOutOfGas}
		}
		cost += delta
	}
	out := arithmeticResult{seq: t.seq, cost: cost}
	if p.path.Outcome == "return" {
		if useNative {
			out.result = nativeResult
		} else {
			out.result = *r.plan.value(r.plan.results[index], f.scratch)
		}
	}

	if p.path.Outcome != "return" { // notify before constructing the recovery payload
		f.checkFailures++
		out.err = ErrExecutionReverted
		out.data = make([]byte, 36)
		copy(out.data, []byte{0x4e, 0x48, 0x7b, 0x71})
		out.data[35] = 0x11
		if p.path.Outcome == "panic_0x12" {
			out.data[35] = 0x12
		}
	}
	if t.gas < cost {
		out.err = ErrOutOfGas
		out.data = nil
	} else {
		out.gas = t.gas - cost
	}
	return out
}

func (f *arithmeticFrame) finish() *arithmeticResult { return f.finishSynchronous() }

// try replaces only a complete prepared region. All preconditions are tested
// before touching live state; unsupported cases execute the original bytecode.
func (f *arithmeticFrame) try(pc *uint64, scope *ScopeContext, r *arithmeticRegion) (bool, *arithmeticResult) {
	return f.tryMode(pc, scope, r, f.config.Mode)
}

func (f *arithmeticFrame) trySynchronous(pc *uint64, s *ScopeContext, r *arithmeticRegion) (bool, *arithmeticResult) {
	return f.tryMode(pc, s, r, ArithmeticSync)
}

func (f *arithmeticFrame) tryMode(pc *uint64, scope *ScopeContext, r *arithmeticRegion, mode ArithmeticMode) (bool, *arithmeticResult) {
	fallback := func() (bool, *arithmeticResult) { f.fallbacks++; return false, f.first }
	stack, contract := scope.Stack, scope.Contract
	if stack.len() < r.inputCount || stack.len()-r.inputCount+r.peak > 1024 {
		return fallback()
	}
	input := stack.data[stack.len()-r.inputCount:]
	a, b, k := r.a.get(input), r.b.get(input), r.k.get(input)
	if (r.domain.Bits != 256 && (!r.domain.canonical(a) || !r.domain.canonical(b))) || (!r.constantContinuation && !contract.validJumpdest(&k)) {
		return fallback()
	}

	cost := uint64(0)
	if f.config.GasPolicy == OriginalArithmeticGas {
		cost = r.successGas + r.extraGas
	}
	// Success cost cannot establish OOG: the real failure path may cost less.
	if contract.Gas < cost {
		return fallback()
	}
	f.sequence++
	task := arithmeticTask{seq: f.sequence, region: r, a: a, b: b, gas: contract.Gas, memoryBytes: uint64(scope.Memory.Len())}
	var result uint256.Int
	{
		out := f.evaluate(task)
		if out.fallback {
			return fallback()
		}
		result = out.result
		if out.err != nil {
			f.first = new(arithmeticResult)
			*f.first = out
			f.recordPrune(r)
			return true, f.first
		}
		cost = out.cost
	}

	// Read preserved sources BEFORE shortening/overwriting the input slice.
	if r.prefixIdentity {
		stack.data = stack.data[:stack.len()-r.inputCount+len(r.prefix)]
	} else {
		restoreArithmeticPrefix(stack, input, r)
	}

	stack.push(&result)
	contract.Gas -= cost
	*pc = k.Uint64()
	f.recordPrune(r)
	return true, nil
}

func (f *arithmeticFrame) recordPrune(r *arithmeticRegion) {
	f.pruned++
	if r.entry != r.helper {
		f.extended++
	}
}

func (evm *EVM) arithmeticLookup(c *Contract, k arithmeticCacheKey) *boundArithmetic {
	return evm.Config.Arithmetic.Engine.bind(c.Code, k)
}
