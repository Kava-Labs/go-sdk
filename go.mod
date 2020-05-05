module github.com/kava-labs/go-sdk

go 1.13

require (
	github.com/btcsuite/btcd v0.20.1-beta
	github.com/btcsuite/btcutil v1.0.2
	github.com/cosmos/go-bip39 v0.0.0-20180819234021-555e2067c45d
	github.com/kava-labs/cosmos-sdk v0.34.4-0.20200505055524-c0acebc54d70
	github.com/kava-labs/tendermint v0.33.4-0.20200505050845-6c848ee6dc48
	github.com/pkg/errors v0.9.1
	github.com/stretchr/testify v1.5.1
	github.com/stumble/gorocksdb v0.0.3 // indirect
	github.com/tendermint/go-amino v0.15.1
	github.com/zondax/ledger-go v0.11.0 // indirect
)

replace github.com/zondax/ledger-go => github.com/binance-chain/ledger-go v0.9.1
