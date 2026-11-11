package prepare

import "testing"

func TestPriorityReplacementAtSoftAdmissionBound(t *testing.T) {
	mode, budget, replacement := valueCompileMode, deliveryCompileBudget, deliveryPriorityReplacement
	total, outstanding := deliveryCompilerTotal, deliveryCompilerOutstanding
	t.Cleanup(func() {
		valueCompileMode, deliveryCompileBudget, deliveryPriorityReplacement = mode, budget, replacement
		deliveryCompilerTotal, deliveryCompilerOutstanding = total, outstanding
	})
	valueCompileMode, deliveryCompileBudget, deliveryPriorityReplacement = "selective", true, true
	deliveryCompilerTotal, deliveryCompilerOutstanding = 4, 4
	m := &Manager{slots: map[string]*Slot{}, capacity: 100, queueCapacity: 100, queueWake: make(chan struct{}, 1)}
	for i, key := range []string{"running", "head", "low", "medium"} {
		s := &Slot{key: key}
		s.priority.Store(int64(i))
		s.started.Store(true)
		m.slots[key] = s
		m.pending.Add(1)
		if key != "running" {
			m.queue = append(m.queue, s)
		}
	}
	old := m.slots["low"]
	high := &Slot{key: "high"}
	high.priority.Store(10)
	m.admitLocked(high)
	if len(m.slots) != 4 || len(m.queue) != 3 || m.slots["running"] == nil || m.queue[0].key != "head" || m.slots["high"] != high || m.slots["low"] != nil || old.PreparationStarted() {
		t.Fatal("replacement changed bounds, active work, or head reservation")
	}
	if !high.PreparationStarted() || m.displaced != 1 {
		t.Fatal("new work not admitted")
	}
	low := &Slot{key: "lower"}
	m.admitLocked(low)
	if low.PreparationStarted() || len(m.slots) != 4 || m.displaced != 1 {
		t.Fatal("lower priority displaced better work")
	}
	deliveryPriorityReplacement = false
	higher := &Slot{key: "higher"}
	higher.priority.Store(20)
	m.admitLocked(higher)
	if higher.PreparationStarted() || m.displaced != 1 {
		t.Fatal("default control policy changed")
	}
	// Four pending jobs remain, not five: displacement balanced the waitgroup.
	for range m.slots {
		m.pending.Done()
	}
	m.pending.Wait()
}
