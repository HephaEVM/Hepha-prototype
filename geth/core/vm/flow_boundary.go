package vm

import "github.com/holiman/uint256"

func buildProtectedFlowPath(b *boundArithmetic, code []byte, table *JumpTable, start uint64, choices uint, maxBranches int, bitmaps ...BitVec) *hotRegion {
	return buildFlowPathWithABI(b, code, table, start, choices, maxBranches, nil, nil, bitmaps...)
}

// Optional summaries and protected entries are shared with the offline backend.
// Existing Geth callers retain the original boundary policy.
func buildFlowPathWithABI(b *boundArithmetic, code []byte, table *JumpTable, start uint64, choices uint, maxBranches int, abi map[uint64]*abiRegionPrepared, barriers map[uint64]bool, bitmaps ...BitVec) *hotRegion {
	stopAtHelper := paperAblationLevel == 2
	limit := ^uint64(0)
	branches := 0
	r := &hotRegion{Entry: start}
	r.pipeline = true
	stack := make([]int, hotRegionInputs, 64)
	for i := range stack {
		stack[i] = i
	}
	known := map[int]uint256.Int{}
	visited := map[uint64]bool{}
	bitmap := flowBitmap(code, bitmaps)
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
		// Binding-time boundary: do not consume another fused plan's entry.
		if pc != start && (flowFusionBoundary(b, pc) || barriers[pc]) {
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
		if (!hotRegionSupported(op) && !memoryOp && op != JUMPI) || ins == nil || ins.undefined || (!memoryOp && (ins.dynamicGas != nil || ins.memorySize != nil)) || len(stack) < ins.minStack {
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
		case op == JUMPI:
			if branches >= maxBranches {
				goto done
			}
			choice := choices&(1<<uint(branches)) != 0
			w, ok := known[stack[len(stack)-1]]
			if !ok || !jumpOK(w) {
				goto done
			}
			add(hotRegionNode{op: JUMPI, args: []int{stack[len(stack)-2]}, nonzero: choice})
			stack = stack[:len(stack)-2]
			branches++
			if choice {
				next = w.Uint64()
			}
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
	if branches == 0 || pc == start || (!stopAtHelper && (r.Instructions < 16 || r.Plain+r.Helpers < 2)) {
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

// Only already-bound effect chains are protected; no contract allowlist and
// no runtime lookup. A future composite executor may explicitly absorb them.
func flowFusionBoundary(b *boundArithmetic, pc uint64) bool {
	if b == nil || pc >= uint64(len(b.entryIndex)) {
		return false
	}
	id := b.entryIndex[pc]
	if id == 0 || int(id) >= len(b.computeEntries) {
		return false
	}
	r := b.computeEntries[id]
	return r != nil && r.address != nil && r.address.chain != nil
}
