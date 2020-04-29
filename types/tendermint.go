package types

import (
	"github.com/tendermint/tendermint/crypto"
	"github.com/tendermint/tendermint/crypto/tmhash"
	tmbytes "github.com/tendermint/tendermint/libs/bytes"
	"github.com/tendermint/tendermint/libs/log"
	// ctypes "github.com/tendermint/tendermint/rpc/core/types"
)

// HexBytes is a wrapper around tmbytes.HexBytes
type HexBytes tmbytes.HexBytes

// Sum is a wrapper around tmhash.Sum
func Sum(data []byte) []byte {
	return tmhash.Sum(data)
}

// Logger is a wrapper around log.Logger
type Logger log.Logger

type PrivKey interface {
	crypto.PrivKey
}

// // CheckTx result
// type ResultBroadcastTx struct {
// 	ctypes.ResultBroadcastTx
// }

// // CheckTx and DeliverTx results
// type ResultBroadcastTxCommit struct {
// 	ctypes.ResultBroadcastTxCommit
// }

// const (
// 	CodeOk int32 = 0
// )

// // ResultBroadcastTx is a replacement for ctypes.ResultBroadcastTx
// type ResultBroadcastTx struct {
// 	Hash string `json:"hash"`
// 	Log  string `json:"log"`
// 	Data string `json:"data"`
// 	Code int32  `json:"code"`
// }

// // ResultBroadcastTxCommit is a replacement for ctypes.ResultBroadcastTxCommit
// type ResultBroadcastTxCommit struct {
// 	Ok   bool   `json:"ok"`
// 	Log  string `json:"log"`
// 	Hash string `json:"hash"`
// 	Code int32  `json:"code"`
// 	Data string `json:"data"`
// }
