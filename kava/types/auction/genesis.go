package auction

import (
	sdk "github.com/kava-labs/cosmos-sdk/types"
)

// GenesisAuction is an interface that extends the auction interface to add functionality needed for initializing auctions from genesis.
type GenesisAuction interface {
	Auction
	GetModuleAccountCoins() sdk.Coins
	Validate() error
}

// GenesisAuctions is a slice of genesis auctions.
type GenesisAuctions []GenesisAuction
