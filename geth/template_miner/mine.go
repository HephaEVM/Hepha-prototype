package templateminer

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"
	"sort"

	"github.com/ethereum/go-ethereum/crypto"
)

type NamedTrace struct {
	Name  string `json:"name"`
	Trace *Trace `json:"trace"`
}
type PatternInstruction struct {
	Role         string `json:"role"`
	Op           string `json:"op"`
	Literal      string `json:"literal,omitempty"`
	AddressRole  string `json:"addressRole,omitempty"`
	DeadConstant bool   `json:"deadConstant,omitempty"`
	CodeData     string `json:"codeData,omitempty"`
}
type PatternPath struct {
	FailureProgram  []PatternInstruction `json:"failureProgram,omitempty"`
	FailureArgument string               `json:"failureArgument,omitempty"`
	Instructions    []PatternInstruction `json:"instructions"`
	Conditions      []Condition          `json:"conditions"`
	Outcome         string               `json:"outcome"`
	Result          *Expr                `json:"result,omitempty"`
	PeakStack       int                  `json:"peakStack"`
	FailureData     string               `json:"failureData,omitempty"`
	FailureMask     string               `json:"failureDataKnownMask,omitempty"`
	FailureMemory   *Expr                `json:"failureMemoryOffset,omitempty"`
}
type Template struct {
	Semantics       string        `json:"semantics,omitempty"`
	FailureHandling string        `json:"failureHandling,omitempty"`
	ID              string        `json:"id"`
	Operation       string        `json:"operation"`
	IntegerType     string        `json:"integerType,omitempty"`
	InputStack      []string      `json:"inputStackBottomToTop"`
	OperandOrder    []string      `json:"operandOrder"`
	OutputStack     []string      `json:"outputStackBottomToTop"`
	Paths           []PatternPath `json:"paths"`
}
type BoundInstruction struct {
	PC   uint64 `json:"pc"`
	Op   string `json:"op"`
	Data string `json:"pushData,omitempty"`
}
type Instance struct {
	TemplateID      string             `json:"templateId"`
	Helper          Helper             `json:"helper"`
	AddressBindings map[string]uint64  `json:"addressBindings"`
	Instructions    []BoundInstruction `json:"instructions"`
	Analysis        *SymbolicAnalysis  `json:"analysis"`
	Witnesses       []Witness          `json:"witnesses"`
	Validation      Validation         `json:"validation"`
}
type Witness struct {
	Trace        string   `json:"trace"`
	Sites        []int    `json:"sourceSiteIds"`
	EntryStep    int      `json:"entryStep"`
	EndStep      int      `json:"endStepExclusive"`
	AnchorPC     uint64   `json:"arithmeticPc"`
	OperandIDs   []int    `json:"operandValueIds"`
	ResultID     int      `json:"resultValueId"`
	GuardPCs     []uint64 `json:"resultDependentGuardPcs"`
	Outcome      string   `json:"outcome"`
	Continuation uint64   `json:"continuationPc"`
}
type Validation struct {
	InputPairs       int    `json:"symbolicDifferentialInputPairs"`
	SuccessWitnesses int    `json:"successWitnesses"`
	PanicWitnesses   int    `json:"panicWitnesses"`
	Status           string `json:"status"`
}
type Bundle struct {
	SchemaVersion   int             `json:"schemaVersion"`
	CompilerVersion string          `json:"compilerVersion"`
	Settings        json.RawMessage `json:"compilerSettings"`
	Contract        string          `json:"contract"`
	RuntimeCodeHash string          `json:"runtimeCodeHash"`
	Mode            string          `json:"mode"`
	Sites           []Site          `json:"sites"`
	Templates       []Template      `json:"templates"`
	Instances       []Instance      `json:"instances"`
	Diagnostics     []string        `json:"diagnostics"`
}

