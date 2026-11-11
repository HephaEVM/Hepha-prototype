// Package prepare owns preparation policy separately from the stable executor ABI.
package prepare

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/holiman/uint256"
	"go-eth1.16.5-evm/regionexec"
	"os"
	"os/exec"
	"path/filepath"
	"plugin"
	"regexp"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

var executorName = regexp.MustCompile(`\bExec(?:_[0-9]+)?\b`)

type Executor = regionexec.Executor
type Factory = func([]uint256.Int, [4]uint64) regionexec.Executor

type Builder func() (string, error)
type Resolver func() (key, source string, err error)
type Slot struct {
	factory      bool
	factoryReady atomic.Pointer[Factory]
	template     *Slot
	values       []uint256.Int
	exits        [4]uint64
	key          string
	build        Builder
	resolve      Resolver // cold lowering, invoked only for an executed candidate
	artifactKey  string
	manager      *Manager
	priority     atomic.Int64 // preparation only; never read by an executor
	started      atomic.Bool
	alias        atomic.Pointer[Slot]
	ready        atomic.Pointer[Executor]
	finished     atomic.Bool
}

// ReadySlot publishes a statically linked executor through the same hot path
// as a loaded plugin. It does not create a manager or enter its work queue.
func ReadySlot(fn Executor) *Slot {
	s := new(Slot)
	s.ready.Store(&fn)
	s.started.Store(true)
	return s
}

// Published is the inlineable hot read. It never resolves factories or admits
// work; a caller that still needs preparation must use Ready on a miss.
func (s *Slot) Published() Executor {
	if p := s.ready.Load(); p != nil {
		return *p
	}
	return nil
}

func (s *Slot) Ready() Executor {
	if p := s.ready.Load(); p != nil {
		return *p
	}
	if s.template != nil {
		s.template.Ready()
		if f := s.template.factoryReady.Load(); f != nil {
			fn := (*f)(s.values, s.exits)
			s.ready.Store(&fn)
			return fn
		}
		return nil
	}
	if a := s.alias.Load(); a != nil {
		a.Ready()
		if f := a.factoryReady.Load(); f != nil {
			s.factoryReady.Store(f)
		}
		if fn := a.Ready(); fn != nil {
			s.ready.Store(&fn)
			return fn
		}
		return nil
	}
	if !s.started.Load() {
		s.manager.activate(s)
	}
	if p := s.ready.Load(); p != nil {
		return *p
	}
	return nil
}

// Loaded resolves an already loaded factory, but never admits compiler work.
// The background publisher uses this before making an immutable executable view.
func (s *Slot) Loaded() Executor {
	if p := s.ready.Load(); p != nil {
		return *p
	}
	if s.template != nil {
		s.template.Loaded()
		if f := s.template.factoryReady.Load(); f != nil {
			fn := (*f)(s.values, s.exits)
			s.ready.Store(&fn)
			return fn
		}
	} else if a := s.alias.Load(); a != nil {
		fn := a.Loaded()
		if f := a.factoryReady.Load(); f != nil {
			s.factoryReady.Store(f)
		}
		if fn != nil {
			s.ready.Store(&fn)
			return fn
		}
	}
	return nil
}

// Reuse attaches work another candidate has already admitted or completed.
// It never invokes a builder, allocates a queue entry, or spends compiler
// capacity. Only preparation workers use it; Published remains the hot read.
func (s *Slot) Reuse() Executor {
	if fn := s.Loaded(); fn != nil {
		return fn
	}
	if s.template != nil {
		s.template.Reuse()
	} else if a := s.alias.Load(); a != nil {
		a.Reuse()
	} else if !s.started.Load() && s.manager != nil {
		m := s.manager
		m.mu.Lock()
		if !s.started.Load() {
			m.reuseLocked(s)
		}
		m.mu.Unlock()
	}
	return s.Loaded()
}

// Settled also includes a failed/rejected preparation, which must remain native.
func (s *Slot) Settled() bool {
	if s.ready.Load() != nil || s.factoryReady.Load() != nil || s.finished.Load() {
		return true
	}
	if s.template != nil {
		return s.template.Settled()
	}
	if a := s.alias.Load(); a != nil {
		return a.Settled()
	}
	return false
}

// Only executed candidates enter the shared registry. Discarded binding
// alternatives must not retain their source strings or consume active slots.
func (m *Manager) activate(s *Slot) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if s.started.Load() {
		return
	}
	if m.reuseLocked(s) {
		return
	}
	m.admitLocked(s)
}

