package main

import (
	"encoding/hex"
	"sort"

	"go-eth1.16.5-evm/core/tracing"
	"go-eth1.16.5-evm/core/vm"
)

// Export the vanilla execution's PCs alongside its immutable state fixture.
// This limits offline C++ compilation to observed sites; verification remains
// the region miner's responsibility and is not replaced by a trace.
type paperSites struct {
	count  uint64
	frames map[*vm.Contract]map[uint64]bool
	codes  map[string]map[uint64]bool
}

func (p *paperSites) opcode(pc uint64, op byte, gas, cost uint64, scope tracing.OpContext, ret []byte, depth int, err error) {
	p.count++
	s, ok := scope.(*vm.ScopeContext)
	if !ok {
		return
	}
	pcs := p.frames[s.Contract]
	if pcs == nil {
		code := hex.EncodeToString(s.Contract.Code)
		pcs = p.codes[code]
		if pcs == nil {
			pcs = map[uint64]bool{}
			p.codes[code] = pcs
		}
		p.frames[s.Contract] = pcs
	}
	pcs[pc] = true
}
func (p *paperSites) export() map[string][]uint64 {
	r := map[string][]uint64{}
	for code, pcs := range p.codes {
		for pc := range pcs {
			r[code] = append(r[code], pc)
		}
		sort.Slice(r[code], func(i, j int) bool { return r[code][i] < r[code][j] })
	}
	return r
}
