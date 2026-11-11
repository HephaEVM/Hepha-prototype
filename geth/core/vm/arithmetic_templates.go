package vm

import (
	"encoding/json"
	"fmt"
	"io"
	"sort"
	"sync"
	"sync/atomic"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/holiman/uint256"
	miner "go-eth1.16.5-evm/template_miner"
	"golang.org/x/sync/singleflight"
)

type ArithmeticMode uint8

const (
	ArithmeticObserve ArithmeticMode = iota // V1: bind/cache, execute original instructions
	ArithmeticSync                          // V2
)

type ArithmeticGasPolicy uint8

const (
	OriginalArithmeticGas ArithmeticGasPolicy = iota
)

// ArithmeticConfig is copied into each frame. Do not change it during execution.
// A nil Config.Arithmetic selects the unmodified instruction interpreter (B).
type ArithmeticConfig struct {
	Engine               *ArithmeticEngine
	Mode                 ArithmeticMode
	GasPolicy            ArithmeticGasPolicy
	HashPreparation      HashPreparationMode
	HashPreparationStats *HashPreparationStats
	ComputeStats         *ComputeHelperStats
}

type ArithmeticEngineOptions struct{ CacheEntries, FeedbackEntries int }
type ArithmeticStats struct {
	CacheHits, CacheMisses, MatchedRegions, RejectedRegions           uint64
	Pruned, Extended, Checked, Failures, Fallbacks, ModelMismatches   uint64
	FrameContexts, NativeChecks, FixedStackPrunes, ConstantPairPrunes uint64
}
type arithmeticCounters struct {
	hits, misses, matched, rejected                            atomic.Uint64
	pruned, extended, checked, failures, fallbacks, mismatches atomic.Uint64
	frames, native, fixed, constantPairs                       atomic.Uint64
}
type arithmeticCacheKey struct {
	ordinary *HotRegionRuntime

	hash       common.Hash
	rules      *JumpTable
	helperOnly bool // separate immutable dispatch view; never tag disabled extended entries
	reduced    bool // identifies predicate lowering in the immutable binding

	compute uint8
}
type boundArithmetic struct {
	// Immutable opcode boundaries; shared with fresh execution frames.
	bitmap                              BitVec
	observationOwned, observationNative bool
	observationRemaining                int
	compactEntries                      []compactDemandEntry
	nativeOnly                          bool

	regions map[uint64]*arithmeticRegion
	// Zero denotes an ordinary PC. Lookup is two bytes per PC and never hashes
	// the PC in the interpreter's ordinary-instruction path.
	dispatchCode   []byte // private opcode-fetch view; original code remains unchanged
	entryIndex     []uint16
	entryRegions   []*arithmeticRegion
	err            error
	computeEntries []*computeHelperRegion // aligned with entryRegions; no per-opcode lookup
}
type arithmeticCacheEntry struct {
	key   arithmeticCacheKey
	value *boundArithmetic
}

// ArithmeticEngine owns one immutable template-library version. Sharing this
// object across EVMs shares positive AND negative code-hash/rules bindings.
type ArithmeticEngine struct {
	preparationYield func() // optional, private background analyzer only
	matcher          *miner.BytecodeMatcher
	version          string
	templates        []miner.Template
	mu               sync.Mutex
	cache            map[arithmeticCacheKey]*boundArithmetic
	fifo             []arithmeticCacheKey
	capacity         int
	flights          singleflight.Group
	feedback         *arithmeticFeedback
	stats            arithmeticCounters
}

func LoadArithmeticEngine(r io.Reader, options ArithmeticEngineOptions) (*ArithmeticEngine, error) {
	var bundle struct {
		SchemaVersion int              `json:"schemaVersion"`
		Templates     []miner.Template `json:"templates"`
	}
	if err := json.NewDecoder(io.LimitReader(r, 64<<20)).Decode(&bundle); err != nil {
		return nil, err
	}
	if bundle.SchemaVersion != 1 {
		return nil, fmt.Errorf("unsupported template bundle schema %d", bundle.SchemaVersion)
	}
	return NewArithmeticEngine(bundle.Templates, options)
}

