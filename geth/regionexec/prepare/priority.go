package prepare

import "sort"

// Called under manager.mu. Only queued, not executing, work can be displaced.
// Keep the queue head's reservation. This is scheduling, not cache
// eviction: its original slot remains valid and can request preparation again.
func (m *Manager) replaceQueued(s *Slot) bool {
	lowest := -1
	score := s.priority.Load()
	for i := 1; i < len(m.queue); i++ {
		if p := m.queue[i].priority.Load(); p < score {
			lowest = i
			score = p
		}
	}
	if lowest < 0 {
		return false
	}
	old := m.queue[lowest]
	delete(m.slots, old.key)
	old.started.Store(false)
	m.queue[lowest] = s
	m.pending.Done()
	m.displaced++
	return true
}

// Prioritize changes cold work ordering only. Multiple bindings to a shared
// factory contribute their best opportunity, without repeated-visit inflation.
func (s *Slot) Prioritize(score int64) {
	if s.template != nil {
		s.template.Prioritize(score)
		return
	}
	if a := s.alias.Load(); a != nil {
		a.Prioritize(score)
		return
	}
	for old := s.priority.Load(); score > old; old = s.priority.Load() {
		if s.priority.CompareAndSwap(old, score) {
			return
		}
	}
}

func (m *Manager) takeBatch() []*Slot {
	m.mu.Lock()
	defer m.mu.Unlock()
	if len(m.queue) == 0 {
		return nil
	}
	// Reserve the current queue head, then prefer useful work over small
	// upgrades. This is best-effort progress, not FIFO fairness: the remaining
	// queue is priority ordered after each batch.
	first := m.queue[0]
	sort.SliceStable(m.queue[1:], func(i, j int) bool { return m.queue[i+1].priority.Load() > m.queue[j+1].priority.Load() })
	n := min(128, len(m.queue))
	if valueCompileMode == "selective" {
		n = min(n, 32)
		if deliveryCompileBudget {
			n = min(n, deliveryCompilerBatch)
		}
	}
	if valueCompileMode == "boundary" {
		n = min(n, 16)
	}
	batch := append([]*Slot(nil), m.queue[:n]...)
	batch[0] = first
	copy(m.queue, m.queue[n:])
	clear(m.queue[len(m.queue)-n:])
	m.queue = m.queue[:len(m.queue)-n]
	if len(m.queue) > 0 {
		select {
		case m.queueWake <- struct{}{}:
		default:
		}
	}
	return batch
}
