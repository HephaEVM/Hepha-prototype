package main

import (
	"errors"
	"testing"

	"go-eth1.16.5-evm/ethdb/memorydb"
)

func TestBenchmarkReadCacheStrictAndOwnedBytes(t *testing.T) {
	db := memorydb.New()
	defer db.Close()
	_ = db.Put([]byte("key"), []byte("value"))
	c := newBenchmarkReadCache(db)
	first, err := c.Get([]byte("key"))
	if err != nil {
		t.Fatal(err)
	}
	first[0] = 'x'
	_, absentErr := c.Get([]byte("absent"))
	if absentErr == nil {
		t.Fatal("expected absent")
	}
	_, _ = c.Has([]byte("absent"))
	diskBefore, _ := c.counts()
	c.setStrict(true)
	value, err := c.Get([]byte("key"))
	if err != nil || string(value) != "value" {
		t.Fatalf("cache alias: %q %v", value, err)
	}
	value[0] = 'z'
	value, _ = c.Get([]byte("key"))
	if string(value) != "value" {
		t.Fatal("returned cache buffer was mutable")
	}
	if ok, err := c.Has([]byte("key")); !ok || err != nil {
		t.Fatal(ok, err)
	}
	if _, err := c.Get([]byte("absent")); err != absentErr {
		t.Fatal("negative lookup not cached")
	}
	if ok, err := c.Has([]byte("absent")); ok || err != nil {
		t.Fatal(ok, err)
	}
	if _, err := c.Get([]byte("cold")); !errors.Is(err, errBenchmarkColdRead) {
		t.Fatal(err)
	}
	if _, err := c.Has([]byte("cold")); !errors.Is(err, errBenchmarkColdRead) {
		t.Fatal(err)
	}
	diskAfter, blocked := c.counts()
	if diskAfter != diskBefore || blocked != 2 {
		t.Fatal(diskBefore, diskAfter, blocked)
	}
}

func TestBenchmarkOverlayWritesTakePrecedence(t *testing.T) {
	base := memorydb.New()
	defer base.Close()
	_ = base.Put([]byte("key"), []byte("old"))
	d := &overlayDB{Database: memorydb.New(), deleted: map[string]bool{}, reads: newBenchmarkReadCache(base)}
	defer d.Database.Close()
	_, _ = d.Get([]byte("key"))
	d.reads.setStrict(true)
	batch := d.NewBatch()
	_ = batch.Put([]byte("key"), []byte("new"))
	_ = batch.Write()
	if value, err := d.Get([]byte("key")); err != nil || string(value) != "new" {
		t.Fatal(string(value), err)
	}
	_ = d.Delete([]byte("key"))
	if ok, err := d.Has([]byte("key")); ok || err != nil {
		t.Fatal(ok, err)
	}
	if _, err := d.Get([]byte("key")); err == nil {
		t.Fatal("deleted base value resurfaced")
	}
	_ = d.Put([]byte("key"), []byte("again"))
	if value, err := d.Get([]byte("key")); err != nil || string(value) != "again" {
		t.Fatal(string(value), err)
	}
}
