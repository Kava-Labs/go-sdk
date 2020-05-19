module github.com/kava-labs/go-sdk

go 1.13

require (
	github.com/btcsuite/btcd v0.20.1-beta
	github.com/btcsuite/btcutil v1.0.2
	github.com/cosmos/go-bip39 v0.0.0-20180819234021-555e2067c45d
	github.com/kava-labs/cosmos-sdk v0.38.3-stable
	github.com/kava-labs/tendermint v0.33.3-stable
	github.com/pkg/errors v0.9.1
	github.com/spf13/cobra v1.0.0
	github.com/stretchr/testify v1.5.1
	github.com/stumble/gorocksdb v0.0.3 // indirect
	github.com/tendermint/go-amino v0.15.1
	github.com/zondax/ledger-go v0.11.0 // indirect
	gopkg.in/yaml.v2 v2.2.8
)

replace github.com/zondax/ledger-go => github.com/binance-chain/ledger-go v0.9.1

replace github.com/tendermint/tm-db => github.com/kava-labs/tm-db v0.4.1-stable

replace github.com/tendermint/tendermint => github.com/kava-labs/tendermint v0.33.3-stable

replace github.com/tendermint/iavl => github.com/kava-labs/iavl v0.13.3-stable

replace github.com/cosmos/cosmos-sdk => github.com/kava-labs/cosmos-sdk v0.38.3-stable
