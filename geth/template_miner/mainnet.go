package templateminer

// This path discovers regions from executed JUMPDESTs, without a compiler AST,
// a pre-trained template library, opcode-distance rules, or handwritten cleanup
// patterns. The bytecode analyzer must close BOTH normal and panic paths before
// a real execution can support a normalized template.
import (
	"bytes"
	"compress/gzip"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/params"
)

type ChainLabel struct {
	Address           string          `json:"address"`
	Name              string          `json:"name"`
	CompilerVersion   string          `json:"compilerVersion"`
	OptimizerEnabled  *bool           `json:"optimizerEnabled"`
	OptimizerMode     string          `json:"optimizerMode"`
	OptimizerDetails  json.RawMessage `json:"optimizerDetails,omitempty"`
	Runs              *int            `json:"runs,omitempty"`
	ViaIR             *bool           `json:"viaIR,omitempty"`
	EvidenceFile      string          `json:"evidenceFile,omitempty"`
	EvidenceSHA256    string          `json:"evidenceSHA256,omitempty"`
	SettingsSourceURL string          `json:"settingsSourceURL,omitempty"`
	SettingsStatus    string          `json:"settingsStatus"`
	Selection         string          `json:"selection,omitempty"`
	ObservedCalls     int             `json:"observedCalls,omitempty"`
	ActivityRank      int             `json:"activityRank,omitempty"`
}

// AuditRecord is the streaming format emitted by cmd/optimizer-audit. Frames
// are kept separate even when nested calls interleave their opcode records.
type AuditRecord struct {
	Memory           string   `json:"entryMemory,omitempty"`
	Kind             string   `json:"kind"`
	CallType         string   `json:"callType"`
	Block            uint64   `json:"block"`
	Tx               string   `json:"tx"`
	CallID           uint64   `json:"callId"`
	CodeAddress      string   `json:"codeAddress"`
	ContextAddress   string   `json:"contextAddress"`
	CodeHash         string   `json:"codeHash"`
	MetadataCompiler string   `json:"metadataCompiler"`
	Depth            int      `json:"depth"`
	PC               uint64   `json:"pc"`
	Step             int      `json:"step"`
	Op               string   `json:"op"`
	Stack            []string `json:"stackBeforeBottomToTop"`
	Error            string   `json:"error"`
	Output           string   `json:"output"`
	Reverted         bool     `json:"reverted"`
}

type JumpOrigin struct {
	JumpPC     uint64  `json:"jumpPc"`
	TargetPC   uint64  `json:"targetPc"`
	ValueID    int     `json:"valueId"`
	ProducerPC *uint64 `json:"producerPc,omitempty"`
	FromEntry  bool    `json:"fromEntry"`
	Taken      bool    `json:"taken"`
}

type ChainWitness struct {
	Block            uint64       `json:"block"`
	Tx               string       `json:"tx"`
	CallID           uint64       `json:"callId"`
	EntryStep        int          `json:"entryStep"`
	InstructionCount int          `json:"instructionCount"`
	Continuation     uint64       `json:"continuationPc"`
	Operands         []string     `json:"operands"`
	Outcome          string       `json:"outcome"`
	Result           string       `json:"result,omitempty"`
	FailureData      string       `json:"observedFailureData,omitempty"`
	Jumps            []JumpOrigin `json:"jumpOrigins"`
	ReturnedValueID  int          `json:"returnedValueId,omitempty"`
}

type ChainInstance struct {
	TemplateID      string            `json:"templateId"`
	Entry           uint64            `json:"entry"`
	Operation       string            `json:"operation"`
	IntegerType     string            `json:"integerType"`
	AddressBindings map[string]uint64 `json:"addressBindings"`
	Executions      int               `json:"observedExecutions"`
	Successes       int               `json:"successes"`
	Panics          int               `json:"panics"`
	LegacyFailures  map[string]int    `json:"legacyFailureOutcomes,omitempty"`
	ValidationPairs int               `json:"validationInputPairs"`
	Witnesses       []ChainWitness    `json:"sampleWitnesses"`
}

type ChainContract struct {
	ChainLabel
	CodeHash           string          `json:"codeHash"`
	CapturedFrames     int             `json:"capturedFrames"`
	Instances          []ChainInstance `json:"instances"`
	RejectedCandidates map[string]int  `json:"rejectedCandidates"`
}

