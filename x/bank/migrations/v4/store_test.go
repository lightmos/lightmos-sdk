package v4_test

import (
	"testing"

	"github.com/lightmos/lightmos-sdk/testutil"
	sdk "github.com/lightmos/lightmos-sdk/types"
	moduletestutil "github.com/lightmos/lightmos-sdk/types/module/testutil"
	"github.com/lightmos/lightmos-sdk/x/bank"
	"github.com/lightmos/lightmos-sdk/x/bank/exported"
	v4 "github.com/lightmos/lightmos-sdk/x/bank/migrations/v4"
	"github.com/lightmos/lightmos-sdk/x/bank/types"
	"github.com/stretchr/testify/require"
)

type mockSubspace struct {
	ps types.Params
}

func newMockSubspace(ps types.Params) mockSubspace {
	return mockSubspace{ps: ps}
}

func (ms mockSubspace) GetParamSet(ctx sdk.Context, ps exported.ParamSet) {
	*ps.(*types.Params) = ms.ps
}

func TestMigrate(t *testing.T) {
	encCfg := moduletestutil.MakeTestEncodingConfig(bank.AppModuleBasic{})
	cdc := encCfg.Codec

	storeKey := sdk.NewKVStoreKey(v4.ModuleName)
	tKey := sdk.NewTransientStoreKey("transient_test")
	ctx := testutil.DefaultContext(storeKey, tKey)
	store := ctx.KVStore(storeKey)

	legacySubspace := newMockSubspace(types.DefaultParams())
	require.NoError(t, v4.MigrateStore(ctx, storeKey, legacySubspace, cdc))

	var res types.Params
	bz := store.Get(v4.ParamsKey)
	require.NoError(t, cdc.Unmarshal(bz, &res))
	require.Equal(t, legacySubspace.ps, res)
}
