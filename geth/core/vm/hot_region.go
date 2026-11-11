package vm

// Structural region analysis shared by offline generation and runtime binding.
// The original builder accepts pure stack/gas regions. The explicit ABI
// pipeline extension also stages ordered word-memory effects and expansion gas.
// Storage, calls, unmodeled control flow and GAS remain boundaries.
import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"reflect"
	"runtime"
	"slices"
	"sort"
	"strings"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/holiman/uint256"
)

const hotRegionInputs = 32

type hotRegionNode struct {
	op      OpCode
	args    []int
	value   uint256.Int
	helper  *arithmeticRegion
	nonzero bool // JUMPI nodes are imported ABI guard obligations, not branches.
}
type hotRegion struct {
	Entry, Exit, Gas                                     uint64
	Inputs, Drop, MaxStack, Instructions, Plain, Helpers int
	nodes                                                []hotRegionNode
	outputs                                              []int
	key                                                  string
	pipeline                                             bool
	decodes, encodes                                     int
}
type hotRegionFunc func(*uint64, *ScopeContext, *hotRegion) bool

var hotRegionGenerated = map[string]hotRegionFunc{}

func hotRegionPure(op OpCode) bool {
	switch op {
	case ADD, SUB, MUL, DIV, SDIV, MOD, SMOD, ADDMOD, MULMOD, SIGNEXTEND,
		LT, GT, SLT, SGT, EQ, ISZERO, AND, OR, XOR, NOT, BYTE, SHL, SHR, SAR,
		CALLDATALOAD, CALLDATASIZE:
		return true
	}
	return false
}
func hotRegionSupported(op OpCode) bool {
	return hotRegionPure(op) || op >= PUSH0 && op <= PUSH32 || op >= DUP1 && op <= SWAP16 || op == POP || op == JUMP || op == JUMPDEST
}

func buildHotRegion(b *boundArithmetic, code []byte, table *JumpTable, start uint64) *hotRegion {
	return buildHotRegionBoundary(b, code, table, start, false)
}

// stopAtHelper is an offline ablation policy, never a runtime branch.
func buildHotRegionBoundary(b *boundArithmetic, code []byte, table *JumpTable, start uint64, stopAtHelper bool) *hotRegion {
	return buildHotRegionLimited(b, code, table, start, stopAtHelper, ^uint64(0))
}

// limit is a shared benchmark endpoint, not an observed branch decision.
func buildHotRegionLimited(b *boundArithmetic, code []byte, table *JumpTable, start uint64, stopAtHelper bool, limit uint64) *hotRegion {
	return buildHotRegionWithABI(b, code, table, start, stopAtHelper, limit, nil)
}

