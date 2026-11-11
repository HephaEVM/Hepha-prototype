package vm

import (
	"encoding/hex"
	"os"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/crypto"
)

// Two distinct signed MUL plans in the recorded 100-block mainnet census.
// The original deployed code is hash-checked. Only the testing dispatcher is
// constructed; these repeated loops are not historical transactions.
var deployedArithmeticHelpers = []struct {
	hash, template string
	helper         uint64
}{
	{"0x52693c83cb9f3adad533fb1a6a4e2d222472d6c958f211d3b4d9b229d304a9db", "arith-7e4d717d1f4b7543a06b64dc", 6605},
	{"0xfa8807899571ad619b60cd5d589160602b93b98dc3da870589bea4136582dcbe", "arith-f062aa0b08811c3a0bed5623", 20298},
}

func deployedArithmeticHelper(t testing.TB, index int) ([]byte, *ArithmeticEngine, *arithmeticRegion) {
	t.Helper()
	x := deployedArithmeticHelpers[index]
	hexCode, err := os.ReadFile("testdata/regions/" + x.hash + ".hex")
	if err != nil {
		t.Fatal(err)
	}
	code, err := hex.DecodeString(strings.TrimPrefix(strings.TrimSpace(string(hexCode)), "0x"))
	if err != nil || crypto.Keccak256Hash(code).Hex() != x.hash {
		t.Fatal("deployed bytecode identity mismatch", err)
	}
	file, err := os.Open("../../assets/arithmetic.json")
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()
	e, err := LoadArithmeticEngine(file, ArithmeticEngineOptions{})
	if err != nil {
		t.Fatal(err)
	}
	bound := e.bind(code, arithmeticCacheKey{hash: crypto.Keccak256Hash(code), rules: &cancunInstructionSet, reduced: true})
	r := bound.regions[x.helper]
	if r == nil || r.template != x.template || r.plan.nativeReduced == nil || r.resultFast == nil || r.domain.Bits != 256 || !r.domain.Signed {
		t.Fatal("expected exact native signed multiplication helper")
	}
	return code, e, r
}

func deployedArithmeticConfig(e *ArithmeticEngine, _ string) *ArithmeticConfig {
	return &ArithmeticConfig{Engine: e, Mode: ArithmeticSync}
}
