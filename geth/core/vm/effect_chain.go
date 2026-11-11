package vm

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/holiman/uint256"
	"strings"
)

// Analysis is performed once at binding. Values are SSA names; a storage load
// always gets a fresh name, even when its address is reused across a write.
const chainInput OpCode = 0xe0
const chainConstant OpCode = 0xe1
const chainHelper OpCode = 0xe2
const chainHashReuse OpCode = 0xe3

type chainNode struct {
	op          OpCode
	a, b, depth int
	c           uint256.Int
}
type chainWrite struct{ offset, id int }
type chainAction struct {
	fastEnd, fastOffset, fastExtra uint64
	op                             OpCode
	pc, start                      uint64
	gas                            uint64
	args                           []int
	result                         int
	reuse                          int
	before, checkpoint             []int
	helper                         *arithmeticRegion
	startStep, opStep              int
}
type chainEvent struct{ node, action int }
type chainRepeat struct{ action, node int }
type effectChain struct {
	directFast                                   func(*effectChain, *uint64, *ScopeContext, *EVM) (bool, error)
	entry, exit                                  uint64
	minStack, peak, delta                        int
	nodes                                        []chainNode
	actions                                      []chainAction
	events                                       []chainEvent
	repeats                                      []chainRepeat
	repeatAt                                     [][]int
	writes                                       []chainWrite
	pcs                                          []uint64
	hashes, reused, loads, stores, logs, helpers int
	shape                                        string
	direct                                       func(*effectChain, *uint64, *ScopeContext, *EVM) (bool, error)
}
type ChainCounters struct{ Entries, Commits, Resumes, Hashes, Reused, StackWrites uint64 }

func chainPure(op OpCode, a, b uint256.Int) uint256.Int {
	var z uint256.Int
	switch op {
	case ADD:
		z.Add(&a, &b)
	case SUB:
		z.Sub(&a, &b)
	case MUL:
		z.Mul(&a, &b)
	case DIV:
		z.Div(&a, &b)
	case AND:
		z.And(&a, &b)
	case OR:
		z.Or(&a, &b)
	case XOR:
		z.Xor(&a, &b)
	case NOT:
		z.Not(&a)
	case EQ:
		if a == b {
			z.SetOne()
		}
	case LT:
		if a.Lt(&b) {
			z.SetOne()
		}
	case GT:
		if a.Gt(&b) {
			z.SetOne()
		}
	case ISZERO:
		if a.IsZero() {
			z.SetOne()
		}
	case SHL:
		if a.LtUint64(256) {
			z.Lsh(&b, uint(a.Uint64()))
		}
	case SHR:
		if a.LtUint64(256) {
			z.Rsh(&b, uint(a.Uint64()))
		}
	}
	return z
}
func chainPureOp(op OpCode) bool {
	switch op {
	case ADD, SUB, MUL, DIV, AND, OR, XOR, NOT, EQ, LT, GT, ISZERO, SHL, SHR:
		return true
	}
	return false
}

