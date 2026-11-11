package vm

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"
	"strings"
	"sync"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/params"
	"github.com/holiman/uint256"
	"go-eth1.16.5-evm/core/state"
	"go-eth1.16.5-evm/core/tracing"
	"go-eth1.16.5-evm/core/types"
	miner "go-eth1.16.5-evm/template_miner"
)

func testMinedEngine(t testing.TB, code []byte) *ArithmeticEngine {
	t.Helper()
	report, err := miner.MatchBytecode(code, nil)
	if err != nil {
		t.Fatal(err)
	}
	feedback := miner.FeedbackBundle{SchemaVersion: 1, Codes: []miner.FeedbackCode{{CodeHash: report.RuntimeCodeHash, Bytecode: hex.EncodeToString(code), Candidates: report.Candidates}}}
	bundle, err := miner.MineFeedback(feedback)
	if err != nil {
		t.Fatal(err)
	}
	if len(bundle.Templates) == 0 {
		t.Fatal("no mined templates")
	}
	e, err := NewArithmeticEngine(bundle.Templates, ArithmeticEngineOptions{})
	if err != nil {
		t.Fatal(err)
	}
	return e
}

type arithmeticTestRun struct {
	data   []byte
	gas    uint64
	err    string
	root   common.Hash
	refund uint64
	logs   int
}

func testArithmeticEVM(t testing.TB, code []byte, cfg *ArithmeticConfig) (*EVM, *state.StateDB, common.Address) {
	t.Helper()
	s, err := state.New(types.EmptyRootHash, state.NewDatabaseForTesting())
	if err != nil {
		t.Fatal(err)
	}
	addr := common.HexToAddress("0x1000")
	s.CreateAccount(addr)
	s.SetCode(addr, code, tracing.CodeChangeUnspecified)
	random := common.Hash{}
	evm := NewEVM(BlockContext{BlockNumber: big.NewInt(1), Time: 1, Random: &random, GasLimit: 30_000_000, BaseFee: big.NewInt(0), BlobBaseFee: big.NewInt(0), Difficulty: big.NewInt(0), CanTransfer: func(s StateDB, a common.Address, v *uint256.Int) bool { return s.GetBalance(a).Cmp(v) >= 0 }, Transfer: func(s StateDB, a, b common.Address, v *uint256.Int) {
		s.SubBalance(a, v, tracing.BalanceChangeTransfer)
		s.AddBalance(b, v, tracing.BalanceChangeTransfer)
	}}, s, params.AllDevChainProtocolChanges, Config{Arithmetic: cfg})
	evm.SetTxContext(TxContext{GasPrice: big.NewInt(0)})
	s.Prepare(evm.chainRules, common.Address{}, common.Address{}, &addr, ActivePrecompiles(evm.chainRules), nil)
	return evm, s, addr
}
func runArithmeticTest(t *testing.T, code, input []byte, gas uint64, cfg *ArithmeticConfig) arithmeticTestRun {
	evm, s, addr := testArithmeticEVM(t, code, cfg)
	ret, left, err := evm.Call(common.Address{}, addr, input, gas, new(uint256.Int))
	errText := ""
	if err != nil {
		errText = err.Error()
	}
	return arithmeticTestRun{append([]byte(nil), ret...), left, errText, s.IntermediateRoot(true), s.GetRefund(), len(s.Logs())}
}
func assertArithmeticSame(t *testing.T, want, got arithmeticTestRun) {
	t.Helper()
	if !bytes.Equal(want.data, got.data) || want.gas != got.gas || want.err != got.err || want.root != got.root || want.refund != got.refund || want.logs != got.logs {
		t.Fatalf("differential mismatch\noriginal=%+v\noptimized=%+v", want, got)
	}
}

