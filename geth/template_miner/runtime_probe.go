package templateminer

// Read-only streaming diagnostics. This API validates candidate semantics once
// per code/entry, then selects the path for each actual invocation. A consumer
// must match its PCs and terminal stack before counting a completed region.
import (
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/holiman/uint256"
)

type RuntimeProbe struct {
	p          *Program
	compiler   string
	candidates map[uint64]*candidate
	guards     map[uint64]map[uint64]string
}
type InvocationPlan struct {
	Entry                                         uint64
	TemplateID, Semantics, Operation, IntegerType string
	PCs                                           []uint64
	Outcome                                       string
	Consumed                                      int
	Continuation                                  uint64
	Result                                        *big.Int
	Prefix                                        []uint256.Int
	Guards                                        map[uint64]string
	FailureData, FailureMask                      string
	path                                          *SymbolicPath
	argumentData                                  []byte
}

func NewRuntimeProbe(code []byte) (*RuntimeProbe, error) {
	ins, e := Disassemble(code)
	if e != nil {
		return nil, e
	}
	p := &Program{Code: code, Instructions: ins, ByPC: map[uint64]Instruction{}}
	for _, i := range ins {
		p.ByPC[i.PC] = i
	}
	return &RuntimeProbe{p: p, compiler: compilerFromRuntime(code), candidates: map[uint64]*candidate{}, guards: map[uint64]map[uint64]string{}}, nil
}
func (r *RuntimeProbe) Plan(entry uint64, stack []uint256.Int, memory []byte) (*InvocationPlan, error) {
	c := r.candidates[entry]
	if c == nil {
		// Discard business labels cheaply before invoking symbolic exploration.
		plausible := false
		for _, n := range []int{3, 4} {
			if len(stack) >= n && stack[len(stack)-n].IsUint64() && r.p.ByPC[stack[len(stack)-n].Uint64()].Op == vm.JUMPDEST {
				plausible = true
			}
		}
		if !plausible {
			return nil, nil
		}
		if strings.HasPrefix(r.compiler, "0.8.") || r.compiler == "" {
			c = discoverCandidate(r.p, entry)
		} else {
			c = discoverLegacyCandidate(r.p, entry)
		}
		if c.err != "" && r.compiler == "" {
			c = discoverLegacyCandidate(r.p, entry)
		}
		r.candidates[entry] = c
		if c.err == "" {
			r.guards[entry] = safetyGuards(c.analysis)
		}
	}
	if c.err != "" {
		return nil, nil
	}
	a := c.analysis
	layout := analysisInputs(a)
	n := len(stack)
	if n < len(layout) {
		return nil, fmt.Errorf("candidate entry lacks inputs")
	}
	inputs := map[string]*big.Int{}
	for i, name := range layout {
		inputs[name] = stack[n-len(layout)+i].ToBig()
	}
	ret := inputs["continuation"]
	if !ret.IsUint64() || r.p.ByPC[ret.Uint64()].Op != vm.JUMPDEST {
		return nil, fmt.Errorf("candidate continuation not JUMPDEST")
	}
	if e := CheckInputs(a, a.Operation, inputs[a.OperandOrder[0]], inputs[a.OperandOrder[1]]); e != nil {
		return nil, e
	}
	path, e := matchingPath(a, inputs)
	if e != nil {
		return nil, e
	}
	plan := &InvocationPlan{Entry: entry, TemplateID: c.template.ID, Semantics: a.Semantics, Operation: a.Operation, IntegerType: a.IntegerType, PCs: path.PCs, Outcome: path.Outcome, Consumed: len(layout), Continuation: ret.Uint64(), Prefix: append([]uint256.Int(nil), stack[:n-len(layout)]...), Guards: r.guards[entry], FailureData: path.FailureData, FailureMask: path.FailureMask, path: path}
	if path.Outcome == "return" {
		plan.Result, e = path.Result.Evaluate(inputs)
		return plan, e
	}
	if path.Outcome == "revert_argument" {
		vals := []*big.Int{}
		for _, v := range stack {
			vals = append(vals, v.ToBig())
		}
		run, e := executeLocal(r.p, entry, vals, memory, ret.Uint64())
		if e != nil {
			return nil, e
		}
		if run.Outcome != "revert" {
			return nil, fmt.Errorf("argument failure did not revert")
		}
		plan.PCs = run.PCs
		plan.Outcome = "revert"
		plan.argumentData = run.Data
	}
	return plan, nil
}
func (p *InvocationPlan) MatchReturn(pc uint64, stack []uint256.Int) bool {
	if p.Outcome != "return" || pc != p.Continuation || len(stack) != len(p.Prefix)+1 {
		return false
	}
	for i, v := range p.Prefix {
		if !v.Eq(&stack[i]) {
			return false
		}
	}
	return stack[len(stack)-1].ToBig().Cmp(p.Result) == 0
}
func (p *InvocationPlan) MatchFailure(data []byte, err error, reverted bool) bool {
	if p.Outcome == "return" {
		return false
	}
	if p.Outcome == "invalid" {
		return err != nil && strings.Contains(strings.ToLower(err.Error()), "invalid opcode") && len(data) == 0
	}
	if !reverted {
		return false
	}
	if p.argumentData != nil {
		return string(data) == string(p.argumentData)
	}
	if isLegacySemantics(p.Semantics) {
		return legacyFailureMatches(p.path, fmt.Sprintf("%x", data))
	}
	return panicOutcome(fmt.Sprintf("%x", data)) == p.Outcome
}
func (r *RuntimeProbe) Templates() []Template {
	out := []Template{}
	for _, c := range r.candidates {
		if c.err == "" {
			out = append(out, c.template)
		}
	}
	return out
}

