package templateminer

import (
	"encoding/json"
	"math/big"
	"os"
	"sort"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/core/vm"
)

func fixture(t *testing.T, name string) *Program {
	t.Helper()
	p, err := Load("testdata/"+name, "Arithmetic.sol", "Arithmetic")
	if err != nil {
		t.Fatal(err)
	}
	return p
}
func scenarios(t *testing.T, p *Program) []NamedTrace {
	t.Helper()
	raw, err := os.ReadFile("testdata/scenarios.json")
	if err != nil {
		t.Fatal(err)
	}
	var calls []struct {
		Name, Method string
		Args         []string
	}
	if err = json.Unmarshal(raw, &calls); err != nil {
		t.Fatal(err)
	}
	var out []NamedTrace
	for _, c := range calls {
		input, err := EncodeCall(p, c.Method, c.Args)
		if err != nil {
			t.Fatal(err)
		}
		trace := Replay(p, input)
		if trace.TrackingError != "" {
			t.Fatal(trace.TrackingError)
		}
		out = append(out, NamedTrace{c.Name, trace})
	}
	return out
}
func mined(t *testing.T, p *Program) *Bundle {
	t.Helper()
	b, err := Mine(p, scenarios(t, p))
	if err != nil {
		t.Fatal(err)
	}
	return b
}
func ids(b *Bundle) []string {
	var r []string
	for _, tpl := range b.Templates {
		r = append(r, tpl.ID)
	}
	sort.Strings(r)
	return r
}

func TestSourceMapInheritanceAndPushInstructionIndex(t *testing.T) {
	loc, err := ParseSourceMap("10:4:0:i:1;:2::o;::1;:::-:0")
	if err != nil {
		t.Fatal(err)
	}
	if len(loc) != 4 || loc[1].Start != 10 || loc[1].Length != 2 || loc[1].Jump != "o" || loc[2].File != 1 || loc[2].ModifierDepth != 1 || loc[3].ModifierDepth != 0 {
		t.Fatalf("wrong inheritance: %+v", loc)
	}
	ins, err := Disassemble([]byte{0x61, 0x5b, 0x01, 0x5b, 0x5f})
	if err != nil {
		t.Fatal(err)
	}
	if len(ins) != 3 || ins[1].PC != 3 || ins[0].Op != vm.PUSH2 {
		t.Fatal("PUSH data mistaken for instructions")
	}
	if _, err = ParseSourceMap("0:1:0:q"); err == nil {
		t.Fatal("invalid map accepted")
	}
}

func TestMiningVersionsOptimizationAndBytecodeOnlyMatching(t *testing.T) {
	var original *Bundle
	for _, name := range []string{"arithmetic-0824.json", "arithmetic-0819.json", "arithmetic-0824-optimized.json"} {
		t.Run(name, func(t *testing.T) {
			p := fixture(t, name)
			b := mined(t, p)
			if len(b.Templates) != 3 || len(b.Instances) != 3 {
				t.Fatalf("expected three arithmetic families, got %d/%d; %v", len(b.Templates), len(b.Instances), b.Diagnostics)
			}
			for _, i := range b.Instances {
				if i.Validation.SuccessWitnesses == 0 || i.Validation.PanicWitnesses == 0 || i.Validation.InputPairs != 228 {
					t.Fatalf("missing validation evidence: %+v", i.Validation)
				}
			}
			unchecked := false
			for _, s := range b.Sites {
				if !s.Checked {
					unchecked = true
					if s.Supported {
						t.Fatal("unchecked site accepted")
					}
				}
			}
			if !unchecked {
				t.Fatal("unchecked label missing")
			}
			report, err := MatchBytecode(p.Code, b.Templates)
			if err != nil {
				t.Fatal(err)
			}
			if len(report.Bindings) != 3 {
				t.Fatalf("self match found %d helpers", len(report.Bindings))
			}
			if original == nil {
				original = b
			}
			if name == "arithmetic-0819.json" {
				if strings.Join(ids(original), ",") != strings.Join(ids(b), ",") {
					t.Fatal("PUSH0/PUSH1 and relocated addresses did not normalize into same families")
				}
				report, err = MatchBytecode(p.Code, original.Templates)
				if err != nil || len(report.Bindings) != 3 {
					t.Fatalf("cross-version bytecode-only matching failed: %v", err)
				}
			}
		})
	}
}

func TestValueIdentityAndNestedSourceSites(t *testing.T) {
	p := fixture(t, "arithmetic-0824.json")
	b := mined(t, p)
	alias, chain := false, 0
	for _, instance := range b.Instances {
		for _, w := range instance.Witnesses {
			if w.Trace == "same-value-alias" {
				alias = true
				if len(w.OperandIDs) != 2 || w.OperandIDs[0] != w.OperandIDs[1] {
					t.Fatal("DUP identity not retained")
				}
			}
			if w.Trace == "chained-sites" {
				chain++
				if len(w.Sites) != 1 {
					t.Fatalf("ambiguous nested source anchor: %v", w.Sites)
				}
			}
			if w.Trace == "unchecked-negative-control" {
				t.Fatal("unchecked execution became a checked witness")
			}
		}
	}
	if !alias || chain != 2 {
		t.Fatalf("missing alias/chain evidence: %v %d", alias, chain)
	}
	input, err := EncodeCall(p, "add", []string{"7", "7"})
	if err != nil {
		t.Fatal(err)
	}
	trace := Replay(p, input)
	w, err := extractWitnesses(p, p.Helpers[0], NamedTrace{"equal-values", trace})
	if err != nil {
		t.Fatal(err)
	}
	if len(w) != 1 || w[0].OperandIDs[0] == w[0].OperandIDs[1] {
		t.Fatal("equal values from distinct calldata slots conflated")
	}
}

