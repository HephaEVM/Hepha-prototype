//go:build darwin || linux

package prepare

import (
	"bytes"
	"fmt"
	"os/exec"
	"syscall"
	"time"
)

// Run only our own compiler and descendants in a private process group. The
// monitor, never the foreground, pauses it while foreground work is active.
// The timeout counts running time: a busy client must not exhaust it by yielding.
func (m *Manager) compilerOutput(cmd *exec.Cmd, limit time.Duration) ([]byte, time.Duration, error) {
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
	var output bytes.Buffer
	cmd.Stdout, cmd.Stderr = &output, &output
	if err := cmd.Start(); err != nil {
		return nil, 0, err
	}
	done := make(chan struct{})
	type result struct {
		paused   time.Duration
		timedOut bool
	}
	observed := make(chan result, 1)
	go func() {
		tick := time.NewTicker(time.Millisecond)
		defer tick.Stop()
		var r result
		var running time.Duration
		paused := false
		last := time.Now()
		account := func() {
			now := time.Now()
			if paused {
				r.paused += now.Sub(last)
			} else {
				running += now.Sub(last)
			}
			last = now
		}
		for {
			select {
			case <-done:
				account()
				observed <- r
				return
			case <-tick.C:
			}
			select {
			case <-done:
				account()
				observed <- r
				return
			default:
			}
			account()
			if running >= limit {
				// SIGKILL also terminates a stopped group and its output writers.
				syscall.Kill(-cmd.Process.Pid, syscall.SIGKILL)
				r.timedOut = true
				observed <- r
				return
			}
			busy := m.activity != nil && m.activity.active.Load() > 0 && m.draining.Load() == 0
			if busy != paused {
				sig := syscall.SIGCONT
				if busy {
					sig = syscall.SIGSTOP
				}
				if syscall.Kill(-cmd.Process.Pid, sig) == nil {
					paused = busy
				}
			}
		}
	}()
	err := cmd.Wait()
	close(done)
	r := <-observed
	if r.timedOut {
		err = fmt.Errorf("compiler exceeded %s of active time: %w", limit, err)
	}
	return output.Bytes(), r.paused, err
}
