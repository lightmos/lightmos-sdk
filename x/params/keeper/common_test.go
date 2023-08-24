package keeper_test

import (
	"cosmossdk.io/depinject"
	"github.com/lightmos/lightmos-sdk/codec"
	storetypes "github.com/lightmos/lightmos-sdk/store/types"
	sdktestutil "github.com/lightmos/lightmos-sdk/testutil"
	sdk "github.com/lightmos/lightmos-sdk/types"
	paramskeeper "github.com/lightmos/lightmos-sdk/x/params/keeper"
	"github.com/lightmos/lightmos-sdk/x/params/testutil"
)

func testComponents() (*codec.LegacyAmino, sdk.Context, storetypes.StoreKey, storetypes.StoreKey, paramskeeper.Keeper) {
	var cdc codec.Codec
	if err := depinject.Inject(testutil.AppConfig, &cdc); err != nil {
		panic(err)
	}

	legacyAmino := createTestCodec()
	mkey := sdk.NewKVStoreKey("test")
	tkey := sdk.NewTransientStoreKey("transient_test")
	ctx := sdktestutil.DefaultContext(mkey, tkey)
	keeper := paramskeeper.NewKeeper(cdc, legacyAmino, mkey, tkey)

	return legacyAmino, ctx, mkey, tkey, keeper
}

type invalid struct{}

type s struct {
	I int
}

func createTestCodec() *codec.LegacyAmino {
	cdc := codec.NewLegacyAmino()
	sdk.RegisterLegacyAminoCodec(cdc)
	cdc.RegisterConcrete(s{}, "test/s", nil)
	cdc.RegisterConcrete(invalid{}, "test/invalid", nil)
	return cdc
}
