package vm

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/ethereum/go-ethereum/params"
	"github.com/holiman/uint256"
	miner "go-eth1.16.5-evm/template_miner"
)

// Ordered native operations contain no DUP/SWAP/POP or internal helper jumps.
// Effects remain speculative until every original guard and gas obligation holds.
const (
	abiPlanConst = 256 + iota
	abiPlanInput
	abiPlanGuard
)

type abiPlanStep struct {
	op           int
	dst, a, b, c int
	value        uint256.Int
	nonzero      bool
}
type abiRegionPlan struct {
	steps              []abiPlanStep
	outputs            []int
	cost, instructions uint64
	peak, slots        int
}
type abiRegionPrepared struct {
	native *abiNativeRegion
	region *miner.ABIRegion
	paths  []abiRegionPlan
	graph  []abiGraphPlan
}
type abiGraphPlan struct {
	plan                   abiRegionPlan
	next, other, condition int
	outcome                string
}
type abiMemoryWrite struct {
	offset, length, source uint64
	word                   [32]byte
	calldata               bool
}
type abiRegionScratch struct {
	values  [512]uint256.Int
	writes  [256]abiMemoryWrite
	virtual [2][128]uint256.Int
}
type abiRegionTxn struct {
	memSize, memFee, cost, instructions uint64
	nw                                  int
	signature                           string
}

func LoadABIRegionDiagnostic(scalars, regions string) (*ABIScalarDiagnostic, error) {
	e, err := LoadABIScalarDiagnostic(scalars)
	if err != nil {
		return nil, err
	}
	b, err := os.ReadFile(regions)
	if err != nil {
		return nil, err
	}
	var l miner.ABIRegionLibrary
	if err = json.Unmarshal(b, &l); err != nil {
		return nil, err
	}
	if l.Version != 1 {
		return nil, fmt.Errorf("unsupported ABI region library")
	}
	e.Regions = &l
	return e, nil
}
func prepareABIRegion(r *miner.ABIRegion, code []byte, table *JumpTable) (*abiRegionPrepared, error) {
	p := &abiRegionPrepared{region: r}
	if len(r.Graph) > 0 {
		for _, block := range r.Graph {
			b := abiGraphPlan{next: block.Next, other: block.Else, condition: -1, outcome: block.Path.Outcome}
			if b.outcome != "revert" {
				path := block.Path
				path.Outcome = "return"
				path.Outputs = append([]int(nil), path.Outputs...)
				if block.Condition >= 0 {
					path.Outputs = append(path.Outputs, block.Condition)
				}
				compiled, err := prepareABIRegion(&miner.ABIRegion{Paths: []miner.ABIRegionPath{path}}, code, table)
				if err != nil {
					return nil, err
				}
				b.plan = compiled.paths[0]
				if block.Condition >= 0 {
					b.condition = b.plan.outputs[len(b.plan.outputs)-1]
					b.plan.outputs = b.plan.outputs[:len(b.plan.outputs)-1]
				}
			}
			p.graph = append(p.graph, b)
		}
		return p, nil
	}
	for _, path := range r.Paths {
		if path.Outcome != "return" {
			continue
		}
		plan := abiRegionPlan{peak: path.Peak, instructions: uint64(len(path.PCs))}
		for _, pc := range path.PCs {
			op := OpCode(code[pc])
			o := table[op]
			if o == nil || o.undefined {
				return nil, fmt.Errorf("undefined opcode")
			}
			if o.dynamicGas != nil && op != MLOAD && op != MSTORE && op != MSTORE8 && op != CALLDATACOPY {
				return nil, fmt.Errorf("unmodeled dynamic gas %s", op)
			}
			plan.cost += o.constantGas
		}
		ids := map[int]int{}
		var compile func(int) (int, error)
		alloc := func(id int) int { n := plan.slots; plan.slots++; ids[id] = n; return n }
		compile = func(id int) (int, error) {
			if n, ok := ids[id]; ok {
				return n, nil
			}
			if id < 0 || id >= len(path.Nodes) {
				return 0, fmt.Errorf("node index")
			}
			v := path.Nodes[id]
			step := abiPlanStep{}
			switch v.Op {
			case "input":
				step.op = abiPlanInput
				step.a = v.Input
			case "const":
				step.op = abiPlanConst
				if err := step.value.SetFromHex(v.Value); err != nil {
					return 0, err
				}
			case "read":
				return 0, fmt.Errorf("read before definition")
			default:
				op := StringToOp(v.Op)
				if op == STOP {
					return 0, fmt.Errorf("unknown expression %s", v.Op)
				}
				step.op = int(op)
				for i, arg := range v.Args {
					n, e := compile(arg)
					if e != nil {
						return 0, e
					}
					switch i {
					case 0:
						step.a = n
					case 1:
						step.b = n
					default:
						return 0, fmt.Errorf("arity")
					}
				}
			}
			n := alloc(id)
			step.dst = n
			plan.steps = append(plan.steps, step)
			return n, nil
		}
		for _, action := range path.Actions {
			step := abiPlanStep{op: int(StringToOp(action.Op)), nonzero: action.Nonzero}
			if action.Op == "guard" {
				step.op = abiPlanGuard
			}
			for i, arg := range action.Args {
				n, e := compile(arg)
				if e != nil {
					return nil, e
				}
				switch i {
				case 0:
					step.a = n
				case 1:
					step.b = n
				case 2:
					step.c = n
				}
			}
			switch action.Op {
			case "CALLDATALOAD", "MLOAD", "MSIZE", "CALLDATASIZE":
				step.dst = alloc(action.Result)
			}
			plan.steps = append(plan.steps, step)
		}
		for _, output := range path.Outputs {
			n, e := compile(output)
			if e != nil {
				return nil, e
			}
			plan.outputs = append(plan.outputs, n)
		}
		if plan.slots > 512 {
			return nil, fmt.Errorf("scratch bound")
		}
		p.paths = append(p.paths, plan)
	}
	if len(p.paths) == 0 {
		return nil, fmt.Errorf("no success")
	}
	return p, nil
}

