// Copyright 2014 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package vm

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/holiman/uint256"
	"go-eth1.16.5-evm/core/tracing"
)

// Config are the configuration options for the Interpreter
type Config struct {
	StreamRegions   *StreamRegions         // Optional code-hash metadata shared by the stream execution path.
	MetadataSeed    map[common.Hash]BitVec // Diagnostic equal-cache native control.
	tradeLegacyRoot bool
	tradeRoot       common.Address
	PreparationTier uint32 // root policy inherited by every internal frame
	FrameTiming     *InterpreterLoopTiming

	RegionExecutionStats *RegionExecutionStats

	HotRegions *HotRegionRuntime // optional prebuilt region library
	Pruning    *PruningConfig    // nil selects native execution
	ABIScalar  *ABIScalarConfig  // internal resolved ABI plan configuration

	// Arithmetic is the resolved internal configuration. Use PruningConfig in new callers.
	Arithmetic *ArithmeticConfig
	// Optional diagnostics: top-level interpreter-loop wall time, without
	// double-counting child frames. Not safe to share between concurrent EVMs.
	LoopTiming *InterpreterLoopTiming

	Tracer                  *tracing.Hooks
	NoBaseFee               bool  // Forces the EIP-1559 baseFee to 0 (needed for 0 price calls)
	EnablePreimageRecording bool  // Enables recording of SHA3/keccak preimages
	ExtraEips               []int // Additional EIPS that are to be enabled

	StatelessSelfValidation bool // Generate execution witnesses and self-check against them (testing purpose)
	EnableWitnessStats      bool // Whether trie access statistics collection is enabled
}

// ScopeContext contains the things that are per-call, such as stack and memory,
// but not transients like pc and gas
type ScopeContext struct {
	Memory    *Memory
	Stack     *Stack
	Contract  *Contract
	Signature string
}

// MemoryData returns the underlying memory slice. Callers must not modify the contents
// of the returned data.
func (ctx *ScopeContext) MemoryData() []byte {
	if ctx.Memory == nil {
		return nil
	}
	return ctx.Memory.Data()
}

// StackData returns the stack data. Callers must not modify the contents
// of the returned data.
func (ctx *ScopeContext) StackData() []uint256.Int {
	if ctx.Stack == nil {
		return nil
	}
	return ctx.Stack.Data()
}

// Caller returns the current caller.
func (ctx *ScopeContext) Caller() common.Address {
	return ctx.Contract.Caller()
}

// Address returns the address where this scope of execution is taking place.
func (ctx *ScopeContext) Address() common.Address {
	return ctx.Contract.Address()
}

// CallValue returns the value supplied with this call.
func (ctx *ScopeContext) CallValue() *uint256.Int {
	return ctx.Contract.Value()
}

// CallInput returns the input/calldata with this call. Callers must not modify
// the contents of the returned data.
func (ctx *ScopeContext) CallInput() []byte {
	return ctx.Contract.Input
}

// ContractCode returns the code of the contract being executed.
func (ctx *ScopeContext) ContractCode() []byte {
	return ctx.Contract.Code
}
