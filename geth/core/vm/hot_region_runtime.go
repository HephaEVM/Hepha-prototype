package vm

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"os"
)

var hotRegionTable *JumpTable

func init() { hotRegionTable = &pragueInstructionSet }

type HotRegionStats struct{ Commits, Instructions, Helpers uint64 }

func (a HotRegionStats) Sub(b HotRegionStats) HotRegionStats {
	return HotRegionStats{a.Commits - b.Commits, a.Instructions - b.Instructions, a.Helpers - b.Helpers}
}

// Immutable library/namespace, with runner-local counters; one sequential owner.
type HotRegionRuntime struct {
	codes  map[common.Hash]*hotRegionPrepared
	engine *ArithmeticEngine
	Stats  HotRegionStats
}
type hotRegionExecutable struct {
	region *hotRegion
	fn     hotRegionFunc
}
type hotRegionPrepared struct {
	entries []hotRegionExecutable
	code    []byte
}

func (evm *EVM) ordinaryLibrary(c *ArithmeticConfig) *HotRegionRuntime {
	p := evm.Config.HotRegions
	if p == nil || p.engine != c.Engine || evm.table != hotRegionTable || c.Mode != ArithmeticSync || c.GasPolicy != OriginalArithmeticGas {
		return nil
	}
	return p
}

// Cold path ONLY. Publish the final table once; never mutate a cached binding.
func (p *HotRegionRuntime) bind(b *boundArithmetic, code []byte, key arithmeticCacheKey) {
	selected := p.codes[key.hash]
	if selected == nil || key.rules != hotRegionTable || !bytes.Equal(code, selected.code) {
		return
	}
	if b.dispatchCode == nil {
		b.dispatchCode = append([]byte(nil), code...)
		b.entryIndex = make([]uint16, len(code))
		b.entryRegions = make([]*arithmeticRegion, 1)
	}
	if b.computeEntries == nil {
		b.computeEntries = make([]*computeHelperRegion, len(b.entryRegions))
	}
	for pc, e := range selected.entries {
		if e.fn == nil {
			continue
		}
		id := b.entryIndex[pc]
		if id == 0 {
			if len(b.entryRegions) >= 1<<16 {
				continue
			}
			id = uint16(len(b.entryRegions))
			b.entryRegions = append(b.entryRegions, nil)
			b.computeEntries = append(b.computeEntries, nil)
			b.entryIndex[pc] = id
		}
		r := new(computeHelperRegion)
		if old := b.computeEntries[id]; old != nil {
			*r = *old
		}
		entry := e
		r.ordinary = &entry
		b.computeEntries[id] = r
		b.dispatchCode[pc] = byte(arithmeticEntryMarker)
	}
}
func NewHotRegionRuntime(e *ArithmeticEngine, path, selected string) (*HotRegionRuntime, error) {
	if paperAblationLevel == 1 {
		return &HotRegionRuntime{engine: e, codes: map[common.Hash]*hotRegionPrepared{}}, nil
	}
	raw, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var corpus struct{ Codes map[string]string }
	if err = json.Unmarshal(raw, &corpus); err != nil {
		return nil, err
	}
	raw, err = os.ReadFile(selected)
	if err != nil {
		return nil, err
	}
	var choices []*HotRegionEvidence
	if err = json.Unmarshal(raw, &choices); err != nil {
		return nil, err
	}
	out := &HotRegionRuntime{codes: map[common.Hash]*hotRegionPrepared{}, engine: e}
	plans := map[string]map[uint64]*hotRegion{}
	for _, d := range choices {
		hash := common.HexToHash(d.CodeHash)
		if plans[d.CodeHash] == nil {
			code, err := hex.DecodeString(corpus.Codes[d.CodeHash])
			if err != nil || crypto.Keccak256Hash(code) != hash {
				return nil, fmt.Errorf("code identity mismatch")
			}
			_, plans[d.CodeHash] = discoverHotRegions(e, code)
			out.codes[hash] = &hotRegionPrepared{entries: make([]hotRegionExecutable, len(code)), code: code}
		}
		r := plans[d.CodeHash][d.Entry]
		fn := hotRegionGenerated[d.Key]
		// The no-passing builder terminates at helper boundaries. A catalog
		// entry for the original longer graph cannot execute that shorter graph.
		if paperAblationLevel == 2 && (r == nil || r.key != d.Key || r.Helpers > 0) {
			continue
		}
		if r == nil || r.key != d.Key || fn == nil {
			return nil, fmt.Errorf("compiled proof mismatch")
		}
		if !paperStaticChecksAvailable(r, d.Key) {
			continue
		}
		out.codes[hash].entries[d.Entry] = hotRegionExecutable{r, fn}
	}
	return out, nil
}