type ChainStratum struct {
	CompilerVersion    string   `json:"compilerVersion"`
	Optimizer          string   `json:"optimizer"`
	ViaIR              string   `json:"viaIR"`
	Addresses          []string `json:"addresses"`
	CodeHashes         []string `json:"codeHashes"`
	TemplateIDs        []string `json:"templateIds"`
	ObservedExecutions int      `json:"observedExecutions"`
}

type FamilySupport struct {
	TemplateID         string   `json:"templateId"`
	Addresses          []string `json:"addresses"`
	CodeHashes         []string `json:"codeHashes"`
	CompilerVersions   []string `json:"compilerVersions"`
	OptimizerGroups    []string `json:"optimizerGroups"`
	ObservedExecutions int      `json:"observedExecutions"`
}

type ChainMiningReport struct {
	SchemaVersion   int              `json:"schemaVersion"`
	Mode            string           `json:"mode"`
	ArithmeticModel string           `json:"arithmeticModel,omitempty"`
	TraceSHA256     string           `json:"traceSHA256"`
	LabelsSHA256    string           `json:"labelsSHA256,omitempty"`
	Start           uint64           `json:"start"`
	End             uint64           `json:"end"`
	ValidatedBlocks int              `json:"validatedBlocks"`
	Transactions    int              `json:"transactions"`
	Frames          int              `json:"frames"`
	ExcludedFrames  map[string]int   `json:"excludedFrames"`
	Templates       []Template       `json:"templates"`
	Contracts       []*ChainContract `json:"contracts"`
	Families        []FamilySupport  `json:"families"`
	Strata          []ChainStratum   `json:"strata"`
	Limitations     []string         `json:"limitations"`
}

type chainFrame struct {
	header AuditRecord
	ops    []AuditRecord
}
type candidate struct {
	helper   Helper
	analysis *SymbolicAnalysis
	template Template
	bindings map[string]uint64
	err      string
}
type chainCode struct {
	program    *Program
	compiler   string
	candidates map[uint64]*candidate
}

// MineAudit consumes a completed replay, authenticates its bytecodes, discovers
// templates independently in every code hash, and aggregates cross-code support.
// Settings are annotations, never part of the template ID and never inferred
// from opcode layout. A labelled contract may legitimately produce no template.
func MineAudit(dir, labelPath string) (*ChainMiningReport, error) {
	return MineAuditWithModel(dir, labelPath, "auto")
}

