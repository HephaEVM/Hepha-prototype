package prepare

import (
	"path/filepath"
	"testing"

	"github.com/holiman/uint256"
	"go-eth1.16.5-evm/regionexec"
)

func TestFactorySharesCompilationAndRebinds(t *testing.T) {
	root, _ := filepath.Abs("../..")
	m, err := NewManager(root, t.TempDir())
	if err != nil {
		t.Fatal(err)
	}
	body := `func Exec(values []uint256.Int,exits [4]uint64)regionexec.Executor{return func(s *regionexec.Frame)bool{if s.Gas<3{return false};s.Gas-=3;s.PC=exits[0];s.Stack=append(s.Stack,values[0]);return true}}`
	build := func() (string, error) { return body, nil }
	a := m.FactorySlot("test-family", build, []uint256.Int{{7}}, [4]uint64{10})
	b := m.FactorySlot("test-family", build, []uint256.Int{{9}}, [4]uint64{20})
	a.Ready()
	b.Ready()
	m.Drain()
	for i, s := range []*Slot{a, b} {
		f := &regionexec.Frame{Gas: 5}
		fn := s.Ready()
		if fn == nil || !fn(f) || f.PC != uint64(10+i*10) || f.Stack[0][0] != uint64(7+i*2) || f.Gas != 2 {
			t.Fatal("bad bound executor", i, f, m.Events())
		}
	}
	events := m.Events()
	if len(events) != 1 || len(events[0].Keys) != 1 || events[0].Error != "" {
		t.Fatal(events)
	}
	// A newly bound instance reuses the already loaded factory without compilation.
	c := m.FactorySlot("test-family", build, []uint256.Int{{11}}, [4]uint64{30})
	f := &regionexec.Frame{Gas: 2}
	if fn := c.Ready(); fn == nil || fn(f) || f.Gas != 2 || len(f.Stack) != 0 {
		t.Fatal("fallback changed frame")
	}
	if len(m.Events()) != 1 {
		t.Fatal("recompiled a known family")
	}
	// A fresh manager indexes the persisted factory without eagerly loading it.
	// Its first request restores the artifact and binds different immutable data.
	reopened, err := NewManager(root, m.dir)
	if err != nil {
		t.Fatal(err)
	}
	d := reopened.FactorySlot("test-family", func() (string, error) {
		t.Error("restored family unexpectedly rebuilt")
		return body, nil
	}, []uint256.Int{{13}}, [4]uint64{40})
	d.Ready()
	reopened.Drain()
	f = &regionexec.Frame{Gas: 5}
	if fn := d.Ready(); fn == nil || !fn(f) || f.Stack[0][0] != 13 || f.PC != 40 || f.Gas != 2 {
		t.Fatal("persisted factory failed to rebind", f, reopened.Events())
	}
	if events := reopened.Events(); len(events) != 1 || !events[0].CacheHit || events[0].Error != "" {
		t.Fatal("expected one artifact restoration", events)
	}
}
