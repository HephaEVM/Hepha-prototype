package prepare

import (
	"github.com/holiman/uint256"
	"go-eth1.16.5-evm/regionexec"
	"testing"
)

func TestLoadedDoesNotActivate(t *testing.T) {
	m := testManager(8, 8)
	s := m.Slot("pending", nil)
	m.mu.Lock()
	if s.Published() != nil || s.Loaded() != nil || s.Settled() {
		t.Fatal("unprepared slot ready")
	}
	m.mu.Unlock()
	if len(m.queue) != 0 || len(m.slots) != 0 {
		t.Fatal("read-only load activated work")
	}
	s.ready.Store(new(Executor))
}

func TestLoadedFactoryAndAlias(t *testing.T) {
	m := testManager(8, 8)
	s := m.FactorySlot("family", nil, []uint256.Int{{7}}, [4]uint64{23})
	var f Factory = func(v []uint256.Int, e [4]uint64) regionexec.Executor {
		return func(out *regionexec.Frame) bool { out.PC = e[0] + v[0].Uint64(); return true }
	}
	s.template.factoryReady.Store(&f)
	if s.Published() != nil {
		t.Fatal("unbound factory published as an executable")
	}
	fn := s.Loaded()
	frame := new(regionexec.Frame)
	if fn == nil || !fn(frame) || frame.PC != 30 || !s.Settled() {
		t.Fatal("factory binding lost arguments")
	}
	if fn = s.Published(); fn == nil || !fn(frame) || frame.PC != 30 {
		t.Fatal("hot read lost the published executor")
	}
	if len(m.queue) != 0 {
		t.Fatal("loaded factory requested compilation")
	}
}
