package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"math/big"
	"os"
	"path/filepath"
	"runtime"
	"runtime/debug"
	"runtime/pprof"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/params"
	"go-eth1.16.5-evm/config"
	"go-eth1.16.5-evm/core"
	"go-eth1.16.5-evm/core/state"
	"go-eth1.16.5-evm/core/types"
	"go-eth1.16.5-evm/core/vm"
	"go-eth1.16.5-evm/database"
	"go-eth1.16.5-evm/ethdb"
	"go-eth1.16.5-evm/trie"
)

type benchmarkTx struct {
	preparation vm.HashPreparationStats
	compute     vm.ComputeHelperStats
	abi         vm.ABIScalarStats

	index       int
	hash        common.Hash
	to          string
	elapsed     time.Duration
	gas, status uint64
	stats       vm.ArithmeticStats
	loop        vm.InterpreterLoopTiming
	hot         vm.HotRegionStats
	regions     [10]uint64
	contract    bool
	timeline    float64
}

func regionCounters(s vm.RegionExecutionStats) [10]uint64 {
	return [10]uint64{s.Commits, s.DemandCommits, s.DemandInstructions, s.TransformCommits, s.BoundaryCommits, s.RelocatedCommits, s.LoopCommits, s.PartialCommits, s.BranchCommits, s.Bindings}
}
func regionDelta(a, b [10]uint64) [10]uint64 {
	for i := range a {
		a[i] -= b[i]
	}
	return a
}

func benchmarkStatsDelta(a, b vm.ArithmeticStats) vm.ArithmeticStats {
	return vm.ArithmeticStats{CacheHits: a.CacheHits - b.CacheHits, CacheMisses: a.CacheMisses - b.CacheMisses,
		MatchedRegions: a.MatchedRegions - b.MatchedRegions, RejectedRegions: a.RejectedRegions - b.RejectedRegions,
		Pruned: a.Pruned - b.Pruned, Extended: a.Extended - b.Extended, Checked: a.Checked - b.Checked,
		Failures: a.Failures - b.Failures, Fallbacks: a.Fallbacks - b.Fallbacks, ModelMismatches: a.ModelMismatches - b.ModelMismatches, FrameContexts: a.FrameContexts - b.FrameContexts, NativeChecks: a.NativeChecks - b.NativeChecks, FixedStackPrunes: a.FixedStackPrunes - b.FixedStackPrunes, ConstantPairPrunes: a.ConstantPairPrunes - b.ConstantPairPrunes}
}

func benchmarkValidate(block *types.Block, s *state.StateDB, result *core.ProcessResult) (common.Hash, common.Hash, error) {
	root := s.IntermediateRoot(config.MainnetChainConfig.IsEIP158(block.Number()))
	receipts := types.DeriveSha(result.Receipts, trie.NewStackTrie(nil))
	if err := s.Error(); err != nil {
		return root, receipts, err
	}
	if root != block.Root() || receipts != block.ReceiptHash() || result.GasUsed != block.GasUsed() {
		return root, receipts, fmt.Errorf("block %d validation mismatch: state=%s want=%s receipts=%s want=%s gas=%d want=%d", block.NumberU64(), root, block.Root(), receipts, block.ReceiptHash(), result.GasUsed, block.GasUsed())
	}
	return root, receipts, nil
}

func benchmarkCPU() (int64, error) {
	var r syscall.Rusage
	if err := syscall.Getrusage(syscall.RUSAGE_SELF, &r); err != nil {
		return 0, err
	}
	return r.Utime.Nano() + r.Stime.Nano(), nil
}

