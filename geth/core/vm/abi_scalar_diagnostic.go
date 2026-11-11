package vm

// ABI summaries are bound to immutable bytecode and shared by the region executor.
import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/holiman/uint256"
	miner "go-eth1.16.5-evm/template_miner"
)

const abiScalarMarker OpCode = 0x0d

type ABIScalarStats struct{ Attempts, Skipped, Instructions, Gas, Fallbacks, Misses uint64 }

func (a ABIScalarStats) Sub(b ABIScalarStats) ABIScalarStats {
	return ABIScalarStats{a.Attempts - b.Attempts, a.Skipped - b.Skipped, a.Instructions - b.Instructions, a.Gas - b.Gas, a.Fallbacks - b.Fallbacks, a.Misses - b.Misses}
}

type ABIScalarDiagnostic struct {
	preparationYield func() // optional worker scheduling, never used at execution
	// Set before binding; each engine owns an immutable template/strategy version.
	Optimize     bool
	execution    map[abiExecutionKey]*abiExecutionPlan
	policies     map[abiExecutionKey]*abiPolicyCache
	direct       [64]abiExecutionSlot
	Known        map[string]bool
	cache        map[abiScalarKey]*abiScalarBound
	Stats        ABIScalarStats
	Capture      bool
	Visits       map[string]uint64
	Regions      *miner.ABIRegionLibrary
	RegionVisits map[string]uint64
	NativeVisits map[string]uint64
	scratch      abiRegionScratch
}
type ABIScalarConfig struct {
	Engine *ABIScalarDiagnostic
	Tape   *ABIScalarTape
	// Opt-in control: prepare the joint binding policy at the outer Run and
	// reuse it in nested calls. Config must remain immutable during execution.
	ReuseBindingPolicy bool
}
type abiScalarKey struct {
	hash        common.Hash
	table       *JumpTable
	arithmetic  *boundArithmetic
	specialized bool
}
type abiScalarBound struct {
	view    []byte
	entries []*abiScalarPrepared
}
type abiScalarPrepared struct {
	full       *abiRegionPrepared
	region     *miner.ABIScalarRegion
	mask, sign uint256.Int
	cost       uint64
	kind       byte
	label      string
}

