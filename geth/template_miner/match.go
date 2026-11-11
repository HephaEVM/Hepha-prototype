package templateminer

import (
	"encoding/hex"
	"fmt"
	"sort"

	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/crypto"
)

type RuntimeBinding struct {
	Semantics   string `json:"semantics,omitempty"`
	TemplateID  string `json:"templateId"`
	Entry       uint64 `json:"entry"`
	Operation   string `json:"operation"`
	IntegerType string `json:"integerType,omitempty"`
	// 0 denotes top of stack. Extra-argument helpers have a deeper continuation.
	ContinuationOffset int                `json:"continuationOffsetFromTop"`
	InputStack         []string           `json:"inputStackBottomToTop"`
	OperandOffsets     []int              `json:"operandOffsetsFromTop"`
	AddressBindings    map[string]uint64  `json:"addressBindings"`
	Instructions       []BoundInstruction `json:"instructions"`
	// Analysis is rebuilt from this exact bytecode, never trusted from a bundle.
	Analysis *SymbolicAnalysis `json:"-"`
}
type CandidateHint struct {
	Entry      uint64   `json:"entry"`
	Reason     string   `json:"reason"`
	TemplateID string   `json:"templateId,omitempty"`
	Operands   []string `json:"observedInputsAThenB,omitempty"`
}
type MatchReport struct {
	SchemaVersion   int              `json:"schemaVersion"`
	RuntimeCodeHash string           `json:"runtimeCodeHash"`
	Mode            string           `json:"mode"`
	Bindings        []RuntimeBinding `json:"bindings"`
	Candidates      []CandidateHint  `json:"candidates,omitempty"`
}

// MatchBytecode needs only deployed runtime bytecode and learned templates. It
// resolves jumps anew and compares the entire closed region, including panic
// guards. It never consults Solidity metadata or existing manual constants.
// This deliberately conservative matcher excludes unresolved/inlined layouts.
type BytecodeMatcher struct {
	known    map[string]Template
	models   map[string]bool
	prefixes map[string][][]string
}

// NewBytecodeMatcher validates and indexes an immutable library once, shared by
// all code bindings. No bytecode-dependent analysis is reused across code hashes.
func NewBytecodeMatcher(templates []Template) (*BytecodeMatcher, error) {
	known := map[string]Template{}
	models := map[string]bool{"": true}
	for _, tpl := range templates {
		expected, err := templateID(tpl)
		if err != nil || tpl.ID != expected {
			return nil, fmt.Errorf("template content does not match its ID")
		}
		if !binaryOperation(tpl.Operation) {
			return nil, fmt.Errorf("unsupported template operation")
		}
		if tpl.Semantics != "" && !isLegacySemantics(tpl.Semantics) {
			return nil, fmt.Errorf("unknown arithmetic semantics")
		}
		models[tpl.Semantics] = true
		known[tpl.ID] = tpl
	}
	m := &BytecodeMatcher{known: known, models: models, prefixes: map[string][][]string{}}
	seen := map[string]bool{}
	for _, t := range templates {
		if len(t.Paths) == 0 {
			return nil, fmt.Errorf("template has no paths")
		}
		prefix := []string{}
		for _, ins := range t.Paths[0].Instructions {
			prefix = append(prefix, ins.Op)
			if len(prefix) == 12 || ins.Op == "JUMP" || ins.Op == "JUMPI" || ins.Op == "REVERT" || ins.Op == "INVALID" {
				break
			}
		}
		key := t.Semantics + fmt.Sprint(prefix)
		if !seen[key] {
			m.prefixes[t.Semantics] = append(m.prefixes[t.Semantics], prefix)
			seen[key] = true
		}
	}
	return m, nil
}

func (m *BytecodeMatcher) possible(p *Program, entry uint64, semantics string) bool {
	for _, prefix := range m.prefixes[semantics] {
		pc := entry
		ok := true
		for _, op := range prefix {
			ins, exists := p.ByPC[pc]
			if !exists || (op == "PUSH" && !(ins.Op >= vm.PUSH0 && ins.Op <= vm.PUSH32)) || (op != "PUSH" && op != ins.Op.String()) {
				ok = false
				break
			}
			pc += uint64(1 + len(ins.Data))
		}
		if ok {
			return true
		}
	}
	return false
}

// Match uses only a necessary opcode-prefix filter; complete symbolic binding
// and template-ID comparison still decide admission. Unmatched hints remain
// evidence for the offline miner, never automatic pruning authority.
func (m *BytecodeMatcher) Match(code []byte) (*MatchReport, error) {
	return m.match(code, len(m.known) > 0)
}

