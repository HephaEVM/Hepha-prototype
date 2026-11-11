package main

// An explicitly separate fixed-working-set experiment. Native history restores
// fixtures before workers exist. Repeated calls then train ordinary preparation;
// frozen measurement probes never train it. No precompiled cache is imported.
import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/holiman/uint256"
	"go-eth1.16.5-evm/config"
	"go-eth1.16.5-evm/core"
	"go-eth1.16.5-evm/core/state"
	"go-eth1.16.5-evm/core/types"
	"go-eth1.16.5-evm/core/vm"
)

func workingSetCheckpoint(pass, last int) bool {
	return pass == 0 || pass == 1 || pass == 2 || pass == 5 || pass == 10 || pass == 20 || pass == 50 || pass == 100 || pass == 200 || pass == last
}

func (q *readinessStudy) runWorkingSet(out string, headers core.ChainContext, cfg vm.Config, p *vm.BackgroundPreparation, sr, cr *benchmarkReadCache, manifest map[string]any) error {
	if len(q.calls) != len(q.wanted) || !vm.RegionWindowEnabled() {
		return fmt.Errorf("fixed working set requires all fixed inputs and bounded synchronous windows")
	}
	passes, err := strconv.Atoi(os.Getenv("PAPER_WORKING_SET_PASSES"))
	if err != nil || passes < 1 || passes > 400 {
		return fmt.Errorf("PAPER_WORKING_SET_PASSES must be 1..400")
	}
	defer p.Close()
	q.fixedWorkload = true
	// A broad-cohort steady-state diagnostic needs only the final snapshot.
	// Retaining every checkpoint also retains a state copy for every call and
	// method at measurement time; omit those copies without changing training.
	finalOnly := os.Getenv("PAPER_FIXED_FINAL_ONLY") == "1"
	type caller struct {
		c *studyCall
		s *state.StateDB
		e *vm.EVM
	}
	var callers []caller
	for _, c := range q.calls {
		s := c.f.pre.Copy()
		cc := vm.ExperimentRootConfig(cfg, &c.f.to)
		cc.Tracer, cc.LoopTiming = nil, nil
		e := vm.NewEVM(core.NewEVMBlockContext(c.block.Header(), headers, nil), s, config.MainnetChainConfig, cc)
		msg, err := core.TransactionToMessage(c.tx, types.MakeSigner(config.MainnetChainConfig, c.block.Number(), c.block.Time()), c.block.BaseFee())
		if err != nil {
			return err
		}
		e.SetTxContext(core.NewEVMTxContext(msg))
		callers = append(callers, caller{c, s, e})
	}
	started := time.Now()
	if !finalOnly {
		q.points = append(q.points, studyPoint{0, 0, p.ExperimentCapture(q.codes)})
	}
	var rows []map[string]any
	var preparationNS int64
	for pass := 1; pass <= passes; pass++ {
		before := time.Now()
		for _, r := range callers {
			code := r.s.GetCode(r.c.f.to)
			if target, ok := types.ParseDelegation(code); ok {
				code = r.s.GetCode(target)
			}
			r.e.OfferRegionCode(r.c.f.to, crypto.Keccak256Hash(code), code)
		}
		step := p.AdvanceValuePreparation()
		prep := time.Since(before).Nanoseconds()
		preparationNS += prep
		for j := range callers {
			r := &callers[(j+pass-1)%len(callers)]
			f := r.c.f
			snapshot := r.s.Snapshot()
			output, left, execErr := r.e.Call(f.from, f.to, f.input, f.gas, uint256.MustFromBig(f.value))
			failure := ""
			if execErr != nil {
				failure = execErr.Error()
			}
			if left != f.left || failure != f.failure || !bytes.Equal(output, f.output) || r.s.Copy().IntermediateRoot(true) != f.root || r.s.GetRefund() != f.refund || !reflect.DeepEqual(componentLogs(r.s), f.logs) {
				return fmt.Errorf("working set mismatch pass=%d tx=%s", pass, r.c.tx.Hash())
			}
			r.s.RevertToSnapshot(snapshot)
		}
		if pass == passes || !finalOnly && workingSetCheckpoint(pass, passes) {
			q.points = append(q.points, studyPoint{uint64(pass), time.Since(started).Seconds(), p.ExperimentCapture(q.codes)})
		}
		rows = append(rows, map[string]any{"pass": pass, "calls": len(callers), "elapsedSeconds": time.Since(started).Seconds(), "preparationNS": prep, "step": step, "background": p.Stats(), "resources": vm.FlowPreparationResources()})
		if pass%10 == 0 || pass == passes {
			fmt.Printf("fixed working set pass=%d/%d calls=%d elapsed=%.3f preparation=%.3f resources=%v\n", pass, passes, len(callers), time.Since(started).Seconds(), float64(preparationNS)/1e9, vm.FlowPreparationResources())
		}
	}
	manifest["scope"] = "Fixed working set, repeated original calls restored to identical historical prestate. Not a chronological block stream. Native preload only restores inputs and state. Each pass trains preparation through actual execution; no measurement-probe training. All preparation windows are separately charged."
	manifest["workingSetCalls"], manifest["workingSetPasses"] = len(callers), passes
	manifest["finalCheckpointOnly"] = finalOnly
	manifest["workingSetSeconds"], manifest["preparationSeconds"] = time.Since(started).Seconds(), float64(preparationNS)/1e9
	manifest["foregroundEndResources"], manifest["foregroundEndBackground"] = vm.FlowPreparationResources(), p.Stats()
	manifest["foregroundEndCatalog"], manifest["foregroundEndFamilies"] = vm.PrecompiledFlowStats(), vm.FlowFamilyStats()
	manifest["deliveryPolicy"], manifest["hotcallPolicy"] = vm.DeliveryPolicyReport(), vm.HotcallPolicyReport()
	if err := writeJSON(filepath.Join(out, "working-set-progress.json"), rows); err != nil {
		return err
	}
	p.Close()
	p.CancelValueCompilation()
	events, err := vm.WaitDemand()
	if err != nil {
		return err
	}
	if err := writeJSON(filepath.Join(out, "demand-events.json"), events); err != nil {
		return err
	}
	if os.Getenv("PAPER_DEMAND_AUDIT") == "1" {
		if err := writeJSON(filepath.Join(out, "preparation-candidate-audit.json"), p.ExperimentCandidateAudit()); err != nil {
			return err
		}
	}
	if err := q.finish(out, headers, cfg, sr, cr); err != nil {
		return err
	}
	if os.Getenv("PAPER_FIXED_MEMORY_COMPARE") == "1" {
		// Diagnostic only: remeasure the SAME immutable checkpoints. No work
		// is admitted and no fixture is reselected. Keep the primary output.
		oldGC, hadGC := os.LookupEnv("PAPER_FIXED_PRE_GC")
		defer func() {
			if hadGC {
				os.Setenv("PAPER_FIXED_PRE_GC", oldGC)
			} else {
				os.Unsetenv("PAPER_FIXED_PRE_GC")
			}
		}()
		os.Setenv("PAPER_FIXED_PRE_GC", "1")
		for _, name := range []string{"all-checkpoints-gc", "final-checkpoint-gc"} {
			if name == "final-checkpoint-gc" {
				last := q.points[len(q.points)-1]
				clear(q.points)
				q.points = []studyPoint{last}
			}
			dir := filepath.Join(out, name)
			if err := os.Mkdir(dir, 0755); err != nil {
				return err
			}
			fmt.Printf("fixed memory comparison %s checkpoints=%d\n", name, len(q.points))
			if err := q.finish(dir, headers, cfg, sr, cr); err != nil {
				return err
			}
		}
	}
	if os.Getenv("PAPER_FIXED_PAPER_TIMER") == "1" {
		if err := q.finishWithPaperTimer(filepath.Join(out, "final-paper-timer"), headers, cfg, sr, cr); err != nil {
			return err
		}
	}
	manifest["status"] = "complete"
	return writeJSON(filepath.Join(out, "benchmark-manifest.json"), manifest)
}
