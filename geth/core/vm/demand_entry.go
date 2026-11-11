package vm

// tryOutlinedDemandEntry handles tagged regions outside the main instruction
// loop. With no failure, a false result falls back to the original opcode.
// Publish a lazily allocated frame before executing its first check so the
// caller's existing deferred cleanup also handles arithmetic failures.
func (evm *EVM) tryOutlinedDemandEntry(op OpCode, pc *uint64, scope *ScopeContext, boundArithmetic *boundArithmetic, abiBound *abiScalarBound, frame **arithmeticFrame) (bool, *arithmeticResult) {
	arithmetic := *frame
	// Entry tags share the ordinary opcode fetch. Normal instructions need no
	// separate PC lookup; the region index is read only at a tagged entry.
	if op == abiScalarMarker && abiBound != nil {
		if r := abiBound.entries[*pc]; r != nil && !evm.abort.Load() && evm.Config.ABIScalar.try(r, pc, scope, nil) {
			return true, nil
		}
	}
	if op == arithmeticEntryMarker && boundArithmetic != nil {
		id := boundArithmetic.entryIndex[*pc]
		compactTried := false
		if int(id) < len(boundArithmetic.compactEntries) && !evm.abort.Load() && evm.Config.Tracer == nil && !evm.Config.EnablePreimageRecording {
			q := &boundArithmetic.compactEntries[id]
			if q.plan != nil {
				compactTried = true
				if q.try(pc, scope, evm) {
					return true, nil
				}
			}
		}
		if boundArithmetic.computeEntries != nil {
			if region := boundArithmetic.computeEntries[id]; region != nil && !evm.abort.Load() {
				if x := region.extension; x != nil && evm.Config.Tracer == nil && !evm.Config.EnablePreimageRecording {
					stats := evm.Config.RegionExecutionStats
					if q := x.probe; q != nil {
						if stats != nil {
							stats.PartialAttempts++
						}
						pure := x.probeOnly
						q.tradeRecordInput(scope.Contract.Input)
						if q.saturateObservation() {
							x.probe = nil
							if pure {
								retireObservedEntry(*pc, scope.Contract.Code, boundArithmetic, abiBound)
							}
						}
						if pure {
							return false, nil
						}
					}
					if !compactTried && x.demand != nil && x.demand.try(pc, scope, evm) {
						return true, nil
					}
					if x.loop != nil {
						if stats != nil {
							stats.LoopAttempts++

						}
						if x.loop.try(pc, scope, evm) {
							if stats != nil {
								stats.LoopCommits++

							}
							return true, nil
						}
					}
					if x.fn != nil {
						if stats != nil {
							stats.PartialAttempts++
						}
						if x.fn(pc, scope, x.partial) {
							if stats != nil {
								stats.PartialCommits++
								if x.newEntry {
									stats.BoundaryCommits++
									stats.BoundaryInstructions += uint64(x.partial.Instructions)
								}
								if x.relocated {
									stats.RelocatedCommits++
								}
								stats.PartialInstructions += uint64(x.partial.Instructions)
							}
							return true, nil
						}
					}
				}

				if p := region.address; p != nil {
					stats := evm.Config.RegionExecutionStats
					if stats != nil {
						stats.Attempts++
					}
					taken := false
					if p.chain != nil {
						var err error
						{
							taken, err = p.chain.tryFast(pc, scope, evm, true)
						}

						if !taken && err == nil && p.previous != nil {
							p = p.previous
							taken = p.try(pc, scope, evm, true)
						}
						if err != nil {
							return false, &arithmeticResult{err: err}
						}
					} else {
						taken = p.try(pc, scope, evm, true)
					}
					if taken {
						if stats != nil {
							if *pc == p.exit && (p.chain == nil || evm.chainProgress == len(p.chain.pcs)) {
								stats.Commits++
								if stats.Capture != nil {
									stats.record(p, scope.Contract)
								}
							} else {
								stats.Resumes++
							}
						}
						return true, nil
					}
				}
				if entry := region.ordinary; entry != nil && entry.fn(pc, scope, entry.region) {
					stats := &evm.Config.HotRegions.Stats
					stats.Commits++
					stats.Instructions += uint64(entry.region.Instructions)
					stats.Helpers += uint64(entry.region.Helpers)
					return true, nil
				}
				if region.preparation != nil {
					if region.preparation.try(pc, scope, evm, evm.Config.Arithmetic.HashPreparationStats) {
						return true, nil
					}
				} else if (region.muldiv != nil) && region.try(pc, scope, evm.Config.Arithmetic.ComputeStats) {
					return true, nil
				}
			}
		}
		if id != 0 && boundArithmetic.entryRegions[id] != nil {

			if arithmetic == nil {
				arithmetic = newPreparedArithmeticFrame(*evm.Config.Arithmetic, boundArithmetic)
				*frame = arithmetic
			}
			region := boundArithmetic.entryRegions[id]
			var skipped bool
			var failure *arithmeticResult
			if region.syncHot.constantPairLayout != 0 {
				if region.syncHot.constantPairLayout <= 2 {
					skipped, failure = arithmetic.trySyncConstantPair(pc, scope, region)
				} else {
					skipped, failure = arithmetic.trySyncPreservedPair(pc, scope, region)
				}
			} else if region.syncHot.fixedLayout != 0 {
				skipped, failure = arithmetic.trySyncFixed(pc, scope, region)
			} else {
				skipped, failure = arithmetic.trySyncHot(pc, scope, region)
			}
			if failure != nil {
				return false, failure
			}
			if skipped {
				return true, nil
			}
		}
	}

	return false, nil
}