func Mine(p *Program, traces []NamedTrace) (*Bundle, error) {
	b := &Bundle{SchemaVersion: 1, CompilerVersion: p.Artifact.CompilerVersion, Settings: p.Artifact.Settings, Contract: p.Source + ":" + p.Name, RuntimeCodeHash: crypto.Keccak256Hash(p.Code).Hex(), Mode: "offline_candidates_not_enabled", Sites: p.Sites, Templates: []Template{}, Instances: []Instance{}, Diagnostics: append([]string{}, p.Diagnostics...)}
	families := map[string]bool{}
	for _, helper := range p.Helpers {
		a, err := Analyze(p, helper)
		if err != nil {
			b.Diagnostics = append(b.Diagnostics, helper.Name+": "+err.Error())
			continue
		}
		var witnesses []Witness
		for _, nt := range traces {
			if nt.Trace.TrackingError != "" {
				return nil, fmt.Errorf("trace %s: %s", nt.Name, nt.Trace.TrackingError)
			}
			w, err := extractWitnesses(p, helper, nt)
			if err != nil {
				return nil, err
			}
			witnesses = append(witnesses, w...)
		}
		// Metadata alone is insufficient: require source-anchored runtime
		// evidence before emitting a family. Unsupported/unexecuted sites stay
		// visible in the report instead of being silently called covered.
		anchored := false
		for _, w := range witnesses {
			if len(w.Sites) > 0 && w.Outcome == "return" {
				anchored = true
			}
		}
		if !anchored {
			b.Diagnostics = append(b.Diagnostics, helper.Name+": no source-anchored successful execution witness")
			continue
		}
		pairs := validationPairs()
		for _, pair := range pairs {
			if err := CheckInputs(a, helper.Operation, pair[0], pair[1]); err != nil {
				return nil, fmt.Errorf("%s differential validation: %w", helper.Name, err)
			}
		}
		tpl, bindings, err := normalize(p, helper, a)
		if err != nil {
			return nil, err
		}
		inst := Instance{TemplateID: tpl.ID, Helper: helper, AddressBindings: bindings, Analysis: a, Witnesses: witnesses, Validation: Validation{InputPairs: len(pairs), Status: "closed_control_flow_and_finite_differential_evidence_not_formal_equivalence"}}
		pcs := append([]uint64(nil), a.ReachablePCs...)
		sort.Slice(pcs, func(i, j int) bool { return pcs[i] < pcs[j] })
		for _, pc := range pcs {
			ins := p.ByPC[pc]
			inst.Instructions = append(inst.Instructions, BoundInstruction{pc, ins.Op.String(), hex.EncodeToString(ins.Data)})
		}
		for _, w := range witnesses {
			if w.Outcome == "return" {
				inst.Validation.SuccessWitnesses++
			} else {
				inst.Validation.PanicWitnesses++
			}
		}
		b.Instances = append(b.Instances, inst)
		if !families[tpl.ID] {
			families[tpl.ID] = true
			b.Templates = append(b.Templates, tpl)
		}
	}
	return b, nil
}

func extractWitnesses(p *Program, h Helper, nt NamedTrace) ([]Witness, error) {
	var out []Witness
	t := nt.Trace
	for start, entry := range t.Steps {
		if entry.PC != h.Entry {
			continue
		}
		n := len(entry.Stack)
		if n < 3 {
			return nil, fmt.Errorf("helper entry lacks three stack slots")
		}
		r := entry.Stack[n-3]
		continuation, ok := new(big.Int).SetString(r.Value, 0)
		if !ok || !continuation.IsUint64() {
			return nil, fmt.Errorf("invalid continuation")
		}
		// Stack slot provenance is retained even when the two operands have
		// identical concrete values, or DUP gives them the same value ID.
		operands := []int{entry.Stack[n-1].ID, entry.Stack[n-2].ID}
		w := Witness{Trace: nt.Name, EntryStep: start, Continuation: continuation.Uint64()}
		if start > 0 {
			loc := t.Steps[start-1].Source
			for _, site := range p.Sites {
				if site.Supported && contains(site.Location, loc) && sourceOperator(site.Operator) == h.Operation {
					w.Sites = append(w.Sites, site.ID)
				}
			}
		}
		anchor := -1
		for i := start; i < len(t.Steps); i++ {
			step := t.Steps[i]
			if step.Op == h.Operation && len(step.Inputs) == 2 && anchor < 0 {
				if t.Depends(step.Inputs[0], operands[0]) && t.Depends(step.Inputs[1], operands[1]) || t.Depends(step.Inputs[0], operands[1]) && t.Depends(step.Inputs[1], operands[0]) {
					anchor = i
					w.AnchorPC = step.PC
					w.OperandIDs = append([]int(nil), step.Inputs...)
					w.ResultID = step.Output
				}
			}
			if anchor >= 0 && step.Op == "JUMPI" && len(step.Inputs) == 2 && t.Depends(step.Inputs[1], w.ResultID) {
				w.GuardPCs = append(w.GuardPCs, step.PC)
			}
			if step.Op == "JUMP" && len(step.Stack) > 0 && step.Stack[len(step.Stack)-1].ID == r.ID {
				if i+1 >= len(t.Steps) || t.Steps[i+1].PC != w.Continuation {
					return nil, fmt.Errorf("return continuation mismatch")
				}
				exit := t.Steps[i+1].Stack
				if len(exit) != n-2 {
					return nil, fmt.Errorf("helper did not replace three inputs with one result")
				}
				for j := 0; j < n-3; j++ {
					if exit[j].ID != entry.Stack[j].ID {
						return nil, fmt.Errorf("helper changed caller stack prefix")
					}
				}
				if anchor < 0 || !t.Depends(exit[len(exit)-1].ID, w.ResultID) {
					return nil, fmt.Errorf("return not dependent on arithmetic anchor")
				}
				w.Outcome = "return"
				w.EndStep = i + 1
				out = append(out, w)
				break
			}
			if step.Op == "REVERT" {
				if t.Return != "4e487b710000000000000000000000000000000000000000000000000000000000000011" || t.ExecutionError == "" {
					return nil, fmt.Errorf("failure witness is not overflow panic")
				}
				// Optimizers may check SUB before computing the result. Keep
				// such evidence as a panic witness without inventing an anchor.
				w.Outcome = "panic_0x11"
				w.EndStep = i + 1
				out = append(out, w)
				break
			}
		}
	}
	return out, nil
}
func sourceOperator(op string) string {
	switch op {
	case "+":
		return "ADD"
	case "-":
		return "SUB"
	case "*":
		return "MUL"
	}
	return ""
}

