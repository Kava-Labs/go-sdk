module github.com/kava-labs/go-sdk

go 1.13

require (
	github.com/btcsuite/btcd v0.20.1-beta
	github.com/btcsuite/btcutil v1.0.2
	github.com/cosmos/go-bip39 v0.0.0-20180819234021-555e2067c45d
	github.com/kava-labs/cosmos-sdk v0.34.4-0.20200506043356-5d772797f9a3
	github.com/kava-labs/tendermint v0.33.4-0.20200506042050-c611c5308a53
	github.com/pkg/errors v0.9.1
	github.com/spf13/cobra v1.0.0
	github.com/stretchr/testify v1.5.1
	github.com/stumble/gorocksdb v0.0.3 // indirect
	github.com/tendermint/go-amino v0.15.1
	github.com/zondax/ledger-go v0.11.0 // indirect
	gopkg.in/yaml.v2 v2.2.8
)

replace github.com/zondax/ledger-go => github.com/binance-chain/ledger-go v0.9.1

replace github.com/tendermint/tendermint => github.com/kava-labs/tendermint v0.33.4-0.20200506042050-c611c5308a53

replace github.com/tendermint/iavl => github.com/kava-labs/iavl v0.13.4-0.20200506042627-f849adb79934

replace github.com/tendermint/tm-db => github.com/kava-labs/tm-db v0.4.2-0.20200506040135-3f7b09feebcd

replace github.com/cosmos/cosmos-sdk => github.com/kava-labs/cosmos-sdk v0.34.4-0.20200506043356-5d772797f9a3
