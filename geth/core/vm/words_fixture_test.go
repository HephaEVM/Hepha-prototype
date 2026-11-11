package vm

import (
	"github.com/holiman/uint256"
)

func makeWords(n int) []uint256.Int {
	w := make([]uint256.Int, n)
	for i := range w {
		w[i].SetUint64(128)
	}
	return w
}
