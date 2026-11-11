package vm

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"math/rand"
	"os"
	"path/filepath"
	"reflect"
	"slices"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/holiman/uint256"
	miner "go-eth1.16.5-evm/template_miner"
)

func hashPlans(t testing.TB, code []byte) []*miner.HashPreparation {
	t.Helper()
	ins, err := miner.Disassemble(code)
	if err != nil {
		t.Fatal(err)
	}
	return miner.DiscoverHashPreparations(ins, nil)
}
func hashReference(t testing.TB, p *hashPreparationPlan, s *ScopeContext, e *EVM) {
	t.Helper()
	for _, addr := range p.region.PCs {
		pc := addr
		op := pragueInstructionSet[s.Contract.Code[pc]]
		if len(s.Stack.data) < op.minStack || len(s.Stack.data) > op.maxStack {
			t.Fatal("invalid reference stack")
		}
		cost := op.constantGas
		size := uint64(0)
		if op.memorySize != nil {
			var overflow bool
			size, overflow = op.memorySize(s.Stack)
			if overflow {
				t.Fatal("reference memory overflow")
			}
			size = toWordSize(size) * 32
		}
		if op.dynamicGas != nil {
			d, err := op.dynamicGas(e, s.Contract, s.Stack, s.Memory, size)
			if err != nil {
				t.Fatal(err)
			}
			cost += d
		}
		if cost > s.Contract.Gas {
			t.Fatal("reference OOG")
		}
		s.Contract.Gas -= cost
		s.Memory.Resize(size)
		if _, err := op.execute(&pc, e, s); err != nil {
			t.Fatal(err)
		}
	}
}
func hashScope(code []byte, words []uint256.Int, memory []byte) *ScopeContext {
	c := NewContract(common.HexToAddress("0x12345678"), common.HexToAddress("0x98765432"), new(uint256.Int), 1000000, nil)
	c.setResolvedCallCode(crypto.Keccak256Hash(code), code)
	m := NewMemory()
	m.Resize(uint64(len(memory)))
	copy(m.store, memory)
	w := uint64(len(memory)) / 32
	m.lastGasCost = w*3 + w*w/512
	return &ScopeContext{Contract: c, Stack: &Stack{data: append([]uint256.Int(nil), words...)}, Memory: m}
}
func hashDifferential(t testing.TB, code []byte, r *miner.HashPreparation, seed int64, samples int) {
	t.Helper()
	rng := rand.New(rand.NewSource(seed))
	for sample := 0; sample < samples; sample++ {
		words := make([]uint256.Int, r.Inputs+sample%4)
		for i := range words {
			words[i] = uint256.Int{rng.Uint64(), rng.Uint64(), rng.Uint64(), rng.Uint64()}
		}
		mem := make([]byte, 32*(sample%5)*40)
		rng.Read(mem)
		e := &EVM{hasher: crypto.NewKeccakState()}
		p := prepareHashPreparation(r, code, &pragueInstructionSet, HashPreparationDirect)
		if p == nil {
			t.Fatal("plan rejected")
		}
		want := hashScope(code, words, mem)
		hashReference(t, p, want, e)
		for _, variant := range []int{0, 1, 2} {
			mode := HashPreparationDirect
			if variant == 0 {
				mode = HashPreparationGeneric
			}
			p.mode = mode
			if variant == 2 {
				p.prepareInputs()
			}
			got := hashScope(code, words, mem)
			pc := r.Entry
			if !p.try(&pc, got, e, nil) || pc != r.Exit || got.Contract.Gas != want.Contract.Gas || got.Memory.lastGasCost != want.Memory.lastGasCost || !bytes.Equal(got.Memory.store, want.Memory.store) || !reflect.DeepEqual(got.Stack.data, want.Stack.data) {
				t.Fatalf("differential entry=%d mode=%d sample=%d", r.Entry, mode, sample)
			}
			got.Memory.Free()
			// Failed admission must be atomic even when the original would write
			// memory before eventually running out of gas.
			got = hashScope(code, words, mem)
			got.Contract.Gas = 1000000 - want.Contract.Gas - 1
			gas := got.Contract.Gas
			pc = r.Entry
			if p.try(&pc, got, e, nil) || pc != r.Entry || got.Contract.Gas != gas || !bytes.Equal(got.Memory.store, mem) || !slices.Equal(got.Stack.data, words) {
				t.Fatal("low gas mutated state")
			}
			got.Memory.Free()
		}
		want.Memory.Free()
	}
}

