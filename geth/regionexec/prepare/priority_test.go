package prepare

import "testing"

func TestPriorityAdmissionCanReplaceAndRetry(t *testing.T) {
	m := testManager(10, 2)
	oldest, low, high := m.Slot("oldest", nil), m.Slot("low", nil), m.Slot("high", nil)
	alias := m.Slot("low", nil)
	oldest.Ready()
	low.Prioritize(10)
	low.Ready()
	alias.Ready()
	high.Prioritize(20)
	high.Ready()
	if m.displaced != 1 || len(m.queue) != 2 || m.queue[0] != oldest || m.queue[1] != high || low.started.Load() || low.Settled() {
		t.Fatal("invalid displacement")
	}
	alias.Ready()
	if low.started.Load() {
		t.Fatal("low score displaced high score")
	}
	for _, s := range m.takeBatch() {
		delete(m.slots, s.key)
		m.pending.Done()
	}
	alias.Ready()
	if !low.started.Load() || len(m.queue) != 1 || m.queue[0] != low {
		t.Fatal("displaced slot could not retry")
	}
	m.takeBatch()
	m.pending.Done()
	m.Drain()
}

func TestPriorityKeepsHeadAndBoundsQueue(t *testing.T) {
	m := testManager(300, 256)
	for i := 0; i < 200; i++ {
		s := &Slot{key: string(rune(i)), manager: m}
		s.Prioritize(int64(i))
		s.Ready()
	}
	batch := m.takeBatch()
	if len(batch) != 128 || len(m.queue) != 72 || batch[0].priority.Load() != 0 || batch[1].priority.Load() != 199 || batch[127].priority.Load() != 73 {
		t.Fatal("priority or fairness lost")
	}
	for range batch {
		m.pending.Done()
	}
	for range m.takeBatch() {
		m.pending.Done()
	}
	m.Drain()
}
