package prepare

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func TestFailedBatchPublishesRevision(t *testing.T) {
	m, err := NewManager(t.TempDir(), t.TempDir())
	if err != nil {
		t.Fatal(err)
	}
	s := m.DeferredSlot("bad-analysis", func() (string, string, error) { return "", "", fmt.Errorf("unsupported fixture") })
	s.Ready()
	m.Drain()
	until := time.Now().Add(time.Second)
	for m.Revision() == 0 && time.Now().Before(until) {
		runtime.Gosched()
	}
	if m.Revision() != 1 || !s.Settled() || s.Loaded() != nil {
		t.Fatal("failed preparation did not release publication waiters")
	}
}
