package msg

import (
	"github.com/tendermint/go-amino"
)

// TODO: Register all custom msg types

var MsgCdc = amino.NewCodec()

func RegisterCodec(cdc *amino.Codec) {
	cdc.RegisterConcrete(MsgCreateAtomicSwap{}, "bep3/MsgCreateAtomicSwap", nil)
	cdc.RegisterConcrete(MsgClaimAtomicSwap{}, "bep3/MsgClaimAtomicSwap", nil)
	cdc.RegisterConcrete(MsgRefundAtomicSwap{}, "bep3/MsgRefundAtomicSwap", nil)
}

func init() {
	RegisterCodec(MsgCdc)
}
