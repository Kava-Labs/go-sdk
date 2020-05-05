package msgs

import (
	"github.com/kava-labs/cosmos-sdk/codec"
)

var MsgCdc = codec.New()

// RegisterCodec registers concrete types on amino
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgCreateAtomicSwap{}, "bep3/MsgCreateAtomicSwap", nil)
	cdc.RegisterConcrete(MsgRefundAtomicSwap{}, "bep3/MsgRefundAtomicSwap", nil)
	cdc.RegisterConcrete(MsgClaimAtomicSwap{}, "bep3/MsgClaimAtomicSwap", nil)
}

func init() {
	RegisterCodec(MsgCdc)
}
