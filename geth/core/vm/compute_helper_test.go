package vm

import (
	"encoding/hex"
	"math/rand"
	"os"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/holiman/uint256"
	"go-eth1.16.5-evm/core/tracing"
	miner "go-eth1.16.5-evm/template_miner"
)

func computeTestEngine(t testing.TB) *ArithmeticEngine {
	t.Helper()
	e, err := NewArithmeticEngine(nil, ArithmeticEngineOptions{CacheEntries: 1024})
	if err != nil {
		t.Fatal(err)
	}
	return e
}
func computeMulDivCode(t *testing.T) []byte {
	t.Helper()
	raw, e := os.ReadFile("../../template_miner/testdata/muldiv-deployed.hex")
	if e != nil {
		t.Fatal(e)
	}
	body, e := hex.DecodeString(strings.TrimSpace(string(raw)))
	if e != nil {
		t.Fatal(e)
	}
	const entry = 256
	ins, e := miner.Disassemble(body)
	if e != nil {
		t.Fatal(e)
	}
	for _, in := range ins {
		if OpCode(in.Op) == PUSH2 {
			dest := int(in.Data[0])*256 + int(in.Data[1])
			if dest >= len(body) || body[dest] != byte(JUMPDEST) {
				t.Fatal("unexpected fixture PUSH2")
			}
			dest += entry
			body[in.PC+1] = byte(dest >> 8)
			body[in.PC+2] = byte(dest)
		}
	}
	code := []byte{byte(PUSH2), 0, 0, byte(PUSH0), byte(CALLDATALOAD), byte(PUSH1), 32, byte(CALLDATALOAD), byte(PUSH1), 64, byte(CALLDATALOAD), byte(PUSH2), 1, 0, byte(JUMP)}
	ret := len(code)
	code[1], code[2] = byte(ret>>8), byte(ret)
	code = append(code, byte(JUMPDEST), byte(PUSH0), byte(MSTORE), byte(PUSH1), 32, byte(PUSH0), byte(RETURN))
	code = append(code, make([]byte, entry-len(code))...)
	return append(code, body...)
}
func computeInput(words ...uint256.Int) []byte {
	var data []byte
	for _, w := range words {
		b := w.Bytes32()
		data = append(data, b[:]...)
	}
	return data
}
func TestComputeMulDivFullEVM(t *testing.T) {
	code := computeMulDivCode(t)
	e := computeTestEngine(t)
	stats := new(ComputeHelperStats)
	cfg := &ArithmeticConfig{Engine: e, Mode: ArithmeticSync, ComputeStats: stats}
	max := *new(uint256.Int).SetAllOne()
	half := *new(uint256.Int).Lsh(uint256.NewInt(1), 255)
	cases := [][3]uint256.Int{{*uint256.NewInt(7), *uint256.NewInt(9), *uint256.NewInt(2)}, {max, max, max}, {half, half, max}, {max, max, *uint256.NewInt(1)}, {max, max, {}}, {}, {*uint256.NewInt(1), {}, *uint256.NewInt(1)}}
	rng := rand.New(rand.NewSource(20261020))
	for i := 0; i < 128; i++ {
		var c [3]uint256.Int
		for j := range c {
			c[j] = uint256.Int{rng.Uint64(), rng.Uint64(), rng.Uint64(), rng.Uint64()}
			c[j].Rsh(&c[j], uint(rng.Intn(257)))
		}
		cases = append(cases, c)
	}
	n := 0
	for _, c := range cases {
		input := computeInput(c[:]...)
		for _, gas := range []uint64{0, 1, 30, 100, 200, 400, 600, 1000, 100000} {
			want := runArithmeticTest(t, code, input, gas, nil)
			got := runArithmeticTest(t, code, input, gas, cfg)
			assertArithmeticSame(t, want, got)
			n++
		}
	}
	// Sweep every gas amount near the helper's branches, including success
	// followed by an OOG in the caller and return-data memory expansion.
	for _, c := range cases[:3] {
		for gas := uint64(100); gas < 800; gas++ {
			input := computeInput(c[:]...)
			assertArithmeticSame(t, runArithmeticTest(t, code, input, gas, nil), runArithmeticTest(t, code, input, gas, cfg))
			n++
		}
	}
	if stats.MulDiv == 0 || stats.MulDivWide == 0 || stats.MulDivFallbacks == 0 {
		t.Fatal(*stats)
	}
	if e.Stats().FrameContexts != 0 {
		t.Fatal("compute-only code allocated arithmetic frames")
	}
	t.Log("full EVM cases", n, "stats", *stats)
}

