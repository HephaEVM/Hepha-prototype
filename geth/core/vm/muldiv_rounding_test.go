package vm

import (
	"encoding/hex"
	"math/big"
	"math/rand"
	"os"
	"reflect"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/holiman/uint256"
)

func roundingDeployedCode(t testing.TB, calls int) []byte {
	t.Helper()
	raw, err := os.ReadFile("../../template_miner/testdata/muldiv-rounding-deployed.hex")
	if err != nil {
		t.Fatal(err)
	}
	code, err := hex.DecodeString(strings.TrimSpace(string(raw)))
	if err != nil {
		t.Fatal(err)
	}
	if crypto.Keccak256Hash(code).Hex() != "0x61bb013bc7f8bed8e3c050464e93985787d254f4236d2e3ef846700c07e00aed" {
		t.Fatal("deployed rounding fixture hash changed")
	}
	// Preserve all deployed helper PCs. Only the synthetic caller/dispatcher is
	// changed. This is a differential fixture, not a historical transaction.
	entry := len(code)
	code[0], code[1], code[2], code[3], code[4] = byte(PUSH3), byte(entry>>16), byte(entry>>8), byte(entry), byte(JUMP)
	code = append(code, byte(JUMPDEST))
	for i := 0; i < calls; i++ {
		start := len(code)
		ret := start + 15
		code = append(code, byte(PUSH2), byte(ret>>8), byte(ret),
			byte(PUSH0), byte(CALLDATALOAD), byte(PUSH1), 32, byte(CALLDATALOAD),
			byte(PUSH1), 64, byte(CALLDATALOAD), byte(PUSH2), 0x41, 0xa9, byte(JUMP),
			byte(JUMPDEST))
		if i+1 < calls {
			code = append(code, byte(POP))
		}
	}
	return append(code, byte(PUSH0), byte(MSTORE), byte(PUSH1), 32, byte(PUSH0), byte(RETURN))
}

func roundingInputs() [][3]uint256.Int {
	max := *new(uint256.Int).SetAllOne()
	m1 := *new(uint256.Int).SubUint64(&max, 1)
	m2 := *new(uint256.Int).SubUint64(&max, 2)
	half := *new(uint256.Int).Lsh(uint256.NewInt(1), 255)
	return [][3]uint256.Int{
		{*uint256.NewInt(6), *uint256.NewInt(9), *uint256.NewInt(3)}, // short, exact
		{*uint256.NewInt(7), *uint256.NewInt(9), *uint256.NewInt(2)}, // short, ceil
		{max, max, max},                // wide, exact
		{half, half, max},              // wide, ceil
		{m1, m1, m2},                   // floor == MAX, remainder == 1
		{max, max, *uint256.NewInt(1)}, // floor overflow
		{max, max, {}}, {}, {*uint256.NewInt(1), {}, *uint256.NewInt(1)},
	}
}

func TestMulDivRoundingIndependentSpec(t *testing.T) {
	cases := roundingInputs()
	rng := rand.New(rand.NewSource(20261020))
	for i := 0; i < 10000; i++ {
		var c [3]uint256.Int
		for j := range c {
			c[j] = uint256.Int{rng.Uint64(), rng.Uint64(), rng.Uint64(), rng.Uint64()}
			if i%2 == 0 {
				c[j].Rsh(&c[j], uint(rng.Intn(257)))
			}
		}
		cases = append(cases, c)
	}
	for i, c := range cases {
		original := c
		got, _, _, ok := mulDivRoundingResult(&c[0], &c[1], &c[2])
		want := new(big.Int).Mul(c[0].ToBig(), c[1].ToBig())
		valid := !c[2].IsZero()
		if valid {
			r := new(big.Int)
			want.QuoRem(want, c[2].ToBig(), r)
			if r.Sign() != 0 {
				want.Add(want, big.NewInt(1))
			}
			valid = want.BitLen() <= 256
		}
		if ok != valid || ok && got.ToBig().Cmp(want) != 0 || c != original {
			t.Fatalf("case %d got %s valid=%v expected %s valid=%v", i, got.Dec(), ok, want, valid)
		}
	}
}