func MineAuditWithModel(dir, labelPath, model string) (*ChainMiningReport, error) {
	if model != "auto" && model != "solidity08" && model != "safemath" {
		return nil, fmt.Errorf("arithmetic model must be auto, solidity08, or safemath")
	}
	r := &ChainMiningReport{SchemaVersion: 1, Mode: "mainnet_offline_candidates_not_enabled", Templates: []Template{}, Contracts: []*ChainContract{}, ExcludedFrames: map[string]int{}, Limitations: []string{
		"Solidity 0.8.x outlined binary uint/int helpers, plus pre-0.8 uint256/int256 SafeMath behavior and four-slot errorMessage argument layouts. Inlined, unary, tryMath and successful memory-writing wrappers remain unsupported.",
		"Legacy failure preserves INVALID versus REVERT, fixed revert bytes, and explicitly masked ABI padding read from caller memory; failure must execute its original path. Normal paths must preserve memory. Division guards alone cannot identify library authorship.",
		"Narrow and signed inputs must use canonical zero/sign extension. Inferred domains describe validated bytecode behavior, not guaranteed recovery of the source type.",
		"Normal paths and arithmetic predicates are symbolically analyzed; error-argument encoders are retained as original failure continuations and checked with varied strings. The 228 finite specification inputs and concrete witnesses are not a full-input equivalence proof.",
		"Execution counts describe captured calls only; first-N sampling is not an unbiased estimate of transaction-wide coverage.",
		"Optimizer and viaIR labels require a public verified runtime identical to the historical captured bytecode; unknown is separate from false.",
		"Multiple addresses sharing a code hash are reported separately from distinct bytecodes; distinct code hashes can still be clones.",
	}}
	r.ArithmeticModel = model
	if err := validateReplay(dir, r); err != nil {
		return nil, err
	}
	labels := map[string]ChainLabel{}
	if labelPath != "" {
		data, err := os.ReadFile(labelPath)
		if err != nil {
			return nil, err
		}
		digest := sha256.Sum256(data)
		r.LabelsSHA256 = hex.EncodeToString(digest[:])
		var rows []ChainLabel
		if err = json.Unmarshal(data, &rows); err != nil {
			return nil, err
		}
		for _, l := range rows {
			labels[strings.ToLower(l.Address)] = l
		}
	}
	codes := map[string]*chainCode{}
	contracts := map[string]*ChainContract{}
	families := map[string]Template{}
	process := func(f *chainFrame, end AuditRecord) error {
		r.Frames++
		if f.header.CallType == "CREATE" || f.header.CallType == "CREATE2" {
			r.ExcludedFrames["creation code, not deployed runtime"]++
			return nil
		}
		address := strings.ToLower(f.header.CodeAddress)
		hash := strings.ToLower(f.header.CodeHash)
		if len(hash) != 66 || !strings.HasPrefix(hash, "0x") {
			return fmt.Errorf("invalid code hash")
		}
		if _, err := hex.DecodeString(hash[2:]); err != nil {
			return err
		}
		cc := codes[hash]
		if cc == nil {
			data, err := os.ReadFile(filepath.Join(dir, "bytecode", hash+".hex"))
			if err != nil {
				return err
			}
			code, err := hex.DecodeString(strings.TrimPrefix(strings.TrimSpace(string(data)), "0x"))
			if err != nil {
				return err
			}
			if crypto.Keccak256Hash(code).Hex() != hash {
				return fmt.Errorf("captured bytecode hash mismatch %s", hash)
			}
			ins, err := Disassemble(code)
			if err != nil {
				return err
			}
			p := &Program{Code: code, Instructions: ins, ByPC: map[uint64]Instruction{}}
			for _, i := range ins {
				p.ByPC[i.PC] = i
			}
			cc = &chainCode{p, compilerFromRuntime(code), map[uint64]*candidate{}}
			codes[hash] = cc
		}
		key := address + ":" + hash
		contract := contracts[key]
		if contract == nil {
			label := labels[address]
			label.Address = address
			if label.Name == "" {
				label.Name = address
			}
			label = verifyChainLabel(label, filepath.Dir(labelPath), cc.program.Code, cc.compiler)
			contract = &ChainContract{ChainLabel: label, CodeHash: hash, Instances: []ChainInstance{}, RejectedCandidates: map[string]int{}}
			contracts[key] = contract
		}
		contract.CapturedFrames++
		legacy := legacyCompiler(contract.CompilerVersion) || (model == "safemath" && contract.CompilerVersion == "")
		if (model == "solidity08" && !strings.HasPrefix(contract.CompilerVersion, "0.8.")) || (model == "safemath" && !legacy) || (model == "auto" && !legacy && !strings.HasPrefix(contract.CompilerVersion, "0.8.")) {
			contract.RejectedCandidates["compiler outside requested arithmetic model (unknown requires explicit safemath mode)"]++
			return nil
		}
		type foundRegion struct {
			c *candidate
			w *ChainWitness
		}
		var found []foundRegion
		for i, s := range f.ops {
			ins, ok := cc.program.ByPC[s.PC]
			if !ok || ins.Op.String() != s.Op {
				return fmt.Errorf("trace opcode/bytecode mismatch %s PC %d", hash, s.PC)
			}
			if s.Step != i {
				return fmt.Errorf("non-contiguous frame %d step %d", f.header.CallID, s.Step)
			}
			if s.Op != "JUMPDEST" || len(s.Stack) < 3 || (i > 0 && f.ops[i-1].Op != "JUMP") {
				continue
			}
			if legacy {
				validReturn := false
				for _, arity := range []int{3, 4} {
					if len(s.Stack) >= arity {
						r, e := word(s.Stack[len(s.Stack)-arity])
						if e == nil && r.IsUint64() && cc.program.ByPC[r.Uint64()].Op == vm.JUMPDEST {
							validReturn = true
						}
					}
				}
				if !validReturn {
					continue
				}
			}
			c := cc.candidates[s.PC]
			if c == nil {
				if legacy {
					c = discoverLegacyCandidate(cc.program, s.PC)
				} else {
					c = discoverCandidate(cc.program, s.PC)
				}
				cc.candidates[s.PC] = c
			}
			if c.err != "" {
				contract.RejectedCandidates[c.err]++
				continue
			}
			w, err := chainWitness(cc.program, f, end, i, c)
			if err != nil {
				contract.RejectedCandidates["execution witness: "+err.Error()]++
				continue
			}
			found = append(found, foundRegion{c, w})
		}
		// A pure Solidity wrapper may also satisfy the input/output contract.
		// Keep the innermost complete checked region, never partial cleanup or
		// guard blocks, so one arithmetic execution is not counted repeatedly.
		for _, region := range found {
			c, w := region.c, region.w
			nested := false
			for _, other := range found {
				if other.c.helper.Operation == c.helper.Operation && other.w.EntryStep >= w.EntryStep && other.w.EntryStep+other.w.InstructionCount <= w.EntryStep+w.InstructionCount && other.w.InstructionCount < w.InstructionCount {
					nested = true
					break
				}
			}
			if nested {
				contract.RejectedCandidates["nested pure wrapper excluded"]++
				continue
			}
			var instance *ChainInstance
			for j := range contract.Instances {
				if contract.Instances[j].Entry == c.helper.Entry {
					instance = &contract.Instances[j]
					break
				}
			}
			if instance == nil {
				contract.Instances = append(contract.Instances, ChainInstance{TemplateID: c.template.ID, Entry: c.helper.Entry, Operation: c.helper.Operation, IntegerType: c.analysis.IntegerType, AddressBindings: c.bindings, ValidationPairs: len(validationPairs()), Witnesses: []ChainWitness{}})
				instance = &contract.Instances[len(contract.Instances)-1]
			}
			instance.Executions++
			if w.Outcome == "return" {
				instance.Successes++
			} else if isLegacySemantics(c.analysis.Semantics) {
				if instance.LegacyFailures == nil {
					instance.LegacyFailures = map[string]int{}
				}
				instance.LegacyFailures[w.Outcome]++
			} else {
				instance.Panics++
			}
			if len(instance.Witnesses) < 3 {
				instance.Witnesses = append(instance.Witnesses, *w)
			}
			families[c.template.ID] = c.template
		}
		return nil
	}
	file, err := os.Open(filepath.Join(dir, "traces.jsonl.gz"))
	if err != nil {
		return nil, err
	}
	defer file.Close()
	digest := sha256.New()
	z, err := gzip.NewReader(io.TeeReader(file, digest))
	if err != nil {
		return nil, err
	}
	defer z.Close()
	decoder := json.NewDecoder(z)
	frames := map[uint64]*chainFrame{}
	for {
		var rec AuditRecord
		err = decoder.Decode(&rec)
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("trace read: %w", err)
		}
		switch rec.Kind {
		case "frame":
			if _, ok := frames[rec.CallID]; ok {
				return nil, fmt.Errorf("duplicate frame %d", rec.CallID)
			}
			if rec.Block < r.Start || rec.Block > r.End {
				return nil, fmt.Errorf("frame outside validated block range")
			}
			frames[rec.CallID] = &chainFrame{header: rec}
		case "opcode":
			f := frames[rec.CallID]
			if f == nil {
				return nil, fmt.Errorf("opcode without frame")
			}
			f.ops = append(f.ops, rec)
		case "exit":
			f := frames[rec.CallID]
			if f == nil {
				return nil, fmt.Errorf("exit without frame")
			}
			if err = process(f, rec); err != nil {
				return nil, err
			}
			delete(frames, rec.CallID)
		}
	}
	if len(frames) > 0 {
		return nil, fmt.Errorf("incomplete trace: %d unfinished frames", len(frames))
	}
	r.TraceSHA256 = hex.EncodeToString(digest.Sum(nil))
	for _, k := range sortedKeys(contracts) {
		r.Contracts = append(r.Contracts, contracts[k])
		sort.Slice(contracts[k].Instances, func(i, j int) bool { return contracts[k].Instances[i].Entry < contracts[k].Instances[j].Entry })
	}
	for _, k := range sortedKeys(families) {
		r.Templates = append(r.Templates, families[k])
	}
	aggregateChainSupport(r)
	return r, nil
}

