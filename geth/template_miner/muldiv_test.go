package templateminer

import (
	"encoding/hex"
	"os"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/core/vm"
)

func mulDivFixture(t *testing.T) *Program {
	t.Helper()
	// Actual deployed code hash 61bb013b...0aed, bytes [13017,13193).
	// Only internal PUSH2 jump destinations are relocated to zero-based PCs.
	raw, e := os.ReadFile("testdata/muldiv-deployed.hex")
	if e != nil {
		t.Fatal(e)
	}
	code, e := hex.DecodeString(strings.TrimSpace(string(raw)))
	if e != nil {
		t.Fatal(e)
	}
	ins, e := Disassemble(code)
	if e != nil {
		t.Fatal(e)
	}
	p := &Program{Code: code, Instructions: ins, ByPC: map[uint64]Instruction{}}
	for _, in := range ins {
		p.ByPC[in.PC] = in
	}
	return p
}
func TestMulDivCompleteSuccessProof(t *testing.T) {
	p := mulDivFixture(t)
	r, e := AnalyzeMulDiv(p, 0)
	if e != nil {
		t.Fatal(e)
	}
	if r.A != 1 || r.B != 2 || r.D != 3 || len(r.Short.PCs) != 48 || len(r.Wide.PCs) != 138 {
		t.Fatalf("%+v", r)
	}
	t.Log("short", len(r.Short.PCs), "wide", len(r.Wide.PCs))
	// A prefix-only matcher would accept these altered guards/results.
	for _, pc := range []uint64{47, 56, 57, 165} {
		q := mulDivFixture(t)
		in, ok := q.ByPC[pc]
		if !ok {
			continue
		}
		old := in.Op
		switch old {
		case vm.DIV:
			in.Op = vm.MOD
		case vm.MUL:
			in.Op = vm.ADD
		case vm.GT:
			in.Op = vm.LT
		default:
			continue
		}
		q.ByPC[pc] = in
		q.Code[pc] = byte(in.Op)
		if _, e := AnalyzeMulDiv(q, 0); e == nil {
			t.Fatalf("accepted mutation at %d: %s", pc, old)
		}
	}
}
