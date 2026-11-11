package vm

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/holiman/uint256"
	"math/rand"
	"os"
	"path/filepath"
	"sort"
	"testing"
	"time"
)

type regionJoinScope struct {
	PC, Gas, Fee uint64
	Stack        []uint256.Int
	Memory       []byte
}
type regionJoinFixture struct {
	CodeHash, Code, Input, Caller string
	A, AE, B, BE                  uint64
	AK, BK                        string
	Count                         int
	Before, After                 regionJoinScope
}
type regionJoinCase struct {
	f           regionJoinFixture
	code, input []byte
	a, b, j     *demandPlan
}

var regionJoinCases []regionJoinCase

func regionJoinPrepare(t testing.TB) []regionJoinCase {
	requireFixtureEnv(t, "REGION_JOIN_FIXTURE", "DEMAND_DIR")
	if regionJoinCases != nil {
		return regionJoinCases
	}
	root, _ := filepath.Abs("../..")
	os.Setenv("DEMAND_ROOT", root)
	raw, err := os.ReadFile(os.Getenv("REGION_JOIN_FIXTURE"))
	if err != nil {
		t.Fatal(err)
	}
	var data map[string]regionJoinFixture
	if err = json.Unmarshal(raw, &data); err != nil {
		t.Fatal(err)
	}
	var fs []regionJoinFixture
	for _, f := range data {
		if f.AK == "demand" && f.BK == "demand" {
			fs = append(fs, f)
		}
	}
	sort.Slice(fs, func(i, j int) bool {
		if fs[i].Count != fs[j].Count {
			return fs[i].Count > fs[j].Count
		}
		return fs[i].CodeHash < fs[j].CodeHash
	})
	eng := regionArithmeticTestEngine(t)
	abi := chainABI(t)
	bindings := map[string]*boundArithmetic{}
	for _, f := range fs {
		code, _ := hex.DecodeString(f.Code)
		input, _ := hex.DecodeString(f.Input)
		b := bindings[f.CodeHash]
		if b == nil {
			e, _, _ := testArithmeticEVM(t, code, nil)
			e.Config = Config{}.WithPruning(PruningConfig{Engine: eng, ABI: abi})
			scope := hashScope(code, f.Before.Stack, f.Before.Memory)
			b, _ = e.abiExecutionBinding(scope.Contract)
			scope.Memory.Free()
			bindings[f.CodeHash] = b
		}
		get := func(pc uint64) *demandPlan {
			if b == nil || pc >= uint64(len(b.entryIndex)) {
				return nil
			}
			id := b.entryIndex[pc]
			if id == 0 || int(id) >= len(b.computeEntries) {
				return nil
			}
			c := b.computeEntries[id]
			if c == nil || c.extension == nil {
				return nil
			}
			return c.extension.demand
		}
		a, z := get(f.A), get(f.B)
		if a == nil || z == nil {
			continue
		}
		j := joinDemandRegions(a, map[uint64]*demandPlan{f.B: z}, code, &pragueInstructionSet)
		if j == nil {
			continue
		}
		a.slot.Ready()
		z.slot.Ready()
		j.slot.Ready()
		regionJoinCases = append(regionJoinCases, regionJoinCase{f, code, input, a, z, j})
		if len(regionJoinCases) == 12 {
			break
		}
	}
	events, err := WaitDemand()
	if err != nil {
		t.Fatal(err)
	}
	for _, v := range events {
		if v.Error != "" {
			t.Fatal(v.Error)
		}
	}
	if len(regionJoinCases) == 0 {
		t.Fatal("no candidates")
	}
	return regionJoinCases
}
func (c regionJoinCase) scope() *ScopeContext {
	s := hashScope(c.code, c.f.Before.Stack, c.f.Before.Memory)
	s.Contract.Input = c.input
	s.Contract.Gas = c.f.Before.Gas
	s.Memory.lastGasCost = c.f.Before.Fee
	return s
}
func TestRegionJoinDifferential(t *testing.T) {
	rng := rand.New(rand.NewSource(20261028))
	accepted, rejected := 0, 0
	for _, c := range regionJoinPrepare(t) {
		ca, cr := 0, 0
		for trial := 0; trial < 256; trial++ {
			s, w := c.scope(), c.scope()
			e := &EVM{table: &pragueInstructionSet, hasher: crypto.NewKeccakState()}
			switch trial % 8 {
			case 1:
				s.Contract.Gas = 0
			case 2:
				s.Stack.data = nil
			case 3:
				s.Contract.Input = append([]byte(nil), c.input...)
				rng.Read(s.Contract.Input)
			case 4:
				for i := range s.Stack.data {
					s.Stack.data[i].SetUint64(uint64(rng.Intn(128)))
				}
			case 5:
				s.Stack.data = makeWords(1024)
			case 6:
				s.Contract.Gas = c.j.regions[0].Gas - 1
			case 7:
				s.Contract.Input = nil
			}
			w.Stack.data = append([]uint256.Int(nil), s.Stack.data...)
			w.Contract.Gas = s.Contract.Gas
			w.Contract.Input = s.Contract.Input
			pc := c.f.A
			before := pipelineSnapshotOf(s, pc, nil, nil)
			if c.j.try(&pc, s, e) {
				r := c.j.regions[e.demandFrame.Path]
				wp := r.Entry
				if err := regionReference(&wp, r.Exit, w, e, r.Gas); err != nil {
					t.Fatal(c.f.CodeHash, trial, err)
				}
				if !samePipelineState(pipelineSnapshotOf(s, pc, nil, nil), pipelineSnapshotOf(w, wp, nil, nil)) {
					t.Fatal("mismatch", c.f.CodeHash, c.f.A, trial)
				}
				ca++
			} else {
				if !samePipelineState(before, pipelineSnapshotOf(s, pc, nil, nil)) {
					t.Fatal("mutated reject")
				}
				cr++
			}
			s.Memory.Free()
			w.Memory.Free()
		}
		accepted += ca
		rejected += cr
		t.Logf("%s/%d→%d count=%d paths=%d accepted=%d rejected=%d", c.f.CodeHash[:12], c.f.A, c.f.B, c.f.Count, len(c.j.regions), ca, cr)
	}
	if accepted == 0 || rejected == 0 {
		t.Fatal(accepted, rejected)
	}
	t.Log("accepted", accepted, "rejected", rejected)
}
func TestRegionJoinPaired(t *testing.T) {
	type row struct {
		Case              string
		Count, Round, Arm int
		NS                float64
	}
	var rows []row
	for _, c := range regionJoinPrepare(t) {
		s := c.scope()
		e := &EVM{table: &pragueInstructionSet, hasher: crypto.NewKeccakState()}
		pc := c.f.A
		if !c.j.try(&pc, s, e) || pc != c.f.BE {
			t.Log("sample not joined", c.f.CodeHash, c.f.A, pc)
			s.Memory.Free()
			continue
		}
		s.Memory.Free()
		for round := 0; round < 12; round++ {
			for order := 0; order < 3; order++ {
				arm := (round + order) % 3
				s := c.scope()
				s.Stack.data = append(make([]uint256.Int, 0, 1024), s.Stack.data...)
				e := &EVM{table: &pragueInstructionSet, hasher: crypto.NewKeccakState()}
				const reps = 10000
				start := time.Now()
				for i := 0; i < reps; i++ {
					s.Stack.data = s.Stack.data[:len(c.f.Before.Stack)]
					copy(s.Stack.data, c.f.Before.Stack)
					copy(s.Memory.store, c.f.Before.Memory)
					s.Contract.Gas = c.f.Before.Gas
					s.Memory.lastGasCost = c.f.Before.Fee
					pc := c.f.A
					switch arm {
					case 0:
						if err := regionReference(&pc, c.f.BE, s, e); err != nil {
							t.Fatal(err)
						}
					case 1:
						if !c.a.try(&pc, s, e) {
							t.Fatal("A reject")
						}
						if err := regionReference(&pc, c.f.B, s, e); err != nil {
							t.Fatal(err)
						}
						if !c.b.try(&pc, s, e) {
							t.Fatal("B reject")
						}
					case 2:
						if !c.j.try(&pc, s, e) {
							t.Fatal("join reject")
						}
					}
					if pc != c.f.BE {
						t.Fatal("exit", pc, c.f.BE)
					}
				}
				rows = append(rows, row{fmt.Sprintf("%s/%d/%d", c.f.CodeHash[:12], c.f.A, c.f.B), c.f.Count, round, arm, float64(time.Since(start).Nanoseconds()) / reps})
				s.Memory.Free()
			}
		}
	}
	out := os.Getenv("REGION_JOIN_REPORT")
	if out == "" {
		t.Fatal("REGION_JOIN_REPORT required")
	}
	raw, _ := json.MarshalIndent(rows, "", "  ")
	if err := os.WriteFile(out, raw, 0644); err != nil {
		t.Fatal(err)
	}
}

