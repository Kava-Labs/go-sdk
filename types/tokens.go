package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Token definition
type Token struct {
	Name        string         `json:"name"`
	Symbol      string         `json:"symbol"`
	OrigSymbol  string         `json:"original_symbol"`
	TotalSupply Fixed8         `json:"total_supply"`
	Owner       sdk.AccAddress `json:"owner"`
	Mintable    bool           `json:"mintable"`
}

type TokenBalance struct {
	Symbol string `json:"symbol"`
	Free   Fixed8 `json:"free"`
	Locked Fixed8 `json:"locked"`
	Frozen Fixed8 `json:"frozen"`
}
