// Package templateminer discovers checked arithmetic helpers from compiler
// metadata and execution, independently of core/vm's hand-written templates.
package templateminer

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/core/vm"
)

type Artifact struct {
	SchemaVersion   int             `json:"schemaVersion"`
	CompilerVersion string          `json:"compilerVersion"`
	Settings        json.RawMessage `json:"settings"`
	InputSources    map[string]struct {
		Content string `json:"content"`
	} `json:"inputSources"`
	Output struct {
		Sources map[string]struct {
			ID  int            `json:"id"`
			AST map[string]any `json:"ast"`
		} `json:"sources"`
		Contracts map[string]map[string]Contract `json:"contracts"`
	} `json:"output"`
}

type Contract struct {
	ABI json.RawMessage `json:"abi"`
	EVM struct {
		DeployedBytecode Bytecode `json:"deployedBytecode"`
	} `json:"evm"`
}
type Bytecode struct {
	Object            string                   `json:"object"`
	SourceMap         string                   `json:"sourceMap"`
	FunctionDebugData map[string]FunctionDebug `json:"functionDebugData"`
	GeneratedSources  []struct {
		ID       int            `json:"id"`
		AST      map[string]any `json:"ast"`
		Contents string         `json:"contents"`
	} `json:"generatedSources"`
}
type FunctionDebug struct {
	EntryPoint     *uint64 `json:"entryPoint"`
	ParameterSlots int     `json:"parameterSlots"`
	ReturnSlots    int     `json:"returnSlots"`
}
type SourceLocation struct {
	Start         int    `json:"start"`
	Length        int    `json:"length"`
	File          int    `json:"file"`
	Jump          string `json:"jump,omitempty"`
	ModifierDepth int    `json:"modifierDepth,omitempty"`
}
type Site struct {
	ID        int            `json:"astId"`
	Source    string         `json:"source"`
	Operator  string         `json:"operator"`
	Type      string         `json:"type"`
	Location  SourceLocation `json:"location"`
	Checked   bool           `json:"checked"`
	Supported bool           `json:"supported"`
}
type Instruction struct {
	PC     uint64
	Op     vm.OpCode
	Data   []byte
	Source SourceLocation
}
type Helper struct {
	Name      string         `json:"name"`
	Operation string         `json:"operation"`
	Entry     uint64         `json:"entry"`
	Location  SourceLocation `json:"location"`
}
type Program struct {
	Artifact     *Artifact
	Source, Name string
	Contract     Contract
	Code         []byte
	Instructions []Instruction
	ByPC         map[uint64]Instruction
	Sites        []Site
	Helpers      []Helper
	Diagnostics  []string
}

func Load(path, source, name string) (*Program, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var a Artifact
	if err = json.Unmarshal(b, &a); err != nil {
		return nil, err
	}
	if a.SchemaVersion != 1 || (!strings.HasPrefix(a.CompilerVersion, "0.8.") && !legacyCompiler(a.CompilerVersion)) {
		return nil, fmt.Errorf("expected schemaVersion=1 and Solidity 0.4.x–0.8.x compiler metadata")
	}
	c, ok := a.Output.Contracts[source][name]
	if !ok {
		return nil, fmt.Errorf("contract %s:%s not found", source, name)
	}
	code, err := hex.DecodeString(strings.TrimPrefix(c.EVM.DeployedBytecode.Object, "0x"))
	if err != nil {
		return nil, fmt.Errorf("unlinked or invalid runtime bytecode: %w", err)
	}
	if len(code) == 0 {
		return nil, fmt.Errorf("empty runtime bytecode")
	}
	ins, err := Disassemble(code)
	if err != nil {
		return nil, err
	}
	locations, err := ParseSourceMap(c.EVM.DeployedBytecode.SourceMap)
	if err != nil {
		return nil, err
	}
	p := &Program{Artifact: &a, Source: source, Name: name, Contract: c, Code: code, Instructions: ins, ByPC: map[uint64]Instruction{}}
	for i := range p.Instructions {
		p.Instructions[i].Source = SourceLocation{File: -1}
		if i < len(locations) {
			p.Instructions[i].Source = locations[i]
		}
		p.ByPC[p.Instructions[i].PC] = p.Instructions[i]
	}
	if legacyCompiler(a.CompilerVersion) {
		p.Diagnostics = append(p.Diagnostics, "Pre-0.8 artifacts use execution-driven SafeMath discovery; source arithmetic nodes are not implicitly checked.")
		return p, nil
	}
	names := sortedKeys(a.Output.Sources)
	for _, filename := range names {
		walkSites(a.Output.Sources[filename].AST, filename, false, &p.Sites)
	}
	// Generated Yul provides semantic seeds, not opcode recognizers. The actual
	// operation must also occur in its AST, and the runtime region is checked
	// independently by value provenance and symbolic control-flow exploration.
	definitions := map[string]map[string]any{}
	for _, s := range c.EVM.DeployedBytecode.GeneratedSources {
		walk(s.AST, func(n map[string]any) {
			if str(n["nodeType"]) == "YulFunctionDefinition" {
				definitions[str(n["name"])] = n
			}
		})
	}
	for _, fname := range sortedKeys(c.EVM.DeployedBytecode.FunctionDebugData) {
		debug := c.EVM.DeployedBytecode.FunctionDebugData[fname]
		if !strings.HasPrefix(fname, "checked_") {
			continue
		}
		if !strings.HasSuffix(fname, "_t_uint256") || debug.ParameterSlots != 2 || debug.ReturnSlots != 1 {
			p.Diagnostics = append(p.Diagnostics, fname+": outside first-version uint256 binary-helper scope")
			continue
		}
		n, ok := definitions[fname]
		if !ok || debug.EntryPoint == nil {
			p.Diagnostics = append(p.Diagnostics, fname+": missing generated Yul AST or entry point")
			continue
		}
		ops := map[string]bool{}
		panicCall := false
		walk(n, func(x map[string]any) {
			if str(x["nodeType"]) != "YulFunctionCall" {
				return
			}
			f, _ := x["functionName"].(map[string]any)
			fn := str(f["name"])
			if fn == "add" || fn == "sub" || fn == "mul" {
				ops[fn] = true
			}
			if fn == "panic_error_0x11" {
				panicCall = true
			}
		})
		if len(ops) != 1 || !panicCall {
			p.Diagnostics = append(p.Diagnostics, fname+": ambiguous arithmetic or no overflow-panic semantic seed")
			continue
		}
		for op := range ops {
			p.Helpers = append(p.Helpers, Helper{fname, strings.ToUpper(op), *debug.EntryPoint, parseSrc(str(n["src"]))})
		}
	}
	if len(p.Helpers) == 0 {
		p.Diagnostics = append(p.Diagnostics, "no supported outlined checked helpers; inlined/viaIR regions are not inferred in this version")
	}
	return p, nil
}

