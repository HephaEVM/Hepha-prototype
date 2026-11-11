package templateminer

import (
	"fmt"
	"sort"
	"strings"

	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/holiman/uint256"
)

// Paths are indexed by 2*wide + nonzeroRemainder. Each includes the complete
// inner mulDiv, remainder test, overflow guard, increment and final return.
// Failure leaves are never replaced by the runtime.
type MulDivRoundingRegion struct {
	Entry, Inner uint64
	A, B, D      int
	Paths        [4]ABIRegionPath
	StructuralID string
}

func guardSet(xs []string) string {
	seen := map[string]bool{}
	for _, x := range xs {
		seen[x] = true
	}
	keys := make([]string, 0, len(seen))
	for x := range seen {
		keys = append(keys, x)
	}
	sort.Strings(keys)
	return strings.Join(keys, ",")
}

func AnalyzeMulDivRounding(p *Program, entry, inner uint64) (*MulDivRoundingRegion, error) {
	floor, err := AnalyzeMulDiv(p, inner)
	if err != nil {
		return nil, fmt.Errorf("inner floor proof: %w", err)
	}
	r, err := analyzeABIRegion(p, entry, 4, true)
	if err != nil {
		return nil, err
	}
	var paths []ABIRegionPath
	for _, path := range r.Paths {
		if path.Outcome == "return" {
			paths = append(paths, path)
		}
	}
	if len(paths) != 4 {
		return nil, fmt.Errorf("expected four rounding success paths, got %d", len(paths))
	}
	for _, perm := range [][3]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}} {
		short, wide, sg, wg := wordSpec(perm[0], perm[1], perm[2])
		rem := wordOp("MULMOD", wordInput(perm[0]), wordInput(perm[1]), wordInput(perm[2]))
		nonzeroD := wordGuard(wordInput(perm[2]), true)
		out := &MulDivRoundingRegion{Entry: entry, Inner: inner, A: perm[0], B: perm[1], D: perm[2], StructuralID: r.ID}
		valid := true
		for _, path := range paths {
			x, g, ok := mulDivPathShape(path)
			if !ok || x == nil {
				valid = false
				break
			}
			matched := -1
			for index := 0; index < 4; index++ {
				q, guards, innerPCs := short, sg, floor.Short.PCs
				if index >= 2 {
					q, guards, innerPCs = wide, wg, floor.Wide.PCs
				}
				// An actual, complete inner helper must occur exactly once.
				// This also makes primitive/fused instruction counters additive.
				if !containsMulDivPathOnce(path.PCs, innerPCs, inner) {
					continue
				}
				expected := append([]string(nil), guards...)
				expected = append(expected, wordGuard(rem, index%2 == 1))
				answer := q
				if index%2 == 1 {
					var max uint256.Int
					max.SetAllOne()
					expected = append(expected, wordGuard(wordOp("LT", q, wordConst(max)), true))
					answer = wordOp("ADD", q, wordSmall(1))
				}
				// Solidity may retain an extra divisor != 0 guard for MULMOD.
				// It is implied by either floor-success predicate.
				withD := append(append([]string(nil), expected...), nonzeroD)
				if x.key == answer.key && (guardSet(g) == guardSet(expected) || guardSet(g) == guardSet(withD)) {
					matched = index
					break
				}
			}
			if matched < 0 || len(out.Paths[matched].PCs) > 0 {
				valid = false
				break
			}
			out.Paths[matched] = path
		}
		if valid {
			return out, nil
		}
	}
	return nil, fmt.Errorf("complete expressions/guards are not checked ceil(a*b/d)")
}

func containsMulDivPathOnce(path, innerPath []uint64, entry uint64) bool {
	at, count := -1, 0
	for i, pc := range path {
		if pc == entry {
			at, count = i, count+1
		}
	}
	if count != 1 || at+len(innerPath) > len(path) {
		return false
	}
	for i, pc := range innerPath {
		if path[at+i] != pc {
			return false
		}
	}
	return true
}

// Static direct calls to already proved floor helpers are discovery hints.
// The nearby MULMOD only filters work; AnalyzeMulDivRounding is the authority.
func MulDivRoundingCandidates(p *Program, floors map[uint64]*MulDivRegion) map[uint64]uint64 {
	out := map[uint64]uint64{}
	for i, in := range p.Instructions {
		if in.Op != vm.JUMP || i == 0 {
			continue
		}
		push := p.Instructions[i-1]
		if push.Op < vm.PUSH1 || push.Op > vm.PUSH32 {
			continue
		}
		var dest uint256.Int
		dest.SetBytes(push.Data)
		if !dest.IsUint64() || floors[dest.Uint64()] == nil {
			continue
		}
		mod := false
		for j := i + 1; j < len(p.Instructions) && j <= i+32; j++ {
			mod = mod || p.Instructions[j].Op == vm.MULMOD
		}
		if !mod {
			continue
		}
		for j := i - 2; j >= 0 && j >= i-24; j-- {
			op := p.Instructions[j].Op
			if op == vm.JUMPDEST {
				out[p.Instructions[j].PC] = dest.Uint64()
				break
			}
			if op == vm.JUMP || op == vm.JUMPI || op == vm.REVERT || op == vm.STOP || op == vm.RETURN {
				break
			}
		}
	}
	return out
}