func TestComputeHelperPolicyIsolation(t *testing.T) {
	code := computeMulDivCode(t)
	e := computeTestEngine(t)
	key := arithmeticCacheKey{hash: crypto.Keccak256Hash(code), rules: &cancunInstructionSet}
	old := e.bind(code, key)
	key.compute = 1
	active := e.bind(code, key)
	if old == active || old.computeEntries != nil || active.computeEntries == nil {
		t.Fatal("cache policy alias")
	}
	if e.bind(code, key) != active {
		t.Fatal("cache missed")
	}
	cfg := ArithmeticConfig{Mode: ArithmeticSync}
	if cfg.computeHelperPolicy() != 5 {
		t.Fatal("disabled")
	}
	cfg.Mode = ArithmeticMode(255)
	if cfg.computeHelperPolicy() != 0 {
		t.Fatal("async admitted")
	}
	cfg.Mode = ArithmeticSync
	cfg.GasPolicy = ArithmeticGasPolicy(255)
	if cfg.computeHelperPolicy() != 0 {
		t.Fatal("replacement gas admitted")
	}
}

func TestComputeMulDivNestedFailure(t *testing.T) {
	child := computeMulDivCode(t)
	// Mutate storage before entering the helper; a rejected division must
	// revert this child write while the caller handles failure and continues.
	prefix := []byte{byte(PUSH1), 9, byte(PUSH0), byte(SSTORE)}
	child = append(append(append([]byte(nil), prefix...), child[:256-len(prefix)]...), child[256:]...)
	child[5], child[6] = 0, byte(15+len(prefix))
	childAddr := common.HexToAddress("0x2000")
	e := computeTestEngine(t)
	stats := new(ComputeHelperStats)
	cfg := &ArithmeticConfig{Engine: e, Mode: ArithmeticSync, ComputeStats: stats}
	max := *new(uint256.Int).SetAllOne()
	inputs := [][]byte{computeInput(*uint256.NewInt(9), *uint256.NewInt(8), *uint256.NewInt(7)), computeInput(max, max, max), computeInput(max, max, *uint256.NewInt(1)), computeInput(max, max, uint256.Int{})}
	for _, childGas := range []uint64{100, 22000, 24000, 100000} {
		// Copy 96 bytes of inputs, call child, then return output, success,
		// remaining gas and return-data length; also commit a caller write.
		parent := []byte{byte(PUSH1), 96, byte(PUSH0), byte(PUSH0), byte(CALLDATACOPY), byte(PUSH1), 32, byte(PUSH0), byte(PUSH1), 96, byte(PUSH0), byte(PUSH0), byte(PUSH2), 0x20, 0, byte(PUSH3), byte(childGas >> 16), byte(childGas >> 8), byte(childGas), byte(CALL), byte(PUSH1), 96, byte(MSTORE), byte(GAS), byte(PUSH1), 128, byte(MSTORE), byte(RETURNDATASIZE), byte(PUSH1), 160, byte(MSTORE), byte(PUSH1), 1, byte(PUSH0), byte(SSTORE), byte(PUSH1), 192, byte(PUSH0), byte(RETURN)}
		run := func(input []byte, cfg *ArithmeticConfig) arithmeticTestRun {
			evm, s, addr := testArithmeticEVM(t, parent, cfg)
			s.CreateAccount(childAddr)
			s.SetCode(childAddr, child, tracing.CodeChangeUnspecified)
			s.SetState(childAddr, common.Hash{}, common.HexToHash("0x7"))
			data, gas, err := evm.Call(common.Address{}, addr, input, 1000000, new(uint256.Int))
			text := ""
			if err != nil {
				text = err.Error()
			}
			return arithmeticTestRun{data, gas, text, s.IntermediateRoot(true), s.GetRefund(), len(s.Logs())}
		}
		for _, input := range inputs {
			assertArithmeticSame(t, run(input, nil), run(input, cfg))
		}
	}
	if stats.MulDiv == 0 || stats.MulDivWide == 0 || stats.MulDivFallbacks == 0 {
		t.Fatal("nested paths not exercised", *stats)
	}
}
