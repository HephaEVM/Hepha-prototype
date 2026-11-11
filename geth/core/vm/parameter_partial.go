package vm

// Only Exit is late-bound. Operand constants, gas, node/helper layouts and
// safety obligations remain part of the key. Never wildcard type/memory guards.
func parameterPartialKey(r *hotRegion) string { c := *r; c.Exit = 0; return c.fingerprint() }

var parameterPartialGenerated = map[string]hotRegionFunc{}
