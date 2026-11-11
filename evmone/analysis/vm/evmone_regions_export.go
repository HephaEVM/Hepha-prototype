//go:build hepha_export

package vm

// This offline export reuses the existing miners and admission proofs. It does
// not introduce new matching rules or change Geth's execution path.
import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	miner "go-eth1.16.5-evm/template_miner"
	"os"
	"sort"
)

type evmoneAffine struct {
	Input int    `json:"input"`
	Add   string `json:"add"`
}
type evmoneClean struct {
	Kind     byte   `json:"kind"`
	Argument string `json:"argument"`
}
type evmoneRead struct {
	Offset evmoneAffine `json:"offset"`
	Memory bool         `json:"memory"`
	Clean  evmoneClean  `json:"clean"`
}
type evmoneWrite struct {
	Offset evmoneAffine `json:"offset"`
	Input  int          `json:"input"`
	Clean  evmoneClean  `json:"clean"`
}
type evmoneBound struct {
	End   int    `json:"end"`
	Start int    `json:"start"`
	Size  string `json:"size"`
}
type EvmoneRegion struct {
	Stage            int            `json:"stage,omitempty"`
	Native           any            `json:"native,omitempty"`
	Kind             int            `json:"kind"`
	Entry            uint64         `json:"entry"`
	Exit             uint64         `json:"exit"`
	Inputs           int            `json:"inputs"`
	Peak             int            `json:"peak"`
	Cost             uint64         `json:"cost"`
	LoopCost         uint64         `json:"loop_cost"`
	Instructions     uint64         `json:"instructions"`
	LoopInstructions uint64         `json:"loop_instructions"`
	Bounds           []evmoneBound  `json:"bounds"`
	Reads            []evmoneRead   `json:"reads"`
	Writes           []evmoneWrite  `json:"writes"`
	Outputs          []evmoneAffine `json:"outputs"`
	ReadOutputs      []int          `json:"read_outputs"`
	Evidence         string         `json:"evidence"`
}

func ExportEvmoneRegions(code []byte, libraryPath string) ([]EvmoneRegion, error) {
	raw, err := os.ReadFile(libraryPath)
	if err != nil {
		return nil, err
	}
	var library miner.ABIRegionLibrary
	if err = json.Unmarshal(raw, &library); err != nil {
		return nil, err
	}
	if library.Version != 1 {
		return nil, fmt.Errorf("unsupported ABI library")
	}
	return exportEvmoneRegions(code, &library, nil), nil
}

// Share the same ABI proof with composition; callers own the immutable library.
func exportEvmoneRegions(code []byte, library *miner.ABIRegionLibrary, pipeline map[uint64]*abiRegionPrepared) []EvmoneRegion {
	result := []EvmoneRegion{}
	used := map[uint64]bool{}
	bitmap := codeBitmap(code)
	for pc := uint64(0); pc < uint64(len(code)); pc++ {
		op := OpCode(code[pc])
		if op >= PUSH1 && op <= PUSH32 {
			pc += uint64(op - PUSH0)
			continue
		}
		if op != JUMPDEST {
			continue
		}
		if p := bindWordLoop(code, pc, &pragueInstructionSet, bitmap); p != nil {
			result = append(result, EvmoneRegion{Kind: 4, Entry: pc, Exit: p.exit, Inputs: 4, Peak: 7, Cost: p.guardGas, LoopCost: p.bodyGas, Instructions: 7, LoopInstructions: 19, Evidence: "bindWordLoop: CFG/stride/stack/effects/Prague gas verified"})
			used[pc] = true
		}
	}
	affine := func(a abiAffine) evmoneAffine { return evmoneAffine{a.input, a.add.Hex()} }
	clean := func(c abiClean) evmoneClean { return evmoneClean{c.kind, c.argument.Hex()} }
	for _, r := range miner.DiscoverABIRegions(code, library) {
		excluded := used[r.Entry] || r.Inputs < 1 || r.Inputs > 32
		if pipeline == nil && excluded {
			continue
		}
		p, err := prepareABIRegion(r, code, &pragueInstructionSet)
		if err != nil {
			continue
		}
		n := prepareABINative(p)
		if n == nil {
			continue
		}
		if pipeline != nil && (n.kind == 1 || n.kind == 2) {
			p.native = n
			n.prepareFixed()
			n.prepareFixedEncode()
			pipeline[r.Entry] = p
		}
		// Standalone export has stricter admission than ABI composition. Reusing
		// its preparation must not remove a summary from the composition map.
		if excluded {
			continue
		}
		v := EvmoneRegion{Kind: int(n.kind), Entry: r.Entry, Inputs: r.Inputs, Peak: n.peak, Cost: n.cost, LoopCost: n.loopCost, Instructions: n.instructions, LoopInstructions: n.loopInstructions, Evidence: r.ID}
		for _, b := range n.bounds {
			v.Bounds = append(v.Bounds, evmoneBound{b.end, b.start, b.size.Hex()})
		}
		for _, x := range n.reads {
			v.Reads = append(v.Reads, evmoneRead{affine(x.offset), x.memory, clean(x.clean)})
		}
		for _, x := range n.writes {
			v.Writes = append(v.Writes, evmoneWrite{affine(x.offset), x.input, clean(x.clean)})
		}
		for _, x := range n.outputs {
			v.Outputs = append(v.Outputs, affine(x))
		}
		v.ReadOutputs = n.readOutputs
		result = append(result, v)
	}
	sort.Slice(result, func(i, j int) bool { return result[i].Entry < result[j].Entry })
	return result
}

// Export a single implementation-generated fixture without a second code parser.
func EvmoneWordLoopFixture() string {
	return hex.EncodeToString([]byte{byte(JUMPDEST), byte(DUP4), byte(DUP2), byte(LT), byte(ISZERO), byte(PUSH2), 0, 24, byte(JUMPI), byte(DUP2), byte(DUP2), byte(ADD), byte(MLOAD), byte(DUP4), byte(DUP3), byte(ADD), byte(MSTORE), byte(PUSH1), 32, byte(ADD), byte(PUSH2), 0, 0, byte(JUMP), byte(JUMPDEST), byte(STOP)})
}
