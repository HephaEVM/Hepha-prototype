package vm

import (
	"bytes"
	"encoding/hex"
	"os"
	"strings"
	"testing"
)

func preparationPaths(t testing.TB) []*hotRegion {
	t.Helper()
	raw, err := os.ReadFile("testdata/regions/fusion-boundary.hex")
	if err != nil {
		t.Fatal(err)
	}
	code, err := hex.DecodeString(strings.TrimSpace(strings.TrimPrefix(string(raw), "0x")))
	if err != nil {
		t.Fatal(err)
	}
	b, _ := discoverHotRegions(wethAddressEngine(t), code)
	seen := map[string]bool{}
	var rs []*hotRegion
	for mask := uint(0); mask < 4; mask++ {
		r := buildProtectedFlowPath(b, code, &pragueInstructionSet, 54, mask, 2)
		if r != nil && !seen[r.key] {
			seen[r.key] = true
			rs = append(rs, r)
		}
	}
	if len(rs) < 2 {
		t.Fatal("real branch fixture missing")
	}
	return rs
}
func TestPreparedFlowImmutable(t *testing.T) {
	rs := preparationPaths(t)
	for _, r := range rs {
		before := r.fingerprint()
		p := &preparedFlow{original: r}
		plain, pb := p.optimized(false)
		proof, fb := p.optimized(true)
		again, ab := p.optimized(true)
		if proof != again || &fb[0] != &ab[0] {
			t.Fatal("proof not reused")
		}
		old, _, ob, _ := optimizeHotRegion(r)
		next, nb, _ := optimizeFlowRegion(r)
		if plain.fingerprint() != old.fingerprint() || proof.fingerprint() != next.fingerprint() || len(pb) != len(ob) || len(fb) != len(nb) {
			t.Fatal("analysis changed")
		}
		if r.fingerprint() != before {
			t.Fatal("mutated source path")
		}
	}
	a, _, err := emitFlowShared(rs, true)
	if err != nil {
		t.Fatal(err)
	}
	b, _, err := emitPreparedFlowShared(prepareFlows(rs), true)
	if err != nil || !bytes.Equal([]byte(a), []byte(b)) {
		t.Fatal("unstable emission", err)
	}
}
