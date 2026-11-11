// optimizer-audit reads historical blocks and uses the repository's normal
// StateProcessor with BaselineProcess-style state setup and optional prefetch.
// Candidate addresses/settings are supplied separately; layouts never imply
// an optimizer setting. The template miner independently verifies evidence.
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"strings"

	"go-eth1.16.5-evm/database"
)

var arithmeticTemplatesPath, arithmeticExecutionMode string
var benchmarkRepeats, benchmarkProcs int
var benchmarkModes string
var benchmarkLoopTiming bool
var benchmarkColdBindings bool

func main() {
	chain := flag.String("chain", "", "read-only chain database")
	snapshot := flag.String("snapshot", "", "read-only state at start-1")
	start := flag.Uint64("start", 23400001, "first block, inclusive")
	end := flag.Uint64("end", 23401000, "last block, inclusive")
	mode := flag.String("mode", "fixed", "fixed or cold")
	out := flag.String("out", "", "new output directory")
	flag.StringVar(&arithmeticTemplatesPath, "arithmetic-templates", "assets/arithmetic.json", "verified arithmetic template library")
	flag.StringVar(&benchmarkModes, "benchmark-modes", "baseline,sync", "baseline or sync, or a comma-separated pair")
	flag.IntVar(&benchmarkRepeats, "benchmark-repeats", 1, "repetitions per block")
	flag.IntVar(&benchmarkProcs, "benchmark-procs", 4, "GOMAXPROCS")
	flag.BoolVar(&benchmarkLoopTiming, "benchmark-loop-timing", true, "record EVM instruction execution latency")
	flag.BoolVar(&benchmarkColdBindings, "benchmark-cold-bindings", false, "independent initially empty plan caches")
	flag.Parse()
	if *chain == "" || *snapshot == "" || *out == "" || *start == 0 || *end < *start {
		fmt.Fprintln(os.Stderr, "chain, snapshot, out and a valid inclusive range are required")
		os.Exit(1)
	}
	if *mode != "fixed" && *mode != "cold" {
		fmt.Fprintln(os.Stderr, "mode must be fixed or cold")
		os.Exit(1)
	}
	for _, m := range strings.Split(benchmarkModes, ",") {
		if m != "baseline" && m != "sync" {
			fmt.Fprintln(os.Stderr, "only baseline and sync modes are supported")
			os.Exit(1)
		}
	}
	// Public experiments always execute the original deployed bytecode.
	os.Setenv("HISTORICAL_UNMODIFIED", "1")
	if *mode == "fixed" {
		os.Setenv("PAPER_WARM", "1")
		os.Unsetenv("REGION_POLICY")
	}
	if err := run(*chain, *snapshot, *start, *end, *mode, *out); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
func run(chainPath, snapshotPath string, start, end uint64, mode, out string) error {
	if err := os.MkdirAll(out, 0755); err != nil {
		return err
	}
	isolatedCallComparison = true
	cfg := *database.DefaultPebbleConfig
	cfg.File, cfg.Readonly = chainPath, true
	cfg.Cache, cfg.Handles = 512, 1024
	db, err := database.OpenDatabase(&cfg)
	if err != nil {
		return fmt.Errorf("read-only chain open: %w", err)
	}
	defer db.Close()
	if mode == "fixed" {
		return hashAblation(db, snapshotPath, start, end, out)
	}
	return benchmark(db, chainPath, snapshotPath, start, end, out)
}
func writeJSON(path string, value interface{}) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	enc := json.NewEncoder(f)
	enc.SetIndent("", "  ")
	return enc.Encode(value)
}
