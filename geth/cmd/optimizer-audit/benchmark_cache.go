package main

import (
	"bytes"
	"errors"
	"sync"

	"go-eth1.16.5-evm/ethdb"
)

var errBenchmarkColdRead = errors.New("benchmark attempted an unprepared disk read")

type benchmarkReadValue struct {
	value []byte
	err   error
}
type benchmarkHasValue struct {
	present bool
	err     error
}

// This cache is only valid for an immutable read-only database. State mutations
// live in overlayDB and take precedence over this cache. During measurement a
// miss fails closed: it never falls through to the disk, even for absent keys.
type benchmarkReadCache struct {
	mu                 sync.Mutex
	source             ethdb.KeyValueReader
	values             map[string]benchmarkReadValue
	presence           map[string]benchmarkHasValue
	strict             bool
	diskReads, blocked uint64
}

func newBenchmarkReadCache(source ethdb.KeyValueReader) *benchmarkReadCache {
	return &benchmarkReadCache{source: source, values: make(map[string]benchmarkReadValue), presence: make(map[string]benchmarkHasValue)}
}
func (c *benchmarkReadCache) Get(key []byte) ([]byte, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if v, ok := c.values[string(key)]; ok {
		return bytes.Clone(v.value), v.err
	}
	if c.strict {
		c.blocked++
		return nil, errBenchmarkColdRead
	}
	c.diskReads++
	v, err := c.source.Get(key)
	c.values[string(key)] = benchmarkReadValue{bytes.Clone(v), err}
	if err == nil {
		c.presence[string(key)] = benchmarkHasValue{true, nil}
	}
	return v, err
}
func (c *benchmarkReadCache) Has(key []byte) (bool, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if v, ok := c.presence[string(key)]; ok {
		return v.present, v.err
	}
	if c.strict {
		c.blocked++
		return false, errBenchmarkColdRead
	}
	c.diskReads++
	v, err := c.source.Has(key)
	c.presence[string(key)] = benchmarkHasValue{v, err}
	return v, err
}
func (c *benchmarkReadCache) setStrict(strict bool) { c.mu.Lock(); c.strict = strict; c.mu.Unlock() }
func (c *benchmarkReadCache) counts() (uint64, uint64) {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.diskReads, c.blocked
}
func (c *benchmarkReadCache) forbidIO() bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.strict {
		c.blocked++
		return true
	}
	return false
}

type benchmarkChainDB struct {
	ethdb.Database
	reads *benchmarkReadCache
}

func (d *benchmarkChainDB) Get(k []byte) ([]byte, error) { return d.reads.Get(k) }
func (d *benchmarkChainDB) Has(k []byte) (bool, error)   { return d.reads.Has(k) }
func (d *benchmarkChainDB) ReadAncients(fn func(ethdb.AncientReaderOp) error) error {
	// HeaderChain is warmed by full block execution; no ancient storage access
	// is allowed during timing. An unexpected access invalidates the run.
	if d.reads.forbidIO() {
		return errBenchmarkColdRead
	}
	return d.Database.ReadAncients(fn)
}