func TestArithmeticPipelineCompilerFixtures(t *testing.T) {
	for _, fixture := range []string{"arithmetic-0819.json", "arithmetic-0824.json", "arithmetic-0824-optimized.json"} {
		t.Run(fixture, func(t *testing.T) {
			p, err := miner.Load("../../template_miner/testdata/"+fixture, "Arithmetic.sol", "Arithmetic")
			if err != nil {
				t.Fatal(err)
			}
			e := testMinedEngine(t, p.Code)
			max := new(big.Int).Sub(new(big.Int).Lsh(big.NewInt(1), 256), big.NewInt(1)).String()
			for _, method := range []string{"add", "sub", "mul"} {
				for _, args := range [][]string{{"9", "7"}, {"0", "1"}, {max, "2"}, {max, "0"}} {
					input, err := miner.EncodeCall(p, method, args)
					if err != nil {
						t.Fatal(err)
					}
					for _, gas := range []uint64{100000, 0, 50, 200, 400, 600} {
						want := runArithmeticTest(t, p.Code, input, gas, nil)
						for _, mode := range []ArithmeticMode{ArithmeticObserve, ArithmeticSync} {
							got := runArithmeticTest(t, p.Code, input, gas, &ArithmeticConfig{Engine: e, Mode: mode})
							assertArithmeticSame(t, want, got)
						}
						assertArithmeticSame(t, want, runArithmeticTest(t, p.Code, input, gas, &ArithmeticConfig{Engine: e, Mode: ArithmeticSync}))
						assertArithmeticSame(t, want, runArithmeticTest(t, p.Code, input, gas, &ArithmeticConfig{Engine: e, Mode: ArithmeticSync}))
						for range []bool{false} {
							assertArithmeticSame(t, want, runArithmeticTest(t, p.Code, input, gas, &ArithmeticConfig{Engine: e, Mode: ArithmeticSync}))
						}
					}
				}
			}
			if e.Stats().Pruned == 0 || e.Stats().Checked == 0 {
				t.Fatalf("pipeline not exercised: %+v", e.Stats())
			}
			t.Logf("%s: %+v", fixture, e.Stats())
		})
	}
}

func TestArithmeticPipelineIntegerMatrix(t *testing.T) {
	for _, contract := range []string{"UnsignedMatrix", "SignedMatrix"} {
		p, err := miner.Load("../../template_miner/testdata/integer-0824.json", "IntegerMatrix.sol", contract)
		if err != nil {
			t.Fatal(err)
		}
		e := testMinedEngine(t, p.Code)
		for _, width := range []uint{8, 64, 128, 256} {
			signed := contract == "SignedMatrix"
			prefix := "uint"
			bits := width
			if signed {
				prefix = "int"
				bits--
			}
			hi := new(big.Int).Sub(new(big.Int).Lsh(big.NewInt(1), bits), big.NewInt(1)).String()
			pairs := [][]string{{"7", "2"}, {"0", "0"}, {hi, "1"}, {hi, "2"}, {"0", "1"}}
			if signed {
				lo := new(big.Int).Neg(new(big.Int).Lsh(big.NewInt(1), bits)).String()
				pairs = append(pairs, []string{lo, "-1"}, []string{"-7", "2"})
			}
			for _, op := range []string{"add", "sub", "mul", "div", "mod"} {
				for _, pair := range pairs {
					method := fmt.Sprintf("%s_%s%d", op, prefix, width)
					input, err := miner.EncodeCall(p, method, pair)
					if err != nil {
						t.Fatal(err)
					}
					want := runArithmeticTest(t, p.Code, input, 100000, nil)
					assertArithmeticSame(t, want, runArithmeticTest(t, p.Code, input, 100000, &ArithmeticConfig{Engine: e, Mode: ArithmeticSync}))
					assertArithmeticSame(t, want, runArithmeticTest(t, p.Code, input, 100000, &ArithmeticConfig{Engine: e, Mode: ArithmeticSync}))
					assertArithmeticSame(t, want, runArithmeticTest(t, p.Code, input, 100000, &ArithmeticConfig{Engine: e, Mode: ArithmeticSync}))
					for _, mode := range []ArithmeticMode{ArithmeticSync} {
						t.Run(method+strings.Join(pair, ",")+fmt.Sprint(mode), func(t *testing.T) {
							assertArithmeticSame(t, want, runArithmeticTest(t, p.Code, input, 100000, &ArithmeticConfig{Engine: e, Mode: mode}))
						})
					}
					for _, mode := range []ArithmeticMode{ArithmeticSync} {
						assertArithmeticSame(t, want, runArithmeticTest(t, p.Code, input, 100000, &ArithmeticConfig{Engine: e, Mode: mode}))
					}
				}
			}
		}
		if e.Stats().Pruned == 0 || e.Stats().ModelMismatches != 0 {
			t.Fatalf("unexpected stats %+v", e.Stats())
		}
		t.Logf("%s: %+v", contract, e.Stats())
	}
}

