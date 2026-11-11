package prepare

import (
	"testing"

	"github.com/holiman/uint256"
	"go-eth1.16.5-evm/regionexec"
)

func TestReuseCannotStartUnknownWork(t *testing.T) {
	m := testManager(1, 1)
	s := m.Slot("unknown", func() (string, error) { t.Fatal("reuse invoked compiler builder"); return "", nil })
	if s.Reuse() != nil || s.PreparationStarted() || len(m.slots) != 0 || len(m.queue) != 0 {
		t.Fatal("cache-only reuse admitted new work")
	}
}

func TestReuseLateFactoryPreservesDistinctArguments(t *testing.T) {
	m := testManager(1, 1)
	a := m.FactorySlot("shared", nil, []uint256.Int{{7}}, [4]uint64{10})
	b := m.FactorySlot("shared", nil, []uint256.Int{{3}}, [4]uint64{20})
	// The factory became available after these candidates were discovered.
	m.factories = map[string]Factory{"shared": func(v []uint256.Int, e [4]uint64) regionexec.Executor {
		return func(f *regionexec.Frame) bool { f.PC = e[0] + v[0].Uint64(); return true }
	}}
	m.completed = 4096
	for i, s := range []*Slot{a, b} {
		fn := s.Reuse()
		f := new(regionexec.Frame)
		want := []uint64{17, 23}[i]
		if fn == nil || !fn(f) || f.PC != want || !s.Settled() {
			t.Fatal("late shared factory lost immutable binding", i, f.PC)
		}
	}
	if len(m.queue) != 0 || len(m.slots) != 0 || m.completed != 4096 {
		t.Fatal("ready factory consumed compiler quota")
	}
}

func TestReusePendingWorkDoesNotDuplicateAdmission(t *testing.T) {
	m := testManager(1, 1)
	a, b := m.Slot("same", nil), m.Slot("same", nil)
	a.Ready()
	if b.Reuse() != nil || b.alias.Load() != a || len(m.queue) != 1 {
		t.Fatal("pending shared work was duplicated")
	}
	var fn Executor = func(f *regionexec.Frame) bool { f.PC = 42; return true }
	a.ready.Store(&fn)
	f := new(regionexec.Frame)
	if got := b.Reuse(); got == nil || !got(f) || f.PC != 42 {
		t.Fatal("completion did not reach waiting alias")
	}
	m.takeBatch()
	m.pending.Done()
	m.Drain()
}
