package templateminer

import (
	"encoding/hex"
	"encoding/json"
	"math/big"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

func TestLegacyPublishedLibrariesAndNegativeControls(t *testing.T) {
	base := "../template_miner/testdata/legacy-corpus"
	data, err := os.ReadFile(filepath.Join(base, "manifest.json"))
	if err != nil {
		t.Fatal(err)
	}
	var entries []struct{ Artifact, Source, Contract string }
	if err = json.Unmarshal(data, &entries); err != nil {
		t.Fatal(err)
	}
	covered, gaps, negative := 0, 0, 0
	for _, e := range entries {
		p, err := Load(filepath.Join(base, e.Artifact), e.Source, e.Contract)
		if err != nil {
			t.Fatal(err)
		}
		methods, err := abi.JSON(strings.NewReader(string(p.Contract.ABI)))
		if err != nil {
			t.Fatal(err)
		}
		cache := map[uint64]*candidate{}
		for name := range methods.Methods {
			isNegative := strings.HasPrefix(name, "unchecked") || strings.HasPrefix(name, "misleading")
			gap := false
			pairs := [][2]*big.Int{{big.NewInt(7), big.NewInt(2)}, {big.NewInt(0), big.NewInt(0)}, {new(big.Int).Set(mask256), big.NewInt(2)}, {big.NewInt(1), big.NewInt(2)}, {big.NewInt(11), big.NewInt(0)}}
			seen := false
			for _, pair := range pairs {
				input, err := EncodeCall(p, name, []string{pair[0].String(), pair[1].String()})
				if err != nil {
					t.Fatal(err)
				}
				tr := Replay(p, input)
				if tr.TrackingError != "" {
					t.Fatal(tr.TrackingError)
				}
				regions, rejected := mineTraceRegionsModel(p, tr, cache, true)
				if isNegative || gap {
					if len(regions) > 0 {
						t.Fatalf("unsupported/negative %s %s accepted", e.Artifact, name)
					}
					continue
				}
				if len(regions) == 0 {
					t.Fatalf("missing %s %s (%s,%s): %v", e.Artifact, name, pair[0], pair[1], rejected)
				}
				for _, r := range regions {
					if r.w.Outcome != legacyActualOutcome(tr) {
						t.Fatal("lost legacy failure category")
					}
					if r.c.template.Semantics != legacySemantics || r.c.template.FailureHandling != "execute_original_failure_path" {
						t.Fatal("missing failure execution requirement")
					}
					if _, err := Analyze(p, r.c.helper); err == nil {
						t.Fatal("legacy helper accepted as a Solidity 0.8 panic helper")
					}
				}
				seen = true
			}
			if isNegative {
				negative++
			} else if gap {
				gaps++
			} else if seen {
				covered++
			}
		}
	}
	if covered != 54 || gaps != 0 || negative != 24 {
		t.Fatalf("unexpected matrix coverage %d/%d negative=%d", covered, gaps, negative)
	}
}

func TestLegacyCodeDataRelocationAndFailurePadding(t *testing.T) {
	p, err := Load("../template_miner/testdata/legacy-corpus/oz-3.4.2-solc-0612-optimized.json", "LegacyHarness.sol", "LegacyHarness")
	if err != nil {
		t.Fatal(err)
	}
	input, _ := EncodeCall(p, "mul_uint256", []string{"7", "2"})
	found, rejected := mineTraceRegionsModel(p, Replay(p, input), map[uint64]*candidate{}, true)
	if len(found) != 1 {
		t.Fatalf("missing MUL: %v", rejected)
	}
	c := found[0].c
	if len(c.analysis.CodeDataProducers) != 1 {
		t.Fatal("expected one code-backed error string")
	}
	var producer uint64
	var data string
	for producer, data = range c.analysis.CodeDataProducers {
	}
	var masked *SymbolicPath
	for i := range c.analysis.Paths {
		if c.analysis.Paths[i].FailureMask != "" {
			masked = &c.analysis.Paths[i]
		}
	}
	if masked == nil || masked.FailureMemory == nil {
		t.Fatal("caller padding dependency was erased")
	}
	bytes, _ := hex.DecodeString(masked.FailureData)
	mask, _ := hex.DecodeString(masked.FailureMask)
	for i, m := range mask {
		if m == 0 {
			bytes[i] = 0xa5
		}
	}
	if !legacyFailureMatches(masked, hex.EncodeToString(bytes)) {
		t.Fatal("padding was incorrectly fixed to zero")
	}
	bytes[0] ^= 1
	if legacyFailureMatches(masked, hex.EncodeToString(bytes)) {
		t.Fatal("selector mutation accepted")
	}
	// Relocate just the error string. Its code offset is not a semantic
	// constant; its bytes and failure category are.
	ins := p.ByPC[producer]
	blob, _ := hex.DecodeString(data)
	newOffset := len(p.Code)
	code := append(append([]byte(nil), p.Code...), blob...)
	new(big.Int).SetInt64(int64(newOffset)).FillBytes(code[int(producer)+1 : int(producer)+1+len(ins.Data)])
	makeProgram := func(code []byte) *Program {
		instructions, err := Disassemble(code)
		if err != nil {
			t.Fatal(err)
		}
		q := &Program{Code: code, Instructions: instructions, ByPC: map[uint64]Instruction{}}
		for _, i := range instructions {
			q.ByPC[i.PC] = i
		}
		return q
	}
	relocated := discoverLegacyCandidate(makeProgram(code), c.helper.Entry)
	if relocated.err != "" || relocated.template.ID != c.template.ID {
		t.Fatalf("code-data relocation changed family: %s", relocated.err)
	}
	code[newOffset] ^= 1
	changed := discoverLegacyCandidate(makeProgram(code), c.helper.Entry)
	if changed.err != "" || changed.template.ID == c.template.ID {
		t.Fatal("different revert string merged into the same template")
	}
	matched, err := MatchBytecode(p.Code, []Template{c.template})
	if err != nil {
		t.Fatal(err)
	}
	if len(matched.Bindings) == 0 || matched.Bindings[0].Semantics != legacySemantics {
		t.Fatal("legacy offline matching lost semantics")
	}
}
