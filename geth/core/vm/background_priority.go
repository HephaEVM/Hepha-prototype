package vm

// This is an admission proxy, not a CPU time estimate: minimum additional
// instruction coverage per graph node, weighted by bounded observed code heat.
// Loaded fallbacks already deliver part of the benefit; do not count it twice.
// No runtime stack values, timing probes or new hot-loop counters are needed.
func backgroundPriority(d *demandPlan, hits uint64) int64 {
	nodes := 0
	for _, r := range d.regions {
		nodes += len(r.nodes)
	}
	delta := deliveryAdditional(d)
	return int64(delta) * int64(min(uint64(64), max(uint64(1), hits))) * 1024 / int64(8+nodes)
}
