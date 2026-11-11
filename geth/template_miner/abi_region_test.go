package templateminer

import (
	"encoding/hex"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestABIRegionCorpus(t *testing.T) {
	paths, _ := filepath.Glob("../template_miner/testdata/abi-region-corpus/*.json")
	scalar, _ := filepath.Glob("../template_miner/testdata/abi-corpus/*.json")
	paths = append(paths, scalar...)
	lib, labels, err := ABIRegionTraining(paths)
	if err != nil {
		t.Fatal(err)
	}
	if len(lib.Templates) == 0 {
		t.Fatal("no templates")
	}
	categories := map[string]int{}
	rejects := map[string]int{}
	for _, x := range labels {
		if why, ok := x["rejected"].(string); ok {
			rejects[why]++
		} else {
			categories[x["category"].(string)]++
		}
	}
	t.Log("templates", len(lib.Templates), "accepted labels", categories, "rejected", rejects)
	out := os.Getenv("ABI_REGION_OUT")
	if out == "" {
		return
	}
	if err = os.MkdirAll(out, 0755); err != nil {
		t.Fatal(err)
	}
	write := func(name string, v any) {
		b, e := json.MarshalIndent(v, "", "  ")
		if e != nil {
			t.Fatal(e)
		}
		if e = os.WriteFile(filepath.Join(out, name), append(b, '\n'), 0644); e != nil {
			t.Fatal(e)
		}
	}
	write("templates.json", lib)
	write("training.json", labels)
	if os.Getenv("ABI_REGION_CENSUS") == "" {
		return
	}
	files, _ := filepath.Glob("../core/vm/testdata/regions/0x*.hex")
	bindings := map[string][]*ABIRegion{}
	total := 0
	counts := map[string]int{}
	for i, file := range files {
		b, e := os.ReadFile(file)
		if e != nil {
			t.Fatal(e)
		}
		code, e := hex.DecodeString(strings.TrimSpace(string(b)))
		if e != nil {
			t.Fatal(e)
		}
		rs := DiscoverABIRegions(code, lib)
		if len(rs) > 0 {
			bindings[strings.TrimSuffix(filepath.Base(file), ".hex")] = rs
			total += len(rs)
			for _, r := range rs {
				counts[r.Category]++
			}
		}
		if i%250 == 0 {
			t.Log("scanned", i, "matched", total)
		}
	}
	write("bindings.json", bindings)
	write("summary.json", map[string]any{"codes": len(files), "matchedCodes": len(bindings), "regions": total, "categories": counts, "templates": len(lib.Templates), "trainingRejected": rejects})
	t.Log("census", len(files), len(bindings), total, counts)
}
