package types

import (
	tmbytes "github.com/kava-labs/tendermint/libs/bytes"
)

// QueryAtomicSwapByID contains the params for query 'custom/bep3/swap'
type QueryAtomicSwapByID struct {
	SwapID tmbytes.HexBytes `json:"swap_id" yaml:"swap_id"`
}

// NewQueryAtomicSwapByID creates a new QueryAtomicSwapByID
func NewQueryAtomicSwapByID(swapBytes tmbytes.HexBytes) QueryAtomicSwapByID {
	return QueryAtomicSwapByID{
		SwapID: swapBytes,
	}
}