func (m *Manager) reuseLocked(s *Slot) bool {
	if old := m.slots[s.key]; old != nil && old != s {
		old.Prioritize(s.priority.Load())
		s.alias.Store(old)
		s.build = nil
		s.resolve = nil
		s.started.Store(true)
		if f := old.factoryReady.Load(); f != nil {
			s.factoryReady.Store(f)
		}
		if fn := old.ready.Load(); fn != nil {
			s.ready.Store(fn)
		}
		return true
	}
	if f := m.factories[s.key]; s.factory && f != nil {
		s.factoryReady.Store(&f)
		s.build = nil
		s.resolve = nil
		s.started.Store(true)
		return true
	}
	if fn := m.loaded[s.key]; fn != nil {
		s.ready.Store(&fn)
		s.build = nil
		s.resolve = nil
		s.started.Store(true)
		return true
	}
	return false
}

func (m *Manager) admitLocked(s *Slot) {
	if valueCompileMode == "none" {
		s.started.Store(true)
		s.finished.Store(true)
		s.build = nil
		s.resolve = nil
		return
	}
	// Soft admission bound: defer new compilation without rejecting its proof
	// or hiding a statically available executor. Shared slots resolve above.
	if m.hotcallCompileBlocked() {
		// Ranking must also apply at the soft outstanding/lifetime admission
		// bound, not only at the much larger generic queue-capacity bound.
		// Replace queued work only; keep the head reservation and all executing
		// jobs. Completed+pending and outstanding counts cannot grow here.
		if deliveryPriorityReplacement && len(m.queue) > 1 {
			m.pending.Add(1)
			if m.replaceQueued(s) {
				s.started.Store(true)
				m.slots[s.key] = s
				return
			}
			m.pending.Done()
		}
		return
	}
	if valueCompileMode != "selective" && os.Getenv("COLD_HOTSET") == "1" && len(m.slots) >= 128 {
		return
	}
	if len(m.slots) >= m.capacity {
		m.rejected++
		s.started.Store(true)
		s.build = nil
		s.resolve = nil
		s.finished.Store(true)
		return
	}
	m.pending.Add(1)
	s.started.Store(true)
	m.slots[s.key] = s
	if len(m.queue) < m.queueCapacity {
		m.queue = append(m.queue, s)
		select {
		case m.queueWake <- struct{}{}:
		default:
		}
	} else if m.replaceQueued(s) {
		// Existing bindings still own the displaced slot and can retry it.
	} else {
		delete(m.slots, s.key)
		s.started.Store(false)
		m.pending.Done()
	}
}

type Event struct {
	Keys                                                                                              []string
	Seconds                                                                                           float64
	GenerateSeconds, CompileSeconds, LoadSeconds, IdleSeconds, CompileCPUSeconds, CompilePauseSeconds float64 `json:",omitempty"`
	Error                                                                                             string
	SourceHash, BinaryHash                                                                            string
	CacheHit                                                                                          bool `json:",omitempty"`
}
type Manager struct {
	oracleArtifacts, oracleExecutors, oracleFactories int
	oraclePreloadNS                                   int64
	oracleMissing                                     map[string]bool
	oracleSources                                     map[string]oracleSource
	activity                                          *Activity
	priorityIdle                                      time.Duration // worker-owned idle time, excluded from pacing
	dir, root                                         string
	mu                                                sync.Mutex
	slots                                             map[string]*Slot
	loaded                                            map[string]Executor
	factories                                         map[string]Factory
	queue                                             []*Slot
	queueWake                                         chan struct{}
	queueCapacity                                     int
	pending                                           sync.WaitGroup
	events                                            []Event
	capacity, rejected                                int
	displaced                                         int
	duty                                              atomic.Int64
	draining                                          atomic.Int64
	budgetWake                                        chan struct{}
	completed                                         int
	revision                                          atomic.Uint64
	artifacts                                         map[string]*artifact
	attempted                                         map[string]bool
}

