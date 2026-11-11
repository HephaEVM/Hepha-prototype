// Author: ClutchZsJ
// This file describes the template matcher built upon go-ethereum library.

package vm

import (
	"math"
	"sort"
	"sync"

	"github.com/ethereum/go-ethereum/common"
)

var (
	// For Solidity Complier v0.8.x overflow check patterns
	// overflow_templateEntrance: entrance pattern [PUSH0/PUSH1, PUSH2 aa, DUP3, PUSH2 bb]
	overflow_templateEntrance = []OpCode{PUSH0, PUSH2, DUP3, PUSH2}
	// overflow_templateReuse: Check segment at 'bb': [PUSH0/PUSH1, DUP2, SWAP1, POP, SWAP2, SWAP1, POP]
	overflow_templateReuse = []OpCode{PUSH0, DUP2, SWAP1, POP, SWAP2, SWAP1, POP}
	// overflow_templateLast: Check segment at 'aa': [SWAP2, POP, PUSH2 cxc, DUP4, PUSH2 bb]
	overflow_templateLast = []OpCode{SWAP2, POP, PUSH2, DUP4, PUSH2}

	// overflow_templateReuseWithMask: Reuse segment with AND mask in the case of smaller uints (<256 bits)
	overflow_templateReuseWithMask = []OpCode{DUP3, AND, SWAP1, POP, SWAP2, SWAP1, POP}
	// overflow_templateReuseWithSignExtend: Reuse segment with SIGNEXTEND for int types
	overflow_templateReuseWithSignExtend = []OpCode{PUSH0, DUP2, PUSH1, SIGNEXTEND, SWAP1, POP, SWAP2, SWAP1, POP}

	// arithmeticPrefixEntrance: Common prefix for arithmetic operations [SWAP3, POP, DUP3]
	arithmeticPrefixEntrance = []OpCode{SWAP3, POP, DUP3}
	// arithmeticPrefixMulAddSub: Next opcode for MUL/ADD/SUB [DUP3]
	arithmeticPrefixMulAddSub = []OpCode{DUP3}
	// arithmeticPrefixDivMod: Next opcodes for DIV/MOD [PUSH2 xx, JUMPI]
	arithmeticPrefixDivMod = []OpCode{PUSH2, JUMPI}
	// divModPrefix: Common prefix in the final DIV/MOD block [DUP3, DUP3]
	divModPrefix = []OpCode{DUP3, DUP3}

	// jump entrance for the check of MIN_INT * (-1)
	jumpEntrance = []OpCode{PUSH2}
	// template for MIN_INT * (-1) check
	minIntMulNeg1Check = []OpCode{SWAP2, POP, PUSH32, DUP5, EQ}
	// template for MIN_INT / (-1) check
	minIntDivNeg1Check = []OpCode{PUSH1, PUSH0, SUB, DUP4, EQ, PUSH32}

	overflow_inlineAdd  = []OpCode{DUP1, DUP3, ADD, DUP1, DUP3, GT, ISZERO, PUSH2, JUMPI}
	overflow_inlineAdd2 = []OpCode{PUSH1, DUP3, NOT, DUP3, GT, ISZERO, PUSH2, JUMPI}
	overflow_inlineSub  = []OpCode{DUP2, DUP2, SUB, DUP2, DUP2, GT, ISZERO, PUSH2, JUMPI}
)

var (
	// For Solidity Complier (<v0.8.0) safe math overflow check patterns
	safemath_Add  = []OpCode{DUP3, DUP5, ADD, SWAP1, POP, DUP4, DUP2, LT, ISZERO} // OpenZeppelin solidity version<0.6.0
	safemath_Add2 = []OpCode{DUP3, ADD, DUP4, DUP2, LT, ISZERO}                   // OpenZeppelin solidity version>=0.6.0<0.8.0
	safemath_Add3 = []OpCode{DUP3, DUP5, ADD, SWAP2, POP, DUP2, LT, ISZERO}       // SafeTransfer

	safemath_SubEntrance  = []OpCode{DUP3, GT, ISZERO, PUSH2}                 // OpenZeppelin solidity version>=0.6.0<0.8.0
	safemath_SubEntrance2 = []OpCode{DUP3, GT, ISZERO, ISZERO, ISZERO, PUSH2} // OpenZeppelin solidity version<0.6.0
	safemath_SubEntrance3 = []OpCode{DUP4, LT, ISZERO, ISZERO, ISZERO, PUSH2}
	safemath_SubEntrance4 = []OpCode{DUP3, DUP5, SUB, SWAP2, POP, DUP2, GT, ISZERO} // SafeTransfer

	safemath_MulEntrance                  = []OpCode{PUSH1, DUP1, PUSH1, DUP5, EQ, ISZERO, PUSH2}
	safemath_MulCheck                     = []OpCode{DUP3, DUP5, MUL, SWAP1, POP, DUP3, DUP5, DUP3, DUP2}
	safemath_DivEntranceWithoutZeroCheck  = []OpCode{DUP2, PUSH2}
	safemath_DivEntranceWithoutZeroCheck2 = []OpCode{DUP2, ISZERO, ISZERO, PUSH2} // JUMPI is checked together with the entrance pattern
	safemath_DivEntranceWithZeroCheck     = []OpCode{DUP3, GT, PUSH2}
	safemath_DivModReuse                  = []OpCode{DUP3, DUP5, DUP2, PUSH2}
	safemath_ModEntrance                  = []OpCode{DUP3, SUB, PUSH2}
	//safemath_intSub           = []OpCode{DUP3, DUP5, SUB, SWAP1, POP, PUSH0, DUP4, SLT}
	//safemath_mulEntrance = []OpCode{DUP4, SUB, PUSH2}
	//safemath_uintMulCheck = []OpCode{PUSH0, DUP3, DUP5, MUL}
	//safemath_intMulCheck  = []OpCode{PUSH32, DUP4, EQ, DUP1, ISZERO, PUSH2}
	//safemath_intDivCheck       = []OpCode{PUSH32, DUP3, EQ, DUP1, ISZERO, PUSH2}
)