func discoverCandidate(p *Program, entry uint64) *candidate {
	c := &candidate{err: "not a supported closed integer binary helper"}
	h := Helper{Entry: entry}
	a, err := Analyze(p, h)
	if err != nil {
		return c
	}
	if err = validateIntegerDomain(a); err != nil {
		c.err = err.Error()
		return c
	}
	h.Operation = a.Operation
	t, bindings, err := normalize(p, h, a)
	if err != nil {
		c.err = err.Error()
		return c
	}
	return &candidate{helper: h, analysis: a, template: t, bindings: bindings}
}

func word(s string) (*big.Int, error) {
	n, ok := new(big.Int).SetString(strings.TrimPrefix(s, "0x"), 16)
	if !ok || n.Sign() < 0 || n.BitLen() > 256 {
		return nil, fmt.Errorf("invalid stack word")
	}
	return n, nil
}

// chainWitness selects a symbolic path using the recorded entry inputs and
// requires an exact instruction-path match. Dynamic IDs are then reconstructed
// within the region to corroborate return-address identity and caller-prefix
// preservation and expose PUSH->JUMP provenance in the output.
func chainWitness(p *Program, f *chainFrame, end AuditRecord, start int, c *candidate) (*ChainWitness, error) {
	if len(analysisInputs(c.analysis)) == 4 {
		return argumentWitness(p, f, end, start, c)
	}
	entry := f.ops[start].Stack
	n := len(entry)
	a, err := word(entry[n-1])
	if err != nil {
		return nil, err
	}
	b, err := word(entry[n-2])
	if err != nil {
		return nil, err
	}
	ret, err := word(entry[n-3])
	if err != nil || !ret.IsUint64() {
		return nil, fmt.Errorf("invalid return address")
	}
	inputs := map[string]*big.Int{"a": a, "b": b}
	if err = CheckInputs(c.analysis, c.helper.Operation, inputs[c.analysis.OperandOrder[0]], inputs[c.analysis.OperandOrder[1]]); err != nil {
		return nil, err
	}
	var path *SymbolicPath
	for i := range c.analysis.Paths {
		x := &c.analysis.Paths[i]
		holds := true
		for _, condition := range x.Conditions {
			v, e := condition.Expression.Evaluate(inputs)
			if e != nil {
				return nil, e
			}
			if (v.Sign() != 0) != condition.Nonzero {
				holds = false
				break
			}
		}
		if holds {
			if path != nil {
				return nil, fmt.Errorf("ambiguous path")
			}
			path = x
		}
	}
	if path == nil || start+len(path.PCs) > len(f.ops) {
		return nil, fmt.Errorf("missing complete path")
	}
	w := &ChainWitness{Block: f.header.Block, Tx: f.header.Tx, CallID: f.header.CallID, EntryStep: start, InstructionCount: len(path.PCs), Continuation: ret.Uint64(), Operands: []string{entry[n-1], entry[n-2]}, Outcome: path.Outcome, Jumps: []JumpOrigin{}}
	ids := make([]int, n)
	values := make([]*big.Int, n)
	for i := range ids {
		ids[i] = i + 1
		values[i], err = word(entry[i])
		if err != nil {
			return nil, err
		}
	}
	nextID := n + 1
	origins := map[int]uint64{}
	returnID := ids[n-3]
	memory := map[uint64]byte{}
	table, _ := vm.LookupInstructionSet(params.Rules{IsCancun: true})
	for k, pc := range path.PCs {
		s := f.ops[start+k]
		ins := p.ByPC[pc]
		if s.PC != pc || s.Op != ins.Op.String() || (s.Error != "" && !(ins.Op == vm.INVALID && path.Outcome == "invalid")) || len(s.Stack) != len(ids) {
			return nil, fmt.Errorf("path or stack mismatch")
		}
		for i, text := range s.Stack {
			v, e := word(text)
			if e != nil {
				return nil, e
			}
			if v.Cmp(values[i]) != 0 {
				return nil, fmt.Errorf("intermediate stack value mismatch")
			}
		}
		nn := len(ids)
		switch {
		case ins.Op >= vm.DUP1 && ins.Op <= vm.DUP16:
			d := int(ins.Op-vm.DUP1) + 1
			if nn < d {
				return nil, fmt.Errorf("DUP underflow")
			}
			ids = append(ids, ids[nn-d])
			values = append(values, values[nn-d])
		case ins.Op >= vm.SWAP1 && ins.Op <= vm.SWAP16:
			d := int(ins.Op-vm.SWAP1) + 1
			if nn <= d {
				return nil, fmt.Errorf("SWAP underflow")
			}
			ids[nn-1], ids[nn-1-d] = ids[nn-1-d], ids[nn-1]
			values[nn-1], values[nn-1-d] = values[nn-1-d], values[nn-1]
		default:
			if ins.Op == vm.INVALID && path.Outcome == "invalid" {
				continue
			}
			pops, limit := table[byte(ins.Op)].Stack()
			pushes := pops + int(params.StackLimit) - limit
			if pops > nn || pushes < 0 || pushes > 1 {
				return nil, fmt.Errorf("stack effect mismatch")
			}
			if ins.Op == vm.JUMP || ins.Op == vm.JUMPI {
				v, e := word(s.Stack[nn-1])
				if e != nil || !v.IsUint64() {
					return nil, fmt.Errorf("invalid jump")
				}
				id := ids[nn-1]
				j := JumpOrigin{JumpPC: pc, TargetPC: v.Uint64(), ValueID: id, FromEntry: id <= n, Taken: ins.Op == vm.JUMP || values[nn-2].Sign() != 0}
				if origin, ok := origins[id]; ok {
					j.ProducerPC = &origin
				}
				w.Jumps = append(w.Jumps, j)
				if k == len(path.PCs)-1 && path.Outcome == "return" && id != returnID {
					return nil, fmt.Errorf("return identity mismatch")
				}
			}
			var produced *big.Int
			if ins.Op == vm.CODECOPY {
				to, from, size := values[nn-1], values[nn-2], values[nn-3]
				if !to.IsUint64() || !from.IsUint64() || !size.IsUint64() || size.Uint64() > 4096 || from.Uint64() > uint64(len(p.Code)) || size.Uint64() > uint64(len(p.Code))-from.Uint64() {
					return nil, fmt.Errorf("CODECOPY witness outside model")
				}
				for j, b := range p.Code[from.Uint64() : from.Uint64()+size.Uint64()] {
					memory[to.Uint64()+uint64(j)] = b
				}
			}
			if ins.Op == vm.MSTORE {
				if !values[nn-1].IsUint64() {
					return nil, fmt.Errorf("memory offset outside witness model")
				}
				at := values[nn-1].Uint64()
				for j, b := range values[nn-2].FillBytes(make([]byte, 32)) {
					memory[at+uint64(j)] = b
				}
			}
			if pushes == 1 {
				if ins.Op.IsPush() {
					produced = new(big.Int).SetBytes(ins.Data)
				} else if ins.Op == vm.MLOAD && isLegacySemantics(c.analysis.Semantics) {
					if !values[nn-1].IsUint64() {
						return nil, fmt.Errorf("memory read outside witness model")
					}
					at := values[nn-1].Uint64()
					data := make([]byte, 32)
					initialized := true
					for j := range data {
						b, ok := memory[at+uint64(j)]
						data[j] = b
						initialized = initialized && ok
					}
					if initialized {
						produced = new(big.Int).SetBytes(data)
					} else if at == 64 && start+k+1 < len(f.ops) {
						following := f.ops[start+k+1].Stack
						if len(following) != nn {
							return nil, fmt.Errorf("missing free-memory pointer witness")
						}
						produced, err = word(following[len(following)-1])
						if err != nil || !produced.IsUint64() || produced.Uint64() < 128 || produced.Uint64()%32 != 0 {
							return nil, fmt.Errorf("noncanonical free-memory pointer")
						}
					} else {
						return nil, fmt.Errorf("uninitialized memory witness")
					}
				} else {
					args := make([]*Expr, pops)
					for j := range args {
						args[j] = constant(values[nn-1-j], nil)
					}
					produced, err = (&Expr{Op: ins.Op.String(), Args: args}).Evaluate(nil)
					if err != nil {
						return nil, err
					}
				}
			}
			ids = ids[:nn-pops]
			values = values[:nn-pops]
			if pushes == 1 {
				origins[nextID] = pc
				ids = append(ids, nextID)
				values = append(values, produced)
				nextID++
			}
		}
	}
	if path.Outcome == "return" {
		i := start + len(path.PCs)
		if i >= len(f.ops) {
			return nil, fmt.Errorf("missing continuation")
		}
		out := f.ops[i]
		if out.PC != ret.Uint64() || len(out.Stack) != n-2 || len(ids) != n-2 {
			return nil, fmt.Errorf("return layout mismatch")
		}
		for j := 0; j < n-3; j++ {
			x, e := word(out.Stack[j])
			if e != nil {
				return nil, e
			}
			y, e := word(entry[j])
			if e != nil {
				return nil, e
			}
			if ids[j] != j+1 || x.Cmp(y) != 0 {
				return nil, fmt.Errorf("caller prefix changed")
			}
		}
		got, e := word(out.Stack[len(out.Stack)-1])
		if e != nil {
			return nil, e
		}
		expected, e := path.Result.Evaluate(inputs)
		if e != nil {
			return nil, e
		}
		if got.Cmp(expected) != 0 {
			return nil, fmt.Errorf("concrete result mismatch")
		}
		w.Result = got.Text(16)
		w.ReturnedValueID = ids[len(ids)-1]
	} else {
		validFailure := end.Reverted && panicOutcome(end.Output) == path.Outcome
		if isLegacySemantics(c.analysis.Semantics) {
			validFailure = (path.Outcome == "revert" && end.Reverted && legacyFailureMatches(path, end.Output)) || (path.Outcome == "invalid" && strings.Contains(strings.ToLower(end.Error), "invalid opcode") && strings.TrimPrefix(end.Output, "0x") == "")
			w.FailureData = strings.TrimPrefix(end.Output, "0x")
		}
		if start+len(path.PCs) != len(f.ops) || !validFailure {
			return nil, fmt.Errorf("failure is not exact expected arithmetic panic")
		}
	}
	return w, nil
}

