package client

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	amino "github.com/tendermint/go-amino"

	depCommon "github.com/binance-chain/bep3-deputy/common"
	"github.com/binance-chain/bep3-deputy/store"
	"github.com/kava-labs/kava-deputy/go-sdk/keys"
	"github.com/kava-labs/kava/app"
)

// Client for interacting with kava
type Client struct {
	Cdc  *amino.Codec
	Url  string
	Keys keys.KeyManager
}

// NewClient initializes a new client
func NewClient(url string, keyManager keys.KeyManager) Client {
	cdc := app.MakeCodec()

	config := sdk.GetConfig()
	app.SetBech32AddressPrefixes(config)
	config.Seal()

	client := Client{
		Cdc:  cdc,
		Url:  url,
		Keys: keyManager,
	}

	// 	TODO: client.Start()

	return client
}

// GetHeight returns current height of chain
func (c *Client) GetHeight() (int64, error) {
	// TODO:
	return 0, nil
}

// GetFetchInterval returns fetch interval of the chain like average blocking time, it is used in observer
func (c *Client) GetFetchInterval() time.Duration {
	// TODO:
	return time.Duration(1)
}

// GetBlockAndTxs returns block info and txs included in this block
func (c *Client) GetBlockAndTxs(height int64) (*depCommon.BlockAndTxLogs, error) {
	// TODO:

	txLogs := make([]*store.TxLog, 0)

	blockAndTxLogs := depCommon.BlockAndTxLogs{
		Height:          int64(0),
		BlockHash:       "",
		ParentBlockHash: "",
		BlockTime:       int64(0),
		TxLogs:          txLogs,
	}

	return &blockAndTxLogs, nil
}

// GetSentTxStatus returns status of tx sent
func (c *Client) GetSentTxStatus(hash string) store.TxStatus {
	// TODO:
	return "tx_status"
}

// GetBalanceAlertMsg returns balance alert message if necessary, like account balance is less than amount in config
func (c *Client) GetBalanceAlertMsg() (string, error) {
	// TODO:
	return "", nil
}

// Claimable returns is swap claimable
func (c *Client) Claimable(swapId common.Hash) (bool, error) {
	// TODO:
	return false, nil
}

// Refundable returns is swap refundable
func (c *Client) Refundable(swapId common.Hash) (bool, error) {
	// TODO:
	return false, nil
}

// HasSwap returns does swap exist
func (c *Client) HasSwap(swapId common.Hash) (bool, error) {
	// TODO:
	return false, nil
}
