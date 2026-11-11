package vm

import "os"

var experimentFocusedPreparation = os.Getenv("COLD_FOCUSED_PREPARATION") == "1"

// Preparation only. Never disable independent loops, partial functions,
// arithmetic, ABI or address summaries. Preserve the normal yielding point.
func experimentSkipRejectedAnalysis() bool {
	return experimentFocusedPreparation && experimentWorkPolicy == "none"
}