func compilerFromRuntime(code []byte) string {
	if len(code) < 2 {
		return ""
	}
	n := int(code[len(code)-2])<<8 | int(code[len(code)-1])
	if n+2 > len(code) {
		return ""
	}
	meta := code[len(code)-2-n : len(code)-2]
	needle := []byte{0x64, 's', 'o', 'l', 'c', 0x43}
	i := bytes.Index(meta, needle)
	if i < 0 || i+len(needle)+3 > len(meta) {
		return ""
	}
	v := meta[i+len(needle):]
	return fmt.Sprintf("%d.%d.%d", v[0], v[1], v[2])
}

func verifyChainLabel(l ChainLabel, base string, code []byte, compiler string) ChainLabel {
	l.CompilerVersion = compiler
	l.OptimizerEnabled = nil
	l.OptimizerMode = "unknown"
	l.OptimizerDetails = nil
	l.EvidenceSHA256 = ""
	l.Runs = nil
	l.ViaIR = nil
	l.SettingsStatus = "unknown: no verified historical runtime evidence"
	if l.EvidenceFile == "" {
		return l
	}
	var e struct {
		SourceURL string `json:"sourceURL"`
		Response  struct {
			Address      string `json:"address"`
			ChainID      string `json:"chainId"`
			RuntimeMatch string `json:"runtimeMatch"`
			Runtime      struct {
				Onchain string `json:"onchainBytecode"`
			} `json:"runtimeBytecode"`
			Compilation struct {
				Version  string `json:"compilerVersion"`
				Settings struct {
					Optimizer struct {
						Enabled *bool           `json:"enabled"`
						Runs    *int            `json:"runs"`
						Details json.RawMessage `json:"details"`
					} `json:"optimizer"`
					ViaIR *bool `json:"viaIR"`
				} `json:"compilerSettings"`
			} `json:"compilation"`
		} `json:"response"`
	}
	path := l.EvidenceFile
	if !filepath.IsAbs(path) {
		path = filepath.Join(base, path)
	}
	data, err := os.ReadFile(path)
	if err != nil || json.Unmarshal(data, &e) != nil {
		l.SettingsStatus = "unknown: unreadable evidence"
		return l
	}
	actual, err := hex.DecodeString(strings.TrimPrefix(e.Response.Runtime.Onchain, "0x"))
	if err != nil || !bytes.Equal(actual, code) || (e.Response.RuntimeMatch != "match" && e.Response.RuntimeMatch != "exact_match") || e.Response.ChainID != "1" || !strings.EqualFold(e.Response.Address, l.Address) {
		l.SettingsStatus = "unknown: historical runtime or address mismatch"
		return l
	}
	if compiler != "" && strings.Split(e.Response.Compilation.Version, "+")[0] != compiler {
		l.SettingsStatus = "unknown: compiler metadata conflict"
		return l
	}
	l.CompilerVersion = strings.Split(e.Response.Compilation.Version, "+")[0]
	l.OptimizerEnabled = e.Response.Compilation.Settings.Optimizer.Enabled
	l.OptimizerMode = flagName(l.OptimizerEnabled)
	l.OptimizerDetails = e.Response.Compilation.Settings.Optimizer.Details
	if l.OptimizerEnabled == nil && len(l.OptimizerDetails) > 0 && string(l.OptimizerDetails) != "null" && string(l.OptimizerDetails) != "{}" {
		l.OptimizerMode = "custom_details"
	}
	digest := sha256.Sum256(data)
	l.EvidenceSHA256 = hex.EncodeToString(digest[:])
	l.Runs = e.Response.Compilation.Settings.Optimizer.Runs
	l.ViaIR = e.Response.Compilation.Settings.ViaIR
	if l.ViaIR == nil {
		v := false
		l.ViaIR = &v
	}
	l.SettingsSourceURL = e.SourceURL
	l.SettingsStatus = "verified: public runtime equals historical captured runtime"
	return l
}

