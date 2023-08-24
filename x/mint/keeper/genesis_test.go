package keeper_test

import (
	"testing"

	"cosmossdk.io/math"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"

	"github.com/lightmos/lightmos-sdk/codec"
	storetypes "github.com/lightmos/lightmos-sdk/store/types"
	"github.com/lightmos/lightmos-sdk/testutil"
	sdk "github.com/lightmos/lightmos-sdk/types"
	moduletestutil "github.com/lightmos/lightmos-sdk/types/module/testutil"
	authtypes "github.com/lightmos/lightmos-sdk/x/auth/types"
	"github.com/lightmos/lightmos-sdk/x/mint"
	"github.com/lightmos/lightmos-sdk/x/mint/keeper"
	minttestutil "github.com/lightmos/lightmos-sdk/x/mint/testutil"
	"github.com/lightmos/lightmos-sdk/x/mint/types"
)

var minterAcc = authtypes.NewEmptyModuleAccount(types.ModuleName, authtypes.Minter)

type GenesisTestSuite struct {
	suite.Suite

	sdkCtx        sdk.Context
	keeper        keeper.Keeper
	cdc           codec.BinaryCodec
	accountKeeper types.AccountKeeper
	key           *storetypes.KVStoreKey
}

func TestGenesisTestSuite(t *testing.T) {
	suite.Run(t, new(GenesisTestSuite))
}

func (s *GenesisTestSuite) SetupTest() {
	key := sdk.NewKVStoreKey(types.StoreKey)
	testCtx := testutil.DefaultContextWithDB(s.T(), key, sdk.NewTransientStoreKey("transient_test"))
	encCfg := moduletestutil.MakeTestEncodingConfig(mint.AppModuleBasic{})

	// gomock initializations
	ctrl := gomock.NewController(s.T())
	s.cdc = codec.NewProtoCodec(encCfg.InterfaceRegistry)
	s.sdkCtx = testCtx.Ctx
	s.key = key

	stakingKeeper := minttestutil.NewMockStakingKeeper(ctrl)
	accountKeeper := minttestutil.NewMockAccountKeeper(ctrl)
	bankKeeper := minttestutil.NewMockBankKeeper(ctrl)
	s.accountKeeper = accountKeeper
	accountKeeper.EXPECT().GetModuleAddress(minterAcc.Name).Return(minterAcc.GetAddress())
	accountKeeper.EXPECT().GetModuleAccount(s.sdkCtx, minterAcc.Name).Return(minterAcc)

	s.keeper = keeper.NewKeeper(s.cdc, key, stakingKeeper, accountKeeper, bankKeeper, "", "")
}

func (s *GenesisTestSuite) TestImportExportGenesis() {
	genesisState := types.DefaultGenesisState()
	genesisState.Minter = types.NewMinter(sdk.NewDecWithPrec(20, 2), math.LegacyNewDec(1))
	genesisState.Params = types.NewParams(
		"testDenom",
		sdk.NewDecWithPrec(15, 2),
		sdk.NewDecWithPrec(22, 2),
		sdk.NewDecWithPrec(9, 2),
		sdk.NewDecWithPrec(69, 2),
		uint64(60*60*8766/5),
	)

	s.keeper.InitGenesis(s.sdkCtx, s.accountKeeper, genesisState)

	minter := s.keeper.GetMinter(s.sdkCtx)
	s.Require().Equal(genesisState.Minter, minter)

	invalidCtx := testutil.DefaultContextWithDB(s.T(), s.key, sdk.NewTransientStoreKey("transient_test"))
	s.Require().Panics(func() { s.keeper.GetMinter(invalidCtx.Ctx) }, "stored minter should not have been nil")
	params := s.keeper.GetParams(s.sdkCtx)
	s.Require().Equal(genesisState.Params, params)

	genesisState2 := s.keeper.ExportGenesis(s.sdkCtx)
	s.Require().Equal(genesisState, genesisState2)
}
