package bep3

import (
	"encoding/json"

	sdk "github.com/cosmos/cosmos-sdk/types"
	cmn "github.com/tendermint/tendermint/libs/common"
)

// AtomicSwap contains the information for an atomic swap
type AtomicSwap struct {
	Amount              sdk.Coins      `json:"amount"  yaml:"amount"`
	RandomNumberHash    cmn.HexBytes   `json:"random_number_hash"  yaml:"random_number_hash"`
	ExpireHeight        int64          `json:"expire_height"  yaml:"expire_height"`
	Timestamp           int64          `json:"timestamp"  yaml:"timestamp"`
	Sender              sdk.AccAddress `json:"sender"  yaml:"sender"`
	Recipient           sdk.AccAddress `json:"recipient"  yaml:"recipient"`
	SenderOtherChain    string         `json:"sender_other_chain"  yaml:"sender_other_chain"`
	RecipientOtherChain string         `json:"recipient_other_chain"  yaml:"recipient_other_chain"`
	ClosedBlock         int64          `json:"closed_block"  yaml:"closed_block"`
	Status              SwapStatus     `json:"status"  yaml:"status"`
	CrossChain          bool           `json:"cross_chain"  yaml:"cross_chain"`
	Direction           SwapDirection  `json:"direction"  yaml:"direction"`
}

// SwapStatus is the status of an AtomicSwap
type SwapStatus byte

const (
	NULL      SwapStatus = 0x00
	Open      SwapStatus = 0x01
	Completed SwapStatus = 0x02
	Expired   SwapStatus = 0x03
)

// NewSwapStatusFromString converts string to SwapStatus type
func NewSwapStatusFromString(str string) SwapStatus {
	switch str {
	case "Open", "open":
		return Open
	case "Completed", "completed":
		return Completed
	case "Expired", "expired":
		return Expired
	default:
		return NULL
	}
}

// String returns the string representation of a SwapStatus
func (status SwapStatus) String() string {
	switch status {
	case Open:
		return "Open"
	case Completed:
		return "Completed"
	case Expired:
		return "Expired"
	default:
		return "NULL"
	}
}

// MarshalJSON marshals the SwapStatus
func (status SwapStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(status.String())
}

// UnmarshalJSON unmarshals the SwapStatus
func (status *SwapStatus) UnmarshalJSON(data []byte) error {
	var s string
	err := json.Unmarshal(data, &s)
	if err != nil {
		return err
	}
	*status = NewSwapStatusFromString(s)
	return nil
}

// SwapDirection is the direction of an AtomicSwap
type SwapDirection byte

const (
	INVALID  SwapDirection = 0x00
	Incoming SwapDirection = 0x01
	Outgoing SwapDirection = 0x02
)

// NewSwapDirectionFromString converts string to SwapDirection type
func NewSwapDirectionFromString(str string) SwapDirection {
	switch str {
	case "Incoming", "incoming", "inc", "I", "i":
		return Incoming
	case "Outgoing", "outgoing", "out", "O", "o":
		return Outgoing
	default:
		return INVALID
	}
}

// String returns the string representation of a SwapDirection
func (direction SwapDirection) String() string {
	switch direction {
	case Incoming:
		return "Incoming"
	case Outgoing:
		return "Outgoing"
	default:
		return "INVALID"
	}
}

// MarshalJSON marshals the SwapDirection
func (direction SwapDirection) MarshalJSON() ([]byte, error) {
	return json.Marshal(direction.String())
}

// UnmarshalJSON unmarshals the SwapDirection
func (direction *SwapDirection) UnmarshalJSON(data []byte) error {
	var s string
	err := json.Unmarshal(data, &s)
	if err != nil {
		return err
	}
	*direction = NewSwapDirectionFromString(s)
	return nil
}
