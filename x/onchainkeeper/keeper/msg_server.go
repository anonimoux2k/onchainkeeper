package keeper

import (
	"context"
	"wasmapp/x/onchainkeeper/types"
)

type msgServer struct {
	Keeper
}

// AcceptPendingRegisteredCronContract implements types.MsgServer.
func (k msgServer) AcceptPendingRegisteredCronContract(context.Context, *types.MsgAcceptPendingRegisteredCronContract) (*types.MsgAcceptPendingRegisteredCronContractResponse, error) {
	panic("unimplemented")
}

// ReactivateDeactivatedCronContract implements types.MsgServer.
func (k msgServer) ReactivateDeactivatedCronContract(context.Context, *types.MsgReactivateDeactivatedCronContract) (*types.MsgReactivateDeactivatedCronContractResponse, error) {
	panic("unimplemented")
}

// RegisterCronContract implements types.MsgServer.
func (k msgServer) RegisterCronContract(context.Context, *types.MsgRegisterCronContract) (*types.MsgRegisterCronContractResponse, error) {
	panic("unimplemented")
}

// UnregisterCronContract implements types.MsgServer.
func (k msgServer) UnregisterCronContract(context.Context, *types.MsgUnregisterCronContract) (*types.MsgUnregisterCronContractResponse, error) {
	panic("unimplemented")
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}