type artifact struct {
	stem  string
	event Event
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
	m := &Manager{activity: NewActivity(), dir: dir, root: root, capacity: capacity, slots: map[string]*Slot{}, loaded: map[string]Executor{}, factories: map[string]Factory{}, queueWake: make(chan struct{}, 1), queueCapacity: 2048, artifacts: map[string]*artifact{}, attempted: map[string]bool{}, budgetWake: make(chan struct{}, 1)}
	metas, _ := filepath.Glob(filepath.Join(dir, "*.json"))
	sort.Strings(metas)
	for _, f := range metas {
		var v Event
		b, e := os.ReadFile(f)
		if e != nil || json.Unmarshal(b, &v) != nil || v.Error != "" {
			continue
		}
		a := &artifact{strings.TrimSuffix(f, ".json"), v}
		for _, key := range v.Keys {
			m.artifacts[key] = a
		}
	}
	if err := m.experimentOracleInit(); err != nil {
		return nil, err
	}
	if !windowBoundaryCompiler {
		go m.worker()
	}
	return m, nil
}
func digest(b []byte) string { x := sha256.Sum256(b); return hex.EncodeToString(x[:]) }
func (m *Manager) Slot(key string, b Builder) *Slot {
	m.mu.Lock()
	defer m.mu.Unlock()
	if s := m.slots[key]; s != nil {
		return s
	}
	if fn := m.loaded[key]; fn != nil {
		s := &Slot{key: key, manager: m}
		s.ready.Store(&fn)
		s.started.Store(true)
		return s
	}
	return &Slot{key: key, build: b, manager: m}
}

// DeferredSlot deduplicates structural candidates before source formatting.
// The resolver returns the historical source-based key, preserving disk cache
// compatibility. It runs on the preparation worker, never on the EVM thread.
func (m *Manager) DeferredSlot(key string, resolve Resolver) *Slot {
	m.mu.Lock()
	defer m.mu.Unlock()
	if s := m.slots[key]; s != nil {
		return s
	}
	if fn := m.loaded[key]; fn != nil {
		s := &Slot{key: key, manager: m}
		s.ready.Store(&fn)
		s.started.Store(true)
		return s
	}
	return &Slot{key: key, resolve: resolve, manager: m}
}

// FactorySlot shares native code across immutable argument bindings. Only the
// first use requests compilation; pending or failed work leaves execution native.
func (m *Manager) FactorySlot(key string, build Builder, values []uint256.Int, exits [4]uint64) *Slot {
	m.mu.Lock()
	defer m.mu.Unlock()
	parent := m.slots[key]
	if parent == nil {
		parent = &Slot{key: key, build: build, manager: m, factory: true}
		if f := m.factories[key]; f != nil {
			parent.factoryReady.Store(&f)
			parent.started.Store(true)
			parent.build = nil
		}
	}
	return &Slot{template: parent, values: append([]uint256.Int(nil), values...), exits: exits}
}

func (s *Slot) cacheKey() string {
	if s.artifactKey != "" {
		return s.artifactKey
	}
	return s.key
}

// Revision advances after a preparation batch publishes success or failure.
// A consumer can avoid polling every pending slot between publications.
func (m *Manager) Revision() uint64 { return m.revision.Load() }

func (m *Manager) Activity() *Activity { return m.activity }
func (m *Manager) waitForForeground() {
	if m.activity != nil {
		t := time.Now()
		m.activity.Wait(nil, &m.draining)
		m.priorityIdle += time.Since(t)
	}
}

func (m *Manager) Drain() { done := m.Unthrottle(); defer done(); m.pending.Wait() }

// SetDutyPercent bounds average active preparation time with cooperative rests.
// Foreground coordination is independent of this soft budget; it is not an OS CPU quota.
func (m *Manager) SetDutyPercent(percent int) { m.duty.Store(int64(percent)) }
func (m *Manager) Unthrottle() func() {
	m.draining.Add(1)
	select {
	case m.budgetWake <- struct{}{}:
	default:
	}
	return func() { m.draining.Add(-1) }
}
func (m *Manager) rest(elapsed time.Duration) {
	p := m.duty.Load()
	if p <= 0 || p >= 100 || m.draining.Load() > 0 {
		return
	}
	t := time.NewTimer(elapsed * time.Duration(100-p) / time.Duration(p))
	defer t.Stop()
	for m.draining.Load() == 0 {
		select {
		case <-t.C:
			return
		case <-m.budgetWake:
		}
	}
}

