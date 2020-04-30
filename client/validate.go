package client

import (
	"fmt"

	tmbytes "github.com/tendermint/tendermint/libs/bytes"
	tmtypes "github.com/tendermint/tendermint/types"
)

const (
	maxABCIPathLength = 1024
	maxABCIDataLength = 1024 * 1024
)

var (
	ExceedABCIPathLengthError = fmt.Errorf("the abci path exceeds max length %d ", maxABCIPathLength)
	ExceedABCIDataLengthError = fmt.Errorf("the abci data exceeds max length %d ", maxABCIDataLength)
)

// ValidateTx validates a Tendermint transaction
func ValidateTx(tx tmtypes.Tx) error {
	maxTxLength := 1024 * 1024
	exceedTxLengthError := fmt.Errorf("the tx data exceeds max length %d ", maxTxLength)

	if len(tx) > maxTxLength {
		return exceedTxLengthError
	}
	return nil
}

// ValidateABCIQuery validates an ABCI query
func ValidateABCIQuery(path string, data tmbytes.HexBytes) error {
	if err := ValidateABCIPath(path); err != nil {
		return err
	}
	if err := ValidateABCIData(data); err != nil {
		return err
	}
	return nil
}

// ValidateABCIPath validates an ABCI query's path
func ValidateABCIPath(path string) error {
	if len(path) > maxABCIPathLength {
		return ExceedABCIPathLengthError
	}
	return nil
}

// ValidateABCIData validates an ABCI query's data
func ValidateABCIData(data tmbytes.HexBytes) error {
	if len(data) > maxABCIDataLength {
		return ExceedABCIPathLengthError
	}
	return nil
}