func bindEffectChain(code []byte, entry uint64, bound *boundArithmetic, table *JumpTable) *effectChain {
	if bound == nil {
		return nil
	}
	p := &effectChain{entry: entry}
	st := make([]int, 32)
	memo := map[chainNode]int{}
	add := func(v chainNode) int {
		if id, ok := memo[v]; ok {
			if paperWithoutReuse && chainPureOp(v.op) {
				p.repeats = append(p.repeats, chainRepeat{len(p.actions), id})
			}
			return id
		}
		id := len(p.nodes)
		p.nodes = append(p.nodes, v)
		memo[v] = id
		p.events = append(p.events, chainEvent{node: id, action: -1})
		return id
	}
	for i := range st {
		st[i] = add(chainNode{op: chainInput, depth: 32 - i})
	}
	constant := func(x uint64) int { return add(chainNode{op: chainConstant, c: *uint256.NewInt(x)}) }
	known := func(id int) (uint64, bool) {
		v := p.nodes[id]
		return v.c.Uint64(), v.op == chainConstant && v.c.IsUint64()
	}
	pop := func() int { id := st[len(st)-1]; st = st[:len(st)-1]; return id }
	push := func(id int) {
		st = append(st, id)
		if len(st)-32 > p.peak {
			p.peak = len(st) - 32
		}
	}
	need := func(n int) bool {
		if len(st) < n {
			return false
		}
		if d := 32 - len(st) + n; d > p.minStack {
			p.minStack = d
		}
		return p.minStack <= 32
	}
	copyStack := func() []int { return append([]int(nil), st...) }
	phaseStack, phasePC := copyStack(), entry
	phaseStep := 0
	pc, pending := entry, uint64(0)
	bitmap := codeBitmap(code)
	visits := map[uint64]int{}
	mem := map[uint64]int{}
	hashes := map[[2]int]int{}
	var last *effectChain
	snapshot := func() {
		if p.reused < 2 || p.hashes < 3 || p.loads < 2 {
			return
		}
		q := *p
		q.repeats = append([]chainRepeat(nil), p.repeats...)
		q.exit = pc + 1
		q.delta = len(st) - 32
		q.nodes = append([]chainNode(nil), p.nodes...)
		q.actions = append([]chainAction(nil), p.actions...)
		q.events = append([]chainEvent(nil), p.events...)
		q.pcs = append([]uint64(nil), p.pcs...)
		for i, id := range st {
			if id == i && i < 32 {
				continue
			}
			q.writes = append(q.writes, chainWrite{i - 32, id})
		}
		last = &q
	}
	for step := 0; step < 640 && pc < uint64(len(code)) && len(p.nodes) < 240 && len(p.actions) < 96; step++ {
		visits[pc]++
		if visits[pc] > 4 {
			break
		}
		op := OpCode(code[pc])
		o := table[op]
		if o == nil || o.undefined || !need(o.minStack) {
			break
		}
		before := copyStack()
		p.pcs = append(p.pcs, pc)
		if op >= PUSH0 && op <= PUSH32 {
			n := uint64(op - PUSH0)
			if pc+n >= uint64(len(code)) {
				break
			}
			var v uint256.Int
			v.SetBytes(code[pc+1 : pc+1+n])
			push(add(chainNode{op: chainConstant, c: v}))
			pending += o.constantGas
			pc += n + 1
			continue
		}
		if op >= DUP1 && op <= DUP16 {
			push(st[len(st)-int(op-DUP1)-1])
			pending += o.constantGas
			pc++
			continue
		}
		if op >= SWAP1 && op <= SWAP16 {
			i := len(st) - int(op-SWAP1) - 2
			st[i], st[len(st)-1] = st[len(st)-1], st[i]
			pending += o.constantGas
			pc++
			continue
		}
		if op == POP {
			pop()
			pending += o.constantGas
			pc++
			continue
		}
		if op == JUMPDEST {
			pending += o.constantGas
			pc++
			continue
		}
		if op == CALLER {
			push(add(chainNode{op: CALLER}))
			pending += o.constantGas
			pc++
			continue
		}
		if chainPureOp(op) {
			a := pop()
			c := constant(0)
			if op != NOT && op != ISZERO {
				c = pop()
			}
			x, y := p.nodes[a], p.nodes[c]
			id := -1
			if x.op == chainConstant && y.op == chainConstant {
				id = add(chainNode{op: chainConstant, c: chainPure(op, x.c, y.c)})
			} else {
				if op == AND || op == OR || op == XOR || op == ADD || op == MUL || op == EQ {
					if a > c {
						a, c = c, a
					}
				}
				id = add(chainNode{op: op, a: a, b: c})
			}
			push(id)
			pending += o.constantGas
			pc++
			continue
		}
		a := chainAction{op: op, pc: pc, start: phasePC, startStep: phaseStep, opStep: len(p.pcs) - 1, gas: pending + o.constantGas, result: -1, reuse: -1, before: before, checkpoint: phaseStack}
		appendAction := func() {
			p.actions = append(p.actions, a)
			p.events = append(p.events, chainEvent{node: -1, action: len(p.actions) - 1})
			pending = 0
			phaseStack = copyStack()
			phasePC = pc + 1
			phaseStep = len(p.pcs)
		}
		fresh := func() int { id := len(p.nodes); p.nodes = append(p.nodes, chainNode{op: 0xfe}); return id }
		stop := false
		switch op {
		case JUMP:
			dest, ok := known(pop())
			if !ok || dest >= uint64(len(code)) || OpCode(code[dest]) != JUMPDEST || !bitmap.codeSegment(dest) {
				stop = true
				break
			}
			h := bound.regions[dest]
			if h != nil && h.syncHot.native != nil && h.syncHot.fixedLayout != 0 && h.extraGas == 0 && h.constantSuccessGas && !h.pathDependentResult && need(3) {
				ret, ok := known(st[len(st)-3])
				if !ok || ret >= uint64(len(code)) || OpCode(code[ret]) != JUMPDEST || !bitmap.codeSegment(ret) {
					stop = true
					break
				}
				x, y := st[len(st)-2], st[len(st)-1]
				if h.syncHot.fixedLayout == 2 {
					x, y = y, x
				}
				a.op = chainHelper
				a.args = []int{x, y}
				a.helper = h
				a.gas += h.syncHot.admissionGas
				a.result = fresh()
				if peak := len(st) - 32 + h.peak - 3; peak > p.peak {
					p.peak = peak
				}
				st = st[:len(st)-3]
				push(a.result)
				p.helpers++
				appendAction()
				phasePC = ret
				pc = ret
				continue
			}
			pending += o.constantGas
			pc = dest
			continue
		case JUMPI:
			dest, ok := known(pop())
			v := pop()
			if !ok || dest <= pc || dest >= uint64(len(code)) || OpCode(code[dest]) != JUMPDEST || !bitmap.codeSegment(dest) {
				stop = true
				break
			}
			a.args = []int{v}
			appendAction()
			phasePC = dest
			pc = dest
			continue
		case MSTORE:
			off, v := pop(), pop()
			a.args = []int{v, off}
			if k, ok := known(off); ok {
				for x := range mem {
					if x < k+32 && k < x+32 {
						delete(mem, x)
					}
				}
				mem[k] = v
			} else {
				mem = map[uint64]int{}
			}
			appendAction()
		case MLOAD:
			off := pop()
			a.args = []int{off}
			a.result = fresh()
			push(a.result)
			appendAction()
		case KECCAK256:
			off, size := pop(), pop()
			x, ok := known(off)
			y, ok2 := known(size)
			v0, k0 := mem[x]
			v1, k1 := mem[x+32]
			if !ok || !ok2 || x > 4096 || y != 64 || !k0 || !k1 {
				stop = true
				break
			}
			a.args = []int{size, off}
			p.hashes++
			key := [2]int{v0, v1}
			if id, found := hashes[key]; found {
				a.reuse = id
				push(id)
				p.reused++
			} else {
				a.result = fresh()
				hashes[key] = a.result
				push(a.result)
			}
			appendAction()
			snapshot()
		case SLOAD:
			a.args = []int{pop()}
			a.result = fresh()
			push(a.result)
			p.loads++
			appendAction()
		case SSTORE:
			key, v := pop(), pop()
			a.args = []int{v, key}
			p.stores++
			appendAction()
			snapshot()
		case LOG0, LOG1, LOG2, LOG3, LOG4:
			n := 2 + int(op-LOG0)
			a.args = append([]int(nil), st[len(st)-n:]...)
			st = st[:len(st)-n]
			p.logs++
			appendAction()
			snapshot()
		default:
			stop = true
		}
		if stop {
			break
		}
		pc++
	}
	if last == nil || last.stores == 0 {
		return nil
	}
	// Shapes exclude constants, source depths, PCs and contract identity. Direct
	// executors are generated from this structural form and receive bound values.
	s := ""
	for _, v := range last.nodes {
		s += fmt.Sprintf("%x/%d/%d;", byte(v.op), v.a, v.b)
	}
	for _, a := range last.actions {
		s += fmt.Sprintf("A%x/%v/%d/%d;", byte(a.op), a.args, a.result, a.reuse)
	}
	for _, v := range last.events {
		s += fmt.Sprintf("E%d/%d;", v.node, v.action)
	}
	h := sha256.Sum256([]byte(s))
	last.shape = hex.EncodeToString(h[:])
	last.direct = chainDirect[last.shape]
	if paperWithoutReuse {
		paperDisableChainReuse(last)
	}
	return last
}

