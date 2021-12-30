package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	capabilitykeeper "github.com/cosmos/cosmos-sdk/x/capability/keeper"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"

	"github.com/cosmos/interchain-accounts/x/inter-tx/types"

	icacontrollerkeeper "github.com/cosmos/ibc-go/v3/modules/apps/27-interchain-accounts/controller/keeper"
)

type Keeper struct {
	cdc codec.Codec

	storeKey   sdk.StoreKey
	paramSpace paramtypes.Subspace

	scopedKeeper        capabilitykeeper.ScopedKeeper
	icaControllerKeeper icacontrollerkeeper.Keeper
}

func NewKeeper(
	cdc codec.Codec,
	storeKey sdk.StoreKey,
	paramSpace paramtypes.Subspace,
	iaKeeper icacontrollerkeeper.Keeper,
	scopedKeeper capabilitykeeper.ScopedKeeper,
) Keeper {

	return Keeper{
		cdc:                 cdc,
		storeKey:            storeKey,
		paramSpace:          paramSpace,
		scopedKeeper:        scopedKeeper,
		icaControllerKeeper: iaKeeper,
	}
}

// ClaimCapability claims the channel capability passed via the OnOpenChanInit callback
func (k *Keeper) ClaimCapability(ctx sdk.Context, cap *capabilitytypes.Capability, name string) error {
	return k.scopedKeeper.ClaimCapability(ctx, cap, name)
}

// GetParams returns the parameters for the fundraising module.
func (k Keeper) GetParams(ctx sdk.Context) (params types.Params) {
	k.paramSpace.GetParamSet(ctx, &params)
	return params
}

// SetParams sets the parameters for the fundraising module.
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramSpace.SetParamSet(ctx, &params)
}
