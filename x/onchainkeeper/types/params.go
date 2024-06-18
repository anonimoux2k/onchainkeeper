package types

import (
	errorsmod "cosmossdk.io/errors"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

var _ paramtypes.ParamSet = (*Params)(nil)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams(
	adminAddress string,
	cronGasLimit uint64,
) Params {
	return Params{
		AdminAddress: adminAddress,
		CronGasLimit: cronGasLimit,
	}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	// Using mock admin address for testing purpose
	// mnemonic: derive monitor world brick tired furnace
	//			 double label theory impose sense board
	//			 come patrol critic west predict shaft
	//			 arena over earth small fashion judge
	return NewParams("cosmos105khsf5w9x08alwcqx055nhag2xm7dcahxegda", 100_000)
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{}
}

// Validate validates the set of params
func (p Params) Validate() error {
	minimumGas := uint64(100_000)
	if p.CronGasLimit < minimumGas {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "invalid cron gas limit: %d. Must be above %d", p.CronGasLimit, minimumGas)
	}
	return nil
}
