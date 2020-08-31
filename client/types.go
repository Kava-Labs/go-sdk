package client

// SyncType is the method type for sending transactions
type SyncType int

const (
	Async SyncType = iota
	Sync
	Commit
)

// ChainNetwork is the name of the blockchain
type ChainNetwork uint8

const (
	LocalNetwork ChainNetwork = iota
	TestNetwork
	ProdNetwork
)

const (
	// LocalChainID is for local development
	LocalChainID = "testing"
	// TestChainID is Kava's latest testnet
	TestChainID = "kava-testnet-9000"
	// ProdChainID is Kava's mainnet
	ProdChainID = "kava-3"
)
