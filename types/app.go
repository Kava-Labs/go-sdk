package types

import (
	"github.com/kava-labs/cosmos-sdk/codec"
	sdk "github.com/kava-labs/cosmos-sdk/types"
	authtypes "github.com/kava-labs/cosmos-sdk/x/auth/types"
	"github.com/kava-labs/cosmos-sdk/x/auth/vesting"

	// "github.com/kava-labs/kava/app"
	"github.com/tendermint/go-amino"
)

const (
	Bech32MainPrefix = "kava"
	Bip44CoinType    = 459
)

// SetBech32AddressPrefixes sets the global prefix to be used when serializing addresses to bech32 strings.
func SetBech32AddressPrefixes(config *sdk.Config) {
	config.SetBech32PrefixForAccount(Bech32MainPrefix, Bech32MainPrefix+sdk.PrefixPublic)
	config.SetBech32PrefixForValidator(Bech32MainPrefix+sdk.PrefixValidator+sdk.PrefixOperator, Bech32MainPrefix+sdk.PrefixValidator+sdk.PrefixOperator+sdk.PrefixPublic)
	config.SetBech32PrefixForConsensusNode(Bech32MainPrefix+sdk.PrefixValidator+sdk.PrefixConsensus, Bech32MainPrefix+sdk.PrefixValidator+sdk.PrefixConsensus+sdk.PrefixPublic)
}

// --------------------------- Codec ---------------------------

// Cdc is the global app codec
var Cdc *codec.Codec

// MakeCodec registers Kava app and locally defined txs/msgs
func MakeCodec() *codec.Codec {
	Cdc = codec.New()

	// Register Kava app
	// app.ModuleBasics.RegisterCodec(Cdc)
	authtypes.RegisterCodec(Cdc)
	vesting.RegisterCodec(Cdc)
	sdk.RegisterCodec(Cdc)
	codec.RegisterCrypto(Cdc)
	codec.RegisterEvidences(Cdc)

	// Register local interfaces, types
	RegisterCodec(Cdc)
	return Cdc
}

// RegisterCodec registers local messages
func RegisterCodec(cdc *amino.Codec) {
	cdc.RegisterInterface((*Tx)(nil), nil)
}

// Tx is a wrapper around sdk.Tx
type Tx sdk.Tx