// Usage exposes admission saturation so experiments cannot silently compare
// configurations with different coverage because they share a bounded cache.
func (m *Manager) Usage() (registered, limit, rejected int) {
	m.mu.Lock()
	defer m.mu.Unlock()
	return len(m.slots), m.capacity, m.rejected
}
func (m *Manager) Resources() map[string]int {
	m.mu.Lock()
	defer m.mu.Unlock()
	return map[string]int{"pending": len(m.slots), "completed": m.completed, "pendingCapacity": m.capacity, "rejected": m.rejected, "displaced": m.displaced, "loadedKeys": len(m.loaded), "loadedFamilies": len(m.factories), "indexedKeys": len(m.artifacts), "attemptedArtifacts": len(m.attempted), "oracleArtifacts": m.oracleArtifacts, "oracleExecutors": m.oracleExecutors, "oracleFactories": m.oracleFactories, "oraclePreloadNS": int(m.oraclePreloadNS), "oracleMissing": len(m.oracleMissing)}
}
func (m *Manager) Events() []Event {
	m.mu.Lock()
	defer m.mu.Unlock()
	return append([]Event(nil), m.events...)
}
func (m *Manager) load(path string) error {
	m.waitForForeground()
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
	var factories map[string]Factory
	if x, err := p.Lookup("Factories"); err == nil {
		f, ok := x.(*map[string]Factory)
		if !ok {
			return fmt.Errorf("unexpected factory ABI")
		}
		factories = *f
	}

	m.mu.Lock()
	defer m.mu.Unlock()
	for k, f := range factories {
		m.factories[k] = f
		if s := m.slots[k]; s != nil {
			fn := f
			s.factoryReady.Store(&fn)
		}
	}
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
	for range m.queueWake {
		// Keep the existing collection window. Reorder only admitted work;
		// keep a bounded queue and reserve its current head in each batch.
		timer := time.NewTimer(40 * time.Millisecond)
	collect:
		for {
			m.mu.Lock()
			n := len(m.queue)
			m.mu.Unlock()
			if n >= 128 {
				break
			}
			select {
			case _, ok := <-m.queueWake:
				if !ok {
					timer.Stop()
					return
				}
			case <-timer.C:
				break collect
			}
		}
		timer.Stop()
		m.waitForForeground()
		batch := m.takeBatch()
		if len(batch) == 0 {
			continue
		}
		elapsed := m.processBatch(batch)
		m.rest(elapsed)
	}
}

// A process indexes metadata at startup, then validates and loads only the
// artifacts requested by execution. The worker serializes plugin loading.
func (m *Manager) restore(s *Slot) bool {
	m.mu.Lock()
	fn := m.loaded[s.cacheKey()]
	factory := m.factories[s.cacheKey()]
	m.mu.Unlock()
	if s.factory && factory != nil {
		s.factoryReady.Store(&factory)
		return true
	}
	if fn != nil {
		s.ready.Store(&fn)
		return true
	}
	a := m.artifacts[s.cacheKey()]
	m.mu.Lock()
	if a == nil || m.attempted[a.stem] {
		m.mu.Unlock()
		return false
	}
	m.attempted[a.stem] = true
	m.mu.Unlock()
	started := time.Now()
	src, err := os.ReadFile(a.stem + ".go.txt")
	if err != nil || digest(src) != a.event.SourceHash {
		return false
	}
	bin, err := os.ReadFile(a.stem + ".so")
	if err != nil || digest(bin) != a.event.BinaryHash {
		return false
	}
	if err = m.load(a.stem + ".so"); err != nil {
		return false
	}
	m.mu.Lock()
	fn = m.loaded[s.cacheKey()]
	factory = m.factories[s.cacheKey()]
	found := fn != nil || s.factory && factory != nil
	if s.factory && factory != nil {
		s.factoryReady.Store(&factory)
	}
	if found {
		m.events = append(m.events, Event{Keys: []string{s.key}, Seconds: time.Since(started).Seconds(), CacheHit: true})
	}
	m.mu.Unlock()
	return found
}

