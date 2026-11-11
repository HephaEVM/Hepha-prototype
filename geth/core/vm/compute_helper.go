package vm

import (
	"github.com/holiman/uint256"
	miner "go-eth1.16.5-evm/template_miner"
)

type ComputeHelperStats struct {
	MulDiv, MulDivWide, MulDivInstructions, MulDivFallbacks                        uint64
	Packed, PackedInstructions, PackedFallbacks, PackedFields                      uint64
	MulDivFusions, MulDivFusionInstructions, MulDivRoundUps, MulDivFusionFallbacks uint64
}

func (a ComputeHelperStats) Sub(b ComputeHelperStats) ComputeHelperStats {
	return ComputeHelperStats{a.MulDiv - b.MulDiv, a.MulDivWide - b.MulDivWide, a.MulDivInstructions - b.MulDivInstructions, a.MulDivFallbacks - b.MulDivFallbacks, a.Packed - b.Packed, a.PackedInstructions - b.PackedInstructions, a.PackedFallbacks - b.PackedFallbacks, a.PackedFields - b.PackedFields,
		a.MulDivFusions - b.MulDivFusions, a.MulDivFusionInstructions - b.MulDivFusionInstructions, a.MulDivRoundUps - b.MulDivRoundUps, a.MulDivFusionFallbacks - b.MulDivFusionFallbacks}
}
func (c *ArithmeticConfig) computeHelperPolicy() uint8 {

	if (c.Mode != ArithmeticSync) || c.GasPolicy != OriginalArithmeticGas {
		return 0
	}
	var p uint8
	{
		p |= 1
		{
			p |= 4
		}
	}

	if c.Mode == ArithmeticSync && c.HashPreparation <= HashPreparationDirect {
		p |= uint8(c.HashPreparation) << 3
	}
	return p
}

type computeHelperPath struct {
	gas, instructions uint64
	peak              int
}
type computeHelperRegion struct {
	extension   *regionExtension
	address     *addressReusePlan
	ordinary    *hotRegionExecutable
	muldiv      *miner.MulDivRegion
	short, wide computeHelperPath

	preparation *hashPreparationPlan
	rounding    *miner.MulDivRoundingRegion
	rounded     [4]computeHelperPath
}

