package vm

// Join two already verified, finite regions. The bridge contains only stack
// preparation and statically resolved jumps. Existing state/call boundaries
// and all original gas/stack obligations remain intact.
func regionBridge(a *hotRegion, targets map[uint64]*demandPlan, code []byte, table *JumpTable, bitmap BitVec) (*hotRegion, []int, *demandPlan) {
	q := *a
	q.nodes = a.nodes[:len(a.nodes):len(a.nodes)]
	st := make([]int, hotRegionInputs-a.Drop)
	for i := range st {
		st[i] = i
	}
	st = append(st, a.outputs...)
	known := func(id int) (uint64, bool) {
		if id < hotRegionInputs || id-hotRegionInputs >= len(q.nodes) {
			return 0, false
		}
		n := q.nodes[id-hotRegionInputs]
		return n.value.Uint64(), n.op == PUSH32 && n.value.IsUint64()
	}
	pc := a.Exit
	seen := map[uint64]bool{}
	for step := 0; ; step++ {
		if p := targets[pc]; p != nil && pc != a.Entry {
			return &q, st, p
		}
		if step >= 12 || pc >= uint64(len(code)) || !bitmap.codeSegment(pc) || seen[pc] {
			return nil, nil, nil
		}
		seen[pc] = true
		op := OpCode(code[pc])
		ins := table[op]
		if ins == nil || ins.undefined || ins.dynamicGas != nil || ins.memorySize != nil || len(st) < ins.minStack {
			return nil, nil, nil
		}
		delta := len(st) - hotRegionInputs
		q.Inputs = max(q.Inputs, ins.minStack-delta)
		q.MaxStack = min(q.MaxStack, ins.maxStack-delta)
		q.Gas += ins.constantGas
		q.Instructions++
		next := pc + 1
		switch {
		case op >= PUSH0 && op <= PUSH32:
			n := uint64(op - PUSH0)
			if pc+1+n > uint64(len(code)) {
				return nil, nil, nil
			}
			x := hotRegionNode{op: PUSH32}
			x.value.SetBytes(code[pc+1 : pc+1+n])
			st = append(st, hotRegionInputs+len(q.nodes))
			q.nodes = append(q.nodes, x)
			next += n
		case op >= DUP1 && op <= DUP16:
			st = append(st, st[len(st)-int(op-DUP1)-1])
		case op >= SWAP1 && op <= SWAP16:
			i, j := len(st)-1, len(st)-int(op-SWAP1)-2
			st[i], st[j] = st[j], st[i]
		case op == POP:
			st = st[:len(st)-1]
		case op == JUMPDEST:
		case op == JUMP:
			target, ok := known(st[len(st)-1])
			if !ok || target >= uint64(len(code)) || OpCode(code[target]) != JUMPDEST || !bitmap.codeSegment(target) {
				return nil, nil, nil
			}
			next = target
			st = st[:len(st)-1]
		default:
			return nil, nil, nil
		}
		pc = next
	}
}

func joinRegionTail(prefix *hotRegion, st []int, b *hotRegion) *hotRegion {
	q := *prefix
	q.nodes = append([]hotRegionNode(nil), prefix.nodes...)
	delta := len(st) - hotRegionInputs
	q.Inputs = max(q.Inputs, b.Inputs-delta)
	q.MaxStack = min(q.MaxStack, b.MaxStack-delta)
	if q.Inputs > hotRegionInputs || q.Inputs < 0 || q.Inputs > q.MaxStack || len(st) < b.Drop {
		return nil
	}
	base := hotRegionInputs + len(q.nodes)
	remap := func(id int) int {
		if id >= hotRegionInputs {
			return base + id - hotRegionInputs
		}
		at := len(st) - hotRegionInputs + id
		if at < 0 || at >= len(st) {
			return -1
		}
		return st[at]
	}
	for _, n := range b.nodes {
		n.args = append([]int(nil), n.args...)
		for i, id := range n.args {
			n.args[i] = remap(id)
			if n.args[i] < 0 {
				return nil
			}
		}
		q.nodes = append(q.nodes, n)
	}
	out := append([]int(nil), st[:len(st)-b.Drop]...)
	for _, id := range b.outputs {
		v := remap(id)
		if v < 0 {
			return nil
		}
		out = append(out, v)
	}
	trim := 0
	for trim < len(out) && trim < hotRegionInputs && out[trim] == trim {
		trim++
	}
	q.Drop = hotRegionInputs - trim
	q.outputs = out[trim:]
	q.Inputs = max(q.Inputs, q.Drop)
	q.Exit = b.Exit
	q.Gas += b.Gas
	q.Instructions += b.Instructions
	q.Plain += b.Plain
	q.Helpers += b.Helpers
	q.pipeline = q.pipeline || b.pipeline
	q.decodes += b.decodes
	q.encodes += b.encodes
	if q.Inputs > hotRegionInputs || q.Inputs > q.MaxStack || len(q.nodes) > 320 {
		return nil
	}
	q.key = q.fingerprint()
	return &q
}

