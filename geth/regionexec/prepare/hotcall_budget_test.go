package prepare

import (
	"strconv"
	"testing"
)

func TestHotcallCompilerBudgets(t *testing.T) {
	old := valueCompileMode
	valueCompileMode = "selective"
	defer func() { valueCompileMode = old }()
	m := &Manager{slots: map[string]*Slot{}}
	for i := 0; i < 128; i++ {
		m.slots[string(rune(i))] = new(Slot)
	}
	if !m.hotcallCompileBlocked() {
		t.Fatal("inflight cap ignored")
	}
	m.completed = 64
	if !m.hotcallCompileBlocked() {
		t.Fatal("completed jobs incorrectly subtracted from outstanding")
	}
	for i := 0; i < 64; i++ {
		delete(m.slots, string(rune(i)))
	}
	if m.hotcallCompileBlocked() {
		t.Fatal("finished-and-removed jobs blocked next wave")
	}
	m.slots = map[string]*Slot{}
	m.completed = 512
	if !m.hotcallCompileBlocked() {
		t.Fatal("lifetime artifact cap ignored")
	}
	s := new(Slot)
	if s.PreparationStarted() {
		t.Fatal("new slot started")
	}
	s.started.Store(true)
	alias := &Slot{template: s}
	if !alias.PreparationStarted() {
		t.Fatal("shared work not recognized")
	}
}

func TestExpandedCompilerBudgetStillBoundsOutstandingWork(t *testing.T) {
	mode, enabled := valueCompileMode, deliveryCompileBudget
	total, outstanding := deliveryCompilerTotal, deliveryCompilerOutstanding
	t.Cleanup(func() {
		valueCompileMode, deliveryCompileBudget = mode, enabled
		deliveryCompilerTotal, deliveryCompilerOutstanding = total, outstanding
	})
	valueCompileMode, deliveryCompileBudget = "selective", true
	deliveryCompilerTotal, deliveryCompilerOutstanding = 256, 8
	m := &Manager{completed: 32, slots: map[string]*Slot{}}
	if m.hotcallCompileBlocked() {
		t.Fatal("old lifetime threshold still blocks an admitted later upgrade")
	}
	for i := 0; i < 8; i++ {
		m.slots[strconv.Itoa(i)] = new(Slot)
	}
	if !m.hotcallCompileBlocked() {
		t.Fatal("expanded lifetime budget removed outstanding-work bound")
	}
	delete(m.slots, "0")
	if m.hotcallCompileBlocked() {
		t.Fatal("completion failed to release outstanding-work capacity")
	}
	m.completed = 249
	if !m.hotcallCompileBlocked() {
		t.Fatal("completed plus pending exceeded lifetime budget")
	}
}
