package templateminer

import (
	"fmt"
	"strings"
	"testing"
)

func TestSignedLegacyLibraryMatrix(t *testing.T) {
	for _, v := range []struct{ profile, compiler string }{{"2.5.1", "0517"}, {"3.4.2", "0612"}, {"3.4.2", "0706"}} {
		for _, opt := range []string{"unoptimized", "optimized"} {
			file := fmt.Sprintf("../template_miner/testdata/legacy-corpus/signed-oz-%s-solc-%s-%s.json", v.profile, v.compiler, opt)
			p, e := Load(file, "LegacyHarness.sol", "LegacyHarness")
			if e != nil {
				t.Fatal(e)
			}
			cache := map[uint64]*candidate{}
			for _, op := range []string{"add", "sub", "mul", "div"} {
				for _, pair := range actualStudyPairs(integerType{bits: 256, signed: true}) {
					input, e := EncodeCall(p, op+"_int256", []string{pair[0].String(), pair[1].String()})
					if e != nil {
						t.Fatal(e)
					}
					tr := Replay(p, input)
					rs, reject := mineTraceRegionsModel(p, tr, cache, true)
					if len(rs) != 1 {
						t.Fatalf("%s %s (%s,%s): %v", file, op, pair[0], pair[1], reject)
					}
					c := rs[0].c
					if c.analysis.IntegerType != "int256" || c.template.Semantics != legacySignedSemantics || rs[0].w.Outcome != legacyActualOutcome(tr) {
						t.Fatal("signed semantics lost")
					}
					if err := CheckInputs(c.analysis, c.analysis.Operation, pair[0], pair[1]); err != nil {
						t.Fatal(err)
					}
				}
			}
		}
	}
}

func TestLegacyErrorArgumentAcrossStrings(t *testing.T) {
	for _, opt := range []string{"unoptimized", "optimized"} {
		p, e := Load("../template_miner/testdata/legacy-corpus/argument-oz-2.5.1-solc-0517-"+opt+".json", "LegacyHarness.sol", "LegacyHarness")
		if e != nil {
			t.Fatal(e)
		}
		cache := map[uint64]*candidate{}
		ids := map[string]string{}
		for _, op := range []string{"sub", "div", "mod"} {
			for _, n := range []int{0, 1, 31, 32, 33, 64, 65, 257, 513} {
				for _, failure := range []bool{false, true} {
					args := []string{"7", "2", strings.Repeat("x", n)}
					if failure {
						if op == "sub" {
							args[0] = "1"
						} else {
							args[1] = "0"
						}
					}
					input, e := EncodeCall(p, op+"_message", args)
					if e != nil {
						t.Fatal(e)
					}
					tr := Replay(p, input)
					rs, reject := mineTraceRegionsModel(p, tr, cache, true)
					if len(rs) != 1 {
						t.Fatalf("%s %s len=%d failure=%v: %v", opt, op, n, failure, reject)
					}
					c := rs[0].c
					if len(c.template.InputStack) != 4 {
						t.Fatal("lost extra input")
					}
					if old := ids[op]; old != "" && old != c.template.ID {
						t.Fatal("template identity depends on runtime message")
					}
					ids[op] = c.template.ID
					if rs[0].w.Outcome != legacyActualOutcome(tr) {
						t.Fatal("lost failure")
					}
					if failure {
						f, end := frameFromTrace(tr)
						at := rs[0].w.EntryStep
						f.ops[at].Memory = ""
						if _, e := chainWitness(p, f, end, at, c); e == nil {
							t.Fatal("accepted missing failure memory evidence")
						}
					}
				}
			}
		}
	}
}
