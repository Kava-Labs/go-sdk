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
	TestNetwork ChainNetwork = iota
	ProdNetwork
)

const (
	ProdChainID = "testing"
	TestChainID = "testing"
)
