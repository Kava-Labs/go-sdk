package keys

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (m *keyManager) Sign(msg sdk.Msg) ([]byte, error) {
	sig, err := m.makeSignature(msg)
	if err != nil {
		return nil, err
	}

	return sig, nil

	// TODO:
	// newTx := NewStdTx(msg.Msgs, []StdSignature{sig}, msg.Memo, msg.Source, msg.Data)
	// bz, err := Cdc.MarshalBinaryLengthPrefixed(&newTx)
	// if err != nil {
	// 	return nil, err
	// }
}

func (m *keyManager) makeSignature(msg sdk.Msg) (sig []byte, err error) {
	if err != nil {
		return
	}
	sigBytes, err := m.privKey.Sign(msg.GetSignBytes())
	if err != nil {
		return
	}
	return sigBytes, nil

	// TODO:
	// return StdSignature{
	// 	AccountNumber: msg.AccountNumber,
	// 	Sequence:      msg.Sequence,
	// 	PubKey:        m.privKey.PubKey(),
	// 	Signature:     sigBytes,
	// }, nil
}
