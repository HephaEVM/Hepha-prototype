package templateminer

import (
	"encoding/json"
	"fmt"
	"math/big"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/crypto"
)

type IntegerStudy struct {
	SchemaVersion int                   `json:"schemaVersion"`
	Mode          string                `json:"mode"`
	Programs      []IntegerStudyProgram `json:"programs"`
	Templates     []Template            `json:"templates"`
	Limitations   []string              `json:"limitations"`
}
type IntegerStudyProgram struct {
	Artifact        string             `json:"artifact"`
	CompilerVersion string             `json:"compilerVersion"`
	Settings        json.RawMessage    `json:"settings"`
	Contract        string             `json:"contract"`
	CodeHash        string             `json:"codeHash"`
	Cases           []IntegerStudyCase `json:"cases"`
}
type IntegerStudyCase struct {
	Method          string               `json:"method"`
	SourceType      string               `json:"sourceType"`
	Operation       string               `json:"operation"`
	NegativeControl bool                 `json:"negativeControl"`
	ActualEVMChecks int                  `json:"actualEvmChecks"`
	Outcomes        map[string]int       `json:"actualOutcomes"`
	Regions         []IntegerStudyRegion `json:"regions"`
	Rejected        map[string]int       `json:"rejectedCandidates,omitempty"`
}
type IntegerStudyRegion struct {
	TemplateID  string            `json:"templateId"`
	Entry       uint64            `json:"entry"`
	IntegerType string            `json:"integerType"`
	Bindings    map[string]uint64 `json:"addressBindings"`
	Executions  int               `json:"executions"`
	Outcomes    map[string]int    `json:"outcomes"`
	Witnesses   []ChainWitness    `json:"sampleWitnesses"`
}

func frameFromTrace(t *Trace) (*chainFrame, AuditRecord) {
	f := &chainFrame{}
	for i, s := range t.Steps {
		r := AuditRecord{Step: i, PC: s.PC, Op: s.Op, Memory: s.Memory}
		for _, v := range s.Stack {
			r.Stack = append(r.Stack, v.Value)
		}
		f.ops = append(f.ops, r)
	}
	return f, AuditRecord{Output: t.Return, Reverted: t.ExecutionError != "", Error: t.ExecutionError}
}

type witnessedRegion struct {
	c *candidate
	w *ChainWitness
}

func mineTraceRegions(p *Program, tr *Trace, cache map[uint64]*candidate) ([]witnessedRegion, map[string]int) {
	return mineTraceRegionsModel(p, tr, cache, false)
}

func mineTraceRegionsModel(p *Program, tr *Trace, cache map[uint64]*candidate, legacy bool) ([]witnessedRegion, map[string]int) {
	f, end := frameFromTrace(tr)
	rejected := map[string]int{}
	var found []witnessedRegion
	for i, s := range f.ops {
		if s.Op != "JUMPDEST" || len(s.Stack) < 3 || (i > 0 && f.ops[i-1].Op != "JUMP") {
			continue
		}
		c := cache[s.PC]
		if c == nil {
			if legacy {
				c = discoverLegacyCandidate(p, s.PC)
			} else {
				c = discoverCandidate(p, s.PC)
			}
			cache[s.PC] = c
		}
		if c.err != "" {
			rejected[c.err]++
			continue
		}
		w, err := chainWitness(p, f, end, i, c)
		if err != nil {
			rejected["witness: "+err.Error()]++
			continue
		}
		found = append(found, witnessedRegion{c, w})
	}
	var kept []witnessedRegion
	for _, r := range found {
		nested := false
		for _, other := range found {
			if other.c.helper.Operation == r.c.helper.Operation && other.w.EntryStep >= r.w.EntryStep && other.w.EntryStep+other.w.InstructionCount <= r.w.EntryStep+r.w.InstructionCount && other.w.InstructionCount < r.w.InstructionCount {
				nested = true
				break
			}
		}
		if !nested {
			kept = append(kept, r)
		}
	}
	return kept, rejected
}

func actualStudyPairs(t integerType) [][2]*big.Int {
	lo, hi := t.bounds()
	out := [][2]*big.Int{{big.NewInt(7), big.NewInt(2)}, {big.NewInt(0), big.NewInt(0)}, {hi, big.NewInt(1)}, {hi, big.NewInt(2)}, {lo, big.NewInt(1)}}
	if t.signed {
		out = append(out, [2]*big.Int{lo, big.NewInt(-1)}, [2]*big.Int{hi, big.NewInt(-1)}, [2]*big.Int{big.NewInt(-7), big.NewInt(2)}, [2]*big.Int{big.NewInt(7), big.NewInt(-2)}, [2]*big.Int{big.NewInt(-7), big.NewInt(-2)})
	}
	return out
}