func validateReplay(dir string, r *ChainMiningReport) error {
	data, err := os.ReadFile(filepath.Join(dir, "block-validation.json"))
	if err != nil {
		return err
	}
	var rows []struct {
		Number          uint64 `json:"number"`
		Transactions    int    `json:"transactions"`
		Gas             uint64 `json:"gasUsed"`
		ExpectedGas     uint64 `json:"expectedGasUsed"`
		State           string `json:"stateRoot"`
		ExpectedState   string `json:"expectedStateRoot"`
		Receipt         string `json:"receiptRoot"`
		ExpectedReceipt string `json:"expectedReceiptRoot"`
	}
	if err = json.Unmarshal(data, &rows); err != nil {
		return err
	}
	if len(rows) == 0 {
		return fmt.Errorf("empty block validation")
	}
	for i, b := range rows {
		if b.State == "" || b.Receipt == "" || b.Gas != b.ExpectedGas || b.State != b.ExpectedState || b.Receipt != b.ExpectedReceipt || (i > 0 && b.Number != rows[i-1].Number+1) {
			return fmt.Errorf("invalid block validation at %d", b.Number)
		}
		r.Transactions += b.Transactions
	}
	r.Start = rows[0].Number
	r.End = rows[len(rows)-1].Number
	r.ValidatedBlocks = len(rows)
	return nil
}

