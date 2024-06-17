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
func (k msgServer) RegisterCronContract(goCtx context.Context, req *types.MsgRegisterCronContract) (*types.MsgRegisterCronContractResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := req.ValidateBasic(); err != nil {
		return nil, err
	}

	return &types.MsgRegisterCronContractResponse{}, k.RegisterContract(ctx, req.SenderAddress, req.ContractAddress)
}

// UnregisterCronContract implements types.MsgServer.
func (k msgServer) UnregisterCronContract(goCtx context.Context, req *types.MsgUnregisterCronContract) (*types.MsgUnregisterCronContractResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := req.ValidateBasic(); err != nil {
		return nil, err
	}

	return &types.MsgUnregisterCronContractResponse{}, k.UnregisterContract(ctx, req.SenderAddress, req.ContractAddress)
}

// AcceptPendingRegisteredCronContract implements types.MsgServer.
func (k msgServer) AcceptPendingRegisteredCronContract(goCtx context.Context, req *types.MsgAcceptPendingRegisteredCronContract) (*types.MsgAcceptPendingRegisteredCronContractResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := req.ValidateBasic(); err != nil {
		return nil, err
	}

	return &types.MsgAcceptPendingRegisteredCronContractResponse{}, k.ActivateCronContract(ctx, req.Authority, req.ContractAddress)
}

// ReactivateDeactivatedCronContract implements types.MsgServer.
func (k msgServer) ReactivateDeactivatedCronContract(goCtx context.Context, req *types.MsgReactivateDeactivatedCronContract) (*types.MsgReactivateDeactivatedCronContractResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := req.ValidateBasic(); err != nil {
		return nil, err
	}

	return &types.MsgReactivateDeactivatedCronContractResponse{}, k.ActivateCronContract(ctx, req.Authority, req.ContractAddress)
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
