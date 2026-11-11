package vm

import (
	"os"

	"testing"
)

func wethAddressEngine(t testing.TB) *ArithmeticEngine {
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
