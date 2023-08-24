package ibc

import (
	"testing"

	"github.com/lightmos/lightmos-sdk/x/restaking/keeper"
	"github.com/lightmos/lightmos-sdk/x/restaking/types"

	tmdb "github.com/cometbft/cometbft-db"
	"github.com/cometbft/cometbft/libs/log"
	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	clienttypes "github.com/cosmos/ibc-go/v7/modules/core/02-client/types"
	channeltypes "github.com/cosmos/ibc-go/v7/modules/core/04-channel/types"
	"github.com/lightmos/lightmos-sdk/codec"
	codectypes "github.com/lightmos/lightmos-sdk/codec/types"
	"github.com/lightmos/lightmos-sdk/store"
	storetypes "github.com/lightmos/lightmos-sdk/store/types"
	sdk "github.com/lightmos/lightmos-sdk/types"
	capabilitykeeper "github.com/lightmos/lightmos-sdk/x/capability/keeper"
	capabilitytypes "github.com/lightmos/lightmos-sdk/x/capability/types"
	typesparams "github.com/lightmos/lightmos-sdk/x/params/types"
	"github.com/stretchr/testify/require"
)

// restakingChannelKeeper is a stub of cosmosibckeeper.ChannelKeeper.
type restakingChannelKeeper struct{}

func (restakingChannelKeeper) GetChannel(ctx sdk.Context, portID, channelID string) (channeltypes.Channel, bool) {
	return channeltypes.Channel{}, false
}

func (restakingChannelKeeper) GetNextSequenceSend(ctx sdk.Context, portID, channelID string) (uint64, bool) {
	return 0, false
}

func (restakingChannelKeeper) SendPacket(
	ctx sdk.Context,
	channelCap *capabilitytypes.Capability,
	sourcePort string,
	sourceChannel string,
	timeoutHeight clienttypes.Height,
	timeoutTimestamp uint64,
	data []byte,
) (uint64, error) {
	return 0, nil
}

func (restakingChannelKeeper) ChanCloseInit(ctx sdk.Context, portID, channelID string, chanCap *capabilitytypes.Capability) error {
	return nil
}

// restakingportKeeper is a stub of cosmosibckeeper.PortKeeper
type restakingPortKeeper struct{}

func (restakingPortKeeper) BindPort(ctx sdk.Context, portID string) *capabilitytypes.Capability {
	return &capabilitytypes.Capability{}
}

func RestakingKeeper(t testing.TB) (*keeper.Keeper, sdk.Context) {
	logger := log.NewNopLogger()

	storeKey := sdk.NewKVStoreKey(types.StoreKey)
	memStoreKey := storetypes.NewMemoryStoreKey(types.MemStoreKey)

	db := tmdb.NewMemDB()
	stateStore := store.NewCommitMultiStore(db)
	stateStore.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(memStoreKey, storetypes.StoreTypeMemory, nil)
	require.NoError(t, stateStore.LoadLatestVersion())

	registry := codectypes.NewInterfaceRegistry()
	appCodec := codec.NewProtoCodec(registry)
	capabilityKeeper := capabilitykeeper.NewKeeper(appCodec, storeKey, memStoreKey)

	paramsSubspace := typesparams.NewSubspace(appCodec,
		types.Amino,
		storeKey,
		memStoreKey,
		"RestakingParams",
	)
	k := keeper.NewKeeper(
		appCodec,
		storeKey,
		memStoreKey,
		paramsSubspace,
		restakingChannelKeeper{},
		restakingPortKeeper{},
		capabilityKeeper.ScopeToModule("RestakingScopedKeeper"),
		nil,
		nil,
		nil,
	)

	ctx := sdk.NewContext(stateStore, tmproto.Header{}, false, logger)

	// Initialize params
	k.SetParams(ctx, types.DefaultParams())

	return k, ctx
}
