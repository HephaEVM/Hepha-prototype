package vm

import (
	"crypto/sha256"
	"fmt"
	"strings"
)

// DemandExecutionVisit is an opt-in, untimed diagnostic. It describes the
// actual chosen region, including failed attempts and shorter fallbacks, rather
// than interpreting the number of published entries as executed coverage.
type DemandExecutionVisit struct {
	CodeHash, ContextAddress, Plan string
	Entry                          uint64
	Shared, Joined                 bool
	Attempts, Commits              uint64
	Paths                          []DemandExecutionPath
}

type DemandExecutionPath struct {
	Exit, Gas                    uint64
	Instructions, Nodes, Helpers int
	Inputs, Drop                 int
	Commits                      uint64
	// Untimed source-level accounting, not CPU-retired check instructions.
	OriginalHelpers, RetainedHelpers                        int
	OriginalGuards, RetainedGuards, ConstantGuardConditions int
	PlainGuardFolds                                         int
}

func demandCapturePlan(p *demandPlan) string {
	var identity strings.Builder
	for _, r := range p.regions {
		fmt.Fprintf(&identity, "%d/%d/%d/%d/%s;", r.Entry, r.Exit, r.Gas, r.Instructions, r.key)
	}
	return fmt.Sprintf("%x", sha256.Sum256([]byte(identity.String())))
}

func (s *RegionExecutionStats) recordDemand(p *demandPlan, c *Contract, path int) {
	if s.DemandCapture == nil {
		return
	}
	digest := demandCapturePlan(p)
	key := fmt.Sprintf("%s/%s/%d/%s", c.CodeHash.Hex(), c.Address().Hex(), p.regions[0].Entry, digest)
	v := s.DemandCapture[key]
	if v == nil {
		v = &DemandExecutionVisit{CodeHash: c.CodeHash.Hex(), ContextAddress: c.Address().Hex(), Plan: digest, Entry: p.regions[0].Entry, Shared: p.shared, Joined: p.fallback != nil}
		for _, r := range p.regions {
			q := DemandExecutionPath{Exit: r.Exit, Gas: r.Gas, Instructions: r.Instructions, Nodes: len(r.nodes), Helpers: r.Helpers, Inputs: r.Inputs, Drop: r.Drop}
			for _, n := range r.nodes {
				if n.helper != nil {
					q.OriginalHelpers++
				} else if n.op == JUMPI {
					q.OriginalGuards++
				}
			}
			_, _, _, plainReport := optimizeHotRegionChecks(r, true)
			q.PlainGuardFolds = plainReport.GuardConstants
			var lowered *hotRegion
			if len(p.flow) > 0 {
				lowered, _, _ = optimizeFlowRegionChecks(r, !paperWithoutCheckElimination)
			} else {
				lowered, _, _, _ = optimizeHotRegionChecks(r, !paperWithoutCheckElimination)
			}
			for _, n := range lowered.nodes {
				if n.helper != nil {
					q.RetainedHelpers++
				} else if n.op == JUMPI {
					q.RetainedGuards++
					id := n.args[0]
					if id >= hotRegionInputs && lowered.nodes[id-hotRegionInputs].op == PUSH32 {
						q.ConstantGuardConditions++
					}
				}
			}
			v.Paths = append(v.Paths, q)
		}
		s.DemandCapture[key] = v
	}
	if path < 0 {
		v.Attempts++
	} else {
		v.Commits++
		v.Paths[path].Commits++
	}
}
