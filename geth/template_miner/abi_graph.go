package templateminer

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/core/vm"
)

// Cyclic ABI programs are lowered by basic block. Abstract jump identities are
// retained across blocks; data values are widened, so dynamic loops converge.
// Every original conditional edge and memory effect remains in the graph.
type ABIRegionBlock struct {
	Path      ABIRegionPath `json:"path"`
	Next      int           `json:"next"`
	Else      int           `json:"else"`
	Condition int           `json:"condition"`
}
type abiAbstract struct {
	kind   byte
	value  string
	origin uint64
}
type abiBlockEntry struct {
	pc    uint64
	stack []abiAbstract
}

func AnalyzeABIGraph(p *Program, entry uint64, inputs int) (*ABIRegion, error) {
	if inputs < 1 || inputs > 18 || p.ByPC[entry].Op != vm.JUMPDEST {
		return nil, fmt.Errorf("graph entry/arity")
	}
	r := &ABIRegion{Entry: entry, Inputs: inputs}
	jumps := map[uint64]uint64{}
	indexes := map[string]int{}
	var queue []abiBlockEntry
	addBlock := func(pc uint64, stack []abiAbstract) int {
		// Only the caller continuation and directly pushed, valid code targets need
		// identities across blocks. Numeric counters/pointers become unknown.
		widened := make([]abiAbstract, len(stack))
		var key strings.Builder
		fmt.Fprintf(&key, "%d", pc)
		for i, a := range stack {
			if a.kind == 2 {
				widened[i] = a
			} else if a.kind == 1 {
				v, ok := new(big.Int).SetString(a.value[2:], 16)
				if ok && v.IsUint64() && p.ByPC[v.Uint64()].Op == vm.JUMPDEST && p.ByPC[a.origin].Op >= vm.PUSH0 && p.ByPC[a.origin].Op <= vm.PUSH32 {
					widened[i] = a
				}
			}
			fmt.Fprintf(&key, "/%d:%s:%d", widened[i].kind, widened[i].value, widened[i].origin)
		}
		k := key.String()
		if i, ok := indexes[k]; ok {
			return i
		}
		i := len(queue)
		indexes[k] = i
		queue = append(queue, abiBlockEntry{pc, widened})
		return i
	}
	initial := make([]abiAbstract, inputs)
	initial[0].kind = 2
	addBlock(entry, initial)
	total := 0
	returned := false
	for bi := 0; bi < len(queue); bi++ {
		if len(queue) > 192 || total > 8192 {
			return nil, fmt.Errorf("graph bound")
		}
		begin := queue[bi]
		s := abiRegionState{pc: begin.pc}
		abstract := map[int]abiAbstract{}
		for i, a := range begin.stack {
			id := s.node(ABIValue{Op: "input", Input: i})
			s.push(id)
			abstract[id] = a
		}
		block := ABIRegionBlock{Next: -1, Else: -1, Condition: -1}
		finish := func(pc uint64) int {
			as := make([]abiAbstract, len(s.stack))
			for i, id := range s.stack {
				as[i] = abstract[id]
			}
			return addBlock(pc, as)
		}
		for {
			total++
			if len(s.path.PCs) > 512 || len(s.path.Nodes) > 512 || len(s.stack) > 128 {
				return nil, fmt.Errorf("block bound")
			}
			ins, ok := p.ByPC[s.pc]
			if !ok {
				return nil, fmt.Errorf("graph instruction boundary")
			}
			s.path.PCs = append(s.path.PCs, s.pc)
			s.pc += uint64(1 + len(ins.Data))
			need := func(n int) bool { return len(s.stack) >= n }
			dest := func(id int) (uint64, error) {
				a := abstract[id]
				if a.kind != 1 {
					return 0, fmt.Errorf("graph unresolved jump")
				}
				v, ok := new(big.Int).SetString(a.value[2:], 16)
				if !ok || !v.IsUint64() || p.ByPC[v.Uint64()].Op != vm.JUMPDEST {
					return 0, fmt.Errorf("graph invalid jump")
				}
				jumps[a.origin] = v.Uint64()
				return v.Uint64(), nil
			}
			switch {
			case ins.Op >= vm.PUSH0 && ins.Op <= vm.PUSH32:
				pc := ins.PC
				v := "0x" + new(big.Int).SetBytes(ins.Data).Text(16)
				id := s.node(ABIValue{Op: "const", Value: v, Origin: &pc})
				s.push(id)
				abstract[id] = abiAbstract{kind: 1, value: v, origin: pc}
			case ins.Op >= vm.DUP1 && ins.Op <= vm.DUP16:
				n := int(ins.Op-vm.DUP1) + 1
				if !need(n) {
					return nil, fmt.Errorf("graph stack underflow")
				}
				s.push(s.stack[len(s.stack)-n])
			case ins.Op >= vm.SWAP1 && ins.Op <= vm.SWAP16:
				n := int(ins.Op-vm.SWAP1) + 1
				if !need(n + 1) {
					return nil, fmt.Errorf("graph stack underflow")
				}
				last := len(s.stack) - 1
				s.stack[last], s.stack[last-n] = s.stack[last-n], s.stack[last]
			case ins.Op == vm.POP:
				if !need(1) {
					return nil, fmt.Errorf("graph stack underflow")
				}
				s.pop()
			case ins.Op == vm.JUMPDEST:
			case abiPureArity(ins.Op) > 0:
				n := abiPureArity(ins.Op)
				if !need(n) {
					return nil, fmt.Errorf("graph stack underflow")
				}
				args := make([]int, n)
				for i := range args {
					args[i] = s.pop()
				}
				id := s.node(ABIValue{Op: ins.Op.String(), Args: args})
				s.push(id)
				// Computed code destinations are deliberately not inferred from data.
			case ins.Op == vm.CALLDATALOAD || ins.Op == vm.MLOAD:
				if !need(1) {
					return nil, fmt.Errorf("graph stack underflow")
				}
				a := s.pop()
				id := s.node(ABIValue{Op: "read"})
				s.path.Actions = append(s.path.Actions, ABIAction{Op: ins.Op.String(), Args: []int{a}, Result: id})
				s.push(id)
			case ins.Op == vm.CALLDATASIZE || ins.Op == vm.MSIZE:
				id := s.node(ABIValue{Op: "read"})
				s.path.Actions = append(s.path.Actions, ABIAction{Op: ins.Op.String(), Result: id})
				s.push(id)
			case ins.Op == vm.MSTORE || ins.Op == vm.MSTORE8 || ins.Op == vm.CALLDATACOPY:
				n := 2
				if ins.Op == vm.CALLDATACOPY {
					n = 3
				}
				if !need(n) {
					return nil, fmt.Errorf("graph stack underflow")
				}
				args := make([]int, n)
				for i := range args {
					args[i] = s.pop()
				}
				s.path.Actions = append(s.path.Actions, ABIAction{Op: ins.Op.String(), Args: args})
			case ins.Op == vm.JUMP:
				if !need(1) {
					return nil, fmt.Errorf("graph stack underflow")
				}
				a := s.pop()
				if abstract[a].kind == 2 {
					s.path.Outcome = "return"
					returned = true
					goto done
				}
				target, err := dest(a)
				if err != nil {
					return nil, err
				}
				block.Next = finish(target)
				s.path.Outcome = "next"
				goto done
			case ins.Op == vm.JUMPI:
				if !need(2) {
					return nil, fmt.Errorf("graph stack underflow")
				}
				a, c := s.pop(), s.pop()
				target, err := dest(a)
				if err != nil {
					return nil, err
				}
				block.Condition = c
				block.Next = finish(target)
				block.Else = finish(s.pc)
				s.path.Outcome = "branch"
				goto done
			case ins.Op == vm.REVERT:
				if !need(2) {
					return nil, fmt.Errorf("graph stack underflow")
				}
				s.path.Outcome = "revert"
				goto done
			default:
				return nil, fmt.Errorf("graph unsupported %s", ins.Op)
			}
		}
	done:
		s.path.Outputs = append([]int(nil), s.stack...)
		block.Path = s.path
		r.Graph = append(r.Graph, block)
	}
	if !returned {
		return nil, fmt.Errorf("graph no return")
	}
	labels := map[uint64]int{}
	for _, block := range r.Graph {
		for _, pc := range block.Path.PCs {
			if _, ok := labels[pc]; !ok {
				labels[pc] = len(labels)
			}
		}
	}
	var normalized []string
	for _, block := range r.Graph {
		normalized = append(normalized, fmt.Sprintf("%s:%d:%d", block.Path.Outcome, block.Next, block.Else))
		for _, pc := range block.Path.PCs {
			in := p.ByPC[pc]
			v := in.Op.String()
			if in.Op >= vm.PUSH0 && in.Op <= vm.PUSH32 {
				if dst, ok := jumps[pc]; ok {
					v = fmt.Sprintf("PUSH L%d", labels[dst])
				} else {
					v = "PUSH 0x" + new(big.Int).SetBytes(in.Data).Text(16)
				}
			}
			normalized = append(normalized, fmt.Sprintf("L%d %s", labels[pc], v))
		}
	}
	b, _ := json.Marshal(struct {
		Inputs int
		Code   []string
	}{inputs, normalized})
	h := sha256.Sum256(b)
	r.ID = hex.EncodeToString(h[:])
	return r, nil
}
