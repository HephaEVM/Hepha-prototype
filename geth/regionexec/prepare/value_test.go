package prepare

import (
	"testing"
	"time"
)

func TestValueBoundaryBatchAndCancellation(t *testing.T) {
	old, oldBoundary := valueCompileMode, windowBoundaryCompiler
	valueCompileMode = "boundary"
	windowBoundaryCompiler = true
	defer func() { valueCompileMode, windowBoundaryCompiler = old, oldBoundary }()
	m := &Manager{slots: map[string]*Slot{}, queueWake: make(chan struct{}, 1)}
	for i := 0; i < 20; i++ {
		s := &Slot{}
		m.queue = append(m.queue, s)
		m.pending.Add(1)
	}
	b := m.takeBatch()
	if len(b) != 16 || len(m.queue) != 4 {
		t.Fatal("unbounded compiler batch")
	}
	for range b {
		m.pending.Done()
	}
	m.CancelBoundary()
	done := make(chan struct{})
	go func() { m.pending.Wait(); close(done) }()
	select {
	case <-done:
	case <-time.After(time.Second):
		t.Fatal("cancellation did not release pending work")
	}
	if len(m.queue) != 0 {
		t.Fatal("unconsumed compiler work retained")
	}
}
