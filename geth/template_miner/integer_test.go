package templateminer

import (
	"encoding/hex"
	"math/big"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/core/vm"
)

func TestSignedExpressionSemanticsAgainstActualEVM(t *testing.T) {
	min := new(big.Int).Neg(new(big.Int).Lsh(big.NewInt(1), 255))
	cases := []struct {
		op   vm.OpCode
		l, r *big.Int
	}{
		{vm.SDIV, big.NewInt(-7), big.NewInt(2)}, {vm.SDIV, big.NewInt(7), big.NewInt(-2)},
		{vm.SDIV, min, big.NewInt(-1)}, {vm.SDIV, big.NewInt(7), big.NewInt(0)},
		{vm.SMOD, big.NewInt(-7), big.NewInt(2)}, {vm.SMOD, big.NewInt(7), big.NewInt(-2)},
		{vm.SMOD, min, big.NewInt(-1)}, {vm.SMOD, min, big.NewInt(0)},
		{vm.SLT, big.NewInt(-1), big.NewInt(0)}, {vm.SGT, min, big.NewInt(1)},
		{vm.SIGNEXTEND, big.NewInt(0), big.NewInt(128)}, {vm.SIGNEXTEND, big.NewInt(0), big.NewInt(127)},
		{vm.SIGNEXTEND, big.NewInt(1), big.NewInt(32768)}, {vm.SIGNEXTEND, big.NewInt(31), min},
		{vm.SIGNEXTEND, big.NewInt(32), big.NewInt(128)}, {vm.SIGNEXTEND, new(big.Int).Set(mask256), big.NewInt(128)},
		{vm.SAR, big.NewInt(1), big.NewInt(-7)}, {vm.SAR, big.NewInt(256), big.NewInt(-7)},
		{vm.SAR, new(big.Int).Set(mask256), big.NewInt(7)},
	}
	for _, c := range cases {
		var code []byte
		for _, n := range []*big.Int{c.r, c.l} {
			code = append(code, byte(vm.PUSH32))
			code = append(code, new(big.Int).Mod(n, modulus).FillBytes(make([]byte, 32))...)
		}
		code = append(code, byte(c.op), byte(vm.PUSH0), byte(vm.MSTORE), byte(vm.PUSH1), 32, byte(vm.PUSH0), byte(vm.RETURN))
		instructions, _ := Disassemble(code)
		p := &Program{Code: code, ByPC: map[uint64]Instruction{}}
		for _, i := range instructions {
			p.ByPC[i.PC] = i
		}
		tr := Replay(p, nil)
		got, err := (&Expr{Op: c.op.String(), Args: []*Expr{constant(c.l, nil), constant(c.r, nil)}}).Evaluate(nil)
		if err != nil || tr.TrackingError != "" || tr.ExecutionError != "" || hex.EncodeToString(got.FillBytes(make([]byte, 32))) != tr.Return {
			t.Fatalf("%s(%s,%s) != EVM: %v %s", c.op, c.l, c.r, err, tr.Return)
		}
	}
}

func TestTypedDiscoveryUsesBytecodeNotSourceLabels(t *testing.T) {
	for _, contract := range []string{"UnsignedMatrix", "SignedMatrix"} {
		p, err := Load("testdata/integer-0824.json", "IntegerMatrix.sol", contract)
		if err != nil {
			t.Fatal(err)
		}
		cache := map[uint64]*candidate{}
		for _, width := range []int{8, 64, 128, 256} {
			typ := integerType{width, contract == "SignedMatrix"}
			for _, op := range []string{"ADD", "SUB", "MUL", "DIV", "MOD"} {
				evmOp := op
				if typ.signed && (op == "DIV" || op == "MOD") {
					evmOp = "S" + op
				}
				name := strings.ToLower(op) + "_" + typ.name()
				for _, pair := range actualStudyPairs(typ) {
					input, err := EncodeCall(p, name, []string{pair[0].String(), pair[1].String()})
					if err != nil {
						t.Fatal(err)
					}
					tr := Replay(p, input)
					if tr.TrackingError != "" {
						t.Fatal(tr.TrackingError)
					}
					found, rejected := mineTraceRegions(p, tr, cache)
					if len(found) != 1 {
						t.Fatalf("%s(%s,%s): %d regions, %v", name, pair[0], pair[1], len(found), rejected)
					}
					c := found[0].c
					if c.analysis.IntegerType != typ.name() || c.helper.Operation != evmOp {
						t.Fatalf("wrong inferred domain: %s %s", c.analysis.IntegerType, c.helper.Operation)
					}
					_, want, err := arithmeticSpec(evmOp, typ, pair[0], pair[1])
					if err != nil || found[0].w.Outcome != want {
						t.Fatalf("wrong panic class for %s", name)
					}
				}
			}
		}
		for _, width := range []int{8, 128, 256} {
			typ := integerType{width, contract == "SignedMatrix"}
			_, hi := typ.bounds()
			input, err := EncodeCall(p, "uncheckedAdd_"+typ.name(), []string{hi.String(), "1"})
			if err != nil {
				t.Fatal(err)
			}
			found, _ := mineTraceRegions(p, Replay(p, input), cache)
			if len(found) != 0 {
				t.Fatalf("unchecked %s accepted", typ.name())
			}
		}
	}
}

