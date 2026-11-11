package vm

import (
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/holiman/uint256"
	"reflect"
)

type wordLoopPlan struct {
	argument                       uint256.Int
	transformOp                    OpCode
	entry, exit, bodyGas, guardGas uint64
}
type regionExtension struct {
	probe               *backgroundEntry
	probeOnly           bool
	demand              *demandPlan
	newEntry, relocated bool
	loop                *wordLoopPlan
	partial             *hotRegion
	fn                  hotRegionFunc
}

var wordLoopCanonical *JumpTable

func init() { wordLoopCanonical = &pragueInstructionSet }

// A structural recipe mined from observed paths; addresses and PUSH widths bind
// from the code. The backedge, exit, stride, stack transfer and active rules are
// rechecked independently of those observations.
func bindWordLoop(code []byte, start uint64, table *JumpTable, bitmap BitVec) *wordLoopPlan {
	pattern := []OpCode{JUMPDEST, DUP4, DUP2, LT, ISZERO, PUSH2, JUMPI, DUP2, DUP2, ADD, MLOAD, DUP4, DUP3, ADD, MSTORE, PUSH1, ADD, PUSH2, JUMP}
	p := &wordLoopPlan{entry: start}
	pc := start
	for i, want := range pattern {
		if pc >= uint64(len(code)) || !bitmap.codeSegment(pc) {
			return nil
		}
		op := OpCode(code[pc])
		ins := table[op]
		push := i == 5 || i == 15 || i == 17
		if (!push && op != want) || (push && (op < PUSH1 || op > PUSH8)) || ins == nil || ins.undefined {
			return nil
		}
		ref := wordLoopCanonical[op]
		if reflect.ValueOf(ins.execute).Pointer() != reflect.ValueOf(ref.execute).Pointer() {
			return nil
		}
		if op == MLOAD || op == MSTORE {
			if ins.dynamicGas == nil || ins.memorySize == nil || reflect.ValueOf(ins.dynamicGas).Pointer() != reflect.ValueOf(ref.dynamicGas).Pointer() || reflect.ValueOf(ins.memorySize).Pointer() != reflect.ValueOf(ref.memorySize).Pointer() {
				return nil
			}
		} else if ins.dynamicGas != nil {
			return nil
		}
		p.bodyGas += ins.constantGas
		if i <= 6 {
			p.guardGas += ins.constantGas
		}
		if push {
			n := uint64(op - PUSH0)
			if pc+n >= uint64(len(code)) {
				return nil
			}
			v := uint64(0)
			for _, b := range code[pc+1 : pc+1+n] {
				v = v<<8 | uint64(b)
			}
			if i == 5 {
				p.exit = v
			}
			if i == 15 && v != 32 {
				return nil
			}
			if i == 17 && v != start {
				return nil
			}
			pc += n
		}
		pc++
	}
	if p.exit < pc || p.exit >= uint64(len(code)) || OpCode(code[p.exit]) != JUMPDEST || !bitmap.codeSegment(p.exit) {
		return nil
	}
	return p
}
func (p *wordLoopPlan) try(pc *uint64, s *ScopeContext, e *EVM) bool {
	n := len(s.Stack.data)
	if n < 4 || n > 1021 || e.abort.Load() || e.Config.Tracer != nil || e.Config.EnablePreimageRecording {
		return false
	}
	st := s.Stack.data
	limit, offset := &st[n-4], &st[n-1]
	if !offset.Lt(limit) {
		if s.Contract.Gas < p.guardGas {
			return false
		}
		s.Contract.Gas -= p.guardGas
		*pc = p.exit
		return true
	}
	if !limit.IsUint64() || !offset.IsUint64() || !st[n-3].IsUint64() || !st[n-2].IsUint64() {
		return false
	}
	i, l := offset.Uint64(), limit.Uint64()
	diff := l - i
	words := diff / 32
	if diff%32 != 0 {
		words++
	}
	if words > 1<<18 {
		return false
	}
	size := words * 32
	next, of := math.SafeAdd(i, size)
	if of {
		return false
	}
	src, of := math.SafeAdd(st[n-2].Uint64(), i)
	if of {
		return false
	}
	dst, of := math.SafeAdd(st[n-3].Uint64(), i)
	if of {
		return false
	}
	srcEnd, of := math.SafeAdd(src, size)
	if of {
		return false
	}
	dstEnd, of := math.SafeAdd(dst, size)
	if of {
		return false
	}
	if src != dst && src < dstEnd && dst < srcEnd {
		return false
	}
	fixed, of := math.SafeMul(words, p.bodyGas)
	if of {
		return false
	}
	fixed, of = math.SafeAdd(fixed, p.guardGas)
	if of {
		return false
	}
	memSize, fee, total, ok := abiNativeMemory(s.Memory, max(srcEnd, dstEnd), fixed, s.Contract.Gas)
	if !ok {
		return false
	}
	s.Memory.Resize(memSize)
	s.Memory.lastGasCost = fee
	{
		copy(s.Memory.store[dst:dstEnd], s.Memory.store[src:srcEnd])
	}

	s.Contract.Gas -= total
	offset.SetUint64(next)
	*pc = p.exit
	if stats := e.Config.RegionExecutionStats; stats != nil {
		stats.LoopWords += words

		stats.LoopBytes += size
	}
	return true
}

