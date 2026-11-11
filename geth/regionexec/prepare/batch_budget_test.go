package prepare

import (
	"strconv"
	"testing"
)

func TestDeliveryBatchPreservesOrderingAndResourceBounds(t *testing.T) {
	mode, enabled := valueCompileMode, deliveryCompileBudget
	batchLimit, outstanding, total := deliveryCompilerBatch, deliveryCompilerOutstanding, deliveryCompilerTotal
	t.Cleanup(func() {
		valueCompileMode, deliveryCompileBudget = mode, enabled
		deliveryCompilerBatch, deliveryCompilerOutstanding, deliveryCompilerTotal = batchLimit, outstanding, total
	})
	valueCompileMode, deliveryCompileBudget = "selective", true
	for _, limit := range []int{8, 32} {
		deliveryCompilerBatch, deliveryCompilerOutstanding, deliveryCompilerTotal = limit, 40, 2048
		m := &Manager{slots: map[string]*Slot{}, queueWake: make(chan struct{}, 1)}
		for i := 0; i < 40; i++ {
			s := &Slot{key: strconv.Itoa(i)}
			s.priority.Store(int64(i))
			m.slots[s.key] = s
			m.queue = append(m.queue, s)
		}
		batch := m.takeBatch()
		if len(batch) != limit || len(m.queue) != 40-limit || batch[0].key != "0" || batch[1].key != "39" || batch[limit-1].priority.Load() != int64(41-limit) {
			t.Fatal("batch bound, head reservation or priority changed", limit)
		}
		if !m.hotcallCompileBlocked() {
			t.Fatal("taken batch ceased to count as outstanding work")
		}
		for _, s := range batch {
			delete(m.slots, s.key)
			m.completed++
		}
		if m.hotcallCompileBlocked() {
			t.Fatal("completed batch did not release outstanding capacity")
		}
		m.completed = 2048 - len(m.slots)
		if !m.hotcallCompileBlocked() {
			t.Fatal("larger batch bypassed completed plus pending lifetime bound")
		}
	}
}
