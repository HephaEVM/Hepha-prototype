package prepare

import "os"

var valueCompileMode = os.Getenv("COLD_VALUE_COMPILER")

// Exactly one admitted batch, bounded by takeBatch's policy. Called exclusively
// at a block boundary; no compiler worker exists in this mode. No global Drain.
func (m *Manager) AdvanceBoundary() bool {
	if !windowBoundaryCompiler {
		return false
	}
	batch := m.takeBatch()
	if len(batch) == 0 {
		return false
	}
	done := m.Unthrottle()
	defer done()
	m.processBatch(batch)
	return true
}

// Work not used before the end of the window is cancelled, not secretly
// prepared after timing. No worker can own a batch in boundary mode.
func (m *Manager) CancelBoundary() {
	if !windowBoundaryCompiler {
		return
	}
	m.mu.Lock()
	defer m.mu.Unlock()
	for _, s := range m.queue {
		delete(m.slots, s.key)
		s.build = nil
		s.resolve = nil
		s.finished.Store(true)
		m.pending.Done()
	}
	m.queue = nil
}

var windowBoundaryCompiler = valueCompileMode == "boundary" || os.Getenv("COLD_WINDOW_COMPILER") == "1"