func TestArithmeticRangeExactBeforeWrapping(t *testing.T) {
	max := new(uint256.Int).SetAllOne()
	zero := uint256.Int{}
	one := *uint256.NewInt(1)
	for _, c := range []struct {
		op   string
		a, b uint256.Int
		code byte
	}{{"ADD", *max, one, 0x11}, {"SUB", zero, one, 0x11}, {"MUL", *max, *uint256.NewInt(2), 0x11}, {"DIV", one, zero, 0x12}, {"MOD", one, zero, 0x12}} {
		_, code, err := CheckArithmeticRange(c.op, IntegerDomain{256, false}, c.a, c.b)
		if err != nil || code != c.code {
			t.Fatalf("%s: %x %v", c.op, code, err)
		}
	}
	min := new(uint256.Int).Lsh(uint256.NewInt(1), 255)
	_, code, err := CheckArithmeticRange("SDIV", IntegerDomain{256, true}, *min, *max)
	if err != nil || code != 0x11 {
		t.Fatal(code, err)
	}
	_, code, err = CheckArithmeticRange("SMOD", IntegerDomain{256, true}, *min, *max)
	if err != nil || code != 0 {
		t.Fatal(code, err)
	}
}

func TestArithmeticFeedbackRoundTripAndCacheIsolation(t *testing.T) {
	p, err := miner.Load("../../template_miner/testdata/arithmetic-0824.json", "Arithmetic.sol", "Arithmetic")
	if err != nil {
		t.Fatal(err)
	}
	empty, _ := NewArithmeticEngine(nil, ArithmeticEngineOptions{})
	input, _ := miner.EncodeCall(p, "add", []string{"2", "3"})
	cfg := &ArithmeticConfig{Engine: empty, Mode: ArithmeticSync}
	want := runArithmeticTest(t, p.Code, input, 100000, nil)
	assertArithmeticSame(t, want, runArithmeticTest(t, p.Code, input, 100000, cfg))
	assertArithmeticSame(t, want, runArithmeticTest(t, p.Code, input, 100000, cfg))
	if empty.Stats().CacheMisses != 1 || empty.Stats().CacheHits == 0 || empty.Stats().Pruned != 0 {
		t.Fatal(empty.Stats())
	}
	var buf bytes.Buffer
	if err = empty.WriteFeedback(&buf); err != nil {
		t.Fatal(err)
	}
	var evidence miner.FeedbackBundle
	if err = json.Unmarshal(buf.Bytes(), &evidence); err != nil {
		t.Fatal(err)
	}
	learned, err := miner.MineFeedback(evidence)
	if err != nil {
		t.Fatal(err)
	}
	data, _ := json.Marshal(learned)
	next, err := LoadArithmeticEngine(bytes.NewReader(data), ArithmeticEngineOptions{})
	if err != nil {
		t.Fatal(err)
	}
	assertArithmeticSame(t, want, runArithmeticTest(t, p.Code, input, 100000, &ArithmeticConfig{Engine: next, Mode: ArithmeticSync}))
	if next.Version() == empty.Version() || next.Stats().Pruned == 0 || empty.Stats().Pruned != 0 {
		t.Fatal("library publication did not isolate old negative cache")
	}
	var wg sync.WaitGroup
	for i := 0; i < 8; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			assertArithmeticSame(t, want, runArithmeticTest(t, p.Code, input, 100000, &ArithmeticConfig{Engine: next, Mode: ArithmeticSync}))
		}()
	}
	wg.Wait()
	if next.Stats().CacheMisses != 1 {
		t.Fatal(next.Stats())
	}
}

