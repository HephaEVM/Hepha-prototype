//go:build !paper_count

package vm

const paperCountInstructions = false

var paperInterpreted uint64

func PaperInterpretedInstructions() uint64 { return 0 }