func buildHotRegionWithABI(b *boundArithmetic, code []byte, table *JumpTable, start uint64, stopAtHelper bool, limit uint64, abi map[uint64]*abiRegionPrepared, barriers ...map[uint64]bool) *hotRegion {
	if paperAblationLevel == 2 {
		stopAtHelper = true
	}
	r := &hotRegion{Entry: start}
	r.pipeline = abi != nil
	stack := make([]int, hotRegionInputs, 64)
	for i := range stack {
		stack[i] = i
	}
	known := map[int]uint256.Int{}
	visited := map[uint64]bool{}
	bitmap := codeBitmap(code)
	pc, need, peak := start, 0, 0
	add := func(n hotRegionNode) int {
		id := hotRegionInputs + len(r.nodes)
		r.nodes = append(r.nodes, n)
		return id
	}
	constant := func(w uint256.Int) int { id := add(hotRegionNode{op: PUSH32, value: w}); known[id] = w; return id }
	jumpOK := func(w uint256.Int) bool {
		return w.IsUint64() && w.Uint64() < uint64(len(code)) && code[w.Uint64()] == byte(JUMPDEST) && bitmap.codeSegment(w.Uint64())
	}
	for step := 0; step < 256 && pc < uint64(len(code)) && !visited[pc] && len(r.nodes) < 480; step++ {
		if pc != start && len(barriers) > 0 && barriers[0][pc] {
			break
		}
		if pc == limit {
			break
		}
		visited[pc] = true
		if p := abi[pc]; p != nil && len(stack) >= p.region.Inputs {
			base := len(stack) - p.region.Inputs
			if dest, ok := known[stack[base]]; ok && jumpOK(dest) {
				if outputs, ok := importPipelineABI(r, p, stack[base:], known); ok {
					need = max(need, p.region.Inputs-(len(stack)-hotRegionInputs))
					peak = max(peak, base-hotRegionInputs+p.paths[0].peak)
					stack = append(stack[:base], outputs...)
					r.Gas += p.paths[0].cost
					r.Instructions += int(p.paths[0].instructions)
					if p.native.kind == 1 {
						r.decodes++
					} else {
						r.encodes++
					}
					pc = dest.Uint64()
					continue
				}
			}
		}
		if stopAtHelper && b.regions[pc] != nil {
			break
		}
		if h := b.regions[pc]; h != nil && h.syncHot.native != nil && !h.pathDependentResult && h.constantSuccessGas && len(stack) >= h.inputCount {
			base := len(stack) - h.inputCount
			resolve := func(s stackSource) int {
				if s.index < 0 {
					return constant(s.constant)
				}
				return stack[base+s.index]
			}
			k := resolve(h.k)
			dest, ok := known[k]
			if ok && jumpOK(dest) {
				// Only constant-cost, memory-free successful helper paths. All other
				// outcomes reject the speculative locals and execute original code.
				length := 0
				for i, cost := range h.syncHot.costs {
					if cost != syncHotFallback && cost == h.syncHot.admissionGas {
						length = max(length, len(h.paths[i].path.PCs))
					}
				}
				prefixSteps, err := hotRegionPrefixSteps(h)
				if err != nil || length == 0 {
					break
				}
				a, z := resolve(h.a), resolve(h.b)
				prefix := make([]int, len(h.prefix))
				for i, s := range h.prefix {
					prefix[i] = resolve(s)
				}
				need = max(need, h.inputCount-(len(stack)-hotRegionInputs))
				peak = max(peak, len(stack)-hotRegionInputs-h.inputCount+h.peak)
				id := add(hotRegionNode{args: []int{a, z}, helper: h})
				stack = append(stack[:base], prefix...)
				stack = append(stack, id)
				r.Gas += h.syncHot.admissionGas
				r.Instructions += length + prefixSteps
				r.Helpers++
				pc = dest.Uint64()
				continue
			}
		}
		op := OpCode(code[pc])
		ins := table[op]
		memoryOp := r.pipeline && (op == MLOAD || op == MSTORE)
		if (!hotRegionSupported(op) && !memoryOp) || ins == nil || ins.undefined || (!memoryOp && (ins.dynamicGas != nil || ins.memorySize != nil)) || len(stack) < ins.minStack {
			break
		}
		need = max(need, ins.minStack-(len(stack)-hotRegionInputs))
		peak = max(peak, len(stack)-hotRegionInputs+1024-ins.maxStack)
		next := pc + 1
		switch {
		case op >= PUSH0 && op <= PUSH32:
			n := uint64(op - PUSH0)
			if pc+1+n > uint64(len(code)) {
				return nil
			}
			var w uint256.Int
			w.SetBytes(code[pc+1 : pc+1+n])
			stack = append(stack, constant(w))
			next += n
		case op >= DUP1 && op <= DUP16:
			stack = append(stack, stack[len(stack)-int(op-DUP1)-1])
		case op >= SWAP1 && op <= SWAP16:
			i, j := len(stack)-1, len(stack)-int(op-SWAP1)-2
			stack[i], stack[j] = stack[j], stack[i]
		case op == POP:
			stack = stack[:len(stack)-1]
		case op == JUMPDEST:
		case op == JUMP:
			w, ok := known[stack[len(stack)-1]]
			if !ok || !jumpOK(w) {
				goto done
			}
			next = w.Uint64()
			stack = stack[:len(stack)-1]
		case op == MSTORE:
			add(hotRegionNode{op: MSTORE, args: []int{stack[len(stack)-1], stack[len(stack)-2]}})
			stack = stack[:len(stack)-2]
		default:
			arity := 2
			switch op {
			case NOT, ISZERO, CALLDATALOAD, MLOAD:
				arity = 1
			case CALLDATASIZE:
				arity = 0
			case ADDMOD, MULMOD:
				arity = 3
			}
			args := make([]int, arity)
			for i := range args {
				args[i] = stack[len(stack)-1-i]
			}
			id := add(hotRegionNode{op: op, args: args})
			stack = append(stack[:len(stack)-arity], id)
			r.Plain++
		}
		r.Gas += ins.constantGas
		r.Instructions++
		pc = next
	}
done:
	if pc == start || (!stopAtHelper && (r.Instructions < 16 || r.Plain+r.Helpers < 2)) {
		return nil
	}
	trim := 0
	for trim < len(stack) && trim < hotRegionInputs && stack[trim] == trim {
		trim++
	}
	r.Drop = hotRegionInputs - trim
	r.Inputs = max(need, r.Drop)
	r.MaxStack = 1024 - peak
	r.Exit = pc
	r.outputs = append([]int(nil), stack[trim:]...)
	if r.Inputs > hotRegionInputs || r.MaxStack < r.Inputs {
		return nil
	}
	r.key = r.fingerprint()
	return r
}