var (
	// For Solidity Complier out-of-bounds check patterns
	//outOfBound_AllConstant1           = []OpCode{PUSH1, PUSH1, PUSH1, PUSH1, DUP2, LT}
	//outOfBound_AllConstant1_ZeroIndex = []OpCode{PUSH1, PUSH1, PUSH0, PUSH1, DUP2, LT}
	//outOfBound_AllConstant2           = []OpCode{PUSH1, DUP1, PUSH1, PUSH1, DUP2, LT}
	//outOfBound_AllConstant2_ZeroIndex = []OpCode{PUSH1, DUP1, PUSH0, PUSH1, DUP2, LT}
	//outOfBound_AllConstant3           = []OpCode{PUSH1, PUSH0, PUSH1, PUSH1, DUP2, LT} // the start storage slot of this array is 0
	//outOfBound_AllConstant3_ZeroIndex = []OpCode{PUSH1, PUSH0, PUSH0, PUSH1, DUP2, LT}

	outOfBound_AllConstantSuffix = []OpCode{PUSH1, DUP2, LT}
	outOfBound_ConstantVarSuffix = []OpCode{DUP2, SLOAD, DUP2, LT}

	//outOfBound_ConstantVar            = []OpCode{PUSH1, PUSH1, PUSH1, DUP2, SLOAD, DUP2, LT}
	//outOfBound_ConstantVar_ZeroIndex  = []OpCode{PUSH1, PUSH1, PUSH0, DUP2, SLOAD, DUP2, LT}
	//outOfBound_ConstantVar2           = []OpCode{PUSH1, PUSH0, PUSH1, DUP2, SLOAD, DUP2, LT} // the storage slot of the array length is 0
	//outOfBound_ConstantVar2_ZeroIndex = []OpCode{PUSH1, PUSH0, PUSH0, DUP2, SLOAD, DUP2, LT}
	//
	//outOfBound_AllConstantSingle            = []OpCode{PUSH1, PUSH1, PUSH1, DUP2, LT} // e.g., only exists a[i] without any operands
	//outOfBound_AllConstantSingle_ZeroIndex  = []OpCode{PUSH1, PUSH0, PUSH1, DUP2, LT}
	//outOfBound_AllConstantSingle2           = []OpCode{PUSH0, PUSH1, PUSH1, DUP2, LT} // the start storage slot of this array is 0
	//outOfBound_AllConstantSingle2_ZeroIndex = []OpCode{PUSH0, PUSH0, PUSH1, DUP2, LT}
	//
	//outOfBound_ConstantVarSingle            = []OpCode{PUSH1, PUSH1, DUP2, SLOAD, DUP2, LT}
	//outOfBound_ConstantVarSingle_ZeroIndex  = []OpCode{PUSH1, PUSH0, DUP2, SLOAD, DUP2, LT}
	//outOfBound_ConstantVarSingle2           = []OpCode{PUSH0, PUSH1, DUP2, SLOAD, DUP2, LT} // the storage slot of the array length is 0
	//outOfBound_ConstantVarSingle2_ZeroIndex = []OpCode{PUSH0, PUSH0, DUP2, SLOAD, DUP2, LT}

	// Loop patterns
	//loopForward_Prefix1 = []OpCode{PUSH1, SLOAD}            // check: i<num or i<=num (num is from storage)
	//loopForward_Prefix2 = []OpCode{PUSH0, SLOAD}            // check: i<num or i<=num (num is from storage, slot is 0)
	//loopForward_Prefix3 = []OpCode{PUSH1}                   // check: i<num or i<=num (num is constant)
	//loopForward_Prefix4 = []OpCode{DUP2}                    // check: i<num or i<=num (num is temporary variable)
	loopForward_Suffix1 = []OpCode{DUP2, LT, ISZERO, PUSH2} // check: i<num
	loopForward_Suffix2 = []OpCode{DUP2, GT, PUSH2}         // check: i<=num

	//loopBackward_Prefix1 = []OpCode{PUSH1}                   // check: i>num (num>0)
	//loopBackward_Prefix2 = []OpCode{PUSH0}                   // check: i>0
	loopBackward_Suffix1 = []OpCode{DUP2, GT, ISZERO, PUSH2} // check: i>0
	loopBackward_Suffix2 = []OpCode{DUP2, LT, PUSH2}         // check: i>=0

	// exists a[i] with operands
	//outOfBound_LoopAllConstant1 = []OpCode{PUSH1, PUSH0, DUP3, PUSH1, DUP2, LT}
	//outOfBound_LoopAllConstant2 = []OpCode{PUSH1, PUSH1, DUP3, PUSH1, DUP2, LT}
	//outOfBound_LoopAllConstant3 = []OpCode{PUSH0, PUSH0, DUP3, PUSH1, DUP2, LT}
	//outOfBound_LoopAllConstant4 = []OpCode{PUSH0, PUSH1, DUP3, PUSH1, DUP2, LT}
	//outOfBound_LoopConstantVar1 = []OpCode{PUSH1, PUSH0, DUP3, DUP2, SLOAD, DUP2, LT}
	//outOfBound_LoopConstantVar2 = []OpCode{PUSH1, PUSH1, DUP3, DUP2, SLOAD, DUP2, LT}
	//outOfBound_LoopConstantVar3 = []OpCode{PUSH0, PUSH0, DUP3, DUP2, SLOAD, DUP2, LT}
	//outOfBound_LoopConstantVar4 = []OpCode{PUSH0, PUSH1, DUP3, DUP2, SLOAD, DUP2, LT}
	outOfBound_LoopConstantSuffix = []OpCode{DUP3, PUSH1, DUP2, LT}
	outOfBound_LoopVarSuffix      = []OpCode{DUP3, DUP2, SLOAD, DUP2, LT}

	// only exists a[i] without any operands
	//outOfBound_LoopAllConstantSingle1 = []OpCode{PUSH1, DUP2, PUSH1, DUP2, LT}
	//outOfBound_LoopAllConstantSingle2 = []OpCode{PUSH0, DUP2, PUSH1, DUP2, LT}
	//outOfBound_LoopConstantVarSingle1 = []OpCode{PUSH1, DUP2, DUP2, SLOAD, DUP2, LT}
	//outOfBound_LoopConstantVarSingle2 = []OpCode{PUSH0, DUP2, DUP2, SLOAD, DUP2, LT}
	outOfBound_LoopConstantSingleSuffix = []OpCode{DUP2, PUSH1, DUP2, LT}
	outOfBound_LoopVarSingleSuffix      = []OpCode{DUP2, DUP2, SLOAD, DUP2, LT}
)

var (
	// panic_trigger: instructions that terminate execution immediately
	panic_trigger = []OpCode{PUSH32, PUSH0, MSTORE, PUSH1, PUSH1, MSTORE, PUSH1, PUSH0, REVERT}
)

// OperationType represents the type of arithmetic operation
type OperationType int
type OutOfBoundType int
type LoopType int
type OverflowOperandOrder int

const (
	OpMul OperationType = iota
	OpAdd
	OpSub
	OpDiv
	OpMod
	OpSDiv
	OpSMod
)

const (
	OutOfBound_AllConstant        OutOfBoundType = iota // length and index are both constants
	OutOfBound_AllConstantSingle                        // length and index are both constants, and only exists a[i] without any operands
	OutOfBound_ConstantVar                              // length is variable (from storage), index is constant
	OutOfBound_ConstantVarSingle                        // length is variable (from storage), index is constant, and only exists a[i] without any operands
	OutOfBound_LoopConstant                             // length is constant
	OutOfBound_LoopConstantSingle                       // length is constant, and only exists a[i] without any operands
	OutOfBound_LoopVar                                  // length is variable (from storage)
	OutOfBound_LoopVarSingle                            // length is variable (from storage), and only exists a[i] without any operands
)

const (
	LoopType_Forward_Constant LoopType = iota
	LoopType_Forward_Var
	LoopType_Backward
)

const (
	// OperandOrderByVersion means the final operand order should be derived
	// from the matched pattern/version information.
	OperandOrderByVersion OverflowOperandOrder = iota
	// OperandOrderLegacy means lhs=second-from-top, rhs=top.
	OperandOrderLegacy
	// OperandOrderModern means lhs=top, rhs=second-from-top.
	OperandOrderModern
)

func (o OperationType) String() string {
	switch o {
	case OpMul:
		return "MUL"
	case OpAdd:
		return "ADD"
	case OpSub:
		return "SUB"
	case OpDiv:
		return "DIV"
	case OpMod:
		return "MOD"
	case OpSDiv:
		return "SDIV"
	case OpSMod:
		return "SMOD"
	default:
		return "UNKNOWN"
	}
}

func (o OutOfBoundType) String() string {
	switch o {
	case OutOfBound_AllConstant:
		return "AllConstant"
	case OutOfBound_AllConstantSingle:
		return "AllConstantSingle"
	case OutOfBound_ConstantVar:
		return "ConstantVar"
	case OutOfBound_ConstantVarSingle:
		return "ConstantVarSingle"
	case OutOfBound_LoopConstant:
		return "LoopConstant"
	case OutOfBound_LoopConstantSingle:
		return "LoopConstantSingle"
	case OutOfBound_LoopVar:
		return "LoopVar"
	case OutOfBound_LoopVarSingle:
		return "LoopVarSingle"
	default:
		return "UNKNOWN"
	}
}

func (l LoopType) String() string {
	switch l {
	case LoopType_Forward_Constant:
		return "Forward_Constant"
	case LoopType_Forward_Var:
		return "Forward_Var"
	case LoopType_Backward:
		return "Backward"
	default:
		return "UNKNOWN"
	}
}

// OverflowCheckPattern represents a matched overflow check pattern
type OverflowCheckPattern struct {
	StartPositions []uint64      // Positions of code segments that match the pattern
	Operation      OperationType // The arithmetic operation type
	UintSize       int           // The size of the uint in bits
	IntSize        int           // The size of the int in bits
	OldVersion     bool          // Indicate whether it is the solidity version before 0.8.0
}

type OverflowJumpPred struct {
	SegmentStart uint64
	TriggerPC    uint64
	OperandOrder OverflowOperandOrder
}

// TemplateRuntime stores only the runtime metadata needed during execution.
type TemplateRuntime struct {
	contractAddr  common.Address
	runtimeLookup []*SegmentRuntimeMeta
	smallestPC    uint64
	panicTrigger  map[uint64]uint64
	matcherExist  bool
}

// SegmentRuntimeMeta stores the overflow-skip runtime metadata for one code segment.
type SegmentRuntimeMeta struct {
	Start        uint64
	End          uint64
	Pattern      *OverflowCheckPattern
	IsWrapper    bool
	IsAhead      bool
	TriggerPC    uint64
	OperandOrder OverflowOperandOrder
}

// OutOfBoundCheckPattern represents a matched out-of-bounds check pattern
type OutOfBoundCheckPattern struct {
	StartPosition uint64
	CheckPosition uint64 // position of the check start (only used in loop cases)
	JumpPosition  uint64
	Type          OutOfBoundType
	Index         uint64
	Length        uint64 // for ConstantVar pattern, this is the storage slot
	Stack         []uint64
}

type LoopInfo struct {
	StartPosition uint64
	Patterns      []OutOfBoundCheckPattern
	Type          LoopType
	LoopNum       uint64 // 0 means require check in the first loop execution (or is the storage slot of loop num)
}

