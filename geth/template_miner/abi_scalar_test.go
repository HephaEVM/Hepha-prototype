package templateminer

import (
	"encoding/hex"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestABIScalarCorpus(t *testing.T) {
	paths, err := filepath.Glob("../template_miner/testdata/integer-corpus/matrix-*.json")
	if err != nil {
		t.Fatal(err)
	}
	extra, err := filepath.Glob("../template_miner/testdata/abi-corpus/*.json")
	if err != nil {
		t.Fatal(err)
	}
	paths = append(paths, extra...)
	known, labels, err := ABIScalarTraining(paths)
	if err != nil {
		t.Fatal(err)
	}
	if len(known) == 0 {
		t.Fatal("no scalar templates")
	}
	t.Log("templates", len(known), "compiler labels", len(labels))
	out := os.Getenv("ABI_SCALAR_CENSUS_OUT")
	if out == "" {
		return
	}
	if err := os.MkdirAll(out, 0755); err != nil {
		t.Fatal(err)
	}
	files, err := filepath.Glob("../core/vm/testdata/regions/0x*.hex")
	if err != nil {
		t.Fatal(err)
	}
	bindings := map[string][]*ABIScalarRegion{}
	total := 0
	for _, file := range files {
		b, err := os.ReadFile(file)
		if err != nil {
			t.Fatal(err)
		}
		code, err := hex.DecodeString(strings.TrimSpace(string(b)))
		if err != nil {
			t.Fatal(err)
		}
		regions := DiscoverABIScalars(code, known)
		if len(regions) > 0 {
			bindings[strings.TrimSuffix(filepath.Base(file), ".hex")] = regions
			total += len(regions)
		}
	}
	for name, v := range map[string]any{"training.json": labels, "templates.json": known, "bindings.json": bindings, "summary.json": map[string]any{"codeFiles": len(files), "matchedCodeFiles": len(bindings), "matchedRegions": total, "templates": len(known), "trainingArtifacts": paths}} {
		b, err := json.MarshalIndent(v, "", "  ")
		if err != nil {
			t.Fatal(err)
		}
		if err = os.WriteFile(filepath.Join(out, name), append(b, '\n'), 0644); err != nil {
			t.Fatal(err)
		}
	}
	t.Log("code", len(files), "matched code", len(bindings), "regions", total)
}
