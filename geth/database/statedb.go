package database

import (
	"fmt"
	"math/big"
	"path/filepath"

	"github.com/ethereum/go-ethereum/common"
	"go-eth1.16.5-evm/core/state"
	"go-eth1.16.5-evm/ethdb/leveldb"
	"go-eth1.16.5-evm/triedb"
)

type AllDBForState struct {
	DiskDB            *leveldb.Database
	TrieDB            *triedb.Database
	BlockChainStateDB *state.CachingDB
	StateDB           *state.StateDB
	stateRoot         common.Hash
}

func NewAllDBForState(blockNumber *big.Int, stateRoot common.Hash, isVerkle, readonly bool) (*AllDBForState, error) {
	diskPath := filepath.Join(defaultStateDBConfig.path, "snapshot_"+blockNumber.String())

	levelDB, err := OpenStateDatabase(diskPath, readonly)
	if err != nil {
		return nil, fmt.Errorf("open leveldb error: %w", err)
	}

	all, err := NewAllDBForStateFromStore(stateRoot, isVerkle, levelDB)
	if err != nil {
		_ = levelDB.Close()
		return nil, err
	}
	all.DiskDB = levelDB
	return all, nil
}

func (a *AllDBForState) UpdateStateDB(stateRoot common.Hash) error {
	if stateRoot == a.stateRoot {
		return nil
	}
	statedb, err := state.New(stateRoot, a.BlockChainStateDB)
	if err != nil {
		return fmt.Errorf("create state db: %w", err)
	}
	a.StateDB = statedb
	a.stateRoot = stateRoot
	return nil
}

func (a *AllDBForState) Close() {
	_ = a.TrieDB.Close()
	if a.DiskDB != nil {
		_ = a.DiskDB.Close()
	}
}

func (a *AllDBForState) Copy() *state.StateDB {
	return a.StateDB.Copy()
}

