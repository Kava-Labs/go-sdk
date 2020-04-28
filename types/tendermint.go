package types

import (
	"github.com/tendermint/tendermint/crypto/tmhash"
	tmbytes "github.com/tendermint/tendermint/libs/bytes"
	"github.com/tendermint/tendermint/libs/log"
)

// HexBytes is a wrapper around tmbytes.HexBytes
type HexBytes tmbytes.HexBytes

// Sum is a wrapper around tmhash.Sum
func Sum(data []byte) []byte {
	return tmhash.Sum(data)
}

// Logger is a wrapper around log.Logger
type Logger log.Logger
