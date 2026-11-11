package vm

import (
	"fmt"
	"github.com/holiman/uint256"
)

// Memory forwarding preserves every original bound/gas check and the high-water
// mark. Dead stores are omitted only from final publication; staged reads retain
// the original ordered writes. Unknown aliasing always stops forwarding.
func memoryDemand(r *hotRegion) (map[int]int, map[int]bool) {
	forward := map[int]int{}
	dead := map[int]bool{}
	interns := map[string]int{}
	intern := func(k string) int {
		if id := interns[k]; id != 0 {
			return id
		}
		id := len(interns) + 1
		interns[k] = id
		return id
	}
	keys := make([]int, hotRegionInputs+len(r.nodes))
	constants := map[int]uint256.Int{}
	for i := 0; i < hotRegionInputs; i++ {
		keys[i] = intern(fmt.Sprintf("in%d", i))
	}
	type access struct{ i, addr, value, write int }
	var history []access
	last := map[int]int{}
	nw := 0
	disjoint := func(a, b int) bool {
		x, xok := constants[a]
		y, yok := constants[b]
		if !xok || !yok || !x.IsUint64() || !y.IsUint64() {
			return false
		}
		u, v := x.Uint64(), y.Uint64()
		if u > abiMaxMemory-32 || v > abiMaxMemory-32 {
			return false
		}
		return u+32 <= v || v+32 <= u
	}
	for i, n := range r.nodes {
		id := hotRegionInputs + i
		if n.op == PUSH32 {
			keys[id] = intern(n.value.Hex())
			constants[id] = n.value
		} else {
			keys[id] = intern(fmt.Sprintf("n%d", i))
		}
		if n.helper == nil && n.op != MLOAD && n.op != MSTORE && n.op != JUMPI && n.op != PUSH32 {
			k := fmt.Sprintf("op%d", n.op)
			for _, a := range n.args {
				k += fmt.Sprintf("(%d)", keys[a])
			}
			keys[id] = intern(k)
		}
		if n.op == MLOAD {
			addr := n.args[0]
			for j := len(history) - 1; j >= 0; j-- {
				p := history[j]
				if keys[p.addr] == keys[addr] {
					forward[i] = p.value
					keys[id] = keys[p.value]
					break
				}
				if p.write >= 0 && !disjoint(p.addr, addr) {
					break
				}
			}
			history = append(history, access{i, addr, id, -1})
		} else if n.op == MSTORE {
			addr := n.args[0]
			k := keys[addr]
			if prev, ok := last[k]; ok {
				dead[prev] = true
			}
			last[k] = nw
			history = append(history, access{i, addr, n.args[1], nw})
			nw++
		}
	}
	return forward, dead
}
