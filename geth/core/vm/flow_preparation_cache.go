package vm

// This is a cache of deterministic lowering, not a cache of transaction values.
// Rebinding still derives the structural regions from the supplied bytecode.
// The embedded lowering implementation invalidates records independently of
// benchmark harness changes. Native executors retain their own ABI validation.
import (
	"crypto/sha256"
	"embed"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"sync"
	"sync/atomic"

	"github.com/golang/snappy"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/opt"
	"go-eth1.16.5-evm/regionexec"
	"go-eth1.16.5-evm/regionexec/prepare"
)

const flowPreparationCacheLimit = 1 << 30

// Include every lowering/proof implementation; changes to region operands,
// gas, helper semantics or layouts are additionally covered by each record key.
//
//go:embed flow_preparation_cache.go flow_runtime.go flow_prepared.go flow_shared.go flow_codegen.go flow_proof.go flow_helper.go hot_region.go hot_region_optimize.go demand_codegen.go demand_memory.go opcodes.go
var flowPreparationSources embed.FS

func flowPreparationVersion() string {
	h := sha256.New()
	fmt.Fprintln(h, "lowering-v2", runtime.Version())
	fmt.Fprintln(h, "extended-ce-v1", paperCEExtendedProof)
	files, _ := flowPreparationSources.ReadDir(".")
	for _, f := range files {
		b, _ := flowPreparationSources.ReadFile(f.Name())
		fmt.Fprintln(h, f.Name())
		h.Write(b)
	}
	return fmt.Sprintf("%x", h.Sum(nil))
}

type flowPreparationRecord struct {
	Source, Digest, Error string
	Helpers               [][]int
	Reports               []flowReport
}

type flowPreparationCache struct {
	db    *leveldb.DB
	mu    sync.Mutex
	bytes uint64
}

var flowCacheOnce sync.Once
var flowCache *flowPreparationCache
var flowCacheHits, flowCacheMisses, flowCacheWrites atomic.Uint64

func PreparationCacheStats() map[string]uint64 {
	s := map[string]uint64{"hits": flowCacheHits.Load(), "misses": flowCacheMisses.Load(), "writes": flowCacheWrites.Load(), "limitBytes": flowPreparationCacheLimit}
	if c := getFlowPreparationCache(); c != nil {
		c.mu.Lock()
		s["storedBytes"] = c.bytes
		c.mu.Unlock()
	}
	return s
}

func getFlowPreparationCache() *flowPreparationCache {
	flowCacheOnce.Do(func() {
		dir := os.Getenv("DEMAND_DIR")
		if dir == "" {
			return
		}
		path := filepath.Join(dir, "lowering-v2-"+flowPreparationVersion())
		flowCache, _ = openFlowPreparationCache(path)
	})
	return flowCache
}

func openFlowPreparationCache(path string) (*flowPreparationCache, error) {
	db, err := leveldb.OpenFile(path, &opt.Options{BlockCacheCapacity: 16 << 20, WriteBuffer: 8 << 20, OpenFilesCacheCapacity: 64})
	if err != nil {
		return nil, err
	}
	c := &flowPreparationCache{db: db}
	if raw, err := db.Get([]byte("bytes"), nil); err == nil && len(raw) == 8 {
		c.bytes = binary.LittleEndian.Uint64(raw)
	}
	return c, nil
}

// Include helper aliasing across paths: shared lowering deduplicates helpers
// by identity, even when two helpers have identical arithmetic semantics.
func flowPreparationKey(ps []*preparedFlow, proof, shared bool) ([]byte, []*arithmeticRegion) {
	h := sha256.New()
	fmt.Fprintf(h, "flow-lowering-v1/%t/%t/", proof, shared)
	if paperWithoutReuse && shared {
		fmt.Fprint(h, "independent-prefixes-v1/")
	}
	if paperWithoutCheckElimination {
		fmt.Fprint(h, "checks-retained/")
	}
	ids := map[*arithmeticRegion]int{}
	var refs []*arithmeticRegion
	for _, p := range ps {
		r := p.original
		// Builders publish immutable paths with their full semantic fingerprint.
		key := r.key
		if key == "" {
			key = r.fingerprint()
		}
		fmt.Fprintf(h, "%s;", key)
		for _, n := range r.nodes {
			if a := n.helper; a != nil {
				id, ok := ids[a]
				if !ok {
					id = len(refs)
					ids[a] = id
					refs = append(refs, a)
				}
				fmt.Fprintf(h, "%d:%d;", id, a.syncHot.admissionGas)
			}
		}
	}
	return append([]byte("record/"), h.Sum(nil)...), refs
}

