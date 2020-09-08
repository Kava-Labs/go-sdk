package client

// SyncType is the method type for sending transactions
type SyncType int

const (
	Async SyncType = iota
	Sync
	Commit
)
