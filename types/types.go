package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// NewInt64Coin is a wrapper around sdk.NewInt64Coin
func NewInt64Coin(denom string, amount int64) sdk.Coin {
	return sdk.NewInt64Coin(denom, amount)
}

// NewCoins is a wrapper around sdk.NewCoins
func NewCoins(coins ...sdk.Coin) sdk.Coins {
	return sdk.NewCoins(coins...)
}

// AccAddressFromBech32 is a wrapper around sdk.AccAddressFromBech32
func AccAddressFromBech32(address string) (addr sdk.AccAddress, err error) {
	accAddress, err := sdk.AccAddressFromBech32(address)
	if err != nil {
		return nil, err
	}
	return accAddress, nil
}

// Tx is a wrapper around sdk.Tx
type Tx sdk.Tx

// GetConfig is a wrapper around sdk.GetConfig
func GetConfig() *sdk.Config {
	return sdk.GetConfig()
}
