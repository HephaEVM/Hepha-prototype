package templateminer

// SafeMath is a behavior profile, not a table of opcode signatures. Candidate
// discovery and normalization are shared with compiler-generated checked math.
import (
	"fmt"
	"math/big"
	"strings"
)

const legacySemantics = "safemath_uint256"
const legacySignedSemantics = "safemath_int256"

func isLegacySemantics(s string) bool { return s == legacySemantics || s == legacySignedSemantics }
func analysisInputs(a *SymbolicAnalysis) []string {
	if len(a.InputStack) > 0 {
		return a.InputStack
	}
	return []string{"continuation", "b", "a"}
}
func inputOffset(a *SymbolicAnalysis, name string) int {
	v := analysisInputs(a)
	for i, s := range v {
		if s == name {
			return len(v) - 1 - i
		}
	}
	return -1
}

func legacyCompiler(version string) bool {
	return strings.HasPrefix(version, "0.4.") || strings.HasPrefix(version, "0.5.") || strings.HasPrefix(version, "0.6.") || strings.HasPrefix(version, "0.7.")
}

func expressionSize(e *Expr, limit int) int {
	if e == nil {
		return 0
	}
	n := 1
	for _, a := range e.Args {
		n += expressionSize(a, limit-n)
		if n > limit {
			return n
		}
	}
	return n
}

// A failure-path address is either constant or F+k, where F is the caller's
// free-memory pointer. F is not silently replaced with 0x80: real callers often
// allocated memory before invoking SafeMath. Only bounded offsets are accepted.
func memoryAffine(e *Expr) (int, *big.Int, bool) {
	if n, ok := constValue(e); ok {
		return 0, n, true
	}
	if e.Op == "input" && e.Input == "freeMemory" {
		return 1, new(big.Int), true
	}
	if len(e.Args) == 2 && (e.Op == "ADD" || e.Op == "SUB") {
		a, x, ok := memoryAffine(e.Args[0])
		b, y, ok2 := memoryAffine(e.Args[1])
		if !ok || !ok2 {
			return 0, nil, false
		}
		if e.Op == "ADD" {
			return a + b, new(big.Int).Add(x, y), true
		}
		return a - b, new(big.Int).Sub(x, y), true
	}
	return 0, nil, false
}

func memoryAddress(e *Expr) (bool, uint64, error) {
	k, n, ok := memoryAffine(e)
	if !ok || (k != 0 && k != 1) || !n.IsUint64() || n.Uint64() > 65536 {
		return false, 0, fmt.Errorf("unsupported symbolic failure-memory address")
	}
	return k == 1, n.Uint64(), nil
}

func legacyStore(s *symbolicState, offset, value *Expr) error {
	relative, at, err := memoryAddress(offset)
	if err != nil {
		return err
	}
	if !relative && at == 64 {
		s.freeMemory = value
		return nil
	}
	n, ok := constValue(value)
	if !ok {
		return fmt.Errorf("nonconstant failure payload; dynamic error strings unsupported")
	}
	mem := s.memory
	if relative {
		mem = s.relativeMemory
	}
	for i, b := range n.FillBytes(make([]byte, 32)) {
		mem[at+uint64(i)] = b
	}
	return nil
}

func legacyLoad(s *symbolicState, offset *Expr) (*Expr, error) {
	relative, at, err := memoryAddress(offset)
	if err != nil {
		return nil, err
	}
	if !relative && at == 64 {
		return s.freeMemory, nil
	}
	mem := s.memory
	if relative {
		mem = s.relativeMemory
	}
	b := make([]byte, 32)
	for i := range b {
		var ok bool
		b[i], ok = mem[at+uint64(i)]
		if !ok {
			return nil, fmt.Errorf("reads caller memory outside free-memory pointer")
		}
	}
	return constant(new(big.Int).SetBytes(b), nil), nil
}

