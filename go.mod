module github.com/kava-labs/go-sdk

go 1.13

require (
	github.com/btcsuite/btcd v0.20.1-beta
	github.com/coreos/go-etcd v2.0.0+incompatible // indirect
	github.com/cosmos/cosmos-sdk v0.38.3
	github.com/cosmos/go-bip39 v0.0.0-20180819234021-555e2067c45d
	github.com/cpuguy83/go-md2man v1.0.10 // indirect
	github.com/kava-labs/kava v0.7.1-0.20200427144034-ae4aee46ff44
	github.com/stretchr/testify v1.5.1
	github.com/stumble/gorocksdb v0.0.3 // indirect
	github.com/tendermint/go-amino v0.15.1
	github.com/tendermint/tendermint v0.33.3
	github.com/ugorji/go/codec v0.0.0-20181204163529-d75b2dcb6bc8 // indirect
)

replace github.com/tendermint/tendermint => github.com/tendermint/tendermint v0.33.3