func NewArithmeticEngine(templates []miner.Template, options ArithmeticEngineOptions) (*ArithmeticEngine, error) {
	// Deep copy: callers cannot alter published templates or expression trees.
	data, err := json.Marshal(templates)
	if err != nil {
		return nil, err
	}
	var frozen []miner.Template
	if err = json.Unmarshal(data, &frozen); err != nil {
		return nil, err
	}
	matcher, err := miner.NewBytecodeMatcher(frozen)
	if err != nil {
		return nil, err
	}
	sort.Slice(frozen, func(i, j int) bool { return frozen[i].ID < frozen[j].ID })
	data, _ = json.Marshal(frozen)
	if options.CacheEntries <= 0 {
		options.CacheEntries = 512
	}
	if options.FeedbackEntries <= 0 {
		options.FeedbackEntries = 128
	}
	return &ArithmeticEngine{matcher: matcher, version: crypto.Keccak256Hash(data).Hex(), templates: frozen, cache: make(map[arithmeticCacheKey]*boundArithmetic), capacity: options.CacheEntries, feedback: newArithmeticFeedback(options.FeedbackEntries)}, nil
}
func (e *ArithmeticEngine) Version() string { return e.version }
func (e *ArithmeticEngine) Stats() ArithmeticStats {
	s := &e.stats
	return ArithmeticStats{CacheHits: s.hits.Load(), CacheMisses: s.misses.Load(), MatchedRegions: s.matched.Load(), RejectedRegions: s.rejected.Load(), Pruned: s.pruned.Load(), Extended: s.extended.Load(), Checked: s.checked.Load(), Failures: s.failures.Load(), Fallbacks: s.fallbacks.Load(), ModelMismatches: s.mismatches.Load(), FrameContexts: s.frames.Load(), NativeChecks: s.native.Load(), FixedStackPrunes: s.fixed.Load(), ConstantPairPrunes: s.constantPairs.Load()}
}

func (e *ArithmeticEngine) bind(code []byte, key arithmeticCacheKey) *boundArithmetic {
	e.mu.Lock()
	cached := e.cache[key]
	e.mu.Unlock()
	if cached != nil {
		e.stats.hits.Add(1)
		return cached
	}
	result, _, _ := e.flights.Do(fmt.Sprintf("%p/%s/%p/%t/%t/%d", key.ordinary, key.hash, key.rules, key.helperOnly, key.reduced, key.compute), func() (any, error) {
		e.mu.Lock()
		cached := e.cache[key]
		e.mu.Unlock()
		if cached != nil {
			return cached, nil
		}
		e.stats.misses.Add(1)
		bound := &boundArithmetic{regions: map[uint64]*arithmeticRegion{}, bitmap: codeBitmap(code)}
		report, err := e.matcher.MatchWithYield(code, e.preparationYield)
		if paperAblationLevel == 1 {
			report.Bindings = nil
		}
		bound.err = err
		if err == nil {
			ownedCode := append([]byte(nil), code...)
			for _, b := range report.Bindings {
				r, err := prepareArithmeticRegion(b, code, key.rules)
				if err != nil {
					e.stats.rejected.Add(1)
					report.Candidates = append(report.Candidates, miner.CandidateHint{Entry: b.Entry, TemplateID: b.TemplateID, Reason: err.Error()})
					continue
				}
				if key.reduced && r.plan.nativeReduced != nil {
					r.plan.nativeHot = r.plan.nativeReduced
					r.plan.guardHot = r.plan.guardReduced
				}

				bound.regions[r.entry] = r
				r.code = ownedCode
				e.stats.matched.Add(1)
			}
			if !key.helperOnly {
				extendArithmeticEntries(bound, code, key.rules)
			}
			bound.indexEntries(code)

			// Walk actual opcode boundaries, never PUSH immediate bytes. Reuse
			// the entry-tag dispatch branch instead of polling each opcode.

			if key.compute != 0 && paperAblationLevel != 1 {
				bound.prepareComputeHelpers(code, key)
			}
			if key.ordinary != nil {
				key.ordinary.bind(bound, code, key)
			}
			if paperAblationLevel == 1 && bound.dispatchCode == nil {
				bound.dispatchCode = append([]byte(nil), code...)
				bound.entryIndex = make([]uint16, len(code))
				bound.entryRegions = make([]*arithmeticRegion, 1)
				bound.computeEntries = make([]*computeHelperRegion, 1)
			}
			e.feedback.record(e.version, code, report.Candidates)
		}
		e.mu.Lock()
		if len(e.fifo) == e.capacity {
			delete(e.cache, e.fifo[0])
			e.fifo = e.fifo[1:]
		}
		e.fifo = append(e.fifo, key)
		e.cache[key] = bound
		e.mu.Unlock()
		return bound, nil
	})
	return result.(*boundArithmetic)
}

