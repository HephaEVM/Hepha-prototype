package vm

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/holiman/uint256"
)

// Compare the preserved executor with the fused executor, including source
// aliases, constants, cyclic output moves, stack growth, and failing admission.
// The compiler/legacy corpus tests separately compare complete EVM outcomes.
func TestArithmeticSyncFusedLayouts(t *testing.T) {
	code, engine, helper := arithmeticAddFixture(t)
	base := engine.bind(code, arithmeticCacheKey{hash: crypto.Keccak256Hash(code), rules: &cancunInstructionSet}).regions[helper]
	if base.plan.native == nil {
		t.Fatal("fixture has no native plan")
	}
	constSource := stackSource{index: -1, constant: *uint256.NewInt(17)}
	prefixes := [][]stackSource{nil, {{index: 0}}, {{index: 5}}, {constSource}, {{index: 5}, {index: 4}}, {{index: 5}, {index: 5}}, {constSource, {index: 1}}, {{index: 5}, {index: 4}, {index: 3}}}
	large := make([]stackSource, 48)
	for i := range large {
		large[i] = stackSource{index: (47 - i) % 6}
	}
	prefixes = append(prefixes, large)
	for mask := uint8(0); mask < 8; mask++ {
		for pi, prefix := range prefixes {
			for _, scenario := range []string{"success", "overflow", "lowgas", "badjump", "underflow", "peak", "uint8_noncanonical", "int8_noncanonical"} {
				t.Run(fmt.Sprintf("sources%d/prefix%d/%s", mask, pi, scenario), func(t *testing.T) {
					r := *base
					r.inputCount, r.peak = 6, 64
					r.a, r.b, r.k = stackSource{index: 0}, stackSource{index: 1}, stackSource{index: 2}
					r.prefix = prefix
					r.prefixIdentity = true
					for i, s := range prefix {
						if s.index != i {
							r.prefixIdentity = false
						}
					}
					words := make([]uint256.Int, 9)
					for i := range words {
						words[i].SetUint64(uint64(i + 1))
					}
					input := words[3:]
					input[2].SetUint64(helper)
					gas := uint64(100000)
					switch scenario {
					case "overflow":
						input[0].SetAllOne()
					case "lowgas":
						gas = 1
					case "badjump":
						input[2].SetAllOne()
					case "peak":
						r.peak = 1024
					case "uint8_noncanonical":
						r.domain, _ = ParseIntegerDomain("uint8")
						input[0].SetUint64(256)
					case "int8_noncanonical":
						r.domain, _ = ParseIntegerDomain("int8")
						input[0].SetUint64(128)
					}
					for i, source := range []*stackSource{&r.a, &r.b, &r.k} {
						if mask&(1<<i) != 0 {
							*source = stackSource{index: -1, constant: input[i]}
						}
					}
					r.constantContinuation = mask&4 != 0 && scenario != "badjump"
					r.prepareSyncLayout()
					r.prepareSyncHot()
					if scenario == "underflow" {
						words = words[:2]
					}
					type snapshot struct {
						stack                                   []uint256.Int
						gas, pc                                 uint64
						skipped                                 bool
						failure                                 *arithmeticResult
						seq, pruned, checked, native, fallbacks uint64
					}
					run := func(mode int) snapshot {
						f := &arithmeticFrame{config: ArithmeticConfig{Engine: engine, Mode: ArithmeticSync}}
						contract := NewContract(common.Address{}, common.Address{}, new(uint256.Int), gas, nil)
						contract.setResolvedCallCode(crypto.Keccak256Hash(code), code)
						scope := &ScopeContext{Contract: contract, Stack: &Stack{data: append([]uint256.Int(nil), words...)}, Memory: NewMemory()}
						defer scope.Memory.Free()
						pc := helper
						var ok bool
						var failure *arithmeticResult
						if mode == 2 {
							ok, failure = f.trySyncHot(&pc, scope, &r)
						} else if mode == 1 {
							ok, failure = f.trySync(&pc, scope, &r)
						} else {
							ok, failure = f.try(&pc, scope, &r)
						}
						return snapshot{scope.Stack.data, contract.Gas, pc, ok, failure, f.sequence, f.pruned, f.checked, f.nativeChecked, f.fallbacks}
					}
					want, got := run(0), run(1)
					if !reflect.DeepEqual(want, got) {
						t.Fatalf("reference=%+v\nfused=%+v", want, got)
					}
					if got := run(2); !reflect.DeepEqual(want, got) {
						t.Fatalf("reference=%+v\ntuned=%+v", want, got)
					}
				})
			}
		}
	}
}

func BenchmarkArithmeticSyncFusion(b *testing.B) {
	code, engine, helper := arithmeticAddFixture(b)
	bound := engine.bind(code, arithmeticCacheKey{hash: crypto.Keccak256Hash(code), rules: &cancunInstructionSet})
	for _, extended := range []bool{false, true} {
		r := bound.regions[helper]
		if extended {
			first := ^uint64(0)
			for _, candidate := range bound.regions {
				if candidate.entry != candidate.helper && candidate.entry < first && candidate.op == "ADD" && candidate.domain.Bits == 256 && candidate.plan.native != nil {
					r = candidate
					first = candidate.entry
				}
			}
			if r.entry == r.helper {
				b.Fatal("missing extended fixture")
			}
		}
		for _, mode := range []string{"generic", "fused", "tuned"} {
			b.Run(fmt.Sprintf("extended_%t/%s", extended, mode), func(b *testing.B) {
				f := &arithmeticFrame{config: ArithmeticConfig{Engine: engine, Mode: ArithmeticSync}}
				contract := NewContract(common.Address{}, common.Address{}, new(uint256.Int), 100000, nil)
				contract.setResolvedCallCode(crypto.Keccak256Hash(code), code)
				scope := &ScopeContext{Contract: contract, Stack: newstack(), Memory: NewMemory()}
				defer returnStack(scope.Stack)
				defer scope.Memory.Free()
				b.ReportAllocs()
				for i := 0; i < b.N; i++ {
					scope.Stack.data = scope.Stack.data[:r.inputCount]
					if r.a.index >= 0 {
						scope.Stack.data[r.a.index].SetUint64(9)
					}
					if r.b.index >= 0 {
						scope.Stack.data[r.b.index].SetUint64(7)
					}
					if r.k.index >= 0 {
						scope.Stack.data[r.k.index].SetUint64(helper)
					}
					contract.Gas = 100000
					pc := r.entry
					var ok bool
					var failure *arithmeticResult
					if mode == "tuned" {
						ok, failure = f.trySyncHot(&pc, scope, r)
					} else if mode == "fused" {
						ok, failure = f.trySync(&pc, scope, r)
					} else {
						ok, failure = f.try(&pc, scope, r)
					}
					if !ok || failure != nil {
						b.Fatal(ok, failure)
					}
				}
			})
		}
	}
}