// CodeSegment represents a basic segment of bytecode
type CodeSegment struct {
	Start uint64   // Starting position
	End   uint64   // Ending position
	Ops   []OpCode // Opcodes in the segment (excluding operands)
	OpPos []uint64 // Positions of opcodes in the original bytecode
}

// TemplateMatcher handles bytecode template matching for overflow/out-of-bounds checks
type TemplateMatcher struct {
	contractAddr        common.Address
	bytecode            []byte
	segments            []CodeSegment
	segmentIndexByStart map[uint64]int
	runtimeLookup       []*SegmentRuntimeMeta
	smallestPC          uint64
	matcherExist        bool
	runtime             *TemplateRuntime
	//overflowPatterns      map[uint64]*OverflowCheckPattern // keyed by segment start position
	overflowJumpPreds map[uint64][]OverflowJumpPred
	//outOfBoundPatterns            []OutOfBoundCheckPattern
	//outOfBoundCheckPatternsInLoop map[uint64]*LoopInfo // keyed by loop segment start position
	panicTrigger map[uint64]uint64 // mapping from panic number to position
	mu           sync.RWMutex
}

// NewTemplateMatcher creates a new template matcher for the given bytecode
func NewTemplateMatcher(addr common.Address, bytecode []byte) *TemplateMatcher {
	return &TemplateMatcher{
		contractAddr:        addr,
		bytecode:            bytecode,
		segments:            make([]CodeSegment, 0),
		segmentIndexByStart: make(map[uint64]int),
		runtimeLookup:       make([]*SegmentRuntimeMeta, len(bytecode)),
		smallestPC:          math.MaxUint64,
		//overflowPatterns:      make(map[uint64]*OverflowCheckPattern),
		overflowJumpPreds: make(map[uint64][]OverflowJumpPred),
		//outOfBoundPatterns:            make([]OutOfBoundCheckPattern, 0),
		//outOfBoundCheckPatternsInLoop: make(map[uint64]*LoopInfo),
		panicTrigger: make(map[uint64]uint64),
	}
}

// extractOperand extracts the operand value from bytecode at the given position
func (tm *TemplateMatcher) extractOperand(pos uint64, size int) uint64 {
	if int(pos)+size > len(tm.bytecode) {
		return 0
	}

	var result uint64
	for i := 0; i < size; i++ {
		result = (result << 8) | uint64(tm.bytecode[pos+uint64(i)])
	}
	return result
}

// parseCodeSegments extracts all code segments based on basic block analysis.
// A basic block starts at a "block header" and ends at the instruction before the next "block header"
// or at a terminating instruction. We are interested in segments that start with JUMPDEST and end with JUMP.
func (tm *TemplateMatcher) parseCodeSegments() {
	if len(tm.bytecode) == 0 {
		return
	}

	// Step 1: Identify all block headers
	headers := map[uint64]struct{}{0: {}} // PC=0 is always a header
	pc := uint64(0)
	for pc < uint64(len(tm.bytecode)) {
		op := OpCode(tm.bytecode[pc])

		if op == JUMPDEST {
			headers[pc] = struct{}{}
		} else if op.IsTerminal() { // JUMP, STOP, REVERT, RETURN, SELFDESTRUCT, INVALID
			if pc+1 < uint64(len(tm.bytecode)) {
				headers[pc+1] = struct{}{}
			}
		}

		// Move to the next instruction
		if op >= PUSH0 && op <= PUSH32 {
			if op == PUSH0 {
				// PUSH0 has no operands
				pc++
			} else {
				pushSize := int(op - PUSH1 + 1)
				pc += uint64(pushSize) + 1
			}
		} else {
			pc++
		}
	}

	// Convert map keys to a slice and sort them
	sortedHeaders := make([]uint64, 0, len(headers))
	for h := range headers {
		sortedHeaders = append(sortedHeaders, h)
	}
	sort.Slice(sortedHeaders, func(i, j int) bool {
		return sortedHeaders[i] < sortedHeaders[j]
	})

	// Step 2: Form basic blocks
	tm.segments = make([]CodeSegment, 0, len(sortedHeaders))
	tm.segmentIndexByStart = make(map[uint64]int, len(sortedHeaders))
	for i := 0; i < len(sortedHeaders); i++ {
		startPC := sortedHeaders[i]
		endPC := uint64(len(tm.bytecode))
		if i+1 < len(sortedHeaders) {
			endPC = sortedHeaders[i+1]
		}

		if startPC >= endPC {
			continue
		}

		segment := CodeSegment{
			Start: startPC,
			Ops:   make([]OpCode, 0),
			OpPos: make([]uint64, 0),
		}

		currentPC := startPC
		for currentPC < endPC {
			op := OpCode(tm.bytecode[currentPC])
			segment.Ops = append(segment.Ops, op)
			segment.OpPos = append(segment.OpPos, currentPC)

			// Move to the next instruction
			if op >= PUSH0 && op <= PUSH32 {
				if op == PUSH0 {
					// PUSH0 has no operands
					currentPC++
				} else {
					pushSize := int(op - PUSH1 + 1)
					currentPC += uint64(pushSize) + 1
				}
			} else {
				currentPC++
			}
		}
		segment.End = segment.OpPos[len(segment.OpPos)-1]

		// Identify potential overflow-check wrappers and record their jump targets
		// together with operand-order metadata for runtime skipping.
		for j := 0; j < len(segment.Ops)-1; j++ {
			if segment.Ops[j] != PUSH2 {
				continue
			}
			jumpPC := segment.OpPos[j]
			if jumpPC+1 > uint64(len(tm.bytecode)) {
				continue
			}

			order := OperandOrderByVersion
			matched := false
			switch {
			case segment.Ops[j+1] == JUMP || segment.Ops[j+1] == JUMPI:
				matched = true
			case (j+6) < len(segment.Ops) &&
				segment.Ops[j+1] == SWAP1 &&
				segment.Ops[j+2] == SWAP2 &&
				segment.Ops[j+3] == SWAP1 &&
				segment.Ops[j+4] == PUSH4 &&
				segment.Ops[j+5] == AND &&
				segment.Ops[j+6] == JUMP:
				matched = true
				order = OperandOrderModern
			}
			if !matched {
				continue
			}

			target := tm.extractOperand(jumpPC+1, 2)
			//fmt.Println("contract address:", tm.contractAddr.String(), "target is:", target)
			tm.overflowJumpPreds[target] = append(tm.overflowJumpPreds[target], OverflowJumpPred{
				SegmentStart: segment.Start,
				TriggerPC:    jumpPC,
				OperandOrder: order,
			})
		}

		tm.segmentIndexByStart[segment.Start] = len(tm.segments)
		tm.segments = append(tm.segments, segment)
	}
}

// initRuntimeLookup clears the unified runtime lookup before a fresh scan.
func (tm *TemplateMatcher) initRuntimeLookup() {
	for i := range tm.runtimeLookup {
		tm.runtimeLookup[i] = nil
	}
	tm.smallestPC = math.MaxUint64
}

// setSegmentRuntimeMeta records overflow runtime metadata on a jumpdest key.
func (tm *TemplateMatcher) setSegmentRuntimeMeta(segmentStart uint64, pattern *OverflowCheckPattern, isWrapper, isAhead bool, triggerPC uint64, order OverflowOperandOrder) {
	if segmentStart >= uint64(len(tm.runtimeLookup)) {
		return
	}
	if !isAhead {
		if existing := tm.runtimeLookup[segmentStart]; existing != nil && existing.Pattern != nil {
			if existing.IsAhead || !existing.IsWrapper {
				return
			}
		}
	}
	meta := &SegmentRuntimeMeta{
		Start:        segmentStart,
		End:          segmentStart,
		Pattern:      pattern,
		IsWrapper:    isWrapper,
		IsAhead:      isAhead,
		TriggerPC:    triggerPC,
		OperandOrder: order,
	}
	if seg := tm.findSegmentByPosition(segmentStart); seg != nil {
		meta.End = seg.End
	}
	meta.Pattern = pattern
	meta.IsWrapper = isWrapper
	meta.IsAhead = isAhead
	meta.TriggerPC = triggerPC
	meta.OperandOrder = order
	tm.runtimeLookup[segmentStart] = meta
	if pattern != nil && segmentStart < tm.smallestPC {
		tm.smallestPC = segmentStart
	}
	if pattern != nil {
		tm.matcherExist = true
	}
}

