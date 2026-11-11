package templateminer

// Hash preparations are straight-line semantic templates, not observations of
// particular operand values. Only constant-address, completely initialized
// 32/64-byte hash inputs are admitted. Storage accesses are always outside.
import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"

	"github.com/holiman/uint256"
)

const (
	HashInput = iota
	HashConstant
	HashCaller
	HashAddress
	HashResult
)

type HashSource struct {
	Kind, Slot int
	Word, Mask uint256.Int
}
type HashWrite struct {
	Offset uint64
	Value  HashSource
}
type HashPreparation struct {
	Entry, Exit             uint64
	PCs                     []uint64
	Inputs, Peak            int
	Offset, Size, MemoryEnd uint64
	Writes                  []HashWrite
	Outputs                 []HashSource
	StructuralID            string
}

func hashOnes() uint256.Int { var w uint256.Int; w.SetAllOne(); return w }
func hashLiteral(w uint256.Int) HashSource {
	return HashSource{Kind: HashConstant, Word: w, Mask: hashOnes()}
}

// DiscoverHashPreparations selects the longest proved suffix ending at each
// KECCAK256. Barriers are caller-owned entries; no candidate crosses them.
func DiscoverHashPreparations(ins []Instruction, barriers map[uint64]bool) []*HashPreparation {
	var out []*HashPreparation
	lastEnd := uint64(0)
	for end, in := range ins {
		if in.Op != 0x20 {
			continue
		}
		for start := max(0, end-32); start < end; start++ {
			if ins[start].PC < lastEnd {
				continue
			}
			r := analyzeHashPreparation(ins[start:end+1], barriers)
			if r != nil {
				out = append(out, r)
				lastEnd = r.Exit
				break
			}
		}
	}
	return out
}

func analyzeHashPreparation(ins []Instruction, barriers map[uint64]bool) *HashPreparation {
	const initial = 16
	stack := make([]HashSource, initial, initial+33)
	for i := range stack {
		stack[i] = HashSource{Kind: HashInput, Slot: i, Mask: hashOnes()}
	}
	need, peak := initial, initial
	r := &HashPreparation{Entry: ins[0].PC}
	pop := func() HashSource { v := stack[len(stack)-1]; stack = stack[:len(stack)-1]; return v }
	read := func(v HashSource) {
		if v.Kind == HashInput {
			need = min(need, v.Slot)
		}
	}
	for j, in := range ins {
		if barriers[in.PC] {
			return nil
		}
		r.PCs = append(r.PCs, in.PC)
		op := in.Op
		switch {
		case op == 0x5f || op >= 0x60 && op <= 0x7f:
			var w uint256.Int
			w.SetBytes(in.Data)
			stack = append(stack, hashLiteral(w))
		case op >= 0x80 && op <= 0x8f:
			n := int(op - 0x7f)
			if len(stack) < n {
				return nil
			}
			v := stack[len(stack)-n]
			read(v)
			stack = append(stack, v)
		case op >= 0x90 && op <= 0x9f:
			n := int(op - 0x8f)
			if len(stack) <= n {
				return nil
			}
			a, b := len(stack)-1, len(stack)-1-n
			read(stack[a])
			read(stack[b])
			stack[a], stack[b] = stack[b], stack[a]
		case op == 0x50:
			if len(stack) < 1 {
				return nil
			}
			read(pop())
		case op == 0x30 || op == 0x33:
			kind := HashAddress
			if op == 0x33 {
				kind = HashCaller
			}
			stack = append(stack, HashSource{Kind: kind, Mask: hashOnes()})
		case op == 0x16:
			if len(stack) < 2 {
				return nil
			}
			a, b := pop(), pop()
			read(a)
			read(b)
			if a.Kind != HashConstant {
				a, b = b, a
			}
			if a.Kind != HashConstant {
				return nil
			}
			if b.Kind == HashConstant {
				b.Word.And(&b.Word, &a.Word)
			} else {
				b.Mask.And(&b.Mask, &a.Word)
			}
			stack = append(stack, b)
		case op == 0x52:
			if len(stack) < 2 || len(r.Writes) >= 2 {
				return nil
			}
			a, b := pop(), pop()
			read(a)
			read(b)
			if a.Kind != HashConstant || !a.Word.IsUint64() || a.Word.Uint64() > 4096 {
				return nil
			}
			r.Writes = append(r.Writes, HashWrite{a.Word.Uint64(), b})
			r.MemoryEnd = max(r.MemoryEnd, a.Word.Uint64()+32)
		case op == 0x20 && j == len(ins)-1:
			if len(stack) < 2 {
				return nil
			}
			a, b := pop(), pop()
			read(a)
			read(b)
			if a.Kind != HashConstant || b.Kind != HashConstant || !a.Word.IsUint64() || !b.Word.IsUint64() {
				return nil
			}
			r.Offset, r.Size = a.Word.Uint64(), b.Word.Uint64()
			if r.Offset > 4096 || (r.Size != 32 && r.Size != 64) || len(r.Writes) != int(r.Size/32) {
				return nil
			}
			seen := map[uint64]bool{}
			for _, w := range r.Writes {
				if w.Offset < r.Offset || w.Offset >= r.Offset+r.Size || (w.Offset-r.Offset)%32 != 0 || seen[w.Offset] {
					return nil
				}
				seen[w.Offset] = true
			}
			stack = append(stack, HashSource{Kind: HashResult, Mask: hashOnes()})
			r.Exit = in.PC + 1
		default:
			return nil
		}
		peak = max(peak, len(stack))
		if len(stack) > initial+18 {
			return nil
		}
	}
	if r.Exit == 0 {
		return nil
	}
	// Untouched prefix stays in the real stack; reads of duplicate/dead values
	// still contribute to the minimum original stack-depth requirement.
	for i, v := range stack {
		if i >= initial || v.Kind != HashInput || v.Slot != i || v.Mask != hashOnes() {
			need = min(need, i)
			break
		}
	}
	if need > len(stack) || len(stack)-need > 18 {
		return nil
	}
	r.Inputs, r.Peak = initial-need, peak-need
	r.Outputs = append([]HashSource(nil), stack[need:]...)
	for i := range r.Writes {
		if r.Writes[i].Value.Kind == HashInput {
			r.Writes[i].Value.Slot -= need
		}
	}
	for i := range r.Outputs {
		if r.Outputs[i].Kind == HashInput {
			r.Outputs[i].Slot -= need
		}
	}
	// Parameterized family identity: PCs and literal values are instance data;
	// retain source kinds/slots, masks, write order and relative memory layout.
	family := *r
	family.Entry, family.Exit, family.PCs, family.StructuralID = 0, 0, nil, ""
	family.Writes = append([]HashWrite(nil), r.Writes...)
	family.Outputs = append([]HashSource(nil), r.Outputs...)
	family.Offset, family.MemoryEnd = 0, r.Size
	for i := range family.Writes {
		family.Writes[i].Offset -= r.Offset
		family.Writes[i].Value.Word.Clear()
	}
	for i := range family.Outputs {
		family.Outputs[i].Word.Clear()
	}
	b, _ := json.Marshal(family)
	sum := sha256.Sum256(b)
	r.StructuralID = hex.EncodeToString(sum[:])
	return r
}
