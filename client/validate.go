package client

import (
	"fmt"

	cmn "github.com/tendermint/tendermint/libs/common"
	tmtypes "github.com/tendermint/tendermint/types"
)

const (
	maxABCIPathLength = 1024
	maxABCIDataLength = 1024 * 1024
)

var (
	ExceedABCIPathLengthError = fmt.Errorf("the abci path exceed max length %d ", maxABCIPathLength)
	ExceedABCIDataLengthError = fmt.Errorf("the abci data exceed max length %d ", maxABCIDataLength)
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
func ValidateABCIQuery(path string, data cmn.HexBytes) error {
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
func ValidateABCIData(data cmn.HexBytes) error {
	if len(data) > maxABCIDataLength {
		return ExceedABCIPathLengthError
	}
	return nil
}