func (m *Manager) compile(batch []*Slot) {
	if m.experimentNoCompile(batch) {
		return
	}
	start := time.Now()
	event := Event{}
	idleStart := m.priorityIdle
	var src strings.Builder
	src.WriteString("package main\nimport(\"math/bits\";\"github.com/holiman/uint256\";\"go-eth1.16.5-evm/regionexec\")\nvar _=bits.Mul64\nvar _ uint256.Int\n")
	var entries, factories []string
	for _, s := range batch {
		m.waitForForeground()
		body, err := s.build()
		if err != nil {
			m.mu.Lock()
			m.events = append(m.events, Event{Keys: []string{s.key}, Error: err.Error()})
			m.mu.Unlock()
			continue
		}
		key := s.cacheKey()
		name := "R" + digest([]byte(key))
		src.WriteString(executorName.ReplaceAllStringFunc(body, func(v string) string { return name + strings.TrimPrefix(v, "Exec") }))
		if s.factory {
			factories = append(factories, fmt.Sprintf("%q:%s", key, name))
		} else {
			entries = append(entries, fmt.Sprintf("%q:%s", key, name))
		}
		event.Keys = append(event.Keys, key)
	}
	if len(entries)+len(factories) == 0 {
		return
	}
	src.WriteString("\nvar Executors=map[string]regionexec.Executor{" + strings.Join(entries, ",") + "}\n")
	src.WriteString("\nvar Factories=map[string]func([]uint256.Int,[4]uint64)regionexec.Executor{" + strings.Join(factories, ",") + "}\n")
	data := []byte(src.String())
	event.SourceHash = digest(data)
	stem := filepath.Join(m.dir, event.SourceHash)
	work, err := os.MkdirTemp(m.dir, "compile-")
	if err != nil {
		event.Error = err.Error()
	} else {
		defer os.RemoveAll(work)
		file := filepath.Join(work, "main.go")
		err = os.WriteFile(file, data, 0600)
		if err == nil {
			args := []string{"build", "-buildmode=plugin"}
			if runtime.GOOS == "darwin" {
				args = append(args, "-ldflags=-extldflags=-Wl,-no_fixup_chains")
			}
			args = append(args, "-o", stem+".so", file)
			cmd := exec.Command(filepath.Join(runtime.GOROOT(), "bin", "go"), args...)
			cmd.Dir = m.root
			procs := "2"
			if m.duty.Load() > 0 {
				procs = "1"
			}
			cmd.Env = append(os.Environ(), "GOPROXY=off", "GOSUMDB=off", "GOMAXPROCS="+procs, "GOTOOLCHAIN=local")
			event.GenerateSeconds = time.Since(start).Seconds()
			ct := time.Now()
			out, paused, e := m.compilerOutput(cmd, 90*time.Second)
			m.priorityIdle += paused
			event.CompilePauseSeconds = paused.Seconds()
			if cmd.ProcessState != nil {
				event.CompileCPUSeconds = (cmd.ProcessState.UserTime() + cmd.ProcessState.SystemTime()).Seconds()
			}
			event.CompileSeconds = time.Since(ct).Seconds()
			if e != nil {
				err = fmt.Errorf("compile: %w: %.1200s", e, out)
			}
		}
		if err == nil {
			lt := time.Now()
			idleBefore := m.priorityIdle
			err = m.load(stem + ".so")
			event.LoadSeconds = (time.Since(lt) - (m.priorityIdle - idleBefore)).Seconds()
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
	event.IdleSeconds = (m.priorityIdle - idleStart).Seconds()
	raw, _ := json.MarshalIndent(event, "", "  ")
	os.WriteFile(stem+".json", raw, 0600)
	m.mu.Lock()
	m.events = append(m.events, event)
	m.mu.Unlock()
}

func (m *Manager) processBatch(batch []*Slot) time.Duration {
	started := time.Now()
	m.waitForForeground()
	started = time.Now() // idle scheduling wait is not active budget consumption
	idleBefore := m.priorityIdle
	pending := make([]*Slot, 0, len(batch))
	seen := make(map[string]bool)
	for _, s := range batch {
		m.waitForForeground()
		if s.resolve != nil {
			key, source, err := s.resolve()
			s.resolve = nil
			if err != nil {
				m.mu.Lock()
				m.events = append(m.events, Event{Keys: []string{s.key}, Error: err.Error()})
				m.mu.Unlock()
				continue
			}
			s.artifactKey = key
			s.build = func() (string, error) { return source, nil }
		}
		if !m.restore(s) {
			if key := s.cacheKey(); !seen[key] {
				pending = append(pending, s)
				seen[key] = true
			}
		}
	}
	if len(pending) > 0 {
		m.compile(pending)
	}
	for _, s := range batch {
		m.mu.Lock()
		if f := m.factories[s.cacheKey()]; s.factory && f != nil {
			m.factories[s.key] = f
			s.factoryReady.Store(&f)
		}
		if fn := m.loaded[s.cacheKey()]; fn != nil {
			m.loaded[s.key] = fn
			s.ready.Store(&fn)
		}
		// Ready slots are owned by their published bindings. The registry is
		// only for pending deduplication, not an ever-growing candidate list.
		delete(m.slots, s.key)
		m.completed++
		m.mu.Unlock()
		s.build = nil
		s.finished.Store(true)
		m.pending.Done()
	}
	m.revision.Add(1)
	return time.Since(started) - (m.priorityIdle - idleBefore)
}