// buildOverflowLookups converts a matched overflow pattern into segment runtime metadata.
func (tm *TemplateMatcher) buildOverflowLookups(pattern *OverflowCheckPattern) {
	if pattern == nil || len(pattern.StartPositions) == 0 {
		return
	}
	startPosition := pattern.StartPositions[0]
	preds := tm.overflowJumpPreds[startPosition]
	defaultOrder := OperandOrderModern
	if pattern.OldVersion {
		defaultOrder = OperandOrderLegacy
	}
	if len(preds) == 0 {
		tm.setSegmentRuntimeMeta(startPosition, pattern, false, false, startPosition, defaultOrder)
		return
	}
	tm.setSegmentRuntimeMeta(startPosition, pattern, true, true, startPosition, defaultOrder)
	for _, pred := range preds {
		order := pred.OperandOrder
		if order == OperandOrderByVersion {
			order = defaultOrder
		}
		tm.setSegmentRuntimeMeta(pred.SegmentStart, pattern, true, false, pred.TriggerPC, order)
	}
}

// matchTemplate checks if a segment matches a given opcode pattern
// Returns true if matched, along with any extracted operand values and the index after matching
func (tm *TemplateMatcher) matchTemplate(segment CodeSegment, template []OpCode, startIdx int) (bool, []uint64, int) {
	if len(segment.Ops)-startIdx < len(template) {
		return false, nil, 0
	}

	operands := make([]uint64, 0)
	segIdx := startIdx

	for tplIdx := 0; tplIdx < len(template); tplIdx++ {
		if segIdx >= len(segment.Ops) {
			return false, nil, 0
		}

		tplOp := template[tplIdx]
		segOp := segment.Ops[segIdx]

		// For PUSH instructions, we match the PUSH type and extract the operand
		if tplOp >= PUSH0 && tplOp <= PUSH32 {
			if segOp == tplOp {
				// Extract operand value
				if segOp == PUSH0 {
					// PUSH0 pushes 0 onto the stack, no operand to extract
					operands = append(operands, 0)
				} else {
					pushSize := int(segOp - PUSH1 + 1)
					operand := tm.extractOperand(segment.OpPos[segIdx]+1, pushSize)
					operands = append(operands, operand)
				}
				segIdx++
			} else {
				return false, nil, 0
			}
		} else {
			// Exact opcode match required
			if segOp != tplOp {
				return false, nil, 0
			}
			segIdx++
		}
	}

	return true, operands, segIdx
}

// findSegmentByPosition finds a code segment that starts at or contains the given position
func (tm *TemplateMatcher) findSegmentByPosition(pos uint64) *CodeSegment {
	if idx, ok := tm.segmentIndexByStart[pos]; ok {
		return &tm.segments[idx]
	}
	return nil
}

// matchReuseSegment checks if a segment matches one of the reuse patterns.
// It returns whether it matched, the size in bits, and the indicator of different cases
func (tm *TemplateMatcher) matchReuseSegment(segment CodeSegment, startIdx int) (bool, int, int) {
	matched, _, _ := tm.matchTemplate(segment, []OpCode{PUSH0}, startIdx)
	matched2, _, _ := tm.matchTemplate(segment, []OpCode{PUSH1}, startIdx)
	if !matched && !matched2 {
		return false, 0, 0
	}

	nextIdx := startIdx + 1
	if nextIdx >= len(segment.Ops) {
		return false, 0, 0
	}

	// 1. Check for the standard uint256/int256 reuse pattern (no AND mask)
	matched, _, _ = tm.matchTemplate(segment, overflow_templateReuse[1:], nextIdx)
	if matched {
		return true, 256, 1
	}

	// 2. Check for reuse pattern with AND mask for smaller uints
	// Check for a PUSH<N> instruction
	op := segment.Ops[nextIdx]
	if op >= PUSH1 && op <= PUSH31 { // PUSH1-PUSH31 for uint8-uint248
		pushSize := int(op - PUSH1 + 1)
		// Match suffix
		matched, _, _ = tm.matchTemplate(segment, overflow_templateReuseWithMask, nextIdx+1)
		if matched {
			return true, pushSize * 8, 2
		}
	}

	// 3. Check for reuse pattern with SIGNEXTEND for int types
	matched, ops, _ := tm.matchTemplate(segment, overflow_templateReuseWithSignExtend[1:], nextIdx)
	if !matched || len(ops) < 1 {
		return false, 0, 0
	}
	opbytes := ops[0]
	return true, int(opbytes+1) * 8, 3
}

// matchOutOfBoundPattern checks if a segment contains out-of-bounds check patterns
func (tm *TemplateMatcher) matchOutOfBoundPattern(segment CodeSegment) *OutOfBoundCheckPattern {
	// Iterate through ops
	var pattern *OutOfBoundCheckPattern
	i := 0
	push1Size := 1
	push2Size := 2
	for i < len(segment.Ops) {
		if segment.Ops[i] == PUSH1 || segment.Ops[i] == PUSH0 {
			var firstOperand uint64
			var secondOperand uint64
			if segment.Ops[i] == PUSH0 {
				firstOperand = 0
			} else {
				firstOperand = tm.extractOperand(segment.OpPos[i]+1, push1Size)
			}

			if segment.Ops[i+1] == PUSH1 || segment.Ops[i+1] == PUSH0 {
				if segment.Ops[i+1] == PUSH0 {
					secondOperand = 0
				} else {
					secondOperand = tm.extractOperand(segment.OpPos[i+1]+1, push1Size)
				}

				matched, operands, nextIdx := tm.matchTemplate(segment, outOfBound_AllConstantSuffix, i+2)
				// AllConstantSingle pattern matched
				if matched && len(operands) > 0 {
					jumpPos := tm.extractOperand(segment.OpPos[nextIdx]+1, push2Size)
					pattern = &OutOfBoundCheckPattern{
						StartPosition: uint64(i),
						JumpPosition:  jumpPos,
						Type:          OutOfBound_AllConstantSingle,
						Index:         secondOperand,
						Length:        operands[0],
						Stack:         []uint64{firstOperand, secondOperand}, // [first storage slot, index]
					}
					break
				}

				matched, operands, nextIdx = tm.matchTemplate(segment, outOfBound_ConstantVarSuffix, i+2)
				// ConstantVarSingle pattern matched
				if matched && len(operands) > 0 {
					jumpPos := tm.extractOperand(segment.OpPos[nextIdx]+1, push2Size)
					pattern = &OutOfBoundCheckPattern{
						StartPosition: uint64(i),
						JumpPosition:  jumpPos,
						Type:          OutOfBound_ConstantVarSingle,
						Index:         secondOperand,
						Length:        firstOperand,
						Stack:         []uint64{firstOperand, secondOperand}, // [length storage slot, index]
					}
					break
				}

				matched, operands, nextIdx = tm.matchTemplate(segment, []OpCode{PUSH1}, i+2)
				if matched && len(operands) > 0 {
					// AllConstant pattern matched
					if matched2, operands2, nextIdx2 := tm.matchTemplate(segment, outOfBound_AllConstantSuffix, nextIdx); matched2 && len(operands2) > 0 {
						jumpPos := tm.extractOperand(segment.OpPos[nextIdx2]+1, push2Size)
						pattern = &OutOfBoundCheckPattern{
							StartPosition: uint64(i),
							JumpPosition:  jumpPos,
							Type:          OutOfBound_AllConstant,
							Index:         operands[0],
							Length:        operands2[0],
							Stack:         []uint64{firstOperand, secondOperand, operands[0]}, // [operand, first storage slot, index]
						}
						break
					}

					// ConstantVar pattern matched
					if matched2, _, nextIdx2 := tm.matchTemplate(segment, outOfBound_ConstantVarSuffix, nextIdx); matched2 {
						jumpPos := tm.extractOperand(segment.OpPos[nextIdx2]+1, push2Size)
						pattern = &OutOfBoundCheckPattern{
							StartPosition: uint64(i),
							JumpPosition:  jumpPos,
							Type:          OutOfBound_ConstantVar,
							Index:         operands[0],
							Length:        secondOperand,
							Stack:         []uint64{firstOperand, secondOperand, operands[0]}, // [operand, length storage slot, index]
						}
						break
					}
				}

				matched, _, nextIdx = tm.matchTemplate(segment, []OpCode{PUSH0}, i+2)
				if matched {
					// AllConstant pattern matched
					if matched2, operands2, nextIdx2 := tm.matchTemplate(segment, outOfBound_AllConstantSuffix, nextIdx); matched2 && len(operands2) > 0 {
						jumpPos := tm.extractOperand(segment.OpPos[nextIdx2]+1, push2Size)
						pattern = &OutOfBoundCheckPattern{
							StartPosition: uint64(i),
							JumpPosition:  jumpPos,
							Type:          OutOfBound_AllConstant,
							Index:         uint64(0),
							Length:        operands2[0],
							Stack:         []uint64{firstOperand, secondOperand, uint64(0)}, // [operand, first storage slot, index]
						}
						break
					}

					// ConstantVar pattern matched
					if matched2, _, nextIdx2 := tm.matchTemplate(segment, outOfBound_ConstantVarSuffix, nextIdx); matched2 {
						jumpPos := tm.extractOperand(segment.OpPos[nextIdx2]+1, push2Size)
						pattern = &OutOfBoundCheckPattern{
							StartPosition: uint64(i),
							JumpPosition:  jumpPos,
							Type:          OutOfBound_ConstantVar,
							Index:         uint64(0),
							Length:        secondOperand,
							Stack:         []uint64{firstOperand, secondOperand, uint64(0)}, // [operand, length storage slot, index]
						}
						break
					}
				}
			}

			if segment.Ops[i+1] == DUP1 {
				secondOperand = firstOperand
				matched, operands, nextIdx := tm.matchTemplate(segment, []OpCode{PUSH1}, i+2)
				if matched && len(operands) > 0 {
					// AllConstant pattern matched
					if matched2, operands2, nextIdx2 := tm.matchTemplate(segment, outOfBound_AllConstantSuffix, nextIdx); matched2 && len(operands2) > 0 {
						jumpPos := tm.extractOperand(segment.OpPos[nextIdx2]+1, push2Size)
						pattern = &OutOfBoundCheckPattern{
							StartPosition: uint64(i),
							JumpPosition:  jumpPos,
							Type:          OutOfBound_AllConstant,
							Index:         operands[0],
							Length:        operands2[0],
							Stack:         []uint64{firstOperand, secondOperand, operands[0]}, // [operand, first storage slot, index]
						}
						break
					}
				}

				matched, _, nextIdx = tm.matchTemplate(segment, []OpCode{PUSH0}, i+2)
				if matched {
					// AllConstant pattern matched
					if matched2, operands2, nextIdx2 := tm.matchTemplate(segment, outOfBound_AllConstantSuffix, nextIdx); matched2 && len(operands2) > 0 {
						jumpPos := tm.extractOperand(segment.OpPos[nextIdx2]+1, push2Size)
						pattern = &OutOfBoundCheckPattern{
							StartPosition: uint64(i),
							JumpPosition:  jumpPos,
							Type:          OutOfBound_AllConstant,
							Index:         uint64(0),
							Length:        operands2[0],
							Stack:         []uint64{firstOperand, secondOperand, uint64(0)}, // [operand, first storage slot, index]
						}
						break
					}
				}
			}
		}
		i++
	}
	return pattern
}

