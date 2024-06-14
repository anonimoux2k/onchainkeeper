package keeper

import (
	"context"
	"wasmapp/x/onchainkeeper/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ types.QueryServer = &Querier{}

type Querier struct {
	keeper Keeper
}

func NewQuerier(k Keeper) Querier {
	return Querier{
		keeper: k,
	}
}

// CronContract implements types.QueryServer.
func (q Querier) CronContract(context.Context, *types.QueryCronContract) (*types.QueryCronContractResponse, error) {
	panic("unimplemented")
}

// CronContracts implements types.QueryServer.
func (q Querier) CronContracts(context.Context, *types.QueryCronContracts) (*types.QueryCronContractsResponse, error) {
	panic("unimplemented")
}

func (q Querier) Params(goCtx context.Context, req *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	return &types.QueryParamsResponse{Params: q.keeper.GetParams(ctx)}, nil
}
