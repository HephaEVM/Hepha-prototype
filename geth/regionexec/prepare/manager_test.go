package prepare

import (
	"fmt"
	"go-eth1.16.5-evm/regionexec"
	"sync/atomic"
	"testing"
)

func testManager(capacity, queue int) *Manager {
	return &Manager{slots: map[string]*Slot{}, loaded: map[string]Executor{}, queueWake: make(chan struct{}, 1), queueCapacity: queue, capacity: capacity}
}

func TestUnexecutedCandidatesDoNotConsumeCapacity(t *testing.T) {
	m := testManager(2, 2)
	for i := 0; i < 10000; i++ {
		if m.Slot(fmt.Sprint(i), nil) == nil {
			t.Fatal("candidate rejected")
		}
	}
	used, _, rejected := m.Usage()
	if used != 0 || rejected != 0 {
		t.Fatalf("unused candidates retained: %d/%d", used, rejected)
	}
	s := m.Slot("executed", nil)
	s.Ready()
	if used, _, _ = m.Usage(); used != 1 {
		t.Fatal(used)
	}
	m.takeBatch()
	m.pending.Done()
	m.Drain()
}

func TestActivationDeduplicatesAndPublishesToAliases(t *testing.T) {
	m := testManager(1, 1)
	a, b := m.Slot("same", nil), m.Slot("same", nil)
	a.Ready()
	b.Ready()
	if len(m.queue) != 1 || b.alias.Load() != a {
		t.Fatal("duplicate preparation")
	}
	var fn Executor = func(*regionexec.Frame) bool { return true }
	a.ready.Store(&fn)
	if b.Ready() == nil || b.ready.Load() == nil {
		t.Fatal("alias did not receive result")
	}
	if m.Slot("same", nil) != a {
		t.Fatal("active slot not reused")
	}
	m.takeBatch()
	m.pending.Done()
	m.Drain()
}

func TestFullQueueCanRetryWithoutLeakingCapacity(t *testing.T) {
	m := testManager(4, 1)
	a, b := m.Slot("a", nil), m.Slot("b", nil)
	a.Ready()
	b.Ready()
	if b.started.Load() || len(m.slots) != 1 {
		t.Fatal("queue rejection leaked active slot")
	}
	m.takeBatch()
	m.pending.Done()
	b.Ready()
	if !b.started.Load() || len(m.queue) != 1 {
		t.Fatal("queue rejection became permanent")
	}
	m.takeBatch()
	m.pending.Done()
	m.Drain()
}

func TestCapacityFailureFallsBackOnce(t *testing.T) {
	m := testManager(1, 1)
	a, b := m.Slot("a", nil), m.Slot("b", nil)
	a.Ready()
	for i := 0; i < 10; i++ {
		if b.Ready() != nil {
			t.Fatal("unexpected executor")
		}
	}
	used, _, rejected := m.Usage()
	if used != 1 || rejected != 1 {
		t.Fatalf("capacity accounting: %d/%d", used, rejected)
	}
	m.takeBatch()
	m.pending.Done()
	m.Drain()
}

func TestDeferredLoweringOnlyOnExecution(t *testing.T) {
	m := testManager(4, 4)
	var calls atomic.Int32
	var fn Executor = func(*regionexec.Frame) bool { return true }
	m.loaded["existing-artifact"] = fn
	resolve := func() (string, string, error) { calls.Add(1); return "existing-artifact", "source", nil }
	m.DeferredSlot("discarded", resolve)
	a, b := m.DeferredSlot("selected", resolve), m.DeferredSlot("selected", resolve)
	if calls.Load() != 0 {
		t.Fatal("eager lowering")
	}
	go m.worker()
	a.Ready()
	b.Ready()
	m.Drain()
	close(m.queueWake)
	if calls.Load() != 1 || a.Ready() == nil || b.Ready() == nil {
		t.Fatal("lowering/restore not shared")
	}
	if len(m.slots) != 0 || m.Slot("selected", nil).Ready() == nil {
		t.Fatal("completed task retained or executor lost")
	}
	if a.build != nil || a.resolve != nil {
		t.Fatal("retained compiler inputs")
	}
}