func TestHashPreparationDifferential(t *testing.T) {
	for _, raw := range []string{
		"6000526020600020",             // one word
		"6007602052600090815260409020", // preserve slot zero; mapping slot 7
		"6020525f5260405f20",           // consume two dynamic words
		"3360005260026020526040600020", // transaction-local caller
		"3060c052600260e052604060c020", // nonzero memory base
		"60ff166000526020600020",       // masked input
		"91906000526020600020",         // output permutation
	} {
		code, _ := hex.DecodeString(raw)
		rs := hashPlans(t, code)
		if len(rs) != 1 {
			t.Fatalf("no template %s", raw)
		}
		hashDifferential(t, code, rs[0], 711, 40)
	}
}
func TestHashPreparationRejects(t *testing.T) {
	for _, raw := range []string{
		"6000526040600020",             // second hash word uninitialized
		"6000526000526040600020",       // duplicate writes
		"6000526020600137206020600020", // calldata copy
		"6000525a506020600020",         // GAS cannot move across grouped charging
		"6000525b6020600020",           // interior entry
		"600052546020600020",           // state read
		"600052602035600020",           // dynamic size
	} {
		code, _ := hex.DecodeString(raw)
		if len(hashPlans(t, code)) != 0 {
			t.Fatalf("unsafe candidate %s", raw)
		}
	}
	code, _ := hex.DecodeString("6000526020600020")
	ins, _ := miner.Disassemble(code)
	if len(miner.DiscoverHashPreparations(ins, map[uint64]bool{2: true})) != 0 {
		t.Fatal("crossed existing entry")
	}
}
func TestHashPreparationFullEVM(t *testing.T) {
	// Return both the hash and the scratch memory so a correct hash alone is
	// insufficient. SSTORE and GAS execute after the replacement boundary.
	code, _ := hex.DecodeString("5f355f52600160205260405f20806002555a60805260605260a05ff3")
	for _, variant := range []int{1, 2, 3, 4} {
		mode := HashPreparationMode(variant)
		if variant == 4 {
			mode = HashPreparationDirect
		}
		stats := new(HashPreparationStats)
		cfg := &ArithmeticConfig{Engine: computeTestEngine(t), Mode: ArithmeticSync, HashPreparation: mode, HashPreparationStats: stats}
		for _, gas := range []uint64{0, 1, 10, 50, 70, 90, 100, 200, 1000, 22000, 50000, 100000} {
			for _, word := range []uint256.Int{{}, *uint256.NewInt(17), *new(uint256.Int).SetAllOne()} {
				input := computeInput(word)
				assertArithmeticSame(t, runArithmeticTest(t, code, input, gas, nil), runArithmeticTest(t, code, input, gas, cfg))
			}
		}
		if stats.Attempts == 0 || mode != HashPreparationObserve && stats.Skipped == 0 {
			t.Fatal("no execution")
		}
	}
}
func TestHashPreparationCorpus(t *testing.T) {
	out := os.Getenv("HASH_PREPARATION_CORPUS")
	if out == "" {
		t.Skip("optional real deployed bytecode corpus")
	}
	files, err := filepath.Glob("../../core/vm/testdata/regions/0x*.hex")
	if err != nil || len(files) == 0 {
		t.Fatal(err)
	}
	var rows []map[string]any
	families := map[string]int{}
	codes, cases, rejected := 0, 0, 0
	for _, file := range files {
		raw, err := os.ReadFile(file)
		if err != nil {
			t.Fatal(err)
		}
		code, err := hex.DecodeString(strings.TrimSpace(string(raw)))
		if err != nil {
			t.Fatal(err)
		}
		ins, err := miner.Disassemble(code)
		if err != nil {
			rejected++
			continue
		}
		rs := miner.DiscoverHashPreparations(ins, nil)
		if len(rs) > 0 {
			codes++
		}
		for i, r := range rs {
			families[r.StructuralID]++
			rows = append(rows, map[string]any{"code": filepath.Base(file), "region": r})
			if i < 3 {
				hashDifferential(t, code, r, int64(i+91), 4)
				cases += 4
			}
		}
	}
	b, err := json.MarshalIndent(map[string]any{"codeFiles": len(files), "codesWithPreparations": codes, "disassembleRejected": rejected, "regions": rows, "families": families, "differentialInputs": cases, "note": "Static opportunities before arithmetic/ABI ownership filtering; not dynamic coverage."}, "", "  ")
	if err != nil {
		t.Fatal(err)
	}
	if err = os.WriteFile(out, b, 0644); err != nil {
		t.Fatal(err)
	}
	t.Logf("codes=%d hitCodes=%d regions=%d families=%d differential=%d", len(files), codes, len(rows), len(families), cases)
}

