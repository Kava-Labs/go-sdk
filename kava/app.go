package kava

import (
	"github.com/kava-labs/cosmos-sdk/codec"
	sdk "github.com/kava-labs/cosmos-sdk/types"
	authtypes "github.com/kava-labs/cosmos-sdk/x/auth"
	"github.com/kava-labs/cosmos-sdk/x/auth/vesting"
	govtypes "github.com/kava-labs/cosmos-sdk/x/gov"
	paramstypes "github.com/kava-labs/cosmos-sdk/x/params"
	supplytypes "github.com/kava-labs/cosmos-sdk/x/supply"

	"github.com/kava-labs/go-sdk/kava/msgs"
)

const (
	Bech32MainPrefix = "kava"
	Bip44CoinType    = 459
)

// --------------------------- Codec ---------------------------

// Cdc is the global app codec
var Cdc *codec.Codec

// MakeCodec registers Kava app and locally defined txs/msgs
func MakeCodec() *codec.Codec {
	Cdc = codec.New()

	// Register custom Kava module msgs
	msgs.RegisterCodec(Cdc)

	// Register cosmos-sdk module msgs
	authtypes.RegisterCodec(Cdc)
	supplytypes.RegisterCodec(Cdc)
	paramstypes.RegisterCodec(Cdc)
	govtypes.RegisterCodec(Cdc)
	vesting.RegisterCodec(Cdc)

	// Register codec, crypto, and evidences
	sdk.RegisterCodec(Cdc)
	codec.RegisterCrypto(Cdc)
	codec.RegisterEvidences(Cdc)

	return Cdc
}

// --------------------------- Config ---------------------------

// SetBech32AddressPrefixes sets the global prefix to be used when serializing addresses to bech32 strings.
func SetBech32AddressPrefixes(config *sdk.Config) {
	config.SetBech32PrefixForAccount(Bech32MainPrefix, Bech32MainPrefix+sdk.PrefixPublic)
	config.SetBech32PrefixForValidator(Bech32MainPrefix+sdk.PrefixValidator+sdk.PrefixOperator, Bech32MainPrefix+sdk.PrefixValidator+sdk.PrefixOperator+sdk.PrefixPublic)
	config.SetBech32PrefixForConsensusNode(Bech32MainPrefix+sdk.PrefixValidator+sdk.PrefixConsensus, Bech32MainPrefix+sdk.PrefixValidator+sdk.PrefixConsensus+sdk.PrefixPublic)
}