// matchOutOfBoundPatternInLoop checks if a loop segment contains out-of-bounds check patterns
func (tm *TemplateMatcher) matchOutOfBoundPatternInLoop(segment CodeSegment) *OutOfBoundCheckPattern {
	// Iterate through ops
	var pattern *OutOfBoundCheckPattern
	i := 0
	push1Size := 1
	push2Size := 2
	for i < len(segment.Ops) {
		if segment.Ops[i] == PUSH1 || segment.Ops[i] == PUSH0 {
			var firstOperand uint64
			if segment.Ops[i] == PUSH0 {
				firstOperand = 0
			} else {
				firstOperand = tm.extractOperand(segment.OpPos[i]+1, push1Size)
			}

			matched, operands, nextIdx := tm.matchTemplate(segment, []OpCode{PUSH0}, i+1)
			if matched && len(operands) > 0 {
				secondOperand := uint64(0)
				// Constant pattern matched
				if matched2, operands2, nextIdx2 := tm.matchTemplate(segment, outOfBound_LoopConstantSuffix, nextIdx); matched2 && len(operands2) > 0 {
					jumpPos := tm.extractOperand(segment.OpPos[nextIdx2]+1, push2Size)
					pattern = &OutOfBoundCheckPattern{
						StartPosition: uint64(i),
						JumpPosition:  jumpPos,
						Type:          OutOfBound_LoopConstant,
						Index:         0, // check in the first loop execution (usually the current first element in the stack)
						Length:        operands2[0],
						Stack:         []uint64{firstOperand, secondOperand, 0}, // [operand, first storage slot, index]
					}
					break
				}

				// Variable pattern matched
				if matched2, _, nextIdx2 := tm.matchTemplate(segment, outOfBound_LoopVarSuffix, nextIdx); matched2 {
					jumpPos := tm.extractOperand(segment.OpPos[nextIdx2]+1, push2Size)
					pattern = &OutOfBoundCheckPattern{
						StartPosition: uint64(i),
						JumpPosition:  jumpPos,
						Type:          OutOfBound_LoopVar,
						Index:         0, // check in the first loop execution (usually the current first element in the stack)
						Length:        secondOperand,
						Stack:         []uint64{firstOperand, secondOperand, 0}, // [operand, length storage slot, index]
					}
					break
				}
			}

			matched, operands, nextIdx = tm.matchTemplate(segment, []OpCode{PUSH1}, i+1)
			if matched && len(operands) > 0 {
				secondOperand := operands[0]
				// Constant pattern matched
				if matched2, operands2, nextIdx2 := tm.matchTemplate(segment, outOfBound_LoopConstantSuffix, nextIdx); matched2 && len(operands2) > 0 {
					jumpPos := tm.extractOperand(segment.OpPos[nextIdx2]+1, push2Size)
					pattern = &OutOfBoundCheckPattern{
						StartPosition: uint64(i),
						JumpPosition:  jumpPos,
						Type:          OutOfBound_LoopConstant,
						Index:         0, // check in the first loop execution (usually the current first element in the stack)
						Length:        operands2[0],
						Stack:         []uint64{firstOperand, secondOperand, 0}, // [operand, first storage slot, index]
					}
					break
				}

				// Variable pattern matched
				if matched2, _, nextIdx2 := tm.matchTemplate(segment, outOfBound_LoopVarSuffix, nextIdx); matched2 {
					jumpPos := tm.extractOperand(segment.OpPos[nextIdx2]+1, push2Size)
					pattern = &OutOfBoundCheckPattern{
						StartPosition: uint64(i),
						JumpPosition:  jumpPos,
						Type:          OutOfBound_LoopVar,
						Index:         0, // check in the first loop execution (usually the current first element in the stack)
						Length:        secondOperand,
						Stack:         []uint64{firstOperand, secondOperand, 0}, // [operand, length storage slot, index]
					}
					break
				}
			}

			matched, operands, nextIdx = tm.matchTemplate(segment, outOfBound_LoopConstantSingleSuffix, i+1)
			// ConstantSingle pattern matched
			if matched && len(operands) > 0 {
				secondOperand := operands[0]
				jumpPos := tm.extractOperand(segment.OpPos[nextIdx]+1, push2Size)
				pattern = &OutOfBoundCheckPattern{
					StartPosition: uint64(i),
					JumpPosition:  jumpPos,
					Type:          OutOfBound_LoopConstantSingle,
					Index:         0, // check in the first loop execution (usually the current first element in the stack)
					Length:        secondOperand,
					Stack:         []uint64{firstOperand, 0}, // [first storage slot, index]
				}
				break
			}

			matched, _, nextIdx = tm.matchTemplate(segment, outOfBound_LoopVarSingleSuffix, i+1)
			// VarSingle pattern matched
			if matched {
				jumpPos := tm.extractOperand(segment.OpPos[nextIdx]+1, push2Size)
				pattern = &OutOfBoundCheckPattern{
					StartPosition: uint64(i),
					JumpPosition:  jumpPos,
					Type:          OutOfBound_LoopVarSingle,
					Index:         0, // check in the first loop execution (usually the current first element in the stack)
					Length:        firstOperand,
					Stack:         []uint64{firstOperand, 0}, // [first length slot, index]
				}
				break
			}
		}
		i++
	}
	return pattern
}