func hotRegionPrefixSteps(r *arithmeticRegion) (int, error) {
	// Reuse the already proved entry extension rather than re-mining opcodes.
	if r.entry == r.helper {
		return 0, nil
	}
	for pc, steps := r.entry, 0; pc < uint64(len(r.code)); {
		op := OpCode(r.code[pc])
		pc++
		steps++
		switch {
		case op >= PUSH1 && op <= PUSH32:
			pc += uint64(op-PUSH1) + 1
		case op == PUSH0 || op == POP || op >= DUP1 && op <= SWAP16:
		case op == JUMP:
			return steps, nil
		default:
			return 0, fmt.Errorf("unexpected entry prefix")
		}
	}
	return 0, fmt.Errorf("unterminated entry prefix")
}

func (r *hotRegion) fingerprint() string {
	var s strings.Builder
	if r.pipeline {
		fmt.Fprint(&s, "abi-pipeline-v1;")
	}
	fmt.Fprintf(&s, "prague-v1;%d/%d/%d/%d/%d;", r.Inputs, r.Drop, r.MaxStack, r.Gas, r.Exit)
	for _, n := range r.nodes {
		fmt.Fprintf(&s, "%d:%v:%s;", n.op, n.args, n.value.Hex())
		if r.pipeline && n.op == JUMPI {
			fmt.Fprintf(&s, "guard:%t;", n.nonzero)
		}
		if n.helper != nil {
			h := n.helper
			fmt.Fprintf(&s, "%s:%s:%d:%t:%v;", h.plan.signature(), runtime.FuncForPC(reflect.ValueOf(h.syncHot.native).Pointer()).Name(), h.domain.Bits, h.domain.Signed, h.syncHot.costs)
		}
	}
	fmt.Fprint(&s, r.outputs)
	return fmt.Sprintf("%x", sha256.Sum256([]byte(s.String())))
}

func discoverHotRegions(e *ArithmeticEngine, code []byte) (*boundArithmetic, map[uint64]*hotRegion) {
	b := e.bind(code, arithmeticCacheKey{hash: crypto.Keccak256Hash(code), rules: &pragueInstructionSet, reduced: true})
	starts := map[uint64]bool{0: true}
	for pc := uint64(0); pc < uint64(len(code)); {
		op := OpCode(code[pc])
		if op == JUMPDEST {
			starts[pc] = true
		}
		pc++
		if op >= PUSH1 && op <= PUSH32 {
			pc += uint64(op-PUSH1) + 1
		}
		if !hotRegionSupported(op) {
			starts[pc] = true
		}
	}
	for pc := range b.regions {
		starts[pc] = true
	}
	out := map[uint64]*hotRegion{}
	for pc := range starts {
		if r := buildHotRegion(b, code, &pragueInstructionSet, pc); r != nil {
			out[pc] = r
		}
	}
	return b, out
}

