package vm

import (
	"encoding/binary"
	"github.com/ethereum/go-ethereum/common"
	"sync/atomic"
)

// StreamRegions shares immutable jump information across EVM instances while
// stack, gas and memory remain private to each frame. Its zero value is ready to use.
type StreamRegions struct {
	slots [2048]atomic.Pointer[streamJumpBits]
}

type streamJumpBits struct {
	hash common.Hash
	bits BitVec
}

func (s *StreamRegions) slot(hash common.Hash) *atomic.Pointer[streamJumpBits] {
	return &s.slots[binary.BigEndian.Uint16(hash[:2])&uint16(len(s.slots)-1)]
}

func (s *StreamRegions) Load(hash common.Hash) (BitVec, bool) {
	record := s.slot(hash).Load()
	if record != nil && record.hash == hash {
		return record.bits, true
	}
	return nil, false
}

func (s *StreamRegions) Store(hash common.Hash, bits BitVec) {
	if len(bits) <= 4096 {
		s.slot(hash).Store(&streamJumpBits{hash, bits})
	}
}

// WithStreamRegions enables direct execution of contiguous supported operations.
// An existing template preparer continues to own its published entry points.
func (c Config) WithStreamRegions(cache *StreamRegions) Config {
	c.StreamRegions = cache
	return c
}

func (c *Contract) useStreamCache(cache *StreamRegions, config *Config) {
	if c.arithmeticHashTrusted && config.MetadataSeed == nil {
		c.jumpDests = cache
	}
}