func TestActualEVMBoundariesAgainstIndependentSpecification(t *testing.T) {
	p := fixture(t, "arithmetic-0824.json")
	pairs := validationPairs()
	// Cartesian boundary values plus 16 deterministic full-width random pairs.
	for _, h := range p.Helpers {
		for _, pair := range pairs[:116] {
			input, err := EncodeCall(p, strings.ToLower(h.Operation), []string{pair[0].String(), pair[1].String()})
			if err != nil {
				t.Fatal(err)
			}
			trace := Replay(p, input)
			if trace.TrackingError != "" {
				t.Fatal(trace.TrackingError)
			}
			exact := new(big.Int)
			switch h.Operation {
			case "ADD":
				exact.Add(pair[0], pair[1])
			case "SUB":
				exact.Sub(pair[0], pair[1])
			case "MUL":
				exact.Mul(pair[0], pair[1])
			}
			overflow := exact.Sign() < 0 || exact.Cmp(modulus) >= 0
			if overflow {
				if trace.ExecutionError == "" || trace.Return != "4e487b710000000000000000000000000000000000000000000000000000000000000011" {
					t.Fatalf("%s failed to reproduce panic", h.Operation)
				}
			} else {
				if trace.ExecutionError != "" {
					t.Fatal(trace.ExecutionError)
				}
				got, ok := new(big.Int).SetString(trace.Return, 16)
				if !ok || got.Cmp(exact) != 0 {
					t.Fatalf("%s result mismatch", h.Operation)
				}
			}
		}
	}
}

func TestRejectChangedOverflowPredicateAndIncompleteEvidence(t *testing.T) {
	p := fixture(t, "arithmetic-0824.json")
	b := mined(t, p)
	var add Helper
	for _, h := range p.Helpers {
		if h.Operation == "ADD" {
			add = h
		}
	}
	a, err := Analyze(p, add)
	if err != nil {
		t.Fatal(err)
	}
	var changed uint64
	for _, pc := range a.ReachablePCs {
		if p.ByPC[pc].Op == vm.GT {
			changed = pc
			break
		}
	}
	if changed == 0 {
		t.Fatal("test fixture lacks GT predicate")
	}
	p.Code = append([]byte(nil), p.Code...)
	p.Code[changed] = byte(vm.LT)
	ins := p.ByPC[changed]
	ins.Op = vm.LT
	p.ByPC[changed] = ins
	mutant, err := Analyze(p, add)
	if err != nil {
		t.Fatal(err)
	}
	rejected := false
	for _, pair := range validationPairs() {
		if CheckInputs(mutant, "ADD", pair[0], pair[1]) != nil {
			rejected = true
			break
		}
	}
	if !rejected {
		t.Fatal("wrong safety predicate passed")
	}
	report, err := MatchBytecode(p.Code, b.Templates)
	if err != nil {
		t.Fatal(err)
	}
	for _, binding := range report.Bindings {
		if binding.Entry == add.Entry {
			t.Fatal("changed guard matched learned template")
		}
	}
	clean := fixture(t, "arithmetic-0824.json")
	empty, err := Mine(clean, nil)
	if err != nil {
		t.Fatal(err)
	}
	if len(empty.Templates) != 0 {
		t.Fatal("unexecuted metadata exported as runtime evidence")
	}
}

func TestPublicERC20CompoundAssignmentAndCoverageDiagnostics(t *testing.T) {
	p, err := Load("testdata/token-0824.json", "TokenHarness.sol", "TokenHarness")
	if err != nil {
		t.Fatal(err)
	}
	input, err := EncodeCall(p, "mint", []string{"0x0000000000000000000000000000000000000001", "100"})
	if err != nil {
		t.Fatal(err)
	}
	trace := Replay(p, input)
	if trace.ExecutionError != "" || trace.TrackingError != "" {
		t.Fatalf("mint failed: %+v", trace)
	}
	b, err := Mine(p, []NamedTrace{{"mint", trace}})
	if err != nil {
		t.Fatal(err)
	}
	if len(b.Instances) != 1 || b.Instances[0].Helper.Operation != "ADD" {
		t.Fatalf("ERC20 checked += not mined: %v", b.Diagnostics)
	}
	if len(b.Instances[0].Witnesses[0].Sites) != 1 {
		t.Fatal("missing public-source site anchor")
	}
	learned := mined(t, fixture(t, "arithmetic-0824.json"))
	report, err := MatchBytecode(p.Code, learned.Templates)
	if err != nil || len(report.Bindings) != 1 {
		t.Fatalf("cross-contract bytecode matching failed: %v", err)
	}
	ir := fixture(t, "arithmetic-0824-viair.json")
	irBundle := mined(t, ir)
	if len(irBundle.Instances) != 0 || len(irBundle.Diagnostics) == 0 {
		t.Fatal("unsupported viaIR falsely reported covered")
	}
}
