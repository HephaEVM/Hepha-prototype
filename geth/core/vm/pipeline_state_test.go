package vm

import (
	"bytes"

	"slices"

	"github.com/holiman/uint256"
)

type pipelineSnapshot struct {
	Stack            []uint256.Int
	Memory, Return   []byte
	Gas, Fee, PC     uint64
	Error, Signature string
}

func pipelineSnapshotOf(s *ScopeContext, pc uint64, ret []byte, err error) pipelineSnapshot {
	e := ""
	if err != nil && err != errStopToken {
		e = err.Error()
	}
	return pipelineSnapshot{append([]uint256.Int(nil), s.Stack.data...), append([]byte(nil), s.Memory.store...), append([]byte(nil), ret...), s.Contract.Gas, s.Memory.lastGasCost, pc, e, s.Signature}
}

func samePipelineState(a, b pipelineSnapshot) bool {
	return a.Gas == b.Gas && a.Fee == b.Fee && a.PC == b.PC && a.Error == b.Error && a.Signature == b.Signature && slices.Equal(a.Stack, b.Stack) && bytes.Equal(a.Memory, b.Memory) && bytes.Equal(a.Return, b.Return)
}
