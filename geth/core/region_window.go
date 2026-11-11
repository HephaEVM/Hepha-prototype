package core

import (
	"go-eth1.16.5-evm/core/types"
	"go-eth1.16.5-evm/core/vm"
)

func (p *StateProcessor) ExperimentPreparationContext(h *types.Header) vm.BlockContext {
	return NewEVMBlockContext(h, p.chain, nil)
}
