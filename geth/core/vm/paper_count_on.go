//go:build paper_count

package vm

// Diagnostic-only sequential runner. Never use this build for latency results.
const paperCountInstructions = true

var paperInterpreted uint64

func PaperInterpretedInstructions() uint64 { return paperInterpreted }
