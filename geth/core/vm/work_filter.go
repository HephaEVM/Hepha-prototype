package vm

import "os"

// Worker-only candidate policy. Full region verification and native fallback
// remain authoritative. Positive saved-work metadata does not guarantee speed.
var experimentWorkPolicy = func() string {
	p := os.Getenv("COLD_WORK_POLICY")
	if p == "" {
		return "all"
	}
	switch p {
	case "all", "semantic", "none":
		return p
	default:
		panic("invalid COLD_WORK_POLICY")
	}
}()

func experimentWorkEligible(p *demandPlan) bool {
	switch experimentWorkPolicy {
	case "all":
		return true
	case "none":
		return false
	}
	if p == nil || len(p.counts) == 0 {
		return false
	}
	for _, q := range p.counts {
		if q.guards == 0 && q.helpers == 0 && q.forwarded == 0 && q.stores == 0 {
			return false
		}
	}
	return true
}