func NewABIScalarDiagnostic(known map[string]bool) *ABIScalarDiagnostic {
	return &ABIScalarDiagnostic{Known: known, cache: map[abiScalarKey]*abiScalarBound{}, Visits: map[string]uint64{}, RegionVisits: map[string]uint64{}, NativeVisits: map[string]uint64{}}
}
func LoadABIScalarDiagnostic(path string) (*ABIScalarDiagnostic, error) {
	b, e := os.ReadFile(path)
	if e != nil {
		return nil, e
	}
	var known map[string]bool
	e = json.Unmarshal(b, &known)
	if e != nil {
		return nil, e
	}
	return NewABIScalarDiagnostic(known), nil
}
func (e *ABIScalarDiagnostic) bind(c *Contract, table *JumpTable, arithmetic *boundArithmetic, view []byte, specialize ...bool) *abiScalarBound {
	if paperAblationLevel == 1 {
		return nil
	}
	h := c.CodeHash
	if !c.arithmeticHashTrusted || h == (common.Hash{}) {
		h = crypto.Keccak256Hash(c.Code)
	}
	fixed := len(specialize) > 0 && specialize[0]
	key := abiScalarKey{h, table, arithmetic, fixed}
	if b, ok := e.cache[key]; ok {
		return b
	}
	e.Stats.Misses++
	var bound *abiScalarBound
	for _, r := range miner.DiscoverABIScalars(c.Code, e.Known, e.preparationYield) {
		valid := true
		cost := uint64(0)
		for _, pc := range r.SuccessPCs {
			op := table[c.Code[pc]]
			if op == nil || op.undefined || op.dynamicGas != nil {
				valid = false
				break
			}
			cost += op.constantGas
		}
		// Arithmetic and ABI entry ownership must never collide.
		if !valid || abiEntryOwned(arithmetic, r.Entry) {
			continue
		}
		if bound == nil {
			bound = &abiScalarBound{view: append([]byte(nil), view...), entries: make([]*abiScalarPrepared, len(c.Code))}
		}
		p := &abiScalarPrepared{region: r, cost: cost, label: fmt.Sprintf("%s/%d", h.Hex(), r.Entry)}
		switch r.Kind {
		case "mask":
			p.kind = 1
			p.mask.SetFromHex(r.Mask)
		case "sign":
			p.kind = 2
			p.sign.SetUint64(r.SignByte)
		case "bool":
			p.kind = 3
		}
		bound.entries[r.Entry] = p
		bound.view[r.Entry] = byte(abiScalarMarker)
	}
	if e.Regions != nil {
		for _, r := range miner.DiscoverABIRegions(c.Code, e.Regions, e.preparationYield) {
			if abiEntryOwned(arithmetic, r.Entry) {
				continue
			}
			full, err := prepareABIRegion(r, c.Code, table)
			if err != nil {
				continue
			}
			if e.Optimize {
				full.native = prepareABINative(full)
				if fixed && full.native != nil {
					full.native.prepareFixed()
				}
			}
			if bound == nil {
				bound = &abiScalarBound{view: append([]byte(nil), view...), entries: make([]*abiScalarPrepared, len(c.Code))}
			}
			// Keep the existing specialized scalar executor at identical entries.
			if bound.entries[r.Entry] != nil {
				continue
			}
			bound.entries[r.Entry] = &abiScalarPrepared{full: full, label: fmt.Sprintf("%s/%d", h.Hex(), r.Entry)}
			bound.view[r.Entry] = byte(abiScalarMarker)
		}
	}
	e.cache[key] = bound
	return bound
}

func abiEntryOwned(b *boundArithmetic, pc uint64) bool {
	return b != nil && (b.regions[pc] != nil || pc < uint64(len(b.entryIndex)) && b.entryIndex[pc] != 0)
}

func (p *abiScalarPrepared) evaluate(scope *ScopeContext) (value uint256.Int, k uint64, prefix int, signature string, ok bool) {
	r := p.region
	stack := scope.Stack.data
	n := len(stack)
	if n < r.Inputs || n-r.Inputs+r.Peak > 1024 || scope.Contract.Gas < p.cost {
		return
	}
	prefix = n - r.Inputs
	ret := &stack[prefix]
	if !scope.Contract.validJumpdest(ret) {
		return
	}
	k = ret.Uint64()
	offset, overflow := stack[prefix+r.Offset].Uint64WithOverflow()
	input := scope.Contract.Input
	if !overflow && offset < uint64(len(input)) {
		if uint64(len(input))-offset >= 32 {
			value.SetBytes(input[offset : offset+32])
		} else {
			var buf [32]byte
			copy(buf[:], input[offset:])
			value.SetBytes(buf[:])
		}
	}
	var cleaned uint256.Int
	switch p.kind {
	case 1:
		cleaned.And(&value, &p.mask)
		if cleaned != value {
			return
		}
	case 2:
		cleaned.ExtendSign(&value, &p.sign)
		if cleaned != value {
			return
		}
	case 3:
		if !value.IsUint64() || value.Uint64() > 1 {
			return
		}
	}
	// opCallDataLoad in this repository also updates ScopeContext.Signature at
	// offset zero. Preserve that extension as well as standard EVM semantics.
	signature = scope.Signature
	if !overflow && offset == 0 {
		word := value.Bytes32()
		signature = string(word[:4])
	}
	ok = true
	return
}

