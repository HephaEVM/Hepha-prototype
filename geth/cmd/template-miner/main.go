package main

import (
	"encoding/hex"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	miner "go-eth1.16.5-evm/template_miner"
)

type scenario struct {
	Name   string   `json:"name"`
	Method string   `json:"method"`
	Args   []string `json:"args"`
	Input  string   `json:"input,omitempty"`
}
type corpusEntry struct {
	Artifact  string `json:"artifact"`
	Source    string `json:"source"`
	Contract  string `json:"contract"`
	Scenarios string `json:"scenarios"`
}
type corpusBundle struct {
	SchemaVersion int              `json:"schemaVersion"`
	Mode          string           `json:"mode"`
	Templates     []miner.Template `json:"templates"`
	Programs      []*miner.Bundle  `json:"programs"`
}

func main() {
	artifact := flag.String("artifact", "", "compiler artifact from compile.cjs")
	feedbackPath := flag.String("feedback", "", "runtime feedback JSON; rediscover offline candidates, never auto-enable")
	source := flag.String("source", "Arithmetic.sol", "contract source key")
	contract := flag.String("contract", "Arithmetic", "contract name")
	scenarios := flag.String("scenarios", "", "JSON array of {name,method,args} or {name,input} calls")
	output := flag.String("out", "", "output JSON bundle/report")
	tracesPath := flag.String("traces", "", "optional annotated execution-trace JSON")
	codePath := flag.String("bytecode", "", "runtime hex file for bytecode-only matching")
	templatesPath := flag.String("templates", "", "learned bundle for bytecode-only matching")
	corpusPath := flag.String("corpus", "", "JSON manifest for batch mining and template-family deduplication")
	auditDir := flag.String("audit-dir", "", "completed optimizer-audit replay directory for real-transaction mining")
	labelsPath := flag.String("labels", "", "public compiler-setting evidence manifest for mainnet mining")
	integerStudy := flag.String("integer-study", "", "controlled compiler corpus manifest: all integer widths and five operations")
	legacyStudy := flag.String("safemath-study", "", "controlled pre-0.8 SafeMath corpus manifest")
	abiStudy := flag.String("abi-study", "", "comma-separated compiler-artifact globs for ABI decode/encode region templates")
	abiCodeDir := flag.String("abi-code-dir", "", "optional deployed .hex corpus to bind with -abi-study")
	model := flag.String("arithmetic-model", "auto", "audit discovery: auto, solidity08, safemath (also permits compiler-unknown runtimes)")
	flag.Parse()
	fail := func(err error) { fmt.Fprintln(os.Stderr, err); os.Exit(1) }
	if *output == "" {
		fail(fmt.Errorf("-out is required"))
	}
	if *abiStudy != "" {
		if *feedbackPath != "" || *artifact != "" || *codePath != "" || *auditDir != "" || *corpusPath != "" || *integerStudy != "" || *legacyStudy != "" || *templatesPath != "" {
			fail(fmt.Errorf("-abi-study cannot be mixed with other mining modes"))
		}
		if err := runABIStudy(*abiStudy, *abiCodeDir, *output); err != nil {
			fail(err)
		}
		return
	}
	if *abiCodeDir != "" {
		fail(fmt.Errorf("-abi-code-dir requires -abi-study"))
	}
	if *feedbackPath != "" {
		if *artifact != "" || *codePath != "" || *auditDir != "" || *corpusPath != "" || *integerStudy != "" || *legacyStudy != "" {
			fail(fmt.Errorf("-feedback cannot be mixed with other mining modes"))
		}
		data, err := os.ReadFile(*feedbackPath)
		if err != nil {
			fail(err)
		}
		var evidence miner.FeedbackBundle
		if err = json.Unmarshal(data, &evidence); err != nil {
			fail(err)
		}
		report, err := miner.MineFeedback(evidence)
		if err != nil {
			fail(err)
		}
		if *templatesPath != "" {
			data, err := os.ReadFile(*templatesPath)
			if err != nil {
				fail(err)
			}
			var previous struct {
				SchemaVersion int              `json:"schemaVersion"`
				Templates     []miner.Template `json:"templates"`
			}
			if err = json.Unmarshal(data, &previous); err != nil {
				fail(err)
			}
			if previous.SchemaVersion != 1 {
				fail(fmt.Errorf("unsupported previous library schema"))
			}
			if _, err = miner.MatchBytecode(nil, previous.Templates); err != nil {
				fail(err)
			}
			seen, blocked := map[string]bool{}, map[string]bool{}
			for _, tpl := range report.Templates {
				seen[tpl.ID] = true
			}
			for _, c := range report.Candidates {
				if strings.HasPrefix(c.Reason, "runtime_counterexample:") {
					blocked[c.TemplateID] = true
				}
			}
			for _, tpl := range previous.Templates {
				if !seen[tpl.ID] && !blocked[tpl.ID] {
					report.Templates = append(report.Templates, tpl)
					seen[tpl.ID] = true
				}
			}
			sort.Slice(report.Templates, func(i, j int) bool { return report.Templates[i].ID < report.Templates[j].ID })
		}
		if err = save(*output, report); err != nil {
			fail(err)
		}
		fmt.Printf("feedback: %d candidates / %d offline template families -> %s\n", len(report.Candidates), len(report.Templates), *output)
		return
	}
	if *legacyStudy != "" {
		if *auditDir != "" || *artifact != "" || *corpusPath != "" || *codePath != "" || *integerStudy != "" {
			fail(fmt.Errorf("-safemath-study cannot be mixed with other modes"))
		}
		r, err := miner.StudyLegacyCorpus(*legacyStudy)
		if err != nil {
			fail(err)
		}
		if err := save(*output, r); err != nil {
			fail(err)
		}
		fmt.Printf("SafeMath corpus: %d programs / %d families -> %s\n", len(r.Programs), len(r.Templates), *output)
		return
	}
	if *integerStudy != "" {
		if *auditDir != "" || *artifact != "" || *corpusPath != "" || *codePath != "" || *templatesPath != "" {
			fail(fmt.Errorf("-integer-study cannot be mixed with other input modes"))
		}
		report, err := miner.StudyIntegerCorpus(*integerStudy)
		if err != nil {
			fail(err)
		}
		if err = save(*output, report); err != nil {
			fail(err)
		}
		fmt.Printf("controlled integer study: %d programs / %d templates -> %s\n", len(report.Programs), len(report.Templates), *output)
		return
	}
	if *auditDir != "" {
		if *artifact != "" || *corpusPath != "" || *codePath != "" || *templatesPath != "" {
			fail(fmt.Errorf("-audit-dir cannot be mixed with artifact or template matching inputs"))
		}
		report, err := miner.MineAuditWithModel(*auditDir, *labelsPath, *model)
		if err != nil {
			fail(err)
		}
		if err = save(*output, report); err != nil {
			fail(err)
		}
		covered, executions := 0, 0
		for _, c := range report.Contracts {
			if len(c.Instances) > 0 {
				covered++
			}
			for _, i := range c.Instances {
				executions += i.Executions
			}
		}
		fmt.Printf("mainnet: %d validated blocks / %d contracts / %d contracts with supported regions / %d templates / %d observed executions -> %s\n", report.ValidatedBlocks, len(report.Contracts), covered, len(report.Templates), executions, *output)
		return
	}
	if *corpusPath != "" {
		raw, err := os.ReadFile(*corpusPath)
		if err != nil {
			fail(err)
		}
		var entries []corpusEntry
		if err = json.Unmarshal(raw, &entries); err != nil {
			fail(err)
		}
		if len(entries) == 0 {
			fail(fmt.Errorf("empty corpus"))
		}
		result := corpusBundle{SchemaVersion: 1, Mode: "offline_candidates_not_enabled"}
		known := map[string]bool{}
		instances := 0
		allTraces := map[string][]miner.NamedTrace{}
		for _, entry := range entries {
			bundle, traces, err := mineProgram(entry)
			if err != nil {
				fail(fmt.Errorf("%s: %w", entry.Artifact, err))
			}
			result.Programs = append(result.Programs, bundle)
			instances += len(bundle.Instances)
			allTraces[entry.Artifact+":"+entry.Contract] = traces
			for _, tpl := range bundle.Templates {
				if !known[tpl.ID] {
					known[tpl.ID] = true
					result.Templates = append(result.Templates, tpl)
				}
			}
		}
		if err = save(*output, result); err != nil {
			fail(err)
		}
		if *tracesPath != "" {
			if err = save(*tracesPath, allTraces); err != nil {
				fail(err)
			}
		}
		fmt.Printf("corpus: %d programs / %d complete helper instances / %d deduplicated families -> %s (offline only)\n", len(result.Programs), instances, len(result.Templates), *output)
		return
	}
	if *codePath != "" {
		if *templatesPath == "" {
			fail(fmt.Errorf("-templates is required with -bytecode"))
		}
		raw, err := os.ReadFile(*codePath)
		if err != nil {
			fail(err)
		}
		code, err := hex.DecodeString(strings.TrimPrefix(strings.TrimSpace(string(raw)), "0x"))
		if err != nil {
			fail(err)
		}
		raw, err = os.ReadFile(*templatesPath)
		if err != nil {
			fail(err)
		}
		var bundle miner.Bundle
		if err = json.Unmarshal(raw, &bundle); err != nil {
			fail(err)
		}
		if bundle.SchemaVersion != 1 {
			fail(fmt.Errorf("unsupported template schema"))
		}
		report, err := miner.MatchBytecode(code, bundle.Templates)
		if err != nil {
			fail(err)
		}
		if err = save(*output, report); err != nil {
			fail(err)
		}
		fmt.Printf("matched %d complete helper regions -> %s (offline only)\n", len(report.Bindings), *output)
		return
	}
	if *artifact == "" || *scenarios == "" {
		fail(fmt.Errorf("-artifact and -scenarios are required for mining"))
	}
	bundle, traces, err := mineProgram(corpusEntry{*artifact, *source, *contract, *scenarios})
	if err != nil {
		fail(err)
	}
	if err = save(*output, bundle); err != nil {
		fail(err)
	}
	if *tracesPath != "" {
		if err = save(*tracesPath, traces); err != nil {
			fail(err)
		}
	}
	fmt.Printf("mined %d families / %d helper instances from %d calls -> %s (offline only)\n", len(bundle.Templates), len(bundle.Instances), len(traces), *output)
	for _, d := range bundle.Diagnostics {
		fmt.Printf("coverage: %s\n", d)
	}
}
func mineProgram(entry corpusEntry) (*miner.Bundle, []miner.NamedTrace, error) {
	p, err := miner.Load(entry.Artifact, entry.Source, entry.Contract)
	if err != nil {
		return nil, nil, err
	}
	raw, err := os.ReadFile(entry.Scenarios)
	if err != nil {
		return nil, nil, err
	}
	var calls []scenario
	if err = json.Unmarshal(raw, &calls); err != nil {
		return nil, nil, err
	}
	if len(calls) == 0 {
		return nil, nil, fmt.Errorf("empty scenario list")
	}
	var traces []miner.NamedTrace
	for i, c := range calls {
		name := c.Name
		if name == "" {
			name = fmt.Sprintf("call-%d", i)
		}
		var input []byte
		if c.Input != "" {
			input, err = hex.DecodeString(strings.TrimPrefix(c.Input, "0x"))
		} else {
			input, err = miner.EncodeCall(p, c.Method, c.Args)
		}
		if err != nil {
			return nil, nil, fmt.Errorf("%s: %w", name, err)
		}
		trace := miner.Replay(p, input)
		if trace.TrackingError != "" {
			return nil, nil, fmt.Errorf("%s: %s", name, trace.TrackingError)
		}
		traces = append(traces, miner.NamedTrace{Name: name, Trace: trace})
	}
	bundle, err := miner.Mine(p, traces)
	return bundle, traces, err
}
func save(path string, v any) error {
	data, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return err
	}
	if err = os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return err
	}
	return os.WriteFile(path, append(data, '\n'), 0644)
}
