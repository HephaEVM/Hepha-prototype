package vm

import "time"

// Run selects the native, synchronous or asynchronous body once per call frame.
func (evm *EVM) Run(contract *Contract, input []byte, readOnly bool) ([]byte, error, time.Duration) {
	return evm.runSelected(contract, input, readOnly)
}
