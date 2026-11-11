//go:build hepha_export

package main

// Untimed, chronological root-frame witnesses for the upstream C++ state host.
// Capture starts AFTER the outer call's value transfer / creation setup. The
// native driver therefore invokes execute directly, and still executes nested
// calls normally. Each original block is separately validated by the replay.
import (
	"compress/gzip"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"go-eth1.16.5-evm/core/state"
	"go-eth1.16.5-evm/core/types"
	"go-eth1.16.5-evm/core/vm"
	"math/big"
	"os"
	"time"
)

type nativeStreamCapture struct {
	file     *os.File
	packed   *gzip.Writer
	enc      *json.Encoder
	tx       *types.Transaction
	index    int
	block    uint64
	nested   bool
	recorded bool
	codes    map[common.Hash]bool
}

func newNativeStreamCapture(path string) (*nativeStreamCapture, error) {
	f, e := os.OpenFile(path, os.O_CREATE|os.O_EXCL|os.O_WRONLY, 0644)
	if e != nil {
		return nil, e
	}
	z, e := gzip.NewWriterLevel(f, gzip.BestSpeed)
	if e != nil {
		f.Close()
		return nil, e
	}
	return &nativeStreamCapture{file: f, packed: z, enc: json.NewEncoder(z), codes: map[common.Hash]bool{}}, nil
}
func (c *nativeStreamCapture) close() {
	if e := c.packed.Close(); e != nil {
		panic(e)
	}
	if e := c.file.Close(); e != nil {
		panic(e)
	}
}
func (c *nativeStreamCapture) write(row map[string]any) {
	newCodes := map[string]string{}
	ref := func(code string) string {
		b, err := hex.DecodeString(code)
		if err != nil {
			panic(err)
		}
		h := crypto.Keccak256Hash(b)
		if !c.codes[h] {
			c.codes[h] = true
			newCodes[h.Hex()] = code
		}
		return h.Hex()
	}
	if code, ok := row["root_code"].(string); ok {
		row["code_ref"] = ref(code)
		delete(row, "root_code")
	}
	for _, part := range []string{"accounts", "post"} {
		if ac, ok := row[part].(map[string]any); ok {
			for _, v := range ac {
				x := v.(map[string]any)
				x["code_ref"] = ref(x["code"].(string))
				delete(x, "code")
			}
		}
	}
	row["new_codes"] = newCodes
	row["hash"] = c.tx.Hash().Hex()
	row["tx_index"] = c.index
	row["block_number"] = c.block
	row["outer_to"] = "CREATE"
	if c.tx.To() != nil {
		row["outer_to"] = c.tx.To().Hex()
	}
	if e := c.enc.Encode(row); e != nil {
		panic(e)
	}
}
func (c *nativeStreamCapture) RunFrame(e *vm.EVM, contract *vm.Contract, input []byte, readOnly bool) (out []byte, err error, lat time.Duration) {
	if c.nested || c.tx == nil {
		return e.RunNative(contract, input, readOnly)
	}
	c.nested = true
	defer func() { c.nested = false }()
	st, ok := e.StateDB.(*state.StateDB)
	if !ok {
		panic("capture requires plain StateDB")
	}
	pre := st.Copy()
	cap := &nativeCaptureState{StateDB: st, pre: pre, accounts: map[common.Address]map[string]any{}, slots: map[common.Address]map[common.Hash]bool{}, transient: map[common.Address]map[common.Hash]bool{}}
	from, to := contract.Caller(), contract.Address()
	cap.account(from)
	cap.account(to)
	beforeLogs, beforeRefund := len(st.Logs()), st.GetRefund()
	gas := contract.Gas
	hashes := map[string]string{}
	getHash := e.Context.GetHash
	e.Context.GetHash = func(n uint64) common.Hash { h := getHash(n); hashes[fmt.Sprint(n)] = h.Hex(); return h }
	e.StateDB = cap
	out, err, lat = e.RunNative(contract, input, readOnly)
	e.StateDB = st
	e.Context.GetHash = getHash
	if x := pre.Error(); x != nil {
		panic(fmt.Sprintf("capture pre tx=%s: %v", c.tx.Hash(), x))
	}
	if x := st.Error(); x != nil {
		panic(fmt.Sprintf("capture execution tx=%s: %v", c.tx.Hash(), x))
	}
	// Failed roots discard all frame writes. Clients may mutate a storage slot
	// before or after charging its gas; those transient differences are not
	// observable once the enclosing frame rolls back.
	postState := st
	if err != nil {
		postState = pre
	}
	post := map[string]any{}
	accounts := map[string]any{}
	for a, v := range cap.accounts {
		accounts[a.Hex()] = v
		slots := map[string]string{}
		for k := range cap.slots[a] {
			slots[k.Hex()] = postState.GetState(a, k).Hex()
			if x := postState.Error(); x != nil {
				panic(fmt.Sprintf("capture post tx=%s account=%s slot=%s: %v", c.tx.Hash(), a, k, x))
			}
		}
		transient := map[string]string{}
		for k := range cap.transient[a] {
			transient[k.Hex()] = postState.GetTransientState(a, k).Hex()
		}
		post[a.Hex()] = map[string]any{"exists": postState.Exist(a), "balance": postState.GetBalance(a).Hex(), "nonce": postState.GetNonce(a), "code": hex.EncodeToString(postState.GetCode(a)), "storage": slots, "transient": transient, "destructed": postState.HasSelfDestructed(a)}
	}
	logs := []any{}
	for _, l := range componentLogs(st)[beforeLogs:] {
		ts := []string{}
		for _, t := range l.Topics {
			ts = append(ts, t.Hex())
		}
		logs = append(logs, map[string]any{"address": l.Address.Hex(), "data": hex.EncodeToString(l.Data), "topics": ts})
	}
	if err != nil {
		logs = []any{}
	}
	hexBig := func(v *big.Int) string {
		if v == nil {
			return "0x0"
		}
		return "0x" + v.Text(16)
	}
	randao := common.Hash{}
	if e.Context.Random != nil {
		randao = *e.Context.Random
	}
	blobs := []string{}
	for _, h := range e.BlobHashes {
		blobs = append(blobs, h.Hex())
	}
	status := 0
	left := contract.Gas
	refund := int64(st.GetRefund()) - int64(beforeRefund)
	if err != nil {
		status = 1
		left = 0
		refund = 0
		if err == vm.ErrExecutionReverted {
			status = 2
			left = contract.Gas
		}
	}
	c.write(map[string]any{"frame": true, "root_code": hex.EncodeToString(contract.Code), "creation": contract.IsDeployment, "read_only": readOnly, "accounts": accounts, "post": post, "block_hashes": hashes, "from": from.Hex(), "to": to.Hex(), "input": hex.EncodeToString(input), "value": contract.Value().Hex(), "gas": gas, "gas_left": left, "refund": refund, "status": status, "output": hex.EncodeToString(out), "logs": logs, "block": map[string]any{"number": e.Context.BlockNumber.Uint64(), "time": e.Context.Time, "gas_limit": e.Context.GasLimit, "coinbase": e.Context.Coinbase.Hex(), "randao": randao.Hex(), "base_fee": hexBig(e.Context.BaseFee), "blob_base_fee": hexBig(e.Context.BlobBaseFee), "gas_price": hexBig(e.GasPrice), "origin": e.Origin.Hex(), "blob_hashes": blobs}})
	c.recorded = true
	return
}