func (p *effectChain) restore(s *ScopeContext, in []uint256.Int, v *[256]uint256.Int, stack []int) {
	n := len(in)
	end := n - 32 + len(stack)
	if cap(s.Stack.data) < end {
		s.Stack.data = append(s.Stack.data, make([]uint256.Int, end-len(s.Stack.data))...)
	} else {
		s.Stack.data = s.Stack.data[:end]
	}
	for i, id := range stack {
		pos := n - 32 + i
		if pos >= 0 {
			s.Stack.data[pos] = v[id]
		}
	}
}
func (p *effectChain) finish(pc *uint64, s *ScopeContext, in []uint256.Int, v *[256]uint256.Int) {
	if paperWithoutReuse {
		paperRepeatChainValues(p, len(p.actions), v)
	}
	n := len(in)
	end := n + p.delta
	if cap(s.Stack.data) < end {
		s.Stack.data = append(s.Stack.data, make([]uint256.Int, end-len(s.Stack.data))...)
	} else {
		s.Stack.data = s.Stack.data[:end]
	}
	for _, w := range p.writes {
		s.Stack.data[n+w.offset] = v[w.id]
	}
	*pc = p.exit
}

// Returns 0 to continue, 1 for an exact native continuation, 2 for an error.
func (p *effectChain) action(i int, pc *uint64, s *ScopeContext, e *EVM, in []uint256.Int, v *[256]uint256.Int) (int, error) {
	a := &p.actions[i]
	if paperWithoutReuse {
		paperRepeatChainValues(p, i, v)
	}
	if s.Contract.Gas < a.gas || e.abort.Load() {
		e.chainProgress = a.startStep
		p.restore(s, in, v, a.checkpoint)
		*pc = a.start
		return 1, nil
	}
	if a.op == JUMPI && v[a.args[0]].IsZero() {
		e.chainProgress = a.opStep
		p.restore(s, in, v, a.before)
		s.Contract.Gas -= a.gas - e.table[JUMPI].constantGas
		*pc = a.pc
		return 1, nil
	}
	if a.op == chainHelper {
		ix, z := a.helper.syncHot.native(v[a.args[0]], v[a.args[1]])
		if a.helper.syncHot.costs[ix] != a.helper.syncHot.admissionGas {
			e.chainProgress = a.opStep
			p.restore(s, in, v, a.before)
			s.Contract.Gas -= a.gas - e.table[JUMP].constantGas - a.helper.syncHot.admissionGas
			*pc = a.pc
			return 1, nil
		}
		v[a.result] = z
		s.Contract.Gas -= a.gas
		return 0, nil
	}
	s.Contract.Gas -= a.gas
	if a.op == JUMPI {
		return 0, nil
	}
	old := s.Stack.data
	// A small separate argument stack prevents temporary host operations from
	// overwriting live EVM inputs. Original dynamic gas and effects stay in order.
	scratch := &e.chainScratch
	for j, id := range a.args {
		scratch[j] = v[id]
	}
	s.Stack.data = scratch[:len(a.args)]
	o := e.table[a.op]
	size := uint64(0)
	var err error
	if o.memorySize != nil {
		var overflow bool
		size, overflow = o.memorySize(s.Stack)
		if overflow {
			err = ErrGasUintOverflow
		} else {
			size, overflow = math.SafeMul(toWordSize(size), 32)
			if overflow {
				err = ErrGasUintOverflow
			}
		}
	}
	if err == nil && o.dynamicGas != nil {
		var cost uint64
		cost, err = o.dynamicGas(e, s.Contract, s.Stack, s.Memory, size)
		if err != nil {
			err = fmt.Errorf("%w: %v", ErrOutOfGas, err)
		} else if s.Contract.Gas < cost {
			err = ErrOutOfGas
		} else {
			s.Contract.Gas -= cost
		}
	}
	if err == nil {
		if size > 0 {
			s.Memory.Resize(size)
		}
		if a.op == KECCAK256 && a.reuse >= 0 { /* Same symbolic preimage; charge original dynamic gas above. */
		} else {
			e.chainExecPC = a.pc
			_, err = o.execute(&e.chainExecPC, e, s)
			if err == nil && a.result >= 0 {
				v[a.result] = *s.Stack.peek()
			}
		}
	}
	s.Stack.data = old
	if err != nil {
		p.restore(s, in, v, a.before)
		*pc = a.pc
		return 2, err
	}
	return 0, nil
}
func (p *effectChain) try(pc *uint64, s *ScopeContext, e *EVM, special bool) (bool, error) {
	n := len(s.Stack.data)
	if n < p.minStack || n+p.peak > 1024 || e.Config.Tracer != nil || e.Config.EnablePreimageRecording || e.abort.Load() {
		return false, nil
	}
	// Reject an entry-only guard before allocating/evaluating the region. No
	// effect or gas charge has occurred, so the original opcode can run directly.
	if len(p.actions) > 0 && p.actions[0].op == JUMPI {
		g := p.nodes[p.actions[0].args[0]]
		if g.op == chainInput && inGuardZero(s, g.depth) {
			return false, nil
		}
	}
	if special && p.direct != nil {
		return p.direct(p, pc, s, e)
	}
	in := s.Stack.data
	var v [256]uint256.Int
	for _, event := range p.events {
		if event.node >= 0 {
			j := event.node
			x := p.nodes[j]
			switch x.op {
			case chainInput:
				if x.depth <= len(in) {
					v[j] = in[len(in)-x.depth]
				}
			case chainConstant:
				v[j] = x.c
			case CALLER:
				v[j].SetBytes(effectiveCaller(s))
			default:
				v[j] = chainPure(x.op, v[x.a], v[x.b])
			}
			continue
		}
		code, err := p.action(event.action, pc, s, e, in, &v)
		if code != 0 {
			return *pc != p.entry || event.action != 0, err
		}
	}
	p.finish(pc, s, in, &v)
	e.chainProgress = len(p.pcs)
	return true, nil
}

