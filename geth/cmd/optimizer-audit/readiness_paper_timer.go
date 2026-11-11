package main

// Use the existing 7.3 timer on a frozen online plan. This is a measurement
// diagnostic, not a warm-reference builder or another preparation policy.
import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"runtime/pprof"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"go-eth1.16.5-evm/config"
	"go-eth1.16.5-evm/core"
	"go-eth1.16.5-evm/core/vm"
)

func (q *readinessStudy) finishWithPaperTimer(out string, headers core.ChainContext, cfg vm.Config, sr, cr *benchmarkReadCache) error {
	if len(q.points) == 0 || len(libraryCases) != 0 || len(libraryPreparers) != 0 {
		return fmt.Errorf("paper timer requires a frozen point and an unused warm harness")
	}
	for k, v := range map[string]string{"PAPER_WARM": "1", "PAPER_ABLATION": "full", "HASH_ITERATIONS": "8", "HASH_ORDER_OFFSET": "0"} {
		old, present := os.LookupEnv(k)
		defer func() {
			if present {
				os.Setenv(k, old)
			} else {
				os.Unsetenv(k)
			}
		}()
		os.Setenv(k, v)
	}
	defer func() { libraryCases = nil }()
	if err := os.Mkdir(out, 0755); err != nil {
		return err
	}
	f, err := os.Create(filepath.Join(out, "observations.jsonl"))
	if err != nil {
		return err
	}
	defer f.Close()
	enc := json.NewEncoder(f)
	last := q.points[len(q.points)-1]
	native := config.DefaultVmConfig

	native.LoopTiming = new(vm.InterpreterLoopTiming)
	arms := []*hashArm{{name: "vanilla-native", cfg: native}, {name: "hepha-frozen", cfg: last.snapshot.Config(cfg)}}
	if os.Getenv("PAPER_FIXED_DISPATCH_COMPARE") == "1" {
		direct, err := last.snapshot.DirectConfig(cfg)
		if err != nil {
			return err
		}
		arms = append(arms, &hashArm{name: "hepha-frozen-direct", cfg: direct})
		if os.Getenv("PAPER_FIXED_EMPTY_FRAME_COMPARE") == "1" {
			lean, err := last.snapshot.EmptyFrameConfig(cfg)
			if err != nil {
				return err
			}
			arms = append(arms, &hashArm{name: "hepha-frozen-direct-lean", cfg: lean})
		}
	}
	for _, a := range arms {
		if a.cfg.Pruning != nil {
			a.engine, a.abi = a.cfg.Pruning.Engine, a.cfg.Pruning.ABI
		}
	}
	before, checks := probeStamp(), 0
	for _, c := range q.calls {
		// Preserve the same per-root preparation entitlement and selector
		// context as readinessStudy.finish; the timer must not select the
		// cheaper cached-only lookup by leaving PreparationTier at zero.
		callArms := make([]*hashArm, len(arms))
		for i, a := range arms {
			copy := *a
			copy.cfg = vm.ExperimentRootConfig(a.cfg, &c.f.to)
			callArms[i] = &copy
		}
		if err := hashRunTarget(c.block, headers, nil, c.tx, 0, 0, nil, common.Hash{}, callArms, 9, 0, sr, cr, enc, &checks, c.f); err != nil {
			return err
		}
	}
	if err := runPaperWarmCases(out, enc, 9); err != nil {
		return err
	}
	if probeStamp() != before {
		return fmt.Errorf("frozen paper timer admitted preparation")
	}
	if os.Getenv("PAPER_FIXED_DISPATCH_PROFILE") == "1" {
		// Profiles follow all latency measurements and never supply paper
		// latency samples. Each method receives the complete same call set.
		for i, a := range arms {
			f, err := os.Create(filepath.Join(out, a.name+".pprof"))
			if err != nil {
				return err
			}
			if err = pprof.StartCPUProfile(f); err != nil {
				f.Close()
				return err
			}
			deadline := time.Now().Add(12 * time.Second)
			for time.Now().Before(deadline) && err == nil {
				for j := i; j < len(libraryCases); j += len(arms) {
					if _, err = libraryCases[j].run("profile-only", 0, 128, true); err != nil {
						break
					}
				}
			}
			pprof.StopCPUProfile()
			f.Close()
			if err != nil {
				return err
			}
		}
		if probeStamp() != before {
			return fmt.Errorf("frozen dispatch profiling admitted preparation")
		}
	}
	return writeJSON(filepath.Join(out, "same-plan-paper-timer.json"), map[string]any{
		"status": "complete", "calls": len(q.calls), "checkpointPass": last.block, "validatedCalls": checks, "preserveRootContext": true,
		"dispatchComparison": os.Getenv("PAPER_FIXED_DISPATCH_COMPARE") == "1", "profilesAfterTiming": os.Getenv("PAPER_FIXED_DISPATCH_PROFILE") == "1",
		"emptyFrameComparison": os.Getenv("PAPER_FIXED_EMPTY_FRAME_COMPARE") == "1",
		"scope":                "Identical final frozen online plan and original calls, measured by the unmodified 7.3 hashRunTarget/runPaperWarmCases harness. Paired method order, per-batch warmup, shared restored fixture state, untimed GC and memory counters match 7.3. No new preparation or input selection; preparation fingerprint checked unchanged.",
	})
}
