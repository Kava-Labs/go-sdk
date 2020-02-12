package client

import (
	"os"

	cfg "github.com/tendermint/tendermint/config"
	"github.com/tendermint/tendermint/libs/log"
	"github.com/tendermint/tendermint/node"
	"github.com/tendermint/tendermint/rpc/client"
)

// NewLocalClient initializes a new local client
func NewLocalClient() (*client.Local, error) {
	logger := log.NewTMLogger(log.NewSyncWriter(os.Stdout))
	config := cfg.DefaultConfig()
	node, err := node.DefaultNewNode(config, logger)
	if err != nil {
		return nil, err
	}

	local := client.NewLocal(node)
	return local, nil
}
