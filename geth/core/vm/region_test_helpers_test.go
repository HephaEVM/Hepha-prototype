package vm

import (
	"github.com/holiman/uint256"
	"os"
	"testing"
)

func abiRegionTestEngine(t testing.TB) *ABIScalarDiagnostic {
	t.Helper()
	e, err := LoadABIRegionDiagnostic("../../assets/abi-scalar.json", "../../assets/abi-regions.json")
	if err != nil {
		t.Fatal(err)
	}
	e.Optimize = true
	return e
}
func regionArithmeticTestEngine(t testing.TB) *ArithmeticEngine {
	t.Helper()
	f, err := os.Open("../../assets/arithmetic.json")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	e, err := LoadArithmeticEngine(f, ArithmeticEngineOptions{})
	if err != nil {
		t.Fatal(err)
	}
	return e
}
func NewEmptyTestEngine() *ArithmeticEngine {
	e, _ := NewArithmeticEngine(nil, ArithmeticEngineOptions{})
	return e
}
func pipelineArithmeticStep(f *arithmeticFrame, pc *uint64, s *ScopeContext, r *arithmeticRegion) bool {
	ok, err := f.trySyncHot(pc, s, r)
	return ok && err == nil
}
func abiTestEngine(t testing.TB) *ABIScalarDiagnostic {
	t.Helper()
	e, err := LoadABIScalarDiagnostic("../../assets/abi-scalar.json")
	if err != nil {
		t.Fatal(err)
	}
	return e
}

var arithmeticDiagnosticFlag bool
var arithmeticDiagnosticWord uint256.Int

func requireFixtureEnv(t testing.TB, names ...string) {
	t.Helper()
	for _, name := range names {
		if os.Getenv(name) == "" {
			t.Skip("external fixture test: set " + name)
		}
	}
}
