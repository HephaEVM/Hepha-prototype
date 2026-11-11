package database

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"go-eth1.16.5-evm/core/rawdb"
	"go-eth1.16.5-evm/core/state"
	"go-eth1.16.5-evm/ethdb"
	"go-eth1.16.5-evm/ethdb/leveldb"
	"go-eth1.16.5-evm/triedb"
)

// OpenStateDatabase uses the same state cache and file cache as BaselineProcess.
func OpenStateDatabase(path string, readonly bool) (*leveldb.Database, error) {
	return leveldb.New(path, defaultStateDBConfig.cache, defaultStateDBConfig.handles, "state_snapshot", readonly)
}

// StateDatabaseCacheConfig reports the actual budget passed to LevelDB.
// LevelDB uses half this MiB budget for its block cache. This is a configured
// limit, not the current resident size, and reporting it does not change it.
func StateDatabaseCacheConfig() (cacheMiB, handles int) {
	return defaultStateDBConfig.cache, defaultStateDBConfig.handles
}

// NewAllDBForStateFromStore reuses the state setup of NewAllDBForState while
// allowing replay to place a memory write layer above a read-only snapshot.
// The caller owns and closes store; Close only closes the trie cache here.
func NewAllDBForStateFromStore(root common.Hash, isVerkle bool, store ethdb.KeyValueStore) (*AllDBForState, error) {
	tdb := triedb.NewDatabase(rawdb.NewDatabase(store), trieDBConfig(core.DefaultConfig(), isVerkle))
	cache := state.NewDatabase(tdb, nil)
	st, err := state.New(root, cache)
	if err != nil {
		_ = tdb.Close()
		return nil, fmt.Errorf("create state db: %w", err)
	}
	return &AllDBForState{TrieDB: tdb, BlockChainStateDB: cache, StateDB: st, stateRoot: root}, nil
}