// Reserved only in our private dispatch view, never in contract bytecode.
const arithmeticEntryMarker OpCode = 0x0c

func (b *boundArithmetic) indexEntries(code []byte) {
	codeSize := len(code)
	if len(b.regions) == 0 || len(b.regions) >= 1<<16 {
		return
	}
	b.dispatchCode = append([]byte(nil), code...)
	b.entryIndex = make([]uint16, codeSize)
	b.entryRegions = make([]*arithmeticRegion, 1, len(b.regions)+1)
	for pc, region := range b.regions {
		if pc >= uint64(codeSize) {
			continue
		}
		region.prepareSyncLayout()
		region.prepareSyncHot()
		b.dispatchCode[pc] = byte(arithmeticEntryMarker)
		b.entryIndex[pc] = uint16(len(b.entryRegions))
		b.entryRegions = append(b.entryRegions, region)
	}
}

// Result expressions are compiled once to native word operations. Unlike finite
// type inference, this preserves the exact returned expression for ALL inputs.
type arithmeticExpression func(a, b uint256.Int) uint256.Int

func compileArithmeticResult(e *miner.Expr) (arithmeticExpression, string, error) {
	return compileArithmeticResultPolicy(e, false)
}

func compileArithmeticResultPolicy(e *miner.Expr, fast bool) (arithmeticExpression, string, error) {
	if e == nil {
		return nil, "", fmt.Errorf("missing result")
	}
	if e.Op == "input" {
		switch e.Input {
		case "a":
			return func(a, b uint256.Int) uint256.Int { return a }, "a", nil
		case "b":
			return func(a, b uint256.Int) uint256.Int { return b }, "b", nil
		}
		return nil, "", fmt.Errorf("result depends on %s", e.Input)
	}
	if e.Op == "const" {
		v, err := uint256.FromHex(e.Value)
		if err != nil {
			return nil, "", err
		}
		word := *v
		return func(a, b uint256.Int) uint256.Int { return word }, word.Hex(), nil
	}
	if (e.Op == "ISZERO" || e.Op == "NOT") && len(e.Args) == 1 {
		f, key, err := compileArithmeticResultPolicy(e.Args[0], fast)
		if err != nil {
			return nil, "", err
		}
		return func(a, b uint256.Int) uint256.Int {
			x := f(a, b)
			var z uint256.Int
			if e.Op == "NOT" {
				z.Not(&x)
			} else if x.IsZero() {
				z.SetOne()
			}
			return z
		}, e.Op + "(" + key + ")", nil
	}
	if len(e.Args) != 2 {
		return nil, "", fmt.Errorf("unsupported result expression")
	}
	l, lk, err := compileArithmeticResultPolicy(e.Args[0], fast)
	if err != nil {
		return nil, "", err
	}
	r, rk, err := compileArithmeticResultPolicy(e.Args[1], fast)
	if err != nil {
		return nil, "", err
	}
	switch e.Op {
	case "ADD", "SUB", "MUL", "DIV", "MOD", "SDIV", "SMOD", "AND", "SIGNEXTEND", "OR", "XOR", "LT", "GT", "SLT", "SGT", "EQ", "SHL", "SHR", "SAR", "BYTE", "EXP":
	default:
		return nil, "", fmt.Errorf("unsupported result operation %s", e.Op)
	}
	if fast {
		if fn := arithmeticDirectResults[e.Op+"("+lk+","+rk+")"]; fn != nil {
			return fn, e.Op + "(" + lk + "," + rk + ")", nil
		}
	}
	return func(a, b uint256.Int) uint256.Int {
		x, y := l(a, b), r(a, b)
		var z uint256.Int
		switch e.Op {
		case "ADD":
			z.Add(&x, &y)
		case "SUB":
			z.Sub(&x, &y)
		case "MUL":
			z.Mul(&x, &y)
		case "DIV":
			z.Div(&x, &y)
		case "MOD":
			z.Mod(&x, &y)
		case "SDIV":
			z.SDiv(&x, &y)
		case "SMOD":
			z.SMod(&x, &y)
		case "AND":
			z.And(&x, &y)
		case "SIGNEXTEND":
			z.ExtendSign(&y, &x)
		case "OR":
			z.Or(&x, &y)
		case "XOR":
			z.Xor(&x, &y)
		case "BYTE":
			z.Set(&y)
			z.Byte(&x)
		case "EXP":
			z.Exp(&x, &y)
		case "LT":
			if x.Lt(&y) {
				z.SetOne()
			}
		case "GT":
			if x.Gt(&y) {
				z.SetOne()
			}
		case "SLT":
			if x.Slt(&y) {
				z.SetOne()
			}
		case "SGT":
			if x.Sgt(&y) {
				z.SetOne()
			}
		case "EQ":
			if x.Eq(&y) {
				z.SetOne()
			}
		case "SHL", "SHR", "SAR":
			shift := uint(256)
			if x.IsUint64() && x.Uint64() < 256 {
				shift = uint(x.Uint64())
			}
			if e.Op == "SHL" {
				z.Lsh(&y, shift)
			} else if e.Op == "SHR" {
				z.Rsh(&y, shift)
			} else {
				z.SRsh(&y, shift)
			}
		}
		return z
	}, e.Op + "(" + lk + "," + rk + ")", nil
}

