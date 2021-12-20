package client

import (
	"context"
	"errors"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	tmbytes "github.com/tendermint/tendermint/libs/bytes"

	bep3 "github.com/kava-labs/kava/x/bep3/types"
)

// GetSwapByID gets an atomic swap on Kava by ID
func (kc *KavaClient) GetSwapByID(ctx context.Context, swapID tmbytes.HexBytes) (swap bep3.AtomicSwap, err error) {
	params := bep3.NewQueryAtomicSwapByID(swapID)
	bz, err := kc.Cdc.MarshalJSON(params)
	if err != nil {
		return bep3.AtomicSwap{}, err
	}

	path := "custom/bep3/swap"

	result, err := kc.ABCIQuery(ctx, path, bz)
	if err != nil {
		return bep3.AtomicSwap{}, err
	}

	err = kc.Cdc.UnmarshalJSON(result, &swap)
	if err != nil {
		return bep3.AtomicSwap{}, err
	}
	return swap, nil
}

// GetAccount gets the account associated with an address on Kava
func (kc *KavaClient) GetAccount(ctx context.Context, addr sdk.AccAddress) (acc authtypes.BaseAccount, err error) {
	params := authtypes.QueryAccountRequest{Address: addr.String()}
	bz, err := kc.Cdc.MarshalJSON(params)
	if err != nil {
		return authtypes.BaseAccount{}, err
	}

	path := fmt.Sprintf("custom/auth/account/%s", addr.String())

	result, err := kc.ABCIQuery(ctx, path, bz)
	if err != nil {
		return authtypes.BaseAccount{}, err
	}

	err = kc.Cdc.UnmarshalJSON(result, &acc)
	if err != nil {
		return authtypes.BaseAccount{}, err
	}

	return acc, err
}

func (kc *KavaClient) GetChainID(ctx context.Context) (string, error) {
	result, err := kc.HTTP.Status(ctx)
	if err != nil {
		return "", err
	}
	return result.NodeInfo.Network, nil
}

// ABCIQuery sends a query to Kava
func (kc *KavaClient) ABCIQuery(ctx context.Context, path string, data tmbytes.HexBytes) ([]byte, error) {
	if err := ValidateABCIQuery(path, data); err != nil {
		return []byte{}, err
	}

	result, err := kc.HTTP.ABCIQuery(ctx, path, data)
	if err != nil {
		return []byte{}, err
	}

	resp := result.Response
	if !resp.IsOK() {
		return []byte{}, errors.New(resp.Log)
	}

	value := result.Response.GetValue()
	if len(value) == 0 {
		return []byte{}, nil
	}

	return value, nil
}
