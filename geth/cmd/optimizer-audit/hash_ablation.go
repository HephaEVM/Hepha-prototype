package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common"

	"github.com/holiman/uint256"
	"go-eth1.16.5-evm/config"
	"go-eth1.16.5-evm/core"
	"go-eth1.16.5-evm/core/state"
	"go-eth1.16.5-evm/core/tracing"
	"go-eth1.16.5-evm/core/types"
	"go-eth1.16.5-evm/core/vm"
	"go-eth1.16.5-evm/database"
	"go-eth1.16.5-evm/ethdb"
	"math/big"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"time"
)

type hashTarget struct {
	Hash  string `json:"hash"`
	Block uint64 `json:"block"`
	Index int    `json:"index"`
}
type hashArm struct {
	name    string
	cfg     vm.Config
	engine  *vm.ArithmeticEngine
	abi     *vm.ABIScalarDiagnostic
	prep    vm.HashPreparationStats
	compute vm.ComputeHelperStats
}
type hashObservation struct {
	Interpreted    uint64 `json:"interpreted_instructions"`
	Analyses       uint64 `json:"analyses"`
	AnalysisNS     uint64 `json:"analysis_ns"`
	AnalyzedBytes  uint64 `json:"analyzed_bytes"`
	FrameNS        int64  `json:"frame_ns"`
	ExternalFrames uint64 `json:"external_frames"`
	HostCallbacks  uint64 `json:"host_callbacks"`

	Address     vm.RegionExecutionStats `json:"address"`
	Direct      uint64                  `json:"direct"`
	EntryChoice uint8                   `json:"entry_choice"`
	RootChecked bool                    `json:"root_checked"`
	GCCycles    uint32                  `json:"gc_cycles"`
	GCPauseNS   uint64                  `json:"gc_pause_ns"`
	AllocBytes  uint64                  `json:"alloc_bytes"`
	HeapBytes   uint64                  `json:"heap_bytes"`

	Hash               string                  `json:"hash"`
	Block              uint64                  `json:"block"`
	Mode               string                  `json:"mode"`
	Phase              string                  `json:"phase"`
	Repeat             int                     `json:"repeat"`
	Order              int                     `json:"order"`
	Lat                int64                   `json:"lat_ns"`
	Elapsed            int64                   `json:"elapsed_ns"`
	ThreadCPU          int64                   `json:"thread_cpu_ns"`
	ProcessCPU         int64                   `json:"process_cpu_ns"`
	BatchCPU           int64                   `json:"batch_cpu_ns,omitempty"`
	BatchElapsed       int64                   `json:"batch_elapsed_ns,omitempty"`
	BatchIterations    int                     `json:"batch_iterations,omitempty"`
	Gas                uint64                  `json:"gas"`
	Status             uint64                  `json:"status"`
	Root               common.Hash             `json:"root"`
	Arithmetic         vm.ArithmeticStats      `json:"arithmetic"`
	ABI                vm.ABIScalarStats       `json:"abi"`
	Prep               vm.HashPreparationStats `json:"prep"`
	Compute            vm.ComputeHelperStats   `json:"compute"`
	BusinessEquivalent bool                    `json:"business_equivalent"`
	ModifiedExecuted   bool                    `json:"modified_executed"`
	Hot                vm.HotRegionStats       `json:"hot"`
}