// One bounded pass over an immutable snapshot: no recursive chain growth.
func joinDemandRegions(a *demandPlan, targets map[uint64]*demandPlan, code []byte, table *JumpTable, bitmaps ...BitVec) *demandPlan {
	var bitmap BitVec
	if len(bitmaps) > 0 {
		bitmap = bitmaps[0]
	} else {
		bitmap = codeBitmap(code)
	}
	if a == nil {
		return nil
	}
	var paths []*hotRegion
	var parents []*hotRegion
	changed := false
	for _, r := range a.regions {
		prefix, st, next := regionBridge(r, targets, code, table, bitmap)
		if next == nil {
			paths = append(paths, r)
			parents = append(parents, r)
			continue
		}
		for _, b := range next.regions {
			q := joinRegionTail(prefix, st, b)
			if q == nil {
				return nil
			}
			paths = append(paths, q)
			parents = append(parents, r)
		}
		changed = true
		if len(paths) > 4 {
			return nil
		}
	}
	if !changed || len(paths) > 4 {
		return nil
	}
	p := makeFlowPlan(paths, true, true)
	if p != nil {
		p.joinParents = parents
	}
	return p
}

// Replace existing entries only; the interpreter gains no new dispatch layer.
func (e *EVM) bindRegionJoins(p *abiExecutionPlan, code []byte) {
	e.bindRegionJoinsAt(p, code, nil)
}

// Revisit an already observed predecessor when its successor is learned later.
// Use original unjoined sources on both sides, keeping the same one-pair scope
// as a single offline pass rather than recursively growing longer chains.
func incrementalJoinSources(targets map[uint64]*demandPlan, code []byte, table *JumpTable, bitmap BitVec, selected map[uint64]bool) (map[uint64]*demandPlan, map[uint64]bool) {
	sources := make(map[uint64]*demandPlan, len(targets))
	for pc, p := range targets {
		if p.joinSource != nil {
			p = p.joinSource
		}
		sources[pc] = p
	}
	updates := make(map[uint64]bool, len(selected))
	for pc, yes := range selected {
		updates[pc] = yes
	}
	for pc, p := range sources {
		if updates[pc] {
			continue
		}
		for _, r := range p.regions {
			_, _, next := regionBridge(r, sources, code, table, bitmap)
			if next != nil && len(next.regions) > 0 && selected[next.regions[0].Entry] {
				updates[pc] = true
				break
			}
		}
	}
	return sources, updates
}

func (e *EVM) bindRegionJoinsAt(p *abiExecutionPlan, code []byte, selected map[uint64]bool) {
	if paperAblationLevel == 2 {
		return
	}
	if experimentSkipRejectedAnalysis() {
		return
	}
	a := p.arithmetic
	if a == nil {
		return
	}
	targets := map[uint64]*demandPlan{}
	for pc, id := range a.entryIndex {
		if id == 0 || int(id) >= len(a.computeEntries) {
			continue
		}
		c := a.computeEntries[id]
		if c != nil && c.extension != nil && c.extension.demand != nil {
			targets[uint64(pc)] = c.extension.demand
		}
	}
	bitmap := codeBitmap(code)
	updates := selected
	incremental := tradeProgressive && tradeJoinUpdates && selected != nil
	if incremental {
		targets, updates = incrementalJoinSources(targets, code, e.table, bitmap, selected)
	}
	for index := range a.entryIndex {
		pc := uint64(index)
		old := targets[pc]
		if old == nil || updates != nil && !updates[pc] {
			continue
		}
		if e.preparationYield != nil {
			e.preparationYield()
		}
		q := joinDemandRegions(old, targets, code, e.table, bitmap)
		if q == nil {
			continue
		}
		id := a.entryIndex[pc]
		c := *a.computeEntries[id]
		x := *c.extension
		q.fallback = x.demand
		q.joinSource = old
		if old.joinSource != nil {
			q.joinSource = old.joinSource
		}
		if incremental {
			if !selected[pc] {
				progressiveJoinRevisits.Add(1)
			}
		}
		x.demand = q
		c.extension = &x
		a.computeEntries[id] = &c
	}
}