// StudyIntegerCorpus uses ABI types ONLY to construct inputs and the independent
// expected answer. Candidate discovery receives deployed bytecode and runtime
// stacks, never function-debug entries, expected widths, or manual templates.
func StudyIntegerCorpus(manifest string) (*IntegerStudy, error) {
	data, err := os.ReadFile(manifest)
	if err != nil {
		return nil, err
	}
	var entries []struct{ Artifact, Source, Contract string }
	if err = json.Unmarshal(data, &entries); err != nil {
		return nil, err
	}
	if len(entries) == 0 {
		return nil, fmt.Errorf("empty integer corpus")
	}
	r := &IntegerStudy{SchemaVersion: 1, Mode: "controlled_compiler_corpus_not_mainnet", Templates: []Template{}, Limitations: []string{"Synthetic source and calls; actual EVM execution, not historical transactions.", "Only outlined binary regions are discovered. Inferred bytecode domain can be wider than the source type if the caller performs narrowing.", "Finite specification checks and concrete executions are not full-input equivalence proofs. No runtime pruning enabled."}}
	known := map[string]Template{}
	for _, entry := range entries {
		file := entry.Artifact
		if !filepath.IsAbs(file) {
			file = filepath.Join(filepath.Dir(manifest), file)
		}
		p, err := Load(file, entry.Source, entry.Contract)
		if err != nil {
			return nil, err
		}
		parsed, err := abi.JSON(strings.NewReader(string(p.Contract.ABI)))
		if err != nil {
			return nil, err
		}
		program := IntegerStudyProgram{Artifact: entry.Artifact, CompilerVersion: p.Artifact.CompilerVersion, Settings: p.Artifact.Settings, Contract: entry.Contract, CodeHash: crypto.Keccak256Hash(p.Code).Hex()}
		cache := map[uint64]*candidate{}
		for _, name := range sortedKeys(parsed.Methods) {
			m := parsed.Methods[name]
			if len(m.Inputs) != 2 || m.Inputs[0].Type.String() != m.Inputs[1].Type.String() {
				continue
			}
			t, err := parseIntegerType(m.Inputs[0].Type.String())
			if err != nil {
				continue
			}
			op := strings.ToUpper(strings.Split(name, "_")[0])
			negative := op == "UNCHECKEDADD"
			if negative {
				op = "ADD"
			}
			if t.signed && (op == "DIV" || op == "MOD") {
				op = "S" + op
			}
			if !binaryOperation(op) {
				continue
			}
			c := IntegerStudyCase{Method: name, SourceType: t.name(), Operation: op, NegativeControl: negative, Outcomes: map[string]int{}, Regions: []IntegerStudyRegion{}, Rejected: map[string]int{}}
			regions := map[string]*IntegerStudyRegion{}
			for _, pair := range actualStudyPairs(t) {
				input, err := EncodeCall(p, name, []string{pair[0].String(), pair[1].String()})
				if err != nil {
					return nil, err
				}
				tr := Replay(p, input)
				if tr.TrackingError != "" {
					return nil, fmt.Errorf("%s: %s", name, tr.TrackingError)
				}
				exact, outcome, err := arithmeticSpec(op, t, pair[0], pair[1])
				if err != nil {
					return nil, err
				}
				if negative {
					outcome = "return"
					exact.Mod(exact, new(big.Int).Lsh(big.NewInt(1), uint(t.bits)))
					if t.signed && exact.Bit(t.bits-1) != 0 {
						exact.Sub(exact, new(big.Int).Lsh(big.NewInt(1), uint(t.bits)))
					}
				}
				if outcome == "return" {
					got, ok := new(big.Int).SetString(tr.Return, 16)
					if tr.ExecutionError != "" || !ok || got.Cmp(new(big.Int).Mod(exact, modulus)) != 0 {
						return nil, fmt.Errorf("%s %s actual EVM result disagrees with independent oracle (%s,%s): %s %s", entry.Artifact, name, pair[0], pair[1], tr.ExecutionError, tr.Return)
					}
				} else if tr.ExecutionError == "" || panicOutcome(tr.Return) != outcome {
					return nil, fmt.Errorf("%s %s actual EVM panic disagrees with oracle", entry.Artifact, name)
				}
				c.ActualEVMChecks++
				c.Outcomes[outcome]++
				found, rejected := mineTraceRegions(p, tr, cache)
				for reason, n := range rejected {
					c.Rejected[reason] += n
				}
				if negative && len(found) > 0 {
					return nil, fmt.Errorf("unchecked negative control %s was accepted", name)
				}
				for _, f := range found {
					if f.c.helper.Operation != op {
						return nil, fmt.Errorf("%s discovered wrong operation %s", name, f.c.helper.Operation)
					}
					key := fmt.Sprintf("%d", f.c.helper.Entry)
					i := regions[key]
					if i == nil {
						i = &IntegerStudyRegion{TemplateID: f.c.template.ID, Entry: f.c.helper.Entry, IntegerType: f.c.analysis.IntegerType, Bindings: f.c.bindings, Outcomes: map[string]int{}}
						regions[key] = i
					}
					i.Executions++
					i.Outcomes[f.w.Outcome]++
					if len(i.Witnesses) < 3 {
						i.Witnesses = append(i.Witnesses, *f.w)
					}
					known[f.c.template.ID] = f.c.template
				}
			}
			for _, key := range sortedKeys(regions) {
				c.Regions = append(c.Regions, *regions[key])
			}
			program.Cases = append(program.Cases, c)
		}
		r.Programs = append(r.Programs, program)
	}
	for _, id := range sortedKeys(known) {
		r.Templates = append(r.Templates, known[id])
	}
	sort.Slice(r.Programs, func(i, j int) bool {
		return r.Programs[i].Artifact+r.Programs[i].Contract < r.Programs[j].Artifact+r.Programs[j].Contract
	})
	return r, nil
}
