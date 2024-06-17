package keeper

import (
	"context"
	"wasmapp/x/onchainkeeper/types"

	onchainkeepertypes "wasmapp/x/onchainkeeper/types"

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

// CronContract returns the cron contract information
func (q Querier) CronContract(goCtx context.Context, req *types.QueryCronContract) (*types.QueryCronContractResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := onchainkeepertypes.ValidateAddresses(req.ContractAddress); err != nil {
		return nil, err
	}

	cron, err := q.keeper.GetCronContract(ctx, req.ContractAddress)
	if err != nil {
		return nil, err
	}

	return &types.QueryCronContractResponse{
		CronContract: *cron,
	}, nil
}

// CronContracts returns all cron contracts info
func (q Querier) CronContracts(goCtx context.Context, req *types.QueryCronContracts) (*types.QueryCronContractsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	crons, err := q.keeper.GetPaginatedCronContracts(ctx, req.Pagination)
	if err != nil {
		return nil, err
	}

	return crons, nil

}

func (q Querier) Params(goCtx context.Context, req *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	return &types.QueryParamsResponse{Params: q.keeper.GetParams(ctx)}, nil
}
