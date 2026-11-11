package vm

// Ready publication already allocates its own opcode views and observer
// extensions, separate from the worker's immutable analysis/base/full plans.
// The ABI engine and its scratch/cache belong to one sequential foreground
// runner. Only this runner removes observation tags from those private views;
// the worker only publishes a new view. Nested calls may see retired tags:
// retirement restores the original opcode and never changes EVM semantics.
func retireObservedEntry(pc uint64, code []byte, a *boundArithmetic, b *abiScalarBound) {
	if !a.observationOwned || pc >= uint64(len(a.dispatchCode)) || a.dispatchCode[pc] != byte(arithmeticEntryMarker) {
		return
	}
	a.dispatchCode[pc] = code[pc]
	if b != nil {
		b.view[pc] = code[pc]
	}
	a.observationRemaining--
	if a.observationRemaining == 0 && a.observationNative {
		a.nativeOnly = true
	}
}

func (q *backgroundEntry) saturateObservation() bool {
	if q.hits.Load() >= q.threshold {
		return true
	}
	n := q.hits.Add(1)
	if n == q.threshold {
		q.owner.observed.Add(1)
		if p := q.owner.owner; p != nil {
			p.observedEntries.Add(1)
			p.signal()
		}
	}
	return n >= q.threshold
}
