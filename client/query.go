package client

import (
	"fmt"
	"math/big"

	depCommon "github.com/binance-chain/bep3-deputy/common"
	"github.com/ethereum/go-ethereum/common"
)

// DeputyAccount interface
type DeputyAccount interface {
	GetBalances()
}

// Deputy
type Deputy struct{}

// GetBalances gets the Deputy's balance
func (d Deputy) GetBalances() {
	// TODO:
	fmt.Println("Get balances")
}

// GetBalance returns balance of swap token
func (c *Client) GetBalance() (*big.Int, error) {
	// TODO:
	return big.NewInt(1), nil
}

// GetSwap returns swap request detail
func (c *Client) GetSwap(swapId common.Hash) (*depCommon.SwapRequest, error) {
	// TODO:
	swapRequest := depCommon.SwapRequest{
		Id:                  common.Hash{},
		RandomNumberHash:    common.Hash{},
		ExpireHeight:        int64(0),
		SenderAddress:       "",
		RecipientAddress:    "",
		RecipientOtherChain: "",
		OutAmount:           big.NewInt(1),
	}

	return &swapRequest, nil
}

// GetStatus returns status of deputy account, like balance of deputy account
func (c *Client) GetStatus() (DeputyAccount, error) {
	// TODO:
	return Deputy{}, nil
}