func Disassemble(code []byte) ([]Instruction, error) {
	var out []Instruction
	for pc := 0; pc < len(code); {
		op := vm.OpCode(code[pc])
		width := 0
		if op >= vm.PUSH1 && op <= vm.PUSH32 {
			width = int(op-vm.PUSH1) + 1
		}
		// Solidity's metadata trailer is data and can contain truncated PUSH.
		// EVM semantics right-pad missing bytes; preserve this convention.
		data := make([]byte, width)
		copy(data, code[pc+1:min(pc+1+width, len(code))])
		out = append(out, Instruction{PC: uint64(pc), Op: op, Data: data})
		pc += 1 + width
	}
	return out, nil
}

// ParseSourceMap follows Solidity's instruction-indexed s:l:f:j:m encoding.
// Empty fields inherit the previous entry; offsets are UTF-8 byte offsets.
func ParseSourceMap(text string) ([]SourceLocation, error) {
	if text == "" {
		return nil, nil
	}
	previous := SourceLocation{File: -1, Jump: "-"}
	var out []SourceLocation
	for _, item := range strings.Split(text, ";") {
		fields := strings.Split(item, ":")
		if len(fields) > 5 {
			return nil, fmt.Errorf("invalid source-map entry %q", item)
		}
		for i, v := range fields {
			if v == "" {
				continue
			}
			if i == 3 {
				if v != "i" && v != "o" && v != "-" {
					return nil, fmt.Errorf("invalid source-map jump %q", v)
				}
				previous.Jump = v
				continue
			}
			n, err := strconv.Atoi(v)
			if err != nil {
				return nil, fmt.Errorf("invalid source-map integer %q", v)
			}
			switch i {
			case 0:
				previous.Start = n
			case 1:
				previous.Length = n
			case 2:
				previous.File = n
			case 4:
				previous.ModifierDepth = n
			}
		}
		out = append(out, previous)
	}
	return out, nil
}

func walkSites(v any, filename string, unchecked bool, sites *[]Site) {
	switch n := v.(type) {
	case map[string]any:
		if str(n["nodeType"]) == "UncheckedBlock" {
			unchecked = true
		}
		if str(n["nodeType"]) == "BinaryOperation" || str(n["nodeType"]) == "Assignment" {
			op := strings.TrimSuffix(str(n["operator"]), "=")
			if op == "+" || op == "-" || op == "*" {
				desc, _ := n["typeDescriptions"].(map[string]any)
				typ := str(desc["typeString"])
				*sites = append(*sites, Site{intNumber(n["id"]), filename, op, typ, parseSrc(str(n["src"])), !unchecked, !unchecked && typ == "uint256"})
			}
		}
		for _, key := range sortedKeys(n) {
			walkSites(n[key], filename, unchecked, sites)
		}
	case []any:
		for _, child := range n {
			walkSites(child, filename, unchecked, sites)
		}
	}
}
func walk(v any, visit func(map[string]any)) {
	switch n := v.(type) {
	case map[string]any:
		visit(n)
		for _, key := range sortedKeys(n) {
			walk(n[key], visit)
		}
	case []any:
		for _, c := range n {
			walk(c, visit)
		}
	}
}
func str(v any) string    { s, _ := v.(string); return s }
func intNumber(v any) int { n, _ := v.(float64); return int(n) }
func parseSrc(s string) SourceLocation {
	parts := strings.Split(s, ":")
	r := SourceLocation{File: -1}
	if len(parts) >= 3 {
		r.Start, _ = strconv.Atoi(parts[0])
		r.Length, _ = strconv.Atoi(parts[1])
		r.File, _ = strconv.Atoi(parts[2])
	}
	return r
}
func sortedKeys[V any](m map[string]V) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}
func contains(outer, inner SourceLocation) bool {
	return outer.File == inner.File && outer.Start <= inner.Start && inner.Start+inner.Length <= outer.Start+outer.Length
}
