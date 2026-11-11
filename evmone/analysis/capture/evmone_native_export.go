package main

// Native host fixtures are collected by an untimed reference execution. Every
// accessed account/slot is materialized from the original call snapshot, never
// from the value observed after an earlier instruction modified it.
import (
	"bytes"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/holiman/uint256"
	"go-eth1.16.5-evm/config"
	"go-eth1.16.5-evm/core"
	"go-eth1.16.5-evm/core/state"
	"go-eth1.16.5-evm/core/tracing"
	"go-eth1.16.5-evm/core/types"
	"go-eth1.16.5-evm/core/vm"
	"math/big"
	"os"
	"path/filepath"
)

type nativeCaptureState struct {
	vm.StateDB
	pre       *state.StateDB
	accounts  map[common.Address]map[string]any
	slots     map[common.Address]map[common.Hash]bool
	transient map[common.Address]map[common.Hash]bool
}

func (s *nativeCaptureState) account(a common.Address) {
	if _, ok := s.accounts[a]; ok {
		return
	}
	s.accounts[a] = map[string]any{"exists": s.pre.Exist(a), "balance": s.pre.GetBalance(a).Hex(), "nonce": s.pre.GetNonce(a), "code": hex.EncodeToString(s.pre.GetCode(a)), "code_hash": s.pre.GetCodeHash(a).Hex(), "warm": s.pre.AddressInAccessList(a), "has_storage": s.pre.GetStorageRoot(a) != types.EmptyRootHash && s.pre.GetStorageRoot(a) != (common.Hash{}), "storage": map[string]any{}, "transient": map[string]any{}, "warm_slots": map[string]bool{}}
}
func (s *nativeCaptureState) slot(a common.Address, k common.Hash) {
	s.account(a)
	if s.slots[a] == nil {
		s.slots[a] = map[common.Hash]bool{}
	}
	if s.slots[a][k] {
		return
	}
	s.slots[a][k] = true
	cur, orig := s.pre.GetStateAndCommittedState(a, k)
	if e := s.pre.Error(); e != nil {
		panic(fmt.Sprintf("capture slot %s %s: %v", a, k, e))
	}
	_, warm := s.pre.SlotInAccessList(a, k)
	s.accounts[a]["storage"].(map[string]any)[k.Hex()] = map[string]any{"current": cur.Hex(), "original": orig.Hex(), "warm": warm}
}