// This small fixture checks composition itself without compiled plugins or
// external transaction data. Numerical execution is covered by the replay test.
func TestRegionJoinStructure(t *testing.T) {
	a := &hotRegion{Entry: 0, Exit: 3, Gas: 6, Inputs: 1, Drop: 1, MaxStack: 1023, Instructions: 2, nodes: []hotRegionNode{{op: PUSH32, value: uint256.Int{1}}, {op: ADD, args: []int{32, 31}}}, outputs: []int{33}}
	b := &hotRegion{Entry: 3, Exit: 6, Gas: 8, Inputs: 1, Drop: 1, MaxStack: 1023, Instructions: 2, nodes: []hotRegionNode{{op: PUSH32, value: uint256.Int{2}}, {op: MUL, args: []int{32, 31}}}, outputs: []int{33}}
	code := []byte{byte(PUSH1), 1, byte(ADD), byte(PUSH1), 2, byte(MUL)}
	p, st, next := regionBridge(a, map[uint64]*demandPlan{3: {regions: []*hotRegion{b}}}, code, &pragueInstructionSet, codeBitmap(code))
	if next == nil {
		t.Fatal("missing successor")
	}
	q := joinRegionTail(p, st, b)
	if q == nil || q.Gas != 14 || q.Drop != 1 || q.Inputs != 1 || q.Instructions != 4 || q.outputs[0] != 35 || q.nodes[3].args[1] != 33 {
		t.Fatal("invalid composition", q)
	}
	// A→SSTORE→B and unresolved dynamic JUMP are not preparation bridges.
	for _, op := range []OpCode{SSTORE, SLOAD, CALL, LOG0, JUMPI, JUMP} {
		bad := append([]byte(nil), code...)
		bad[3] = byte(op)
		_, _, target := regionBridge(a, map[uint64]*demandPlan{5: {regions: []*hotRegion{b}}}, bad, &pragueInstructionSet, codeBitmap(bad))
		if target != nil {
			t.Fatal("crossed boundary", op)
		}
	}
	if a.nodes[1].args[0] != 32 || len(a.nodes) != 2 || b.nodes[1].args[1] != 31 {
		t.Fatal("mutated input summary")
	}
}

func TestRegionJoinOriginalFallback(t *testing.T) {
	c := regionJoinPrepare(t)[0]
	q := *c.j
	q.fallback = c.a
	gas := uint64(0)
	for _, r := range c.a.regions {
		gas = max(gas, r.Gas)
	}
	s, w := c.scope(), c.scope()
	defer s.Memory.Free()
	defer w.Memory.Free()
	s.Contract.Gas = gas
	w.Contract.Gas = gas
	e := &EVM{table: &pragueInstructionSet, hasher: crypto.NewKeccakState()}
	p, wp := c.f.A, c.f.A
	if !c.a.try(&wp, w, e) {
		t.Fatal("original region refused fixture")
	}
	if !q.try(&p, s, e) || p != wp || p == c.f.BE || !samePipelineState(pipelineSnapshotOf(s, p, nil, nil), pipelineSnapshotOf(w, wp, nil, nil)) {
		t.Fatal("failed joined admission did not retain original region", p, wp)
	}
}
