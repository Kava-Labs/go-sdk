package tx

import (
	"github.com/tendermint/go-amino"
	cryptoamino "github.com/tendermint/tendermint/crypto/encoding/amino"

	"github.com/kava-labs/go-sdk/types/msg"
)

// cdc global variable
var Cdc = amino.NewCodec()

// TODO: It appears that we only need to register custom types
// ---------------- sdk ------------------
// sdk "github.com/cosmos/cosmos-sdk/types"
// cdc.RegisterInterface((*sdk.Tx)(nil), nil)

// --------------- auth ------------------
// authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
// cdc.RegisterConcrete(authtypes.StdTx{}, "auth/StdTx", nil)

func RegisterCodec(cdc *amino.Codec) {
	msg.RegisterCodec(cdc)
}

func init() {
	cryptoamino.RegisterAmino(Cdc)
	RegisterCodec(Cdc)
}
