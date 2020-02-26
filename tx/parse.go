package tx

import (
	"github.com/tendermint/go-amino"
)

func ParseTx(cdc *amino.Codec, txBytes []byte) (Tx, error) {
	var parsedTx StdTx
	err := cdc.UnmarshalBinaryLengthPrefixed(txBytes, &parsedTx)

	if err != nil {
		return nil, err
	}

	return parsedTx, nil
}

// parse the indexed txs into an array of Info
func FormatTxResults(cdc *amino.Codec, res []*ResultTx) ([]Info, error) {
	var err error
	out := make([]Info, len(res))
	for i := range res {
		out[i], err = formatTxResult(cdc, res[i])
		if err != nil {
			return nil, err
		}
	}
	return out, nil
}

func formatTxResult(cdc *amino.Codec, res *ResultTx) (Info, error) {
	parsedTx, err := ParseTx(cdc, res.Tx)
	if err != nil {
		return Info{}, err
	}
	return Info{
		Hash:   res.Hash,
		Height: res.Height,
		Tx:     parsedTx,
		Result: res.TxResult,
	}, nil
}
