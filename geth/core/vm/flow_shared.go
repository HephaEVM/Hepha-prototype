package vm

import (
	"fmt"
	"go-eth1.16.5-evm/regionexec"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

var flowInputRE = regexp.MustCompile(`^v([0-9]+)\s*:=\s*s.Stack\[n-([0-9]+)\]$`)
var flowHelperRE = regexp.MustCompile(`s.Helpers\[([0-9]+)\]`)
var flowGuardRE = regexp.MustCompile(`^if (!?)(v[0-9]+)\.IsZero\(\)\s*\{\s*return false\s*\}$`)

func flowStatements(src string) ([]string, error) {
	fset := token.NewFileSet()
	f, e := parser.ParseFile(fset, "", "package main\n"+src, 0)
	if e != nil {
		return nil, e
	}
	var out []string
	source := "package main\n" + src
	fn := f.Decls[0].(*ast.FuncDecl)
	for _, s := range fn.Body.List {
		// Preserve deterministic emitter text. Only the final accepted function
		// is formatted; raw candidates are parsed for structural statement bounds.
		body := source[fset.Position(s.Pos()).Offset:fset.Position(s.End()).Offset]
		out = append(out, strings.ReplaceAll(body, "\n\t", "\n"))
	}
	return out, nil
}

// Factor only structurally identical statements and complementary guards.
// The tree has at most four leaves. Failed admission never retries a path
// with a different condition and never publishes temporary memory writes.
func flowTree(paths [][]string) (string, error) {
	if len(paths) == 0 {
		return "return false\n", nil
	}
	var b strings.Builder
	for len(paths[0]) > 0 {
		first := paths[0][0]
		same := true
		for _, p := range paths[1:] {
			if len(p) == 0 || p[0] != first {
				same = false
				break
			}
		}
		if !same {
			break
		}
		b.WriteString(first)
		b.WriteByte('\n')
		for i := range paths {
			paths[i] = paths[i][1:]
		}
	}
	if len(paths[0]) == 0 {
		for _, p := range paths {
			if len(p) != 0 {
				return "", fmt.Errorf("unequal terminal")
			}
		}
		return b.String(), nil
	}
	m := flowGuardRE.FindStringSubmatch(paths[0][0])
	if m == nil {
		return "", fmt.Errorf("non-guard divergence: %s", paths[0][0])
	}
	var zero, nonzero [][]string
	for _, p := range paths {
		if len(p) == 0 {
			return "", fmt.Errorf("empty branch")
		}
		x := flowGuardRE.FindStringSubmatch(p[0])
		if x == nil || x[2] != m[2] {
			return "", fmt.Errorf("unshared definitions before guard")
		}
		if x[1] == "!" {
			zero = append(zero, p[1:])
		} else {
			nonzero = append(nonzero, p[1:])
		}
	}
	z, e := flowTree(zero)
	if e != nil {
		return "", e
	}
	nz, e := flowTree(nonzero)
	if e != nil {
		return "", e
	}
	fmt.Fprintf(&b, "if %s.IsZero(){\n%s}else{\n%s}\n", m[2], z, nz)
	return b.String(), nil
}
func emitFlowShared(rs []*hotRegion, proof bool) (string, []regionexec.Helper, error) {
	return emitPreparedFlowShared(prepareFlows(rs), proof)
}
func emitPreparedFlowShared(prepared []*preparedFlow, proof bool) (string, []regionexec.Helper, error) {
	src, helpers, err := emitPreparedFlowSharedRaw(prepared, proof)
	if err != nil {
		return "", nil, err
	}
	raw, err := format.Source([]byte("package main\n" + src))
	if err != nil {
		return "", nil, err
	}
	return strings.TrimPrefix(string(raw), "package main\n"), helpers, nil
}
func emitPreparedFlowSharedRaw(prepared []*preparedFlow, proof bool) (string, []regionexec.Helper, error) {
	var paths [][]string
	var independent []string
	var helpers []regionexec.Helper
	helperID := map[*arithmeticRegion]int{}
	inputs := map[int]int{}
	need, peak, writes := 0, 1024, 0
	gas := uint64(0)
	for path, pp := range prepared {
		r := pp.original
		need = max(need, r.Inputs)
		peak = min(peak, r.MaxStack)
		gas = max(gas, r.Gas)
		rr, _ := pp.optimized(proof)
		remap := map[int]int{}
		nh, nw := 0, 0
		for _, n := range rr.nodes {
			if n.op == MSTORE {
				nw++
			}
			if h := n.helper; h != nil {
				ix, ok := helperID[h]
				if !ok {
					ix = len(helpers)
					helperID[h] = ix
					helpers = append(helpers, regionexec.Helper(h.syncHot.native))
				}
				remap[nh] = ix
				nh++
			}
		}
		writes = max(writes, nw)
		raw, _, e := emitPreparedFlowRaw(pp, proof, true)
		if e != nil {
			return "", nil, e
		}
		raw = flowHelperRE.ReplaceAllStringFunc(raw, func(s string) string {
			x := flowHelperRE.FindStringSubmatch(s)
			i, _ := strconv.Atoi(x[1])
			return fmt.Sprintf("s.Helpers[%d]", remap[i])
		})
		raw = strings.ReplaceAll(raw, "s.PC=s.Exit", fmt.Sprintf("s.Path=%d\ns.PC=%d", path, r.Exit))
		independent = append(independent, strings.Replace(raw, "func Exec(", fmt.Sprintf("func Exec_%d(", path), 1))
		ss, e := flowStatements(raw)
		if e != nil {
			return "", nil, e
		}
		var body []string
		for i, s := range ss {
			if i < 2 || s == "var memEnd uint64" || strings.HasPrefix(s, "var writes [") || s == "nw:=0" || strings.HasPrefix(s, "skip:=[") {
				continue
			}
			if m := flowInputRE.FindStringSubmatch(s); m != nil {
				id, _ := strconv.Atoi(m[1])
				k, _ := strconv.Atoi(m[2])
				inputs[id] = k
				continue
			}
			body = append(body, s)
		}
		paths = append(paths, body)
	}
	body, e := flowTree(paths)
	if e != nil {
		return "", nil, e
	}
	// Keep the same recognized leaves and common admission as shared lowering.
	// CR-off restores independent candidate evaluation; every failed candidate
	// keeps its writes staged, and only the successful candidate commits.
	if paperWithoutReuse && len(prepared) > 1 {
		var s strings.Builder
		fmt.Fprintf(&s, "func Exec(s *regionexec.Frame)bool{\nn:=len(s.Stack)\nif n<%d||n>%d||s.Gas<%d{return false}\n", need, peak, gas)
		for i := range independent {
			fmt.Fprintf(&s, "if Exec_%d(s){return true}\n", i)
		}
		s.WriteString("return false\n}\n")
		s.WriteString(strings.Join(independent, "\n"))
		return s.String(), helpers, nil
	}
	var s strings.Builder
	fmt.Fprintf(&s, "func Exec(s *regionexec.Frame)bool{\nn:=len(s.Stack)\nif n<%d||n>%d||s.Gas<%d{return false}\nvar memEnd uint64\nvar writes [%d]regionexec.Write\nnw:=0\nskip:=[%d]bool{}\n", need, peak, gas, writes, writes)
	ids := make([]int, 0, len(inputs))
	for id := range inputs {
		ids = append(ids, id)
	}
	sort.Ints(ids)
	for _, id := range ids {
		fmt.Fprintf(&s, "v%d:=s.Stack[n-%d]\n_ = v%d\n", id, inputs[id], id)
	}
	s.WriteString(body)
	s.WriteString("}\n")
	return s.String(), helpers, nil
}