func benchmark(db ethdb.Database, chainPath, snapshotPath string, start, end uint64, out string) error {
	if path := os.Getenv("PREP_CPU_PROFILE"); path != "" {
		f, err := os.Create(path)
		if err != nil {
			return err
		}
		defer f.Close()
		if err = pprof.StartCPUProfile(f); err != nil {
			return err
		}
		defer pprof.StopCPUProfile()
	}
	if benchmarkRepeats < 1 || benchmarkProcs < 1 || arithmeticTemplatesPath == "" {
		return fmt.Errorf("benchmark needs positive repeats/procs and -arithmetic-templates")
	}
	modes := strings.Split(benchmarkModes, ",")
	if benchmarkModes == "" {
		modes = []string{"baseline", "sync", "baseline-aa", "sync-aa"}
	}
	seen := map[string]bool{}
	for _, m := range modes {

		seen[m] = true
	}
	// Never append to or silently replace an earlier experiment.
	manifestPath := filepath.Join(out, "benchmark-manifest.json")
	if _, err := os.Stat(manifestPath); err == nil {
		return fmt.Errorf("benchmark output already exists: %s", out)
	}
	oldProcs := runtime.GOMAXPROCS(benchmarkProcs)
	defer runtime.GOMAXPROCS(oldProcs)
	library, err := os.ReadFile(arithmeticTemplatesPath)
	if err != nil {
		return err
	}
	engine, err := vm.LoadArithmeticEngine(bytes.NewReader(library), vm.ArithmeticEngineOptions{CacheEntries: 16384})
	if err != nil {
		return err
	}
	if err = os.WriteFile(filepath.Join(out, "frozen-templates.json"), library, 0644); err != nil {
		return err
	}
	measureFrom := start
	if value := os.Getenv("BENCHMARK_MEASURE_FROM"); value != "" {
		v, err := strconv.ParseUint(value, 10, 64)
		if err != nil || v < start || v > end {
			return fmt.Errorf("invalid BENCHMARK_MEASURE_FROM")
		}
		measureFrom = v
	}
	freshTransactions := os.Getenv("BENCHMARK_FRESH_TX_EVM") == "1"
	var freshInstances uint64
	build, _ := debug.ReadBuildInfo()
	manifest := map[string]any{
		"schemaVersion": 2, "status": "running", "nativeBaseline": "RunNative from Geth v1.16.5, shared timing adaptation only", "runtime": "separate native and retained synchronous region loops", "startedAt": time.Now().Format(time.RFC3339),
		"startBlock": measureFrom, "replayStartBlock": start, "endBlock": end, "repetitions": benchmarkRepeats,
		"modes": modes, "gomaxprocs": benchmarkProcs,
		"goVersion": runtime.Version(), "os": runtime.GOOS, "arch": runtime.GOARCH, "logicalCPUs": runtime.NumCPU(), "build": build,
		"chainPath": chainPath, "snapshotPath": snapshotPath, "templatePath": arithmeticTemplatesPath,
		"templateSHA256": fmt.Sprintf("%x", sha256.Sum256(library)), "templateVersion": engine.Version(),
		"bindingCacheColdStart":  true,
		"freshEVMPerTransaction": freshTransactions,
		"completionTimer":        "Process timing; legacy drain column is zero. Background compilation must be completed separately before warm measurements.",
		"bindingCacheEntries":    16384, "gasPolicy": "original",
		"checkStrategy":     "retained synchronous native plans and exact structural fallback; original exceptional behavior",
		"loopTimingEnabled": benchmarkLoopTiming,
		"loopTimer":         "existing EVM.Run loop boundary, top-level only, nanoseconds; early returns and initcode included; excludes current frame setup/cleanup, includes nested calls and state/memory work; legacy Call microseconds also retained",
		"transactionTimer":  "ApplyTransactionWithEVM: complete state transition, checks, finalization and receipt construction; excludes sender recovery and SetTxContext",
		"blockTimer":        "StateProcessor.ProcessWithObserver; includes callbacks/system calls/finalization; excludes load/copy/warmup/root computation/commit/output",
		"warmup":            "Native-only state prefetch; no optimized warmup and no waiting for compilation",
		"diskPolicy":        "warm read-only state/chain caches; abort on any disk-cache miss, state iterator or ancient read during timing",
		"order":             "cyclic Latin-square rotation: (block offset + repetition) mod number_of_modes; every mode executes against an independent copy of identical parent state",
		"gc":                "default Go GC enabled, no forced collection in measured intervals; per-run counts recorded",
		"generalization":    "original deployed bytecode; 1000 consecutive blocks, all transactions; prior development overlapped subsets of this interval, not an independent held-out evaluation",
	}
	manifest["primaryTransactionMetric"] = "elapsed_ns"
	if benchmarkLoopTiming {
		manifest["primaryTransactionMetric"] = "loop_ns"
	}
	if err = writeJSON(manifestPath, manifest); err != nil {
		return err
	}
	chain := &benchmarkChainDB{Database: db, reads: newBenchmarkReadCache(db)}
	processor, err := newReplayProcessor(chain)
	if err != nil {
		return err
	}
	parent, err := database.GetBlockByNumber(chain, new(big.Int).SetUint64(start-1))
	if err != nil {
		return err
	}
	overlay, err := openOverlay(snapshotPath)
	if err != nil {
		return err
	}
	defer overlay.Close()
	overlay.reads = newBenchmarkReadCache(overlay.disk)
	all, err := database.NewAllDBForStateFromStore(parent.Root(), false, overlay)
	if err != nil {
		return err
	}
	defer all.Close()
	if err = benchmarkHistory(all, parent, start); err != nil {
		return err
	}
	file, err := os.Create(filepath.Join(out, "transaction-runs.csv"))
	if err != nil {
		return err
	}
	defer file.Close()
	w := csv.NewWriter(file)
	defer w.Flush()
	if err = w.Write([]string{"block", "tx_index", "tx_hash", "to", "repeat", "mode", "order", "elapsed_ns", "gas_used", "status", "pruned", "extended", "checked", "failures", "fallbacks", "model_mismatches", "cache_hits", "cache_misses", "frame_contexts", "native_checks", "loop_ns", "loop_runs", "loop_early_exits", "loop_creations", "legacy_lat_us", "fixed_stack_prunes", "constant_pair_prunes", "abi_attempts", "abi_skipped", "abi_instructions", "abi_gas", "abi_fallbacks", "abi_misses", "muldiv", "muldiv_wide", "muldiv_instructions", "muldiv_fallbacks", "packed", "packed_instructions", "packed_fallbacks", "packed_fields", "muldiv_fusions", "muldiv_fusion_instructions", "muldiv_roundups", "muldiv_fusion_fallbacks", "prep_attempts", "prep_skipped", "prep_instructions", "prep_fallbacks", "prep_words", "prep_single_output", "region_commits", "region_instructions", "region_helpers", "address_commits", "demand_commits", "demand_instructions", "transform_commits", "boundary_commits", "relocated_commits", "loop_commits", "partial_commits", "branch_commits", "bindings", "contract_call", "timeline_seconds"}); err != nil {
		return err
	}
	blockFile, err := os.Create(filepath.Join(out, "block-runs.csv"))
	if err != nil {
		return err
	}
	defer blockFile.Close()
	bw := csv.NewWriter(blockFile)
	defer bw.Flush()
	if err = bw.Write([]string{"block", "repeat", "mode", "order", "transactions", "elapsed_ns", "sum_tx_ns", "cpu_ns", "gc_cycles", "allocated_bytes", "disk_reads", "blocked_reads", "state_root", "receipt_root", "gas_used", "pruned", "checked", "fallbacks", "model_mismatches", "cache_misses", "preparation_drain_ns", "completion_ns"}); err != nil {
		return err
	}
	abiEngine, err := vm.LoadABIRegionDiagnostic(os.Getenv("ABI_SCALAR_TEMPLATES"), os.Getenv("ABI_EXPANDED_TEMPLATES"))
	if err != nil {
		return err
	}
	abiEngine.Optimize = true
	abiStats := func() vm.ABIScalarStats { return abiEngine.Stats }
	var prepStats vm.HashPreparationStats
	var computeStats vm.ComputeHelperStats
	var regionStats vm.RegionExecutionStats
	// Data warming and fixed-input capture finish before ANY online worker is
	// created. Thus a cold plan cannot gain free progress during native prefetch.
	study, err := newReadinessStudy()
	if err != nil {
		return err
	}
	if study != nil {
		core.ExperimentRegionHistoricalHashObserver = regionCostObserveHash
		defer func() { core.ExperimentRegionHistoricalHashObserver = nil }()
	}
	if os.Getenv("PAPER_STREAM") == "1" {
		preloadEnd := end
		if raw := os.Getenv("BENCHMARK_PRELOAD_END"); raw != "" {
			value, parseErr := strconv.ParseUint(raw, 10, 64)
			if parseErr != nil || value < end {
				return fmt.Errorf("BENCHMARK_PRELOAD_END must be at least the measured end block")
			}
			preloadEnd = value
		}
		manifest["dataPreloadEndBlock"] = preloadEnd
		started := time.Now()
		root := parent.Root()
		previousHash := parent.Hash()
		core.ReplayHistoricalHashes[start-1] = previousHash
		for number := start; number <= preloadEnd; number++ {
			if err = all.UpdateStateDB(root); err != nil {
				return err
			}
			b, er := database.GetBlockByNumber(chain, new(big.Int).SetUint64(number))
			if er != nil {
				return er
			}
			if b.NumberU64() != number || b.ParentHash() != previousHash {
				return fmt.Errorf("noncontiguous paper preload history at %d", number)
			}
			if study != nil {
				regionCostHeaders[number] = b.Hash()
			}
			state := all.Copy()
			c, observer := study.capture(vm.Config{}, state, b)
			result, er := processor.ProcessWithObserver(b, state, c, observer)
			if er != nil {
				return er
			}
			if _, _, er = benchmarkValidate(b, state, result); er != nil {
				return er
			}
			root, _, er = state.CommitWithUpdate(number, config.MainnetChainConfig.IsEIP158(b.Number()), config.MainnetChainConfig.IsCancun(b.Number(), b.Time()))
			if er != nil {
				return er
			}
			if root != b.Root() {
				return fmt.Errorf("preload commit mismatch %d", number)
			}
			// Retain authenticated hashes for replay after the header-chain LRU
			// has advanced beyond this prefix. Both comparison arms share the
			// frozen map; no header lookup can trigger disk I/O during timing.
			previousHash = b.Hash()
			core.ReplayHistoricalHashes[number] = previousHash
			if er = all.TrieDB.Commit(root, false); er != nil {
				return er
			}
			if number%100 == 0 {
				fmt.Printf("paper data preload validated through %d\n", number)
			}
		}
		if err = all.UpdateStateDB(parent.Root()); err != nil {
			return err
		}
		manifest["dataPreloadSeconds"] = time.Since(started).Seconds()
		manifest["dataPreloadBeforeWorkers"] = true
		if os.Getenv("BENCHMARK_PRELOAD_GC") == "1" {
			// Discard preload garbage before any preparer exists. This is a
			// symmetric benchmark startup boundary, not a timed-loop exemption:
			// automatic GC remains enabled throughout the measured replay.
			var before, after runtime.MemStats
			runtime.ReadMemStats(&before)
			gcStart := time.Now()
			runtime.GC()
			gcNS := time.Since(gcStart).Nanoseconds()
			runtime.ReadMemStats(&after)
			manifest["preloadGC"] = map[string]any{
				"nanoseconds": gcNS, "heapBeforeBytes": before.HeapAlloc, "heapAfterBytes": after.HeapAlloc,
				"completedCycles": after.NumGC - before.NumGC, "nextGCBytes": after.NextGC,
				"beforePreparationWorkers": true, "automaticGCDuringReplay": true,
				"scope": "Startup after native data preload; excluded from EVM lat and replay wall time, recorded separately",
			}
			fmt.Printf("preload GC complete elapsed=%s heapMiB=%d->%d\n", time.Duration(gcNS), before.HeapAlloc/(1024*1024), after.HeapAlloc/(1024*1024))
		}
	}
	var preparers []*vm.BackgroundPreparation
	makeConfig := func(e *vm.ArithmeticEngine, mode string) vm.Config {
		cfg := config.DefaultVmConfig
		cfg.Tracer, cfg.Arithmetic, cfg.Pruning, cfg.ABIScalar = nil, nil, nil, nil

		cfg.StreamRegions = nil

		if benchmarkLoopTiming {
			cfg.LoopTiming = new(vm.InterpreterLoopTiming)
		}
		if mode == "sync" || mode == "sync-aa" {
			cfg.RegionExecutionStats = &regionStats
			cfg = cfg.WithPruning(vm.PruningConfig{Engine: e, ABI: abiEngine, ComputeStats: &computeStats, HashStats: &prepStats})
		}
		if (mode == "sync" || mode == "sync-aa") && os.Getenv("HOT_REGION_CORPUS") != "" {
			var er error
			cfg.HotRegions, er = vm.NewHotRegionRuntime(e, os.Getenv("HOT_REGION_CORPUS"), os.Getenv("HOT_REGION_SELECTED"))
			if er != nil {
				panic(er)
			}
		}
		if mode == "sync" {
			p, err := experimentAdmissionPreparation(e, abiEngine)
			if err != nil {
				panic(err)
			}
			cfg.Pruning.Preparation = p
			preparers = append(preparers, p)
		}
		// Optional reporting only. Keep the shared latency timer, arithmetic
		// safety/gas checks and scheduling observations unchanged.
		if mode == "stream" || mode == "sync" && os.Getenv("BENCHMARK_STREAM_REGIONS") == "1" {
			cfg = cfg.WithStreamRegions(new(vm.StreamRegions))
		}
		return executionReportingConfig(cfg, os.Getenv("COLD_EXECUTION_COUNTERS") != "off")
	}
	warmConfigs := make([]vm.Config, len(modes))
	for i, m := range modes {
		armEngine := engine
		if i > 0 {
			armEngine, err = vm.LoadArithmeticEngine(bytes.NewReader(library), vm.ArithmeticEngineOptions{CacheEntries: 16384})
			if err != nil {
				return err
			}
		}
		warmConfigs[i] = makeConfig(armEngine, m)

	}
	if os.Getenv("PAPER_FIXED_WORKING_SET") == "1" {
		if os.Getenv("PAPER_STREAM") != "1" || study == nil || len(preparers) != 1 || len(modes) != 1 || modes[0] != "sync" {
			return fmt.Errorf("fixed working set requires native preload, frozen targets, and one sync preparation arm")
		}
		ce, err := config.CreateConsensusEngine(config.MainnetChainConfig)
		if err != nil {
			return err
		}
		headers, err := core.NewHeaderChain(chain, config.MainnetChainConfig, ce)
		if err != nil {
			return err
		}
		return study.runWorkingSet(out, headers, warmConfigs[0], preparers[0], overlay.reads, chain.reads, manifest)
	}
	measuredConfigs := make([][]vm.Config, benchmarkRepeats)
	for repeat := range measuredConfigs {
		measuredConfigs[repeat] = warmConfigs
		if benchmarkColdBindings {
			measuredConfigs[repeat] = make([]vm.Config, len(modes))
			for i, m := range modes {
				e := engine
				if i > 0 {
					e, err = vm.LoadArithmeticEngine(bytes.NewReader(library), vm.ArithmeticEngineOptions{CacheEntries: 16384})
					if err != nil {
						return err
					}
				}
				measuredConfigs[repeat][i] = makeConfig(e, m)
			}
		}
	}

	checks := []map[string]any{}
	parentRoot, parentHash := parent.Root(), parent.Hash()
	totalTx := 0
	prepVisits := map[string]*vm.HashPreparationVisit{}
	var boundaryRows []map[string]any
	var boundaryNS int64
	var regionWindows []map[string]any
	if study != nil && os.Getenv("PAPER_STREAM") == "1" && len(preparers) > 0 {
		study.checkpoint(preparers[0], 0, end-measureFrom+1, 0)
	}
	runStart := time.Now()
	windowWarm := os.Getenv("COLD_WINDOW_WARM") == "1"
	windowTraining := windowWarm
	windowStable := 0
	windowPrevious := ""
	windowReplayStart := start
	windowReplayRoot, windowReplayHash := parent.Root(), parent.Hash()
	windowTrainingRows := []map[string]any{}
	for windowPass := 0; ; windowPass++ {
		parentRoot, parentHash = windowReplayRoot, windowReplayHash
		totalTx = 0
		checks = nil
		boundaryRows = nil
		regionWindows = nil
		boundaryNS = 0
		if windowWarm && !windowTraining {
			fmt.Println("starting frozen fully-warm measurement")
		}
		for number := windowReplayStart; number <= end; number++ {
			blockStart := time.Now()
			if err = all.UpdateStateDB(parentRoot); err != nil {
				return err
			}
			block, err := database.GetBlockByNumber(chain, new(big.Int).SetUint64(number))
			if err != nil {
				return err
			}
			if block.ParentHash() != parentHash {
				return fmt.Errorf("parent discontinuity at %d", number)
			}
			if study != nil {
				regionCostHeaders[number] = block.Hash()
			}
			if number < measureFrom {
				state := all.Copy()
				result, err := processor.ProcessWithObserver(block, state, makeConfig(engine, "baseline"), &core.TransactionObserver{})
				if err != nil {
					return fmt.Errorf("prefix block %d: %w", number, err)
				}
				if _, _, err = benchmarkValidate(block, state, result); err != nil {
					return err
				}
				root, _, err := state.CommitWithUpdate(number, config.MainnetChainConfig.IsEIP158(block.Number()), config.MainnetChainConfig.IsCancun(block.Number(), block.Time()))
				if err != nil {
					return err
				}
				if root != block.Root() {
					return fmt.Errorf("prefix commit mismatch %d", number)
				}
				if err = all.TrieDB.Commit(root, false); err != nil {
					return err
				}
				parentRoot, parentHash = root, block.Hash()
				if number%20 == 0 {
					fmt.Printf("validated prefix through block=%d\\n", number)
				}
				continue
			}
			if number == measureFrom {
				windowReplayStart = measureFrom
				windowReplayRoot, windowReplayHash = parentRoot, parentHash
				runStart = time.Now()
			}
			if (vm.RegionWindowEnabled() || os.Getenv("FINAL_BACKGROUND") == "1") && (!windowWarm || windowTraining) {
				windowStart := time.Now()
				row := regionBlockWindow(block, all.Copy(), processor, measuredConfigs[0][0], preparers)
				row["totalNS"] = time.Since(windowStart).Nanoseconds()
				regionWindows = append(regionWindows, row)
				if windowTraining {
					for _, p := range preparers {
						if er := p.WindowSettle(); er != nil {
							return er
						}
					}
				}
			}
			warmStart := time.Now()

			warmBefore := engine.Stats()
			if os.Getenv("PAPER_STREAM") != "1" {
				sWarm := all.Copy()
				for _, p := range preparers {
					p.EnterForeground()
				}
				nativeCfg, nativeObs := study.capture(makeConfig(engine, "baseline"), sWarm, block)
				resultWarm, err := processor.ProcessWithObserver(block, sWarm, nativeCfg, nativeObs)
				for _, p := range preparers {
					p.LeaveForeground()
				}
				if err != nil {
					return fmt.Errorf("state prefetch block %d: %w", number, err)
				}
				if _, _, err = benchmarkValidate(block, sWarm, resultWarm); err != nil {
					return err
				}
				sWarm = nil
			}

			warmElapsed := time.Since(warmStart)
			warmStats := benchmarkStatsDelta(engine.Stats(), warmBefore)
			var canonical *state.StateDB
			for repeat := 0; repeat < benchmarkRepeats; repeat++ {
				configs := measuredConfigs[repeat]
				for order := 0; order < len(modes); order++ {
					i := (int(number-start) + repeat + order) % len(modes)
					if len(modes) == 4 {
						i = [4][4]int{{0, 1, 3, 2}, {1, 2, 0, 3}, {2, 3, 1, 0}, {3, 0, 2, 1}}[(int(number-start)+repeat)%4][order]
					}
					activeEngine := engine
					if cfg := configs[i].Arithmetic; cfg != nil {
						activeEngine = cfg.Engine
					}
					// Legacy CSV columns remain zero to keep existing analysis readers usable.
					s := all.Copy()
					rows := make([]benchmarkTx, 0, len(block.Transactions()))
					var prepBefore vm.HashPreparationStats
					var computeBefore vm.ComputeHelperStats
					var abiBefore vm.ABIScalarStats
					var before vm.ArithmeticStats
					var hotBefore vm.HotRegionStats
					var regionBefore [10]uint64
					var contractCall bool
					observer := &core.TransactionObserver{
						FreshEVMPerTransaction: freshTransactions,
						Before: func(index int, tx *types.Transaction) {

							regionBefore = regionCounters(regionStats)
							contractCall = tx.To() != nil && s.GetCodeSize(*tx.To()) > 0
							prepBefore = prepStats
							if h := configs[i].HotRegions; h != nil {
								hotBefore = h.Stats
							}
							computeBefore = computeStats
							abiBefore = abiStats()
							before = activeEngine.Stats()
							if t := configs[i].LoopTiming; t != nil {
								*t = vm.InterpreterLoopTiming{}
							}
						},
						After: func(index int, tx *types.Transaction, receipt *types.Receipt, elapsed time.Duration) {

							to := "CREATE"
							if tx.To() != nil {
								to = tx.To().Hex()
							}
							var loop vm.InterpreterLoopTiming
							if t := configs[i].LoopTiming; t != nil {
								loop = *t
							}
							var hot vm.HotRegionStats
							if h := configs[i].HotRegions; h != nil {
								hot = h.Stats.Sub(hotBefore)
							}
							rows = append(rows, benchmarkTx{prepStats.Sub(prepBefore), computeStats.Sub(computeBefore), abiStats().Sub(abiBefore), index, tx.Hash(), to, elapsed, receipt.GasUsed, receipt.Status, benchmarkStatsDelta(activeEngine.Stats(), before), loop, hot, regionDelta(regionCounters(regionStats), regionBefore), contractCall, time.Since(runStart).Seconds()})
						},
					}

					var memBefore, memAfter runtime.MemStats
					runtime.ReadMemStats(&memBefore)
					cpuBefore, err := benchmarkCPU()
					if err != nil {
						return err
					}
					abiMisses := abiStats().Misses
					statsBefore := activeEngine.Stats()

					sd, sb := overlay.reads.counts()
					cd, cb := chain.reads.counts()
					untimedCapture := false
					overlay.reads.setStrict(!untimedCapture)
					chain.reads.setStrict(!untimedCapture)
					now := time.Now()
					for _, p := range preparers {
						p.EnterForeground()
					}
					result, processErr := processor.ProcessWithObserver(block, s, configs[i], observer)
					for _, p := range preparers {
						p.LeaveForeground()
					}
					elapsed := time.Since(now)
					if freshTransactions && observer.FreshInstances != len(block.Transactions()) {
						return fmt.Errorf("fresh EVM count mismatch at block %d: %d", number, observer.FreshInstances)
					}
					freshInstances += uint64(observer.FreshInstances)

					if false && abiStats().Misses != abiMisses {
						return fmt.Errorf("cold ABI binding")
					}
					var drain time.Duration
					overlay.reads.setStrict(false)
					chain.reads.setStrict(false)
					cpuAfter, err := benchmarkCPU()
					if err != nil {
						return err
					}
					runtime.ReadMemStats(&memAfter)
					stats := benchmarkStatsDelta(activeEngine.Stats(), statsBefore)
					sd2, sb2 := overlay.reads.counts()
					cd2, cb2 := chain.reads.counts()
					diskReads, blocked := sd2-sd+cd2-cd, sb2-sb+cb2-cb
					if !untimedCapture && (diskReads != 0 || blocked != 0) {
						return fmt.Errorf("unprepared I/O invalidates block %d mode %s repeat %d: disk=%d blocked=%d", number, modes[i], repeat, diskReads, blocked)
					}
					if processErr != nil {
						return fmt.Errorf("block %d mode %s repeat %d: %w", number, modes[i], repeat, processErr)
					}
					if false && stats.CacheMisses != 0 {
						return fmt.Errorf("cold binding invalidates block %d mode %s: %d misses", number, modes[i], stats.CacheMisses)
					}
					if stats.ModelMismatches != 0 {
						return fmt.Errorf("arithmetic model mismatch: %d", stats.ModelMismatches)
					}
					root, receiptRoot, err := benchmarkValidate(block, s, result)
					if err != nil {
						return fmt.Errorf("mode %s repeat %d: %w", modes[i], repeat, err)
					}
					if len(rows) != len(block.Transactions()) {
						return fmt.Errorf("incomplete transaction timing")
					}
					var sumTx time.Duration
					for _, row := range rows {
						sumTx += row.elapsed
						fields := []string{fmt.Sprint(number), fmt.Sprint(row.index), row.hash.Hex(), row.to, fmt.Sprint(repeat), modes[i], fmt.Sprint(order), fmt.Sprint(row.elapsed.Nanoseconds()), fmt.Sprint(row.gas), fmt.Sprint(row.status)}
						for _, value := range []uint64{row.stats.Pruned, row.stats.Extended, row.stats.Checked, row.stats.Failures, row.stats.Fallbacks, row.stats.ModelMismatches, row.stats.CacheHits, row.stats.CacheMisses, row.stats.FrameContexts, row.stats.NativeChecks} {
							fields = append(fields, strconv.FormatUint(value, 10))
						}
						fields = append(fields, strconv.FormatInt(row.loop.Nanoseconds, 10), strconv.FormatUint(row.loop.Runs, 10), strconv.FormatUint(row.loop.EarlyExits, 10), strconv.FormatUint(row.loop.CreationRuns, 10), strconv.FormatInt(row.loop.LegacyCallMicros, 10), strconv.FormatUint(row.stats.FixedStackPrunes, 10), strconv.FormatUint(row.stats.ConstantPairPrunes, 10))
						for _, value := range []uint64{row.abi.Attempts, row.abi.Skipped, row.abi.Instructions, row.abi.Gas, row.abi.Fallbacks, row.abi.Misses} {
							fields = append(fields, strconv.FormatUint(value, 10))
						}
						for _, value := range []uint64{row.compute.MulDiv, row.compute.MulDivWide, row.compute.MulDivInstructions, row.compute.MulDivFallbacks, row.compute.Packed, row.compute.PackedInstructions, row.compute.PackedFallbacks, row.compute.PackedFields, row.compute.MulDivFusions, row.compute.MulDivFusionInstructions, row.compute.MulDivRoundUps, row.compute.MulDivFusionFallbacks} {
							fields = append(fields, strconv.FormatUint(value, 10))
						}
						for _, v := range []uint64{row.preparation.Attempts, row.preparation.Skipped, row.preparation.Instructions, row.preparation.Fallbacks, row.preparation.Words, row.preparation.SingleOutput} {
							fields = append(fields, strconv.FormatUint(v, 10))
						}
						fields = append(fields, fmt.Sprint(row.hot.Commits), fmt.Sprint(row.hot.Instructions), fmt.Sprint(row.hot.Helpers))
						for _, v := range row.regions {
							fields = append(fields, strconv.FormatUint(v, 10))
						}
						fields = append(fields, strconv.FormatBool(row.contract), strconv.FormatFloat(row.timeline, 'f', 9, 64))
						if !windowTraining {
							if err = w.Write(fields); err != nil {
								return err
							}
						}
					}
					if !windowTraining {
						if err = bw.Write([]string{fmt.Sprint(number), fmt.Sprint(repeat), modes[i], fmt.Sprint(order), fmt.Sprint(len(rows)), fmt.Sprint(elapsed.Nanoseconds()), fmt.Sprint(sumTx.Nanoseconds()), fmt.Sprint(cpuAfter - cpuBefore), fmt.Sprint(memAfter.NumGC - memBefore.NumGC), fmt.Sprint(memAfter.TotalAlloc - memBefore.TotalAlloc), fmt.Sprint(diskReads), fmt.Sprint(blocked), root.Hex(), receiptRoot.Hex(), fmt.Sprint(result.GasUsed), fmt.Sprint(stats.Pruned), fmt.Sprint(stats.Checked), fmt.Sprint(stats.Fallbacks), fmt.Sprint(stats.ModelMismatches), fmt.Sprint(stats.CacheMisses), fmt.Sprint(drain.Nanoseconds()), fmt.Sprint((elapsed + drain).Nanoseconds())}); err != nil {
							return err
						}
					}
					if i == 0 && canonical == nil {
						canonical = s
					}
				}
			}
			w.Flush()
			bw.Flush()
			if err = w.Error(); err != nil {
				return err
			}
			if err = bw.Error(); err != nil {
				return err
			}
			// Only one already validated baseline state advances the historical chain.
			root, _, err := canonical.CommitWithUpdate(number, config.MainnetChainConfig.IsEIP158(block.Number()), config.MainnetChainConfig.IsCancun(block.Number(), block.Time()))
			if err != nil {
				return err
			}
			if root != block.Root() {
				return fmt.Errorf("commit root mismatch %d", number)
			}
			if err = all.TrieDB.Commit(root, false); err != nil {
				return err
			}
			parentRoot, parentHash = root, block.Hash()
			totalTx += len(block.Transactions())
			if study != nil && !windowTraining && len(preparers) > 0 {
				study.checkpoint(preparers[0], number-measureFrom+1, end-measureFrom+1, time.Since(runStart).Seconds())
			}
			var step vm.ValuePreparationStep
			if number < end {
				for _, p := range preparers {
					q := vm.ValuePreparationStep{}
					if !vm.RegionWindowEnabled() {
						q = p.AdvanceValuePreparation()
					}
					step.Nanoseconds += q.Nanoseconds
					step.Analyses += q.Analyses
					step.CompileBatches += q.CompileBatches
					step.CreditNS = q.CreditNS
				}
			}
			boundaryNS += step.Nanoseconds
			boundaryRows = append(boundaryRows, map[string]any{"block": number, "step": step, "cumulativeNS": boundaryNS})
			registered, capacity, rejected := vm.FlowDemandUsage()
			var footprint runtime.MemStats
			runtime.ReadMemStats(&footprint)
			bg := []vm.BackgroundPreparationStats{}
			for _, p := range preparers {
				bg = append(bg, p.Stats())
			}
			checks = append(checks, map[string]any{"timelineSeconds": time.Since(runStart).Seconds(), "background": bg, "families": vm.FlowFamilyStats(), "preparationResources": vm.FlowPreparationResources(), "activePlans": registered, "planCapacity": capacity, "rejectedPlans": rejected, "heapAllocBytes": footprint.HeapAlloc, "gcCycles": footprint.NumGC, "nextGCBytes": footprint.NextGC, "gcPauseTotalNS": footprint.PauseTotalNs, "preparationCache": vm.PreparationCacheStats(), "block": number, "transactions": len(block.Transactions()), "validatedMeasuredRuns": len(modes) * benchmarkRepeats, "stateRoot": root, "receiptRoot": block.ReceiptHash(), "gasUsed": block.GasUsed(), "warmupSeconds": warmElapsed.Seconds(), "warmupStats": warmStats, "elapsedSeconds": time.Since(blockStart).Seconds()})
			if err = writeJSON(filepath.Join(out, "block-validation.json"), checks); err != nil {
				return err
			}
			fmt.Printf("benchmark block=%d txs=%d validated=%d warmup=%s total=%s progress=%d/%d\n", number, len(block.Transactions()), len(modes)*benchmarkRepeats, warmElapsed.Round(time.Millisecond), time.Since(blockStart).Round(time.Millisecond), number-start+1, end-start+1)
			fmt.Printf("elapsed=%.3fs activePlans=%d heapMiB=%d background=%+v resources=%v\n", time.Since(runStart).Seconds(), registered, footprint.HeapAlloc/(1024*1024), bg, vm.FlowPreparationResources())
		}

		if !windowTraining {
			break
		}
		for _, p := range preparers {
			if er := p.WindowSettle(); er != nil {
				return er
			}
		}
		fingerprint := ""
		for _, p := range preparers {
			fingerprint += p.WindowFingerprint()
		}
		if fingerprint == windowPrevious {
			windowStable++
		} else {
			windowStable = 0
		}
		windowPrevious = fingerprint
		row := map[string]any{"pass": windowPass + 1, "stablePasses": windowStable, "seconds": time.Since(runStart).Seconds(), "resources": vm.FlowPreparationResources(), "fingerprint": fingerprint}
		windowTrainingRows = append(windowTrainingRows, row)
		if err = writeJSON(filepath.Join(out, "warm-training.json"), windowTrainingRows); err != nil {
			return err
		}
		fmt.Printf("warm training pass=%d stable=%d resources=%v\n", windowPass+1, windowStable, vm.FlowPreparationResources())
		if windowStable >= 2 {
			for i, c := range warmConfigs {
				if c.Pruning != nil && c.Pruning.Preparation != nil {
					warmConfigs[i] = c.Pruning.Preparation.WindowFrozenConfig(c)
				}
			}
			for i := range measuredConfigs {
				measuredConfigs[i] = warmConfigs
			}
			abiEngine = warmConfigs[0].Pruning.ABI
			windowTraining = false
		} else if windowPass >= 7 {
			return fmt.Errorf("warm preparation did not stabilize after 8 full passes")
		}
	}
	manifest["windowStudy"] = map[string]any{"fullyWarm": windowWarm, "training": windowTrainingRows, "compilerBoundary": os.Getenv("COLD_WINDOW_COMPILER"), "windowNS": vm.WindowBudgetNS(), "sameAdmission": true, "unobservedGateBypassed": false}
	manifest["status"], manifest["finishedAt"], manifest["transactions"] = "complete", time.Now().Format(time.RFC3339), totalTx
	manifest["measurementEngineStats"] = func() []map[string]any {
		entries := []map[string]any{}
		for rep, configs := range measuredConfigs {
			for i, cfg := range configs {
				if cfg.Arithmetic != nil {
					row := map[string]any{"repeat": rep, "mode": modes[i], "arithmetic": cfg.Arithmetic.Engine.Stats()}
					entries = append(entries, row)
				}
			}
		}
		return entries
	}()
	manifest["elapsedSeconds"], manifest["engineStatsIncludingWarmup"] = time.Since(runStart).Seconds(), engine.Stats()
	rawPreparations, err := json.MarshalIndent(prepVisits, "", "  ")
	if err != nil {
		return err
	}
	if err = os.WriteFile(filepath.Join(out, "preparation-visits.json"), rawPreparations, 0644); err != nil {
		return err
	}
	manifest["computeStatsIncludingWarmup"] = computeStats
	manifest["preparationStatsIncludingWarmup"] = prepStats
	manifest["preparationPolicy"] = "Original gas, proved straight-line constant-address fully initialized 32/64-byte KECCAK inputs. Original memory writes and hash preserved; state accesses excluded. Existing arithmetic/ABI/mulDiv/ceil identical. Direct operand recipes; all observations retained."
	manifest["order"] = "Four-arm Williams balanced order (cyclic rotation for other supported subsets); equal same-arm untimed warmup before every observation."
	manifest["hotRegionScope"] = "Optional prebuilt region corpus selected by HOT_REGION_CORPUS/HOT_REGION_SELECTED, plus structurally discovered connected regions. Selection and training provenance must be reported separately."
	manifest["hotRegionSelection"] = os.Getenv("HOT_REGION_SELECTED")
	manifest["combinedPolicy"] = "Retained synchronous arithmetic, ABI, mulDiv, hash and connected regions; original gas. Background preparation only, no asynchronous arithmetic checks."
	manifest["replicationWarmup"] = "Independent arithmetic/binding namespaces; same-arm untimed replay before every observation. Underlying ABI engine may share plans."
	manifest["databaseCacheMiB"] = 512
	stateCacheMiB, stateHandles := 512, 1024
	manifest["stateDatabaseBudgetMiB"] = stateCacheMiB
	manifest["stateBlockCacheMiB"] = stateCacheMiB / 2
	manifest["databaseHandles"] = 1024
	manifest["stateDatabaseHandles"] = stateHandles

	if err = writeJSON(filepath.Join(out, "abi-visits.json"), abiEngine.Visits); err != nil {
		return err
	}
	if err = writeJSON(filepath.Join(out, "abi-full-visits.json"), abiEngine.Visits); err != nil {
		return err
	}
	if err = writeJSON(filepath.Join(out, "abi-region-visits.json"), abiEngine.RegionVisits); err != nil {
		return err
	}
	if err = writeJSON(filepath.Join(out, "abi-native-visits.json"), abiEngine.NativeVisits); err != nil {
		return err
	}
	// Foreground completion is saved before cleanup. No Wait occurs during replay.
	manifest["bindingCacheColdStart"] = !windowWarm
	manifest["status"] = "foreground-complete"
	manifest["freshTransactionEVMInstances"] = freshInstances
	if freshTransactions {
		manifest["transactionTimer"] = "Fresh EVM construction plus ApplyTransactionWithEVM and receipt construction. Excludes sender recovery, observer callbacks and StateDB.SetTxContext."
	}
	manifest["foregroundEndResources"] = vm.FlowPreparationResources()
	manifest["boundaryPreparationNS"] = boundaryNS
	manifest["valuePolicy"] = vm.ValuePolicyReport()
	manifest["hotcallPolicy"] = vm.HotcallPolicyReport()
	manifest["deliveryPolicy"] = vm.DeliveryPolicyReport()
	manifest["tradePolicy"] = vm.TradePolicyReport()
	manifest["regionPriority"] = vm.RegionPriorityReport()
	if err = writeJSON(filepath.Join(out, "preblock-windows.json"), regionWindows); err != nil {
		return err
	}
	if err = writeJSON(filepath.Join(out, "boundary-preparation.json"), boundaryRows); err != nil {
		return err
	}
	manifest["hotspotPolicy"] = vm.HotspotPolicyReport()
	endBG := []vm.BackgroundPreparationStats{}
	for _, p := range preparers {
		endBG = append(endBG, p.Stats())
	}
	manifest["foregroundEndBackground"] = endBG
	manifest["executionDiagnosticCounters"] = os.Getenv("COLD_EXECUTION_COUNTERS")
	manifest["diagnosticConfigurationVersion"] = "v2-preserved-through-NewEVM"
	manifest["gc"] = "Go GC enabled; identical GOGC and GOMEMLIMIT for paired native/candidate; no forced collection"
	if _, ok := manifest["preloadGC"]; ok {
		manifest["gc"] = "Go GC enabled with unchanged GOGC/GOMEMLIMIT; one startup collection after native preload and before workers for both arms; no forced collection during replay"
	}
	manifest["gogc"] = os.Getenv("GOGC")
	manifest["gomemlimit"] = os.Getenv("GOMEMLIMIT")
	manifest["admissionThreshold"] = os.Getenv("COLD_ADMISSION_THRESHOLD")
	manifest["preparationDuty"] = os.Getenv("COLD_PREPARATION_DUTY")
	manifest["order"] = "One sequential pass; native and optimized are separate processes. Native-only state prefetch before each timed block. No optimized warmup or preparation wait during replay."
	manifest["warmup"] = "Native-only state prefetch; region caches initially empty and persistent across blocks."
	manifest["replicationWarmup"] = "None for optimized execution; exactly one timed pass per process."
	manifest["generalization"] = "100 consecutive blocks, original deployed bytecode, no checked/SafeMath injection; development overlap, not held out."
	if boundaryNS > 0 {
		manifest["order"] = "One sequential pass; native state prefetch; bounded synchronous preparation after each measured block except the last, separately recorded and included in completion elapsed."
	}
	manifest["completionTimer"] = "foreground elapsed includes native state prefetch, timed processing, verification, commits and output; final cleanup is separate."
	if os.Getenv("PAPER_STREAM") == "1" {
		manifest["generalization"] = fmt.Sprintf("Original %d-block deployment trace; development overlap disclosed; no checked/SafeMath injection", end-start+1)
		manifest["warmup"] = "Whole dataset native data preload and fixed-input capture before any online worker; no per-block native prefetch during online progression"
		manifest["completionTimer"] = "Replay wall time after data preload, including windows, online preparation interference, state commits, validation and output; not node TPS"
		if windowWarm {
			manifest["order"] = "Plans converge under the declared admission policy, then immutable ready snapshots are measured; methods rotate by block"
		} else {
			manifest["order"] = "One chronological candidate stream; independent Vanilla processes; no baseline or probe calls interleaved with online preparation"
		}
	}

	if err = writeJSON(manifestPath, manifest); err != nil {
		return err
	}
	fmt.Printf("FOREGROUND COMPLETE blocks=%d txs=%d seconds=%.3f; cleaning up workers separately\n", end-start+1, totalTx, time.Since(runStart).Seconds())
	cleanup := time.Now()
	for _, p := range preparers {
		p.Close()
		p.CancelValueCompilation()
	}
	events, er := vm.WaitDemand()
	if er != nil {
		return er
	}
	if err = writeJSON(filepath.Join(out, "demand-events.json"), events); err != nil {
		return err
	}
	failures := 0
	for _, event := range events {
		if event.Error != "" {
			failures++
		}
	}
	manifest["compileFailures"] = failures
	manifest["cleanupSeconds"] = time.Since(cleanup).Seconds()
	manifest["status"] = "complete"
	fmt.Printf("CLEANUP COMPLETE seconds=%.3f compileFailures=%d\n", time.Since(cleanup).Seconds(), failures)
	if err = writeJSON(manifestPath, manifest); err != nil {
		return err
	}
	if err := vm.ExperimentOracleEnd(); err != nil {
		return err
	}
	if study != nil {
		ce, err := config.CreateConsensusEngine(config.MainnetChainConfig)
		if err != nil {
			return err
		}
		headers, err := core.NewHeaderChain(chain, config.MainnetChainConfig, ce)
		if err != nil {
			return err
		}
		if err = study.finish(out, headers, warmConfigs[0], overlay.reads, chain.reads); err != nil {
			return err
		}
	}
	f, err := os.Create(filepath.Join(out, "arithmetic-feedback.json"))
	if err != nil {
		return err
	}
	writeErr := engine.WriteFeedback(f)
	closeErr := f.Close()
	if writeErr != nil {
		return writeErr
	}
	return closeErr
}

func benchmarkHistory(all *database.AllDBForState, parent *types.Block, start uint64) error {
	if !bytes.Equal(all.StateDB.GetCode(params.HistoryStorageAddress), params.HistoryStorageCode) {
		return fmt.Errorf("expected post-Prague history storage contract")
	}
	h := map[uint64]common.Hash{}
	for n := start - 257; n < start-1; n++ {
		v := all.StateDB.GetState(params.HistoryStorageAddress, common.BigToHash(new(big.Int).SetUint64(n%params.HistoryServeWindow)))
		if v == (common.Hash{}) {
			return fmt.Errorf("missing historical hash %d", n)
		}
		h[n] = v
	}
	if err := all.StateDB.Error(); err != nil {
		return err
	}
	if h[start-2] != parent.ParentHash() {
		return fmt.Errorf("historical parent mismatch")
	}
	core.ReplayHistoricalHashes = h
	return nil
}
