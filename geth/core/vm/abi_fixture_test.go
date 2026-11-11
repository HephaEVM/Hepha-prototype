package vm

import (
	"testing"
)

func chainABI(t testing.TB) *ABIScalarDiagnostic {
	a, err := LoadABIRegionDiagnostic("../../assets/abi-scalar.json", "../../assets/abi-expanded.json")
	if err != nil {
		t.Fatal(err)
	}
	a.Optimize = true
	return a
}
