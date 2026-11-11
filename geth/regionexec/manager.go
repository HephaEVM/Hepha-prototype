package regionexec

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"plugin"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

type Builder func() (string, error)
type Slot struct {
	key     string
	build   Builder
	manager *Manager
	started atomic.Bool
	ready   atomic.Pointer[Executor]
}

func (s *Slot) Ready() Executor {
	if p := s.ready.Load(); p != nil {
		return *p
	}
	if s.started.CompareAndSwap(false, true) {
		s.manager.pending.Add(1)
		select {
		case s.manager.queue <- s:
		default:
			s.manager.pending.Done()
		}
	}
	return nil
}

type Event struct {
	Keys                   []string
	Seconds                float64
	Error                  string
	SourceHash, BinaryHash string
}
type Manager struct {
	dir, root          string
	mu                 sync.Mutex
	slots              map[string]*Slot
	loaded             map[string]Executor
	queue              chan *Slot
	pending            sync.WaitGroup
	events             []Event
	capacity, rejected int
}

func NewManager(root, dir string) (*Manager, error) {
	if err := os.MkdirAll(dir, 0700); err != nil {
		return nil, err
	}
	capacity := 16384
	if raw := os.Getenv("DEMAND_MAX_SLOTS"); raw != "" {
		v, err := strconv.Atoi(raw)
		if err != nil || v < 16384 || v > 262144 {
			return nil, fmt.Errorf("invalid DEMAND_MAX_SLOTS")
		}
		capacity = v
	}
	m := &Manager{dir: dir, root: root, capacity: capacity, slots: map[string]*Slot{}, loaded: map[string]Executor{}, queue: make(chan *Slot, 2048)}
	metas, _ := filepath.Glob(filepath.Join(dir, "*.json"))
	sort.Strings(metas)
	for _, f := range metas {
		var v Event
		b, e := os.ReadFile(f)
		if e != nil || json.Unmarshal(b, &v) != nil || v.Error != "" {
			continue
		}
		stem := strings.TrimSuffix(f, ".json")
		src, e := os.ReadFile(stem + ".go.txt")
		if e != nil || digest(src) != v.SourceHash {
			continue
		}
		bin, e := os.ReadFile(stem + ".so")
		if e != nil || digest(bin) != v.BinaryHash {
			continue
		}
		if e = m.load(stem + ".so"); e != nil {
			m.events = append(m.events, Event{Error: e.Error()})
		}
	}
	go m.worker()
	return m, nil
}
func digest(b []byte) string { x := sha256.Sum256(b); return hex.EncodeToString(x[:]) }
func (m *Manager) Slot(key string, b Builder) *Slot {
	m.mu.Lock()
	defer m.mu.Unlock()
	if s := m.slots[key]; s != nil {
		return s
	}
	if len(m.slots) >= m.capacity {
		m.rejected++
		return nil
	}
	s := &Slot{key: key, build: b, manager: m}
	if f := m.loaded[key]; f != nil {
		s.ready.Store(&f)
		s.started.Store(true)
	}
	m.slots[key] = s
	return s
}
func (m *Manager) Drain() { m.pending.Wait() }

// Usage exposes admission saturation so experiments cannot silently compare
// configurations with different coverage because they share a bounded cache.
func (m *Manager) Usage() (registered, limit, rejected int) {
	m.mu.Lock()
	defer m.mu.Unlock()
	return len(m.slots), m.capacity, m.rejected
}
func (m *Manager) Events() []Event {
	m.mu.Lock()
	defer m.mu.Unlock()
	return append([]Event(nil), m.events...)
}
func (m *Manager) load(path string) error {
	p, err := plugin.Open(path)
	if err != nil {
		return err
	}
	x, err := p.Lookup("Executors")
	if err != nil {
		return err
	}
	v, ok := x.(*map[string]Executor)
	if !ok {
		return fmt.Errorf("unexpected executor ABI")
	}
	m.mu.Lock()
	defer m.mu.Unlock()
	for k, f := range *v {
		m.loaded[k] = f
		if s := m.slots[k]; s != nil {
			fn := f
			s.ready.Store(&fn)
		}
	}
	return nil
}
func (m *Manager) worker() {
	for first := range m.queue {
		batch := []*Slot{first}
		timer := time.NewTimer(40 * time.Millisecond)
	collect:
		for len(batch) < 128 {
			select {
			case s := <-m.queue:
				batch = append(batch, s)
			case <-timer.C:
				break collect
			}
		}
		timer.Stop()
		m.compile(batch)
		for range batch {
			m.pending.Done()
		}
	}
}
func (m *Manager) compile(batch []*Slot) {
	start := time.Now()
	event := Event{}
	var src strings.Builder
	src.WriteString("package main\nimport(\"math/bits\";\"github.com/holiman/uint256\";\"go-eth1.16.5-evm/regionexec\")\nvar _=bits.Mul64\nvar _ uint256.Int\n")
	var entries []string
	for _, s := range batch {
		body, err := s.build()
		if err != nil {
			m.mu.Lock()
			m.events = append(m.events, Event{Keys: []string{s.key}, Error: err.Error()})
			m.mu.Unlock()
			continue
		}
		name := "R" + digest([]byte(s.key))
		src.WriteString(strings.ReplaceAll(body, "Exec", name))
		entries = append(entries, fmt.Sprintf("%q:%s", s.key, name))
		event.Keys = append(event.Keys, s.key)
	}
	if len(entries) == 0 {
		return
	}
	src.WriteString("\nvar Executors=map[string]regionexec.Executor{" + strings.Join(entries, ",") + "}\n")
	data := []byte(src.String())
	event.SourceHash = digest(data)
	stem := filepath.Join(m.dir, event.SourceHash)
	work, err := os.MkdirTemp(m.dir, "compile-")
	if err != nil {
		event.Error = err.Error()
	} else {
		file := filepath.Join(work, "main.go")
		err = os.WriteFile(file, data, 0600)
		if err == nil {
			ctx, cancel := context.WithTimeout(context.Background(), 90*time.Second)
			defer cancel()
			args := []string{"build", "-buildmode=plugin"}
			if runtime.GOOS == "darwin" {
				args = append(args, "-ldflags=-extldflags=-Wl,-no_fixup_chains")
			}
			args = append(args, "-o", stem+".so", file)
			cmd := exec.CommandContext(ctx, filepath.Join(runtime.GOROOT(), "bin", "go"), args...)
			cmd.Dir = m.root
			cmd.Env = append(os.Environ(), "GOPROXY=off", "GOSUMDB=off", "GOMAXPROCS=2", "GOTOOLCHAIN=local")
			out, e := cmd.CombinedOutput()
			if e != nil {
				err = fmt.Errorf("compile: %w: %.1200s", e, out)
			}
		}
		if err == nil {
			err = m.load(stem + ".so")
		}
		if err != nil {
			event.Error = err.Error()
		}
	}
	os.WriteFile(stem+".go.txt", data, 0600)
	if bin, e := os.ReadFile(stem + ".so"); e == nil {
		event.BinaryHash = digest(bin)
	}
	event.Seconds = time.Since(start).Seconds()
	raw, _ := json.MarshalIndent(event, "", "  ")
	os.WriteFile(stem+".json", raw, 0600)
	m.mu.Lock()
	m.events = append(m.events, event)
	m.mu.Unlock()
}
