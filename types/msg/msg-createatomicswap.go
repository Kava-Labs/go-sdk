package msg

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/crypto"
	cmn "github.com/tendermint/tendermint/libs/common"

	"github.com/kava-labs/go-sdk/types/bep3"
)

// ensure Msg interface compliance at compile time
var (
	_                      sdk.Msg = &MsgCreateAtomicSwap{}
	_                      sdk.Msg = &MsgClaimAtomicSwap{}
	_                      sdk.Msg = &MsgRefundAtomicSwap{}
	AtomicSwapCoinsAccAddr         = sdk.AccAddress(crypto.AddressHash([]byte("KavaAtomicSwapCoins")))
	// kava prefix address:  [INSERT BEP3-DEPUTY ADDRESS] // TODO:
	// tkava prefix address: [INSERT BEP3-DEPUTY ADDRESS]
)

// MsgCreateAtomicSwap contains an AtomicSwap struct
type MsgCreateAtomicSwap struct {
	From                sdk.AccAddress `json:"from"  yaml:"from"`
	To                  sdk.AccAddress `json:"to"  yaml:"to"`
	RecipientOtherChain string         `json:"recipient_other_chain"  yaml:"recipient_other_chain"`
	SenderOtherChain    string         `json:"sender_other_chain"  yaml:"sender_other_chain"`
	RandomNumberHash    cmn.HexBytes   `json:"random_number_hash"  yaml:"random_number_hash"`
	Timestamp           int64          `json:"timestamp"  yaml:"timestamp"`
	Amount              sdk.Coins      `json:"amount"  yaml:"amount"`
	ExpectedIncome      string         `json:"expected_income"  yaml:"expected_income"`
	HeightSpan          int64          `json:"height_span"  yaml:"height_span"`
	CrossChain          bool           `json:"cross_chain"  yaml:"cross_chain"`
}

// NewMsgCreateAtomicSwap initializes a new MsgCreateAtomicSwap
func NewMsgCreateAtomicSwap(from sdk.AccAddress, to sdk.AccAddress, recipientOtherChain,
	senderOtherChain string, randomNumberHash cmn.HexBytes, timestamp int64,
	amount sdk.Coins, expectedIncome string, heightSpan int64, crossChain bool) MsgCreateAtomicSwap {
	return MsgCreateAtomicSwap{
		From:                from,
		To:                  to,
		RecipientOtherChain: recipientOtherChain,
		SenderOtherChain:    senderOtherChain,
		RandomNumberHash:    randomNumberHash,
		Timestamp:           timestamp,
		Amount:              amount,
		ExpectedIncome:      expectedIncome,
		HeightSpan:          heightSpan,
		CrossChain:          crossChain,
	}
}

// Route establishes the route for the MsgCreateAtomicSwap
func (msg MsgCreateAtomicSwap) Route() string { return bep3.RouterKey }

// Type is the name of MsgCreateAtomicSwap
func (msg MsgCreateAtomicSwap) Type() string { return bep3.CreateAtomicSwap }

// String prints the MsgCreateAtomicSwap
func (msg MsgCreateAtomicSwap) String() string {
	return fmt.Sprintf("AtomicSwap{%v#%v#%v#%v#%v#%v#%v#%v#%v#%v}",
		msg.From, msg.To, msg.RecipientOtherChain, msg.SenderOtherChain,
		msg.RandomNumberHash, msg.Timestamp, msg.Amount, msg.ExpectedIncome,
		msg.HeightSpan, msg.CrossChain)
}

// GetInvolvedAddresses gets the addresses involved in a MsgCreateAtomicSwap
func (msg MsgCreateAtomicSwap) GetInvolvedAddresses() []sdk.AccAddress {
	return append(msg.GetSigners(), AtomicSwapCoinsAccAddr)
}

// GetSigners gets the signers of a MsgCreateAtomicSwap
func (msg MsgCreateAtomicSwap) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.From}
}

// ValidateBasic validates the MsgCreateAtomicSwap
func (msg MsgCreateAtomicSwap) ValidateBasic() sdk.Error {
	if len(msg.From) != bep3.AddrByteCount {
		return sdk.ErrInternal(fmt.Sprintf("the expected address length is %d, actual length is %d", bep3.AddrByteCount, len(msg.From)))
	}
	if len(msg.To) != bep3.AddrByteCount {
		return sdk.ErrInternal(fmt.Sprintf("the expected address length is %d, actual length is %d", bep3.AddrByteCount, len(msg.To)))
	}
	if !msg.CrossChain && len(msg.RecipientOtherChain) != 0 {
		return sdk.ErrInternal(fmt.Sprintf("must leave recipient address on other chain to empty for single chain swap"))
	}
	if !msg.CrossChain && len(msg.SenderOtherChain) != 0 {
		return sdk.ErrInternal(fmt.Sprintf("must leave sender address on other chain to empty for single chain swap"))
	}
	if msg.CrossChain && len(msg.RecipientOtherChain) == 0 {
		return sdk.ErrInternal(fmt.Sprintf("missing recipient address on other chain for cross chain swap"))
	}
	if len(msg.RecipientOtherChain) > bep3.MaxOtherChainAddrLength {
		return sdk.ErrInternal(fmt.Sprintf("the length of recipient address on other chain should be less than %d", bep3.MaxOtherChainAddrLength))
	}
	if len(msg.SenderOtherChain) > bep3.MaxOtherChainAddrLength {
		return sdk.ErrInternal(fmt.Sprintf("the length of sender address on other chain should be less than %d", bep3.MaxOtherChainAddrLength))
	}
	if len(msg.RandomNumberHash) != bep3.RandomNumberHashLength {
		return sdk.ErrInternal(fmt.Sprintf("the length of random number hash should be %d", bep3.RandomNumberHashLength))
	}
	if msg.Timestamp <= 0 {
		return sdk.ErrInternal("timestamp must be positive")
	}
	if !msg.Amount.IsAllPositive() {
		return sdk.ErrInternal(fmt.Sprintf("the swapped out coin must be positive"))
	}
	if len(msg.ExpectedIncome) > bep3.MaxExpectedIncomeLength {
		return sdk.ErrInternal(fmt.Sprintf("the length of expected income should be less than %d", bep3.MaxExpectedIncomeLength))
	}
	expectedIncomeCoins, err := sdk.ParseCoins(msg.ExpectedIncome)
	if err != nil || expectedIncomeCoins == nil {
		return sdk.ErrInternal(fmt.Sprintf("expected income %s must be in valid format e.g. 10000ukava", msg.ExpectedIncome))
	}
	if expectedIncomeCoins.IsAnyGT(msg.Amount) {
		return sdk.ErrInternal(fmt.Sprintf("expected income %s cannot be greater than amount %s", msg.ExpectedIncome, msg.Amount.String()))
	}
	if msg.HeightSpan <= 0 {
		return sdk.ErrInternal("height span  must be positive")
	}
	return nil
}

// GetSignBytes gets the sign bytes of a MsgCreateAtomicSwap
func (msg MsgCreateAtomicSwap) GetSignBytes() []byte {
	bz := MsgCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}
