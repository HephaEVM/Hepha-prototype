package main

import (
	"encoding/json"
	"fmt"
	"go-eth1.16.5-evm/core/vm"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"time"
)

// Separate from cold-start interference probes: no noise workloads or hidden
// preparation windows are run between paired timed methods.
func runPaperWarmCases(out string, enc *json.Encoder, rounds int) error {
	if len(libraryCases) == 0 {
		return fmt.Errorf("no paper cases")
	}
	n, err := strconv.Atoi(os.Getenv("HASH_ITERATIONS"))
	if err != nil || n < 1 {
		return fmt.Errorf("positive HASH_ITERATIONS required")
	}
	offset, _ := strconv.Atoi(os.Getenv("HASH_ORDER_OFFSET"))
	started := time.Now()
	for i, c := range libraryCases {
		row, err := c.run("warmup", 0, 1, false)
		if err != nil {
			return err
		}
		if err = enc.Encode(row); err != nil {
			return err
		}
		if (i+1)%48 == 0 {
			if err = probeWait(); err != nil {
				return err
			}
			fmt.Printf("paper preparation %d/%d elapsed=%s\n", i+1, len(libraryCases), time.Since(started).Round(time.Second))
		}
	}
	if err = probeWait(); err != nil {
		return err
	}
	if err = probeSettle(libraryCases); err != nil {
		return err
	}
	prepSeconds := time.Since(started).Seconds()
	var groups [][]libraryCase
	for _, c := range libraryCases {
		if len(groups) == 0 || groups[len(groups)-1][0].hash != c.hash {
			groups = append(groups, nil)
		}
		groups[len(groups)-1] = append(groups[len(groups)-1], c)
	}
	wantedArms := 2
	for _, group := range groups {
		if len(group) != wantedArms {
			return fmt.Errorf("expected %d paired methods for %s", wantedArms, group[0].hash)
		}
	}
	iterations := make(map[string]int)

	runtime.GC()
	frozen := probeStamp()

	timed := time.Now()
	for rep := 0; rep < rounds; rep++ {
		for gi := range groups {
			group := groups[(gi+rep)%len(groups)]
			batchN := n
			if k := iterations[group[0].hash]; k > 0 {
				batchN = k
			}
			for order := range group {
				c := group[(rep+offset+order)%len(group)]
				// One untimed call equalizes immediate code/data warming for each arm.
				if _, err = c.run("paper-prewarm", rep, 1, true); err != nil {
					return err
				}
				row, e := c.run("paper-timed", rep, batchN, true)
				if e != nil {
					return e
				}
				row.Order = order
				if err = enc.Encode(row); err != nil {
					return err
				}
			}
		}
		if probeStamp() != frozen {
			return fmt.Errorf("preparation changed during paper timed round %d", rep)
		}
		fmt.Printf("paper timed round=%d/%d calls=%d elapsed=%s\n", rep+1, rounds, len(groups), time.Since(timed).Round(time.Second))
	}
	for _, c := range libraryCases {
		row, e := c.run("postflight", 0, 1, true)
		if e != nil {
			return e
		}
		if err = enc.Encode(row); err != nil {
			return err
		}
	}
	return writeJSON(filepath.Join(out, "warm-validation.json"), map[string]any{"status": "complete", "mechanisms": vm.PaperMechanismConfiguration(), "calls": len(groups), "rounds": rounds, "iterations": n, "iterationOverrides": iterations, "preparationSeconds": prepSeconds, "measurementSeconds": time.Since(timed).Seconds(), "frozenPreparation": frozen, "finalPreparation": probeStamp()})
}
