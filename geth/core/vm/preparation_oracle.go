package vm

import "go-eth1.16.5-evm/regionexec/prepare"

func ExperimentOracleEnd() error {
	prepare.ExperimentOracleEnd()
	if m := getDemandManager(); m != nil {
		return m.ExperimentOracleMisses()
	}
	return nil
}
