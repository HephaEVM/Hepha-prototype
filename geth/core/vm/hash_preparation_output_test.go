package vm

import (
	"encoding/hex"
	"testing"
)

func TestHashPreparationOutputShapes(t *testing.T) {
	for _, tc := range []struct {
		name, raw string
		kind      uint8
	}{
		{"single", "5f5260205f20", hashOutputIdentity},
		{"preserved", "805f5260205f20", hashOutputIdentity},
		{"constants", "5f526001600260205f20", hashOutputConstants},
		{"prefix_constants", "805f52600760205f20", hashOutputConstants},
		{"moved_slot", "905f5260205f20", hashOutputOne},
		{"masked_slot", "60ff16805f5260205f20", hashOutputOne},
		{"permutation", "91905f5260205f20", hashOutputGeneral},
	} {
		t.Run(tc.name, func(t *testing.T) {
			code, _ := hex.DecodeString(tc.raw)
			rs := hashPlans(t, code)
			if len(rs) != 1 {
				t.Fatal("expected one region")
			}
			p := prepareHashPreparation(rs[0], code, &pragueInstructionSet, HashPreparationDirect)
			if p.outputKind != tc.kind {
				t.Fatalf("kind %d, want %d; outputs %+v", p.outputKind, tc.kind, rs[0].Outputs)
			}
			// Vary input values, preserved prefixes and backing-array capacity;
			// compare stack, memory, gas and exit to original opcode execution.
			hashDifferential(t, code, rs[0], 8106, 100)
		})
	}
}
