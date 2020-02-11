package client

// bep3 types
import (
	"fmt"
	"math/big"

	depCommon "github.com/binance-chain/bep3-deputy/common"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	bep3 "github.com/kava-labs/kava/x/bep3/types"
)

// HTLT sends htlt tx
func (c *Client) HTLT(randomNumberHash common.Hash, timestamp int64, heightSpan int64, recipientAddr string,
	otherChainSenderAddr string, otherChainRecipientAddr string, outAmount *big.Int) (string, *depCommon.Error) {

	msgCreateHTLT := bep3.NewHTLTMsg(
		sdk.AccAddress(recipientAddr), // TODO: Deputy address
		sdk.AccAddress(recipientAddr),
		otherChainRecipientAddr,
		otherChainSenderAddr,
		bep3.SwapBytes(randomNumberHash.Bytes()),
		timestamp,
		sdk.NewCoins(sdk.NewInt64Coin("kava", outAmount.Int64())),
		fmt.Sprintf("kava", outAmount),
		heightSpan,
		true,
	)

	// TODO:
	_ = msgCreateHTLT

	return "", depCommon.NewError(nil, false)
}

// Claim sends claim tx
func (c *Client) Claim(swapId common.Hash, randomNumber common.Hash) (string, *depCommon.Error) {
	// TODO:
	return "", depCommon.NewError(nil, false)
}

// Refund sends refund tx
func (c *Client) Refund(swapId common.Hash) (string, *depCommon.Error) {
	// TODO:
	return "", depCommon.NewError(nil, false)
}

func (c *Client) PostTx() {
	//TODO:
}
