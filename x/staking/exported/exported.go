package exported

import (
	sdk "github.com/lightmos/lightmos-sdk/types"
	paramtypes "github.com/lightmos/lightmos-sdk/x/params/types"
)

type (
	// Subspace defines an interface that implements the legacy x/params Subspace
	// type.
	//
	// NOTE: This is used solely for migration of x/params managed parameters.
	Subspace interface {
		GetParamSet(ctx sdk.Context, ps paramtypes.ParamSet)
	}
)