func hashAblation(db ethdb.Database, snapshot string, start, end uint64, out string) error {
	if start != 23400001 {
		return fmt.Errorf("unexpected anchor")
	}
	manifestPath := filepath.Join(out, "manifest.json")
	if _, err := os.Stat(manifestPath); err == nil {
		return fmt.Errorf("refusing overwrite")
	}
	runtime.GOMAXPROCS(benchmarkProcs)
	modes := []string{"vanilla-native", "hepha"}

	rounds := 32
	if v := os.Getenv("HASH_ROUNDS"); v != "" {
		var err error
		rounds, err = strconv.Atoi(v)
		if err != nil || rounds < 1 {
			return fmt.Errorf("bad rounds")
		}
	}
	offset := 0
	if os.Getenv("HASH_ORDER_OFFSET") != "" {
		var e error
		offset, e = strconv.Atoi(os.Getenv("HASH_ORDER_OFFSET"))
		if e != nil || offset < 0 {
			return fmt.Errorf("bad order offset")
		}
	}
	raw, err := os.ReadFile(os.Getenv("HASH_TARGETS"))
	if err != nil {
		return err
	}
	var targets []hashTarget
	if err = json.Unmarshal(raw, &targets); err != nil {
		return err
	}
	wanted := map[string]hashTarget{}
	for _, r := range targets {
		if r.Block <= end {
			wanted[r.Hash] = r
		}
	}
	if len(wanted) == 0 {
		return fmt.Errorf("no targets")
	}
	manifest := map[string]any{"status": "running", "modes": modes, "rounds": rounds, "targets": wanted, "iterations": 128, "scope": "Paired original historical root calls, Vanilla geth versus Hepha Full. Preparation converges before timing. Original gas and full output/state validation. This is fixed-call execution, not full-block wall time.", "start": start, "end": end, "orderOffset": offset, "immutableArmSelection": true, "databaseCacheMiB": 512, "databaseHandles": 1024}

	if raw := os.Getenv("HASH_ITERATIONS"); raw != "" {
		n, e := strconv.Atoi(raw)
		if e != nil || n < 1 {
			return fmt.Errorf("invalid HASH_ITERATIONS")
		}
		manifest["iterations"] = n
	}

	if err = writeJSON(manifestPath, manifest); err != nil {
		return err
	}
	chain := &benchmarkChainDB{Database: db, reads: newBenchmarkReadCache(db)}
	consensus, err := config.CreateConsensusEngine(config.MainnetChainConfig)
	if err != nil {
		return err
	}
	headers, err := core.NewHeaderChain(chain, config.MainnetChainConfig, consensus)
	if err != nil {
		return err
	}
	processor := core.NewStateProcessor(headers)
	parent, err := database.GetBlockByNumber(chain, new(big.Int).SetUint64(start-1))
	if err != nil {
		return err
	}
	overlay, err := openOverlay(snapshot)
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

	library, err := os.ReadFile(arithmeticTemplatesPath)
	if err != nil {
		return err
	}
	arms := make([]*hashArm, len(modes))
	policyArms := map[string]*hashArm{}
	defer func() {
		for _, p := range libraryPreparers {
			p.Close()
		}
	}()
	for i, name := range modes {
		a := &hashArm{name: name}
		arms[i] = a
		a.cfg = config.DefaultVmConfig

		a.cfg.Pruning = nil
		a.cfg.Arithmetic = nil
		a.cfg.ABIScalar = nil
		a.cfg.Tracer = nil

		a.cfg.LoopTiming = new(vm.InterpreterLoopTiming)
		if os.Getenv("PAPER_FRAME_TIMING") == "1" {
			a.cfg.FrameTiming = new(vm.InterpreterLoopTiming)
		}

		if strings.Contains(name, "native") {
			continue
		}
		group := strings.TrimSuffix(name, "-aa")
		if prior, ok := policyArms[group]; ok {
			*a = *prior
			a.name = name
			continue
		}
		policyArms[group] = a

		a.engine, err = vm.LoadArithmeticEngine(bytes.NewReader(library), vm.ArithmeticEngineOptions{CacheEntries: 16384})
		if err != nil {
			return err
		}
		a.abi, err = vm.LoadABIRegionDiagnostic(os.Getenv("ABI_SCALAR_TEMPLATES"), os.Getenv("ABI_EXPANDED_TEMPLATES"))
		if err != nil {
			return err
		}
		a.abi.Optimize = true
		a.cfg.RegionExecutionStats = new(vm.RegionExecutionStats)
		a.cfg = a.cfg.WithPruning(vm.PruningConfig{Engine: a.engine, ABI: a.abi, ComputeStats: &a.compute, HashStats: &a.prep})
		if os.Getenv("PAPER_WARM") == "1" {
			a.cfg.RegionExecutionStats = nil
			a.cfg = a.cfg.WithPruning(vm.PruningConfig{Engine: a.engine, ABI: a.abi})
		}
		a.cfg.HotRegions, err = vm.NewHotRegionRuntime(a.engine, os.Getenv("HOT_REGION_CORPUS"), os.Getenv("HOT_REGION_SELECTED"))
		if err != nil {
			return err
		}

		if os.Getenv("REGION_POLICY") == "new" {
			p, er := vm.NewBackgroundPreparation(a.engine, a.abi, 4096)
			if er != nil {
				return er
			}
			a.cfg.Pruning.Preparation = p
			libraryPreparers = append(libraryPreparers, p)
		}
	}
	file, err := os.Create(filepath.Join(out, "observations.jsonl"))
	if err != nil {
		return err
	}
	defer file.Close()
	enc := json.NewEncoder(file)
	canonical := config.DefaultVmConfig
	canonical.Pruning = nil
	canonical.Arithmetic = nil
	canonical.ABIScalar = nil
	canonical.Tracer = nil

	root := parent.Root()
	parentHash := parent.Hash()
	completed := 0
	checks := 0
	var blockChecks []map[string]any
	for number := start; number <= end; number++ {
		if err = all.UpdateStateDB(root); err != nil {
			return err
		}
		block, e := database.GetBlockByNumber(chain, new(big.Int).SetUint64(number))
		if e != nil {
			return e
		}
		if block.ParentHash() != parentHash {
			return fmt.Errorf("parent mismatch")
		}
		st := all.Copy()
		var pre *state.StateDB
		var prefix uint64
		var problem error
		var fixture *componentCall
		var active bool
		traced := canonical
		traced.Tracer = &tracing.Hooks{
			OnOpcode: func(pc uint64, op byte, gas, cost uint64, scope tracing.OpContext, ret []byte, depth int, err error) {
				if active && fixture != nil {
					if s, ok := scope.(*vm.ScopeContext); ok && len(s.Contract.Code) > 0 {
						fixture.codes[s.Contract.CodeHash.Hex()] = true
					}
				}
			},
			OnEnter: func(depth int, typ byte, from, to common.Address, input []byte, gas uint64, value *big.Int) {
				if active && depth == 0 {
					if vm.OpCode(typ) != vm.CALL {
						problem = fmt.Errorf("non-CALL target")
						return
					}
					fixture = &componentCall{pre: st.Copy(), from: from, to: to, input: bytes.Clone(input), value: new(big.Int).Set(value), gas: gas, codes: map[string]bool{}, executed: map[common.Address]bool{}}
				}
				if active && fixture != nil {
					fixture.executed[to] = true
				}
			},
			OnExit: func(depth int, output []byte, gasUsed uint64, err error, reverted bool) {
				if active && depth == 0 && fixture != nil {
					fixture.output = bytes.Clone(output)
					fixture.left = fixture.gas - gasUsed
					fixture.root = st.Copy().IntermediateRoot(true)
					fixture.refund = st.GetRefund()
					fixture.logs = componentLogs(st)
					if err != nil {
						fixture.failure = err.Error()
					}
				}
			},
		}
		observer := &core.TransactionObserver{
			Before: func(index int, tx *types.Transaction) {
				_, active = wanted[tx.Hash().Hex()]
				fixture = nil
				if _, ok := wanted[tx.Hash().Hex()]; ok && problem == nil {
					pre = st.Copy()
				}
			},
			After: func(index int, tx *types.Transaction, receipt *types.Receipt, _ time.Duration) {
				if target, ok := wanted[tx.Hash().Hex()]; ok && problem == nil {
					if target.Index != index || target.Block != number {
						problem = fmt.Errorf("target identity mismatch")
						return
					}
					expected := st.Copy().IntermediateRoot(true)
					problem = hashRunTarget(block, headers, pre, tx, index, prefix, receipt, expected, arms, rounds, offset, overlay.reads, chain.reads, enc, &checks, fixture)
					pre = nil
					if problem == nil {
						completed++
						fmt.Printf("target block=%d index=%d hash=%s completed=%d/%d checks=%d\n", number, index, tx.Hash(), completed, len(wanted), checks)
					}
				}
				prefix = receipt.CumulativeGasUsed
			},
		}
		targetBlock := false
		for _, target := range wanted {
			if target.Block == number {
				targetBlock = true
				break
			}
		}
		if !targetBlock {
			traced.Tracer = nil
			observer = nil
		}
		result, e := processor.ProcessWithObserver(block, st, traced, observer)
		if e != nil {
			return e
		}
		if problem != nil {
			return problem
		}
		if _, _, e = benchmarkValidate(block, st, result); e != nil {
			return e
		}
		root, _, e = st.CommitWithUpdate(number, true, true)
		if e != nil {
			return e
		}
		if root != block.Root() {
			return fmt.Errorf("commit mismatch")
		}
		parentHash = block.Hash()
		blockChecks = append(blockChecks, map[string]any{"block": number, "root": root, "gas": result.GasUsed, "transactions": len(block.Transactions())})
		if number%10 == 0 {
			fmt.Printf("canonical reconstruction through %d targets=%d/%d\n", number, completed, len(wanted))
		}
	}
	if completed != len(wanted) {
		return fmt.Errorf("missing targets %d/%d", completed, len(wanted))
	}
	if err = writeJSON(filepath.Join(out, "block-validation.json"), blockChecks); err != nil {
		return err
	}
	if err = runLibraryCases(out, enc, rounds); err != nil {
		return err
	}
	manifest["status"] = "complete"
	manifest["validatedReplays"] = checks
	manifest["completedTargets"] = completed
	return writeJSON(manifestPath, manifest)
}

