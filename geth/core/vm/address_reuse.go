package vm

// Experimental composition of a proved SafeMath summary with its surrounding
// mapping-address preparation. The original SLOAD and SSTORE semantics remain.
import (
	"bytes"
	"fmt"
	"reflect"

	"github.com/ethereum/go-ethereum/params"
	"github.com/holiman/uint256"
)

type addressReusePlan struct {
	previous                            *addressReusePlan
	chain                               *effectChain
	summary                             *summaryRegion
	joined                              *addressJoinPlan
	guarded                             *addressGuardPlan
	expanded                            *addressSlicePlan
	entry, exit, continuation, fixedGas uint64
	sourceFromTop                       int // zero = CALLER; otherwise an unchanged entry stack slot
	slot                                uint256.Int
	helper                              *arithmeticRegion
}

func bindAddressReuse(code []byte, entry uint64, b *boundArithmetic, table *JumpTable) *addressReusePlan {
	if p := bindLegacyAddress(code, entry, b, table); p != nil {
		return p
	}
	p := bindAddressSlice(code, entry, b, table)
	if p == nil {
		return nil
	}
	return &addressReusePlan{entry: p.entry, exit: p.exit, continuation: p.continuation, fixedGas: p.fixedGas, helper: p.helper, expanded: p}
}
func bindLegacyAddress(code []byte, entry uint64, b *boundArithmetic, table *JumpTable) *addressReusePlan {
	if entry+16 >= uint64(len(code)) || b == nil {
		return nil
	}
	if table[SLOAD].constantGas != 0 || table[SLOAD].dynamicGas == nil || reflect.ValueOf(table[SLOAD].dynamicGas).Pointer() != reflect.ValueOf(gasSLoadEIP2929).Pointer() {
		return nil
	}
	head := []byte{byte(KECCAK256), byte(SLOAD), byte(PUSH2), 0, 0, byte(SWAP1), byte(SWAP2), byte(SWAP1), byte(PUSH4), 255, 255, 255, 255, byte(AND), byte(JUMP)}
	head[3], head[4] = code[entry+3], code[entry+4]
	if !bytes.Equal(code[entry:entry+15], head) {
		return nil
	}
	h := b.regions[uint64(head[3])<<8|uint64(head[4])]
	if h == nil || h.peak > 6 || h.syncHot.fixedLayout != 2 || h.syncHot.native == nil || h.extraGas != 0 || !h.constantSuccessGas || h.pathDependentResult {
		return nil
	}
	ret := entry + 15
	// Address calculation is structurally verified, then equality to the first
	// hash preimage is guarded at entry. No sampled operand becomes a constant.
	if ret+6 >= uint64(len(code)) || !bytes.Equal(code[ret:ret+2], []byte{byte(JUMPDEST), byte(PUSH1)}) || code[ret+3] != byte(PUSH1) || code[ret+4] != 0 {
		return nil
	}
	source := OpCode(code[ret+5])
	from := 0
	if source != CALLER {
		if source < DUP4 || source > DUP16 {
			return nil
		}
		from = int(source-DUP1) + 2
	}
	tail := []byte{byte(JUMPDEST), byte(PUSH1), code[ret+2], byte(PUSH1), 0, byte(source), byte(PUSH20)}
	tail = append(tail, bytes.Repeat([]byte{255}, 20)...)
	tail = append(tail, byte(AND), byte(PUSH20))
	tail = append(tail, bytes.Repeat([]byte{255}, 20)...)
	tail = append(tail, byte(AND), byte(DUP2), byte(MSTORE), byte(PUSH1), 32, byte(ADD), byte(SWAP1), byte(DUP2), byte(MSTORE), byte(PUSH1), 32, byte(ADD), byte(PUSH1), 0, byte(KECCAK256), byte(DUP2), byte(SWAP1))
	end := ret + uint64(len(tail))
	if end >= uint64(len(code)) || !bytes.Equal(code[ret:end], tail) || OpCode(code[end]) != SSTORE {
		return nil
	}
	p := &addressReusePlan{entry: entry, exit: end, continuation: ret, sourceFromTop: from, helper: h}
	p.slot.SetUint64(uint64(code[ret+2]))
	p.fixedGas = h.syncHot.admissionGas
	for _, segment := range [][2]uint64{{entry, entry + 15}, {ret, end}} {
		for pc := segment[0]; pc < segment[1]; pc++ {
			op := OpCode(code[pc])
			o := table[op]
			if op == SLOAD {
				continue
			}
			if o == nil || o.undefined || o.dynamicGas != nil && op != MSTORE && op != KECCAK256 {
				return nil
			}
			p.fixedGas += o.constantGas
			if op == KECCAK256 {
				p.fixedGas += 2 * params.Keccak256WordGas
			}
			if op >= PUSH1 && op <= PUSH32 {
				pc += uint64(op - PUSH0)
			}
		}
	}
	return p
}

