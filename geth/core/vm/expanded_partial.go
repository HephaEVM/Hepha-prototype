package vm

var expandedPartialGenerated = map[string]hotRegionFunc{}

// Only instruction boundaries. An unsupported effect stays in the interpreter;
// its successor may start a new independent straight-line summary.
func partialStarts(code []byte, expanded bool) map[uint64]bool {
	starts := map[uint64]bool{}
	if expanded && len(code) > 0 {
		starts[0] = true
	}
	for pc := uint64(0); pc < uint64(len(code)); {
		op := OpCode(code[pc])
		if op == JUMPDEST {
			starts[pc] = true
		}
		pc++
		if op >= PUSH1 && op <= PUSH32 {
			pc += uint64(op - PUSH0)
			continue
		}
		if !expanded || pc >= uint64(len(code)) {
			continue
		}
		if op == STOP || op == RETURN || op == REVERT || op == INVALID || op == SELFDESTRUCT || op == JUMP {
			continue
		}
		if !hotRegionSupported(op) && op != MLOAD && op != MSTORE {
			starts[pc] = true
		}
	}
	return starts
}
