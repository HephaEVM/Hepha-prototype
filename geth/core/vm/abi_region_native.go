package vm

import (
	"github.com/holiman/uint256"
	miner "go-eth1.16.5-evm/template_miner"
)

// Recipes are proved from the bound SSA/effect graph, not compiler names or
// observations that a guard happened to pass. Unrecognized graphs retain the
// generic executor. All failure checks precede live memory/stack/gas mutation.
type abiAffine struct {
	input int
	add   uint256.Int
}
type abiClean struct {
	kind     byte
	argument uint256.Int
}
type abiNativeRead struct {
	offset abiAffine
	memory bool
	clean  abiClean
}
type abiNativeWrite struct {
	offset abiAffine
	input  int
	clean  abiClean
}
type abiNativeBound struct {
	end, start int
	size       uint256.Int
}
type abiNativeRegion struct {
	encode                     *abiFixedEncode
	fixed                      *abiFixedDecode
	kind                       byte // 1 decode, 2 encode, 3 forward word copy with trailing zero
	peak                       int
	cost, instructions         uint64
	loopCost, loopInstructions uint64
	bounds                     []abiNativeBound
	reads                      []abiNativeRead
	writes                     []abiNativeWrite
	outputs                    []abiAffine
	readOutputs                []int
}

func abiNode(nodes []miner.ABIValue, id int, op string, arity int) bool {
	return id >= 0 && id < len(nodes) && nodes[id].Op == op && len(nodes[id].Args) == arity
}
func abiInput(nodes []miner.ABIValue, id int) (int, bool) {
	if !abiNode(nodes, id, "input", 0) {
		return 0, false
	}
	return nodes[id].Input, true
}
func abiConst(nodes []miner.ABIValue, id int) (uint256.Int, bool) {
	var v uint256.Int
	if !abiNode(nodes, id, "const", 0) || v.SetFromHex(nodes[id].Value) != nil {
		return v, false
	}
	return v, true
}
func abiAffineOf(nodes []miner.ABIValue, id int) (abiAffine, bool) {
	if n, ok := abiInput(nodes, id); ok {
		return abiAffine{input: n}, true
	}
	if n, ok := abiConst(nodes, id); ok {
		return abiAffine{input: -1, add: n}, true
	}
	if abiNode(nodes, id, "ADD", 2) {
		a, ok := abiAffineOf(nodes, nodes[id].Args[0])
		if !ok {
			return a, false
		}
		b, ok := abiAffineOf(nodes, nodes[id].Args[1])
		if !ok || (a.input >= 0 && b.input >= 0) {
			return a, false
		}
		a.add.Add(&a.add, &b.add)
		if a.input < 0 {
			a.input = b.input
		}
		return a, true
	}
	return abiAffine{}, false
}
func (a *abiAffine) eval(input []uint256.Int) uint256.Int {
	v := a.add
	if a.input >= 0 {
		v.Add(&v, &input[a.input])
	}
	return v
}
func abiCleanupOf(nodes []miner.ABIValue, id, value int) (abiClean, bool) {
	if id == value {
		return abiClean{}, true
	}
	if abiNode(nodes, id, "AND", 2) {
		a, b := nodes[id].Args[0], nodes[id].Args[1]
		if b == value {
			a, b = b, a
		}
		if a == value {
			v, ok := abiConst(nodes, b)
			return abiClean{1, v}, ok
		}
	}
	if abiNode(nodes, id, "SIGNEXTEND", 2) && nodes[id].Args[1] == value {
		v, ok := abiConst(nodes, nodes[id].Args[0])
		return abiClean{2, v}, ok
	}
	if abiNode(nodes, id, "ISZERO", 1) {
		i := nodes[id].Args[0]
		if abiNode(nodes, i, "ISZERO", 1) && nodes[i].Args[0] == value {
			return abiClean{kind: 3}, true
		}
	}
	return abiClean{}, false
}
func (c *abiClean) apply(v *uint256.Int) {
	switch c.kind {
	case 1:
		v.And(v, &c.argument)
	case 2:
		v.ExtendSign(v, &c.argument)
	case 3:
		if v.IsZero() {
			v.Clear()
		} else {
			v.SetOne()
		}
	}
}
func prepareABINative(p *abiRegionPrepared) *abiNativeRegion {
	if len(p.graph) > 0 {
		return prepareABICopy(p)
	}
	if len(p.paths) != 1 {
		return nil
	}
	var path *miner.ABIRegionPath
	for i := range p.region.Paths {
		if p.region.Paths[i].Outcome == "return" {
			path = &p.region.Paths[i]
			break
		}
	}
	if path == nil {
		return nil
	}
	n := &abiNativeRegion{peak: p.paths[0].peak, cost: p.paths[0].cost, instructions: p.paths[0].instructions}
	if prepareABIDecode(n, path) {
		n.kind = 1
		return n
	}
	n = &abiNativeRegion{peak: p.paths[0].peak, cost: p.paths[0].cost, instructions: p.paths[0].instructions}
	if prepareABIEncode(n, path) {
		n.kind = 2
		return n
	}
	return nil
}
func prepareABIDecode(n *abiNativeRegion, p *miner.ABIRegionPath) bool {
	nodes := p.Nodes
	reads := map[int]int{}
	for _, a := range p.Actions {
		switch a.Op {
		case "CALLDATALOAD", "MLOAD":
			if len(a.Args) != 1 || len(n.reads) >= 18 {
				return false
			}
			o, ok := abiAffineOf(nodes, a.Args[0])
			if !ok {
				return false
			}
			reads[a.Result] = len(n.reads)
			n.reads = append(n.reads, abiNativeRead{offset: o, memory: a.Op == "MLOAD"})
		case "guard":
			if len(a.Args) != 1 {
				return false
			}
			id, truth := a.Args[0], a.Nonzero
			for abiNode(nodes, id, "ISZERO", 1) {
				id = nodes[id].Args[0]
				truth = !truth
			}
			if abiNode(nodes, id, "SLT", 2) && !truth {
				sub := nodes[id].Args[0]
				size, ok := abiConst(nodes, nodes[id].Args[1])
				if !ok || !abiNode(nodes, sub, "SUB", 2) {
					return false
				}
				end, eok := abiInput(nodes, nodes[sub].Args[0])
				start, sok := abiInput(nodes, nodes[sub].Args[1])
				if !eok || !sok {
					return false
				}
				n.bounds = append(n.bounds, abiNativeBound{end, start, size})
				continue
			}
			if !abiNode(nodes, id, "EQ", 2) || !truth {
				return false
			}
			left, right := nodes[id].Args[0], nodes[id].Args[1]
			index, ok := reads[left]
			if !ok {
				left, right = right, left
				index, ok = reads[left]
			}
			if !ok {
				return false
			}
			clean, ok := abiCleanupOf(nodes, right, left)
			if !ok {
				return false
			}
			if n.reads[index].clean.kind != 0 && n.reads[index].clean != clean {
				return false
			}
			if clean.kind != 0 {
				n.reads[index].clean = clean
			}
		default:
			return false
		}
	}
	if len(n.reads) == 0 || len(p.Outputs) > 18 {
		return false
	}
	for _, id := range p.Outputs {
		index, ok := reads[id]
		if !ok {
			return false
		}
		n.readOutputs = append(n.readOutputs, index)
	}
	return true
}
func prepareABIEncode(n *abiNativeRegion, p *miner.ABIRegionPath) bool {
	for _, a := range p.Actions {
		if a.Op != "MSTORE" || len(a.Args) != 2 || len(n.writes) >= 18 {
			return false
		}
		o, ok := abiAffineOf(p.Nodes, a.Args[0])
		if !ok {
			return false
		}
		found := false
		for id, v := range p.Nodes {
			if v.Op != "input" {
				continue
			}
			if clean, ok := abiCleanupOf(p.Nodes, a.Args[1], id); ok {
				n.writes = append(n.writes, abiNativeWrite{o, v.Input, clean})
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	if len(n.writes) == 0 || len(p.Outputs) > 18 {
		return false
	}
	for _, id := range p.Outputs {
		v, ok := abiAffineOf(p.Nodes, id)
		if !ok {
			return false
		}
		n.outputs = append(n.outputs, v)
	}
	return true
}

const abiMaxMemory = uint64(0x1FFFFFFFE0)

func abiNativeMemory(m *Memory, end, cost, gas uint64) (uint64, uint64, uint64, bool) {
	if end > abiMaxMemory || cost > gas {
		return 0, 0, 0, false
	}
	size, fee := uint64(m.Len()), m.lastGasCost
	end = (end + 31) &^ uint64(31)
	if end > size {
		words := end / 32
		next := 3*words + words*words/512
		if next < fee || next-fee > gas-cost {
			return 0, 0, 0, false
		}
		cost += next - fee
		size, fee = end, next
	}
	return size, fee, cost, true
}
func (e *ABIScalarDiagnostic) runABINative(n *abiNativeRegion, prefix int, s *ScopeContext) (uint64, bool) {
	if n.fixed != nil {
		if count, ok := e.runABIFixed(n, prefix, s); ok {
			return count, true
		}
	}
	if prefix+n.peak > 1024 || s.Contract.Gas < n.cost {
		return 0, false
	}
	if n.kind == 3 {
		return e.runABICopy(n, prefix, s)
	}
	input := s.Stack.data[prefix:]
	values := e.scratch.values[:]
	var offsets [18]uint64
	var end uint64
	signature := s.Signature
	if n.kind == 1 {
		for _, b := range n.bounds {
			var span uint256.Int
			span.Sub(&input[b.end], &input[b.start])
			if span.Slt(&b.size) {
				return 0, false
			}
		}
		for i := range n.reads {
			r := &n.reads[i]
			off := r.offset.eval(input)
			v := &values[i]
			if r.memory {
				o, ov := off.Uint64WithOverflow()
				if ov || o > abiMaxMemory-32 {
					return 0, false
				}
				end = max(end, o+32)
				if o+32 <= uint64(s.Memory.Len()) {
					v.SetBytes(s.Memory.store[o : o+32])
				} else {
					var word [32]byte
					if o < uint64(s.Memory.Len()) {
						copy(word[:], s.Memory.store[o:])
					}
					v.SetBytes(word[:])
				}
			} else {
				abiLoadCalldata(v, s.Contract.Input, &off)
				if off.IsZero() {
					word := v.Bytes32()
					signature = string(word[:4])
				}
			}
			clean := *v
			r.clean.apply(&clean)
			if clean != *v {
				return 0, false
			}
		}
	} else {
		if n.encode != nil {
			if count, ok := e.runABIFixedEncode(n, prefix, s); ok {
				return count, true
			}
		}
		for i := range n.writes {
			off := n.writes[i].offset.eval(input)
			o, ov := off.Uint64WithOverflow()
			if ov || o > abiMaxMemory-32 {
				return 0, false
			}
			offsets[i] = o
			end = max(end, o+32)
		}
		for i := range n.outputs {
			values[i] = n.outputs[i].eval(input)
		}
	}
	size, fee, cost, ok := abiNativeMemory(s.Memory, end, n.cost, s.Contract.Gas)
	if !ok {
		return 0, false
	}
	// Admission is complete; none of the following operations can reject.
	s.Memory.Resize(size)
	s.Memory.lastGasCost = fee
	if n.kind == 2 {
		for i := range n.writes {
			w := &n.writes[i]
			v := input[w.input]
			w.clean.apply(&v)
			s.Memory.Set32(offsets[i], &v)
		}
	}
	stack := s.Stack.data[:prefix]
	if n.kind == 1 {
		for _, i := range n.readOutputs {
			stack = append(stack, values[i])
		}
	} else {
		for i := range n.outputs {
			stack = append(stack, values[i])
		}
	}
	s.Stack.data = stack
	s.Contract.Gas -= cost
	s.Signature = signature
	e.Stats.Gas += cost
	return n.instructions, true
}
