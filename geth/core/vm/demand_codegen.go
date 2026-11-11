package vm

import (
	"bytes"
	"fmt"
	"go-eth1.16.5-evm/regionexec"
	"go/format"
	"strings"
)

func emitDemand(orig *hotRegion, memory bool) (string, []regionexec.Helper, error) {
	return emitPreparedDemand(&preparedFlow{original: orig}, memory)
}
func emitPreparedDemand(prepared *preparedFlow, memory bool) (string, []regionexec.Helper, error) {
	r, bounds := prepared.optimized(false)
	forward, dead := map[int]int{}, map[int]bool{}
	if memory {
		forward, dead = memoryDemand(r)
	}
	var s bytes.Buffer
	var helpers []regionexec.Helper
	used := make([]bool, hotRegionInputs+len(r.nodes))
	for _, id := range r.outputs {
		used[id] = true
	}
	for i := len(r.nodes) - 1; i >= 0; i-- {
		n := r.nodes[i]
		if n.helper != nil || (r.pipeline && (n.op == JUMPI || n.op == MSTORE || n.op == MLOAD)) {
			used[hotRegionInputs+i] = true
		}
		if f, ok := forward[i]; ok {
			used[f] = true
		}
		if used[hotRegionInputs+i] {
			for _, a := range n.args {
				used[a] = true
			}
		}
	}
	fmt.Fprintf(&s, "func Exec(s *regionexec.Frame) bool {\nn:=len(s.Stack)\nif n<%d || n>%d || s.Gas<%d {return false}\n", r.Inputs, r.MaxStack, r.Gas)
	writes := 0
	if r.pipeline {
		for _, node := range r.nodes {
			if node.op == MSTORE {
				writes++
			}
		}
		fmt.Fprintf(&s, "var memEnd uint64\nvar writes [%d]regionexec.Write\nnw:=0\n", writes)
		fmt.Fprintf(&s, "skip:=[%d]bool{", writes)
		for i, v := range dead {
			if v {
				fmt.Fprintf(&s, "%d:true,", i)
			}
		}
		fmt.Fprint(&s, "}\n")
	}
	for i := 0; i < hotRegionInputs; i++ {
		if used[i] {
			fmt.Fprintf(&s, "v%d:=s.Stack[n-%d]\n", i, hotRegionInputs-i)
		}
	}
	for i, node := range r.nodes {
		id := hotRegionInputs + i
		if !used[id] {
			continue
		}
		v := fmt.Sprintf("v%d", id)
		arg := func(k int) string { return fmt.Sprintf("v%d", node.args[k]) }
		if r.pipeline && node.op == JUMPI {
			if node.nonzero {
				fmt.Fprintf(&s, "if %s.IsZero(){return false}\n", arg(0))
			} else {
				fmt.Fprintf(&s, "if !%s.IsZero(){return false}\n", arg(0))
			}
			continue
		}
		if r.pipeline && (node.op == MLOAD || node.op == MSTORE) {
			fmt.Fprintf(&s, "if !%s.IsUint64() || %s.Uint64()>regionexec.MaxMemory-32{return false}\no%d:=%s.Uint64()\nmemEnd=max(memEnd,o%d+32)\n", arg(0), arg(0), id, arg(0), id)
			if node.op == MLOAD {
				if f, ok := forward[i]; ok {
					fmt.Fprintf(&s, "%s:=v%d\n", v, f)
				} else {
					fmt.Fprintf(&s, "var %s uint256.Int\nregionexec.Load(&%s,s.Memory,o%d,writes[:nw])\n", v, v, id)
				}
				fmt.Fprintf(&s, "_=%s\n", v)
			} else {
				fmt.Fprintf(&s, "writes[nw]=regionexec.Write{o%d,%s}\nnw++\n", id, arg(1))
			}
			continue
		}
		if h := node.helper; h != nil {
			if h.domain.Bits != 256 {
				if !h.domain.Signed {
					fmt.Fprintf(&s, "if %s.BitLen()>%d||%s.BitLen()>%d{return false}\n", arg(0), h.domain.Bits, arg(1), h.domain.Bits)
				} else {
					fmt.Fprintf(&s, "var ca%d,cb%d uint256.Int\nix%d:=uint256.Int{%d}\nca%d.ExtendSign(&%s,&ix%d);cb%d.ExtendSign(&%s,&ix%d)\nif ca%d!=%s||cb%d!=%s{return false}\n", id, id, id, h.domain.Bits/8-1, id, arg(0), id, id, arg(1), id, id, arg(0), id, arg(1))
				}
			}
			fmt.Fprintf(&s, "p%d,%s:=s.Helpers[%d](%s,%s)\n", id, v, len(helpers), arg(0), arg(1))
			helpers = append(helpers, regionexec.Helper(h.syncHot.native))
			var cond []string
			for p, cost := range h.syncHot.costs {
				if cost != syncHotFallback && cost == h.syncHot.admissionGas {
					cond = append(cond, fmt.Sprintf("p%d!=%d", id, p))
				}
			}
			if len(cond) == 0 {
				return "", nil, fmt.Errorf("no admissible helper path")
			}
			fmt.Fprintf(&s, "if %s{return false}\n_=%s\n", strings.Join(cond, "&&"), v)
			continue
		}
		fmt.Fprintf(&s, "var %s uint256.Int\n", v)
		if node.op == PUSH32 {
			fmt.Fprintf(&s, "%s=uint256.Int{%d,%d,%d,%d}\n", v, node.value[0], node.value[1], node.value[2], node.value[3])
			continue
		}
		if len(bounds) > id && emitDemandNarrow(&s, node, id, bounds) {
			continue
		}
		binary := map[OpCode]string{ADD: "Add", SUB: "Sub", MUL: "Mul", DIV: "Div", SDIV: "SDiv", MOD: "Mod", SMOD: "SMod", AND: "And", OR: "Or", XOR: "Xor"}
		if method := binary[node.op]; method != "" {
			fmt.Fprintf(&s, "%s.%s(&%s,&%s)\n", v, method, arg(0), arg(1))
			continue
		}
		compare := map[OpCode]string{EQ: "Eq", LT: "Lt", GT: "Gt", SLT: "Slt", SGT: "Sgt"}
		if method := compare[node.op]; method != "" {
			fmt.Fprintf(&s, "if %s.%s(&%s){%s.SetOne()}\n", arg(0), method, arg(1), v)
			continue
		}
		switch node.op {
		case ADDMOD, MULMOD:
			method := "AddMod"
			if node.op == MULMOD {
				method = "MulMod"
			}
			fmt.Fprintf(&s, "%s.%s(&%s,&%s,&%s)\n", v, method, arg(0), arg(1), arg(2))
		case ISZERO:
			fmt.Fprintf(&s, "if %s.IsZero(){%s.SetOne()}\n", arg(0), v)
		case NOT:
			fmt.Fprintf(&s, "%s.Not(&%s)\n", v, arg(0))
		case BYTE:
			fmt.Fprintf(&s, "%s=%s\n%s.Byte(&%s)\n", v, arg(1), v, arg(0))
		case SIGNEXTEND:
			fmt.Fprintf(&s, "%s.ExtendSign(&%s,&%s)\n", v, arg(1), arg(0))
		case SHL, SHR:
			method := "Lsh"
			if node.op == SHR {
				method = "Rsh"
			}
			fmt.Fprintf(&s, "if %s.LtUint64(256){%s.%s(&%s,uint(%s.Uint64()))}\n", arg(0), v, method, arg(1), arg(0))
		case SAR:
			fmt.Fprintf(&s, "shift%d:=uint(256)\nif %s.LtUint64(256){shift%d=uint(%s.Uint64())}\n%s.SRsh(&%s,shift%d)\n", id, arg(0), id, arg(0), v, arg(1), id)
		case CALLDATASIZE:
			fmt.Fprintf(&s, "%s.SetUint64(uint64(len(s.Input)))\n", v)
		case CALLDATALOAD:
			fmt.Fprintf(&s, "if %s.IsUint64(){offset%d:=%s.Uint64();if offset%d<uint64(len(s.Input)){var word%d [32]byte;copy(word%d[:],s.Input[offset%d:min(offset%d+32,uint64(len(s.Input)))]);%s.SetBytes(word%d[:])}}\n", arg(0), id, arg(0), id, id, id, id, id, v, id)
		default:
			return "", nil, fmt.Errorf("missing emitter %s", node.op)
		}
	}
	if r.pipeline {
		fmt.Fprintf(&s, "memSize,memFee,totalGas,ok:=regionexec.MemoryCost(s,memEnd,%d)\nif !ok{return false}\nregionexec.Resize(s,memSize)\ns.MemoryGas=memFee\nfor i:=0;i<nw;i++{if !skip[i]{writes[i].Value.PutUint256(s.Memory[writes[i].Offset:])}}\n", r.Gas)
	}
	fmt.Fprintf(&s, "base:=n-%d\nend:=base+%d\nif end>cap(s.Stack){s.Stack=append(s.Stack,make([]uint256.Int,end-n)...)}\ns.Stack=s.Stack[:end]\n", r.Drop, len(r.outputs))
	for i, id := range r.outputs {
		fmt.Fprintf(&s, "s.Stack[base+%d]=v%d\n", i, id)
	}
	if r.pipeline {
		fmt.Fprint(&s, "s.Gas-=totalGas\ns.PC=s.Exit\nreturn true\n}\n")
	} else {
		fmt.Fprintf(&s, "s.Gas-=%d\ns.PC=s.Exit\nreturn true\n}\n", r.Gas)
	}

	raw, err := format.Source(append([]byte("package main\n"), s.Bytes()...))
	if err != nil {
		return "", nil, err
	}
	return strings.TrimPrefix(string(raw), "package main\n"), helpers, nil
}
func emitDemandNarrow(s *bytes.Buffer, n hotRegionNode, id int, ranges []hotInterval) bool {
	if len(n.args) != 2 {
		return false
	}
	a, b := n.args[0], n.args[1]
	width := ranges[id].hi.BitLen()
	if width <= 64 && (n.op == AND || ((n.op == ADD || n.op == SUB || n.op == MUL) && ranges[a].hi.BitLen() <= 64 && ranges[b].hi.BitLen() <= 64)) {
		op := map[OpCode]string{AND: "&", ADD: "+", SUB: "-", MUL: "*"}[n.op]
		fmt.Fprintf(s, "v%d[0]=v%d[0]%sv%d[0]\n", id, a, op, b)
		return true
	}
	if n.op == MUL && ranges[a].hi.BitLen() <= 64 && ranges[b].hi.BitLen() <= 64 {
		fmt.Fprintf(s, "v%d[1],v%d[0]=bits.Mul64(v%d[0],v%d[0])\n", id, id, a, b)
		return true
	}
	if n.op == ADD && width <= 128 && ranges[a].hi.BitLen() <= 128 && ranges[b].hi.BitLen() <= 128 {
		fmt.Fprintf(s, "var carry%d uint64\nv%d[0],carry%d=bits.Add64(v%d[0],v%d[0],0)\nv%d[1],_=bits.Add64(v%d[1],v%d[1],carry%d)\n", id, id, id, a, b, id, a, b, id)
		return true
	}
	return false
}
