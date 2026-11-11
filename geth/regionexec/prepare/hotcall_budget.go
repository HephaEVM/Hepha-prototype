package prepare

import (
	"os"
	"strconv"
)

var deliveryCompileBudget = os.Getenv("COLD_SELECTIVE_DELIVERY") == "1"
var deliveryCompilerTotal = compilerBudgetInt("COMPILER_TOTAL_LIMIT", 32, 16384)
var deliveryCompilerOutstanding = compilerBudgetInt("COMPILER_OUTSTANDING_LIMIT", 8, 128)
var deliveryCompilerBatch = compilerBudgetInt("COMPILER_BATCH_LIMIT", 8, 32)
var deliveryPriorityReplacement = os.Getenv("COLD_TRADE_COMPILER_REPLACE_QUEUED") == "1"

func compilerBudgetInt(name string, fallback, maximum int) int {
	s := os.Getenv("COLD_TRADE_" + name)
	if s == "" {
		return fallback
	}
	n, err := strconv.Atoi(s)
	if err != nil || n < 1 || n > maximum {
		panic("invalid COLD_TRADE_" + name)
	}
	return n
}

// No compiler work on a foreground path. This observation is only used by the
// preparation worker to avoid counting shared queued/loaded work twice.
func (s *Slot) PreparationStarted() bool {
	if s.template != nil {
		return s.template.PreparationStarted()
	}
	if a := s.alias.Load(); a != nil {
		return a.PreparationStarted()
	}
	return s.started.Load()
}
func (m *Manager) hotcallCompileBlocked() bool {
	if valueCompileMode != "selective" {
		return false
	}
	if deliveryCompileBudget {
		return m.completed+len(m.slots) >= deliveryCompilerTotal || len(m.slots) >= deliveryCompilerOutstanding
	}
	// Lifetime artifact budget and instantaneous queue budget are distinct.
	// Completed slots leave the pending registry; completed+pending bounds lifetime work.
	if m.completed+len(m.slots) >= 512 {
		return true
	}
	return len(m.slots) >= 128
}
