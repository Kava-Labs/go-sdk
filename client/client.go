package client

import (
	"fmt"

	"github.com/binance-chain/bep3-deputy/util"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/tendermint/go-amino"
	"github.com/tendermint/tendermint/rpc/client"
	ctypes "github.com/tendermint/tendermint/rpc/core/types"
	tmtypes "github.com/tendermint/tendermint/types"

	"github.com/kava-labs/go-sdk/keys"
)

// KavaClient facilitates interaction with the Kava blockchain
type KavaClient struct {
	Network ChainNetwork
	HTTP    *client.HTTP
	Keybase keys.KeyManager
}

// NewKavaClient creates a new KavaClient
func NewKavaClient(cdc *amino.Codec, mnemonic string, rpcAddr string, networkType ChainNetwork) *KavaClient {
	// Set up HTTP client
	http := client.NewHTTP(rpcAddr, "/websocket")

	// TODO: import chain?
	http.Logger = util.SdkLogger

	// Set up key manager
	keyManager, err := keys.NewMnemonicKeyManager(mnemonic)
	if err != nil {
		panic(fmt.Sprintf("new key manager from mnenomic err, err=%s", err.Error()))
	}

	return &KavaClient{
		Network: networkType,
		HTTP:    http,
		Keybase: keyManager,
	}
}

// TODO: "options ...tx.Option"
func (kc *KavaClient) broadcast(m sdk.Msg, syncType SyncType) (*ctypes.ResultBroadcastTx, error) {
	signBz, err := kc.sign(m) // TODO: "options..."
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

// TODO: options ...tx.Option
func (kc *KavaClient) sign(m sdk.Msg) ([]byte, error) {
	if kc.Keybase == nil {
		return nil, fmt.Errorf("Keys are missing, must to set key")
	}

	chainID := TestChainID
	if kc.Network != ProdNetwork {
		chainID = TestChainID
	}

	signMsg := &authtypes.StdSignMsg{
		ChainID:       chainID,
		AccountNumber: 0, // TODO: -1
		Sequence:      0,
		Fee:           authtypes.NewStdFee(200000, sdk.NewCoins(sdk.NewCoin("ukava", sdk.NewInt(250000)))),
		Msgs:          []sdk.Msg{m},
		Memo:          "",
	}

	// for _, op := range options {
	// 	signMsg = op(signMsg)
	// }

	if signMsg.Sequence == 0 || signMsg.AccountNumber == 0 { // TODO: -1
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

	// TODO: remove print
	fmt.Println("Attempting to sign msg:", *signMsg)

	signedMsg, err := kc.Keybase.Sign(*signMsg)
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
