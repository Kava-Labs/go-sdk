package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/vesting"
	"github.com/kava-labs/kava/app"
	"github.com/tendermint/go-amino"
)

// --------------------------- Codec ---------------------------

// Cdc is the global app codec
var Cdc *codec.Codec

// MakeCodec registers Kava app and locally defined txs/msgs
func MakeCodec() *codec.Codec {
	Cdc = codec.New()

	// Register Kava app
	app.ModuleBasics.RegisterCodec(Cdc)
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