func TestHashPreparationSafety(t *testing.T) {
	code, _ := hex.DecodeString("6000526020600020")
	r := hashPlans(t, code)[0]
	p := prepareHashPreparation(r, code, &pragueInstructionSet, HashPreparationDirect)
	e := &EVM{hasher: crypto.NewKeccakState()}
	for _, n := range []int{0, 1024} {
		s := hashScope(code, make([]uint256.Int, n), nil)
		pc := r.Entry
		gas := s.Contract.Gas
		if p.try(&pc, s, e, nil) || pc != r.Entry || len(s.Memory.store) != 0 || s.Contract.Gas != gas || len(s.Stack.data) != n {
			t.Fatal("stack rejection mutated state")
		}
		s.Memory.Free()
	}
	engine := computeTestEngine(t)
	seen := map[*boundArithmetic]bool{}
	for _, mode := range []HashPreparationMode{HashPreparationOff, HashPreparationObserve, HashPreparationGeneric, HashPreparationDirect} {
		cfg := &ArithmeticConfig{Mode: ArithmeticSync, HashPreparation: mode}
		b := engine.bind(code, arithmeticCacheKey{hash: crypto.Keccak256Hash(code), rules: &pragueInstructionSet, compute: cfg.computeHelperPolicy()})
		if seen[b] {
			t.Fatal("policy cache alias")
		}
		seen[b] = true
	}
	for _, mode := range []ArithmeticMode{ArithmeticObserve} {
		cfg := &ArithmeticConfig{Mode: mode, HashPreparation: HashPreparationDirect}
		if cfg.computeHelperPolicy() != 0 {
			t.Fatal("unsupported policy admitted")
		}
	}
	// Preimage recording remains available, and original GAS observations after
	// the region agree at every nearby gas boundary (including mid-region OOG).
	full, _ := hex.DecodeString("5f355f5260205f205a6020525f5260405ff3")
	var preimages map[common.Hash][]byte
	for _, mode := range []HashPreparationMode{HashPreparationOff, HashPreparationGeneric, HashPreparationDirect} {
		stats := new(HashPreparationStats)
		cfg := &ArithmeticConfig{Engine: engine, Mode: ArithmeticSync, HashPreparation: mode, HashPreparationStats: stats}
		evm, state, addr := testArithmeticEVM(t, full, cfg)
		evm.Config.EnablePreimageRecording = true
		_, _, err := evm.Call(common.Address{}, addr, computeInput(*uint256.NewInt(123)), 10000, new(uint256.Int))
		if err != nil {
			t.Fatal(err)
		}
		if mode == HashPreparationOff {
			preimages = state.Preimages()
		} else if !reflect.DeepEqual(preimages, state.Preimages()) || stats.Skipped == 0 {
			t.Fatal("preimage mismatch")
		}
		for gas := uint64(0); gas < 140; gas++ {
			assertArithmeticSame(t, runArithmeticTest(t, full, nil, gas, nil), runArithmeticTest(t, full, nil, gas, cfg))
		}
	}
}