// Reuse an actual compiled helper while replacing only the dispatcher. All
// absolute helper/cleanup/panic addresses remain those of the deployed bytecode.
func arithmeticWrapper(t testing.TB, original []byte, helper uint64, pairs [][2]uint256.Int, before, after []byte) []byte {
	t.Helper()
	prefix := append([]byte(nil), before...)
	for _, pair := range pairs {
		label := len(prefix) + 1
		prefix = append(prefix, byte(PUSH2), 0, 0)
		for _, w := range []uint256.Int{pair[1], pair[0]} {
			bytes := w.Bytes32()
			prefix = append(prefix, byte(PUSH32))
			prefix = append(prefix, bytes[:]...)
		}
		prefix = append(prefix, byte(PUSH2), byte(helper>>8), byte(helper), byte(JUMP))
		k := len(prefix)
		prefix[label], prefix[label+1] = byte(k>>8), byte(k)
		prefix = append(prefix, byte(JUMPDEST), byte(POP))
	}
	prefix = append(prefix, after...)
	report, err := miner.MatchBytecode(original, nil)
	if err != nil {
		t.Fatal(err)
	}
	// Helpers may invoke cleanup functions before their own entry.
	for _, hint := range report.Candidates {
		if hint.Reason != "unknown_closed_region" {
			continue
		}
		if uint64(len(prefix)) >= hint.Entry {
			t.Fatalf("test dispatcher overlaps helper: %d >= %d", len(prefix), hint.Entry)
		}
	}
	code := append([]byte(nil), original...)
	copy(code, prefix)
	ins, _ := miner.Disassemble(original)
	for _, i := range ins {
		if i.PC >= uint64(len(prefix)) {
			for j := len(prefix); j < int(i.PC); j++ {
				code[j] = byte(STOP)
			}
			break
		}
	}
	return code
}
func arithmeticAddFixture(t testing.TB) ([]byte, *ArithmeticEngine, uint64) {
	p, err := miner.Load("../../template_miner/testdata/arithmetic-0824.json", "Arithmetic.sol", "Arithmetic")
	if err != nil {
		t.Fatal(err)
	}
	e := testMinedEngine(t, p.Code)
	report, err := miner.MatchBytecode(p.Code, e.templates)
	if err != nil {
		t.Fatal(err)
	}
	for _, b := range report.Bindings {
		if b.Operation == "ADD" && (b.IntegerType == "" || b.IntegerType == "uint256") {
			return p.Code, e, b.Entry
		}
	}
	t.Fatal("missing add helper")
	return nil, nil, 0
}

func TestArithmeticRollbackGasEdgesAndLaterOOG(t *testing.T) {
	code, e, helper := arithmeticAddFixture(t)
	max := *new(uint256.Int).SetAllOne()
	one := *uint256.NewInt(1)
	before := []byte{byte(PUSH1), 9, byte(PUSH0), byte(SSTORE)}
	// Speculative state/log/transient writes are journalled, and a later memory
	// overflow/OOG must never replace the earlier arithmetic Panic(0x11).
	after := []byte{byte(PUSH1), 7, byte(PUSH0), byte(SSTORE), byte(PUSH1), 3, byte(PUSH0), byte(TSTORE), byte(PUSH0), byte(PUSH0), byte(LOG0), byte(PUSH32)}
	w := max.Bytes32()
	after = append(after, w[:]...)
	after = append(after, byte(MLOAD), byte(STOP))
	wrapped := arithmeticWrapper(t, code, helper, [][2]uint256.Int{{max, one}}, before, after)
	want := runArithmeticTest(t, wrapped, nil, 100000, nil)
	if want.err != ErrExecutionReverted.Error() || len(want.data) != 36 || want.data[35] != 0x11 {
		t.Fatal(want)
	}
	need := uint64(100000) - want.gas
	for _, mode := range []ArithmeticMode{ArithmeticSync} {
		for range []bool{false} {
			cfg := &ArithmeticConfig{Engine: e, Mode: mode}
			for gas := need - 15; gas <= need+30; gas++ {
				assertArithmeticSame(t, runArithmeticTest(t, wrapped, nil, gas, nil), runArithmeticTest(t, wrapped, nil, gas, cfg))
			}
			assertArithmeticSame(t, want, runArithmeticTest(t, wrapped, nil, 100000, cfg))
		}
	}
	if e.Stats().Failures == 0 {
		t.Fatalf("failure recovery not exercised: %+v", e.Stats())
	}
}