func MatchBytecode(code []byte, templates []Template) (*MatchReport, error) {
	m, err := NewBytecodeMatcher(templates)
	if err != nil {
		return nil, err
	}
	return m.match(code, false)
}

// MatchWithYield preserves binding and evidence; the callback only schedules
// worker analysis between candidate entries and never receives EVM state.
func (m *BytecodeMatcher) MatchWithYield(code []byte, yield func()) (*MatchReport, error) {
	return m.match(code, len(m.known) > 0, yield)
}
func (m *BytecodeMatcher) match(code []byte, filter bool, yields ...func()) (*MatchReport, error) {
	var pause func()
	if len(yields) > 0 {
		pause = yields[0]
	}
	instructions, err := Disassemble(code)
	if err != nil {
		return nil, err
	}
	p := &Program{Code: code, Instructions: instructions, ByPC: map[uint64]Instruction{}}
	for _, ins := range instructions {
		p.ByPC[ins.PC] = ins
	}
	report := &MatchReport{SchemaVersion: 1, RuntimeCodeHash: crypto.Keccak256Hash(code).Hex(), Mode: "offline_bindings_not_enabled", Bindings: []RuntimeBinding{}}
	for _, ins := range instructions {
		if ins.Op != vm.JUMPDEST {
			continue
		}
		if pause != nil {
			pause()
		}
		var legacyCandidate *candidate
		for _, semantics := range []string{"", legacySemantics, legacySignedSemantics} {
			if !m.models[semantics] {
				continue
			}
			if filter && !m.possible(p, ins.PC, semantics) {
				if semantics == "" && len(report.Candidates) < 64 && suspiciousArithmeticWindow(p, ins.PC) {
					report.Candidates = append(report.Candidates, CandidateHint{Entry: ins.PC, Reason: "unmatched_entry_prefix"})
				}
				continue
			}
			var c *candidate
			if isLegacySemantics(semantics) {
				if legacyCandidate == nil {
					legacyCandidate = discoverLegacyCandidate(p, ins.PC)
				}
				c = legacyCandidate
			} else {
				c = discoverCandidate(p, ins.PC)
			}
			if c.err != "" {
				if semantics == "" && len(report.Candidates) < 64 && suspiciousArithmeticWindow(p, ins.PC) {
					report.Candidates = append(report.Candidates, CandidateHint{Entry: ins.PC, Reason: "unresolved_arithmetic_region: " + c.err})
				}
				continue
			}
			if c.analysis.Semantics != semantics {
				continue
			}
			a, tpl, addresses := c.analysis, c.template, c.bindings
			if _, ok := m.known[tpl.ID]; !ok {
				report.Candidates = append(report.Candidates, CandidateHint{Entry: ins.PC, Reason: "unknown_closed_region", TemplateID: tpl.ID})
				continue
			}

			offsets := []int{inputOffset(a, a.OperandOrder[0]), inputOffset(a, a.OperandOrder[1])}
			binding := RuntimeBinding{Analysis: a, ContinuationOffset: inputOffset(a, "continuation"), InputStack: analysisInputs(a), Semantics: a.Semantics, TemplateID: tpl.ID, Entry: ins.PC, Operation: a.Operation, IntegerType: tpl.IntegerType, OperandOffsets: offsets, AddressBindings: addresses}
			pcs := append([]uint64(nil), a.ReachablePCs...)
			sort.Slice(pcs, func(i, j int) bool { return pcs[i] < pcs[j] })
			for _, pc := range pcs {
				i := p.ByPC[pc]
				binding.Instructions = append(binding.Instructions, BoundInstruction{pc, i.Op.String(), hex.EncodeToString(i.Data)})
			}
			report.Bindings = append(report.Bindings, binding)
		}
	}
	return report, nil
}

// A cheap evidence filter only. Business branches also satisfy it, so these
// hints NEVER produce bindings or pruning authority.
func suspiciousArithmeticWindow(p *Program, entry uint64) bool {
	pc := entry
	arithmetic := false
	for n := 0; n < 32; n++ {
		i, ok := p.ByPC[pc]
		if !ok {
			return false
		}
		switch i.Op {
		case vm.ADD, vm.SUB, vm.MUL, vm.DIV, vm.MOD, vm.SDIV, vm.SMOD:
			arithmetic = true
		case vm.JUMPI:
			return arithmetic
		case vm.JUMP, vm.STOP, vm.RETURN, vm.REVERT, vm.INVALID:
			return false
		}
		pc += uint64(1 + len(i.Data))
	}
	return false
}