func (cfg *ABIScalarConfig) try(p *abiScalarPrepared, pc *uint64, scope *ScopeContext, frame *abiScalarTapeFrame) bool {
	e := cfg.Engine
	e.Stats.Attempts++
	if p.full != nil {
		if frame != nil {
			panic("full ABI regions do not support recorded-result oracle")
		}
		if e.tryABIRegion(p, pc, scope) {
			return true
		}
		e.Stats.Fallbacks++
		return false
	}
	if frame != nil && !frame.tape.record {
		if frame.next >= len(frame.events) {
			panic("ABI oracle extra entry")
		}
		ev := &frame.events[frame.next]
		frame.next++
		if ev.entry != *pc {
			panic("ABI oracle entry mismatch")
		}
		if frame.tape.verify {
			if scope.Contract.Gas != ev.gas || len(scope.Stack.data) != len(ev.before) {
				panic("ABI oracle admission mismatch")
			}
			for i := range ev.before {
				if ev.before[i] != scope.Stack.data[i] {
					panic("ABI oracle changed input")
				}
			}
		}
		if ev.skip {
			scope.Stack.data[ev.prefix] = ev.value
			scope.Stack.data = scope.Stack.data[:ev.prefix+1]
			scope.Signature = ev.signature
			scope.Contract.Gas -= p.cost
			*pc = ev.exit
			e.Stats.Skipped++
			e.Stats.Instructions += uint64(len(p.region.SuccessPCs))
			e.Stats.Gas += p.cost
			return true
		}
	}
	var event abiScalarTapeEvent
	if frame != nil && frame.tape.record {
		event = abiScalarTapeEvent{entry: *pc, gas: scope.Contract.Gas, before: append([]uint256.Int(nil), scope.Stack.data...)}
	}
	value, k, prefix, sig, ok := p.evaluate(scope)
	if frame != nil && frame.tape.record {
		event.skip, event.value, event.exit, event.prefix, event.signature = ok, value, k, prefix, sig
		frame.events = append(frame.events, event)
	}
	if !ok {
		e.Stats.Fallbacks++
		return false
	}
	scope.Stack.data[prefix] = value
	scope.Stack.data = scope.Stack.data[:prefix+1]
	scope.Signature = sig
	scope.Contract.Gas -= p.cost
	*pc = k
	e.Stats.Skipped++
	e.Stats.Instructions += uint64(len(p.region.SuccessPCs))
	e.Stats.Gas += p.cost
	if e.Capture {
		e.Visits[p.label]++
	}
	return true
}

type ABIScalarTape struct {
	frames         []*abiScalarTapeFrame
	next           int
	record, verify bool
}
type abiScalarTapeFrame struct {
	tape        *ABIScalarTape
	code, input []byte
	depth, next int
	events      []abiScalarTapeEvent
}
type abiScalarTapeEvent struct {
	entry, exit, gas uint64
	prefix           int
	value            uint256.Int
	before           []uint256.Int
	signature        string
	skip             bool
}

func (t *ABIScalarTape) BeginRecord()            { t.frames = nil; t.next = 0; t.record = true; t.verify = false }
func (t *ABIScalarTape) BeginReplay(verify bool) { t.next = 0; t.record = false; t.verify = verify }
func (t *ABIScalarTape) End() {
	if !t.record && t.next != len(t.frames) {
		panic("ABI oracle unconsumed frames")
	}
}
func (t *ABIScalarTape) enter(c *Contract, depth int) *abiScalarTapeFrame {
	if t.record {
		f := &abiScalarTapeFrame{tape: t, code: append([]byte(nil), c.Code...), input: append([]byte(nil), c.Input...), depth: depth}
		t.frames = append(t.frames, f)
		return f
	}
	if t.next >= len(t.frames) {
		panic("ABI oracle extra frame")
	}
	f := t.frames[t.next]
	t.next++
	f.next = 0
	if t.verify && (depth != f.depth || !bytes.Equal(c.Code, f.code) || !bytes.Equal(c.Input, f.input)) {
		panic("ABI oracle frame mismatch")
	}
	return f
}
func (f *abiScalarTapeFrame) finish() {
	if !f.tape.record && f.next != len(f.events) {
		panic("ABI oracle unconsumed entries")
	}
}