func flagName(b *bool) string {
	if b == nil {
		return "unknown"
	}
	if *b {
		return "true"
	}
	return "false"
}
func uniqueStrings(a []string) []string {
	m := map[string]bool{}
	for _, s := range a {
		m[s] = true
	}
	return sortedKeys(m)
}
func aggregateChainSupport(r *ChainMiningReport) {
	families := map[string]*FamilySupport{}
	strata := map[string]*ChainStratum{}
	for _, c := range r.Contracts {
		key := c.CompilerVersion + "/optimizer=" + c.OptimizerMode + "/viaIR=" + flagName(c.ViaIR)
		s := strata[key]
		if s == nil {
			s = &ChainStratum{CompilerVersion: c.CompilerVersion, Optimizer: c.OptimizerMode, ViaIR: flagName(c.ViaIR), TemplateIDs: []string{}}
			strata[key] = s
		}
		s.Addresses = append(s.Addresses, c.Address)
		s.CodeHashes = append(s.CodeHashes, c.CodeHash)
		for _, inst := range c.Instances {
			f := families[inst.TemplateID]
			if f == nil {
				f = &FamilySupport{TemplateID: inst.TemplateID}
				families[inst.TemplateID] = f
			}
			f.Addresses = append(f.Addresses, c.Address)
			f.CodeHashes = append(f.CodeHashes, c.CodeHash)
			f.CompilerVersions = append(f.CompilerVersions, c.CompilerVersion)
			f.OptimizerGroups = append(f.OptimizerGroups, c.OptimizerMode)
			f.ObservedExecutions += inst.Executions
			s.TemplateIDs = append(s.TemplateIDs, inst.TemplateID)
			s.ObservedExecutions += inst.Executions
		}
	}
	for _, k := range sortedKeys(families) {
		f := families[k]
		f.Addresses = uniqueStrings(f.Addresses)
		f.CodeHashes = uniqueStrings(f.CodeHashes)
		f.CompilerVersions = uniqueStrings(f.CompilerVersions)
		f.OptimizerGroups = uniqueStrings(f.OptimizerGroups)
		r.Families = append(r.Families, *f)
	}
	for _, k := range sortedKeys(strata) {
		s := strata[k]
		s.Addresses = uniqueStrings(s.Addresses)
		s.CodeHashes = uniqueStrings(s.CodeHashes)
		s.TemplateIDs = uniqueStrings(s.TemplateIDs)
		r.Strata = append(r.Strata, *s)
	}
}
