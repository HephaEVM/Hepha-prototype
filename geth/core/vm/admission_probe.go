package vm

import (
	"encoding/binary"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"sync/atomic"
)

// Admission hints only: no execution plan is ever keyed by a fingerprint.
// A collision can delay or hasten preparation; the existing full key, bytecode
// copy, proof and publication remain the only authority for execution.
type preparationProbation struct {
	threshold uint64
	slots     [1024]atomic.Uint64
}

func (p *preparationProbation) observe(hash common.Hash) bool {
	word := binary.LittleEndian.Uint64(hash[:8])
	tag := word &^ uint64(65535)
	slot := &p.slots[word&1023]
	for {
		old := slot.Load()
		next := tag | 1
		if old&^uint64(65535) == tag {
			if old&65535 >= p.threshold {
				return true
			}
			next = old + 1
		}
		if slot.CompareAndSwap(old, next) {
			return next&65535 >= p.threshold
		}
	}
}

// Experiment-only constructor option. Below-threshold observations retain no
// code or request and are not admitted work; Wait drains admitted requests.
// Configure before any calls. Zero/one preserves existing eager admission.
func (p *BackgroundPreparation) ExperimentAdmissionProbation(threshold uint64) error {
	if threshold > 65535 {
		return fmt.Errorf("probation threshold exceeds counter range")
	}
	if threshold > 1 {
		p.probation = &preparationProbation{threshold: threshold}
	}
	return nil
}
