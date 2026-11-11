package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/big"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"sort"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/holiman/uint256"
	"go-eth1.16.5-evm/config"
	"go-eth1.16.5-evm/core"
	"go-eth1.16.5-evm/core/state"
	"go-eth1.16.5-evm/core/tracing"
	"go-eth1.16.5-evm/core/types"
	"go-eth1.16.5-evm/core/vm"
)

type studyCall struct {
	f     *componentCall
	block *types.Block
	tx    *types.Transaction
	codes map[common.Hash]bool // actual native execution; collected only in untimed validation
}
type studyPoint struct {
	block    uint64
	seconds  float64
	snapshot *vm.ExperimentReadySnapshot
}
type readinessStudy struct {
	denseCache    vm.ExperimentSnapshotCache
	fixedWorkload bool
	wanted        map[common.Hash]bool
	codes         map[common.Hash]bool
	calls         []*studyCall
	points        []studyPoint
	current       *studyCall
	tx            *types.Transaction
	captureNS     int64
	delegated     int
}

func newReadinessStudy() (*readinessStudy, error) {
	path := os.Getenv("READINESS_TARGETS")
	if path == "" {
		return nil, nil
	}
	raw, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var rows []struct{ Hash common.Hash }
	if err = json.Unmarshal(raw, &rows); err != nil {
		return nil, err
	}
	q := &readinessStudy{wanted: map[common.Hash]bool{}, codes: map[common.Hash]bool{}}
	for _, r := range rows {
		q.wanted[r.Hash] = true
	}
	// Frozen fixture code identities only select what is retained in snapshots;
	// they do not seed, observe, bind or compile any online plan.
	if path := os.Getenv("PAPER_PROBE_CODE_HASHES"); path != "" {
		data, err := os.ReadFile(path)
		if err != nil {
			return nil, err
		}
		var hashes []common.Hash
		if err = json.Unmarshal(data, &hashes); err != nil {
			return nil, err
		}
		for _, h := range hashes {
			q.codes[h] = true
		}
	}
	return q, nil
}
func (q *readinessStudy) capture(c vm.Config, s *state.StateDB, b *types.Block) (vm.Config, *core.TransactionObserver) {
	if q == nil || len(q.calls) == len(q.wanted) {
		return c, &core.TransactionObserver{}
	}
	c.Tracer = &tracing.Hooks{
		OnEnter: func(depth int, typ byte, from, to common.Address, input []byte, gas uint64, value *big.Int) {
			if q.tx == nil {
				return
			}
			started := time.Now()
			defer func() { q.captureNS += time.Since(started).Nanoseconds() }()
			if depth == 0 && vm.OpCode(typ) == vm.CALL {
				q.current = &studyCall{block: b, tx: q.tx, f: &componentCall{pre: s.Copy(), from: from, to: to, input: bytes.Clone(input), value: new(big.Int).Set(value), gas: gas}}
			}
			code := s.GetCode(to)
			if target, ok := types.ParseDelegation(code); ok {
				q.delegated++
				code = s.GetCode(target)
			}
			if vm.OpCode(typ) == vm.CREATE || vm.OpCode(typ) == vm.CREATE2 {
				code = input
			}
			if len(code) > 0 {
				q.codes[crypto.Keccak256Hash(code)] = true
			}
		},
		OnExit: func(depth int, out []byte, used uint64, err error, reverted bool) {
			if depth != 0 || q.current == nil {
				return
			}
			started := time.Now()
			defer func() { q.captureNS += time.Since(started).Nanoseconds() }()
			f := q.current.f
			f.output = bytes.Clone(out)
			f.left = f.gas - used
			f.root = s.Copy().IntermediateRoot(true)
			f.refund = s.GetRefund()
			f.logs = componentLogs(s)
			if err != nil {
				f.failure = err.Error()
			}
		},
	}
	obs := &core.TransactionObserver{Before: func(_ int, tx *types.Transaction) {
		q.tx = nil
		q.current = nil
		if q.wanted[tx.Hash()] {
			q.tx = tx
		}
	}, After: func(_ int, _ *types.Transaction, _ *types.Receipt, _ time.Duration) {
		if q.current != nil {
			q.calls = append(q.calls, q.current)
			q.wanted[q.current.tx.Hash()] = false
		}
		q.current = nil
		q.tx = nil
	}}
	return c, obs
}
func (q *readinessStudy) checkpoint(p *vm.BackgroundPreparation, n, final uint64, seconds float64) {
	if q == nil {
		return
	}
	dense := os.Getenv("PAPER_FIXED_DENSE") == "1"
	if dense {
		if n != final && !((n <= 100 && n%10 == 0) || (n > 100 && n%50 == 0)) {
			return
		}
	} else if os.Getenv("PAPER_FIXED_FINAL_ONLY") == "1" {
		if n != final {
			return
		}
	} else if os.Getenv("PAPER_STREAM") == "1" {
		if n != 0 && n != 50 && n != 100 && n != 200 && n != 400 && n != 800 && n != 1000 {
			return
		}
	} else if n != 1 && n != 5 && n%10 != 0 {
		return
	}
	p.EnterForeground()
	started := time.Now()
	var cache *vm.ExperimentSnapshotCache
	if dense {
		cache = &q.denseCache
	}
	s := p.ExperimentCaptureCached(q.codes, cache)
	q.captureNS += time.Since(started).Nanoseconds()
	p.LeaveForeground()
	q.points = append(q.points, studyPoint{n, seconds, s})
}
func (q *readinessStudy) finish(out string, headers core.ChainContext, cfg vm.Config, sr, cr *benchmarkReadCache) error {
	if q == nil {
		return nil
	}
	// Time one frozen checkpoint at a time: one state per method and input,
	// rather than keeping thousands of additional runner states alive.
	if os.Getenv("PAPER_FIXED_DENSE") == "1" && len(q.points) > 1 {
		q.denseCache = vm.ExperimentSnapshotCache{}
		for i, point := range q.points {
			sub := *q
			sub.points = []studyPoint{point}
			dir := filepath.Join(out, fmt.Sprintf("checkpoint-%d", point.block))
			if err := os.Mkdir(dir, 0755); err != nil {
				return err
			}
			if err := sub.finish(dir, headers, cfg, sr, cr); err != nil {
				return err
			}
			fmt.Printf("measured frozen checkpoint %d/%d block=%d calls=%d\n", i+1, len(q.points), point.block, len(q.calls))
		}
		return writeJSON(filepath.Join(out, "readiness-protocol.json"), map[string]any{"checkpoints": len(q.points), "calls": len(q.calls), "rounds": 9, "iterations": 8, "onlineTrainingByProbe": false, "timedPreparation": false, "frozenReadyViews": true, "scope": "paired Vanilla and Hepha medians independently measured for every immutable checkpoint after the full chronological replay"})
	}
	restoreHistory, err := regionCostFreezeHistory()
	if err != nil {
		return err
	}
	defer restoreHistory()
	if len(q.calls) != len(q.wanted) {
		return fmt.Errorf("readiness fixtures %d/%d", len(q.calls), len(q.wanted))
	}
	type arm struct {
		name    string
		cfg     vm.Config
		seconds float64
		block   uint64
	}
	native := config.DefaultVmConfig
	native.Tracer = nil
	native.Pruning = nil
	native.Arithmetic = nil
	native.ABIScalar = nil
	native.HotRegions = nil
	native.LoopTiming = new(vm.InterpreterLoopTiming)
	arms := []arm{{name: "native", cfg: native}}
	for _, p := range q.points {
		unit := "block"
		if q.fixedWorkload {
			unit = "pass"
		}
		arms = append(arms, arm{name: fmt.Sprintf("%s-%d", unit, p.block), cfg: p.snapshot.Config(cfg), seconds: p.seconds, block: p.block})
	}

	type runner struct {
		a arm
		c *studyCall
		s *state.StateDB
		e *vm.EVM
	}
	var runners []*runner
	for _, a := range arms {
		for _, c := range q.calls {
			a.cfg = vm.ExperimentRootConfig(a.cfg, &c.f.to)
			s := c.f.pre.Copy()
			ctx := core.NewEVMBlockContext(c.block.Header(), headers, nil)
			e := vm.NewEVM(ctx, s, config.MainnetChainConfig, a.cfg)
			msg, err := core.TransactionToMessage(c.tx, types.MakeSigner(config.MainnetChainConfig, c.block.Number(), c.block.Time()), c.block.BaseFee())
			if err != nil {
				return err
			}
			e.SetTxContext(core.NewEVMTxContext(msg))
			runners = append(runners, &runner{a, c, s, e})
		}
	}
	freshDefault := os.Getenv("COLD_FIXED_REUSE") != "1"
	compareFresh := os.Getenv("PAPER_FIXED_COMPARE_FRESH") == "1"
	if compareFresh && freshDefault {
		return fmt.Errorf("fresh-instance comparison requires reused EVM as the main arm")
	}
	run := func(r *runner, n int, validate, fresh bool) (int64, error) {
		*r.a.cfg.LoopTiming = vm.InterpreterLoopTiming{}
		f := r.c.f
		for j := 0; j < n; j++ {
			// Isolate fresh-instance and repeated-call cache lifecycles.
			e := r.e
			if fresh {
				e = vm.NewEVM(r.e.Context, r.s, config.MainnetChainConfig, r.e.Config)
				e.SetTxContext(r.e.TxContext)
			}
			snap := r.s.Snapshot()
			output, left, err := e.Call(f.from, f.to, f.input, f.gas, uint256.MustFromBig(f.value))
			failure := ""
			if err != nil {
				failure = err.Error()
			}
			if left != f.left || failure != f.failure || !bytes.Equal(output, f.output) {
				return 0, fmt.Errorf("readiness call mismatch %s/%s", r.a.name, r.c.tx.Hash())
			}
			if validate {
				if r.s.Copy().IntermediateRoot(true) != f.root || r.s.GetRefund() != f.refund {
					return 0, fmt.Errorf("readiness state mismatch %s/%s", r.a.name, r.c.tx.Hash())
				}
				if logs := componentLogs(r.s); !reflect.DeepEqual(logs, f.logs) {
					return 0, fmt.Errorf("readiness logs mismatch")
				}
			}
			r.s.RevertToSnapshot(snap)
		}
		return r.a.cfg.LoopTiming.Nanoseconds / int64(n), nil
	}
	// Independently check actual executing bytecode against the capture set.
	actual := map[common.Hash]bool{}
	for _, r := range runners {
		if r.a.name == "native" {
			r.c.codes = map[common.Hash]bool{}
			seen := map[*vm.Contract]bool{}
			r.e.Config.Tracer = &tracing.Hooks{OnOpcode: func(_ uint64, _ byte, _ uint64, _ uint64, scope tracing.OpContext, _ []byte, _ int, _ error) {
				if c, ok := scope.(*vm.ScopeContext); ok && !seen[c.Contract] {
					seen[c.Contract] = true
					hash := crypto.Keccak256Hash(c.Contract.Code)
					actual[hash] = true
					r.c.codes[hash] = true
				}
			}}
			if _, err := run(r, 1, true, freshDefault); err != nil {
				return err
			}
			r.e.Config.Tracer = nil
		}
	}
	for hash := range actual {
		if !q.codes[hash] {
			return fmt.Errorf("snapshot missing actual code hash %s", hash)
		}
	}
	f, err := os.Create(filepath.Join(out, "same-call-readiness.jsonl"))
	if err != nil {
		return err
	}
	defer f.Close()
	enc := json.NewEncoder(f)
	// Untimed native state warming and cache lookup only; frozen plans have no
	// observation callbacks and cannot generate work. Then rotate all arms.
	for _, r := range runners {
		if _, err := run(r, 1, true, freshDefault); err != nil {
			return err
		}
	}
	if compareFresh {
		for _, r := range runners {
			if _, err := run(r, 1, true, true); err != nil {
				return err
			}
		}
	}
	forcedGC := os.Getenv("PAPER_FIXED_PRE_GC") == "1"
	if forcedGC {
		runtime.GC() // Same untimed collection used by the paper warm harness.
	}
	var memoryBefore, memoryAfter runtime.MemStats
	if os.Getenv("PAPER_FIXED_MEMORY_COMPARE") == "1" {
		runtime.ReadMemStats(&memoryBefore)
	}
	resourcesBefore := vm.FlowPreparationResources()
	sd, sb := sr.counts()
	cd, cb := cr.counts()
	sr.setStrict(true)
	cr.setStrict(true)
	// The companion mode isolates immediate CPU-cache warming with the SAME
	// frozen plans, calls and timer. Rotate its order; it is not online training.
	encoders := []*json.Encoder{enc}
	warmModes := []bool{os.Getenv("PAPER_FIXED_BATCH_PREWARM") == "1"}
	if os.Getenv("PAPER_FIXED_COMPARE_PREWARM") == "1" {
		companion, err := os.Create(filepath.Join(out, "same-call-readiness-batchwarm.jsonl"))
		if err != nil {
			return err
		}
		defer companion.Close()
		encoders = append(encoders, json.NewEncoder(companion))
		warmModes = []bool{false, true}
	}
	freshMode := -1
	if compareFresh {
		companion, err := os.Create(filepath.Join(out, "same-call-readiness-fresh.jsonl"))
		if err != nil {
			return err
		}
		defer companion.Close()
		freshMode = len(encoders)
		encoders = append(encoders, json.NewEncoder(companion))
		warmModes = append(warmModes, warmModes[0])
	}
	for round := 0; round < 9; round++ {
		for mi := range encoders {
			mode := (mi + round) % len(encoders)
			for j := range runners {
				r := runners[(j+round)%len(runners)]
				if warmModes[mode] {
					if _, err := run(r, 1, false, freshDefault || mode == freshMode); err != nil {
						return err
					}
				}
				lat, err := run(r, 8, false, freshDefault || mode == freshMode)
				if err != nil {
					return err
				}
				row := map[string]any{"freshEVMPerCall": freshDefault || mode == freshMode, "batchPrewarm": warmModes[mode], "condition": r.a.name, "checkpointBlock": r.a.block, "onlineSeconds": r.a.seconds, "round": round, "hash": r.c.tx.Hash(), "to": r.c.f.to, "latNS": lat}
				if q.fixedWorkload {
					delete(row, "checkpointBlock")
					row["checkpointPass"] = r.a.block
				}
				if err = encoders[mode].Encode(row); err != nil {
					return err
				}
			}
		}
	}
	if os.Getenv("PAPER_FIXED_MEMORY_COMPARE") == "1" {
		runtime.ReadMemStats(&memoryAfter)
		if err := writeJSON(filepath.Join(out, "readiness-memory.json"), map[string]any{
			"residentCheckpoints": len(q.points), "runners": len(runners), "forcedPreGC": forcedGC,
			"heapBefore": memoryBefore.HeapAlloc, "heapAfter": memoryAfter.HeapAlloc,
			"heapObjectsBefore": memoryBefore.HeapObjects, "heapObjectsAfter": memoryAfter.HeapObjects,
			"gcCycles":       memoryAfter.NumGC - memoryBefore.NumGC,
			"gcPauseNS":      memoryAfter.PauseTotalNs - memoryBefore.PauseTotalNs,
			"allocatedBytes": memoryAfter.TotalAlloc - memoryBefore.TotalAlloc,
			"scope":          "Whole frozen measurement loop, all methods; includes untimed prewarm and result encoding. GC enabled normally. Memory counters are outside individual timed batches, not per-method attribution.",
		}); err != nil {
			return err
		}
	}
	resourcesAfter := vm.FlowPreparationResources()
	if resourcesBefore["pending"] != 0 || resourcesAfter["pending"] != 0 || resourcesBefore["completed"] != resourcesAfter["completed"] {
		return fmt.Errorf("readiness timing admitted preparation")
	}
	sr.setStrict(false)
	cr.setStrict(false)
	sd1, sb1 := sr.counts()
	cd1, cb1 := cr.counts()
	if sd != sd1 || sb != sb1 || cd != cd1 || cb != cb1 {
		return fmt.Errorf("readiness timed database access")
	}
	var demandAudit []map[string]any
	for _, r := range runners {
		if os.Getenv("PAPER_DEMAND_AUDIT") == "1" && r.a.block == q.points[len(q.points)-1].block {
			r.e.Config.RegionExecutionStats = &vm.RegionExecutionStats{DemandCapture: make(map[string]*vm.DemandExecutionVisit)}
		}
		if _, err := run(r, 1, true, freshDefault); err != nil {
			return err
		}
		if stats := r.e.Config.RegionExecutionStats; stats != nil && stats.DemandCapture != nil {
			demandAudit = append(demandAudit, map[string]any{"hash": r.c.tx.Hash(), "condition": r.a.name, "stats": stats})
			r.e.Config.RegionExecutionStats = nil
		}
	}
	if len(demandAudit) > 0 {
		if err = writeJSON(filepath.Join(out, "readiness-demand-audit.json"), demandAudit); err != nil {
			return err
		}
	}
	var codePoints []map[string]any
	for _, p := range q.points {
		row := map[string]any{"block": p.block, "onlineSeconds": p.seconds, "codes": p.snapshot.CodeReadiness()}
		if q.fixedWorkload {
			delete(row, "block")
			row["pass"] = p.block
		}
		codePoints = append(codePoints, row)
	}
	var callCodes []map[string]any
	for _, c := range q.calls {
		var hashes []string
		for h := range c.codes {
			hashes = append(hashes, h.Hex())
		}
		sort.Strings(hashes)
		callCodes = append(callCodes, map[string]any{"hash": c.tx.Hash(), "codes": hashes})
	}
	if err = writeJSON(filepath.Join(out, "readiness-code-status.json"), map[string]any{"calls": callCodes, "checkpoints": codePoints, "scope": "Untimed frozen delivery metadata for actually executed native code. Ready entry counts are not dynamic instruction coverage. No probe training."}); err != nil {
		return err
	}
	return writeJSON(filepath.Join(out, "readiness-protocol.json"), map[string]any{"companionFreshEVM": compareFresh, "companionBatchPrewarm": os.Getenv("PAPER_FIXED_COMPARE_PREWARM") == "1", "batchPrewarm": os.Getenv("PAPER_FIXED_BATCH_PREWARM") == "1", "untimedDemandAudit": len(demandAudit) > 0, "calls": len(q.calls), "delegatedFrames": q.delegated, "actualCodeHashesVerified": len(actual), "timedPreparation": false, "checkpoints": len(q.points), "captureNS": q.captureNS, "rounds": 9, "iterations": 8, "onlineTrainingByProbe": false, "preparationUsesSameFixedCalls": q.fixedWorkload, "frozenReadyViews": true, "freshEVMPerCall": os.Getenv("COLD_FIXED_REUSE") != "1", "warmReferencePreparedAfterReplay": false, "scope": "Fixed original calls with identical input/state/context. Frozen observations removed; lat measured with no background work. This diagnoses readiness, not online foreground interference. State, gas, output, failure and logs validated; executing code hashes independently verified. If preparationUsesSameFixedCalls is true, preparation uses repeated fixed workload passes, NOT a chronological block stream."})
}