// This interpreter is used ONLY by the offline witness collector. Generated
// functions contain direct operations, never a traversal of these nodes.
func hotRegionValue(op OpCode, args []uint256.Int, input []byte) uint256.Int {
	var z uint256.Int
	switch op {
	case ADD:
		z.Add(&args[0], &args[1])
	case SUB:
		z.Sub(&args[0], &args[1])
	case MUL:
		z.Mul(&args[0], &args[1])
	case DIV:
		z.Div(&args[0], &args[1])
	case SDIV:
		z.SDiv(&args[0], &args[1])
	case MOD:
		z.Mod(&args[0], &args[1])
	case SMOD:
		z.SMod(&args[0], &args[1])
	case ADDMOD:
		z.AddMod(&args[0], &args[1], &args[2])
	case MULMOD:
		z.MulMod(&args[0], &args[1], &args[2])
	case AND:
		z.And(&args[0], &args[1])
	case OR:
		z.Or(&args[0], &args[1])
	case XOR:
		z.Xor(&args[0], &args[1])
	case NOT:
		z.Not(&args[0])
	case EQ:
		if args[0].Eq(&args[1]) {
			z.SetOne()
		}
	case LT:
		if args[0].Lt(&args[1]) {
			z.SetOne()
		}
	case GT:
		if args[0].Gt(&args[1]) {
			z.SetOne()
		}
	case SLT:
		if args[0].Slt(&args[1]) {
			z.SetOne()
		}
	case SGT:
		if args[0].Sgt(&args[1]) {
			z.SetOne()
		}
	case ISZERO:
		if args[0].IsZero() {
			z.SetOne()
		}
	case SHL:
		if args[0].LtUint64(256) {
			z.Lsh(&args[1], uint(args[0].Uint64()))
		}
	case SHR:
		if args[0].LtUint64(256) {
			z.Rsh(&args[1], uint(args[0].Uint64()))
		}
	case SAR:
		shift := uint(256)
		if args[0].LtUint64(256) {
			shift = uint(args[0].Uint64())
		}
		z.SRsh(&args[1], shift)
	case BYTE:
		z = args[1]
		z.Byte(&args[0])
	case SIGNEXTEND:
		z.ExtendSign(&args[1], &args[0])
	case CALLDATASIZE:
		z.SetUint64(uint64(len(input)))
	case CALLDATALOAD:
		if args[0].IsUint64() {
			offset := args[0].Uint64()
			if offset < uint64(len(input)) {
				var word [32]byte
				copy(word[:], input[offset:min(offset+32, uint64(len(input)))])
				z.SetBytes(word[:])
			}
		}
	default:
		panic("unsupported hot region opcode")
	}
	return z
}
func (r *hotRegion) evaluate(words []uint256.Int, input []byte, gas uint64) ([]uint256.Int, int, bool) {
	// The pure witness evaluator has no memory input. Effectful pipeline regions
	// use their generated executor and the full opcode differential runner.
	if r.pipeline {
		return nil, 0, false
	}
	if len(words) < r.Inputs || len(words) > r.MaxStack || gas < r.Gas {
		return nil, 0, false
	}
	steps := r.Instructions
	values := make([]uint256.Int, hotRegionInputs+len(r.nodes))
	copy(values[hotRegionInputs-r.Inputs:hotRegionInputs], words[len(words)-r.Inputs:])
	for i, n := range r.nodes {
		var z uint256.Int
		if h := n.helper; h != nil {
			a, b := values[n.args[0]], values[n.args[1]]
			if !h.domain.canonical(a) || !h.domain.canonical(b) {
				return nil, 0, false
			}
			p, v := h.syncHot.native(a, b)
			if p < 0 || p >= len(h.syncHot.costs) || h.syncHot.costs[p] == syncHotFallback || h.syncHot.costs[p] != h.syncHot.admissionGas {
				return nil, 0, false
			}
			maxLength := 0
			for i, cost := range h.syncHot.costs {
				if cost != syncHotFallback && cost == h.syncHot.admissionGas {
					maxLength = max(maxLength, len(h.paths[i].path.PCs))
				}
			}
			steps += len(h.paths[p].path.PCs) - maxLength
			z = v
		} else if n.op == PUSH32 {
			z = n.value
		} else {
			var args []uint256.Int
			for _, a := range n.args {
				args = append(args, values[a])
			}
			z = hotRegionValue(n.op, args, input)
		}
		values[hotRegionInputs+i] = z
	}
	out := append([]uint256.Int(nil), words[:len(words)-r.Drop]...)
	for _, id := range r.outputs {
		out = append(out, values[id])
	}
	return out, steps, true
}

