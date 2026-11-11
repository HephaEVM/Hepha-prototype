package vm

import (
	"github.com/holiman/uint256"
	miner "go-eth1.16.5-evm/template_miner"
)

// Import a re-proved ABI success path into the same value graph as ordinary
// computation and checked arithmetic. Labels alone never authorize replacement.
// Restrict admission to existing native scalar/tuple decode and word encode
// recipes; other ABI graphs keep their original execution.
func importPipelineABI(r *hotRegion, p *abiRegionPrepared, inputs []int, known map[int]uint256.Int) ([]int, bool) {
	if paperAblationLevel == 2 {
		return nil, false
	}
	if p.native == nil || (p.native.kind != 1 && p.native.kind != 2) || len(p.paths) != 1 || len(p.graph) != 0 {
		return nil, false
	}
	var path *miner.ABIRegionPath
	for i := range p.region.Paths {
		if p.region.Paths[i].Outcome == "return" {
			path = &p.region.Paths[i]
			break
		}
	}
	if path == nil {
		return nil, false
	}
	var nodes []hotRegionNode
	ids := map[int]int{}
	base := hotRegionInputs + len(r.nodes)
	add := func(n hotRegionNode) int { id := base + len(nodes); nodes = append(nodes, n); return id }
	var value func(int) (int, bool)
	value = func(id int) (int, bool) {
		if v, ok := ids[id]; ok {
			return v, true
		}
		if id < 0 || id >= len(path.Nodes) {
			return 0, false
		}
		n := path.Nodes[id]
		if n.Op == "input" {
			if n.Input < 0 || n.Input >= len(inputs) {
				return 0, false
			}
			ids[id] = inputs[n.Input]
			return inputs[n.Input], true
		}
		if n.Op == "read" {
			return 0, false
		}
		x := hotRegionNode{}
		if n.Op == "const" {
			x.op = PUSH32
			if x.value.SetFromHex(n.Value) != nil {
				return 0, false
			}
		} else {
			x.op = StringToOp(n.Op)
			if !hotRegionPure(x.op) {
				return 0, false
			}
			for _, a := range n.Args {
				v, ok := value(a)
				if !ok {
					return 0, false
				}
				x.args = append(x.args, v)
			}
		}
		v := add(x)
		ids[id] = v
		return v, true
	}
	for _, a := range path.Actions {
		n := hotRegionNode{op: StringToOp(a.Op), nonzero: a.Nonzero}
		if a.Op == "guard" {
			n.op = JUMPI
		}
		switch n.op {
		case JUMPI, CALLDATALOAD, CALLDATASIZE, MLOAD, MSTORE:
		default:
			return nil, false
		}
		for _, arg := range a.Args {
			v, ok := value(arg)
			if !ok {
				return nil, false
			}
			n.args = append(n.args, v)
		}
		v := add(n)
		if n.op == CALLDATALOAD || n.op == CALLDATASIZE || n.op == MLOAD {
			ids[a.Result] = v
		}
	}
	var outputs []int
	for _, o := range path.Outputs {
		v, ok := value(o)
		if !ok {
			return nil, false
		}
		outputs = append(outputs, v)
	}
	for i, n := range nodes {
		if n.op == PUSH32 {
			known[base+i] = n.value
		}
	}
	r.nodes = append(r.nodes, nodes...)
	return outputs, true
}

func discoverPipelineABI(code []byte, library *miner.ABIRegionLibrary) map[uint64]*abiRegionPrepared {
	out := map[uint64]*abiRegionPrepared{}
	for _, r := range miner.DiscoverABIRegions(code, library) {
		p, err := prepareABIRegion(r, code, &pragueInstructionSet)
		if err != nil {
			continue
		}
		p.native = prepareABINative(p)
		if p.native == nil || (p.native.kind != 1 && p.native.kind != 2) {
			continue
		}
		p.native.prepareFixed()
		p.native.prepareFixedEncode()
		out[r.Entry] = p
	}
	return out
}

type pipelineWrite struct {
	offset uint64
	value  uint256.Int
}

// Exact virtual memory read, including partial overlaps and last-write-wins.
// No live writes or resizing occur before the final gas/guard admission.
func pipelineLoad(dst *uint256.Int, memory *Memory, offset uint64, writes []pipelineWrite) {
	var word [32]byte
	if offset < uint64(memory.Len()) {
		copy(word[:], memory.store[offset:min(offset+32, uint64(memory.Len()))])
	}
	for i := range writes {
		w := &writes[i]
		start, end := max(offset, w.offset), min(offset+32, w.offset+32)
		if start < end {
			b := w.value.Bytes32()
			copy(word[start-offset:end-offset], b[start-w.offset:end-w.offset])
		}
	}
	dst.SetBytes(word[:])
}