// Count safety branches, not every branch inside a helper. Zero short-circuits
// are excluded: a safety arm must lead exclusively to failure. Legacy failure
// classes come from independent signed/unsigned boundary inputs, not strings.
func safetyGuards(a *SymbolicAnalysis) map[uint64]string {
	type arm struct {
		returns bool
		kinds   map[string]bool
		failure bool
	}
	arms := map[uint64]map[bool]*arm{}
	pathKinds := map[*SymbolicPath]map[string]bool{}
	typ, _ := parseIntegerType(a.IntegerType)
	for _, pair := range typedValidationPairs(typ) {
		inputs := map[string]*big.Int{a.OperandOrder[0]: pair[0], a.OperandOrder[1]: pair[1]}
		q, e := matchingPath(a, inputs)
		if e != nil || q.Outcome == "return" {
			continue
		}
		_, out, e := arithmeticSpec(a.Operation, typ, pair[0], pair[1])
		if e != nil {
			continue
		}
		if pathKinds[q] == nil {
			pathKinds[q] = map[string]bool{}
		}
		pathKinds[q][out] = true
	}
	for i := range a.Paths {
		q := &a.Paths[i]
		for _, c := range q.Conditions {
			if arms[c.GuardPC] == nil {
				arms[c.GuardPC] = map[bool]*arm{}
			}
			s := arms[c.GuardPC][c.Nonzero]
			if s == nil {
				s = &arm{kinds: map[string]bool{}}
				arms[c.GuardPC][c.Nonzero] = s
			}
			if q.Outcome == "return" {
				s.returns = true
			} else {
				s.failure = true
				for k := range pathKinds[q] {
					s.kinds[k] = true
				}
			}
		}
	}
	result := map[uint64]string{}
	for pc, sides := range arms {
		for _, s := range sides {
			if s.returns || !s.failure {
				continue
			}
			if s.kinds["panic_0x11"] {
				result[pc] = "overflow"
			} else if s.kinds["panic_0x12"] {
				result[pc] = "zero_divisor"
			}
		}
	}
	return result
}
