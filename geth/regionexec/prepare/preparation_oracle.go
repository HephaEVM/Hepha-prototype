package prepare

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"sync/atomic"
	"time"
)

// Diagnostic only: executable artifacts are supplied using prior observations
// of this very workload. Analysis/binding caches remain empty. This is not a
// deployment claim or a held-out test. No mutable transaction data is cached.
var oracleSuppressCompiler atomic.Bool

type oracleSource struct {
	Key, Source string
	Factory     bool
}

func (m *Manager) experimentOracleInit() error {
	dir := os.Getenv("ORACLE_ARTIFACTS")
	if dir == "" {
		return nil
	}
	oracleSuppressCompiler.Store(true)
	started := time.Now()
	var files []string
	for _, part := range filepath.SplitList(dir) {
		fs, err := filepath.Glob(filepath.Join(part, "*.json"))
		if err != nil {
			return err
		}
		files = append(files, fs...)
	}
	sort.Strings(files)
	seen := map[string]bool{}
	for _, p := range files {
		b, err := os.ReadFile(p)
		if err != nil {
			return err
		}
		var e Event
		if json.Unmarshal(b, &e) != nil || e.Error != "" || e.SourceHash == "" || e.BinaryHash == "" {
			continue
		}
		if seen[e.SourceHash] {
			continue
		}
		seen[e.SourceHash] = true
		stem := p[:len(p)-5]
		src, err := os.ReadFile(stem + ".go.txt")
		if err != nil || digest(src) != e.SourceHash {
			return fmt.Errorf("oracle source integrity: %s", p)
		}
		bin, err := os.ReadFile(stem + ".so")
		if err != nil || digest(bin) != e.BinaryHash {
			return fmt.Errorf("oracle binary integrity: %s", p)
		}
		if err = m.load(stem + ".so"); err != nil {
			return fmt.Errorf("oracle load %s: %w", p, err)
		}
		m.oracleArtifacts++
	}
	m.oracleExecutors = len(m.loaded)
	m.oracleFactories = len(m.factories)
	m.oraclePreloadNS = time.Since(started).Nanoseconds()
	if m.oracleArtifacts == 0 {
		return fmt.Errorf("oracle corpus has no verified artifacts")
	}
	return nil
}
func (m *Manager) experimentNoCompile(batch []*Slot) bool {
	if !oracleSuppressCompiler.Load() {
		return false
	}
	if os.Getenv("ORACLE_COLLECT") == "1" {
		if m.oracleSources == nil {
			m.oracleSources = map[string]oracleSource{}
		}
		for _, s := range batch {
			key := s.cacheKey()
			if _, ok := m.oracleSources[key]; !ok {
				body, err := s.build()
				if err != nil {
					panic(err)
				}
				m.oracleSources[key] = oracleSource{key, body, s.factory}
			}
		}
	}
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.oracleMissing == nil {
		m.oracleMissing = map[string]bool{}
	}
	for _, s := range batch {
		m.oracleMissing[s.cacheKey()] = true
	}
	return true
}

// Called only after online snapshots are frozen and the pending queue drained.
// Missing executors may then be generated for the separate warm reference.
func ExperimentOracleEnd() { oracleSuppressCompiler.Store(false) }
func (m *Manager) ExperimentOracleMisses() error {
	m.mu.Lock()
	keys := make([]string, 0, len(m.oracleMissing))
	for k := range m.oracleMissing {
		keys = append(keys, k)
	}
	m.mu.Unlock()
	sort.Strings(keys)
	b, err := json.Marshal(keys)
	if err != nil {
		return err
	}
	if err = os.WriteFile(filepath.Join(m.dir, "oracle-missing-keys.json"), b, 0600); err != nil {
		return err
	}
	if len(m.oracleSources) > 0 {
		rows := make([]oracleSource, 0, len(m.oracleSources))
		for _, k := range keys {
			if s, ok := m.oracleSources[k]; ok {
				rows = append(rows, s)
			}
		}
		b, err = json.Marshal(rows)
		if err != nil {
			return err
		}
		return os.WriteFile(filepath.Join(m.dir, "oracle-missing-sources.json"), b, 0600)
	}
	return nil
}
