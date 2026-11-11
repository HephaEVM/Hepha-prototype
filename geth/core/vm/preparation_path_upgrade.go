package vm

// Admission proxy for a verified pair composition. An unchanged short branch
// must not erase a longer branch's upgrade. Match construction provenance,
// never exits or path length alone. This is potential benefit, not a claim
// about which branch executes or its CPU cost. Only preparation calls this.
func joinedPathAdditional(d *demandPlan) int {
	if d == nil || len(d.joinParents) != len(d.regions) {
		return 0
	}
	same := func(a, b *hotRegion) bool {
		return a != nil && b != nil && (a == b || a.key != "" && a.key == b.key)
	}
	minimumPositive := 0
	for i, path := range d.regions {
		parent := d.joinParents[i]
		if parent == nil || same(path, parent) {
			continue
		}
		covered, found, duplicate := 0, false, false
		for f := d.fallback; f != nil; f = f.fallback {
			if f.slot == nil || f.slot.Loaded() == nil {
				continue
			}
			for j, prior := range f.regions {
				if same(path, prior) {
					duplicate = true
				}
				if same(parent, prior) || len(f.joinParents) == len(f.regions) && same(parent, f.joinParents[j]) {
					found = true
					covered = max(covered, prior.Instructions)
				}
			}
		}
		if !found || duplicate {
			continue
		}
		if delta := path.Instructions - covered; delta > 0 && (minimumPositive == 0 || delta < minimumPositive) {
			minimumPositive = delta
		}
	}
	return minimumPositive
}
