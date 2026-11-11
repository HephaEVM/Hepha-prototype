package vm

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/params"
	"github.com/holiman/uint256"
	miner "go-eth1.16.5-evm/template_miner"
)

type HashPreparationMode uint8

const (
	HashPreparationOff HashPreparationMode = iota
	HashPreparationObserve
	HashPreparationGeneric
	HashPreparationDirect
)

type HashPreparationStats struct {
	Attempts, Skipped, Instructions, Fallbacks, Words, SingleOutput uint64
	// Optional warmup/census feedback, never enabled during timing.
	Capture map[string]*HashPreparationVisit `json:",omitempty"`
}
type HashPreparationVisit struct {
	CodeHash string
	Region   *miner.HashPreparation
	Count    uint64
	Inputs   []uint256.Int
}

func (a HashPreparationStats) Sub(b HashPreparationStats) HashPreparationStats {
	return HashPreparationStats{Attempts: a.Attempts - b.Attempts, Skipped: a.Skipped - b.Skipped, Instructions: a.Instructions - b.Instructions, Fallbacks: a.Fallbacks - b.Fallbacks, Words: a.Words - b.Words, SingleOutput: a.SingleOutput - b.SingleOutput}
}

type hashPreparationPlan struct {
	fixedInputs        bool
	inputSlots         [2]int
	inputWords         [2]uint256.Int
	fixedEnd, fixedFee uint64
	region             *miner.HashPreparation
	mode               HashPreparationMode
	gas                uint64
	ops                []*operation
	single             bool
	// Proved output shapes avoid interpreting and copying unchanged stack slots.
	outputKind      uint8
	keep            int
	outputConstants []uint256.Int
	outputOne       *miner.HashSource
}

const (
	hashOutputGeneral uint8 = iota
	hashOutputIdentity
	hashOutputConstants
	hashOutputOne
)

func prepareHashPreparation(r *miner.HashPreparation, code []byte, table *JumpTable, mode HashPreparationMode) *hashPreparationPlan {
	p := &hashPreparationPlan{region: r, mode: mode}
	for _, pc := range r.PCs {
		op := table[code[pc]]
		if op == nil || op.undefined || (op.dynamicGas != nil && OpCode(code[pc]) != MSTORE && OpCode(code[pc]) != KECCAK256) {
			return nil
		}
		p.gas += op.constantGas
		p.ops = append(p.ops, op)
	}
	p.gas += params.Keccak256WordGas * (r.Size / 32)
	p.single = len(r.Outputs) == 1 && r.Outputs[0].Kind == miner.HashResult
	if mode == HashPreparationDirect {
		p.prepareOutput()
	}
	return p
}

func (p *hashPreparationPlan) prepareOutput() {
	outputs := p.region.Outputs
	var ones uint256.Int
	ones.SetAllOne()
	if len(outputs) == 0 || outputs[len(outputs)-1].Kind != miner.HashResult || outputs[len(outputs)-1].Mask != ones {
		return
	}
	prefix := outputs[:len(outputs)-1]
	for p.keep < len(prefix) {
		v := &prefix[p.keep]
		if v.Kind != miner.HashInput || v.Slot != p.keep || v.Mask != ones {
			break
		}
		p.keep++
	}
	tail := prefix[p.keep:]
	if len(tail) == 0 {
		p.outputKind = hashOutputIdentity
		return
	}
	allConstants := true
	for i := range tail {
		if tail[i].Kind != miner.HashConstant {
			allConstants = false
			break
		}
	}
	if allConstants {
		p.outputKind = hashOutputConstants
		p.outputConstants = make([]uint256.Int, len(tail))
		for i := range tail {
			p.outputConstants[i] = tail[i].Word
		}
	} else if len(tail) == 1 {
		p.outputKind, p.outputOne = hashOutputOne, &tail[0]
	}
}

func hashSourceValue(v *miner.HashSource, input []uint256.Int, c *Contract, result uint256.Int) uint256.Int {
	var w uint256.Int
	switch v.Kind {
	case miner.HashInput:
		w = input[v.Slot]
	case miner.HashConstant:
		return v.Word
	case miner.HashCaller:
		w.SetBytes(c.Caller().Bytes())
	case miner.HashAddress:
		w.SetBytes(c.Address().Bytes())
	case miner.HashResult:
		w = result
	}
	w.And(&w, &v.Mask)
	return w
}