// matchLoopPattern attempts to match loop patterns (the comparison between index (i) and loop count (num))
func (tm *TemplateMatcher) matchLoopPattern(startSeg CodeSegment) *LoopInfo {
	// The entry segment for our pattern must start with JUMPDEST.
	if len(startSeg.Ops) == 0 || startSeg.Ops[0] != JUMPDEST {
		return nil
	}

	// Skip JUMPDEST for matching
	startIdx := 1

	matched, operands, nextIdx := tm.matchTemplate(startSeg, []OpCode{PUSH1}, startIdx)
	if matched && len(operands) > 0 {
		if matched2, _, nextIdx2 := tm.matchTemplate(startSeg, []OpCode{SLOAD}, nextIdx); matched2 {
			// i < num (num is from storage)
			if matched3, _, _ := tm.matchTemplate(startSeg, loopForward_Suffix1, nextIdx2); matched3 {
				return &LoopInfo{
					StartPosition: startSeg.Start,
					Type:          LoopType_Forward_Var,
					LoopNum:       operands[0],
				}
			}
			// i <= num (num is from storage)
			if matched3, _, _ := tm.matchTemplate(startSeg, loopForward_Suffix2, nextIdx2); matched3 {
				return &LoopInfo{
					StartPosition: startSeg.Start,
					Type:          LoopType_Forward_Var,
					LoopNum:       operands[0],
				}
			}
		}

		// i < num (num is constant)
		if matched2, _, _ := tm.matchTemplate(startSeg, loopForward_Suffix1, nextIdx); matched2 {
			return &LoopInfo{
				StartPosition: startSeg.Start,
				Type:          LoopType_Forward_Constant,
				LoopNum:       operands[0],
			}
		}

		// i <= num (num is constant)
		if matched2, _, _ := tm.matchTemplate(startSeg, loopForward_Suffix2, nextIdx); matched2 {
			return &LoopInfo{
				StartPosition: startSeg.Start,
				Type:          LoopType_Forward_Constant,
				LoopNum:       operands[0],
			}
		}

		// i>num (num>0)
		if matched2, _, _ := tm.matchTemplate(startSeg, loopBackward_Suffix1, nextIdx); matched2 {
			return &LoopInfo{
				StartPosition: startSeg.Start,
				Type:          LoopType_Backward,
				LoopNum:       0,
			}
		}

		// i>=num (num>=0)
		if matched2, _, _ := tm.matchTemplate(startSeg, loopBackward_Suffix2, nextIdx); matched2 {
			return &LoopInfo{
				StartPosition: startSeg.Start,
				Type:          LoopType_Backward,
				LoopNum:       0,
			}
		}
	}

	matched, _, nextIdx = tm.matchTemplate(startSeg, []OpCode{PUSH0}, startIdx)
	if matched {
		if matched2, _, nextIdx2 := tm.matchTemplate(startSeg, []OpCode{SLOAD}, nextIdx); matched2 {
			// i < num (num is from storage, slot is 0)
			if matched3, _, _ := tm.matchTemplate(startSeg, loopForward_Suffix1, nextIdx2); matched3 {
				return &LoopInfo{
					StartPosition: startSeg.Start,
					Type:          LoopType_Forward_Var,
					LoopNum:       0,
				}
			}
			// i <= num (num is from storage)
			if matched3, _, _ := tm.matchTemplate(startSeg, loopForward_Suffix2, nextIdx2); matched3 {
				return &LoopInfo{
					StartPosition: startSeg.Start,
					Type:          LoopType_Forward_Var,
					LoopNum:       0,
				}
			}
		}

		// i>0
		if matched2, _, _ := tm.matchTemplate(startSeg, loopBackward_Suffix1, nextIdx); matched2 {
			return &LoopInfo{
				StartPosition: startSeg.Start,
				Type:          LoopType_Backward,
				LoopNum:       0,
			}
		}

		// i>=0
		if matched2, _, _ := tm.matchTemplate(startSeg, loopBackward_Suffix2, nextIdx); matched2 {
			return &LoopInfo{
				StartPosition: startSeg.Start,
				Type:          LoopType_Backward,
				LoopNum:       0,
			}
		}
	}

	matched, _, nextIdx = tm.matchTemplate(startSeg, []OpCode{DUP2}, startIdx)
	if matched {
		// i < num (num is constant)
		if matched2, _, _ := tm.matchTemplate(startSeg, loopForward_Suffix1, nextIdx); matched2 {
			return &LoopInfo{
				StartPosition: startSeg.Start,
				Type:          LoopType_Forward_Constant,
				LoopNum:       0,
			}
		}

		// i <= num (num is constant)
		if matched2, _, _ := tm.matchTemplate(startSeg, loopForward_Suffix2, nextIdx); matched2 {
			return &LoopInfo{
				StartPosition: startSeg.Start,
				Type:          LoopType_Forward_Constant,
				LoopNum:       0,
			}
		}
	}
	return nil
}

// matchOverflowPattern attempts to match the complete overflow check pattern
// starting from the initial segment
func (tm *TemplateMatcher) matchOverflowPattern(startSeg CodeSegment) *OverflowCheckPattern {
	// The entry segment for our pattern must start with JUMPDEST.
	if len(startSeg.Ops) == 0 || startSeg.Ops[0] != JUMPDEST {
		return nil
	}

	// Skip JUMPDEST for matching
	startIdx := 1

	// Check inline overflow patterns (ADD/SUB)
	matched, _, _ := tm.matchTemplate(startSeg, overflow_inlineAdd, startIdx)
	if matched {
		return &OverflowCheckPattern{StartPositions: []uint64{startSeg.Start}, Operation: OpAdd, UintSize: 256, IntSize: 0}
	}
	matched, _, _ = tm.matchTemplate(startSeg, overflow_inlineAdd2, startIdx)
	if matched {
		return &OverflowCheckPattern{StartPositions: []uint64{startSeg.Start}, Operation: OpAdd, UintSize: 256, IntSize: 0}
	}
	matched, _, _ = tm.matchTemplate(startSeg, overflow_inlineSub, startIdx)
	if matched {
		return &OverflowCheckPattern{StartPositions: []uint64{startSeg.Start}, Operation: OpSub, UintSize: 256, IntSize: 0}
	}

	// Check overflow templates: initial pattern
	matched, _, _ = tm.matchTemplate(startSeg, []OpCode{PUSH0}, startIdx)
	matched2, _, _ := tm.matchTemplate(startSeg, []OpCode{PUSH1}, startIdx)
	if !matched && !matched2 {
		return nil
	}

	matched, operands, _ := tm.matchTemplate(startSeg, overflow_templateEntrance[1:], startIdx+1)
	if !matched || len(operands) < 2 {
		return nil
	}

	// Extract jump targets: operands[0] is 'last', operands[1] is 'reused'
	targetLast := operands[0]
	targetReuse := operands[1]

	// Check segment at reused segment
	segReuse := tm.findSegmentByPosition(targetReuse)
	if segReuse == nil {
		return nil
	}

	// Check for different reuse patterns (standard uint256/int256 or masked smaller uints/ints)
	matched, size, indicator := tm.matchReuseSegment(*segReuse, startIdx)
	if !matched {
		return nil
	}

	// Check segment at last segment
	segLast := tm.findSegmentByPosition(targetLast)
	if segLast == nil {
		return nil
	}

	matched, operandsLast, _ := tm.matchTemplate(*segLast, overflow_templateLast, startIdx)
	if !matched || len(operandsLast) < 2 {
		return nil
	}

	targetArithmetic := operandsLast[0]

	// Check segment for arithmetic operation
	segArithmetic := tm.findSegmentByPosition(targetArithmetic)
	if segArithmetic == nil {
		return nil
	}

	// --- Optimized Arithmetic Operation Matching using matchTemplate ---
	arithmeticStartIdx := 1 // Skip JUMPDEST
	var isInt bool

	// Match common prefix: [SWAP3, POP, DUP3]
	matched, _, nextIdx := tm.matchTemplate(*segArithmetic, arithmeticPrefixEntrance, arithmeticStartIdx)
	if !matched {
		return nil
	}

	// Branch for MUL, ADD, SUB
	// Check for next part: [DUP3]
	matched, _, finalIdx := tm.matchTemplate(*segArithmetic, arithmeticPrefixMulAddSub, nextIdx)
	if matched && finalIdx < len(segArithmetic.Ops) {
		op := segArithmetic.Ops[finalIdx]
		switch op {
		case MUL:
			if indicator == 1 {
				// Distinguish uint256 between int256 (int256 needs to check for MIN_INT * (-1) case)
				matched, entrance, _ := tm.matchTemplate(*segArithmetic, jumpEntrance, finalIdx+1)
				if !matched {
					return nil
				}
				minIntMulNeg1 := tm.findSegmentByPosition(entrance[0])
				if minIntMulNeg1 == nil {
					return nil
				}
				matched, _, _ = tm.matchTemplate(*minIntMulNeg1, minIntMulNeg1Check, startIdx)
				if matched {
					isInt = true
				} else {
					isInt = false
				}
			} else if indicator == 2 {
				isInt = false
			} else {
				isInt = true
			}

			if isInt {
				return &OverflowCheckPattern{StartPositions: []uint64{startSeg.Start, targetReuse, targetLast, targetArithmetic}, Operation: OpMul, UintSize: 0, IntSize: size}
			} else {
				return &OverflowCheckPattern{StartPositions: []uint64{startSeg.Start, targetReuse, targetLast, targetArithmetic}, Operation: OpMul, UintSize: size, IntSize: 0}
			}
		case ADD, SUB:
			opType := OpAdd
			if op == SUB {
				opType = OpSub
			}
			if indicator == 1 {
				// Distinguish uint256 between int256 ([ADD/SUB, SWAP1, POP, DUPx, DUPx, cmpOp])
				cmpOp := segArithmetic.Ops[finalIdx+5]
				switch cmpOp {
				case GT, LT:
					isInt = false
				case SLT, SGT:
					isInt = true
				default:
					return nil
				}
			} else if indicator == 2 {
				isInt = false
			} else {
				isInt = true
			}

			if isInt {
				return &OverflowCheckPattern{StartPositions: []uint64{startSeg.Start, targetReuse, targetLast, targetArithmetic}, Operation: opType, UintSize: 0, IntSize: size}
			} else {
				return &OverflowCheckPattern{StartPositions: []uint64{startSeg.Start, targetReuse, targetLast, targetArithmetic}, Operation: opType, UintSize: size, IntSize: 0}
			}
		}
	}

	// Branch for DIV, MOD
	// Check for next part: [PUSH2 xx, JUMPI]
	matched, operandsDivMod, _ := tm.matchTemplate(*segArithmetic, arithmeticPrefixDivMod, nextIdx)
	if matched && len(operandsDivMod) > 0 {
		target := operandsDivMod[0]
		seg := tm.findSegmentByPosition(target)
		if seg != nil {
			// Match common prefix in DIV/MOD/SDIV/SMOD block:
			matched, _, finalIdx = tm.matchTemplate(*seg, divModPrefix, arithmeticStartIdx)
			if matched && finalIdx < len(seg.Ops) {
				switch seg.Ops[finalIdx] {
				case DIV:
					return &OverflowCheckPattern{StartPositions: []uint64{startSeg.Start, targetReuse, targetLast, targetArithmetic, target}, Operation: OpDiv, UintSize: size, IntSize: 0}
				case MOD:
					return &OverflowCheckPattern{StartPositions: []uint64{startSeg.Start, targetReuse, targetLast, targetArithmetic, target}, Operation: OpMod, UintSize: size, IntSize: 0}
				case SMOD:
					return &OverflowCheckPattern{StartPositions: []uint64{startSeg.Start, targetReuse, targetLast, targetArithmetic, target}, Operation: OpSMod, UintSize: 0, IntSize: size}
				}
			} else {
				// SDIV needs to check MIN_INT / (-1) case
				matched, _, _ = tm.matchTemplate(*seg, minIntDivNeg1Check[2:], arithmeticStartIdx+2)
				if matched {
					return &OverflowCheckPattern{StartPositions: []uint64{startSeg.Start, targetReuse, targetLast, targetArithmetic, target}, Operation: OpSDiv, UintSize: 0, IntSize: size}
				}
			}
		}
	}

	return nil
}

