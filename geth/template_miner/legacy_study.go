package templateminer

import (
	"encoding/json"
	"fmt"
	"math/big"
	"os"
	"path/filepath"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/crypto"
)

func legacyActualOutcome(tr *Trace) string {
	if tr.ExecutionError == "" {
		return "return"
	}
	if strings.Contains(strings.ToLower(tr.ExecutionError), "invalid opcode") {
		return "invalid"
	}
	if tr.ExecutionError == "execution reverted" {
		return "revert"
	}
	return "unsupported_execution_failure"
}

// StudyLegacyCorpus compiles no rules from function names: ABI/name information
// only supplies the independent experiment oracle. The miner gets bytecode and
// execution stacks and must discover its own operation and complete boundary.
func StudyLegacyCorpus(manifest string) (*IntegerStudy, error) {
	data, err := os.ReadFile(manifest)
	if err != nil {
		return nil, err
	}
	var entries []struct{ Artifact, Source, Contract, Profile string }
	if err = json.Unmarshal(data, &entries); err != nil {
		return nil, err
	}
	if len(entries) == 0 {
		return nil, fmt.Errorf("empty SafeMath corpus")
	}
	r := &IntegerStudy{SchemaVersion: 1, Mode: "controlled_pre08_safemath_corpus_not_mainnet", Templates: []Template{}, Limitations: []string{
		"Published historical library sources with synthetic harnesses and calls; separate from real mainnet replay.",
		"uint256/int256 outlined binary helpers, including an errorMessage pointer argument. tryMath, inline checks and successful memory-writing wrappers remain unsupported. Error encoders are retained and finitely validated, not universally proven.",
		"Both success and failure paths are retained, including INVALID versus REVERT, fixed bytes and explicitly masked caller-memory ABI padding. Original failure execution is required. No runtime pruning is enabled.",
		"228 finite specification checks per accepted helper plus actual EVM executions are not a full-input equivalence proof.",
	}}
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
		if !legacyCompiler(p.Artifact.CompilerVersion) {
			return nil, fmt.Errorf("SafeMath corpus requires pre-0.8 compiler")
		}
		parsed, err := abi.JSON(strings.NewReader(string(p.Contract.ABI)))
		if err != nil {
			return nil, err
		}
		program := IntegerStudyProgram{Artifact: entry.Artifact, CompilerVersion: p.Artifact.CompilerVersion, Settings: p.Artifact.Settings, Contract: entry.Contract, CodeHash: crypto.Keccak256Hash(p.Code).Hex()}
		cache := map[uint64]*candidate{}
		for _, name := range sortedKeys(parsed.Methods) {
			method := parsed.Methods[name]
			if len(method.Inputs) != 2 || (method.Inputs[0].Type.String() != "uint256" && method.Inputs[0].Type.String() != "int256") || method.Inputs[1].Type.String() != method.Inputs[0].Type.String() {
				continue
			}
			typ, _ := parseIntegerType(method.Inputs[0].Type.String())
			op := strings.ToUpper(strings.Split(name, "_")[0])
			if typ.signed && (op == "DIV" || op == "MOD") {
				op = "S" + op
			}
			negative := op == "UNCHECKEDADD" || op == "MISLEADINGADD"
			if negative {
				op = "ADD"
			}
			if !binaryOperation(op) {
				continue
			}
			c := IntegerStudyCase{Method: name, SourceType: typ.name(), Operation: op, NegativeControl: negative, Outcomes: map[string]int{}, Regions: []IntegerStudyRegion{}, Rejected: map[string]int{}}
			regions := map[string]*IntegerStudyRegion{}
			pairs := typedValidationPairs(typ)
			if negative {
				pairs = append(actualStudyPairs(typ), [2]*big.Int{big.NewInt(11), big.NewInt(0)})
			}
			for _, pair := range pairs {
				input, err := EncodeCall(p, name, []string{pair[0].String(), pair[1].String()})
				if err != nil {
					return nil, err
				}
				tr := Replay(p, input)
				if tr.TrackingError != "" {
					return nil, fmt.Errorf("%s %s: %s", entry.Artifact, name, tr.TrackingError)
				}
				actual := legacyActualOutcome(tr)
				if actual == "unsupported_execution_failure" {
					return nil, fmt.Errorf("unexpected EVM failure: %s", tr.ExecutionError)
				}
				if !negative {
					exact, want, err := arithmeticSpec(op, typ, pair[0], pair[1])
					if err != nil {
						return nil, err
					}
					if want == "return" {
						got, ok := new(big.Int).SetString(tr.Return, 16)
						if actual != "return" || !ok || got.Cmp(new(big.Int).Mod(exact, modulus)) != 0 {
							return nil, fmt.Errorf("legacy actual result disagrees with oracle: %s %s (%s,%s)", entry.Artifact, name, pair[0], pair[1])
						}
					} else if actual == "return" {
						return nil, fmt.Errorf("missing actual failure: %s %s", entry.Artifact, name)
					}
				}
				c.ActualEVMChecks++
				c.Outcomes[actual]++
				found, rejected := mineTraceRegionsModel(p, tr, cache, true)
				for reason, n := range rejected {
					c.Rejected[reason] += n
				}
				if negative && len(found) > 0 {
					return nil, fmt.Errorf("negative control incorrectly accepted: %s %s", entry.Artifact, name)
				}
				for _, f := range found {
					if f.c.helper.Operation != op || f.c.analysis.IntegerType != typ.name() {
						return nil, fmt.Errorf("wrong mined operation for %s", name)
					}
					key := fmt.Sprint(f.c.helper.Entry)
					inst := regions[key]
					if inst == nil {
						inst = &IntegerStudyRegion{TemplateID: f.c.template.ID, Entry: f.c.helper.Entry, IntegerType: f.c.analysis.IntegerType, Bindings: f.c.bindings, Outcomes: map[string]int{}}
						regions[key] = inst
					}
					if len(inst.Witnesses) < 3 || inst.Outcomes[f.w.Outcome] == 0 {
						inst.Witnesses = append(inst.Witnesses, *f.w)
					}
					inst.Executions++
					inst.Outcomes[f.w.Outcome]++
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
	return r, nil
}