type arithmeticPath struct {
	originalFailure bool
	path            miner.SymbolicPath
	gas             uint64
}
type stackSource struct {
	index    int
	constant uint256.Int
} // index from bottom of consumed input; -1 = constant
func (s stackSource) get(input []uint256.Int) uint256.Int {
	if s.index < 0 {
		return s.constant
	}
	return input[s.index]
}

type arithmeticRegion struct {
	syncHot              arithmeticSyncHot
	syncLayout           arithmeticSyncLayout
	plan                 *arithmeticPlan
	legacy               bool
	pathDependentResult  bool
	code                 []byte // immutable copy, shared by regions; retained for rare counterexamples
	entry, helper        uint64
	template             string
	op                   string
	domain               IntegerDomain
	inputCount, peak     int
	a, b, k              stackSource
	lhsIsA               bool
	prefix               []stackSource
	prefixIdentity       bool
	constantContinuation bool
	result               arithmeticExpression
	resultFast           arithmeticExpression
	directMultiply       bool // canonical result expression is exactly MUL(a,b) or MUL(b,a)
	paths                []arithmeticPath
	successGas           uint64
	constantSuccessGas   bool
	extraGas             uint64
}

func prepareArithmeticRegion(b miner.RuntimeBinding, code []byte, table *JumpTable) (*arithmeticRegion, error) {
	legacy := b.Semantics == "safemath_uint256" || b.Semantics == "safemath_int256"
	if (!legacy && b.Semantics != "") || (len(b.InputStack) != 3 && !(legacy && len(b.InputStack) == 4)) || b.Analysis == nil {
		return nil, fmt.Errorf("runtime_requires_closed_panic_or_safemath_helper")
	}
	d, err := ParseIntegerDomain(b.IntegerType)
	if err != nil {
		return nil, err
	}
	r := &arithmeticRegion{entry: b.Entry, helper: b.Entry, template: b.TemplateID, op: b.Operation, domain: d, legacy: legacy, inputCount: len(b.InputStack), prefixIdentity: true, constantSuccessGas: true, lhsIsA: b.Analysis.OperandOrder[0] == "a"}
	for i, name := range b.InputStack {
		switch name {
		case "a":
			r.a = stackSource{index: i}
		case "b":
			r.b = stackSource{index: i}
		case "continuation":
			r.k = stackSource{index: i}
		case "errorMessage":
			if !legacy {
				return nil, fmt.Errorf("unexpected error input")
			}
		default:
			return nil, fmt.Errorf("unknown input")
		}
	}
	first := true
	resultKey := ""
	for _, path := range b.Analysis.Paths {
		p := arithmeticPath{path: path, originalFailure: legacy && path.Outcome != "return"}
		for _, pc := range path.PCs {
			if p.originalFailure {
				break
			}
			op := OpCode(code[pc])
			entry := table[op]
			if entry.undefined || (entry.dynamicGas != nil && op != MSTORE && op != REVERT) {
				return nil, fmt.Errorf("unsupported_dynamic_or_inactive_opcode_%s", op)
			}
			p.gas += entry.constantGas
		}
		if path.Outcome == "return" {
			f, key, err := compileArithmeticResult(path.Result)
			if err != nil {
				return nil, err
			}
			if resultKey != "" && resultKey != key {
				r.pathDependentResult = true
			}
			resultKey = key
			r.result = f
			var fastKey string
			r.resultFast, fastKey, err = compileArithmeticResultPolicy(canonicalArithmeticInputs(path.Result, d), true)
			if err != nil {
				return nil, err
			}
			r.directMultiply = fastKey == "MUL(a,b)" || fastKey == "MUL(b,a)"
			if first {
				r.successGas = p.gas
				first = false
			} else if r.successGas != p.gas {
				r.constantSuccessGas = false
			}
		} else if !legacy && path.Outcome != "panic_0x11" && path.Outcome != "panic_0x12" {
			return nil, fmt.Errorf("unsupported_failure_encoder")
		}
		r.peak = max(r.peak, path.PeakStack)
		r.paths = append(r.paths, p)
	}
	if first {
		return nil, fmt.Errorf("no success path")
	}
	r.plan, err = buildArithmeticPlan(r.paths)
	if err != nil {
		return nil, err
	}
	return r, nil
}

