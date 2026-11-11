package vm

import "github.com/holiman/uint256"

// Keep the uncommon large permutation buffer out of try's stack frame. Most
// preambles preserve only a few caller slots. Sources are saved before any
// destination is overwritten, including cycles, constants and repeated inputs.
func restoreArithmeticPrefix(stack *Stack, input []uint256.Int, r *arithmeticRegion) {
	if len(r.prefix) > 8 {
		restoreLargeArithmeticPrefix(stack, input, r)
		return
	}
	var saved [8]uint256.Int
	for i, source := range r.prefix {
		saved[i] = source.get(input)
	}
	stack.data = stack.data[:stack.len()-r.inputCount]
	stack.data = append(stack.data, saved[:len(r.prefix)]...)
}

func restoreLargeArithmeticPrefix(stack *Stack, input []uint256.Int, r *arithmeticRegion) {
	var saved [64]uint256.Int
	for i, source := range r.prefix {
		saved[i] = source.get(input)
	}
	stack.data = stack.data[:stack.len()-r.inputCount]
	stack.data = append(stack.data, saved[:len(r.prefix)]...)
}
