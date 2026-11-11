package templateminer

import (
	"encoding/hex"
	"os"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/crypto"
)

func roundingProgram(t *testing.T) *Program {
	t.Helper()
	raw, err := os.ReadFile("testdata/muldiv-rounding-deployed.hex")
	if err != nil {
		t.Fatal(err)
	}
	code, err := hex.DecodeString(strings.TrimSpace(string(raw)))
	if err != nil {
		t.Fatal(err)
	}
	if crypto.Keccak256Hash(code).Hex() != "0x61bb013bc7f8bed8e3c050464e93985787d254f4236d2e3ef846700c07e00aed" {
		t.Fatal("deployed rounding fixture hash changed")
	}
	ins, err := Disassemble(code)
	if err != nil {
		t.Fatal(err)
	}
	p := &Program{Code: code, Instructions: ins, ByPC: map[uint64]Instruction{}}
	for _, in := range ins {
		p.ByPC[in.PC] = in
	}
	return p
}

func TestMulDivRoundingCompleteProof(t *testing.T) {
	p := roundingProgram(t)
	floor, err := AnalyzeMulDiv(p, 13017)
	if err != nil {
		t.Fatal(err)
	}
	candidates := MulDivRoundingCandidates(p, map[uint64]*MulDivRegion{floor.Entry: floor})
	if candidates[16809] != 13017 {
		t.Fatal("missing observed wrapper", candidates)
	}
	r, err := AnalyzeMulDivRounding(p, 16809, 13017)
	if err != nil {
		t.Fatal(err)
	}
	for i, path := range r.Paths {
		t.Log(i, "instructions", len(path.PCs), "peak", path.Peak)
	}
	for pc, replacement := range map[uint64]vm.OpCode{16837: vm.ADDMOD, 16848: vm.GT, 16860: vm.SUB} {
		q := roundingProgram(t)
		in := q.ByPC[pc]
		t.Log("mutation", pc, in.Op, replacement)
		in.Op = replacement
		q.ByPC[pc], q.Code[pc] = in, byte(replacement)
		if _, err := AnalyzeMulDivRounding(q, 16809, 13017); err == nil {
			t.Fatal("accepted changed rounding semantics", pc)
		}
	}
}
