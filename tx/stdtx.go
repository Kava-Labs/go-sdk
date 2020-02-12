package tx

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const Source int64 = 0

type Tx interface {
	// Gets the Msg.
	GetMsgs() []sdk.Msg
}

// StdTx def
type StdTx struct {
	Msgs       []sdk.Msg      `json:"msg"`
	Signatures []StdSignature `json:"signatures"`
	Memo       string         `json:"memo"`
	Source     int64          `json:"source"`
	Data       []byte         `json:"data"`
}

// NewStdTx to instantiate an instance
func NewStdTx(msgs []sdk.Msg, sigs []StdSignature, memo string, source int64, data []byte) StdTx {
	return StdTx{
		Msgs:       msgs,
		Signatures: sigs,
		Memo:       memo,
		Source:     source,
		Data:       data,
	}
}

// GetMsgs def
func (tx StdTx) GetMsgs() []sdk.Msg { return tx.Msgs }
