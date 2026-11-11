package vm

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/params"
	"github.com/holiman/uint256"
	"reflect"
)

var chainFastDirect = map[string]func(*effectChain, *uint64, *ScopeContext, *EVM) (bool, error){}

func prepareChainFast(p *effectChain, table *JumpTable) {
	same := func(a, b any) bool { return reflect.ValueOf(a).Pointer() == reflect.ValueOf(b).Pointer() }
	for i := range p.actions {
		a := &p.actions[i]
		o := table[a.op]
		var off int
		var length uint64
		switch a.op {
		case MSTORE:
			if !same(o.execute, opMstore) || !same(o.dynamicGas, gasMStore) || !same(o.memorySize, memoryMStore) {
				continue
			}
			off = a.args[1]
			length = 32
		case MLOAD:
			if !same(o.execute, opMload) || !same(o.dynamicGas, gasMLoad) || !same(o.memorySize, memoryMLoad) {
				continue
			}
			off = a.args[0]
			length = 32
		case KECCAK256:
			if !same(o.execute, opKeccak256) || !same(o.dynamicGas, gasKeccak256) || !same(o.memorySize, memoryKeccak256) {
				continue
			}
			off = a.args[1]
			length = 64
			a.fastExtra = 2 * params.Keccak256WordGas
		default:
			continue
		}
		n := p.nodes[off]
		if n.op != chainConstant || !n.c.IsUint64() || n.c.Uint64() > 4096 {
			continue
		}
		a.fastOffset = n.c.Uint64()
		a.fastEnd = ((a.fastOffset + length + 31) / 32) * 32
	}
	p.directFast = chainFastDirect[p.shape]
}
func (p *effectChain) actionFast(i int, pc *uint64, s *ScopeContext, e *EVM, in []uint256.Int, v *[256]uint256.Int) (int, error) {
	a := &p.actions[i]
	if paperWithoutReuse {
		paperRepeatChainValues(p, i, v)
	}
	if s.Contract.Gas < a.gas || e.abort.Load() {
		e.chainProgress = a.startStep
		p.restore(s, in, v, a.checkpoint)
		*pc = a.start
		return 1, nil
	}
	if a.op == JUMPI && v[a.args[0]].IsZero() {
		e.chainProgress = a.opStep
		p.restore(s, in, v, a.before)
		s.Contract.Gas -= a.gas - e.table[JUMPI].constantGas
		*pc = a.pc
		return 1, nil
	}
	if a.op == chainHelper {
		ix, z := a.helper.syncHot.native(v[a.args[0]], v[a.args[1]])
		if a.helper.syncHot.costs[ix] != a.helper.syncHot.admissionGas {
			e.chainProgress = a.opStep
			p.restore(s, in, v, a.before)
			s.Contract.Gas -= a.gas - e.table[JUMP].constantGas - a.helper.syncHot.admissionGas
			*pc = a.pc
			return 1, nil
		}
		v[a.result] = z
		s.Contract.Gas -= a.gas
		return 0, nil
	}
	s.Contract.Gas -= a.gas
	if a.op == JUMPI {
		return 0, nil
	}
	if a.fastEnd != 0 && uint64(s.Memory.Len()) >= a.fastEnd {
		if s.Contract.Gas < a.fastExtra {
			p.restore(s, in, v, a.before)
			*pc = a.pc
			return 2, ErrOutOfGas
		}
		s.Contract.Gas -= a.fastExtra
		switch a.op {
		case MSTORE:
			s.Memory.Set32(a.fastOffset, &v[a.args[0]])
		case MLOAD:
			v[a.result].SetBytes(s.Memory.GetPtr(a.fastOffset, 32))
		case KECCAK256:
			if a.reuse < 0 {
				e.hasher.Reset()
				e.hasher.Write(s.Memory.GetPtr(a.fastOffset, 64))
				e.hasher.Read(e.hasherBuf[:])
				v[a.result].SetBytes(e.hasherBuf[:])
			}
		}
		return 0, nil
	}
	old := s.Stack.data
	// A small separate argument stack prevents temporary host operations from
	// overwriting live EVM inputs. Original dynamic gas and effects stay in order.
	scratch := &e.chainScratch
	for j, id := range a.args {
		scratch[j] = v[id]
	}
	s.Stack.data = scratch[:len(a.args)]
	o := e.table[a.op]
	size := uint64(0)
	var err error
	if o.memorySize != nil {
		var overflow bool
		size, overflow = o.memorySize(s.Stack)
		if overflow {
			err = ErrGasUintOverflow
		} else {
			size, overflow = math.SafeMul(toWordSize(size), 32)
			if overflow {
				err = ErrGasUintOverflow
			}
		}
	}
	if err == nil && o.dynamicGas != nil {
		var cost uint64
		cost, err = o.dynamicGas(e, s.Contract, s.Stack, s.Memory, size)
		if err != nil {
			err = fmt.Errorf("%w: %v", ErrOutOfGas, err)
		} else if s.Contract.Gas < cost {
			err = ErrOutOfGas
		} else {
			s.Contract.Gas -= cost
		}
	}
	if err == nil {
		if size > 0 {
			s.Memory.Resize(size)
		}
		if a.op == KECCAK256 && a.reuse >= 0 { /* Same symbolic preimage; charge original dynamic gas above. */
		} else {
			e.chainExecPC = a.pc
			_, err = o.execute(&e.chainExecPC, e, s)
			if err == nil && a.result >= 0 {
				v[a.result] = *s.Stack.peek()
			}
		}
	}
	s.Stack.data = old
	if err != nil {
		p.restore(s, in, v, a.before)
		*pc = a.pc
		return 2, err
	}
	return 0, nil
}
func (p *effectChain) tryFast(pc *uint64, s *ScopeContext, e *EVM, special bool) (bool, error) {
	n := len(s.Stack.data)
	if n < p.minStack || n+p.peak > 1024 || e.Config.Tracer != nil || e.Config.EnablePreimageRecording || e.abort.Load() {
		return false, nil
	}
	// Reject an entry-only guard before allocating/evaluating the region. No
	// effect or gas charge has occurred, so the original opcode can run directly.
	if len(p.actions) > 0 && p.actions[0].op == JUMPI {
		g := p.nodes[p.actions[0].args[0]]
		if g.op == chainInput && inGuardZero(s, g.depth) {
			return false, nil
		}
	}
	if e.chainValues == nil {
		e.chainValues = new([256]uint256.Int)
	}
	if special && p.directFast != nil {
		return p.directFast(p, pc, s, e)
	}
	in := s.Stack.data
	v := e.chainValues
	for _, event := range p.events {
		if event.node >= 0 {
			j := event.node
			x := p.nodes[j]
			switch x.op {
			case chainInput:
				if x.depth <= len(in) {
					v[j] = in[len(in)-x.depth]
				}
			case chainConstant:
				v[j] = x.c
			case CALLER:
				v[j].SetBytes(effectiveCaller(s))
			default:
				v[j] = chainPure(x.op, v[x.a], v[x.b])
			}
			continue
		}
		code, err := p.actionFast(event.action, pc, s, e, in, v)
		if code != 0 {
			return *pc != p.entry || event.action != 0, err
		}
	}
	p.finish(pc, s, in, v)
	e.chainProgress = len(p.pcs)
	return true, nil
}
