package types

// DONTCOVER

import (
	sdkerrors "cosmossdk.io/errors"
)

// x/onchainkeeper module sentinel errors
var (
	ErrInvalidSigner             = sdkerrors.Register(ModuleName, 1100, "expected gov account as only signer for proposal message")
	ErrInvalidAdmin              = sdkerrors.Register(ModuleName, 1101, "expected security admin as only signer for this action")
	ErrInvalidCWContract         = sdkerrors.Register(ModuleName, 1102, "invalid CosmWasm contract")
	ErrContractNotAdmin          = sdkerrors.Register(ModuleName, 1103, "sender is not the contract admin")
	ErrContractNotCreator        = sdkerrors.Register(ModuleName, 1104, "sender is not the contract creator")
	ErrContractNotRegistered     = sdkerrors.Register(ModuleName, 1105, "contract not registered")
	ErrContractAlreadyRegistered = sdkerrors.Register(ModuleName, 1106, "contract already registered")
	ErrContractAlreadyActivated  = sdkerrors.Register(ModuleName, 1107, "contract already activated")
)
