

package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"flag"
	"fmt"
	vm "go-eth1.16.5-evm/core/vm"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

func main() {
	dirs := flag.String("fixtures", "", "comma-separated fixture directories")
	out := flag.String("out", "", "new output file")
	library := flag.String("library", "assets/abi-expanded.json", "ABI library")
	extra := flag.String("extra", "", "optional JSON array of additional code hex strings")
	effects := flag.Bool("effects", false, "export stage 3 effect/address chains only")
	full := flag.Bool("full", false, "export stage 1 and 2 native programs")
	arithmetic := flag.String("arithmetic", "assets/arithmetic.json", "arithmetic template library")
	sitesFile := flag.String("sites", "", "optional untimed executed-PC census (extra fixtures remain unrestricted)")
	flag.Parse()
	var engine *vm.ArithmeticEngine
	if *full || *effects {
		f, err := os.Open(*arithmetic)
		must(err)
		engine, err = vm.LoadArithmeticEngine(f, vm.ArithmeticEngineOptions{})
		must(err)
		must(f.Close())
	}
	var exporter *vm.EvmoneProgramExporter
	if *full {
		var err error
		exporter, err = vm.NewEvmoneProgramExporter(*library, engine)
		must(err)
	}
	codes := map[string]bool{}
	sites := map[string][]uint64{}
	if *sitesFile != "" {
		raw, err := os.ReadFile(*sitesFile)
		must(err)
		must(json.Unmarshal(raw, &sites))
	}
	extraCodes := map[string]bool{}
	// Include executed init/runtime code that only appears during a nested
	// creation and therefore is absent from the fixture's initial accounts.
	for code := range sites {
		codes[code] = true
	}
	for _, d := range strings.Split(*dirs, ",") {
		files, err := filepath.Glob(filepath.Join(d, "*.json"))
		must(err)
		for _, p := range files {
			raw, err := os.ReadFile(p)
			must(err)
			var f struct {
				Accounts map[string]struct {
					Code string `json:"code"`
				} `json:"accounts"`
			}
			must(json.Unmarshal(raw, &f))
			for _, a := range f.Accounts {
				if a.Code != "" {
					codes[a.Code] = true
				}
			}
		}
	}
	if *extra != "" {
		raw, err := os.ReadFile(*extra)
		must(err)
		var xs []string
		must(json.Unmarshal(raw, &xs))
		for _, s := range xs {
			codes[s] = true
			extraCodes[s] = true
		}
	}
	codes[vm.EvmoneWordLoopFixture()] = true
	keys := []string{}
	for s := range codes {
		keys = append(keys, s)
	}
	sort.Strings(keys)
	rows := []any{}
	counts := map[int]int{}
	for i, s := range keys {
		code, err := hex.DecodeString(s)
		must(err)
		var regions []vm.EvmoneRegion
		if *full || *effects {
			var selected map[uint64]bool
			if *sitesFile != "" && !extraCodes[s] {
				selected = map[uint64]bool{}
				for _, pc := range sites[s] {
					selected[pc] = true
				}
			}
			if *effects {
				regions, err = vm.ExportEvmoneEffects(code, engine, selected)
			} else {
				regions, err = exporter.Regions(code, selected)
			}
		} else {
			regions, err = vm.ExportEvmoneRegions(code, *library)
		}
		must(err)
		for _, r := range regions {
			counts[r.Kind]++
		}
		rows = append(rows, map[string]any{"code": s, "sha256": fmt.Sprintf("%x", sha256.Sum256(code)), "regions": regions})
		if i%50 == 0 {
			fmt.Printf("analyzed %d/%d\n", i+1, len(keys))
		}
	}
	raw, err := json.MarshalIndent(map[string]any{"version": 1, "fork": "Prague", "codes": rows, "counts": counts}, "", "  ")
	must(err)
	f, err := os.OpenFile(*out, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0644)
	must(err)
	_, err = f.Write(append(raw, '\n'))
	must(err)
	must(f.Close())
	fmt.Println("regions", counts)
}
func must(err error) {
	if err != nil {
		panic(err)
	}
}
