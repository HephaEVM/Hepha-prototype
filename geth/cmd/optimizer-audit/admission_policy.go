package main

import (
	"fmt"
	"go-eth1.16.5-evm/core/vm"
	"os"
	"strconv"
)

func experimentAdmissionPreparation(e *vm.ArithmeticEngine, a *vm.ABIScalarDiagnostic) (*vm.BackgroundPreparation, error) {
	read := func(k string, def int) (int, error) {
		s := os.Getenv(k)
		if s == "" {
			return def, nil
		}
		return strconv.Atoi(s)
	}
	threshold, err := read("COLD_ADMISSION_THRESHOLD", 0)
	if err != nil {
		return nil, err
	}
	duty, err := read("COLD_PREPARATION_DUTY", 100)
	if err != nil {
		return nil, err
	}
	if threshold < 0 || threshold > 65535 || duty < 1 || duty > 100 {
		return nil, fmt.Errorf("invalid experimental admission/duty")
	}
	minimum, err := read("COLD_PREPARATION_MIN_HITS", 4)
	if err != nil || minimum < 1 || minimum > 64 {
		return nil, fmt.Errorf("invalid code preparation threshold")
	}
	entryMinimum, err := read("COLD_ENTRY_MIN_HITS", 2)
	if err != nil || entryMinimum < 1 || entryMinimum > 64 {
		return nil, fmt.Errorf("invalid entry preparation threshold")
	}
	p, err := vm.NewBackgroundPreparationWithLimits(e, a, hotsetCapacity(), vm.BackgroundPreparationLimits{uint64(minimum), uint64(entryMinimum), duty})
	if err != nil {
		return nil, err
	}
	if err = p.ExperimentAdmissionProbation(uint64(threshold)); err != nil {
		p.Close()
		return nil, err
	}
	return p, nil
}

func hotsetCapacity() int {
	return vm.ExperimentPreparationCapacity()
}
