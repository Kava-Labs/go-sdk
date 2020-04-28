package bep3

import cmn "github.com/tendermint/tendermint/libs/common"

const (
	// QueryGetAtomicSwap command for getting info about an atomic swap
	QueryGetAtomicSwap = "swap"
	// QueryGetParams command for getting module params
	QueryGetParams = "parameters"
)

// QueryAtomicSwapByID contains the params for query 'custom/bep3/swap'
type QueryAtomicSwapByID struct {
	SwapID cmn.HexBytes `json:"swap_id" yaml:"swap_id"`
}

// NewQueryAtomicSwapByID creates a new QueryAtomicSwapByID
func NewQueryAtomicSwapByID(swapBytes cmn.HexBytes) QueryAtomicSwapByID {
	return QueryAtomicSwapByID{
		SwapID: swapBytes,
	}
}
