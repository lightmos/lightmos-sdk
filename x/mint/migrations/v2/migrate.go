package v2

import (
	"github.com/lightmos/lightmos-sdk/codec"
	sdk "github.com/lightmos/lightmos-sdk/types"
	"github.com/lightmos/lightmos-sdk/x/mint/exported"
	"github.com/lightmos/lightmos-sdk/x/mint/types"
)

const (
	ModuleName = "mint"
)

var ParamsKey = []byte{0x01}

// Migrate migrates the x/mint module state from the consensus version 1 to
// version 2. Specifically, it takes the parameters that are currently stored
// and managed by the x/params modules and stores them directly into the x/mint
// module state.
func Migrate(
	ctx sdk.Context,
	store sdk.KVStore,
	legacySubspace exported.Subspace,
	cdc codec.BinaryCodec,
) error {
	var currParams types.Params
	legacySubspace.GetParamSet(ctx, &currParams)

	if err := currParams.Validate(); err != nil {
		return err
	}

	bz := cdc.MustMarshal(&currParams)
	store.Set(ParamsKey, bz)

	return nil
}
