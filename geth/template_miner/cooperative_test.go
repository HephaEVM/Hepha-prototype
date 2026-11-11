package templateminer

import (
	"encoding/hex"
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
	"time"
)

func cooperativeFixtures(t *testing.T) ([][]byte, *BytecodeMatcher, *ABIRegionLibrary, map[string]bool) {
	t.Helper()
	var bundle struct{ Templates []Template }
	raw, err := os.ReadFile("../assets/arithmetic.json")
	if err != nil {
		t.Fatal(err)
	}
	if err = json.Unmarshal(raw, &bundle); err != nil {
		t.Fatal(err)
	}
	matcher, err := NewBytecodeMatcher(bundle.Templates)
	if err != nil {
		t.Fatal(err)
	}
	raw, err = os.ReadFile("../assets/abi-expanded.json")
	if err != nil {
		t.Fatal(err)
	}
	var lib ABIRegionLibrary
	if err = json.Unmarshal(raw, &lib); err != nil {
		t.Fatal(err)
	}
	raw, err = os.ReadFile("../assets/abi-scalar.json")
	if err != nil {
		t.Fatal(err)
	}
	var known map[string]bool
	if err = json.Unmarshal(raw, &known); err != nil {
		t.Fatal(err)
	}
	paths, err := filepath.Glob("../core/vm/testdata/regions/*.hex")
	if err != nil || len(paths) == 0 {
		t.Fatal("missing deployed bytecode fixtures", err)
	}
	codes := [][]byte{{0x60}, {0x60, 0x5b, 0x00}}
	for _, p := range paths {
		raw, err := os.ReadFile(p)
		if err != nil {
			t.Fatal(err)
		}
		code, err := hex.DecodeString(strings.TrimPrefix(strings.TrimSpace(string(raw)), "0x"))
		if err != nil {
			t.Fatal(err)
		}
		codes = append(codes, code)
	}
	return codes, matcher, &lib, known
}

func TestCooperativeDiscoverySameBindings(t *testing.T) {
	codes, m, lib, known := cooperativeFixtures(t)
	calls := 0
	pause := func() { calls++ }
	for i, code := range codes {
		a, ea := m.Match(code)
		b, eb := m.MatchWithYield(code, pause)
		if (ea == nil) != (eb == nil) || !reflect.DeepEqual(a, b) {
			t.Fatal("arithmetic result changed", i, ea, eb)
		}
		a1, b1 := DiscoverABIScalars(code, known), DiscoverABIScalars(code, known, pause)
		if !reflect.DeepEqual(a1, b1) {
			t.Fatal("scalar result changed", i)
		}
		a2, b2 := DiscoverABIRegions(code, lib), DiscoverABIRegions(code, lib, pause)
		if !reflect.DeepEqual(a2, b2) {
			t.Fatal("ABI region result changed", i)
		}
	}
	if calls < 2 {
		t.Fatal("no candidate checkpoints", calls)
	}
	t.Log("identical binding, analysis, and candidate records", len(codes), "codes", calls, "checkpoints")
}

func TestCooperativeDiscoveryCanPauseBetweenCandidates(t *testing.T) {
	_, m, _, _ := cooperativeFixtures(t)
	entered, resume, done := make(chan struct{}), make(chan struct{}), make(chan struct{})
	calls := 0
	go func() {
		defer close(done)
		_, _ = m.MatchWithYield([]byte{0x5b, 0, 0x5b, 0}, func() {
			calls++
			if calls == 2 {
				close(entered)
				<-resume
			}
		})
	}()
	select {
	case <-entered:
	case <-time.After(time.Second):
		t.Fatal("did not reach second candidate")
	}
	select {
	case <-done:
		t.Fatal("analysis bypassed pause")
	default:
	}
	close(resume)
	select {
	case <-done:
	case <-time.After(time.Second):
		t.Fatal("analysis did not resume")
	}
}
