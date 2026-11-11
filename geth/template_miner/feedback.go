package templateminer

import (
	"encoding/hex"
	"fmt"
	"sort"
	"strings"

	"github.com/ethereum/go-ethereum/crypto"
)

// Feedback is evidence of encountered CODE, not proof that every candidate was
// executed, nor authorization to enable its template. Publication stays explicit.
type FeedbackCode struct {
	CodeHash       string          `json:"codeHash"`
	Bytecode       string          `json:"bytecode"`
	LibraryVersion string          `json:"libraryVersion"`
	Candidates     []CandidateHint `json:"candidates"`
}
type FeedbackBundle struct {
	SchemaVersion int            `json:"schemaVersion"`
	Mode          string         `json:"mode"`
	Codes         []FeedbackCode `json:"codes"`
	Dropped       uint64         `json:"dropped"`
}
type FeedbackMining struct {
	SchemaVersion int                 `json:"schemaVersion"`
	Mode          string              `json:"mode"`
	Templates     []Template          `json:"templates"`
	Candidates    []FeedbackCandidate `json:"candidates"`
}
type FeedbackCandidate struct {
	CodeHash   string `json:"codeHash"`
	Entry      uint64 `json:"entry"`
	TemplateID string `json:"templateId,omitempty"`
	Reason     string `json:"reason,omitempty"`
}

// MineFeedback reruns discovery from the retained bytecode. Hints and supplied
// template IDs are never trusted. The result is an OFFLINE candidate bundle;
// a user must explicitly publish/load it as a new immutable runtime library.
func MineFeedback(feedback FeedbackBundle) (*FeedbackMining, error) {
	if feedback.SchemaVersion != 1 {
		return nil, fmt.Errorf("unsupported feedback schema")
	}
	out := &FeedbackMining{SchemaVersion: 1, Mode: "offline_feedback_candidates_not_enabled", Templates: []Template{}}
	known := map[string]bool{}
	blocked := map[string]bool{}
	for _, e := range feedback.Codes {
		code, err := hex.DecodeString(strings.TrimPrefix(e.Bytecode, "0x"))
		if err != nil {
			return nil, err
		}
		if len(code) > 1<<20 || crypto.Keccak256Hash(code).Hex() != e.CodeHash {
			return nil, fmt.Errorf("invalid feedback code hash/size")
		}
		ins, err := Disassemble(code)
		if err != nil {
			return nil, err
		}
		p := &Program{Code: code, Instructions: ins, ByPC: map[uint64]Instruction{}}
		for _, i := range ins {
			p.ByPC[i.PC] = i
		}
		seen := map[uint64]bool{}
		for _, hint := range e.Candidates {
			if seen[hint.Entry] {
				continue
			}
			seen[hint.Entry] = true
			c := discoverCandidate(p, hint.Entry)
			if c.err != "" {
				c = discoverLegacyCandidate(p, hint.Entry)
			}
			record := FeedbackCandidate{CodeHash: e.CodeHash, Entry: hint.Entry, Reason: c.err}
			if c.err == "" {
				record.TemplateID = c.template.ID
				// Check ALL observations for this entry, even when the static hint
				// appeared first. One counterexample excludes the whole family.
				for _, observed := range e.Candidates {
					if observed.Entry != hint.Entry || len(observed.Operands) == 0 {
						continue
					}
					if len(observed.Operands) != 2 {
						return nil, fmt.Errorf("invalid counterexample operands")
					}
					a, err := word(observed.Operands[0])
					if err != nil {
						return nil, err
					}
					b, err := word(observed.Operands[1])
					if err != nil {
						return nil, err
					}
					l, r := a, b
					if c.analysis.OperandOrder[0] != "a" {
						l, r = r, l
					}
					if err := CheckInputs(c.analysis, c.analysis.Operation, l, r); err != nil {
						blocked[c.template.ID] = true
						record.Reason = "runtime_counterexample: " + err.Error()
					}
				}
				if !known[c.template.ID] && !blocked[c.template.ID] {
					out.Templates = append(out.Templates, c.template)
					known[c.template.ID] = true
				}
			}
			out.Candidates = append(out.Candidates, record)
		}
	}
	valid := out.Templates[:0]
	for _, tpl := range out.Templates {
		if !blocked[tpl.ID] {
			valid = append(valid, tpl)
		}
	}
	out.Templates = valid
	sort.Slice(out.Templates, func(i, j int) bool { return out.Templates[i].ID < out.Templates[j].ID })
	return out, nil
}
