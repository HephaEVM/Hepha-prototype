package vm

import "github.com/ethereum/go-ethereum/common"

type metadataJumpCache struct {
	shared map[common.Hash]BitVec
	local  mapJumpDests
}

func metadataNewJumpCache(seed map[common.Hash]BitVec) JumpDestCache {
	if seed == nil {
		return newMapJumpDests()
	}
	return &metadataJumpCache{seed, make(mapJumpDests)}
}
func (j *metadataJumpCache) Load(h common.Hash) (BitVec, bool) {
	if b, ok := j.shared[h]; ok {
		return b, true
	}
	return j.local.Load(h)
}
func (j *metadataJumpCache) Store(h common.Hash, b BitVec) { j.local.Store(h, b) }
