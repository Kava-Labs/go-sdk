package kava

import (
	"github.com/kava-labs/cosmos-sdk/codec"
	sdk "github.com/kava-labs/cosmos-sdk/types"
	"github.com/kava-labs/cosmos-sdk/types/module"

	"github.com/kava-labs/cosmos-sdk/x/auth"
	"github.com/kava-labs/cosmos-sdk/x/auth/vesting"
	"github.com/kava-labs/cosmos-sdk/x/bank"
	"github.com/kava-labs/cosmos-sdk/x/crisis"
	distr "github.com/kava-labs/cosmos-sdk/x/distribution"
	"github.com/kava-labs/cosmos-sdk/x/evidence"
	"github.com/kava-labs/cosmos-sdk/x/genutil"
	"github.com/kava-labs/cosmos-sdk/x/gov"
	"github.com/kava-labs/cosmos-sdk/x/mint"
	"github.com/kava-labs/cosmos-sdk/x/params"
	paramsclient "github.com/kava-labs/cosmos-sdk/x/params/client"
	"github.com/kava-labs/cosmos-sdk/x/slashing"
	"github.com/kava-labs/cosmos-sdk/x/staking"
	"github.com/kava-labs/cosmos-sdk/x/supply"
	"github.com/kava-labs/cosmos-sdk/x/upgrade"
	upgradeclient "github.com/kava-labs/cosmos-sdk/x/upgrade/client"

	"github.com/kava-labs/go-sdk/kava/types/auction"
	"github.com/kava-labs/go-sdk/kava/types/bep3"
	"github.com/kava-labs/go-sdk/kava/types/cdp"
	"github.com/kava-labs/go-sdk/kava/types/committee"
	commclient "github.com/kava-labs/go-sdk/kava/types/committee/client"
	"github.com/kava-labs/go-sdk/kava/types/incentive"
	"github.com/kava-labs/go-sdk/kava/types/kavadist"
	"github.com/kava-labs/go-sdk/kava/types/pricefeed"
	"github.com/kava-labs/go-sdk/kava/types/validatorvesting"
)

const (
	Bech32MainPrefix = "kava"
	Bip44CoinType    = 459
)

// var (
// 	_ module.AppModule           = AppModule{}
// 	_ module.AppModuleBasic      = AppModuleBasic{}
// 	_ module.AppModuleSimulation = AppModule{}
// )

// --------------------------- Codec ---------------------------

// Cdc is the global app codec
// var Cdc *codec.Codec

// MakeCodec registers Kava app and locally defined txs/msgs
func MakeCodec() *codec.Codec {
	cdc := codec.New()

	// ModuleBasics is used for cosmos-sdk codec registration as not all modules have exported RegisterCodec methods
	moduleBasics := module.NewBasicManager(
		genutil.AppModuleBasic{},
		auth.AppModuleBasic{},
		bank.AppModuleBasic{},
		staking.AppModuleBasic{},
		mint.AppModuleBasic{},
		distr.AppModuleBasic{},
		gov.NewAppModuleBasic(
			paramsclient.ProposalHandler, distr.ProposalHandler, commclient.ProposalHandler,
			upgradeclient.ProposalHandler,
		),
		params.AppModuleBasic{},
		crisis.AppModuleBasic{},
		slashing.AppModuleBasic{},
		upgrade.AppModuleBasic{},
		supply.AppModuleBasic{},
		evidence.AppModuleBasic{},
	)
	moduleBasics.RegisterCodec(cdc)

	// Register custom Kava module msgs
	validatorvesting.RegisterCodec(cdc)
	auction.RegisterCodec(cdc)
	cdp.RegisterCodec(cdc)
	pricefeed.RegisterCodec(cdc)
	committee.RegisterCodec(cdc)
	bep3.RegisterCodec(cdc)
	kavadist.RegisterCodec(cdc)
	incentive.RegisterCodec(cdc)

	// Register vesting, codec, crypto, and evidences in the same order as app
	vesting.RegisterCodec(cdc)
	sdk.RegisterCodec(cdc)
	codec.RegisterCrypto(cdc)
	codec.RegisterEvidences(cdc)

	return cdc.Seal()
}

// --------------------------- Config ---------------------------

// SetBech32AddressPrefixes sets the global prefix to be used when serializing addresses to bech32 strings.
func SetBech32AddressPrefixes(config *sdk.Config) {
	config.SetBech32PrefixForAccount(Bech32MainPrefix, Bech32MainPrefix+sdk.PrefixPublic)
	config.SetBech32PrefixForValidator(Bech32MainPrefix+sdk.PrefixValidator+sdk.PrefixOperator, Bech32MainPrefix+sdk.PrefixValidator+sdk.PrefixOperator+sdk.PrefixPublic)
	config.SetBech32PrefixForConsensusNode(Bech32MainPrefix+sdk.PrefixValidator+sdk.PrefixConsensus, Bech32MainPrefix+sdk.PrefixValidator+sdk.PrefixConsensus+sdk.PrefixPublic)
}
