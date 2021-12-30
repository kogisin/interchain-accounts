package client

import (
	govclient "github.com/cosmos/cosmos-sdk/x/gov/client"

	"github.com/cosmos/interchain-accounts/x/inter-tx/client/cli"
	"github.com/cosmos/interchain-accounts/x/inter-tx/client/rest"
)

// ProposalHandler is the register interchain account command handler.
// Note that rest.ProposalRESTHandler will be deprecated in the future.
var (
	ProposalHandler = govclient.NewProposalHandler(cli.GetCmdSubmitRegisterInterchainAccountProposal, rest.ProposalRESTHandler)
)
