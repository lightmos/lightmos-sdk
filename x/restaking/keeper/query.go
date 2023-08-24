package keeper

import (
	"github.com/lightmos/lightmos-sdk/x/restaking/types"
)

var _ types.QueryServer = Keeper{}