// matchOverflowPattern2 attempts to match an alternative overflow check pattern (suitable for older Solidity versions)
func (tm *TemplateMatcher) matchOverflowPattern2(startSeg CodeSegment) *OverflowCheckPattern {
	// The entry segment for our pattern must start with JUMPDEST.
	if len(startSeg.Ops) == 0 || startSeg.Ops[0] != JUMPDEST {
		return nil
	}

	// Skip JUMPDEST and first three opcodes for matching
	startIdx := 4
	// Check for multiplication overflow patterns
	matched, operands, _ := tm.matchTemplate(startSeg, safemath_MulEntrance[3:], startIdx)
	if matched && len(operands) > 0 {
		jumpOp := operands[0]
		checkSegment := tm.findSegmentByPosition(jumpOp)
		if checkSegment == nil {
			return nil
		}

		matched2, _, _ := tm.matchTemplate(*checkSegment, safemath_MulCheck, 1)
		if matched2 {
			return &OverflowCheckPattern{StartPositions: []uint64{startSeg.Start}, Operation: OpMul, UintSize: 256, IntSize: 0, OldVersion: true}
		}
		matched2, _, _ = tm.matchTemplate(*checkSegment, safemath_MulCheck, 2)
		if matched2 {
			return &OverflowCheckPattern{StartPositions: []uint64{startSeg.Start}, Operation: OpMul, UintSize: 256, IntSize: 0, OldVersion: true}
		}
		return nil
	}

	// Skip JUMPDEST and first two opcodes for matching
	startIdx = 3

	// Try matching different entrance branches (including [DUP3] or [DUP4])
	//matched, _, nextIdx := tm.matchTemplate(startSeg, []OpCode{DUP4}, startIdx)
	//if matched {
	//	matched2, operands, _ := tm.matchTemplate(startSeg, safemath_mulEntrance[1:], nextIdx)
	//	if matched2 && len(operands) > 0 {
	//		//fmt.Println("find mul1")
	//		jumpOp := operands[0]
	//		checkSegment := tm.findSegmentByPosition(jumpOp)
	//		if checkSegment == nil {
	//			return nil
	//		}
	//		// Check for multiplication overflow patterns
	//		matched3, _, _ := tm.matchTemplate(*checkSegment, safemath_uintMulCheck, 1)
	//		if matched3 {
	//			//fmt.Println("success mul")
	//			return &OverflowCheckPattern{StartPositions: []uint64{startSeg.Start, jumpOp}, Operation: OpMul, UintSize: 256, IntSize: 0}
	//		}
	//		matched3, _, _ = tm.matchTemplate(*checkSegment, safemath_intMulCheck, 1)
	//		if matched3 {
	//			return &OverflowCheckPattern{StartPositions: []uint64{startSeg.Start, jumpOp}, Operation: OpMul, UintSize: 0, IntSize: 256}
	//		}
	//	}
	//	return nil
	//}

	matched, _, _ = tm.matchTemplate(startSeg, []OpCode{DUP3}, startIdx)
	if !matched {
		// Check for unsigned subtraction overflow pattern
		matched, operands, _ = tm.matchTemplate(startSeg, safemath_SubEntrance3, startIdx)
		if matched && len(operands) > 0 {
			jumpOp := operands[0]
			checkSegment := tm.findSegmentByPosition(jumpOp)
			if checkSegment == nil {
				return nil
			}
			if len(checkSegment.Ops) > 4 && (checkSegment.Ops[3] == SUB || checkSegment.Ops[4] == SUB) {
				return &OverflowCheckPattern{StartPositions: []uint64{startSeg.Start, jumpOp}, Operation: OpSub, UintSize: 256, IntSize: 0, OldVersion: true}
			}
		}
		return nil
	}

	// Check for unsigned addition overflow pattern
	matched, _, _ = tm.matchTemplate(startSeg, safemath_Add2[1:], startIdx+1)
	if matched {
		return &OverflowCheckPattern{StartPositions: []uint64{startSeg.Start}, Operation: OpAdd, UintSize: 256, IntSize: 0, OldVersion: true}
	}

	// Try matching different entrance branches (including [DUP5], [GT] or [SUB])
	matched, _, nextIdx := tm.matchTemplate(startSeg, []OpCode{DUP5}, startIdx+1)
	if matched {
		// Check for unsigned addition overflow pattern
		matched2, _, _ := tm.matchTemplate(startSeg, safemath_Add[2:], nextIdx)
		if matched2 {
			//if startSeg.Ops[nextIdx2+2] == LT {
			//	//fmt.Println("success add")
			//	return &OverflowCheckPattern{StartPositions: []uint64{startSeg.Start}, Operation: OpAdd, UintSize: 256, IntSize: 0, OldVersion: true}
			//}
			//if startSeg.Ops[nextIdx2+2] == SLT {
			//	return &OverflowCheckPattern{StartPositions: []uint64{startSeg.Start}, Operation: OpAdd, UintSize: 0, IntSize: 256}
			//}
			return &OverflowCheckPattern{StartPositions: []uint64{startSeg.Start}, Operation: OpAdd, UintSize: 256, IntSize: 0, OldVersion: true}
		}
		matched2, _, _ = tm.matchTemplate(startSeg, safemath_Add3[2:], nextIdx)
		if matched2 {
			return &OverflowCheckPattern{StartPositions: []uint64{startSeg.Start}, Operation: OpAdd, UintSize: 256, IntSize: 0, OldVersion: true}
		}
		// Check for signed subtraction overflow pattern
		//matched2, _, _ = tm.matchTemplate(startSeg, safemath_intSub[2:], nextIdx)
		//if matched2 {
		//	return &OverflowCheckPattern{StartPositions: []uint64{startSeg.Start}, Operation: OpSub, UintSize: 0, IntSize: 256}
		//}

		// Check for unsigned division overflow pattern (without dividend>0 check)
		matched2, operands2, _ := tm.matchTemplate(startSeg, safemath_DivEntranceWithoutZeroCheck, nextIdx)
		if matched2 && len(operands2) > 0 {
			jumpOp := operands2[0]
			checkSegment := tm.findSegmentByPosition(jumpOp)
			if checkSegment == nil {
				return nil
			}
			if checkSegment.Ops[1] == DIV {
				//fmt.Println("success div")
				return &OverflowCheckPattern{StartPositions: []uint64{startSeg.Start, jumpOp}, Operation: OpDiv, UintSize: 256, IntSize: 0, OldVersion: true}
			}
		}
		matched2, operands2, _ = tm.matchTemplate(startSeg, safemath_DivEntranceWithoutZeroCheck2, nextIdx)
		if matched2 && len(operands2) > 0 {
			jumpOp := operands2[0]
			checkSegment := tm.findSegmentByPosition(jumpOp)
			if checkSegment == nil {
				return nil
			}
			if checkSegment.Ops[1] == DIV {
				//fmt.Println("success div")
				return &OverflowCheckPattern{StartPositions: []uint64{startSeg.Start, jumpOp}, Operation: OpDiv, UintSize: 256, IntSize: 0, OldVersion: true}
			}
		}

		// Check for unsigned subtraction overflow pattern
		matched2, _, _ = tm.matchTemplate(startSeg, safemath_SubEntrance4[2:], nextIdx)
		if matched2 {
			return &OverflowCheckPattern{StartPositions: []uint64{startSeg.Start}, Operation: OpSub, UintSize: 256, IntSize: 0, OldVersion: true}
		}
		return nil
	}

	matched, _, nextIdx = tm.matchTemplate(startSeg, []OpCode{GT}, startIdx+1)
	if matched {
		// Check for unsigned division overflow pattern (with dividend>0 check)
		matched2, operands2, _ := tm.matchTemplate(startSeg, safemath_DivEntranceWithZeroCheck[2:], nextIdx)
		if matched2 && len(operands2) > 0 {
			jumpOp := operands2[0]
			checkSegment := tm.findSegmentByPosition(jumpOp)
			if checkSegment == nil {
				return nil
			}
			matched3, _, _ := tm.matchTemplate(*checkSegment, safemath_DivModReuse, 2)
			if matched3 {
				//fmt.Println("success div")
				return &OverflowCheckPattern{StartPositions: []uint64{startSeg.Start, jumpOp}, Operation: OpDiv, UintSize: 256, IntSize: 0, OldVersion: true}
			}
		}
		// Check for unsigned subtraction overflow pattern
		matched2, operands2, _ = tm.matchTemplate(startSeg, safemath_SubEntrance[2:], nextIdx)
		if matched2 && len(operands2) > 0 {
			jumpOp := operands2[0]
			checkSegment := tm.findSegmentByPosition(jumpOp)
			if checkSegment == nil {
				return nil
			}
			if len(checkSegment.Ops) > 4 && (checkSegment.Ops[3] == SUB || checkSegment.Ops[4] == SUB) {
				return &OverflowCheckPattern{StartPositions: []uint64{startSeg.Start, jumpOp}, Operation: OpSub, UintSize: 256, IntSize: 0, OldVersion: true}
			}
		}
		matched2, operands2, _ = tm.matchTemplate(startSeg, safemath_SubEntrance2[2:], nextIdx)
		if matched2 && len(operands2) > 0 {
			jumpOp := operands2[0]
			checkSegment := tm.findSegmentByPosition(jumpOp)
			if checkSegment == nil {
				return nil
			}
			if checkSegment.Ops[3] == SUB || checkSegment.Ops[4] == SUB {
				return &OverflowCheckPattern{StartPositions: []uint64{startSeg.Start, jumpOp}, Operation: OpSub, UintSize: 256, IntSize: 0, OldVersion: true}
			}
		}
		return nil
	}

	matched, _, nextIdx = tm.matchTemplate(startSeg, []OpCode{SUB}, startIdx+1)
	if matched {
		matched2, operands2, _ := tm.matchTemplate(startSeg, safemath_ModEntrance[2:], nextIdx)
		if matched2 && len(operands2) > 0 {
			jumpOp := operands2[0]
			checkSegment := tm.findSegmentByPosition(jumpOp)
			if checkSegment == nil {
				return nil
			}
			// Check for signed division overflow pattern
			//matched3, _, _ := tm.matchTemplate(*checkSegment, safemath_intDivCheck, 1)
			//if matched3 {
			//	return &OverflowCheckPattern{StartPositions: []uint64{startSeg.Start, jumpOp}, Operation: OpSDiv, UintSize: 0, IntSize: 256}
			//}
			// Check for unsigned modulus overflow pattern
			matched3, _, _ := tm.matchTemplate(*checkSegment, safemath_DivModReuse, 2)
			if matched3 {
				// Safemath does not provide separate SMOD check for signed integers, we only check for unsigned integers here
				return &OverflowCheckPattern{StartPositions: []uint64{startSeg.Start, jumpOp}, Operation: OpMod, UintSize: 256, IntSize: 0, OldVersion: true}
			}
		}
		return nil
	}
	return nil
}

