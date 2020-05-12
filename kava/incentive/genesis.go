package incentive

// GenesisClaimPeriodID stores the next claim id and its corresponding denom
type GenesisClaimPeriodID struct {
	Denom string `json:"denom" yaml:"denom"`
	ID    uint64 `json:"id" yaml:"id"`
}

// GenesisClaimPeriodIDs array of GenesisClaimPeriodID
type GenesisClaimPeriodIDs []GenesisClaimPeriodID