func (p *addressReusePlan) try(pc *uint64, s *ScopeContext, e *EVM, reuse bool) bool {
	if p.summary != nil {
		return p.summary.try(pc, s, e)
	}
	if p.guarded != nil {
		return p.guarded.try(pc, s, e)
	}
	if p.joined != nil {
		return p.joined.try(pc, s, e)
	}
	if p.expanded != nil {
		return p.expanded.try(pc, s, e)
	}
	n := len(s.Stack.data)
	if n < 4 || n < p.sourceFromTop || n > 1022 || s.Memory.Len() < 64 || e.abort.Load() || e.Config.Tracer != nil || e.Config.EnablePreimageRecording {
		return false
	}
	in := s.Stack.data
	if !in[n-1].IsZero() || !in[n-2].IsUint64() || in[n-2].Uint64() != 64 || !in[n-4].IsUint64() || in[n-4].Uint64() != p.continuation {
		return false
	}
	// Conservative admission: insufficient gas falls back before any mutation,
	// including before a cold SLOAD adds an access-list entry.
	if s.Contract.Gas < p.fixedGas+params.ColdSloadCostEIP2929 {
		return false
	}
	var address, first, second uint256.Int
	if p.sourceFromTop == 0 {
		address.SetBytes(effectiveCaller(s))
	} else {
		address = in[n-p.sourceFromTop]
	}
	address[2] &= 0xffffffff
	address[3] = 0
	first.SetBytes(s.Memory.store[:32])
	second.SetBytes(s.Memory.store[32:64])
	if first != address || second != p.slot {
		return false
	}
	delta := in[n-3]
	s.Contract.Gas -= params.Keccak256Gas + 2*params.Keccak256WordGas
	opKeccak256(pc, e, s)
	key := *s.Stack.peek()
	loadCost, _ := gasSLoadEIP2929(e, s.Contract, s.Stack, s.Memory, 0)
	s.Contract.Gas -= loadCost
	opSload(pc, e, s)
	index, result := p.helper.syncHot.native(delta, *s.Stack.peek())
	cost := p.helper.syncHot.costs[index]
	if cost == syncHotFallback || cost != p.helper.syncHot.admissionGas || e.abort.Load() {
		// A real prefix has executed. Resume exactly after SLOAD; the ordinary
		// instructions retain failure order, INVALID behavior and remaining gas.
		*pc = p.entry + 2
		return true
	}
	if !reuse || paperWithoutReuse {
		e.hasher.Reset()
		e.hasher.Write(s.Memory.store[:64])
		e.hasher.Read(e.hasherBuf[:])
		key.SetBytes(e.hasherBuf[:])
	}
	// Both later MSTOREs write the same bytes already present; no intervening
	// helper memory effect exists. Preserve memory length and original gas.
	s.Contract.Gas -= p.fixedGas - (params.Keccak256Gas + 2*params.Keccak256WordGas)
	in[n-4], in[n-3], in[n-2] = result, result, key
	s.Stack.data = in[:n-1]
	*pc = p.exit // SSTORE executes normally, including static-call rejection.
	return true
}

func effectiveCaller(s *ScopeContext) []byte { a := s.Contract.Caller(); return a[:] }

type RegionExecutionVisit struct {
	Chain                                             bool
	OriginalHashes, ReusedHashes, ExitWrites, Helpers int
	ReuseCandidates                                   int
	Semantic                                          bool
	CodeHash, ContextAddress                          string
	Entry                                             uint64
	Count                                             uint64
	CodeBytes                                         int
}
type RegionExecutionStats struct {
	FlowHelpers, FlowGuards, FlowShared, FlowMulti                                                                     uint64
	DemandAttempts, DemandCommits, DemandInstructions, BranchCommits, ForwardedLoads, SkippedStores                    uint64
	TransformAttempts, TransformCommits, TransformWords                                                                uint64
	BoundaryCommits, BoundaryInstructions, RelocatedCommits                                                            uint64
	LoopAttempts, LoopCommits, LoopWords, LoopBytes, PartialAttempts, PartialCommits, PartialInstructions, EmptyFrames uint64
	BindingNS                                                                                                          int64
	Bindings                                                                                                           uint64
	Attempts, Commits, Resumes                                                                                         uint64
	Capture                                                                                                            map[string]*RegionExecutionVisit `json:",omitempty"`
	DemandCapture                                                                                                      map[string]*DemandExecutionVisit `json:",omitempty"`
}