// matchPanicSegment attempts to match the panic trigger pattern in a segment.
func (tm *TemplateMatcher) matchPanicSegment(startSeg CodeSegment) {
	// The entry segment for our pattern must start with JUMPDEST.
	if len(startSeg.Ops) == 0 || startSeg.Ops[0] != JUMPDEST {
		return
	}

	// Skip JUMPDEST for matching
	startIdx := 1

	matched, operands, _ := tm.matchTemplate(startSeg, panic_trigger, startIdx)
	if !matched || len(operands) < 6 {
		return
	}
	panicCode := operands[2]
	tm.panicTrigger[panicCode] = startSeg.Start
}

// GetPanicPosition returns the position of the panic trigger for the given panic code
func (tm *TemplateMatcher) GetPanicPosition(panicCode uint64) (uint64, bool) {
	pos, exists := tm.panicTrigger[panicCode]
	return pos, exists
}

// BuildRuntime exports the minimal runtime metadata required by the interpreter.
func (tm *TemplateMatcher) BuildRuntime() *TemplateRuntime {
	if tm == nil {
		return nil
	}
	if tm.runtime == nil {
		tm.runtime = &TemplateRuntime{
			contractAddr:  tm.contractAddr,
			runtimeLookup: tm.runtimeLookup,
			smallestPC:    tm.smallestPC,
			panicTrigger:  tm.panicTrigger,
			matcherExist:  tm.matcherExist,
		}
	}
	return tm.runtime
}

// ScanPatterns scans the bytecode for overflow check patterns and panic patterns
func (tm *TemplateMatcher) ScanPatterns() {
	// Parse code segments first
	tm.parseCodeSegments()
	// Iterate through all segments sequentially to avoid goroutine scheduling and lock overhead.
	for i := range tm.segments {
		segment := tm.segments[i]
		if pattern := tm.matchOverflowPattern(segment); pattern != nil {
			tm.buildOverflowLookups(pattern)
		} else if pattern2 := tm.matchOverflowPattern2(segment); pattern2 != nil {
			tm.buildOverflowLookups(pattern2)
		}
		tm.matchPanicSegment(segment)
	}
}

// SegmentMetaForPC returns the runtime metadata of the segment containing the given pc.
func (tm *TemplateMatcher) SegmentMetaForPC(pc uint64) *SegmentRuntimeMeta {
	if pc >= uint64(len(tm.runtimeLookup)) {
		return nil
	}
	return tm.runtimeLookup[pc]
}

// RuntimeLookup returns the jumpdest-start runtime lookup table.
func (tm *TemplateMatcher) RuntimeLookup() []*SegmentRuntimeMeta {
	return tm.runtimeLookup
}

// SmallestPC returns the smallest jumpdest pc that owns overflow runtime metadata.
func (tm *TemplateMatcher) SmallestPC() uint64 {
	return tm.smallestPC
}

// IsOverflowCheckEntry checks if the given PC is the entry point of an overflow check pattern
func (tm *TemplateMatcher) IsOverflowCheckEntry(pc uint64) *OverflowCheckPattern {
	if pc >= uint64(len(tm.runtimeLookup)) {
		return nil
	}
	meta := tm.runtimeLookup[pc]
	if meta == nil || meta.Pattern == nil {
		return nil
	}
	if meta.IsWrapper && !meta.IsAhead {
		return nil
	}
	return meta.Pattern
}

// GetPanicPosition returns the panic trigger position for a runtime panic code.
func (tr *TemplateRuntime) GetPanicPosition(panicCode uint64) (uint64, bool) {
	if tr == nil {
		return 0, false
	}
	pos, exists := tr.panicTrigger[panicCode]
	return pos, exists
}

// ContractAddr returns the code owner address associated with this runtime.
func (tr *TemplateRuntime) ContractAddr() common.Address {
	if tr == nil {
		return common.Address{}
	}
	return tr.contractAddr
}