// Extend through a bounded straight-line stack-only preamble. No JUMPDEST may
// occur inside it, so an EVM jump cannot enter its middle. Slots are resolved
// statically, including preserved caller values; runtime provenance is unnecessary.
func extendArithmeticEntries(bound *boundArithmetic, code []byte, table *JumpTable) {
	instructions, err := miner.Disassemble(code)
	if err != nil {
		return
	}
	jumpdests := make(map[uint64]bool)
	for _, ins := range instructions {
		if OpCode(ins.Op) == JUMPDEST {
			jumpdests[ins.PC] = true
		}
	}
	for end, ins := range instructions {
		if ins.Op.String() != "JUMP" || end == 0 {
			continue
		}
		push := instructions[end-1]
		if len(push.Data) == 0 || len(push.Data) > 8 {
			continue
		}
		dest := new(uint256.Int).SetBytes(push.Data).Uint64()
		base := bound.regions[dest]
		if base == nil || base.entry != base.helper {
			continue
		}
		start := end - 1
		for start > 0 && end-start < 16 {
			op := OpCode(instructions[start-1].Op)
			if !(op >= PUSH0 && op <= PUSH32 || op >= DUP1 && op <= DUP16 || op >= SWAP1 && op <= SWAP16 || op == POP) {
				break
			}
			start--
		}
		// Model 32 incoming slots, then trim the untouched prefix.
		stack := make([]stackSource, 32)
		for i := range stack {
			stack[i].index = i
		}
		minRead, peak := 32, 32
		gas := uint64(0)
		valid := true
		for _, i := range instructions[start : end+1] {
			op := OpCode(i.Op)
			operation := table[op]
			if operation.undefined {
				valid = false
				break
			}
			gas += operation.constantGas
			n := len(stack)
			switch {
			case op >= PUSH0 && op <= PUSH32:
				var w uint256.Int
				w.SetBytes(i.Data)
				stack = append(stack, stackSource{index: -1, constant: w})
			case op >= DUP1 && op <= DUP16:
				d := int(op-DUP1) + 1
				if n < d {
					valid = false
					break
				}
				s := stack[n-d]
				if s.index >= 0 {
					minRead = min(minRead, s.index)
				}
				stack = append(stack, s)
			case op >= SWAP1 && op <= SWAP16:
				d := int(op-SWAP1) + 1
				if n <= d {
					valid = false
					break
				}
				for _, s := range []stackSource{stack[n-1], stack[n-1-d]} {
					if s.index >= 0 {
						minRead = min(minRead, s.index)
					}
				}
				stack[n-1], stack[n-1-d] = stack[n-1-d], stack[n-1]
			case op == POP || op == JUMP:
				if n < 1 {
					valid = false
					break
				}
				if stack[n-1].index >= 0 {
					minRead = min(minRead, stack[n-1].index)
				}
				stack = stack[:n-1]
			default:
				valid = false
			}
			peak = max(peak, len(stack))
			if !valid {
				break
			}
		}
		if !valid || len(stack) < base.inputCount {
			continue
		}
		for _, s := range stack[len(stack)-base.inputCount:] {
			if s.index >= 0 {
				minRead = min(minRead, s.index)
			}
		}
		// Only identity slots may be omitted from the preserved output prefix.
		for i := 0; i < minRead; i++ {
			if i >= len(stack)-base.inputCount || stack[i].index != i {
				minRead = i
				break
			}
		}
		adjust := func(s stackSource) stackSource {
			if s.index >= 0 {
				s.index -= minRead
			}
			return s
		}
		r := *base
		r.entry = instructions[start].PC
		r.inputCount = 32 - minRead
		r.extraGas = gas
		if r.inputCount > 32 {
			continue
		}
		args := stack[len(stack)-base.inputCount:]
		r.a = adjust(args[base.a.index])
		r.b = adjust(args[base.b.index])
		r.k = adjust(args[base.k.index])
		r.constantContinuation = r.k.index < 0 && r.k.constant.IsUint64() && jumpdests[r.k.constant.Uint64()]
		r.prefix = nil
		for _, s := range stack[minRead : len(stack)-base.inputCount] {
			r.prefix = append(r.prefix, adjust(s))
		}
		r.prefixIdentity = true
		for i, s := range r.prefix {
			if s.index != i {
				r.prefixIdentity = false
				break
			}
		}
		r.peak = max(peak-minRead, len(stack)-base.inputCount-minRead+base.peak)
		if _, exists := bound.regions[r.entry]; !exists {
			bound.regions[r.entry] = &r
		}
	}
}

func arithmeticBarrier(op OpCode) bool {
	switch op {
	case CALL, CALLCODE, DELEGATECALL, STATICCALL, CREATE, CREATE2, SELFDESTRUCT:
		return true
	}
	return false
}