// Replaces hashRunTarget in the isolated command overlay.
type componentCall struct {
	pre               *state.StateDB
	codes             map[string]bool
	executed          map[common.Address]bool
	from, to          common.Address
	input, output     []byte
	value             *big.Int
	gas, left, refund uint64
	failure           string
	root              common.Hash
	logs              []*types.Log
}

func hashRunTarget(block *types.Block, headers core.ChainContext, _ *state.StateDB, tx *types.Transaction, index int, _ uint64, _ *types.Receipt, _ common.Hash, arms []*hashArm, rounds, offset int, sr, cr *benchmarkReadCache, enc *json.Encoder, checks *int, fixture *componentCall) error {
	if fixture == nil || fixture.pre == nil {
		return fmt.Errorf("missing root call fixture")
	}
	msg, err := core.TransactionToMessage(tx, types.MakeSigner(config.MainnetChainConfig, block.Number(), block.Time()), block.BaseFee())
	if err != nil {
		return err
	}
	checked, audits, equivalent, err := checkedFixture(block, headers, tx, fixture)
	if err != nil {
		return err
	}

	modifiedExecuted := false
	for _, a := range audits {
		modifiedExecuted = modifiedExecuted || (a.Applied && a.Executed)
	}
	if err = enc.Encode(map[string]any{"codeHashes": fixture.codes, "phase": "eligibility", "hash": tx.Hash().Hex(), "block": block.NumberU64(), "to": fixture.to.Hex(), "patches": audits, "business_equivalent": equivalent, "modified_executed": modifiedExecuted, "unchecked_gas": fixture.gas - fixture.left, "checked_gas": checked.gas - checked.left, "unchecked_error": fixture.failure, "checked_error": checked.failure}); err != nil {
		return err
	}
	iterations := 128
	if raw := os.Getenv("HASH_ITERATIONS"); raw != "" {
		iterations, err = strconv.Atoi(raw)
		if err != nil || iterations < 1 {
			return fmt.Errorf("invalid HASH_ITERATIONS")
		}
	}
	type runner struct {
		a       *hashArm
		s       *state.StateDB
		e       *vm.EVM
		name    string
		choice  uint8
		fixture *componentCall
	}
	runners := make([]runner, len(arms))
	sharedState := fixture.pre.Copy()
	checkedState := checked.pre.Copy()
	same := map[string]runner{}
	for i, a := range arms {
		group := strings.TrimSuffix(a.name, "-aa")
		choice := uint8(0)
		if v, ok := same[group]; ok {
			v.name = a.name
			v.choice = choice
			runners[i] = v
			continue
		}
		s := sharedState
		expected := fixture
		if isCheckedArm(a.name) {
			s = checkedState
			expected = checked
		}
		ctx := core.NewEVMBlockContext(block.Header(), headers, nil)
		e := vm.NewEVM(ctx, s, config.MainnetChainConfig, a.cfg)
		e.Config.Arithmetic = a.cfg.Arithmetic
		e.SetTxContext(core.NewEVMTxContext(msg))

		runners[i] = runner{a, s, e, a.name, choice, expected}
		same[group] = runners[i]
	}
	run := func(v *runner, phase string, rep, order, n int, strict bool) (hashObservation, error) {
		a, s, e := v.a, v.s, v.e
		if phase == "postflight" && os.Getenv("PAPER_DEMAND_AUDIT") == "1" {
			previous := a.cfg.RegionExecutionStats
			a.cfg.RegionExecutionStats = &vm.RegionExecutionStats{}
			e.Config.RegionExecutionStats = a.cfg.RegionExecutionStats
			defer func() { a.cfg.RegionExecutionStats = previous; e.Config.RegionExecutionStats = previous }()
		}
		if a.cfg.RegionExecutionStats != nil {
			*a.cfg.RegionExecutionStats = vm.RegionExecutionStats{}
			if phase == "postflight" && os.Getenv("PAPER_DEMAND_AUDIT") == "1" {
				a.cfg.RegionExecutionStats.DemandCapture = make(map[string]*vm.DemandExecutionVisit)
				a.cfg.RegionExecutionStats.Capture = make(map[string]*vm.RegionExecutionVisit)
			}
			if phase == "warmup" && !strings.HasSuffix(v.name, "-aa") {
				a.cfg.RegionExecutionStats.Capture = make(map[string]*vm.RegionExecutionVisit)
			}
		}

		fixture := v.fixture
		*a.cfg.LoopTiming = vm.InterpreterLoopTiming{}
		if a.cfg.FrameTiming != nil {
			*a.cfg.FrameTiming = vm.InterpreterLoopTiming{}
		}
		var before vm.ArithmeticStats
		var abiBefore vm.ABIScalarStats
		var hotBefore vm.HotRegionStats

		if a.engine != nil {
			before = a.engine.Stats()
			abiBefore = a.abi.Stats
			if a.cfg.HotRegions != nil {
				hotBefore = a.cfg.HotRegions.Stats
			}
		}

		prepBefore, computeBefore := a.prep, a.compute
		sd, sb := sr.counts()
		cd, cb := cr.counts()
		sr.setStrict(strict)
		cr.setStrict(strict)
		var memBefore, memAfter runtime.MemStats
		if phase != "stream" {
			runtime.ReadMemStats(&memBefore)
		}
		cpuBefore, processBefore := int64(0), int64(0)
		if phase == "cpu-diagnostic" {
			cpuBefore = int64(0)
			var err error
			processBefore, err = benchmarkCPU()
			if err != nil {
				return hashObservation{}, err
			}
		}
		start := time.Now()
		instructionBefore := vm.PaperInterpretedInstructions()
		for _, p := range libraryPreparers {
			p.EnterForeground()
		}
		defer func() {
			for _, p := range libraryPreparers {
				p.LeaveForeground()
			}
		}()
		lastSnapshot := 0
		for j := 0; j < n; j++ {
			snapshot := s.Snapshot()
			lastSnapshot = snapshot
			out, left, err := e.Call(fixture.from, fixture.to, fixture.input, fixture.gas, uint256.MustFromBig(fixture.value))
			failure := ""
			if err != nil {
				failure = err.Error()
			}
			if left != fixture.left || failure != fixture.failure || !bytes.Equal(out, fixture.output) {
				return hashObservation{}, fmt.Errorf("call result mismatch %s %s left=%d/%d error=%s/%s", tx.Hash(), a.name, left, fixture.left, failure, fixture.failure)
			}
			if j+1 < n {
				s.RevertToSnapshot(snapshot)
			}
		}
		elapsed := time.Since(start).Nanoseconds()
		instructionCount := vm.PaperInterpretedInstructions() - instructionBefore
		cpuAfter, processAfter := int64(0), int64(0)
		if phase == "cpu-diagnostic" {
			cpuAfter = int64(0)
			var err error
			processAfter, err = benchmarkCPU()
			if err != nil {
				return hashObservation{}, err
			}
		}
		if phase != "stream" {
			runtime.ReadMemStats(&memAfter)
		}
		sr.setStrict(false)
		cr.setStrict(false)
		sd1, sb1 := sr.counts()
		cd1, cb1 := cr.counts()
		row := hashObservation{EntryChoice: v.choice, BusinessEquivalent: equivalent, ModifiedExecuted: modifiedExecuted, Hash: tx.Hash().Hex(), Block: block.NumberU64(), Mode: v.name, Phase: phase, Repeat: rep, Order: order, Lat: a.cfg.LoopTiming.Nanoseconds / int64(n), Elapsed: elapsed / int64(n), Gas: fixture.gas - fixture.left, Status: 1, Prep: a.prep.Sub(prepBefore), Compute: a.compute.Sub(computeBefore)}

		row.Interpreted = instructionCount / uint64(n)
		if a.cfg.RegionExecutionStats != nil {
			row.Address = *a.cfg.RegionExecutionStats
		}
		if a.cfg.FrameTiming != nil {
			row.FrameNS = a.cfg.FrameTiming.Nanoseconds / int64(n)
		}

		row.ThreadCPU = (cpuAfter - cpuBefore) / int64(n)
		row.ProcessCPU = (processAfter - processBefore) / int64(n)
		if phase == "cpu-diagnostic" {
			row.BatchCPU, row.BatchElapsed, row.BatchIterations = processAfter-processBefore, elapsed, n
		}
		row.GCCycles = memAfter.NumGC - memBefore.NumGC
		row.GCPauseNS = memAfter.PauseTotalNs - memBefore.PauseTotalNs
		row.AllocBytes = memAfter.TotalAlloc - memBefore.TotalAlloc
		row.HeapBytes = memAfter.HeapAlloc
		if fixture.failure != "" {
			row.Status = 0
		}
		if a.engine != nil {
			row.Arithmetic = benchmarkStatsDelta(a.engine.Stats(), before)
			row.ABI = a.abi.Stats.Sub(abiBefore)
			if a.cfg.HotRegions != nil {
				row.Hot = a.cfg.HotRegions.Stats.Sub(hotBefore)
			}
		}
		if strings.HasPrefix(a.name, "hepha-frozen-direct") && (row.Arithmetic.CacheMisses != 0 || row.ABI.Misses != 0) {
			return row, fmt.Errorf("frozen direct dispatch attempted new binding")
		}

		if strict && (sd != sd1 || sb != sb1 || cd != cd1 || cb != cb1 || (phase != "stream" && phase != "stream-check" && (row.Arithmetic.CacheMisses != 0 || row.ABI.Misses != 0)) || row.Arithmetic.ModelMismatches != 0) {
			return row, fmt.Errorf("timed I/O or binding miss %s", a.name)
		}
		// Validate full state before/after all timed repetitions, separately from
		// the timing stream. Every individual Call still checks return/error/gas.
		row.Root = fixture.root
		row.RootChecked = phase == "warmup" || phase == "postflight" || phase == "stream-check"
		if s.Error() != nil {
			return row, s.Error()
		}
		if row.RootChecked {
			row.Root = s.Copy().IntermediateRoot(true)
			if row.Root != fixture.root || s.GetRefund() != fixture.refund || !reflect.DeepEqual(componentLogs(s), fixture.logs) {
				return row, fmt.Errorf("call state mismatch %s %s root=%s/%s refund=%d/%d logsEqual=%t", tx.Hash(), v.name, row.Root, fixture.root, s.GetRefund(), fixture.refund, reflect.DeepEqual(componentLogs(s), fixture.logs))
			}
		}
		s.RevertToSnapshot(lastSnapshot)
		*checks += n
		return row, nil
	}
	// Keep captured input/state fixed across complete corpus sweeps.
	for i := range runners {
		v := &runners[i]
		libraryCases = append(libraryCases, libraryCase{hash: tx.Hash().Hex(), noise: experimentNoise(tx.Hash().Hex()), native: strings.Contains(v.name, "native"), run: func(phase string, rep, n int, strict bool) (hashObservation, error) {
			return run(v, phase, rep, 0, n, strict)
		}})
	}

	return nil
}

