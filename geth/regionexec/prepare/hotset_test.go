package prepare

import (
	"fmt"
	"go-eth1.16.5-evm/regionexec"
	"testing"
)

func TestHotsetBoundDoesNotRejectOrBlockLoaded(t *testing.T) {
	t.Setenv("COLD_HOTSET", "1")
	m := testManager(1024, 2048)
	for i := 0; i < 128; i++ {
		s := &Slot{key: fmt.Sprint(i), manager: m}
		m.activate(s)
	}
	late := &Slot{key: "late", manager: m}
	m.activate(late)
	if late.started.Load() || late.Settled() || len(m.slots) != 128 {
		t.Fatal("soft bound rejected or admitted excess work")
	}
	shared := &Slot{key: "1", manager: m}
	m.activate(shared)
	if shared.alias.Load() == nil {
		t.Fatal("dedup blocked by cap")
	}
	m.loaded["ready"] = func(*regionexec.Frame) bool { return true }
	loaded := &Slot{key: "ready", manager: m}
	m.activate(loaded)
	if loaded.Loaded() == nil || len(m.slots) != 128 {
		t.Fatal("available executor blocked by compilation cap")
	}
}
