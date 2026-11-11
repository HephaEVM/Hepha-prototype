package vm

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/holiman/uint256"
	miner "go-eth1.16.5-evm/template_miner"
)

func completeArithmeticEngine(t *testing.T) *ArithmeticEngine {
	t.Helper()
	data, err := os.ReadFile("../../assets/arithmetic.json")
	if err != nil {
		t.Fatal(err)
	}
	e, err := LoadArithmeticEngine(bytes.NewReader(data), ArithmeticEngineOptions{CacheEntries: 4096})
	if err != nil {
		t.Fatal(err)
	}
	return e
}

// Success, boundary gas, INVALID, fixed strings and caller-provided error strings
// must preserve return bytes, remaining gas, state, refunds and logs.
func TestArithmeticSyncLegacyCorpus(t *testing.T) {
	for _, reduced := range []bool{false, true} {
		t.Run(fmt.Sprintf("reduced_%t", reduced), func(t *testing.T) { testArithmeticSyncLegacyCorpus(t, reduced) })
	}
}

func testArithmeticSyncLegacyCorpus(t *testing.T, reduced bool, fixed ...bool) {
	useFixed := len(fixed) != 0 && fixed[0]
	e := completeArithmeticEngine(t)
	base := "../../template_miner/testdata/legacy-corpus"
	raw, err := os.ReadFile(filepath.Join(base, "extended-manifest.json"))
	if err != nil {
		t.Fatal(err)
	}
	var entries []struct{ Artifact, Source, Contract string }
	if err = json.Unmarshal(raw, &entries); err != nil {
		t.Fatal(err)
	}
	for _, opt := range []string{"optimized", "unoptimized"} {
		entries = append(entries, struct{ Artifact, Source, Contract string }{"argument-oz-2.5.1-solc-0517-" + opt + ".json", "LegacyHarness.sol", "LegacyHarness"})
	}
	for _, entry := range entries {
		t.Run(entry.Artifact, func(t *testing.T) {
			p, err := miner.Load(filepath.Join(base, entry.Artifact), entry.Source, entry.Contract)
			if err != nil {
				t.Fatal(err)
			}
			methods, err := abi.JSON(strings.NewReader(string(p.Contract.ABI)))
			if err != nil {
				t.Fatal(err)
			}
			for name, m := range methods.Methods {
				negative := strings.HasPrefix(name, "unchecked") || strings.HasPrefix(name, "misleading")
				if len(m.Inputs) < 2 {
					continue
				}
				signed := m.Inputs[0].Type.T == abi.IntTy
				bits := uint(256)
				if signed {
					bits = 255
				}
				hi := new(big.Int).Sub(new(big.Int).Lsh(big.NewInt(1), bits), big.NewInt(1)).String()
				pairs := [][]string{{"7", "2"}, {"0", "0"}, {"0", "1"}, {hi, "1"}, {hi, "2"}, {"1", "0"}}
				if signed {
					pairs = append(pairs, []string{new(big.Int).Neg(new(big.Int).Lsh(big.NewInt(1), 255)).String(), "-1"}, []string{"-7", "2"})
				}
				for _, pair := range pairs {
					lengths := []int{-1}
					if len(m.Inputs) == 3 {
						lengths = []int{0, 1, 31, 32, 33, 257}
					}
					for _, n := range lengths {
						args := append([]string(nil), pair...)
						if n >= 0 {
							args = append(args, strings.Repeat("x", n))
						}
						input, err := miner.EncodeCall(p, name, args)
						if err != nil {
							t.Fatal(err)
						}
						want := runArithmeticTest(t, p.Code, input, 100000, nil)
						before := e.Stats()
						got := runArithmeticTest(t, p.Code, input, 100000, &ArithmeticConfig{Engine: e, Mode: ArithmeticSync})
						assertArithmeticSame(t, want, got)
						assertArithmeticSame(t, want, runArithmeticTest(t, p.Code, input, 100000, &ArithmeticConfig{Engine: e, Mode: ArithmeticSync}))
						if negative && e.Stats().Pruned != before.Pruned {
							t.Fatalf("negative control pruned: %s", name)
						}
						if !negative && want.err == "" && e.Stats().Pruned == before.Pruned {
							t.Fatalf("success not covered: %s %v, stats=%+v", name, args, e.Stats())
						}
						// Include OOG before the guard and immediately around successful return.
						gases := []uint64{50, 200, 400}
						if want.err == "" {
							used := 100000 - want.gas
							gases = append(gases, used-1, used, used+1)
						}
						for _, gas := range gases {
							assertArithmeticSame(t, runArithmeticTest(t, p.Code, input, gas, nil), runArithmeticTest(t, p.Code, input, gas, &ArithmeticConfig{Engine: e, Mode: ArithmeticSync}))
						}
					}
				}
			}
		})
	}
	if useFixed && e.Stats().FixedStackPrunes == 0 {
		t.Fatal("legacy fixed path was not exercised")
	}
}

