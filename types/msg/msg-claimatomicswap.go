package msg

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/kava-labs/go-sdk/types/bep3"
	cmn "github.com/tendermint/tendermint/libs/common"
)

// MsgClaimAtomicSwap defines a AtomicSwap claim
type MsgClaimAtomicSwap struct {
	From         sdk.AccAddress `json:"from"  yaml:"from"`
	SwapID       cmn.HexBytes   `json:"swap_id"  yaml:"swap_id"`
	RandomNumber cmn.HexBytes   `json:"random_number"  yaml:"random_number"`
}

// NewMsgClaimAtomicSwap initializes a new MsgClaimAtomicSwap
func NewMsgClaimAtomicSwap(from sdk.AccAddress, swapID, randomNumber []byte) MsgClaimAtomicSwap {
	return MsgClaimAtomicSwap{
		From:         from,
		SwapID:       swapID,
		RandomNumber: randomNumber,
	}
}

// Route establishes the route for the MsgClaimAtomicSwap
func (msg MsgClaimAtomicSwap) Route() string { return bep3.RouterKey }

// Type is the name of MsgClaimAtomicSwap
func (msg MsgClaimAtomicSwap) Type() string { return bep3.ClaimAtomicSwap }

// String prints the MsgClaimAtomicSwap
func (msg MsgClaimAtomicSwap) String() string {
	return fmt.Sprintf("claimAtomicSwap{%v#%v#%v}", msg.From, msg.SwapID, msg.RandomNumber)
}

// GetInvolvedAddresses gets the addresses involved in a MsgClaimAtomicSwap
func (msg MsgClaimAtomicSwap) GetInvolvedAddresses() []sdk.AccAddress {
	return append(msg.GetSigners(), AtomicSwapCoinsAccAddr)
}

// GetSigners gets the signers of a MsgClaimAtomicSwap
func (msg MsgClaimAtomicSwap) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.From}
}

// ValidateBasic validates the MsgClaimAtomicSwap
func (msg MsgClaimAtomicSwap) ValidateBasic() sdk.Error {
	if len(msg.From) != bep3.AddrByteCount {
		return sdk.ErrInternal(fmt.Sprintf("the expected address length is %d, actual length is %d", bep3.AddrByteCount, len(msg.From)))
	}
	if len(msg.SwapID) != bep3.SwapIDLength {
		return sdk.ErrInternal(fmt.Sprintf("the length of swapID should be %d", bep3.SwapIDLength))
	}
	if len(msg.RandomNumber) == 0 {
		return sdk.ErrInternal("the length of random number cannot be 0")
	}
	return nil
}

// GetSignBytes gets the sign bytes of a MsgClaimAtomicSwap
func (msg MsgClaimAtomicSwap) GetSignBytes() []byte {
	bz := MsgCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}