func legacyReturnData(s *symbolicState, offset, size *Expr) ([]byte, []byte, error) {
	relative, at, err := memoryAddress(offset)
	if err != nil {
		return nil, nil, err
	}
	k, n, ok := memoryAffine(size)
	if !ok || k != 0 || !n.IsUint64() || n.Uint64() > 4096 {
		return nil, nil, fmt.Errorf("unsupported legacy failure data length")
	}
	mem := s.memory
	if relative {
		mem = s.relativeMemory
	}
	b := make([]byte, int(n.Uint64()))
	mask := make([]byte, len(b))
	unknown := false
	for i := range b {
		var ok bool
		b[i], ok = mem[at+uint64(i)]
		if ok {
			mask[i] = 255
		} else {
			unknown = true
		}
	}
	if !unknown {
		return b, nil, nil
	}
	// Solidity's older error encoder can leave only ABI string padding
	// untouched. Preserve these bytes as reads of the caller's entry memory;
	// never assume that memory beyond the free pointer is zero.
	if len(b) < 68 || string(b[:4]) != "\x08\xc3\x79\xa0" {
		return nil, nil, fmt.Errorf("unknown non-string failure payload")
	}
	for _, v := range mask[:68] {
		if v != 255 {
			return nil, nil, fmt.Errorf("unknown error-string header")
		}
	}
	start := new(big.Int).SetBytes(b[4:36])
	length := new(big.Int).SetBytes(b[36:68])
	if start.Cmp(big.NewInt(32)) != 0 || !length.IsUint64() || length.Uint64() > 4096 || 68+int((length.Uint64()+31)/32*32) != len(b) {
		return nil, nil, fmt.Errorf("unsupported error-string layout")
	}
	for _, v := range mask[68 : 68+int(length.Uint64())] {
		if v != 255 {
			return nil, nil, fmt.Errorf("dynamic error-string data unsupported")
		}
	}
	return b, mask, nil
}

func legacyFailureMatches(path *SymbolicPath, data string) bool {
	data = strings.TrimPrefix(data, "0x")
	if path.FailureMask == "" {
		return data == path.FailureData
	}
	if len(data) != len(path.FailureData) || len(path.FailureMask) != len(data) {
		return false
	}
	for i := 0; i < len(data); i += 2 {
		if path.FailureMask[i:i+2] == "ff" && data[i:i+2] != path.FailureData[i:i+2] {
			return false
		}
	}
	return true
}

func discoverLegacyCandidate(p *Program, entry uint64) *candidate {
	layouts := [][]string{nil, {"continuation", "b", "a", "errorMessage"}}
	reasons := []string{}
	for _, layout := range layouts {
		c := discoverLegacyLayout(p, entry, layout)
		if c.err == "" {
			return c
		}
		reasons = append(reasons, c.err)
	}
	return &candidate{err: strings.Join(reasons, "; ")}
}
func discoverLegacyLayout(p *Program, entry uint64, layout []string) *candidate {
	fail := func(err error) *candidate { return &candidate{err: "legacy: " + err.Error()} }
	h := Helper{Entry: entry}
	a, err := analyzeLayout(p, h, true, layout)
	if err != nil {
		return fail(err)
	}
	if !binaryOperation(a.Operation) {
		return fail(fmt.Errorf("not binary arithmetic"))
	}
	var types []integerType
	for _, t := range []integerType{{bits: 256}, {bits: 256, signed: true}} {
		valid := true
		for _, pair := range typedValidationPairs(t) {
			if checkTypedInputs(a, a.Operation, t, pair[0], pair[1]) != nil {
				valid = false
				break
			}
		}
		if valid {
			types = append(types, t)
		}
	}
	if len(types) != 1 {
		return fail(fmt.Errorf("expected unique uint256/int256 domain, got %d", len(types)))
	}
	a.IntegerType = types[0].name()
	if types[0].signed {
		a.Semantics = legacySignedSemantics
	}
	if len(layout) > 0 {
		if err := validateArgumentFailure(p, entry, a, types[0]); err != nil {
			return fail(err)
		}
	}
	h.Operation = a.Operation
	tpl, bindings, err := normalize(p, h, a)
	if err != nil {
		return fail(err)
	}
	return &candidate{helper: h, analysis: a, template: tpl, bindings: bindings}
}
