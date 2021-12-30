package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmos/interchain-accounts/x/inter-tx/types"
)

// HandleRegisterInterchainAccountProposal is a handler for executing a public plan creation proposal.
func HandleRegisterInterchainAccountProposal(ctx sdk.Context, k Keeper, proposal *types.RegisterInterchainAccountProposal) error {
	params := k.GetParams(ctx)
	acc, err := sdk.AccAddressFromBech32(params.InterchainAccountController)
	if err != nil {
		return err
	}

	if err := k.RegisterInterchainAccount(ctx, acc, proposal.ConnectionId, proposal.CounterpartyConnectionId); err != nil {
		return err
	}

	return nil
}