func componentLogs(s *state.StateDB) []*types.Log {
	logs := s.Logs()
	sort.Slice(logs, func(i, j int) bool { return logs[i].Index < logs[j].Index })
	out := make([]*types.Log, len(logs))
	for i, l := range logs {
		v := *l
		out[i] = &v
	}
	return out
}

type libraryCase struct {
	hash          string
	noise, native bool
	run           func(string, int, int, bool) (hashObservation, error)
}

var libraryCases []libraryCase
var libraryPreparers []*vm.BackgroundPreparation
var experimentNoiseHashes map[string]bool

func experimentNoise(hash string) bool {
	if os.Getenv("PAPER_WARM") == "1" {
		return false
	}
	if experimentNoiseHashes == nil {
		experimentNoiseHashes = map[string]bool{}
		raw, err := os.ReadFile(os.Getenv("NOISE_TARGETS"))
		if err != nil {
			panic(err)
		}
		var xs []hashTarget
		if err = json.Unmarshal(raw, &xs); err != nil {
			panic(err)
		}
		for _, x := range xs {
			experimentNoiseHashes[x.Hash] = true
		}
	}
	return experimentNoiseHashes[hash]
}
func libraryBackgroundStats() []vm.BackgroundPreparationStats {
	var r []vm.BackgroundPreparationStats
	for _, p := range libraryPreparers {
		r = append(r, p.Stats())
	}
	return r
}
func probeStamp() string {
	return fmt.Sprint(vm.FlowPreparationResources(), libraryBackgroundStats())
}
func probeWait() error {
	for _, p := range libraryPreparers {
		p.Wait()
	}
	events, err := vm.WaitDemand()
	if err != nil {
		return err
	}
	for _, e := range events {
		if e.Error != "" {
			return fmt.Errorf("compile: %s", e.Error)
		}
	}
	return nil
}
func probeSettle(cases []libraryCase) error {
	// No state changes persist between sweeps. New plans/functions must stop
	// appearing for two complete sweeps of the fixed inputs.
	stable := 0
	for pass := 0; pass < 16; pass++ {
		before := vm.FlowPreparationResources()["completed"]
		enqueued := uint64(0)
		for _, p := range libraryPreparers {
			enqueued += p.Stats().Enqueued
		}
		for _, c := range cases {
			if !c.native {
				if _, err := c.run("stream-check", pass, 1, true); err != nil {
					return err
				}
			}
		}
		if err := probeWait(); err != nil {
			return err
		}
		after := uint64(0)
		for _, p := range libraryPreparers {
			after += p.Stats().Enqueued
		}
		if before == vm.FlowPreparationResources()["completed"] && after == enqueued {
			stable++
		} else {
			stable = 0
		}
		fmt.Printf("settle pass=%d stable=%d %s\n", pass, stable, probeStamp())
		if stable >= 2 {
			return nil
		}
	}
	return fmt.Errorf("fixed input preparation did not converge")
}

func runLibraryCases(out string, enc *json.Encoder, rounds int) error {
	return runPaperWarmCases(out, enc, rounds)
}
