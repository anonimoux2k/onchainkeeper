package keeper

import (
	"wasmapp/x/onchainkeeper/types"
)

var _ types.QueryServer = Keeper{}
