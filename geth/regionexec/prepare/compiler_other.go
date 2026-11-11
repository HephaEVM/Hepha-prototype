//go:build !darwin && !linux

package prepare

import (
	"bytes"
	"os/exec"
	"time"
)

// Other hosts retain bounded compilation without process-group preemption.
func (m *Manager) compilerOutput(cmd *exec.Cmd, limit time.Duration) ([]byte, time.Duration, error) {
	var output bytes.Buffer
	cmd.Stdout, cmd.Stderr = &output, &output
	if err := cmd.Start(); err != nil {
		return nil, 0, err
	}
	done, joined := make(chan struct{}), make(chan struct{})
	go func() {
		defer close(joined)
		t := time.NewTimer(limit)
		defer t.Stop()
		select {
		case <-done:
		case <-t.C:
			cmd.Process.Kill()
		}
	}()
	err := cmd.Wait()
	close(done)
	<-joined
	return output.Bytes(), 0, err
}
