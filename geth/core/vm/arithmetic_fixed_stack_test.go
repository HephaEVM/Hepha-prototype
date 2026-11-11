package vm

import (
	"reflect"
	"testing"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/holiman/uint256"
)

func TestArithmeticFixedBindingProof(t *testing.T) {
	_, _, source := deployedArithmeticHelper(t, 0)
	if source.syncHot.fixedLayout != 2 {
		t.Fatal("expected continuation,b,a layout")
	}
	for _, alter := range []func(*arithmeticRegion){
		func(r *arithmeticRegion) { r.inputCount = 4 },
		func(r *arithmeticRegion) { r.prefix = []stackSource{{index: 0}} },
		func(r *arithmeticRegion) { r.domain.Bits = 128 },
		func(r *arithmeticRegion) { r.constantContinuation = true },
		func(r *arithmeticRegion) { r.k.index = -1 },
		func(r *arithmeticRegion) { r.a.index = r.b.index },
	} {
		r := *source
		alter(&r)
		r.prepareSyncHot()
		if r.syncHot.fixedLayout != 0 {
			t.Fatal("unsupported layout accepted")
		}
	}
}

func TestArithmeticFixedSyncAdmission(t *testing.T) {
	code, e, source := deployedArithmeticHelper(t, 0)
	for _, layout := range []uint8{1, 2} {
		r := *source
		r.a.index, r.b.index = 1, 2
		if layout == 2 {
			r.a.index, r.b.index = 2, 1
		}
		r.prepareSyncLayout()
		r.prepareSyncHot()
		if r.syncHot.fixedLayout != layout {
			t.Fatal("wrong fixed layout")
		}
		for _, n := range []int{0, 1, 2, 3, 5, r.syncHot.maxStack, r.syncHot.maxStack + 1} {
			for _, badJump := range []bool{false, true} {
				for _, gas := range []uint64{r.syncHot.admissionGas - 1, r.syncHot.admissionGas, 10000000} {
					for _, overflow := range []bool{false, true} {
						var wantStack []uint256.Int
						var wantPC, wantGas, wantSequence uint64
						var wantSkip bool
						var wantFailure *arithmeticResult
						for _, fixed := range []bool{false, true} {
							contract := NewContract([20]byte{}, [20]byte{}, new(uint256.Int), gas, nil)
							contract.setResolvedCallCode(crypto.Keccak256Hash(code), code)
							stack := &Stack{data: make([]uint256.Int, n, max(n, 1024))}
							for i := range stack.data {
								stack.data[i].SetUint64(uint64(100 + i))
							}
							if n >= 3 {
								stack.data[n-3].SetUint64(r.helper)
								if badJump {
									stack.data[n-3][3] = 1
								}
								stack.data[n-3+r.a.index], stack.data[n-3+r.b.index] = uint256.Int{7}, uint256.Int{9}
								if overflow {
									stack.data[n-3+r.a.index] = uint256.Int{0, 0, 0, uint64(1) << 63}
									stack.data[n-3+r.b.index].SetAllOne()
								}
							}
							memory := NewMemory()
							f := &arithmeticFrame{config: *deployedArithmeticConfig(e, "sync-reduced")}
							pc := r.helper
							scope := &ScopeContext{Stack: stack, Contract: contract, Memory: memory}
							var skip bool
							var failure *arithmeticResult
							if fixed {
								skip, failure = f.trySyncFixed(&pc, scope, &r)
							} else {
								skip, failure = f.trySyncHot(&pc, scope, &r)
							}
							if !fixed {
								wantStack = append([]uint256.Int{}, stack.data...)
								wantPC, wantGas, wantSequence, wantSkip, wantFailure = pc, contract.Gas, f.sequence, skip, failure
							} else if !reflect.DeepEqual(wantStack, stack.data) || pc != wantPC || contract.Gas != wantGas || f.sequence != wantSequence || skip != wantSkip || !reflect.DeepEqual(failure, wantFailure) {
								t.Fatalf("fixed mismatch layout=%d n=%d gas=%d badJump=%t overflow=%t", layout, n, gas, badJump, overflow)
							}
							memory.Free()
						}
					}
				}
			}
		}
	}
}

func TestArithmeticFixedLegacyCorpus(t *testing.T) {
	testArithmeticSyncLegacyCorpus(t, true, true)
}
