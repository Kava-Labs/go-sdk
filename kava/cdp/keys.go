package cdp

const (
	// ModuleName The name that will be used throughout the module
	ModuleName = "cdp"

	// StoreKey Top level store key where all module items will be stored
	StoreKey = ModuleName

	// RouterKey Top level router key
	RouterKey = ModuleName

	// QuerierRoute Top level query string
	QuerierRoute = ModuleName

	// DefaultParamspace default name for parameter store
	DefaultParamspace = ModuleName

	// LiquidatorMacc module account for liquidator
	LiquidatorMacc = "liquidator"

	// SavingsRateMacc module account for savings rate
	SavingsRateMacc = "savings"
)

// KVStore key prefixes
var (
	CdpIDKeyPrefix              = []byte{0x00}
	CdpKeyPrefix                = []byte{0x01}
	CollateralRatioIndexPrefix  = []byte{0x02}
	CdpIDKey                    = []byte{0x03}
	DebtDenomKey                = []byte{0x04}
	GovDenomKey                 = []byte{0x05}
	DepositKeyPrefix            = []byte{0x06}
	PrincipalKeyPrefix          = []byte{0x07}
	PreviousDistributionTimeKey = []byte{0x08}
	PricefeedStatusKeyPrefix    = []byte{0x09}
)