func TestTypedPanicMutationAndCanonicalInputPrecondition(t *testing.T) {
	p, err := Load("testdata/integer-0824.json", "IntegerMatrix.sol", "SignedMatrix")
	if err != nil {
		t.Fatal(err)
	}
	input, _ := EncodeCall(p, "div_int8", []string{"-128", "-1"})
	tr := Replay(p, input)
	found, _ := mineTraceRegions(p, tr, map[uint64]*candidate{})
	if len(found) != 1 {
		t.Fatal("missing signed division region")
	}
	c := found[0].c
	if err := CheckInputs(c.analysis, "SDIV", big.NewInt(128), big.NewInt(1)); err == nil {
		t.Fatal("non-sign-extended int8 input accepted")
	}
	if err := CheckInputs(c.analysis, "SDIV", big.NewInt(-128), big.NewInt(-1)); err != nil {
		t.Fatal(err)
	}
	f, end := frameFromTrace(tr)
	end.Output = strings.TrimSuffix(end.Output, "11") + "12"
	if _, err := chainWitness(p, f, end, found[0].w.EntryStep, c); err == nil {
		t.Fatal("wrong concrete panic code accepted")
	}
	// Corrupt only the overflow panic constant, leaving zero-divisor panic intact.
	changed := false
	for _, pc := range c.analysis.ReachablePCs {
		i := p.ByPC[pc]
		if i.Op.IsPush() && new(big.Int).SetBytes(i.Data).Cmp(big.NewInt(0x11)) == 0 {
			i.Data = append([]byte(nil), i.Data...)
			i.Data[len(i.Data)-1] = 0x12
			p.ByPC[pc] = i
			changed = true
		}
	}
	if !changed {
		t.Fatal("fixture lacks panic constant")
	}
	if mutated := discoverCandidate(p, c.helper.Entry); mutated.err == "" {
		t.Fatal("wrong symbolic panic code accepted")
	}
}

func TestTypedTemplateMatchingRetainsWidth(t *testing.T) {
	p, err := Load("testdata/integer-0824.json", "IntegerMatrix.sol", "UnsignedMatrix")
	if err != nil {
		t.Fatal(err)
	}
	input, _ := EncodeCall(p, "add_uint8", []string{"7", "2"})
	found, _ := mineTraceRegions(p, Replay(p, input), map[uint64]*candidate{})
	if len(found) != 1 {
		t.Fatal("no uint8 region")
	}
	report, err := MatchBytecode(p.Code, []Template{found[0].c.template})
	if err != nil || len(report.Bindings) != 1 || report.Bindings[0].IntegerType != "uint8" {
		t.Fatalf("typed match lost width: %v %+v", err, report)
	}
}

func TestMainnetModeAcceptsTypedAndDivideByZeroFrames(t *testing.T) {
	u, err := Load("testdata/integer-0824.json", "IntegerMatrix.sol", "UnsignedMatrix")
	if err != nil {
		t.Fatal(err)
	}
	s, err := Load("testdata/integer-0824.json", "IntegerMatrix.sol", "SignedMatrix")
	if err != nil {
		t.Fatal(err)
	}
	dir, records := auditFixture(t, []*Program{u, s}, []string{"div_uint8", "sub_int128"})
	r, err := MineAudit(dir, "")
	if err != nil || len(r.Contracts) != 2 || len(r.Templates) != 2 {
		t.Fatalf("typed mainnet mining: %v", err)
	}
	types := map[string]bool{}
	for _, c := range r.Contracts {
		for _, i := range c.Instances {
			types[i.IntegerType] = true
		}
	}
	if !types["uint8"] || !types["int128"] {
		t.Fatalf("typed domains lost: %v", types)
	}
	// Replace the first frame with a genuine zero-divisor EVM execution.
	input, _ := EncodeCall(u, "div_uint8", []string{"7", "0"})
	tr := Replay(u, input)
	f, end := frameFromTrace(tr)
	var changed []AuditRecord
	changed = append(changed, records[0])
	for _, op := range f.ops {
		op.Kind = "opcode"
		op.CallID = 1
		changed = append(changed, op)
	}
	end.Kind = "exit"
	end.CallID = 1
	changed = append(changed, end)
	for _, rec := range records {
		if rec.CallID == 2 {
			changed = append(changed, rec)
		}
	}
	writeAuditFixture(t, dir, changed)
	r, err = MineAudit(dir, "")
	if err != nil {
		t.Fatal(err)
	}
	seen := false
	for _, c := range r.Contracts {
		for _, i := range c.Instances {
			if i.Operation == "DIV" {
				seen = i.Panics == 1 && i.Witnesses[0].Outcome == "panic_0x12"
			}
		}
	}
	if !seen {
		t.Fatal("zero-divisor panic was not retained in mainnet mode")
	}
}