func normalize(p *Program, h Helper, a *SymbolicAnalysis) (Template, map[string]uint64, error) {
	roles := map[uint64]string{}
	bindings := map[string]uint64{}
	role := func(pc uint64) string {
		if r, ok := roles[pc]; ok {
			return r
		}
		r := fmt.Sprintf("L%d", len(roles))
		roles[pc] = r
		bindings[r] = pc
		return r
	}
	role(h.Entry)
	for _, path := range a.Paths {
		for _, pc := range append(append([]uint64{}, path.PCs...), path.FailureProgram...) {
			role(pc)
		}
	}
	tpl := Template{Semantics: a.Semantics, Operation: h.Operation, InputStack: analysisInputs(a), OperandOrder: a.OperandOrder, OutputStack: []string{"result"}}
	if isLegacySemantics(a.Semantics) {
		tpl.FailureHandling = "execute_original_failure_path"
	}
	if a.IntegerType != "uint256" {
		tpl.IntegerType = a.IntegerType
	}
	for _, path := range a.Paths {
		normalized := PatternPath{Outcome: path.Outcome, PeakStack: path.PeakStack, Result: cleanExpr(path.Result), FailureData: path.FailureData, FailureMask: path.FailureMask, FailureMemory: cleanExpr(path.FailureMemory)}
		for _, condition := range path.Conditions {
			normalized.Conditions = append(normalized.Conditions, Condition{Expression: cleanExpr(condition.Expression), Nonzero: condition.Nonzero})
		}
		allPCs := append(append([]uint64{}, path.PCs...), path.FailureProgram...)
		for index, pc := range allPCs {
			ins := p.ByPC[pc]
			item := PatternInstruction{Role: role(pc), Op: ins.Op.String()}
			if ins.Op.IsPush() {
				item.Op = "PUSH"
				if data, ok := a.CodeDataProducers[pc]; ok {
					item.CodeData = data
				} else if target, ok := a.JumpProducers[pc]; ok {
					item.AddressRole = role(target)
				} else if a.DeadPushes[pc] {
					item.DeadConstant = true
				} else {
					item.Literal = constant(new(big.Int).SetBytes(ins.Data), nil).Value
				}
			}
			if index < len(path.PCs) {
				normalized.Instructions = append(normalized.Instructions, item)
			} else {
				normalized.FailureProgram = append(normalized.FailureProgram, item)
				normalized.FailureArgument = "errorMessage"
			}
		}
		tpl.Paths = append(tpl.Paths, normalized)
	}
	id, err := templateID(tpl)
	if err != nil {
		return tpl, nil, err
	}
	tpl.ID = id
	return tpl, bindings, nil
}
func templateID(tpl Template) (string, error) {
	tpl.ID = ""
	data, err := json.Marshal(tpl)
	if err != nil {
		return "", err
	}
	digest := sha256.Sum256(data)
	return "arith-" + hex.EncodeToString(digest[:12]), nil
}
func cleanExpr(e *Expr) *Expr {
	if e == nil {
		return nil
	}
	c := *e
	c.OriginPC = nil
	c.Args = nil
	for _, a := range e.Args {
		c.Args = append(c.Args, cleanExpr(a))
	}
	return &c
}
func validationPairs() [][2]*big.Int {
	values := []*big.Int{new(big.Int), big.NewInt(1), big.NewInt(2), big.NewInt(3), new(big.Int).Sub(new(big.Int).Lsh(big.NewInt(1), 128), big.NewInt(1)), new(big.Int).Lsh(big.NewInt(1), 128), new(big.Int).Sub(new(big.Int).Lsh(big.NewInt(1), 255), big.NewInt(1)), new(big.Int).Lsh(big.NewInt(1), 255), new(big.Int).Sub(new(big.Int).Set(mask256), big.NewInt(1)), new(big.Int).Set(mask256)}
	var pairs [][2]*big.Int
	for _, a := range values {
		for _, b := range values {
			pairs = append(pairs, [2]*big.Int{a, b})
		}
	}
	// Reproducible pseudo-random full-width cases independent of the bytecode.
	for i := 0; i < 128; i++ {
		left := sha256.Sum256([]byte(fmt.Sprintf("miner-left-%d", i)))
		right := sha256.Sum256([]byte(fmt.Sprintf("miner-right-%d", i)))
		pairs = append(pairs, [2]*big.Int{new(big.Int).SetBytes(left[:]), new(big.Int).SetBytes(right[:])})
	}
	return pairs
}