// Access-list probing can precede an out-of-gas exit. Capture warmth only;
// reading the value here would request state the original execution never used.
func (s *nativeCaptureState) warmSlot(a common.Address, k common.Hash) {
	s.account(a)
	_, warm := s.pre.SlotInAccessList(a, k)
	s.accounts[a]["warm_slots"].(map[string]bool)[k.Hex()] = warm
}
func (s *nativeCaptureState) transientSlot(a common.Address, k common.Hash) {
	s.account(a)
	if s.transient[a] == nil {
		s.transient[a] = map[common.Hash]bool{}
	}
	s.transient[a][k] = true
	s.accounts[a]["transient"].(map[string]any)[k.Hex()] = s.pre.GetTransientState(a, k).Hex()
}
func (s *nativeCaptureState) GetBalance(a common.Address) *uint256.Int {
	s.account(a)
	return s.StateDB.GetBalance(a)
}
func (s *nativeCaptureState) SubBalance(a common.Address, v *uint256.Int, r tracing.BalanceChangeReason) uint256.Int {
	s.account(a)
	return s.StateDB.SubBalance(a, v, r)
}
func (s *nativeCaptureState) AddBalance(a common.Address, v *uint256.Int, r tracing.BalanceChangeReason) uint256.Int {
	s.account(a)
	return s.StateDB.AddBalance(a, v, r)
}
func (s *nativeCaptureState) GetNonce(a common.Address) uint64 {
	s.account(a)
	return s.StateDB.GetNonce(a)
}
func (s *nativeCaptureState) SetNonce(a common.Address, n uint64, r tracing.NonceChangeReason) {
	s.account(a)
	s.StateDB.SetNonce(a, n, r)
}
func (s *nativeCaptureState) GetCode(a common.Address) []byte {
	s.account(a)
	return s.StateDB.GetCode(a)
}
func (s *nativeCaptureState) GetCodeRuntime(a common.Address) ([]byte, common.Hash, any) {
	s.account(a)
	return s.StateDB.GetCodeRuntime(a)
}
func (s *nativeCaptureState) GetCodeHash(a common.Address) common.Hash {
	s.account(a)
	return s.StateDB.GetCodeHash(a)
}
func (s *nativeCaptureState) GetCodeSize(a common.Address) int {
	s.account(a)
	return s.StateDB.GetCodeSize(a)
}
func (s *nativeCaptureState) SetCode(a common.Address, b []byte, r tracing.CodeChangeReason) []byte {
	s.account(a)
	return s.StateDB.SetCode(a, b, r)
}
func (s *nativeCaptureState) Exist(a common.Address) bool { s.account(a); return s.StateDB.Exist(a) }
func (s *nativeCaptureState) Empty(a common.Address) bool { s.account(a); return s.StateDB.Empty(a) }
func (s *nativeCaptureState) GetStorageRoot(a common.Address) common.Hash {
	s.account(a)
	return s.StateDB.GetStorageRoot(a)
}
func (s *nativeCaptureState) CreateAccount(a common.Address) {
	s.account(a)
	s.StateDB.CreateAccount(a)
}
func (s *nativeCaptureState) CreateContract(a common.Address) {
	s.account(a)
	s.StateDB.CreateContract(a)
}
func (s *nativeCaptureState) SelfDestruct(a common.Address) uint256.Int {
	s.account(a)
	return s.StateDB.SelfDestruct(a)
}
func (s *nativeCaptureState) SelfDestruct6780(a common.Address) (uint256.Int, bool) {
	s.account(a)
	return s.StateDB.SelfDestruct6780(a)
}
func (s *nativeCaptureState) HasSelfDestructed(a common.Address) bool {
	s.account(a)
	return s.StateDB.HasSelfDestructed(a)
}
func (s *nativeCaptureState) AddressInAccessList(a common.Address) bool {
	s.account(a)
	return s.StateDB.AddressInAccessList(a)
}
func (s *nativeCaptureState) AddAddressToAccessList(a common.Address) {
	s.account(a)
	s.StateDB.AddAddressToAccessList(a)
}
func (s *nativeCaptureState) GetState(a common.Address, k common.Hash) common.Hash {
	s.slot(a, k)
	return s.StateDB.GetState(a, k)
}
func (s *nativeCaptureState) GetStateAndCommittedState(a common.Address, k common.Hash) (common.Hash, common.Hash) {
	s.slot(a, k)
	return s.StateDB.GetStateAndCommittedState(a, k)
}
func (s *nativeCaptureState) SetState(a common.Address, k, v common.Hash) common.Hash {
	s.slot(a, k)
	return s.StateDB.SetState(a, k, v)
}
func (s *nativeCaptureState) SlotInAccessList(a common.Address, k common.Hash) (bool, bool) {
	s.warmSlot(a, k)
	return s.StateDB.SlotInAccessList(a, k)
}
func (s *nativeCaptureState) AddSlotToAccessList(a common.Address, k common.Hash) {
	s.warmSlot(a, k)
	s.StateDB.AddSlotToAccessList(a, k)
}
func (s *nativeCaptureState) GetTransientState(a common.Address, k common.Hash) common.Hash {
	s.transientSlot(a, k)
	return s.StateDB.GetTransientState(a, k)
}
func (s *nativeCaptureState) SetTransientState(a common.Address, k, v common.Hash) {
	s.transientSlot(a, k)
	s.StateDB.SetTransientState(a, k, v)
}

