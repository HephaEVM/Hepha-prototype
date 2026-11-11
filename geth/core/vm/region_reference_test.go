package vm

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common/math"
)

func regionReference(pc *uint64, end uint64, s *ScopeContext, e *EVM, fixed ...uint64) error {
	spent := uint64(0)
	for step := 0; *pc != end || (len(fixed) > 0 && spent < fixed[0]); step++ {
		if step > 100000 {
			return fmt.Errorf("step bound")
		}
		o := e.table[s.Contract.GetOp(*pc)]
		n := s.Stack.len()
		if n < o.minStack {
			return &ErrStackUnderflow{n, o.minStack}
		}
		if n > o.maxStack {
			return &ErrStackOverflow{n, o.maxStack}
		}
		if s.Contract.Gas < o.constantGas {
			return ErrOutOfGas
		}
		s.Contract.Gas -= o.constantGas
		spent += o.constantGas
		size := uint64(0)
		if o.memorySize != nil {
			v, of := o.memorySize(s.Stack)
			if of {
				return ErrGasUintOverflow
			}
			size, of = math.SafeMul(toWordSize(v), 32)
			if of {
				return ErrGasUintOverflow
			}
		}
		if o.dynamicGas != nil {
			g, err := o.dynamicGas(e, s.Contract, s.Stack, s.Memory, size)
			if err != nil {
				return fmt.Errorf("%w: %v", ErrOutOfGas, err)
			}
			if s.Contract.Gas < g {
				return ErrOutOfGas
			}
			s.Contract.Gas -= g
		}
		if size > 0 {
			s.Memory.Resize(size)
		}
		if _, err := o.execute(pc, e, s); err != nil {
			return err
		}
		*pc++
	}
	return nil
}
