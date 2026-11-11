package main

import (
	"bytes"
	"testing"

	"github.com/syndtr/goleveldb/leveldb"
)

// A replay may commit new state, but the supplied snapshot must stay unchanged.
func TestOverlayLeavesSnapshotUnchanged(t *testing.T) {
	path := t.TempDir()
	disk, err := leveldb.OpenFile(path, nil)
	if err != nil {
		t.Fatal(err)
	}
	for k, v := range map[string]string{"account": "old", "deleted": "keep"} {
		if err := disk.Put([]byte(k), []byte(v), nil); err != nil {
			t.Fatal(err)
		}
	}
	if err := disk.Close(); err != nil {
		t.Fatal(err)
	}
	overlay, err := openOverlay(path)
	if err != nil {
		t.Fatal(err)
	}
	batch := overlay.NewBatch()
	_ = batch.Put([]byte("account"), []byte("new"))
	_ = batch.Put([]byte("created"), []byte("temporary"))
	_ = batch.Delete([]byte("deleted"))
	if err := batch.Write(); err != nil {
		t.Fatal(err)
	}
	got, err := overlay.Get([]byte("account"))
	if err != nil || !bytes.Equal(got, []byte("new")) {
		t.Fatalf("overlay read: %q, %v", got, err)
	}
	if has, err := overlay.Has([]byte("deleted")); err != nil || has {
		t.Fatalf("overlay deletion: %v, %v", has, err)
	}
	if err := overlay.Close(); err != nil {
		t.Fatal(err)
	}
	disk, err = leveldb.OpenFile(path, nil)
	if err != nil {
		t.Fatal(err)
	}
	defer disk.Close()
	for k, want := range map[string]string{"account": "old", "deleted": "keep"} {
		got, err := disk.Get([]byte(k), nil)
		if err != nil || string(got) != want {
			t.Fatalf("snapshot changed at %s: %q, %v", k, got, err)
		}
	}
	if has, err := disk.Has([]byte("created"), nil); err != nil || has {
		t.Fatalf("new state leaked into snapshot: %v, %v", has, err)
	}
}
