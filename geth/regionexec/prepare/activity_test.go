package prepare

import (
	"sync/atomic"
	"testing"
	"time"
)

func TestActivityNestedAndCancellation(t *testing.T) {
	a := NewActivity()
	a.Begin()
	a.Begin()
	stop := make(chan struct{})
	done := make(chan bool, 1)
	go func() { done <- a.Wait(stop, nil) }()
	a.End()
	select {
	case <-done:
		t.Fatal("worker entered while another foreground user was active")
	case <-time.After(5 * time.Millisecond):
	}
	a.End()
	select {
	case ok := <-done:
		if !ok {
			t.Fatal("unexpected cancellation")
		}
	case <-time.After(time.Second):
		t.Fatal("idle worker did not resume")
	}
	a.Begin()
	go func() { done <- a.Wait(stop, nil) }()
	close(stop)
	if <-done {
		t.Fatal("stop must cancel waiting")
	}
	a.End()
}

func TestActivityOfflineDrain(t *testing.T) {
	a := NewActivity()
	a.Begin()
	defer a.End()
	var draining atomic.Int64
	draining.Store(1)
	if !a.Wait(nil, &draining) {
		t.Fatal("offline drain was cancelled")
	}
}