func TestArithmeticSyncMissingHelpersAndLazyContext(t *testing.T) {
	e := completeArithmeticEngine(t)
	raw, err := os.ReadFile("../../core/vm/testdata/regions/0xe3486cf05f2c00e00767dec63eb46d309ac8b8e4d3b870a9c13c3de4605009f9.hex")
	if err != nil {
		t.Fatal(err)
	}
	code, err := hex.DecodeString(strings.TrimPrefix(strings.TrimSpace(string(raw)), "0x"))
	if err != nil {
		t.Fatal(err)
	}
	bound := e.bind(code, arithmeticCacheKey{hash: crypto.Keccak256Hash(code), rules: &cancunInstructionSet})
	for _, entry := range []uint64{17248, 17271} {
		if bound.regions[entry] == nil {
			t.Fatalf("known missing helper still absent: %d", entry)
		}
	}
	// Contract contains a valid helper but execution stops before reaching it.
	fixture, _, _ := arithmeticAddFixture(t)
	fixture[0] = byte(STOP)
	before := e.Stats()
	runArithmeticTest(t, fixture, nil, 100000, &ArithmeticConfig{Engine: e, Mode: ArithmeticSync})
	if e.Stats().MatchedRegions == before.MatchedRegions {
		t.Fatal("lazy test had no positive binding")
	}
	if e.Stats().FrameContexts != before.FrameContexts {
		t.Fatal("unused positive binding allocated a context")
	}
	negative := []byte{byte(STOP)}
	before = e.Stats()
	runArithmeticTest(t, negative, nil, 100000, &ArithmeticConfig{Engine: e, Mode: ArithmeticSync})
	if e.Stats().FrameContexts != before.FrameContexts {
		t.Fatal("negative binding allocated a context")
	}
}

func TestArithmeticSharedExpressionPlan(t *testing.T) {
	a, b := &miner.Expr{Op: "input", Input: "a"}, &miner.Expr{Op: "input", Input: "b"}
	add := &miner.Expr{Op: "ADD", Args: []*miner.Expr{a, b}}
	guard := &miner.Expr{Op: "LT", Args: []*miner.Expr{add, a}}
	paths := []arithmeticPath{{path: miner.SymbolicPath{Outcome: "return", Result: add, Conditions: []miner.Condition{{Expression: guard, Nonzero: false}}}}, {path: miner.SymbolicPath{Outcome: "panic_0x11", Conditions: []miner.Condition{{Expression: guard, Nonzero: true}}}}}
	p, err := buildArithmeticPlan(paths)
	if err != nil {
		t.Fatal(err)
	}
	if len(p.nodes) != 4 {
		t.Fatalf("expected a,b,ADD,LT only; got %d", len(p.nodes))
	}
	s := new(arithmeticScratch)
	if p.selectPath(*uint256.NewInt(5), *uint256.NewInt(7), s) != 0 {
		t.Fatal("wrong success path")
	}
	valid := s.valid
	result := p.value(p.results[0], s)
	if result.Uint64() != 12 || s.valid != valid {
		t.Fatal("result recomputed instead of shared")
	}
	if p.selectPath(*new(uint256.Int).SetAllOne(), *uint256.NewInt(1), s) != 1 {
		t.Fatal("scratch from previous invocation reused")
	}
	// Prefix consistency and completeness are checked before publication.
	if _, err := buildArithmeticPlan(paths[:1]); err == nil {
		t.Fatal("incomplete tree accepted")
	}
}

// The fast opcode prefix is a necessary condition only. It must not discard
// any binding accepted by the full symbolic matcher, including legacy layouts.
func TestArithmeticIndexedMatcherCompleteness(t *testing.T) {
	e := completeArithmeticEngine(t)
	base := "../../template_miner/testdata/legacy-corpus"
	raw, err := os.ReadFile(filepath.Join(base, "extended-manifest.json"))
	if err != nil {
		t.Fatal(err)
	}
	var entries []struct{ Artifact, Source, Contract string }
	if err = json.Unmarshal(raw, &entries); err != nil {
		t.Fatal(err)
	}
	codes := map[string][]byte{}
	for _, entry := range entries {
		p, err := miner.Load(filepath.Join(base, entry.Artifact), entry.Source, entry.Contract)
		if err != nil {
			t.Fatal(err)
		}
		codes[entry.Artifact] = p.Code
	}
	for _, opt := range []string{"unoptimized", "optimized"} {
		p, err := miner.Load(filepath.Join(base, "argument-oz-2.5.1-solc-0517-"+opt+".json"), "LegacyHarness.sol", "LegacyHarness")
		if err != nil {
			t.Fatal(err)
		}
		codes["argument-"+opt] = p.Code
	}
	for _, name := range []string{"arithmetic-0819.json", "arithmetic-0824.json", "arithmetic-0824-optimized.json", "arithmetic-0824-viair.json"} {
		p, err := miner.Load("../../template_miner/testdata/"+name, "Arithmetic.sol", "Arithmetic")
		if err != nil {
			t.Fatal(err)
		}
		codes[name] = p.Code
	}
	raw, err = os.ReadFile("../../core/vm/testdata/regions/0xe3486cf05f2c00e00767dec63eb46d309ac8b8e4d3b870a9c13c3de4605009f9.hex")
	if err != nil {
		t.Fatal(err)
	}
	codes["deployed-missing-helpers"], err = hex.DecodeString(strings.TrimPrefix(strings.TrimSpace(string(raw)), "0x"))
	if err != nil {
		t.Fatal(err)
	}
	for name, code := range codes {
		t.Run(name, func(t *testing.T) {
			want, err := miner.MatchBytecode(code, e.templates)
			if err != nil {
				t.Fatal(err)
			}
			got, err := e.matcher.Match(code)
			if err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(want.Bindings, got.Bindings) {
				t.Fatal("prefix filter changed accepted bindings")
			}
			for _, b := range got.Bindings {
				if _, err = prepareArithmeticRegion(b, code, &cancunInstructionSet); err != nil {
					t.Fatalf("known binding %s at %d not executable: %v", b.TemplateID, b.Entry, err)
				}
			}
		})
	}
}