func TestArithmeticNestedCallAndDelegateCall(t *testing.T) {
	original, e, helper := arithmeticAddFixture(t)
	max := *new(uint256.Int).SetAllOne()
	one := *uint256.NewInt(1)
	child := arithmeticWrapper(t, original, helper, [][2]uint256.Int{{max, one}}, []byte{byte(PUSH1), 7, byte(PUSH0), byte(SSTORE)}, []byte{byte(STOP)})
	for _, callOp := range []OpCode{CALL, DELEGATECALL, STATICCALL} {
		// Save child's success flag and return data; then call a parent helper
		// again, exercising parent binding restoration after a different code hash.
		call := []byte{byte(PUSH0), byte(PUSH0), byte(PUSH0), byte(PUSH0)}
		if callOp == CALL {
			call = append(call, byte(PUSH0))
		}
		call = append(call, byte(PUSH2), 0x20, 0, byte(PUSH3), 1, 0x86, 0xa0, byte(callOp), byte(PUSH0), byte(MSTORE))
		call = append(call, byte(RETURNDATASIZE), byte(PUSH0), byte(PUSH1), 32, byte(RETURNDATACOPY))
		parent := arithmeticWrapper(t, original, helper, [][2]uint256.Int{{one, one}}, call, []byte{byte(PUSH1), 68, byte(PUSH0), byte(RETURN)})
		run := func(cfg *ArithmeticConfig) arithmeticTestRun {
			evm, s, addr := testArithmeticEVM(t, parent, cfg)
			s.SetCode(common.HexToAddress("0x2000"), child, tracing.CodeChangeUnspecified)
			ret, gas, err := evm.Call(common.Address{}, addr, nil, 250000, new(uint256.Int))
			errText := ""
			if err != nil {
				errText = err.Error()
			}
			if s.GetTransientState(addr, common.Hash{}) != (common.Hash{}) {
				t.Fatal("transient state leaked")
			}
			return arithmeticTestRun{append([]byte(nil), ret...), gas, errText, s.IntermediateRoot(true), s.GetRefund(), len(s.Logs())}
		}
		want := run(nil)
		assertArithmeticSame(t, want, run(&ArithmeticConfig{Engine: e, Mode: ArithmeticSync}))
		for _, mode := range []ArithmeticMode{ArithmeticSync} {
			assertArithmeticSame(t, want, run(&ArithmeticConfig{Engine: e, Mode: mode}))
		}
	}
}

func TestArithmeticRetainsGuardsBeyondFiniteMiningWitnesses(t *testing.T) {
	code, _, helper := arithmeticAddFixture(t)
	ins, _ := miner.Disassemble(code)
	last := ins[len(ins)-1]
	end := int(last.PC) + 1 + len(last.Data)
	code = append(code, make([]byte, end-len(code))...)
	entry := len(code)
	panicPC := entry + 15
	// An extra business predicate, chosen outside the miner's finite witness
	// set, deliberately masquerades as arithmetic Panic(0x11).
	code = append(code, byte(JUMPDEST), byte(DUP1), byte(PUSH3), 0x12, 0x34, 0x56, byte(EQ), byte(PUSH2), byte(panicPC>>8), byte(panicPC), byte(JUMPI), byte(PUSH2), byte(helper>>8), byte(helper), byte(JUMP))
	code = append(code, byte(JUMPDEST), byte(PUSH4), 0x4e, 0x48, 0x7b, 0x71, byte(PUSH1), 224, byte(SHL), byte(PUSH0), byte(MSTORE), byte(PUSH1), 0x11, byte(PUSH1), 4, byte(MSTORE), byte(PUSH1), 36, byte(PUSH0), byte(REVERT))
	e := testMinedEngine(t, code)
	wrapped := arithmeticWrapper(t, code, uint64(entry), [][2]uint256.Int{{*uint256.NewInt(0x123456), *uint256.NewInt(1)}}, nil, []byte{byte(STOP)})
	want := runArithmeticTest(t, wrapped, nil, 100000, nil)
	if want.err != ErrExecutionReverted.Error() {
		t.Fatal("invalid guard fixture")
	}
	for _, mode := range []ArithmeticMode{ArithmeticSync} {
		assertArithmeticSame(t, want, runArithmeticTest(t, wrapped, nil, 100000, &ArithmeticConfig{Engine: e, Mode: mode}))
	}
	if e.Stats().Failures == 0 {
		t.Fatal("extra guard was not exercised")
	}
	// The nontriggering input must still execute successfully with identical gas.
	ordinary := arithmeticWrapper(t, code, uint64(entry), [][2]uint256.Int{{*uint256.NewInt(9), *uint256.NewInt(1)}}, nil, []byte{byte(STOP)})
	assertArithmeticSame(t, runArithmeticTest(t, ordinary, nil, 100000, nil), runArithmeticTest(t, ordinary, nil, 100000, &ArithmeticConfig{Engine: e, Mode: ArithmeticSync}))
}
