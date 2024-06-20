package helpers

import (
	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"

	"wasmapp/x/onchainkeeper/types"

	storetypes "cosmossdk.io/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Execute contract, recover from panic (return error)
func ExecuteContract(k wasmtypes.ContractOpsKeeper, ctx sdk.Context, contractAddr sdk.AccAddress, msgBz []byte, err *error) {
	defer func() {
		if recoverError := recover(); recoverError != nil {
			if isOutOfGas, msg := IsOutOfGasError(recoverError); isOutOfGas {
				*err = types.ErrOutOfGas.Wrapf("%s", msg)
			} else {
				*err = types.ErrContractExecutionPanic.Wrapf("%s", recoverError)
			}
		}
	}()

	// Execute contract with sudo
	_, *err = k.Sudo(ctx, contractAddr, msgBz)
}

// Check if error is out of gas error
func IsOutOfGasError(err any) (bool, string) {
	switch e := err.(type) {
	case storetypes.ErrorOutOfGas:
		return true, e.Descriptor
	case storetypes.ErrorGasOverflow:
		return true, e.Descriptor
	default:
		return false, ""
	}
}
