package msg

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/kava-labs/go-sdk/types/bep3"
	cmn "github.com/tendermint/tendermint/libs/common"
)

// MsgRefundAtomicSwap defines a refund msg
type MsgRefundAtomicSwap struct {
	From   sdk.AccAddress `json:"from" yaml:"from"`
	SwapID cmn.HexBytes   `json:"swap_id" yaml:"swap_id"`
}

// NewMsgRefundAtomicSwap initializes a new MsgRefundAtomicSwap
func NewMsgRefundAtomicSwap(from sdk.AccAddress, swapID []byte) MsgRefundAtomicSwap {
	return MsgRefundAtomicSwap{
		From:   from,
		SwapID: swapID,
	}
}

// Route establishes the route for the MsgRefundAtomicSwap
func (msg MsgRefundAtomicSwap) Route() string { return bep3.RouterKey }

// Type is the name of MsgRefundAtomicSwap
func (msg MsgRefundAtomicSwap) Type() string { return bep3.RefundAtomicSwap }

// String prints the MsgRefundAtomicSwap
func (msg MsgRefundAtomicSwap) String() string {
	return fmt.Sprintf("refundAtomicSwap{%v#%v}", msg.From, msg.SwapID)
}

// GetInvolvedAddresses gets the addresses involved in a MsgRefundAtomicSwap
func (msg MsgRefundAtomicSwap) GetInvolvedAddresses() []sdk.AccAddress {
	return append(msg.GetSigners(), AtomicSwapCoinsAccAddr)
}

// GetSigners gets the signers of a MsgRefundAtomicSwap
func (msg MsgRefundAtomicSwap) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.From}
}

// ValidateBasic validates the MsgRefundAtomicSwap
func (msg MsgRefundAtomicSwap) ValidateBasic() sdk.Error {
	if len(msg.From) != bep3.AddrByteCount {
		return sdk.ErrInternal(fmt.Sprintf("the expected address length is %d, actual length is %d", bep3.AddrByteCount, len(msg.From)))
	}
	if len(msg.SwapID) != bep3.SwapIDLength {
		return sdk.ErrInternal(fmt.Sprintf("the length of swapID should be %d", bep3.SwapIDLength))
	}
	return nil
}

// GetSignBytes gets the sign bytes of a MsgRefundAtomicSwap
func (msg MsgRefundAtomicSwap) GetSignBytes() []byte {
	bz := MsgCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}
