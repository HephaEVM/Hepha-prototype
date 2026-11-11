package vm

import (
	miner "go-eth1.16.5-evm/template_miner"
)

func abiIsInput(n []miner.ABIValue, id, input int) bool {
	i, ok := abiInput(n, id)
	return ok && i == input
}
func abiIsConst(n []miner.ABIValue, id int, want uint64) bool {
	v, ok := abiConst(n, id)
	return ok && v.IsUint64() && v.Uint64() == want
}
func abiAddInputs(n []miner.ABIValue, id, a, b int) bool {
	if !abiNode(n, id, "ADD", 2) {
		return false
	}
	x, y := n[id].Args[0], n[id].Args[1]
	return (abiIsInput(n, x, a) && abiIsInput(n, y, b)) || (abiIsInput(n, x, b) && abiIsInput(n, y, a))
}
func abiIdentityOutputs(p *miner.ABIRegionPath, count int) bool {
	if len(p.Outputs) < count {
		return false
	}
	for i := 0; i < count; i++ {
		if !abiIsInput(p.Nodes, p.Outputs[i], i) {
			return false
		}
	}
	return true
}
func abiCopyCondition(b *miner.ABIRegionBlock, initial bool) bool {
	n, id := b.Path.Nodes, b.Condition
	if !abiNode(n, id, "ISZERO", 1) {
		return false
	}
	id = n[id].Args[0]
	if !abiNode(n, id, "LT", 2) || !abiIsInput(n, n[id].Args[1], 1) {
		return false
	}
	if initial {
		return abiIsConst(n, n[id].Args[0], 0)
	}
	return abiIsInput(n, n[id].Args[0], 4)
}
func prepareABICopy(p *abiRegionPrepared) *abiNativeRegion {
	r := p.region
	if r.Inputs != 4 || len(r.Graph) != 4 || len(p.graph) != 4 {
		return nil
	}
	g := r.Graph
	entry := &g[0]
	valid := func(i int) bool { return i >= 0 && i < len(g) }
	if entry.Path.Outcome != "branch" || len(entry.Path.Actions) != 0 || !valid(entry.Next) || !valid(entry.Else) || !abiCopyCondition(entry, true) {
		return nil
	}
	exitID, bodyID := entry.Next, entry.Else
	body, exit := &g[bodyID], &g[exitID]
	if body.Path.Outcome != "next" || !valid(body.Next) {
		return nil
	}
	loopID := body.Next
	loop := &g[loopID]
	if exitID == 0 || bodyID == 0 || loopID == 0 || exitID == bodyID || exitID == loopID || bodyID == loopID {
		return nil
	}
	if loop.Path.Outcome != "branch" || loop.Next != exitID || loop.Else != bodyID || len(loop.Path.Actions) != 0 || !abiCopyCondition(loop, false) {
		return nil
	}
	if len(entry.Path.Outputs) != 5 || !abiIdentityOutputs(&entry.Path, 4) || !abiIsConst(entry.Path.Nodes, entry.Path.Outputs[4], 0) {
		return nil
	}
	if len(loop.Path.Outputs) != 5 || !abiIdentityOutputs(&loop.Path, 5) {
		return nil
	}
	bp := &body.Path
	if len(bp.Actions) != 2 || len(bp.Outputs) != 5 || !abiIdentityOutputs(bp, 4) {
		return nil
	}
	read, write := bp.Actions[0], bp.Actions[1]
	if read.Op != "MLOAD" || len(read.Args) != 1 || !abiAddInputs(bp.Nodes, read.Args[0], 3, 4) {
		return nil
	}
	if write.Op != "MSTORE" || len(write.Args) != 2 || write.Args[1] != read.Result || !abiAddInputs(bp.Nodes, write.Args[0], 2, 4) {
		return nil
	}
	next, ok := abiAffineOf(bp.Nodes, bp.Outputs[4])
	if !ok || next.input != 4 || !next.add.IsUint64() || next.add.Uint64() != 32 {
		return nil
	}
	ep := &exit.Path
	if ep.Outcome != "return" || len(ep.Outputs) != 0 || len(ep.Actions) != 1 {
		return nil
	}
	tail := ep.Actions[0]
	if tail.Op != "MSTORE" || len(tail.Args) != 2 || !abiAddInputs(ep.Nodes, tail.Args[0], 2, 1) || !abiIsConst(ep.Nodes, tail.Args[1], 0) {
		return nil
	}
	n := &abiNativeRegion{kind: 3}
	for i := range p.graph {
		n.peak = max(n.peak, p.graph[i].plan.peak)
	}
	n.cost = p.graph[0].plan.cost + p.graph[exitID].plan.cost
	n.instructions = p.graph[0].plan.instructions + p.graph[exitID].plan.instructions
	n.loopCost = p.graph[bodyID].plan.cost + p.graph[loopID].plan.cost
	n.loopInstructions = p.graph[bodyID].plan.instructions + p.graph[loopID].plan.instructions
	if n.loopCost == 0 {
		return nil
	}
	return n
}
func (e *ABIScalarDiagnostic) runABICopy(n *abiNativeRegion, prefix int, s *ScopeContext) (uint64, bool) {
	input := s.Stack.data[prefix:]
	length, ov := input[1].Uint64WithOverflow()
	if ov || length > 4096*32 {
		return 0, false
	}
	dst, ov := input[2].Uint64WithOverflow()
	if ov || dst > abiMaxMemory-length-32 {
		return 0, false
	}
	words := (length + 31) / 32
	if words > (s.Contract.Gas-n.cost)/n.loopCost {
		return 0, false
	}
	cost := n.cost + words*n.loopCost
	end := dst + length + 32
	src := uint64(0)
	if words > 0 {
		src, ov = input[3].Uint64WithOverflow()
		if ov || src > abiMaxMemory-words*32 {
			return 0, false
		}
		end = max(end, src+words*32)
	}
	size, fee, cost, ok := abiNativeMemory(s.Memory, end, cost, s.Contract.Gas)
	if !ok {
		return 0, false
	}
	s.Memory.Resize(size)
	s.Memory.lastGasCost = fee
	// Each original MLOAD snapshots one word before its MSTORE. A whole-range
	// memmove would change forward-overlap semantics. Keep exactly this order.
	for i := uint64(0); i < words; i++ {
		o := i * 32
		copy(s.Memory.store[dst+o:dst+o+32], s.Memory.store[src+o:src+o+32])
	}
	clear(s.Memory.store[dst+length : dst+length+32])
	s.Stack.data = s.Stack.data[:prefix]
	s.Contract.Gas -= cost
	e.Stats.Gas += cost
	return n.instructions + words*n.loopInstructions, true
}
