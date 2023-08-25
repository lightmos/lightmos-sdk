package keeper

import (
	"github.com/cosmos/cosmos-sdk/x/restaking/types"
)

var _ types.QueryServer = Keeper{}