var chainDirect = map[string]func(*effectChain, *uint64, *ScopeContext, *EVM) (bool, error){}

func inGuardZero(s *ScopeContext, depth int) bool {
	return depth > len(s.Stack.data) || s.Stack.data[len(s.Stack.data)-depth].IsZero()
}

func bindEffectChainShared(code []byte, entry uint64, bound *boundArithmetic, table *JumpTable, bitmap BitVec) *effectChain {
	return bindEffectChainThreshold(code, entry, bound, table, bitmap, 2)
}

// The native backend also lowers shorter, proved address-reuse chains.
func bindEffectChainThreshold(code []byte, entry uint64, bound *boundArithmetic, table *JumpTable, bitmap BitVec, repeats int) *effectChain {
	if bound == nil {
		return nil
	}
	p := &effectChain{entry: entry}
	st := make([]int, 32)
	memo := map[chainNode]int{}
	add := func(v chainNode) int {
		if id, ok := memo[v]; ok {
			if paperWithoutReuse && chainPureOp(v.op) {
				p.repeats = append(p.repeats, chainRepeat{len(p.actions), id})
			}
			return id
		}
		id := len(p.nodes)
		p.nodes = append(p.nodes, v)
		memo[v] = id
		p.events = append(p.events, chainEvent{node: id, action: -1})
		return id
	}
	for i := range st {
		st[i] = add(chainNode{op: chainInput, depth: 32 - i})
	}
	constant := func(x uint64) int { return add(chainNode{op: chainConstant, c: *uint256.NewInt(x)}) }
	known := func(id int) (uint64, bool) {
		v := p.nodes[id]
		return v.c.Uint64(), v.op == chainConstant && v.c.IsUint64()
	}
	pop := func() int { id := st[len(st)-1]; st = st[:len(st)-1]; return id }
	push := func(id int) {
		st = append(st, id)
		if len(st)-32 > p.peak {
			p.peak = len(st) - 32
		}
	}
	need := func(n int) bool {
		if len(st) < n {
			return false
		}
		if d := 32 - len(st) + n; d > p.minStack {
			p.minStack = d
		}
		return p.minStack <= 32
	}
	copyStack := func() []int { return append([]int(nil), st...) }
	phaseStack, phasePC := copyStack(), entry
	phaseStep := 0
	pc, pending := entry, uint64(0)
	visits := map[uint64]int{}
	mem := map[uint64]int{}
	hashes := map[[2]int]int{}
	var last *effectChain
	snapshot := func() {
		if p.reused < repeats || p.hashes < repeats+1 || p.loads < repeats {
			return
		}
		q := *p
		q.repeats = append([]chainRepeat(nil), p.repeats...)
		q.exit = pc + 1
		q.delta = len(st) - 32
		q.nodes = p.nodes
		q.actions = p.actions
		q.events = p.events
		q.pcs = p.pcs
		for i, id := range st {
			if id == i && i < 32 {
				continue
			}
			q.writes = append(q.writes, chainWrite{i - 32, id})
		}
		last = &q
	}
	for step := 0; step < 640 && pc < uint64(len(code)) && len(p.nodes) < 240 && len(p.actions) < 96; step++ {
		visits[pc]++
		if visits[pc] > 4 {
			break
		}
		op := OpCode(code[pc])
		o := table[op]
		if o == nil || o.undefined || !need(o.minStack) {
			break
		}
		before := copyStack()
		p.pcs = append(p.pcs, pc)
		if op >= PUSH0 && op <= PUSH32 {
			n := uint64(op - PUSH0)
			if pc+n >= uint64(len(code)) {
				break
			}
			var v uint256.Int
			v.SetBytes(code[pc+1 : pc+1+n])
			push(add(chainNode{op: chainConstant, c: v}))
			pending += o.constantGas
			pc += n + 1
			continue
		}
		if op >= DUP1 && op <= DUP16 {
			push(st[len(st)-int(op-DUP1)-1])
			pending += o.constantGas
			pc++
			continue
		}
		if op >= SWAP1 && op <= SWAP16 {
			i := len(st) - int(op-SWAP1) - 2
			st[i], st[len(st)-1] = st[len(st)-1], st[i]
			pending += o.constantGas
			pc++
			continue
		}
		if op == POP {
			pop()
			pending += o.constantGas
			pc++
			continue
		}
		if op == JUMPDEST {
			pending += o.constantGas
			pc++
			continue
		}
		if op == CALLER {
			push(add(chainNode{op: CALLER}))
			pending += o.constantGas
			pc++
			continue
		}
		if chainPureOp(op) {
			a := pop()
			c := constant(0)
			if op != NOT && op != ISZERO {
				c = pop()
			}
			x, y := p.nodes[a], p.nodes[c]
			id := -1
			if x.op == chainConstant && y.op == chainConstant {
				id = add(chainNode{op: chainConstant, c: chainPure(op, x.c, y.c)})
			} else {
				if op == AND || op == OR || op == XOR || op == ADD || op == MUL || op == EQ {
					if a > c {
						a, c = c, a
					}
				}
				id = add(chainNode{op: op, a: a, b: c})
			}
			push(id)
			pending += o.constantGas
			pc++
			continue
		}
		a := chainAction{op: op, pc: pc, start: phasePC, startStep: phaseStep, opStep: len(p.pcs) - 1, gas: pending + o.constantGas, result: -1, reuse: -1, before: before, checkpoint: phaseStack}
		appendAction := func() {
			p.actions = append(p.actions, a)
			p.events = append(p.events, chainEvent{node: -1, action: len(p.actions) - 1})
			pending = 0
			phaseStack = copyStack()
			phasePC = pc + 1
			phaseStep = len(p.pcs)
		}
		fresh := func() int { id := len(p.nodes); p.nodes = append(p.nodes, chainNode{op: 0xfe}); return id }
		stop := false
		switch op {
		case JUMP:
			dest, ok := known(pop())
			if !ok || dest >= uint64(len(code)) || OpCode(code[dest]) != JUMPDEST || !bitmap.codeSegment(dest) {
				stop = true
				break
			}
			h := bound.regions[dest]
			if h != nil && h.syncHot.native != nil && h.syncHot.fixedLayout != 0 && h.extraGas == 0 && h.constantSuccessGas && !h.pathDependentResult && need(3) {
				ret, ok := known(st[len(st)-3])
				if !ok || ret >= uint64(len(code)) || OpCode(code[ret]) != JUMPDEST || !bitmap.codeSegment(ret) {
					stop = true
					break
				}
				x, y := st[len(st)-2], st[len(st)-1]
				if h.syncHot.fixedLayout == 2 {
					x, y = y, x
				}
				a.op = chainHelper
				a.args = []int{x, y}
				a.helper = h
				a.gas += h.syncHot.admissionGas
				a.result = fresh()
				if peak := len(st) - 32 + h.peak - 3; peak > p.peak {
					p.peak = peak
				}
				st = st[:len(st)-3]
				push(a.result)
				p.helpers++
				appendAction()
				phasePC = ret
				pc = ret
				continue
			}
			pending += o.constantGas
			pc = dest
			continue
		case JUMPI:
			dest, ok := known(pop())
			v := pop()
			if !ok || dest <= pc || dest >= uint64(len(code)) || OpCode(code[dest]) != JUMPDEST || !bitmap.codeSegment(dest) {
				stop = true
				break
			}
			a.args = []int{v}
			appendAction()
			phasePC = dest
			pc = dest
			continue
		case MSTORE:
			off, v := pop(), pop()
			a.args = []int{v, off}
			if k, ok := known(off); ok {
				for x := range mem {
					if x < k+32 && k < x+32 {
						delete(mem, x)
					}
				}
				mem[k] = v
			} else {
				mem = map[uint64]int{}
			}
			appendAction()
		case MLOAD:
			off := pop()
			a.args = []int{off}
			a.result = fresh()
			push(a.result)
			appendAction()
		case KECCAK256:
			off, size := pop(), pop()
			x, ok := known(off)
			y, ok2 := known(size)
			v0, k0 := mem[x]
			v1, k1 := mem[x+32]
			if !ok || !ok2 || x > 4096 || y != 64 || !k0 || !k1 {
				stop = true
				break
			}
			a.args = []int{size, off}
			p.hashes++
			key := [2]int{v0, v1}
			if id, found := hashes[key]; found {
				a.reuse = id
				push(id)
				p.reused++
			} else {
				a.result = fresh()
				hashes[key] = a.result
				push(a.result)
			}
			appendAction()
			snapshot()
		case SLOAD:
			a.args = []int{pop()}
			a.result = fresh()
			push(a.result)
			p.loads++
			appendAction()
		case SSTORE:
			key, v := pop(), pop()
			a.args = []int{v, key}
			p.stores++
			appendAction()
			snapshot()
		case LOG0, LOG1, LOG2, LOG3, LOG4:
			n := 2 + int(op-LOG0)
			a.args = append([]int(nil), st[len(st)-n:]...)
			st = st[:len(st)-n]
			p.logs++
			appendAction()
			snapshot()
		default:
			stop = true
		}
		if stop {
			break
		}
		pc++
	}
	if last == nil || last.stores == 0 {
		return nil
	}
	// Shapes exclude constants, source depths, PCs and contract identity. Direct
	// executors are generated from this structural form and receive bound values.
	var shape strings.Builder
	for _, v := range last.nodes {
		fmt.Fprintf(&shape, "%x/%d/%d;", byte(v.op), v.a, v.b)
	}
	for _, a := range last.actions {
		fmt.Fprintf(&shape, "A%x/%v/%d/%d;", byte(a.op), a.args, a.result, a.reuse)
	}
	for _, v := range last.events {
		fmt.Fprintf(&shape, "E%d/%d;", v.node, v.action)
	}
	h := sha256.Sum256([]byte(shape.String()))
	last.shape = hex.EncodeToString(h[:])
	last.direct = chainDirect[last.shape]
	if paperWithoutReuse {
		paperDisableChainReuse(last)
	}
	prepareChainFast(last, table)
	return last
}
