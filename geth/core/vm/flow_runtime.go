package vm

import (
	"crypto/sha256"
	"fmt"
	"go-eth1.16.5-evm/regionexec"
	"go-eth1.16.5-evm/regionexec/prepare"
	"go/format"
	"strings"
)

func FlowDemandUsage() (int, int, int) {
	m := getDemandManager()
	if m == nil {
		return 0, 0, 0
	}
	return m.Usage()
}
func makeFlowPlan(rs []*hotRegion, proof, shared bool) *demandPlan {
	ps := prepareFlows(rs)
	if p := bindPrecompiledFlow(ps, proof, shared); p != nil {
		return p
	}
	p := makePreparedFlowPlan(ps, proof, shared)
	if precompiledFlowObserver != nil && p != nil {
		precompiledFlowObserver(ps, proof, shared, p)
	}
	return p
}
func makePreparedFlowPlan(prepared []*preparedFlow, proof, shared bool) *demandPlan {
	rs := make([]*hotRegion, len(prepared))
	for i, p := range prepared {
		rs[i] = p.original
	}
	m := getDemandManager()
	if m == nil {
		return nil
	}
	p := &demandPlan{regions: rs, branch: len(rs) > 1, shared: shared}
	cache := getFlowPreparationCache()
	cacheKey, helperRefs := flowPreparationKey(prepared, proof, shared)
	if record, ok := cache.get(cacheKey, helperRefs, len(rs), shared); ok {
		if record.Error != "" {
			return nil
		}
		p.helpers = record.bindHelpers(helperRefs)
		p.flow = record.Reports
		p.forward = make([]int, len(rs))
		p.dead = make([]int, len(rs))
		p.slot = flowPreparationSlot(m, record.Digest, shared, cache.sourceBuilder(cacheKey))
		return p.finishBinding()
	}
	for _, pp := range prepared {
		old, ob := pp.optimized(false)
		nr, nb := pp.optimized(true)
		count := func(x *hotRegion, b []hotInterval) (int, int) {
			h, g := 0, 0
			for _, n := range x.nodes {
				if n.helper != nil {
					h++
				}
				if n.op == JUMPI && !b[n.args[0]].exact() {
					g++
				}
			}
			return h, g
		}
		oh, og := count(old, ob)
		nh, ng := count(nr, nb)
		p.flow = append(p.flow, flowReport{HelpersRemoved: oh - nh, GuardsRemoved: og - ng})
	}
	var src string
	if shared {
		body, h, e := emitPreparedFlowSharedRaw(prepared, proof)
		if e != nil {
			cache.put(cacheKey, flowPreparationRecord{Error: e.Error()})
			return nil
		}
		src = body
		p.helpers = [][]regionexec.Helper{h}
	} else {
		var s strings.Builder
		for i, pp := range prepared {
			body, h, e := emitPreparedFlowPath(pp, proof, false)
			if e != nil {
				panic(e)
			}
			p.helpers = append(p.helpers, h)
			if len(rs) == 2 {
				body = strings.ReplaceAll(body, "Exec", fmt.Sprintf("Exec_%d", i))
			}
			s.WriteString(body)
		}
		if len(rs) == 2 {
			s.WriteString("\nfunc Exec(s *regionexec.Frame)bool{old:=s.Exit;s.Exit=s.Exits[0];s.Helpers=s.BranchHelpers[0];if Exec_0(s){s.Path=0;return true};s.Exit=s.Exits[1];s.Helpers=s.BranchHelpers[1];if Exec_1(s){s.Path=1;return true};s.Exit=old;return false}\n")
		}
		src = s.String()
	}
	for range rs {
		p.forward = append(p.forward, 0)
		p.dead = append(p.dead, 0)
	}
	build := prepare.Builder(func() (string, error) { return src, nil })
	if cache.put(cacheKey, flowPreparationRecord{Source: src, Reports: p.flow, Helpers: flowPreparationHelpers(prepared, helperRefs, proof, shared)}) {
		build = cache.sourceBuilder(cacheKey)
	}
	p.slot = flowPreparationSlot(m, fmt.Sprintf("%x", sha256.Sum256([]byte(src))), shared, build)
	if p.slot == nil {
		return nil
	}
	return p.finishBinding()
}

func flowPreparationSlot(m *prepare.Manager, digest string, raw bool, build prepare.Builder) *prepare.Slot {
	if !raw {
		return m.Slot("flow-v1-"+digest, build)
	}
	return m.DeferredSlot("flow-raw-v1-"+digest, func() (string, string, error) {
		source, err := build()
		if err != nil {
			return "", "", err
		}
		formatted, err := format.Source([]byte("package main\n" + source))
		if err != nil {
			return "", "", err
		}
		source = strings.TrimPrefix(string(formatted), "package main\n")
		return fmt.Sprintf("flow-v1-%x", sha256.Sum256([]byte(source))), source, nil
	})
}
func (e *EVM) bindFlow(b *boundArithmetic, code []byte, pc uint64, x *regionExtension, bitmaps ...BitVec) *demandPlan {
	bitmap := flowBitmap(code, bitmaps)

	if flowFusionBoundary(b, pc) {
		return nil
	}
	build := buildProtectedFlowPath
	{
		depth := 2
		var rs []*hotRegion
		seen := map[string]bool{}
		for mask := 0; mask < 1<<depth; mask++ {
			r := build(b, code, e.table, pc, uint(mask), depth, bitmap)
			if r == nil || len(r.nodes) > 160 || x.partial != nil && r.Instructions <= x.partial.Instructions {
				continue
			}
			if !seen[r.key] {
				rs = append(rs, r)
				seen[r.key] = true
			}
		}
		if len(rs) >= 2 {
			if p := makeFlowPlan(rs, true, true); p != nil {
				return p
			}
		}
	}
	a := buildDemandBranch(b, code, e.table, pc, false, bitmap)
	z := buildDemandBranch(b, code, e.table, pc, true, bitmap)
	if a != nil && z != nil && len(a.nodes) <= 160 && len(z.nodes) <= 160 && (x.partial == nil || a.Instructions > x.partial.Instructions && z.Instructions > x.partial.Instructions) {
		rs := []*hotRegion{a, z}
		prepared := prepareFlows(rs)
		if !preparedFlowsChanged(prepared) {
			return makePreparedDemandPlan(prepared, false)
		}
		return makePreparedFlowPlan(prepared, true, false)
	}
	if x.partial != nil && x.fn == nil {
		rs := []*hotRegion{x.partial}
		prepared := prepareFlows(rs)
		if !preparedFlowsChanged(prepared) {
			return makePreparedDemandPlan(prepared, false)
		}
		return makePreparedFlowPlan(prepared, true, false)
	}
	return nil
}
func flowChanged(rs []*hotRegion) bool { return preparedFlowsChanged(prepareFlows(rs)) }
