package app

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/kava-labs/go-sdk/types/tx"
	"github.com/tendermint/go-amino"
)

//TODO:  ctypes "github.com/tendermint/tendermint/rpc/core/types"
// TODO: ctypes.RegisterAmino(cdc)

const (
	Bech32MainPrefix = "kava"
	Bip44CoinType    = 459
)

// MakeCodec registers msgs and interfaces.
// Must register any module's msgs before go-sdk can interact with them.
func MakeCodec() *amino.Codec {
	var cdc = codec.New()
	sdk.RegisterCodec(cdc)
	codec.RegisterCrypto(cdc)
	authtypes.RegisterCodec(cdc)
	tx.RegisterCodec(cdc)
	return cdc
}

// SetBech32AddressPrefixes sets the global prefix to be used when serializing addresses to bech32 strings.
func SetBech32AddressPrefixes(config *sdk.Config) {
	config.SetBech32PrefixForAccount(Bech32MainPrefix, Bech32MainPrefix+sdk.PrefixPublic)
	config.SetBech32PrefixForValidator(Bech32MainPrefix+sdk.PrefixValidator+sdk.PrefixOperator, Bech32MainPrefix+sdk.PrefixValidator+sdk.PrefixOperator+sdk.PrefixPublic)
	config.SetBech32PrefixForConsensusNode(Bech32MainPrefix+sdk.PrefixValidator+sdk.PrefixConsensus, Bech32MainPrefix+sdk.PrefixValidator+sdk.PrefixConsensus+sdk.PrefixPublic)
}

// SetBip44CoinType sets the global coin type to be used in hierarchical deterministic wallets.
func SetBip44CoinType(config *sdk.Config) {
	config.SetCoinType(Bip44CoinType)
}
