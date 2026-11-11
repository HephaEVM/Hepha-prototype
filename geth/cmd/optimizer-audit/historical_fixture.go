package main

import (
	"go-eth1.16.5-evm/core"

	"go-eth1.16.5-evm/core/types"

	"strings"
)

type replacementAudit struct {
	Address, Name, Status, HistoricalHash, CheckedHash string
	Applied, Executed                                  bool
}

func checkedFixture(_ *types.Block, _ core.ChainContext, _ *types.Transaction, fixture *componentCall) (*componentCall, []replacementAudit, bool, error) {
	return fixture, nil, true, nil
}

func isCheckedArm(name string) bool { return strings.HasPrefix(name, "checked-") }
