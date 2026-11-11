package main

import (
	"bytes"
	"sync"

	goleveldb "github.com/syndtr/goleveldb/leveldb"
	"go-eth1.16.5-evm/database"
	"go-eth1.16.5-evm/ethdb"
	"go-eth1.16.5-evm/ethdb/leveldb"
	"go-eth1.16.5-evm/ethdb/memorydb"
)

// Isolated warm-call comparisons use the budget recorded in their manifest.
var isolatedCallComparison bool

// Read through the project's LevelDB wrapper; commit mutations to memory only.
type overlayDB struct {
	*memorydb.Database
	disk    *leveldb.Database
	mu      sync.RWMutex
	deleted map[string]bool
	// Opt-in immutable disk read cache, used only by the paired benchmark.
	reads *benchmarkReadCache
}

func openOverlay(path string) (*overlayDB, error) {
	var disk *leveldb.Database
	var err error
	if isolatedCallComparison {
		disk, err = leveldb.New(path, 512, 1024, "state_snapshot", true)
	} else {
		disk, err = database.OpenStateDatabase(path, true)
	}
	if err != nil {
		return nil, err
	}
	return &overlayDB{Database: memorydb.New(), disk: disk, deleted: map[string]bool{}}, nil
}
func (d *overlayDB) Get(k []byte) ([]byte, error) {
	d.mu.RLock()
	defer d.mu.RUnlock()
	if d.deleted[string(k)] {
		return nil, goleveldb.ErrNotFound
	}
	if ok, _ := d.Database.Has(k); ok {
		return d.Database.Get(k)
	}
	if d.reads != nil {
		return d.reads.Get(k)
	}
	return d.disk.Get(k)
}
func (d *overlayDB) Has(k []byte) (bool, error) {
	d.mu.RLock()
	defer d.mu.RUnlock()
	if d.deleted[string(k)] {
		return false, nil
	}
	if ok, _ := d.Database.Has(k); ok {
		return true, nil
	}
	if d.reads != nil {
		return d.reads.Has(k)
	}
	return d.disk.Has(k)
}
func (d *overlayDB) Put(k, v []byte) error {
	d.mu.Lock()
	defer d.mu.Unlock()
	delete(d.deleted, string(k))
	return d.Database.Put(k, v)
}
func (d *overlayDB) Delete(k []byte) error {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.deleted[string(k)] = true
	return d.Database.Delete(k)
}
func (d *overlayDB) Close() error                       { _ = d.Database.Close(); return d.disk.Close() }
func (d *overlayDB) NewBatch() ethdb.Batch              { return &overlayBatch{db: d} }
func (d *overlayDB) NewBatchWithSize(_ int) ethdb.Batch { return d.NewBatch() }
func (d *overlayDB) DeleteRange(start, end []byte) error {
	it := d.NewIterator(nil, start)
	defer it.Release()
	for it.Next() {
		if end != nil && bytes.Compare(it.Key(), end) >= 0 {
			break
		}
		if err := d.Delete(bytes.Clone(it.Key())); err != nil {
			return err
		}
	}
	return it.Error()
}
func (d *overlayDB) NewIterator(prefix, start []byte) ethdb.Iterator {
	merged := memorydb.New()
	if d.reads != nil && d.reads.forbidIO() {
		return &overlayIterator{Iterator: merged.NewIterator(prefix, start), store: merged, err: errBenchmarkColdRead}
	}
	d.mu.RLock()
	defer d.mu.RUnlock()
	it := d.disk.NewIterator(prefix, start)
	for it.Next() {
		if !d.deleted[string(it.Key())] {
			_ = merged.Put(it.Key(), it.Value())
		}
	}
	err := it.Error()
	it.Release()
	mem := d.Database.NewIterator(prefix, start)
	for mem.Next() {
		_ = merged.Put(mem.Key(), mem.Value())
	}
	if err == nil {
		err = mem.Error()
	}
	mem.Release()
	return &overlayIterator{Iterator: merged.NewIterator(prefix, start), store: merged, err: err}
}

type overlayIterator struct {
	ethdb.Iterator
	store *memorydb.Database
	err   error
}

func (i *overlayIterator) Error() error {
	if i.err != nil {
		return i.err
	}
	return i.Iterator.Error()
}
func (i *overlayIterator) Release() { i.Iterator.Release(); _ = i.store.Close() }

type batchOp struct {
	k, v []byte
	del  bool
}
type overlayBatch struct {
	db   *overlayDB
	ops  []batchOp
	size int
}

func (b *overlayBatch) Put(k, v []byte) error {
	b.ops = append(b.ops, batchOp{bytes.Clone(k), bytes.Clone(v), false})
	b.size += len(k) + len(v)
	return nil
}
func (b *overlayBatch) Delete(k []byte) error {
	b.ops = append(b.ops, batchOp{k: bytes.Clone(k), del: true})
	b.size += len(k)
	return nil
}
func (b *overlayBatch) DeleteRange(start, end []byte) error {
	it := b.db.NewIterator(nil, start)
	defer it.Release()
	count := 0
	for it.Next() {
		if end != nil && bytes.Compare(it.Key(), end) >= 0 {
			break
		}
		count++
		if count > 10000 {
			return ethdb.ErrTooManyKeys
		}
		if err := b.Delete(it.Key()); err != nil {
			return err
		}
	}
	return it.Error()
}
func (b *overlayBatch) ValueSize() int { return b.size }
func (b *overlayBatch) Reset()         { b.ops = nil; b.size = 0 }
func (b *overlayBatch) Write() error   { return b.Replay(b.db) }
func (b *overlayBatch) Replay(w ethdb.KeyValueWriter) error {
	for _, op := range b.ops {
		var err error
		if op.del {
			err = w.Delete(op.k)
		} else {
			err = w.Put(op.k, op.v)
		}
		if err != nil {
			return err
		}
	}
	return nil
}
