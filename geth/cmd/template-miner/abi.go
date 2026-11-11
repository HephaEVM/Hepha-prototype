package main

import (
	"encoding/hex"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	miner "go-eth1.16.5-evm/template_miner"
)

func runABIStudy(patterns, codeDir, output string) error {
	seen := map[string]bool{}
	var paths []string
	for _, pattern := range strings.Split(patterns, ",") {
		matches, err := filepath.Glob(strings.TrimSpace(pattern))
		if err != nil {
			return err
		}
		if len(matches) == 0 {
			return fmt.Errorf("no artifacts match %q", pattern)
		}
		for _, p := range matches {
			if !seen[p] {
				seen[p] = true
				paths = append(paths, p)
			}
		}
	}
	sort.Strings(paths)
	library, labels, err := miner.ABIRegionTraining(paths)
	if err != nil {
		return err
	}
	if len(library.Templates) == 0 {
		return fmt.Errorf("no ABI region templates")
	}
	if err = save(output, library); err != nil {
		return err
	}
	if err = save(output+".training.json", labels); err != nil {
		return err
	}
	matched, total := 0, 0
	if codeDir != "" {
		files, err := filepath.Glob(filepath.Join(codeDir, "*.hex"))
		if err != nil {
			return err
		}
		if len(files) == 0 {
			return fmt.Errorf("no deployed hex files")
		}
		bindings := map[string][]*miner.ABIRegion{}
		for _, file := range files {
			raw, err := os.ReadFile(file)
			if err != nil {
				return err
			}
			code, err := hex.DecodeString(strings.TrimPrefix(strings.TrimSpace(string(raw)), "0x"))
			if err != nil {
				return fmt.Errorf("%s: %w", file, err)
			}
			regions := miner.DiscoverABIRegions(code, library)
			if len(regions) > 0 {
				bindings[strings.TrimSuffix(filepath.Base(file), ".hex")] = regions
				matched++
				total += len(regions)
			}
		}
		if err = save(output+".bindings.json", bindings); err != nil {
			return err
		}
	}
	fmt.Printf("ABI: %d compiler labels, %d templates, %d matched codes / %d regions -> %s\n", len(labels), len(library.Templates), matched, total, output)
	return nil
}
