package vm

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
)

type ExperimentReadySnapshot struct {
	plans  map[arithmeticCacheKey]*abiExecutionPlan
	status map[common.Hash]ExperimentCodeReadiness
}

// Untimed diagnostic of plan delivery, not a measurement of executed coverage.
type ExperimentCodeReadiness struct {
	Admitted, BasePublished                 bool
	ObservedEntries, AnalyzedEntries        uint64
	ReadyEntries, DemandEntries, ABIEntries int
}

func (s *ExperimentReadySnapshot) CodeReadiness() map[common.Hash]ExperimentCodeReadiness {
	x := make(map[common.Hash]ExperimentCodeReadiness, len(s.status))
	for h, v := range s.status {
		x[h] = v
	}
	return x
}

// ExperimentSnapshotCache shares frozen copies of the same immutable publication
// across dense checkpoints. Live observation counters are always sampled anew.
type ExperimentSnapshotCache struct {
	source map[arithmeticCacheKey]*abiExecutionPlan
	frozen map[arithmeticCacheKey]*abiExecutionPlan
}

func (c *ExperimentSnapshotCache) freeze(k arithmeticCacheKey, v *abiExecutionPlan, code []byte) *abiExecutionPlan {
	if c == nil {
		return experimentFreezeView(v, code)
	}
	if c.source[k] == v {
		return c.frozen[k]
	}
	if c.source == nil {
		c.source = make(map[arithmeticCacheKey]*abiExecutionPlan)
		c.frozen = make(map[arithmeticCacheKey]*abiExecutionPlan)
	}
	f := experimentFreezeView(v, code)
	c.source[k], c.frozen[k] = v, f
	return f
}
func (p *BackgroundPreparation) ExperimentCapture(hashes map[common.Hash]bool) *ExperimentReadySnapshot {
	return p.ExperimentCaptureCached(hashes, nil)
}
func (p *BackgroundPreparation) ExperimentCaptureCached(hashes map[common.Hash]bool, cache *ExperimentSnapshotCache) *ExperimentReadySnapshot {
	s := &ExperimentReadySnapshot{plans: map[arithmeticCacheKey]*abiExecutionPlan{}, status: map[common.Hash]ExperimentCodeReadiness{}}
	for h := range hashes {
		s.status[h] = ExperimentCodeReadiness{}
	}
	p.mu.RLock()
	defer p.mu.RUnlock()
	for k, slot := range p.slots {
		if hashes[k.hash] {
			r := ExperimentCodeReadiness{Admitted: true, BasePublished: slot.tracking.Load(), ObservedEntries: slot.observed.Load(), AnalyzedEntries: slot.analyzed.Load()}
			if v := slot.plan.Load(); v != nil {
				frozen := cache.freeze(k, v, p.windowCodes[k])
				s.plans[k] = frozen
				if a := frozen.arithmetic; a != nil {
					for _, id := range a.entryIndex {
						if id == 0 {
							continue
						}
						r.ReadyEntries++
						if int(id) < len(a.computeEntries) {
							c := a.computeEntries[id]
							if c != nil && c.extension != nil && c.extension.demand != nil {
								r.DemandEntries++
							}
						}
					}
				}
				if frozen.abi != nil {
					for _, entry := range frozen.abi.entries {
						if entry != nil {
							r.ABIEntries++
						}
					}
				}
			}
			s.status[k.hash] = r
		}
	}
	return s
}
func (s *ExperimentReadySnapshot) Config(c Config) Config {
	old := c.Pruning
	abi := NewABIScalarDiagnostic(old.ABI.Known)
	abi.Regions = old.ABI.Regions
	abi.Optimize = true
	p := &BackgroundPreparation{engine: old.Engine, abi: abi, closed: true, slots: map[arithmeticCacheKey]*backgroundPlan{}}
	for k, v := range s.plans {
		slot := &backgroundPlan{}
		slot.tracking.Store(true)
		slot.plan.Store(v)
		slot.cachedOnly.Store(v)
		p.slots[k] = slot
		if v.hasReadyRegion() {
			p.readyPresence[k.hash[0]&63].Or(uint64(1) << (k.hash[1] & 63))
		}
	}
	pc := *old
	pc.ABI = abi
	pc.Preparation = p
	c = c.WithPruning(pc)
	c.LoopTiming = new(InterpreterLoopTiming)
	return c
}

// Diagnostic only: use the existing cached-only lookup for ALL frozen frames.
// The exact plan pointers, selector gates and missing-plan native fallback are
// preserved. This does not build or admit plans for missing (e.g. init) code.
func (s *ExperimentReadySnapshot) DirectConfig(c Config) (Config, error) {
	for k, p := range s.plans {
		if p.tradeAllowed != nil || p.preparing != nil {
			return Config{}, fmt.Errorf("direct snapshot requires unrestricted frozen plans: %s", k.hash)
		}
	}
	c = s.Config(c)
	c.Pruning.Preparation.frozenReadyOnly = true
	return c, nil
}

// EmptyFrameConfig additionally sends already prepared negative plans through
// the same region interpreter used by fully warm bindings. No regions are
// added; absent plans, unsupported forks and tracing still use RunNative.
func (s *ExperimentReadySnapshot) EmptyFrameConfig(c Config) (Config, error) {
	c, err := s.DirectConfig(c)
	if err == nil {
		p := c.Pruning.Preparation
		p.frozenLeanFallback = true
		// A negative prepared plan has no optimization bit in the ordinary
		// ready-presence filter. Retain its plan identity for this loop test.
		for k := range s.plans {
			p.readyPresence[k.hash[0]&63].Or(uint64(1) << (k.hash[1] & 63))
		}
	}
	return c, err
}
func ExperimentWarmConfig(c Config) Config {
	panic("unmatched warm reference disabled in window study")
}

// ExperimentCandidateAudit must run after Close. Request graphs are owned by
// the preparation worker; a queue lock alone does not make them safe to inspect.
func (p *BackgroundPreparation) ExperimentCandidateAudit() []map[string]any {
	select {
	case <-p.done:
	default:
		panic("candidate audit requires stopped preparation worker")
	}
	return p.candidateAudit
}

// Called by the synchronous preparation owner during shutdown, before pending
// requests are retired. No graph is inspected concurrently with analysis.
func (p *BackgroundPreparation) capturePendingCandidates() []map[string]any {
	var rows []map[string]any
	for _, r := range p.queue {
		for _, d := range r.demands {
			var paths []DemandExecutionPath
			for _, region := range d.regions {
				paths = append(paths, DemandExecutionPath{Exit: region.Exit, Gas: region.Gas, Instructions: region.Instructions, Nodes: len(region.nodes), Helpers: region.Helpers, Inputs: region.Inputs, Drop: region.Drop})
			}
			if len(paths) == 0 {
				continue
			}
			rows = append(rows, map[string]any{"codeHash": r.key.hash, "plan": demandCapturePlan(d), "entry": d.regions[0].Entry, "paths": paths, "hits": r.slot.hits.Load(), "loaded": d.slot.Loaded() != nil, "started": d.slot.PreparationStarted(), "additional": deliveryAdditional(d), "score": backgroundPriority(d, r.slot.hits.Load()), "protected": r.slot.deliveryLegacy.Load(), "admit": deliveryCompileAdmit(r, d, d.slot.PreparationStarted())})
		}
	}
	return rows
}