func (evm *EVM) bindRegionExtensions(result *abiExecutionPlan, code []byte) {
	evm.bindRegionExtensionsAt(result, code, nil)
}

func (evm *EVM) bindRegionExtensionsAt(result *abiExecutionPlan, code []byte, selected map[uint64]bool) {
	a := result.arithmetic
	if a == nil {
		return
	}
	bitmap := codeBitmap(code)
	starts := partialStarts(code, true)
	for pc := uint64(0); pc < uint64(len(code)); pc++ {
		op := OpCode(code[pc])
		if op >= PUSH1 && op <= PUSH32 {
			pc += uint64(op - PUSH0)
			continue
		}
		if !starts[pc] || selected != nil && !selected[pc] {
			continue
		}
		if result.abi != nil && result.abi.entries[pc] != nil {
			continue
		}
		if evm.preparationCandidateYield != nil {
			evm.preparationCandidateYield()
		}
		x := &regionExtension{loop: bindWordLoop(code, pc, evm.table, bitmap)}
		if paperAblationLevel == 1 {
			x.loop = nil
		}

		if x.loop == nil {
			x.partial = buildPartialRegion(a, code, evm.table, pc)
			if x.partial != nil {
				{
					x.fn = parameterPartialGenerated[parameterPartialKey(x.partial)]
					if x.fn == nil {
						x.fn = expandedPartialGenerated[parameterPartialKey(x.partial)]
					}
					if x.fn != nil && !paperStaticChecksAvailable(x.partial, parameterPartialKey(x.partial)) {
						x.fn = nil // Rebuild from the original summary with checks retained.
					}
				}
			}
		}
		if x.loop == nil {
			x.demand = evm.bindDemand(a, code, pc, x, bitmap)
		}
		x.newEntry = op != JUMPDEST
		if x.partial != nil && x.fn != nil {
			x.relocated = parameterPartialGenerated[parameterPartialKey(x.partial)] != nil
		}
		if x.loop == nil && x.fn == nil && x.demand == nil {
			continue
		}
		id := a.entryIndex[pc]
		var c computeHelperRegion
		if id == 0 {
			if len(a.entryRegions) >= 65535 {
				break
			}
			id = uint16(len(a.entryRegions))
			a.entryIndex[pc] = id
			a.entryRegions = append(a.entryRegions, nil)
			a.computeEntries = append(a.computeEntries, nil)
		} else if a.computeEntries[id] != nil {
			c = *a.computeEntries[id]
		}
		c.extension = x
		a.computeEntries[id] = &c
		a.dispatchCode[pc] = byte(arithmeticEntryMarker)
		if result.abi != nil {
			result.abi.view[pc] = byte(arithmeticEntryMarker)
		}
	}
}

func buildPartialRegion(b *boundArithmetic, code []byte, table *JumpTable, start uint64) *hotRegion {
	r := buildHotRegionWithABI(b, code, table, start, false, ^uint64(0), map[uint64]*abiRegionPrepared{})
	if r == nil {
		return nil
	}
	memory := 0
	for _, n := range r.nodes {
		if n.op == MSTORE || n.op == MLOAD {
			memory++
		}
	}
	if memory < 2 || r.Instructions < 24 || len(r.nodes) > 160 {
		return nil
	}
	return r
}
