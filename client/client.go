package client

import (
	"fmt"
	"os"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/tendermint/tendermint/libs/log"
	rpcclient "github.com/tendermint/tendermint/rpc/client/http"
	ctypes "github.com/tendermint/tendermint/rpc/core/types"
	tmtypes "github.com/tendermint/tendermint/types"

	"github.com/kava-labs/go-sdk/keys"
)

// KavaClient facilitates interaction with the Kava blockchain
type KavaClient struct {
	HTTP    *rpcclient.HTTP
	Keybase keys.KeyManager
	Cdc     *codec.Codec
}

// NewKavaClient creates a new KavaClient
func NewKavaClient(cdc *codec.Codec, mnemonic string, coinID uint32, rpcAddr string) *KavaClient {
	// Set up HTTP client
	http, err := rpcclient.New(rpcAddr, "/websocket")
	if err != nil {
		panic(err)
	}
	http.Logger = log.NewTMLogger(log.NewSyncWriter(os.Stdout))

	// Set up key manager
	keyManager, err := keys.NewMnemonicKeyManager(mnemonic, coinID)
	if err != nil {
		panic(fmt.Sprintf("new key manager from mnenomic err, err=%s", err.Error()))
	}

	return &KavaClient{
		HTTP:    http,
		Keybase: keyManager,
		Cdc:     cdc,
	}
}

// Broadcast sends a message to the Kava blockchain as a transaction.
// This pays no transaction fees.
func (kc *KavaClient) Broadcast(m sdk.Msg, syncType SyncType) (*ctypes.ResultBroadcastTx, error) {
	fee := authtypes.NewStdFee(250000, nil)
	return kc.BroadcastWithFee(m, fee, syncType)
}

// BroadcastWithFee sends a message to the Kava blockchain as a transaction, paying the specified transaction fee.
func (kc *KavaClient) BroadcastWithFee(m sdk.Msg, fee authtypes.StdFee, syncType SyncType) (*ctypes.ResultBroadcastTx, error) {
	signBz, err := kc.sign(m, fee)
	if err != nil {
		return nil, err
	}
	switch syncType {
	case Async:
		return kc.BroadcastTxAsync(signBz)
	case Sync:
		return kc.BroadcastTxSync(signBz)
	case Commit:
		commitRes, err := kc.BroadcastTxCommit(signBz)
		if err != nil {
			return nil, err
		}
		if commitRes.CheckTx.IsErr() {
			return &ctypes.ResultBroadcastTx{
				Code: commitRes.CheckTx.Code,
				Log:  commitRes.CheckTx.Log,
				Hash: commitRes.Hash,
				Data: commitRes.CheckTx.Data,
			}, nil
		}
		return &ctypes.ResultBroadcastTx{
			Code: commitRes.DeliverTx.Code,
			Log:  commitRes.DeliverTx.Log,
			Hash: commitRes.Hash,
			Data: commitRes.DeliverTx.Data,
		}, nil
	default:
		return nil, fmt.Errorf("unknown synctype")
	}
}

func (kc *KavaClient) sign(m sdk.Msg, fee authtypes.StdFee) ([]byte, error) {
	if kc.Keybase == nil {
		return nil, fmt.Errorf("Keys are missing, must to set key")
	}

	chainID, err := kc.GetChainID()
	if err != nil {
		return nil, fmt.Errorf("could not fetch chain id: %w", err)
	}

	signMsg := &authtypes.StdSignMsg{
		ChainID:       chainID,
		AccountNumber: 0,
		Sequence:      0,
		Fee:           fee,
		Msgs:          []sdk.Msg{m},
		Memo:          "",
	}

	if signMsg.Sequence == 0 || signMsg.AccountNumber == 0 {
		fromAddr := kc.Keybase.GetAddr()
		acc, err := kc.GetAccount(fromAddr)
		if err != nil {
			return nil, err
		}

		if acc.Address.Empty() {
			return nil, fmt.Errorf("the signer account does not exist on kava")
		}

		signMsg.Sequence = acc.Sequence
		signMsg.AccountNumber = acc.AccountNumber
	}

	for _, m := range signMsg.Msgs {
		if err := m.ValidateBasic(); err != nil {
			return nil, err
		}
	}

	signedMsg, err := kc.Keybase.Sign(*signMsg, kc.Cdc)
	if err != nil {
		return nil, err
	}

	return signedMsg, nil
}

// BroadcastTxCommit sends a transaction using commit
func (kc *KavaClient) BroadcastTxCommit(tx tmtypes.Tx) (*ctypes.ResultBroadcastTxCommit, error) {
	if err := ValidateTx(tx); err != nil {
		return nil, err
	}
	return kc.HTTP.BroadcastTxCommit(tx)
}

// BroadcastTxAsync sends a transaction using async
func (kc *KavaClient) BroadcastTxAsync(tx tmtypes.Tx) (*ctypes.ResultBroadcastTx, error) {
	if err := ValidateTx(tx); err != nil {
		return nil, err
	}
	return kc.HTTP.BroadcastTxAsync(tx)
}

// BroadcastTxSync sends a transaction using sync
func (kc *KavaClient) BroadcastTxSync(tx tmtypes.Tx) (*ctypes.ResultBroadcastTx, error) {
	if err := ValidateTx(tx); err != nil {
		return nil, err
	}
	return kc.HTTP.BroadcastTxSync(tx)
}