func computePath(path miner.ABIRegionPath, code []byte, table *JumpTable) (computeHelperPath, bool) {
	p := computeHelperPath{peak: path.Peak, instructions: uint64(len(path.PCs))}
	for _, pc := range path.PCs {
		if pc >= uint64(len(code)) {
			return p, false
		}
		o := table[code[pc]]
		if o == nil || o.undefined || o.dynamicGas != nil || ^uint64(0)-p.gas < o.constantGas {
			return p, false
		}
		p.gas += o.constantGas
	}
	return p, true
}
func (b *boundArithmetic) prepareComputeHelpers(code []byte, key arithmeticCacheKey) {
	ins, e := miner.Disassemble(code)
	if e != nil {
		return
	}
	p := &miner.Program{Code: code, Instructions: ins, ByPC: map[uint64]miner.Instruction{}}
	for _, in := range ins {
		p.ByPC[in.PC] = in
	}
	found := map[uint64]*computeHelperRegion{}
	if key.compute&1 != 0 {
		for _, entry := range miner.MulDivCandidates(p) {
			r, e := miner.AnalyzeMulDiv(p, entry)
			if e != nil {
				continue
			}
			short, ok := computePath(r.Short, code, key.rules)
			if !ok {
				continue
			}
			wide, ok := computePath(r.Wide, code, key.rules)
			if !ok {
				continue
			}
			found[entry] = &computeHelperRegion{muldiv: r, short: short, wide: wide}
		}
	}
	if key.compute&4 != 0 {
		floors := map[uint64]*miner.MulDivRegion{}
		for entry, r := range found {
			floors[entry] = r.muldiv
		}
		for entry, inner := range miner.MulDivRoundingCandidates(p, floors) {
			rounding, err := miner.AnalyzeMulDivRounding(p, entry, inner)
			if err != nil {
				continue
			}
			floor := found[inner]
			r := &computeHelperRegion{muldiv: floor.muldiv, short: floor.short, wide: floor.wide, rounding: rounding}
			valid := true
			for i, path := range rounding.Paths {
				var ok bool
				r.rounded[i], ok = computePath(path, code, key.rules)
				valid = valid && ok
				// Keep existing arithmetic coverage/accounting unaltered. A
				// future heterogeneous composer can explicitly own overlaps.
				for _, pc := range path.PCs {
					if b.regions[pc] != nil {
						valid = false
					}
				}
			}
			if valid {
				found[entry] = r
			}
		}
	}

	// Preserve existing extended arithmetic entries. Passing one could
	// replace work already removed by arithmetic with a second takeover.

	if mode := HashPreparationMode((key.compute >> 3) & 3); mode != HashPreparationOff {
		barriers := make(map[uint64]bool)
		for pc := range b.regions {
			barriers[pc] = true
		}
		for pc := range found {
			barriers[pc] = true
		}
		for _, r := range miner.DiscoverHashPreparations(ins, barriers) {
			if plan := prepareHashPreparation(r, code, key.rules, mode); plan != nil {

				found[r.Entry] = &computeHelperRegion{preparation: plan}
			}
		}
	}
	if len(found) == 0 {
		return
	}
	if b.dispatchCode == nil {
		b.dispatchCode = append([]byte(nil), code...)
		b.entryIndex = make([]uint16, len(code))
		b.entryRegions = make([]*arithmeticRegion, 1)
	}
	b.computeEntries = make([]*computeHelperRegion, len(b.entryRegions))
	for pc, r := range found {
		i := b.entryIndex[pc]
		if i == 0 {
			if len(b.entryRegions) >= 1<<16 {
				continue
			}
			i = uint16(len(b.entryRegions))
			b.entryRegions = append(b.entryRegions, nil)
			b.computeEntries = append(b.computeEntries, nil)
			b.entryIndex[pc] = i
		}
		b.computeEntries[i] = r
		b.dispatchCode[pc] = byte(arithmeticEntryMarker)
	}
}
func (r *computeHelperRegion) try(pc *uint64, s *ScopeContext, stats *ComputeHelperStats) bool {
	if r.rounding != nil {
		if r.tryMulDivRounding(pc, s, stats) {
			return true
		}
		if stats != nil {
			stats.MulDivFusionFallbacks++
		}
		return false
	}

	if r.tryMulDiv(pc, s, stats) {
		return true
	}
	if stats != nil {
		stats.MulDivFallbacks++
	}
	return false
}
func (r *computeHelperRegion) tryMulDiv(pc *uint64, s *ScopeContext, stats *ComputeHelperStats) bool {
	n := len(s.Stack.data)
	if n < 4 {
		return false
	}
	in := s.Stack.data[n-4:]
	m := r.muldiv
	d := &in[m.D]
	if d.IsZero() {
		return false
	}
	var product, result uint256.Int
	_, wide := product.MulOverflow(&in[m.A], &in[m.B])
	path := &r.short
	if wide {
		path = &r.wide
	}
	if n-4+path.peak > 1024 || s.Contract.Gas < path.gas || !s.Contract.validJumpdest(&in[0]) {
		return false
	}
	if wide {
		if !tryMulDivFloor(&result, &in[m.A], &in[m.B], d) {
			return false
		}
	} else {
		result.Div(&product, d)
	}
	dest := in[0].Uint64()
	s.Stack.data[n-4] = result
	s.Stack.data = s.Stack.data[:n-3]
	s.Contract.Gas -= path.gas
	*pc = dest
	if stats != nil {
		stats.MulDiv++
		stats.MulDivInstructions += path.instructions
		if wide {
			stats.MulDivWide++
		}
	}
	return true
}
