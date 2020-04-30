package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/vesting"
	"github.com/kava-labs/kava/app"
	"github.com/tendermint/go-amino"

	// "github.com/tendermint/go-amino"
	"github.com/cosmos/cosmos-sdk/codec"
)

// --------------------------- Codec ---------------------------

// config := types.GetConfig()
// 	app.SetBech32AddressPrefixes(config)
// 	app.SetBip44CoinType(config)

// TODO: Is this codec why tendermint v0.33.3 is being imported?

// Cdc is the global app codec  // TODO: Cdc = amino.NewCodec()
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
	// cdc.RegisterConcrete(StdTx{}, "auth/StdTx", nil)
	// msg.RegisterCodec(cdc)
}

// Tx is a wrapper around sdk.Tx
type Tx sdk.Tx

// func init() {
// 	// cryptoamino.RegisterAmino(Cdc) // TODO: redundant?
// 	RegisterCodec(Cdc)
// }
