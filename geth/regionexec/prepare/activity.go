package prepare

import (
	"sync/atomic"
	"time"
)

// Activity gives foreground execution priority without making it acquire a
// worker lock or wait for preparation. It is advisory: work already in flight
// finishes at its next preparation boundary. Nested/concurrent users share it.
type Activity struct {
	epoch  time.Time
	active atomic.Int64
	last   atomic.Int64
	waits  atomic.Uint64
	waitNS atomic.Int64
}

func NewActivity() *Activity { return &Activity{epoch: time.Now()} }
func (a *Activity) Begin()   { a.active.Add(1) }
func (a *Activity) End() {
	a.last.Store(time.Since(a.epoch).Nanoseconds())
	if a.active.Add(-1) < 0 {
		panic("unbalanced foreground activity")
	}
}

// A short quiet period avoids treating tiny gaps between calls as an invitation
// to start another allocation-heavy analysis. No foreground caller waits here.
func (a *Activity) Wait(stop <-chan struct{}, draining *atomic.Int64) bool {
	start := time.Now()
	waited := false
	for {
		select {
		case <-stop:
			return false
		default:
		}
		if draining != nil && draining.Load() > 0 {
			break
		}
		if a.active.Load() == 0 && time.Since(a.epoch).Nanoseconds()-a.last.Load() >= int64(500*time.Microsecond) {
			break
		}
		waited = true
		timer := time.NewTimer(time.Millisecond)
		select {
		case <-stop:
			timer.Stop()
			return false
		case <-timer.C:
		}
	}
	if waited {
		a.waits.Add(1)
		a.waitNS.Add(time.Since(start).Nanoseconds())
	}
	return true
}

func (a *Activity) Stats() (uint64, int64) { return a.waits.Load(), a.waitNS.Load() }
