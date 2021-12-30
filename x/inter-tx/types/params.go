package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/address"
)

var (
	// cosmos1mg35ppl39tdl4u7967aha7hafj90k4r08xakqjsm6qc4u2k3660qp43eyg
	InterchainAccountController = sdk.AccAddress(address.Module(ModuleName, []byte("InterchainAccountControllerAcc"))).String()
)