func (p *hashPreparationPlan) try(pc *uint64, s *ScopeContext, evm *EVM, stats *HashPreparationStats) bool {
	if stats != nil {
		stats.Attempts++
		if stats.Capture != nil && len(s.Stack.data) >= p.region.Inputs {
			hash := s.Contract.CodeHash
			if !s.Contract.arithmeticHashTrusted || hash == (common.Hash{}) {
				hash = crypto.Keccak256Hash(s.Contract.Code)
			}
			key := fmt.Sprintf("%s/%d", hash.Hex(), p.region.Entry)
			v := stats.Capture[key]
			if v == nil {
				v = &HashPreparationVisit{CodeHash: hash.Hex(), Region: p.region, Inputs: append([]uint256.Int(nil), s.Stack.data[len(s.Stack.data)-p.region.Inputs:]...)}
				stats.Capture[key] = v
			}
			v.Count++
		}
	}
	if p.mode == HashPreparationObserve {
		return false
	}
	r := p.region
	n := len(s.Stack.data)
	if n < r.Inputs || n-r.Inputs+r.Peak > 1024 {
		if stats != nil {
			stats.Fallbacks++
		}
		return false
	}
	end := (r.MemoryEnd + 31) &^ uint64(31)
	if p.fixedInputs {
		end = p.fixedEnd
	}
	fee, cost := s.Memory.lastGasCost, p.gas
	if end > uint64(s.Memory.Len()) {
		words := end / 32
		if p.fixedInputs {
			fee = p.fixedFee
		} else {
			fee = words*params.MemoryGas + words*words/params.QuadCoeffDiv
		}
		if fee < s.Memory.lastGasCost {
			if stats != nil {
				stats.Fallbacks++
			}
			return false
		}
		cost += fee - s.Memory.lastGasCost
	}
	if s.Contract.Gas < cost {
		if stats != nil {
			stats.Fallbacks++
		}
		return false
	}
	// No remaining rejection: the original peak, all costs and the complete
	// read/write set have been checked before mutating live execution state.
	s.Memory.Resize(end)
	s.Memory.lastGasCost = fee
	if p.mode == HashPreparationGeneric {
		// Diagnostic predecoded execution of precisely the same proved region.
		// This retains every original stack operation; gas is charged as a unit.
		for i, op := range p.ops {
			q := r.PCs[i]
			if _, err := op.execute(&q, evm, s); err != nil {
				panic("proved hash preparation rejected: " + err.Error())
			}
		}
	} else {
		input := s.Stack.data[n-r.Inputs:]
		// Resolve both words before output stack writes; actual memory writes
		// remain in original order, including their externally visible bytes.
		var first uint256.Int
		if p.fixedInputs {
			first = p.inputWord(0, input)
		} else {
			first = hashSourceValue(&r.Writes[0].Value, input, s.Contract, uint256.Int{})
		}
		s.Memory.Set32(r.Writes[0].Offset, &first)
		if len(r.Writes) == 2 {
			var second uint256.Int
			if p.fixedInputs {
				second = p.inputWord(1, input)
			} else {
				second = hashSourceValue(&r.Writes[1].Value, input, s.Contract, uint256.Int{})
			}
			s.Memory.Set32(r.Writes[1].Offset, &second)
		}
		data := s.Memory.GetPtr(r.Offset, r.Size)
		evm.hasher.Reset()
		evm.hasher.Write(data)
		evm.hasher.Read(evm.hasherBuf[:])
		if evm.Config.EnablePreimageRecording {
			evm.StateDB.AddPreimage(evm.hasherBuf, data)
		}
		var result uint256.Int
		result.SetBytes(evm.hasherBuf[:])
		base := n - r.Inputs
		switch p.outputKind {
		case hashOutputIdentity:
			s.Stack.data = append(s.Stack.data[:base+p.keep], result)
		case hashOutputConstants:
			s.Stack.data = append(s.Stack.data[:base+p.keep], p.outputConstants...)
			s.Stack.data = append(s.Stack.data, result)
		case hashOutputOne:
			// Read before append can overwrite or reallocate the input backing store.
			value := hashSourceValue(p.outputOne, input, s.Contract, result)
			s.Stack.data = append(s.Stack.data[:base+p.keep], value, result)
		default:
			var values [18]uint256.Int
			for i := range r.Outputs {
				values[i] = hashSourceValue(&r.Outputs[i], input, s.Contract, result)
			}
			s.Stack.data = append(s.Stack.data[:base], values[:len(r.Outputs)]...)
		}
	}
	s.Contract.Gas -= cost
	*pc = r.Exit
	if stats != nil {
		stats.Skipped++
		stats.Instructions += uint64(len(r.PCs))
		stats.Words += r.Size / 32
		if p.single {
			stats.SingleOutput++
		}
	}
	return true
}

// Selected only from the symbolic source and exact full-word mask, never
// observed values. Fixed memory expansion totals depend only on the plan.
func (p *hashPreparationPlan) prepareInputs() {
	if p.mode != HashPreparationDirect || len(p.region.Writes) > 2 {
		return
	}
	var ones uint256.Int
	ones.SetAllOne()
	for i, w := range p.region.Writes {
		switch w.Value.Kind {
		case miner.HashInput:
			if w.Value.Mask != ones {
				return
			}
			p.inputSlots[i] = w.Value.Slot
		case miner.HashConstant:
			p.inputSlots[i] = -1
			p.inputWords[i] = w.Value.Word
		default:
			return
		}
	}
	p.fixedEnd = (p.region.MemoryEnd + 31) &^ uint64(31)
	words := p.fixedEnd / 32
	p.fixedFee = words*params.MemoryGas + words*words/params.QuadCoeffDiv
	p.fixedInputs = true
}
func (p *hashPreparationPlan) inputWord(i int, input []uint256.Int) uint256.Int {
	if p.inputSlots[i] >= 0 {
		return input[p.inputSlots[i]]
	}
	return p.inputWords[i]
}
