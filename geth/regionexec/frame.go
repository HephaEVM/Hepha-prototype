// Package regionexec is the narrow ABI between the EVM region
// compiler and locally generated Go plugins. It has no chain/database access.
package regionexec

import "github.com/holiman/uint256"

type Helper func(uint256.Int, uint256.Int) (int, uint256.Int)
type Frame struct {
	Helpers                  []Helper
	BranchHelpers            [2][]Helper
	Stack                    []uint256.Int
	Input, Memory            []byte
	Gas, MemoryGas, Exit, PC uint64
	Exits                    [2]uint64
	Path                     uint8
}
type Executor func(*Frame) bool
type Write struct {
	Offset uint64
	Value  uint256.Int
}

const MaxMemory = uint64(0x1FFFFFFFE0)

func MemoryCost(f *Frame, end, cost uint64) (size, fee, total uint64, ok bool) {
	if end > MaxMemory || cost > f.Gas {
		return
	}
	size = uint64(len(f.Memory))
	fee = f.MemoryGas
	if end > size {
		size = (end + 31) / 32 * 32
		words := size / 32
		fee = words*3 + words*words/512
		if fee < f.MemoryGas {
			return
		}
		extra := fee - f.MemoryGas
		if extra > f.Gas-cost {
			return
		}
		cost += extra
	}
	return size, fee, cost, true
}
func Resize(f *Frame, size uint64) {
	if size > uint64(len(f.Memory)) {
		f.Memory = append(f.Memory, make([]byte, size-uint64(len(f.Memory)))...)
	}
}
func Load(dst *uint256.Int, memory []byte, offset uint64, writes []Write) {
	var word [32]byte
	if offset < uint64(len(memory)) {
		copy(word[:], memory[offset:min(offset+32, uint64(len(memory)))])
	}
	for _, w := range writes {
		lo, hi := max(offset, w.Offset), min(offset+32, w.Offset+32)
		if lo < hi {
			v := w.Value.Bytes32()
			copy(word[lo-offset:hi-offset], v[lo-w.Offset:hi-w.Offset])
		}
	}
	dst.SetBytes(word[:])
}
