package onchainkeeper

import (
	"time"

	"cosmossdk.io/log"

	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"wasmapp/x/onchainkeeper/helpers"
	"wasmapp/x/onchainkeeper/keeper"
	"wasmapp/x/onchainkeeper/types"
)

func BeginBlocker(ctx sdk.Context, k keeper.Keeper) {
	defer telemetry.ModuleMeasureSince(types.ModuleName, time.Now(), telemetry.MetricKeyBeginBlocker)

	logger := k.Logger()
	p := k.GetParams(ctx)

	// Get all cron contracts
	crons, err := k.GetAllCronContracts(ctx)
	if err != nil {
		logger.Error("Failed to get cron contracts", "error", err)
		return
	}

	// Track errors
	errorExecs := make([]string, len(crons))
	errorExists := false

	// Execute all cron contracts that are activated
	for idx, cron := range crons {

		// Skip deactivated crons
		if !cron.IsActivated {
			continue
		}

		cronAddr := sdk.MustAccAddressFromBech32(cron.ContractAddress)
		if handleError(ctx, k, logger, errorExecs, &errorExists, err, idx, cron.ContractAddress) {
			continue
		}

		// Create context with gas limit from params
		childCtx := ctx.WithGasMeter(storetypes.NewGasMeter(p.CronGasLimit))

		// Execute cron
		helpers.ExecuteContract(k.GetContractKeeper(), childCtx, cronAddr, []byte(types.BeginBlockSudoMessage), &err)
		if handleError(ctx, k, logger, errorExecs, &errorExists, err, idx, cron.ContractAddress) {
			continue
		}
	}

	// Log errors
	if errorExists {
		logger.Error("Failed to execute cron contracts at the beginning of the block", "crons", errorExecs)
	}
}

func EndBlocker(ctx sdk.Context, k keeper.Keeper) {
	defer telemetry.ModuleMeasureSince(types.ModuleName, time.Now(), telemetry.MetricKeyEndBlocker)

	logger := k.Logger()
	p := k.GetParams(ctx)

	// Get all cron contracts
	crons, err := k.GetAllCronContracts(ctx)
	if err != nil {
		logger.Error("Failed to get cron contracts", "error", err)
		return
	}

	// Track errors
	errorExecs := make([]string, len(crons))
	errorExists := false

	// Execute all cron contracts that are activated
	for idx, cron := range crons {

		// Skip deactivated crons
		if !cron.IsActivated {
			continue
		}

		cronAddr := sdk.MustAccAddressFromBech32(cron.ContractAddress)
		if handleError(ctx, k, logger, errorExecs, &errorExists, err, idx, cron.ContractAddress) {
			continue
		}

		// Create context with gas limit from params
		childCtx := ctx.WithGasMeter(storetypes.NewGasMeter(p.CronGasLimit))

		// Execute cron
		helpers.ExecuteContract(k.GetContractKeeper(), childCtx, cronAddr, []byte(types.EndBlockSudoMessage), &err)
		if handleError(ctx, k, logger, errorExecs, &errorExists, err, idx, cron.ContractAddress) {
			continue
		}
	}

	// Log errors
	if errorExists {
		logger.Error("Failed to execute cron contracts at the end of block", "crons", errorExecs)
	}
}

// Handle contract execution errors. Return true if error is present, false otherwise.
func handleError(
	ctx sdk.Context,
	k keeper.Keeper,
	logger log.Logger,
	errorExecs []string,
	errorExists *bool,
	err error,
	idx int,
	cronAddress string,
) bool {
	if err != nil {
		*errorExists = true
		errorExecs[idx] = cronAddress

		err := k.DeactivateCron(ctx, cronAddress)
		if err != nil {
			logger.Error("Failed to deactivate cron", "cron", cronAddress, "error", err)
		}
	}

	return false
}
