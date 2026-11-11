package prepare

import "testing"

func TestDeliveryCompileCap(t *testing.T) {
	old, mode := deliveryCompileBudget, valueCompileMode
	deliveryCompileBudget = true
	valueCompileMode = "selective"
	defer func() { deliveryCompileBudget = old; valueCompileMode = mode }()
	m := &Manager{slots: map[string]*Slot{}}
	for i := 0; i < 8; i++ {
		m.slots[string(rune(i))] = new(Slot)
	}
	if !m.hotcallCompileBlocked() {
		t.Fatal("inflight cap")
	}
	delete(m.slots, string(rune(0)))
	m.completed = 1
	if m.hotcallCompileBlocked() {
		t.Fatal("completed item incorrectly retained")
	}
	m.completed = 25
	if !m.hotcallCompileBlocked() {
		t.Fatal("lifetime cap")
	}
}