func TestMulDivRoundingFullEVM(t *testing.T) {
	code := roundingDeployedCode(t, 1)
	e := computeTestEngine(t)
	stats := new(ComputeHelperStats)
	cfg := &ArithmeticConfig{Engine: e, Mode: ArithmeticSync, ComputeStats: stats}
	n := 0
	for _, c := range roundingInputs() {
		input := computeInput(c[:]...)
		for gas := uint64(0); gas < 1100; gas++ {
			assertArithmeticSame(t, runArithmeticTest(t, code, input, gas, nil), runArithmeticTest(t, code, input, gas, cfg))
			n++
		}
	}
	if stats.MulDivFusions == 0 || stats.MulDivRoundUps == 0 || stats.MulDivWide == 0 || stats.MulDivFusionFallbacks == 0 {
		t.Fatal(*stats)
	}
	key := arithmeticCacheKey{hash: crypto.Keccak256Hash(code), rules: &cancunInstructionSet, compute: 1}
	floor := e.bind(code, key)
	key.compute = 5
	fused := e.bind(code, key)
	if floor == fused || floor.entryIndex[16809] != 0 || fused.entryIndex[16809] == 0 {
		t.Fatal("fusion policy aliased")
	}
	r := fused.computeEntries[fused.entryIndex[16809]]
	if r.rounding == nil {
		t.Fatal("missing wrapper")
	}
	for i, path := range r.rounded {
		t.Log("path", i, "instructions", path.instructions, "gas", path.gas, "peak", path.peak)
	}
	t.Log("full EVM gas-boundary comparisons", n, "stats", *stats)
}

func TestMulDivRoundingAtomicAdmission(t *testing.T) {
	code := roundingDeployedCode(t, 1)
	e := computeTestEngine(t)
	b := e.bind(code, arithmeticCacheKey{hash: crypto.Keccak256Hash(code), rules: &cancunInstructionSet, compute: 5})
	r := b.computeEntries[b.entryIndex[16809]]
	for _, bad := range []string{"short", "return", "peak", "gas", "ceil-overflow"} {
		c := NewContract(common.Address{}, common.Address{}, new(uint256.Int), 100000, nil)
		c.setResolvedCallCode(crypto.Keccak256Hash(code), code)
		words := roundingInputs()[1]
		if bad == "ceil-overflow" {
			words = roundingInputs()[4]
		}
		data := []uint256.Int{{16809}, words[0], words[1], words[2]}
		switch bad {
		case "short":
			data = data[:3]
		case "return":
			data[0].SetUint64(1)
		case "peak":
			data = append(make([]uint256.Int, 1010), data...)
		case "gas":
			c.Gas = 1
		}
		before := append([]uint256.Int(nil), data...)
		gas, pc := c.Gas, uint64(16809)
		s := &ScopeContext{Contract: c, Stack: &Stack{data: data}}
		if r.tryMulDivRounding(&pc, s, nil) || pc != 16809 || c.Gas != gas || !reflect.DeepEqual(s.Stack.data, before) {
			t.Fatal("non-atomic fallback", bad)
		}
	}
}

func BenchmarkMulDivRoundingComposition(b *testing.B) {
	for i, words := range roundingInputs()[:4] {
		for _, mode := range []string{"original", "floor-only", "fused"} {
			b.Run([]string{"short-exact", "short-round", "wide-exact", "wide-round"}[i]+"/"+mode, func(b *testing.B) {
				code := roundingDeployedCode(b, 64)
				var cfg *ArithmeticConfig
				if mode != "original" {
					cfg = &ArithmeticConfig{Engine: computeTestEngine(b), Mode: ArithmeticSync}
				}
				evm, _, _ := testArithmeticEVM(b, code, cfg)
				contract := NewContract(common.Address{}, common.Address{}, new(uint256.Int), 10000000, nil)
				contract.setResolvedCallCode(crypto.Keccak256Hash(code), code)
				input := computeInput(words[:]...)
				want, _, _, ok := mulDivRoundingResult(&words[0], &words[1], &words[2])
				if !ok {
					b.Fatal("invalid benchmark input")
				}
				result := want.Bytes32()
				ret, err, _ := evm.Run(contract, input, false)
				if err != nil || string(ret) != string(result[:]) {
					b.Fatal(err)
				}
				b.ReportAllocs()
				b.ResetTimer()
				for n := 0; n < b.N; n++ {
					contract.Gas = 10000000
					if _, err, _ := evm.Run(contract, input, false); err != nil {
						b.Fatal(err)
					}
				}
			})
		}
	}
}
