package rest

import (
	"net/http"

	"github.com/kava-labs/cosmos-sdk/client/context"
	sdk "github.com/kava-labs/cosmos-sdk/types"
	"github.com/kava-labs/cosmos-sdk/types/rest"
	"github.com/kava-labs/cosmos-sdk/x/auth/client/utils"
	govrest "github.com/kava-labs/cosmos-sdk/x/gov/client/rest"
	govtypes "github.com/kava-labs/cosmos-sdk/x/gov/types"
)

// PostGovProposalReq is a rest handler for for the gov module, that handles committee change/delete proposals.
type PostGovProposalReq struct {
	BaseReq  rest.BaseReq     `json:"base_req" yaml:"base_req"`
	Content  govtypes.Content `json:"content" yaml:"content"`
	Proposer sdk.AccAddress   `json:"proposer" yaml:"proposer"`
	Deposit  sdk.Coins        `json:"deposit" yaml:"deposit"`
}

func ProposalRESTHandler(cliCtx context.CLIContext) govrest.ProposalRESTHandler {
	return govrest.ProposalRESTHandler{
		SubRoute: "committee",
		Handler:  postGovProposalHandlerFn(cliCtx),
	}
}

func postGovProposalHandlerFn(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// Parse and validate http request body
		var req PostGovProposalReq
		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			return
		}
		req.BaseReq = req.BaseReq.Sanitize()
		if !req.BaseReq.ValidateBasic(w) {
			return
		}
		if err := req.Content.ValidateBasic(); err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		// Create and return a StdTx
		msg := govtypes.NewMsgSubmitProposal(req.Content, req.Deposit, req.Proposer)
		if err := msg.ValidateBasic(); err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		utils.WriteGenerateStdTxResponse(w, cliCtx, req.BaseReq, []sdk.Msg{msg})
	}
}
