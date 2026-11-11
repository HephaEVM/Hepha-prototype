package vm

import (
	"encoding/json"
	"github.com/golang/snappy"
	"path/filepath"
	"testing"
)

func TestFlowPreparationCacheRecovery(t *testing.T) {
	path := filepath.Join(t.TempDir(), "lowering")
	c, err := openFlowPreparationCache(path)
	if err != nil {
		t.Fatal(err)
	}
	ps := prepareFlows(preparationPaths(t))
	key, refs := flowPreparationKey(ps, true, true)
	src, helpers, err := emitPreparedFlowShared(ps, true)
	if err != nil {
		t.Fatal(err)
	}
	rec := flowPreparationRecord{Source: src, Helpers: flowPreparationHelpers(ps, refs, true, true), Reports: make([]flowReport, len(ps))}
	c.put(key, rec)
	size := c.bytes
	if err = c.db.Close(); err != nil {
		t.Fatal(err)
	}
	c, err = openFlowPreparationCache(path)
	if err != nil {
		t.Fatal(err)
	}
	defer c.db.Close()
	got, ok := c.get(key, refs, len(ps), true)
	if !ok || got.Source != src || len(got.bindHelpers(refs)[0]) != len(helpers) || c.bytes != size {
		t.Fatal("reopen lost preparation")
	}
	c.put(key, rec)
	if c.bytes != size {
		t.Fatal("duplicate counted twice")
	}
	other, _ := flowPreparationKey(ps, false, true)
	if _, ok = c.get(other, refs, len(ps), true); ok {
		t.Fatal("proof mode reused")
	}
	other, _ = flowPreparationKey(ps, true, false)
	if _, ok = c.get(other, refs, len(ps), false); ok {
		t.Fatal("layout reused")
	}
	if _, ok = c.get(key, refs, len(ps)+1, true); ok {
		t.Fatal("path count mismatch admitted")
	}
	// Invalid records must fall back to analysis, never publish partial helpers.
	bad := got
	bad.Helpers = nil
	raw, _ := json.Marshal(bad)
	c.db.Put(key, snappy.Encode(nil, raw), nil)
	if _, ok = c.get(key, refs, len(ps), true); ok {
		t.Fatal("invalid helper shape admitted")
	}
	bad = got
	bad.Source += "invalid"
	raw, _ = json.Marshal(bad)
	c.db.Put(key, snappy.Encode(nil, raw), nil)
	if _, ok = c.get(key, refs, len(ps), true); ok {
		t.Fatal("modified source admitted")
	}
	neg := []byte("record/rejected")
	c.put(neg, flowPreparationRecord{Error: "non-guard divergence"})
	if v, ok := c.get(neg, nil, 2, true); !ok || v.Error == "" {
		t.Fatal("negative cache lost")
	}
	c.bytes = flowPreparationCacheLimit
	c.put([]byte("record/full"), rec)
	if _, ok = c.get([]byte("record/full"), refs, len(ps), true); ok {
		t.Fatal("capacity exceeded")
	}
}
