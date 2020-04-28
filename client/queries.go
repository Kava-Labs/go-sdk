package client

import (
	"errors"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	cmn "github.com/tendermint/tendermint/libs/common"

	"github.com/kava-labs/go-sdk/types/bep3"
)

// GetSwapByID gets an atomic swap on Kava by ID
func (kc *KavaClient) GetSwapByID(swapID cmn.HexBytes) (swap bep3.AtomicSwap, err error) {
	params := bep3.NewQueryAtomicSwapByID(swapID)
	bz, err := kc.Keybase.GetCodec().MarshalJSON(params)
	if err != nil {
		return bep3.AtomicSwap{}, err
	}

	path := "custom/bep3/swap"

	result, err := kc.ABCIQuery(path, bz)
	if err != nil {
		return bep3.AtomicSwap{}, err
	}

	err = kc.Keybase.GetCodec().UnmarshalJSON(result, &swap)
	if err != nil {
		return bep3.AtomicSwap{}, err
	}
	return swap, nil
}

// GetAccount gets the account associated with an address on Kava
func (kc *KavaClient) GetAccount(addr sdk.AccAddress) (acc authtypes.BaseAccount, err error) {
	params := authtypes.NewQueryAccountParams(addr)
	bz, err := kc.Keybase.GetCodec().MarshalJSON(params)
	if err != nil {
		return authtypes.BaseAccount{}, err
	}

	path := fmt.Sprintf("custom/acc/account/%s", addr.String())

	result, err := kc.ABCIQuery(path, bz)
	if err != nil {
		return authtypes.BaseAccount{}, err
	}

	// TODO: UnmarshalJSON isn't working
	fmt.Println("result:", result)

	err = kc.Keybase.GetCodec().UnmarshalJSON(result, &acc)
	if err != nil {
		return authtypes.BaseAccount{}, err
	}

	return acc, err
}

// ABCIQuery sends a query to Kava
func (kc *KavaClient) ABCIQuery(path string, data cmn.HexBytes) ([]byte, error) {
	if err := ValidateABCIQuery(path, data); err != nil {
		return []byte{}, err
	}

	result, err := kc.HTTP.ABCIQuery(path, data)
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
