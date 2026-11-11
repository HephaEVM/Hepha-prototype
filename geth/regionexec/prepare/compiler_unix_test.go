//go:build darwin || linux

package prepare

import (
	"os/exec"
	"testing"
	"time"
)

func TestCompilerYieldsWithoutSpendingTimeout(t *testing.T) {
	m := &Manager{activity: NewActivity()}
	m.activity.Begin()
	done := make(chan error, 1)
	go func() {
		out, paused, err := m.compilerOutput(exec.Command("sh", "-c", "sleep 0.03; printf ready"), 60*time.Millisecond)
		if err == nil && (string(out) != "ready" || paused < 50*time.Millisecond) {
			t.Errorf("output=%q paused=%s", out, paused)
		}
		done <- err
	}()
	select {
	case err := <-done:
		t.Fatalf("compiler did not yield: %v", err)
	case <-time.After(100 * time.Millisecond):
	}
	m.activity.End()
	select {
	case err := <-done:
		if err != nil {
			t.Fatal(err)
		}
	case <-time.After(time.Second):
		t.Fatal("compiler did not resume")
	}
}
func TestCompilerActiveTimeoutKillsChildren(t *testing.T) {
	m := &Manager{activity: NewActivity()}
	done := make(chan error, 1)
	go func() {
		_, _, err := m.compilerOutput(exec.Command("sh", "-c", "sleep 10"), 10*time.Millisecond)
		done <- err
	}()
	select {
	case err := <-done:
		if err == nil {
			t.Fatal("timeout succeeded")
		}
	case <-time.After(time.Second):
		t.Fatal("child output pipe survived timeout")
	}
}

func TestCompilerOfflineDrainResumesActiveForeground(t *testing.T) {
	m := &Manager{activity: NewActivity()}
	m.activity.Begin()
	defer m.activity.End()
	done := make(chan error, 1)
	go func() {
		_, _, err := m.compilerOutput(exec.Command("sh", "-c", "sleep 0.03; printf ready"), time.Second)
		done <- err
	}()
	select {
	case err := <-done:
		t.Fatalf("did not yield: %v", err)
	case <-time.After(30 * time.Millisecond):
	}
	m.draining.Add(1)
	defer m.draining.Add(-1)
	select {
	case err := <-done:
		if err != nil {
			t.Fatal(err)
		}
	case <-time.After(time.Second):
		t.Fatal("offline drain did not resume compiler")
	}
}
