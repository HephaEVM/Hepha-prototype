//go:build hepha_export

package vm

import (
	"github.com/ethereum/go-ethereum/crypto"
	"sort"
)

// ExportEvmoneEffects reuses the semantic chain proof. Every state access is
// retained, including loads of a previously written slot. Only its address is
// reused. Checkpoints describe the exact original continuation after effects.
func ExportEvmoneEffects(code []byte, engine *ArithmeticEngine, sites map[uint64]bool) ([]EvmoneRegion, error) {
	b := engine.bind(code, arithmeticCacheKey{hash: crypto.Keccak256Hash(code), rules: &pragueInstructionSet, reduced: true, compute: 1 | 4 | uint8(HashPreparationDirect)<<3})
	if b.err != nil {
		return nil, b.err
	}
	starts := partialStarts(code, true)
	for _, pc := range summaryCandidateEntries(code) {
		starts[pc] = true
	}
	bitmap := codeBitmap(code)
	out := []EvmoneRegion{}
	for pc := range starts {
		if sites != nil && !sites[pc] {
			continue
		}
		p := bindEffectChainThreshold(code, pc, b, &pragueInstructionSet, bitmap, 1)
		if p == nil {
			continue
		}
		nodes := []any{}
		actions := []any{}
		events := [][2]int{}
		writes := [][2]int{}
		for _, v := range p.nodes {
			op := v.op.String()
			if v.op == chainInput {
				op = "input"
			}
			if v.op == chainConstant {
				op = "constant"
			}
			if v.op == 0xfe {
				op = "effect"
			}
			nodes = append(nodes, map[string]any{"op": op, "a": v.a, "b": v.b, "depth": v.depth, "value": v.c.Hex()})
		}
		valid := true
		for _, a := range p.actions {
			op := a.op.String()
			gas := a.gas
			// Geth places the warm SLOAD cost and all fixed LOG cost in dynamicGas;
			// evmone places these in its opcode table, before the primitive.
			if a.op == SLOAD {
				gas += 100
			}
			if a.op >= LOG0 && a.op <= LOG4 {
				gas += uint64(a.op-LOG0+1) * 375
			}
			x := map[string]any{"op": op, "pc": a.pc, "start": a.start, "gas": gas, "args": a.args, "result": a.result, "reuse": a.reuse, "before": a.before, "checkpoint": a.checkpoint, "start_step": a.startStep, "op_step": a.opStep}
			if a.op == chainHelper {
				h := cppArithmetic(a.helper)
				if h == nil {
					valid = false
					break
				}
				x["op"] = "helper"
				x["helper"] = h
				x["helper_gas"] = a.helper.syncHot.admissionGas
			}
			actions = append(actions, x)
		}
		if !valid {
			continue
		}
		for _, v := range p.events {
			events = append(events, [2]int{v.node, v.action})
		}
		for _, v := range p.writes {
			writes = append(writes, [2]int{v.offset, v.id})
		}
		native := map[string]any{"family": "effects", "entry": p.entry, "exit": p.exit, "inputs": p.minStack, "peak": p.peak, "delta": p.delta, "nodes": nodes, "actions": actions, "events": events, "writes": writes, "instructions": len(p.pcs), "hashes": p.hashes, "reused": p.reused, "loads": p.loads, "stores": p.stores, "logs": p.logs}
		out = append(out, EvmoneRegion{Kind: 5, Entry: pc, Stage: 3, Native: native, Evidence: "SSA address identity; ordered effects; exact continuation checkpoints"})
	}
	sort.Slice(out, func(i, j int) bool { return out[i].Entry < out[j].Entry })
	return out, nil
}
