package cdp

import "github.com/kava-labs/cosmos-sdk/codec"

// ModuleCdc generic sealed codec to be used throughout module
var ModuleCdc *codec.Codec

func init() {
	cdc := codec.New()
	RegisterCodec(cdc)
	codec.RegisterCrypto(cdc)
	ModuleCdc = cdc.Seal() // TODO: seal??
}

// RegisterCodec registers the necessary types for cdp module
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgCreateCDP{}, "cdp/MsgCreateCDP", nil)
	cdc.RegisterConcrete(MsgDeposit{}, "cdp/MsgDeposit", nil)
	cdc.RegisterConcrete(MsgWithdraw{}, "cdp/MsgWithdraw", nil)
	cdc.RegisterConcrete(MsgDrawDebt{}, "cdp/MsgDrawDebt", nil)
	cdc.RegisterConcrete(MsgRepayDebt{}, "cdp/MsgRepayDebt", nil)
}