type HotRegionSample struct {
	Stack []uint256.Int
	Input []byte
}
type HotRegionEvidence struct {
	StopAtHelper                         bool `json:",omitempty"`
	CodeHash, Key                        string
	Entry, Exit, Gas                     uint64
	Instructions, Plain, Helpers, Inputs int
	Visits, Rejected                     uint64
	Samples                              []HotRegionSample
}
type HotRegionStudy struct {
	Engine                                               *ArithmeticEngine
	Codes                                                map[string]string
	Evidence                                             map[string]*HotRegionEvidence
	Instructions, ConfirmedInstructions, ConfirmedVisits uint64
	bindings                                             map[string]map[uint64]*hotRegion
	Err                                                  error
}

func NewHotRegionStudy(e *ArithmeticEngine) *HotRegionStudy {
	return &HotRegionStudy{Engine: e, Codes: map[string]string{}, Evidence: map[string]*HotRegionEvidence{}, bindings: map[string]map[uint64]*hotRegion{}}
}

type HotRegionWitness struct {
	study    *HotRegionStudy
	hash     string
	regions  map[uint64]*hotRegion
	pending  *hotRegion
	expected []uint256.Int
	gas      uint64
	steps    int
	limit    int
	sample   HotRegionSample
}

func (s *HotRegionStudy) Frame(code []byte) *HotRegionWitness {
	hash := crypto.Keccak256Hash(code).Hex()
	regions, ok := s.bindings[hash]
	if !ok {
		_, regions = discoverHotRegions(s.Engine, code)
		s.bindings[hash] = regions
		s.Codes[hash] = hex.EncodeToString(code)
	}
	return &HotRegionWitness{study: s, hash: hash, regions: regions}
}
func (w *HotRegionWitness) Observe(pc, gas uint64, words []uint256.Int, input []byte, err error) {
	s := w.study
	s.Instructions++
	if s.Err != nil {
		return
	}
	if r := w.pending; r != nil {
		w.steps++
		if w.steps == w.limit {
			if pc != r.Exit || err != nil || gas != w.gas || !slices.Equal(words, w.expected) {
				s.Err = fmt.Errorf("region witness mismatch %s pc%d exit%d steps%d gas%d/%d", w.hash, r.Entry, pc, w.steps, gas, w.gas)
				return
			}
			key := fmt.Sprintf("%s/%d", w.hash, r.Entry)
			d := s.Evidence[key]
			d.Visits++
			s.ConfirmedVisits++
			s.ConfirmedInstructions += uint64(w.steps)
			if len(d.Samples) < 3 && len(w.sample.Input) <= 16384 {
				d.Samples = append(d.Samples, w.sample)
			}
			w.pending = nil
		} else {
			if w.steps > r.Instructions+256 {
				s.Err = fmt.Errorf("region did not reach exit %s/%d", w.hash, r.Entry)
			}
			return
		}
	}
	if err != nil {
		return
	}
	r := w.regions[pc]
	if r == nil {
		return
	}
	key := fmt.Sprintf("%s/%d", w.hash, r.Entry)
	d := s.Evidence[key]
	if d == nil {
		d = &HotRegionEvidence{CodeHash: w.hash, Key: r.key, Entry: r.Entry, Exit: r.Exit, Gas: r.Gas, Instructions: r.Instructions, Plain: r.Plain, Helpers: r.Helpers, Inputs: r.Inputs}
		s.Evidence[key] = d
	}
	out, steps, ok := r.evaluate(words, input, gas)
	if !ok {
		d.Rejected++
		return
	}
	w.pending = r
	w.expected = out
	w.gas = gas - r.Gas
	w.steps = 0
	w.limit = steps
	w.sample = HotRegionSample{append([]uint256.Int(nil), words...), append([]byte(nil), input...)}
}
func (w *HotRegionWitness) Finish() {
	if w.pending != nil && w.study.Err == nil {
		w.study.Err = fmt.Errorf("unclosed pure region %s/%d", w.hash, w.pending.Entry)
	}
}
func (s *HotRegionStudy) Ranked() []*HotRegionEvidence {
	var out []*HotRegionEvidence
	for _, d := range s.Evidence {
		if d.Visits > 0 {
			out = append(out, d)
		}
	}
	sort.Slice(out, func(i, j int) bool {
		a, b := out[i], out[j]
		if a.Visits*uint64(a.Instructions) != b.Visits*uint64(b.Instructions) {
			return a.Visits*uint64(a.Instructions) > b.Visits*uint64(b.Instructions)
		}
		return a.CodeHash < b.CodeHash || a.CodeHash == b.CodeHash && a.Entry < b.Entry
	})
	return out
}
