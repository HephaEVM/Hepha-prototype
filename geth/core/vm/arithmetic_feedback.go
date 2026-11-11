package vm

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"sort"
	"sync"

	"github.com/ethereum/go-ethereum/crypto"
	miner "go-eth1.16.5-evm/template_miner"
)

type arithmeticFeedback struct {
	mu      sync.Mutex
	limit   int
	entries map[string]miner.FeedbackCode
	dropped uint64
}

func newArithmeticFeedback(limit int) *arithmeticFeedback {
	return &arithmeticFeedback{limit: limit, entries: map[string]miner.FeedbackCode{}}
}
func (f *arithmeticFeedback) record(version string, code []byte, hints []miner.CandidateHint) {
	if len(hints) == 0 {
		return
	}
	f.mu.Lock()
	defer f.mu.Unlock()
	key := crypto.Keccak256Hash(code).Hex()
	if old, ok := f.entries[key]; ok {
		old.Candidates = append([]miner.CandidateHint(nil), old.Candidates...)
		seen := map[string]bool{}
		id := func(h miner.CandidateHint) string { return fmt.Sprintf("%d/%s/%v", h.Entry, h.Reason, h.Operands) }
		for _, h := range old.Candidates {
			seen[id(h)] = true
		}
		for _, h := range hints {
			if seen[id(h)] {
				continue
			}
			seen[id(h)] = true
			if len(old.Candidates) < 64 {
				old.Candidates = append(old.Candidates, h)
			} else {
				f.dropped++
				// A concrete semantic counterexample takes priority over a static hint.
				if len(h.Operands) == 2 {
					for i := len(old.Candidates) - 1; i >= 0; i-- {
						if len(old.Candidates[i].Operands) == 0 {
							old.Candidates[i] = h
							break
						}
					}
				}
			}
		}
		f.entries[key] = old
		return
	}
	if len(f.entries) >= f.limit || len(code) > 1<<20 {
		f.dropped++
		return
	}
	if len(hints) > 64 {
		f.dropped += uint64(len(hints) - 64)
		hints = hints[:64]
	}
	f.entries[key] = miner.FeedbackCode{CodeHash: key, Bytecode: hex.EncodeToString(code), LibraryVersion: version, Candidates: append([]miner.CandidateHint(nil), hints...)}
}
func (e *ArithmeticEngine) WriteFeedback(w io.Writer) error {
	f := e.feedback
	f.mu.Lock()
	out := miner.FeedbackBundle{SchemaVersion: 1, Mode: "untrusted_code_candidates_not_execution_counts", Dropped: f.dropped}
	for _, v := range f.entries {
		out.Codes = append(out.Codes, v)
	}
	f.mu.Unlock()
	sort.Slice(out.Codes, func(i, j int) bool { return out.Codes[i].CodeHash < out.Codes[j].CodeHash })
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "  ")
	return encoder.Encode(out)
}
