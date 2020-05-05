package types

import (
	"encoding/json"
	"fmt"

	"github.com/btcsuite/btcutil/bech32"
	sdk "github.com/kava-labs/cosmos-sdk/types"
	authtypes "github.com/kava-labs/cosmos-sdk/x/auth/types"
	"github.com/pkg/errors"
)

const (
	// KavaPrefix is the address prefix on Kava
	KavaPrefix = "kava"
)

// ------------------------- sdk.Coins -------------------------

// NewInt64Coin is a wrapper around sdk.NewInt64Coin
func NewInt64Coin(denom string, amount int64) sdk.Coin {
	return sdk.NewInt64Coin(denom, amount)
}

// Coins is a wrapper around sdk.Coins
type Coins sdk.Coins

// ToSdk returns sdk.Coins type from a Coins
func (c Coins) ToSdk() sdk.Coins {
	return sdk.Coins(c)
}

// NewCoins is a wrapper around sdk.NewCoins
func NewCoins(coins ...sdk.Coin) sdk.Coins {
	return sdk.NewCoins(coins...)
}

// ----------------------- sdk.AccAddress -----------------------

// AccAddress is a marshalable type convertable to sdk.AccAddress
type AccAddress []byte

// Marshal needed for protobuf compatibility
func (bz AccAddress) Marshal() ([]byte, error) {
	return bz, nil
}

// Unmarshal needed for protobuf compatibility
func (bz *AccAddress) Unmarshal(data []byte) error {
	*bz = data
	return nil
}

// UnmarshalJSON to Unmarshal from JSON assuming Bech32 encoding
func (bz *AccAddress) UnmarshalJSON(data []byte) error {
	var s string
	err := json.Unmarshal(data, &s)
	if err != nil {
		return nil
	}

	bz2, err := AccAddressFromBech32(s)
	if err != nil {
		return err
	}
	*bz = bz2
	return nil
}

// AccAddressFromBech32 to create an AccAddress from a bech32 string
func AccAddressFromBech32(address string) (addr AccAddress, err error) {
	bz, err := GetFromBech32(address, KavaPrefix)
	if err != nil {
		return nil, err
	}
	return AccAddress(bz), nil
}

// GetFromBech32 to decode a bytestring from a bech32-encoded string
func GetFromBech32(bech32str, prefix string) ([]byte, error) {
	if len(bech32str) == 0 {
		return nil, errors.New("decoding bech32 address failed: must provide an address")
	}
	hrp, bz, err := DecodeAndConvert(bech32str)
	if err != nil {
		return nil, err
	}

	if hrp != prefix {
		return nil, fmt.Errorf("invalid bech32 prefix. Expected %s, Got %s", prefix, hrp)
	}

	return bz, nil
}

//ConvertAndEncode converts from a base64 encoded byte string to base32 encoded byte string and then to bech32
func ConvertAndEncode(hrp string, data []byte) (string, error) {
	converted, err := bech32.ConvertBits(data, 8, 5, true)
	if err != nil {
		return "", errors.Wrap(err, "encoding bech32 failed")
	}
	return bech32.Encode(hrp, converted)

}

//DecodeAndConvert decodes a bech32 encoded string and converts to base64 encoded bytes
func DecodeAndConvert(bech string) (string, []byte, error) {
	hrp, data, err := bech32.Decode(bech)
	if err != nil {
		return "", nil, errors.Wrap(err, "decoding bech32 failed")
	}
	converted, err := bech32.ConvertBits(data, 5, 8, false)
	if err != nil {
		return "", nil, errors.Wrap(err, "decoding bech32 failed")
	}
	return hrp, converted, nil
}

// ToSdk returns sdk.AccAddress type from an AccAddress
func (bz AccAddress) ToSdk() sdk.AccAddress {
	return sdk.AccAddress(bz)
}

// ----------------------- sdk.StdSignature -----------------------

// StdSignature represents a sig
type StdSignature struct {
	authtypes.StdSignature
	// crypto.PubKey `json:"pub_key" yaml:"pub_key"` // optional
	// Signature     []byte                          `json:"signature" yaml:"signature"`
}

// TODO: func (std StdSignature) ToSdk() authtypes.StdSignature {
// 	return authtypes.StdSignature(std)
// }

// --------------------------- Other ---------------------------

// GetConfig is a wrapper around sdk.GetConfig
func GetConfig() *sdk.Config {
	return sdk.GetConfig()
}