func (s *RegionExecutionStats) record(p *addressReusePlan, c *Contract) {
	key := fmt.Sprintf("%s/%s/%d", c.CodeHash.Hex(), c.Address().Hex(), p.entry)
	v := s.Capture[key]
	if v == nil {
		v = &RegionExecutionVisit{Semantic: p.summary != nil, CodeHash: c.CodeHash.Hex(), ContextAddress: c.Address().Hex(), Entry: p.entry, CodeBytes: len(c.Code)}
		if p.chain != nil {
			v.Chain = true
			v.OriginalHashes = p.chain.hashes
			v.ReusedHashes = p.chain.reused
			v.ReuseCandidates = p.chain.reused
			if paperWithoutReuse {
				v.ReusedHashes = 0 // candidates remain, but each hash is executed
			}
			v.ExitWrites = len(p.chain.writes)
			v.Helpers = p.chain.helpers
		}
		s.Capture[key] = v
	}
	v.Count++
}

func (evm *EVM) bindComposedRegions(p *abiExecutionPlan, code []byte) *abiExecutionPlan {
	return evm.bindComposedEntries(p, code, nil)
}

// A nil selection preserves complete offline/warm binding. Cold preparation
// visits only entries observed in execution; validation of each region is unchanged.
func (evm *EVM) bindComposedEntries(p *abiExecutionPlan, code []byte, selected map[uint64]bool) *abiExecutionPlan {
	if p.arithmetic == nil || evm.Config.Arithmetic == nil || evm.Config.Arithmetic.Mode != ArithmeticSync || evm.Config.Arithmetic.GasPolicy != OriginalArithmeticGas {
		return p
	}
	a := *p.arithmetic
	a.entryIndex = append([]uint16(nil), a.entryIndex...)
	a.entryRegions = append([]*arithmeticRegion(nil), a.entryRegions...)
	a.computeEntries = append([]*computeHelperRegion(nil), a.computeEntries...)
	a.dispatchCode = append([]byte(nil), a.dispatchCode...)
	if len(a.computeEntries) < len(a.entryRegions) {
		a.computeEntries = append(a.computeEntries, make([]*computeHelperRegion, len(a.entryRegions)-len(a.computeEntries))...)
	}
	result := &abiExecutionPlan{arithmetic: &a, abi: p.abi}
	if p.abi != nil {
		z := *p.abi
		z.view = append([]byte(nil), z.view...)
		result.abi = &z
	}
	if paperAblationLevel == 1 || paperAblationLevel == 2 {
		// The same ordinary-region builder/backend remains available; helper
		// summaries are either absent (Basic) or separate entries (+Summary).
		evm.bindRegionExtensionsAt(result, code, selected)
		a.nativeOnly = true
		for _, id := range a.entryIndex {
			if id != 0 {
				a.nativeOnly = false
				break
			}
		}
		if result.abi != nil {
			for _, r := range result.abi.entries {
				if r != nil {
					a.nativeOnly = false
					break
				}
			}
		}
		return result
	}
	for pc := uint64(0); pc < uint64(len(code)); pc++ {
		op := OpCode(code[pc])
		if op >= PUSH1 && op <= PUSH32 {
			pc += uint64(op - PUSH0)
			continue
		}
		if selected != nil && !selected[pc] {
			continue
		}
		if op != KECCAK256 || a.entryIndex[pc] != 0 || result.abi != nil && result.abi.entries[pc] != nil {
			continue
		}
		if evm.preparationCandidateYield != nil {
			evm.preparationCandidateYield()
		}
		region := bindAddressReuse(code, pc, p.arithmetic, evm.table)
		if region == nil {
			g := bindAddressGuard(code, pc, p.arithmetic, evm.table)
			if g != nil {
				q := g.summary
				region = &addressReusePlan{entry: q.entry, exit: q.exit, continuation: q.continuation, fixedGas: q.fixedGas, helper: q.helper, guarded: g}
			}
		}
		if region == nil {
			continue
		}
		if len(a.entryRegions) >= 65535 {
			break
		}
		a.entryIndex[pc] = uint16(len(a.entryRegions))
		a.entryRegions = append(a.entryRegions, nil)
		a.computeEntries = append(a.computeEntries, &computeHelperRegion{address: region})
		a.dispatchCode[pc] = byte(arithmeticEntryMarker)
		if result.abi != nil {
			result.abi.view[pc] = byte(arithmeticEntryMarker)
		}
	}
	{
		for id, old := range a.computeEntries {
			if old == nil || old.preparation == nil || selected != nil && !selected[old.preparation.region.Entry] {
				continue
			}
			if evm.preparationCandidateYield != nil {
				evm.preparationCandidateYield()
			}
			j := bindAddressJoin(old.preparation, code, p.arithmetic, evm.table)
			if j == nil {
				continue
			}
			q := j.summary
			c := *old
			c.address = &addressReusePlan{entry: q.entry, exit: q.exit, continuation: q.continuation, fixedGas: q.fixedGas, helper: q.helper, joined: j}
			a.computeEntries[id] = &c
		}
	}
	candidates := summaryCandidateEntries(code)
	{
		for _, pc := range candidates {
			if selected != nil && !selected[pc] {
				continue
			}
			if a.entryIndex[pc] != 0 || result.abi != nil && result.abi.entries[pc] != nil {
				continue
			}
			if evm.preparationCandidateYield != nil {
				evm.preparationCandidateYield()
			}
			q := bindSummaryRegion(code, pc, p.arithmetic, evm.table)
			if q == nil || len(a.entryRegions) >= 65535 {
				continue
			}
			a.entryIndex[pc] = uint16(len(a.entryRegions))
			a.entryRegions = append(a.entryRegions, nil)
			a.computeEntries = append(a.computeEntries, &computeHelperRegion{address: &addressReusePlan{entry: q.entry, exit: q.exit, fixedGas: q.gas, summary: q}})
			a.dispatchCode[pc] = byte(arithmeticEntryMarker)
			if result.abi != nil {
				result.abi.view[pc] = byte(arithmeticEntryMarker)
			}
		}
	}

	{
		var sharedBitmap BitVec
		{
			sharedBitmap = codeBitmap(code)
		}

		// Reuse the existing entry table: untagged instructions never gain a lookup.

		for _, pc := range candidates {
			if selected != nil && !selected[pc] {
				continue
			}
			if result.abi != nil && result.abi.entries[pc] != nil {
				continue
			}
			if evm.preparationCandidateYield != nil {
				evm.preparationCandidateYield()
			}
			var q *effectChain
			{
				q = bindEffectChainShared(code, pc, p.arithmetic, evm.table, sharedBitmap)
			}

			if q == nil {
				continue
			}
			id := a.entryIndex[pc]
			var previous *addressReusePlan
			var c computeHelperRegion
			if id == 0 {
				if len(a.entryRegions) >= 65535 {
					continue
				}
				id = uint16(len(a.entryRegions))
				a.entryIndex[pc] = id
				a.entryRegions = append(a.entryRegions, nil)
				a.computeEntries = append(a.computeEntries, nil)
			} else if a.computeEntries[id] != nil {
				c = *a.computeEntries[id]
				previous = c.address
			}
			c.address = &addressReusePlan{entry: q.entry, exit: q.exit, chain: q, previous: previous}
			a.computeEntries[id] = &c
			a.dispatchCode[pc] = byte(arithmeticEntryMarker)
			if result.abi != nil {
				result.abi.view[pc] = byte(arithmeticEntryMarker)
			}
		}
	}

	{
		evm.bindRegionExtensionsAt(result, code, selected)
		evm.bindRegionJoinsAt(result, code, selected)
	}
	{
		empty := true
		for _, id := range a.entryIndex {
			if id != 0 {
				empty = false
				break
			}
		}
		if empty && result.abi != nil {
			for _, r := range result.abi.entries {
				if r != nil {
					empty = false
					break
				}
			}
		}
		a.nativeOnly = empty
	}
	// Warm preparation uses the same direct executable headers as background
	// publication; a pending slot is resolved once on its first ready entry.
	prepareCompactEntries(result)

	return result
}
