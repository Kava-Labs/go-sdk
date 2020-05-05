package types

import (
	"github.com/kava-labs/tendermint/crypto/tmhash"
	tmbytes "github.com/kava-labs/tendermint/libs/bytes"
	"github.com/kava-labs/tendermint/libs/log"
)

// --------------------------- Logger ---------------------------

// TendermintLogger is an interface that enables loggers to extend KavaLogger
type TendermintLogger interface {
	Debug(msg string, keyvals ...interface{})
	Info(msg string, keyvals ...interface{})
	Error(msg string, keyvals ...interface{})
}

// KavaLogger implements method With, abstracting the tendermint dependency
type KavaLogger struct {
	TendermintLogger
}

// With adds a set of key-values to KavaLogger
func (l *KavaLogger) With(keyvals ...interface{}) log.Logger {
	return l
}

// --------------------------- HexBytes ---------------------------

// HexBytes is a wrapper around tmbytes.HexBytes
type HexBytes tmbytes.HexBytes

// Sum is a wrapper around tmhash.Sum
func Sum(data []byte) []byte {
	return tmhash.Sum(data)
}

// ToTendermint returns tmbytes.HexBytes type from HexBytes
func (hb HexBytes) ToTendermint() tmbytes.HexBytes {
	return tmbytes.HexBytes(hb)
}
