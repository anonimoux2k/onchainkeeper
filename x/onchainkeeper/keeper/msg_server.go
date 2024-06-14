package keeper

import (
	"context"

	errorsmod "cosmossdk.io/errors"

	"wasmapp/x/onchainkeeper/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ types.MsgServer = msgServer{}

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

// RegisterCronContract implements types.MsgServer.
func (k msgServer) RegisterCronContract(context.Context, *types.MsgRegisterCronContract) (*types.MsgRegisterCronContractResponse, error) {
	panic("unimplemented")
}

// UnregisterCronContract implements types.MsgServer.
func (k msgServer) UnregisterCronContract(context.Context, *types.MsgUnregisterCronContract) (*types.MsgUnregisterCronContractResponse, error) {
	panic("unimplemented")
}

// AcceptPendingRegisteredCronContract implements types.MsgServer.
func (k msgServer) AcceptPendingRegisteredCronContract(context.Context, *types.MsgAcceptPendingRegisteredCronContract) (*types.MsgAcceptPendingRegisteredCronContractResponse, error) {
	panic("unimplemented")
}

// ReactivateDeactivatedCronContract implements types.MsgServer.
func (k msgServer) ReactivateDeactivatedCronContract(context.Context, *types.MsgReactivateDeactivatedCronContract) (*types.MsgReactivateDeactivatedCronContractResponse, error) {
	panic("unimplemented")
}

func (k msgServer) UpdateParams(goCtx context.Context, req *types.MsgUpdateParams) (*types.MsgUpdateParamsResponse, error) {
	if k.GetAuthority() != req.Authority {
		return nil, errorsmod.Wrapf(types.ErrInvalidSigner, "invalid authority; expected %s, got %s", k.GetAuthority(), req.Authority)
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	if err := k.SetParams(ctx, req.Params); err != nil {
		return nil, err
	}

	return &types.MsgUpdateParamsResponse{}, nil
}