// tryABIRegion never changes live state on rejection, including memory size,
// lastGasCost and the repository's calldata Signature extension.
func (e *ABIScalarDiagnostic) tryABIRegion(p *abiScalarPrepared, pc *uint64, scope *ScopeContext) bool {
	r := p.full.region
	n := len(scope.Stack.data)
	if n < r.Inputs {
		return false
	}
	prefix := n - r.Inputs
	ret := scope.Stack.data[prefix]
	if !scope.Contract.validJumpdest(&ret) {
		return false
	}
	if p.full.native != nil {
		instructions, ok := e.runABINative(p.full.native, prefix, scope)
		if !ok {
			return false
		}
		*pc = ret.Uint64()
		e.Stats.Skipped++
		e.Stats.Instructions += instructions
		if e.Capture {
			e.Visits[p.label]++
			e.NativeVisits[p.label+"/"+strconv.Itoa(int(p.full.native.kind))]++
			path := "0"
			if len(p.full.graph) > 0 {
				path = "graph"
			}
			e.RegionVisits[p.label+"/"+r.Category+"/"+path]++
		}
		return true
	}
	if len(p.full.graph) > 0 {
		if !e.runABIGraph(p.full, prefix, scope) {
			return false
		}
		*pc = ret.Uint64()
		e.Stats.Skipped++
		if e.Capture {
			e.Visits[p.label]++
			e.RegionVisits[p.label+"/"+r.Category+"/graph"]++
		}
		return true
	}
	for pi := range p.full.paths {
		plan := &p.full.paths[pi]
		if prefix+plan.peak > 1024 || scope.Contract.Gas < plan.cost {
			continue
		}
		if e.runABIPlan(plan, r.Inputs, prefix, scope) {
			*pc = ret.Uint64()
			e.Stats.Skipped++
			e.Stats.Instructions += plan.instructions
			if e.Capture {
				e.Visits[p.label]++
				e.RegionVisits[p.label+"/"+r.Category+"/"+strconv.Itoa(pi)]++
			}
			return true
		}
	}
	return false
}
func abiLoadCalldata(dst *uint256.Int, input []byte, offset *uint256.Int) {
	o, overflow := offset.Uint64WithOverflow()
	dst.Clear()
	if overflow || o >= uint64(len(input)) {
		return
	}
	if uint64(len(input))-o >= 32 {
		dst.SetBytes(input[o : o+32])
		return
	}
	var b [32]byte
	copy(b[:], input[o:])
	dst.SetBytes(b[:])
}
func (e *ABIScalarDiagnostic) runABIPlan(p *abiRegionPlan, inputs, prefix int, scope *ScopeContext) bool {
	t := abiRegionTxn{memSize: uint64(scope.Memory.Len()), memFee: scope.Memory.lastGasCost, signature: scope.Signature}
	if !e.evalABIPlan(p, scope.Stack.data[prefix:], &t, scope) {
		return false
	}
	e.commitABIPlan(p.outputs, prefix, &t, scope)
	return true
}
func (e *ABIScalarDiagnostic) runABIGraph(p *abiRegionPrepared, prefix int, scope *ScopeContext) bool {
	t := abiRegionTxn{memSize: uint64(scope.Memory.Len()), memFee: scope.Memory.lastGasCost, signature: scope.Signature}
	side := 0
	stack := e.scratch.virtual[side][:p.region.Inputs]
	copy(stack, scope.Stack.data[prefix:])
	bi := 0
	// A bounded attempt avoids monopolizing the host on very large or malformed
	// loops. Exhaustion leaves the original frame untouched and resumes the EVM.
	for steps := 0; steps < 8192; steps++ {
		if bi < 0 || bi >= len(p.graph) {
			return false
		}
		b := &p.graph[bi]
		if b.outcome == "revert" || prefix+b.plan.peak > 1024 {
			return false
		}
		if !e.evalABIPlan(&b.plan, stack, &t, scope) {
			return false
		}
		if b.outcome == "return" {
			e.commitABIPlan(b.plan.outputs, prefix, &t, scope)
			e.Stats.Instructions += t.instructions
			return true
		}
		bi = b.next
		if b.condition >= 0 && e.scratch.values[b.condition].IsZero() {
			bi = b.other
		}
		if len(b.plan.outputs) > 128 {
			return false
		}
		side = 1 - side
		stack = e.scratch.virtual[side][:len(b.plan.outputs)]
		for i, id := range b.plan.outputs {
			stack[i] = e.scratch.values[id]
		}
	}
	return false
}
func (e *ABIScalarDiagnostic) evalABIPlan(p *abiRegionPlan, input []uint256.Int, t *abiRegionTxn, scope *ScopeContext) bool {
	values := e.scratch.values[:]
	writes := e.scratch.writes[:]
	if p.cost > scope.Contract.Gas-t.cost {
		return false
	}
	nw := t.nw
	memSize := t.memSize
	memFee := t.memFee
	cost := t.cost + p.cost
	signature := t.signature
	expand := func(offset, size uint64) bool {
		if size == 0 {
			return true
		}
		if offset > 0x1FFFFFFFE0 || size > 0x1FFFFFFFE0-offset {
			return false
		}
		end := (offset + size + 31) &^ uint64(31)
		if end > memSize {
			words := end / 32
			fee := words*params.MemoryGas + words*words/params.QuadCoeffDiv
			if fee < memFee || fee-memFee > scope.Contract.Gas-cost {
				return false
			}
			cost += fee - memFee
			memFee = fee
			memSize = end
		}
		return true
	}
	for _, s := range p.steps {
		z := &values[s.dst]
		a, b := &values[s.a], &values[s.b]
		switch s.op {
		case abiPlanConst:
			*z = s.value
		case abiPlanInput:
			if s.a >= len(input) {
				return false
			}
			*z = input[s.a]
		case abiPlanGuard:
			if (!a.IsZero()) != s.nonzero {
				return false
			}
		case int(ADD):
			z.Add(a, b)
		case int(SUB):
			z.Sub(a, b)
		case int(MUL):
			z.Mul(a, b)
		case int(DIV):
			z.Div(a, b)
		case int(MOD):
			z.Mod(a, b)
		case int(SDIV):
			z.SDiv(a, b)
		case int(SMOD):
			z.SMod(a, b)
		case int(AND):
			z.And(a, b)
		case int(OR):
			z.Or(a, b)
		case int(XOR):
			z.Xor(a, b)
		case int(NOT):
			z.Not(a)
		case int(ISZERO):
			z.Clear()
			if a.IsZero() {
				z.SetOne()
			}
		case int(EQ):
			z.Clear()
			if a.Eq(b) {
				z.SetOne()
			}
		case int(LT):
			z.Clear()
			if a.Lt(b) {
				z.SetOne()
			}
		case int(GT):
			z.Clear()
			if a.Gt(b) {
				z.SetOne()
			}
		case int(SLT):
			z.Clear()
			if a.Slt(b) {
				z.SetOne()
			}
		case int(SGT):
			z.Clear()
			if a.Sgt(b) {
				z.SetOne()
			}
		case int(SIGNEXTEND):
			z.ExtendSign(b, a)
		case int(BYTE):
			*z = *b
			z.Byte(a)
		case int(SHL), int(SHR), int(SAR):
			shift, overflow := a.Uint64WithOverflow()
			if overflow || shift >= 256 {
				shift = 256
			}
			if s.op == int(SHL) {
				z.Lsh(b, uint(shift))
			} else if s.op == int(SHR) {
				z.Rsh(b, uint(shift))
			} else {
				z.SRsh(b, uint(shift))
			}
		case int(CALLDATASIZE):
			z.SetUint64(uint64(len(scope.Contract.Input)))
		case int(MSIZE):
			z.SetUint64(memSize)
		case int(CALLDATALOAD):
			abiLoadCalldata(z, scope.Contract.Input, a)
			if a.IsZero() {
				word := z.Bytes32()
				signature = string(word[:4])
			}
		case int(MLOAD):
			o, overflow := a.Uint64WithOverflow()
			if overflow || !expand(o, 32) {
				return false
			}
			var word [32]byte
			if o < uint64(len(scope.Memory.store)) {
				copy(word[:], scope.Memory.store[o:])
			}
			for i := 0; i < nw; i++ {
				w := &writes[i]
				lo := max(o, w.offset)
				hi := min(o+32, w.offset+w.length)
				if hi <= lo {
					continue
				}
				if w.calldata {
					for j := lo; j < hi; j++ {
						src := w.source + (j - w.offset)
						if src >= w.source && src < uint64(len(scope.Contract.Input)) {
							word[j-o] = scope.Contract.Input[src]
						} else {
							word[j-o] = 0
						}
					}
				} else {
					copy(word[lo-o:hi-o], w.word[lo-w.offset:hi-w.offset])
				}
			}
			z.SetBytes(word[:])
		case int(MSTORE), int(MSTORE8):
			o, overflow := a.Uint64WithOverflow()
			size := uint64(32)
			if s.op == int(MSTORE8) {
				size = 1
			}
			if overflow || !expand(o, size) || nw >= len(writes) {
				return false
			}
			w := abiMemoryWrite{offset: o, length: size}
			if size == 1 {
				w.word[0] = byte(b.Uint64())
			} else {
				w.word = b.Bytes32()
			}
			writes[nw] = w
			nw++
		case int(CALLDATACOPY):
			length, ov := values[s.c].Uint64WithOverflow()
			if ov {
				return false
			}
			if length == 0 {
				continue
			}
			off, ov := a.Uint64WithOverflow()
			if ov || !expand(off, length) || nw >= len(writes) {
				return false
			}
			words := (length + 31) / 32
			fee := words * params.CopyGas
			if fee > scope.Contract.Gas-cost {
				return false
			}
			cost += fee
			src, ov := b.Uint64WithOverflow()
			if ov {
				src = ^uint64(0)
			}
			writes[nw] = abiMemoryWrite{offset: off, length: length, source: src, calldata: true}
			nw++
		default:
			return false
		}
	}
	t.nw, t.memSize, t.memFee, t.cost, t.signature = nw, memSize, memFee, cost, signature
	t.instructions += p.instructions
	return true
}
func (e *ABIScalarDiagnostic) commitABIPlan(outputs []int, prefix int, t *abiRegionTxn, scope *ScopeContext) {
	// Original successful path is now fully admitted. Commit effects in order.
	scope.Memory.Resize(t.memSize)
	scope.Memory.lastGasCost = t.memFee
	for i := 0; i < t.nw; i++ {
		w := &e.scratch.writes[i]
		dst := scope.Memory.store[w.offset : w.offset+w.length]
		if !w.calldata {
			copy(dst, w.word[:w.length])
			continue
		}
		copied := 0
		if w.source < uint64(len(scope.Contract.Input)) {
			copied = copy(dst, scope.Contract.Input[w.source:])
		}
		clear(dst[copied:])
	}
	stack := scope.Stack.data[:prefix]
	for _, i := range outputs {
		stack = append(stack, e.scratch.values[i])
	}
	scope.Stack.data = stack
	scope.Contract.Gas -= t.cost
	scope.Signature = t.signature
	e.Stats.Gas += t.cost
}
