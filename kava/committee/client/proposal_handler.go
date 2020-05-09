package client

import (
	govclient "github.com/kava-labs/cosmos-sdk/x/gov/client"

	"github.com/kava-labs/go-sdk/kava/committee/client/cli"
	"github.com/kava-labs/go-sdk/kava/committee/client/rest"
)

// ProposalHandler is a struct containing handler funcs for submiting CommitteeChange/Delete proposal txs to the gov module through the cli or rest.
var ProposalHandler = govclient.NewProposalHandler(cli.GetGovCmdSubmitProposal, rest.ProposalRESTHandler)