func exportNativeHostFixture(block *types.Block, headers core.ChainContext, tx *types.Transaction, f *componentCall) error {
	dir := os.Getenv("EVMONE_NATIVE_EXPORT")
	if dir == "" {
		return nil
	}
	msg, err := core.TransactionToMessage(tx, types.MakeSigner(config.MainnetChainConfig, block.Number(), block.Time()), block.BaseFee())
	if err != nil {
		return err
	}
	pre := f.pre.Copy()
	st := f.pre.Copy()
	cap := &nativeCaptureState{StateDB: st, pre: pre, accounts: map[common.Address]map[string]any{}, slots: map[common.Address]map[common.Hash]bool{}, transient: map[common.Address]map[common.Hash]bool{}}
	ctx := core.NewEVMBlockContext(block.Header(), headers, nil)
	hashes := map[string]string{}
	getHash := ctx.GetHash
	ctx.GetHash = func(n uint64) common.Hash { h := getHash(n); hashes[fmt.Sprint(n)] = h.Hex(); return h }
	exportConfig := vm.Config{}
	var sites *paperSites
	if os.Getenv("PAPER_WARM") == "1" {
		sites = &paperSites{frames: map[*vm.Contract]map[uint64]bool{}, codes: map[string]map[uint64]bool{}}
		exportConfig.Tracer = &tracing.Hooks{OnOpcode: sites.opcode}
	}
	ev := vm.NewEVM(ctx, cap, config.MainnetChainConfig, exportConfig)
	ev.SetTxContext(core.NewEVMTxContext(msg))
	cap.account(f.from)
	cap.account(f.to)
	beforeLogs := len(st.Logs())
	beforeRefund := st.GetRefund()
	out, left, execErr := ev.Call(f.from, f.to, f.input, f.gas, uint256.MustFromBig(f.value))
	failure := ""
	status := 0
	if execErr != nil {
		failure = execErr.Error()
		status = 1
		if execErr == vm.ErrExecutionReverted {
			status = 2
		}
	}
	if left != f.left || failure != f.failure || !bytes.Equal(out, f.output) || st.Copy().IntermediateRoot(true) != f.root {
		return fmt.Errorf("native fixture export mismatch %s", tx.Hash())
	}
	accounts := map[string]any{}
	post := map[string]any{}
	for a, v := range cap.accounts {
		accounts[a.Hex()] = v
		slots := map[string]string{}
		for k := range cap.slots[a] {
			slots[k.Hex()] = st.GetState(a, k).Hex()
		}
		transient := map[string]string{}
		for k := range cap.transient[a] {
			transient[k.Hex()] = st.GetTransientState(a, k).Hex()
		}
		post[a.Hex()] = map[string]any{"exists": st.Exist(a), "balance": st.GetBalance(a).Hex(), "nonce": st.GetNonce(a), "code": hex.EncodeToString(st.GetCode(a)), "storage": slots, "transient": transient, "destructed": st.HasSelfDestructed(a)}
	}
	logs := []any{}
	for _, l := range componentLogs(st)[beforeLogs:] {
		topics := []string{}
		for _, t := range l.Topics {
			topics = append(topics, t.Hex())
		}
		logs = append(logs, map[string]any{"address": l.Address.Hex(), "data": hex.EncodeToString(l.Data), "topics": topics})
	}
	hexBig := func(v *big.Int) string {
		if v == nil {
			return "0x0"
		}
		return "0x" + v.Text(16)
	}
	randao := common.Hash{}
	if ctx.Random != nil {
		randao = *ctx.Random
	}
	blobs := []string{}
	for _, h := range msg.BlobHashes {
		blobs = append(blobs, h.Hex())
	}
	result := map[string]any{"hash": tx.Hash().Hex(), "fork": "Prague", "accounts": accounts, "post": post, "block_hashes": hashes, "from": f.from.Hex(), "to": f.to.Hex(), "input": hex.EncodeToString(f.input), "value": hexBig(f.value), "gas": f.gas, "output": hex.EncodeToString(out), "gas_left": left, "refund": int64(st.GetRefund()) - int64(beforeRefund), "status": status, "logs": logs, "expected_root": f.root.Hex(), "block": map[string]any{"number": block.NumberU64(), "time": ctx.Time, "gas_limit": ctx.GasLimit, "coinbase": ctx.Coinbase.Hex(), "randao": randao.Hex(), "base_fee": hexBig(ctx.BaseFee), "blob_base_fee": hexBig(ctx.BlobBaseFee), "gas_price": hexBig(msg.GasPrice), "origin": msg.From.Hex(), "blob_hashes": blobs}}
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}
	if sites != nil {
		result["executed_pcs"] = sites.export()
		result["original_instructions"] = sites.count
	}
	return writeJSON(filepath.Join(dir, tx.Hash().Hex()+".json"), result)
}
