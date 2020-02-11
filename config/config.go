package config

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Config struct {
	Chain   string         `json:"chain"`
	RpcAddr string         `json:"rpc_addr"`
	Deputy  sdk.AccAddress `json:"deputy"`
	// KeyType                    string `json:"key_type"`
	// AWSRegion                  string `json:"aws_region"`
	// AWSSecretName              string `json:"aws_secret_name"`
	// Mnemonic                   string `json:"mnemonic"`
	// Symbol                     string `json:"symbol"`
	// FetchInterval              int64  `json:"fetch_interval"`
	// TokenBalanceAlertThreshold int64  `json:"token_balance_alert_threshold"`
	// KavaBalanceAlertThreshold  int64  `json:"kava_balance_alert_threshold"`
}

func NewConfig(chain string, rpcAddr string, deputy sdk.AccAddress) Config {
	return Config{
		Chain:   chain,
		RpcAddr: rpcAddr,
		Deputy:  deputy,
	}
}

// GetChain returns unique name of the chain(like BNB, ETH and etc)
func (c *Config) GetChain() string {
	return c.Chain
}

// GetDeputyAddress returns deputy account address
func (c *Config) GetDeputyAddress() sdk.AccAddress {
	return c.Deputy
}
