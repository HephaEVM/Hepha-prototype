package vm

// The synchronous loop never owns worker queues or speculative predictions.
// Publish only counters that can actually change; no asynchronous join is needed.
func (f *arithmeticFrame) finishSynchronous() *arithmeticResult {
	if f.finished {
		return f.first
	}
	s := &f.config.Engine.stats
	if f.pruned != 0 {
		s.pruned.Add(f.pruned)
	}
	if f.extended != 0 {
		s.extended.Add(f.extended)
	}
	if f.fallbacks != 0 {
		s.fallbacks.Add(f.fallbacks)
	}
	if f.checked != 0 {
		s.checked.Add(f.checked)
	}
	if f.nativeChecked != 0 {
		s.native.Add(f.nativeChecked)
	}
	if f.checkFailures != 0 {
		s.failures.Add(f.checkFailures)
	}
	if f.fixedPrunes != 0 {
		s.fixed.Add(f.fixedPrunes)
	}
	if f.constantPairs != 0 {
		s.constantPairs.Add(f.constantPairs)
	}
	f.finished = true
	return f.first
}
