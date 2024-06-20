package keeper

import (
	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"

	"wasmapp/x/onchainkeeper/types"
)

// Store Key for cron contracts
var (
	StoreKeyCronContracts = []byte("cron_contracts")
)

// Get the store for cron contracts.
func (k Keeper) getStore(ctx sdk.Context) prefix.Store {
	return prefix.NewStore(ctx.KVStore(k.storeKey), StoreKeyCronContracts)
}

// Set a cron contract in the KV store.
func (k Keeper) SetCronContract(ctx sdk.Context, cron types.CronContract) error {
	store := k.getStore(ctx)
	bz, err := k.cdc.Marshal(&cron)
	if err != nil {
		return err
	}

	// Set the cron contract
	store.Set([]byte(cron.ContractAddress), bz)
	return nil
}

// Remove a cron contract from the KV store.
func (k Keeper) RemoveContract(ctx sdk.Context, cronContractAddress string) {
	store := k.getStore(ctx)
	key := []byte(cronContractAddress)

	if store.Has(key) {
		store.Delete(key)
	}
}

// Get a cron contract from the KV store.
func (k Keeper) GetCronContract(ctx sdk.Context, cronContractAddress string) (*types.CronContract, error) {
	if !k.IsCronContract(ctx, cronContractAddress) {
		return nil, types.ErrContractNotRegistered
	}
	store := k.getStore(ctx)
	bz := store.Get([]byte(cronContractAddress))

	var contract types.CronContract
	err := k.cdc.Unmarshal(bz, &contract)
	if err != nil {
		return nil, err
	}

	return &contract, nil
}

// Get all cron contract from the KV store.
func (k Keeper) GetAllCronContracts(ctx sdk.Context) ([]types.CronContract, error) {
	store := k.getStore(ctx)

	iterator := storetypes.KVStorePrefixIterator(store, []byte(nil))
	defer iterator.Close()

	contracts := []types.CronContract{}
	for ; iterator.Valid(); iterator.Next() {
		var contract types.CronContract
		err := k.cdc.Unmarshal(iterator.Value(), &contract)
		if err != nil {
			return nil, err
		}

		contracts = append(contracts, contract)
	}

	return contracts, nil
}

// Get all registered cron contracts
func (k Keeper) GetPaginatedCronContracts(ctx sdk.Context, pag *query.PageRequest) (*types.QueryCronContractsResponse, error) {
	store := k.getStore(ctx)

	// Filter and paginate all contracts
	results, pageRes, err := query.GenericFilteredPaginate(
		k.cdc,
		store,
		pag,
		func(_ []byte, value *types.CronContract) (*types.CronContract, error) {
			return value, nil
		},
		func() *types.CronContract {
			return &types.CronContract{}
		},
	)
	if err != nil {
		return nil, err
	}

	// Dereference pointer array of contracts
	var contracts []types.CronContract
	for _, contract := range results {
		contracts = append(contracts, *contract)
	}

	// Return paginated contracts
	return &types.QueryCronContractsResponse{
		CronContracts: contracts,
		Pagination:    pageRes,
	}, nil
}

// Register a cron contract in the KV store.
func (k Keeper) RegisterContract(ctx sdk.Context, senderAddress string, cronContractAddress string) error {
	if k.IsCronContract(ctx, cronContractAddress) {
		return types.ErrContractAlreadyRegistered
	}

	if ok, err := k.IsContractManager(ctx, senderAddress, cronContractAddress); !ok {
		return err
	}

	return k.SetCronContract(ctx, types.CronContract{
		ContractAddress: cronContractAddress,
		IsActivated:     false,
	})
}

// Unregister a cron contract from the KV store.
func (k Keeper) UnregisterContract(ctx sdk.Context, senderAddress string, cronContractAddress string) error {
	if !k.IsCronContract(ctx, cronContractAddress) {
		return types.ErrContractNotRegistered
	}

	if ok, err := k.IsContractManager(ctx, senderAddress, cronContractAddress); !ok {
		return err
	}

	k.RemoveContract(ctx, cronContractAddress)
	return nil
}

// Activate a cron contract (pending registered or deactivated).
func (k Keeper) ActivateContract(ctx sdk.Context, senderAddress string, cronContractAddress string) error {
	if !k.IsCronContract(ctx, cronContractAddress) {
		return types.ErrContractNotRegistered
	}

	cron, _ := k.GetCronContract(ctx, cronContractAddress)

	if cron.IsActivated {
		return types.ErrContractAlreadyActivated
	}

	if ok := k.IsSecurityAdmin(ctx, senderAddress); !ok {
		return types.ErrInvalidAdmin
	}

	cron.IsActivated = true
	k.SetCronContract(ctx, *cron)

	return nil
}

// Deactivate a cron contract
func (k Keeper) DeactivateCron(ctx sdk.Context, cronContractAddress string) error {
	if !k.IsCronContract(ctx, cronContractAddress) {
		return types.ErrContractNotRegistered
	}

	cron, _ := k.GetCronContract(ctx, cronContractAddress)

	if !cron.IsActivated {
		return types.ErrContractAlreadyDeactivated
	}

	cron.IsActivated = false
	k.SetCronContract(ctx, *cron)

	return nil
}

// Check if a cron contract is in the KV store.
func (k Keeper) IsCronContract(ctx sdk.Context, cronContractAddress string) bool {
	store := k.getStore(ctx)
	return store.Has([]byte(cronContractAddress))
}

// Check if the sender is the contract manager for the cron contract or not. If
// an admin is present, they are considered the manager. If there is no admin, the
// contract creator is considered the manager.
func (k Keeper) IsContractManager(ctx sdk.Context, senderAddress string, cronContractAddress string) (bool, error) {
	cron := sdk.MustAccAddressFromBech32(cronContractAddress)

	if ok := k.wasmKeeper.HasContractInfo(ctx, cron); !ok {
		return false, types.ErrInvalidCWContract
	}

	contractInfo := k.wasmKeeper.GetContractInfo(ctx, cron)

	adminExists := len(contractInfo.Admin) > 0
	isSenderAdmin := contractInfo.Admin == senderAddress
	isSenderCreator := contractInfo.Creator == senderAddress

	if adminExists && !isSenderAdmin {
		return false, types.ErrContractNotAdmin
	} else if !adminExists && !isSenderCreator {
		return false, types.ErrContractNotCreator
	}

	return true, nil
}

// Check if the sender is the designated security admin of the module
func (k Keeper) IsSecurityAdmin(ctx sdk.Context, senderAddress string) bool {
	params := k.GetParams(ctx)
	return senderAddress == params.AdminAddress

}