func (c *flowPreparationCache) get(key []byte, refs []*arithmeticRegion, paths int, shared bool) (flowPreparationRecord, bool) {
	if c == nil {
		return flowPreparationRecord{}, false
	}
	raw, err := c.db.Get(key, nil)
	if err != nil {
		flowCacheMisses.Add(1)
		return flowPreparationRecord{}, false
	}
	var v flowPreparationRecord
	raw, err = snappy.Decode(nil, raw)
	if err != nil {
		return v, false
	}
	if json.Unmarshal(raw, &v) != nil {
		return v, false
	}
	if v.Error != "" {
		flowCacheHits.Add(1)
		return v, true
	}
	helperRows := paths
	if shared {
		helperRows = 1
	}
	if v.Source == "" || v.Digest != fmt.Sprintf("%x", sha256.Sum256([]byte(v.Source))) || len(v.Reports) != paths || len(v.Helpers) != helperRows {
		return v, false
	}
	for _, row := range v.Helpers {
		for _, id := range row {
			if id < 0 || id >= len(refs) || refs[id] == nil || refs[id].syncHot.native == nil {
				return v, false
			}
		}
	}
	flowCacheHits.Add(1)
	return v, true
}

func (c *flowPreparationCache) put(key []byte, v flowPreparationRecord) bool {
	if c == nil {
		return false
	}
	if v.Source != "" {
		v.Digest = fmt.Sprintf("%x", sha256.Sum256([]byte(v.Source)))
	}
	raw, err := json.Marshal(v)
	if err != nil {
		return false
	}
	raw = snappy.Encode(nil, raw)
	c.mu.Lock()
	defer c.mu.Unlock()
	old, _ := c.db.Get(key, nil)
	if string(old) == string(raw) {
		return true
	}
	oldSize := uint64(0)
	if len(old) > 0 {
		oldSize = uint64(len(key) + len(old))
	}
	size := uint64(len(key) + len(raw))
	base := c.bytes - min(c.bytes, oldSize)
	if base+size > flowPreparationCacheLimit {
		return false
	}
	var amount [8]byte
	binary.LittleEndian.PutUint64(amount[:], base+size)
	batch := new(leveldb.Batch)
	batch.Put(key, raw)
	batch.Put([]byte("bytes"), amount[:])
	if c.db.Write(batch, nil) == nil {
		c.bytes = base + size
		flowCacheWrites.Add(1)
		return true
	}
	return false
}

// Retain only the compact key in dormant candidates, not their generated
// source. Source is fetched only if an executed plan needs a native executor.
func (c *flowPreparationCache) sourceBuilder(key []byte) prepare.Builder {
	key = append([]byte(nil), key...)
	return func() (string, error) {
		raw, err := c.db.Get(key, nil)
		if err != nil {
			return "", err
		}
		var v flowPreparationRecord
		raw, err = snappy.Decode(nil, raw)
		if err != nil {
			return "", err
		}
		if err = json.Unmarshal(raw, &v); err != nil {
			return "", err
		}
		if v.Source == "" || v.Digest != fmt.Sprintf("%x", sha256.Sum256([]byte(v.Source))) {
			return "", fmt.Errorf("invalid cached lowering")
		}
		return v.Source, nil
	}
}

func flowPreparationHelpers(ps []*preparedFlow, refs []*arithmeticRegion, proof, shared bool) [][]int {
	ids := map[*arithmeticRegion]int{}
	for i, h := range refs {
		ids[h] = i
	}
	var out [][]int
	seen := map[*arithmeticRegion]bool{}
	if shared {
		out = append(out, []int{})
	}
	for _, p := range ps {
		if !shared {
			out = append(out, []int{})
		}
		r, _ := p.optimized(proof)
		for _, n := range r.nodes {
			if h := n.helper; h != nil && (!shared || !seen[h]) {
				out[len(out)-1] = append(out[len(out)-1], ids[h])
				seen[h] = true
			}
		}
	}
	return out
}

func (v flowPreparationRecord) bindHelpers(refs []*arithmeticRegion) [][]regionexec.Helper {
	out := make([][]regionexec.Helper, len(v.Helpers))
	for i, row := range v.Helpers {
		for _, id := range row {
			out[i] = append(out[i], regionexec.Helper(refs[id].syncHot.native))
		}
	}
	return out
}
