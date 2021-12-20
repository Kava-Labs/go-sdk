package client

import (
	"context"
	"fmt"
	"os"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/legacy/legacytx"
	"github.com/tendermint/tendermint/libs/log"
	rpcclient "github.com/tendermint/tendermint/rpc/client/http"

	"github.com/kava-labs/go-sdk/keys"
	"github.com/kava-labs/kava/app"
)

// KavaClient facilitates interaction with the Kava blockchain
type KavaClient struct {
	HTTP    *rpcclient.HTTP
	Keybase keys.KeyManager
	Cdc     *codec.LegacyAmino
}

// NewKavaClient creates a new KavaClient
func NewKavaClient(cdc *codec.LegacyAmino, mnemonic string, coinID uint32, rpcAddr string) *KavaClient {
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
func (kc *KavaClient) Broadcast(clientCtx client.Context, m sdk.Msg, syncType SyncType) (res *sdk.TxResponse, err error) {
	fee := legacytx.NewStdFee(250000, nil)
	return kc.BroadcastWithFee(clientCtx, m, fee, syncType)
}

// BroadcastWithFee sends a message to the Kava blockchain as a transaction, paying the specified transaction fee.
func (kc *KavaClient) BroadcastWithFee(clientCtx client.Context, m sdk.Msg, fee legacytx.StdFee, syncType SyncType) (res *sdk.TxResponse, err error) {
	signBz, err := kc.sign(m, fee)
	if err != nil {
		return nil, err
	}

	// build a legacy transaction
	sigs := []legacytx.StdSignature{legacytx.NewStdSignature(kc.Keybase.GetKeyRing().GetPubKey(), signBz)}
	stdTx := legacytx.NewStdTx([]sdk.Msg{m}, fee, sigs, "")
	stdTx.TimeoutHeight = 100000

	var mode string
	switch syncType {
	case Async:
		mode = "async"
	case Sync:
		mode = "sync"
	case Commit:
		mode = "commit"
	default:
		return nil, fmt.Errorf("unknown synctype")
	}
	legacyTx := app.LegacyTxBroadcastRequest{
		Tx:   stdTx,
		Mode: mode,
	}

	// Build the tx from the legacy tx object
	builder := clientCtx.TxConfig.NewTxBuilder()
	builder.SetFeeAmount(legacyTx.Tx.GetFee())
	builder.SetGasLimit(legacyTx.Tx.GetGas())
	builder.SetMemo(legacyTx.Tx.GetMemo())
	builder.SetTimeoutHeight(legacyTx.Tx.GetTimeoutHeight())

	signatures, err := legacyTx.Tx.GetSignaturesV2()
	if err != nil {
		return nil, err
	}

	// TODO: could use account retreiver here instead i.e.
	//		 clientCtx.AccountRetriever.GetAccountNumberSequence(clientCtx, addr)
	for i, sig := range signatures {
		addr := sdk.AccAddress(sig.PubKey.Address())
		acc, err := kc.GetAccount(context.Background(), addr)
		if err != nil {
			return nil, err
		}
		signatures[i].Sequence = acc.GetSequence()
	}

	err = builder.SetSignatures(signatures...)
	if err != nil {
		return nil, err
	}

	txBytes, err := clientCtx.TxConfig.TxEncoder()(builder.GetTx())
	if err != nil {
		return nil, err
	}

	clientCtx = clientCtx.WithBroadcastMode(legacyTx.Mode)
	res, err = clientCtx.BroadcastTx(txBytes)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (kc *KavaClient) sign(m sdk.Msg, fee legacytx.StdFee) ([]byte, error) {
	if kc.Keybase == nil {
		return nil, fmt.Errorf("Keys are missing, must to set key")
	}
	chainID, err := kc.GetChainID(context.Background())
	if err != nil {
		return nil, fmt.Errorf("could not fetch chain id: %w", err)
	}

	signMsg := &legacytx.StdSignMsg{
		ChainID:       chainID,
		AccountNumber: 0,
		Sequence:      0,
		Fee:           fee,
		Msgs:          []sdk.Msg{m},
		Memo:          "",
	}

	if signMsg.Sequence == 0 || signMsg.AccountNumber == 0 {
		fromAddr := kc.Keybase.GetKeyRing().GetAddress()
		acc, err := kc.GetAccount(context.Background(), fromAddr)
		if err != nil {
			return nil, err
		}

		if acc.GetAddress().Empty() {
			return nil, fmt.Errorf("the signer account does not exist on kava")
		}

		signMsg.Sequence = acc.GetSequence()
		signMsg.AccountNumber = acc.GetAccountNumber()
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
