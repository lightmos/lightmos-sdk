package keeper_test

import (
	"testing"
	"time"

	"github.com/cometbft/cometbft/libs/log"
	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	"github.com/stretchr/testify/suite"

	"github.com/golang/mock/gomock"
	"github.com/lightmos/lightmos-sdk/baseapp"
	"github.com/lightmos/lightmos-sdk/crypto/keys/secp256k1"
	"github.com/lightmos/lightmos-sdk/testutil"
	sdk "github.com/lightmos/lightmos-sdk/types"
	moduletestutil "github.com/lightmos/lightmos-sdk/types/module/testutil"
	"github.com/lightmos/lightmos-sdk/x/authz/keeper"
	authzmodule "github.com/lightmos/lightmos-sdk/x/authz/module"
	authztestutil "github.com/lightmos/lightmos-sdk/x/authz/testutil"
	bank "github.com/lightmos/lightmos-sdk/x/bank/types"
)

type GenesisTestSuite struct {
	suite.Suite

	ctx           sdk.Context
	keeper        keeper.Keeper
	baseApp       *baseapp.BaseApp
	accountKeeper *authztestutil.MockAccountKeeper
	encCfg        moduletestutil.TestEncodingConfig
}

func (suite *GenesisTestSuite) SetupTest() {
	key := sdk.NewKVStoreKey(keeper.StoreKey)
	testCtx := testutil.DefaultContextWithDB(suite.T(), key, sdk.NewTransientStoreKey("transient_test"))
	suite.ctx = testCtx.Ctx.WithBlockHeader(tmproto.Header{Height: 1})
	suite.encCfg = moduletestutil.MakeTestEncodingConfig(authzmodule.AppModuleBasic{})

	// gomock initializations
	ctrl := gomock.NewController(suite.T())
	suite.accountKeeper = authztestutil.NewMockAccountKeeper(ctrl)

	suite.baseApp = baseapp.NewBaseApp(
		"authz",
		log.NewNopLogger(),
		testCtx.DB,
		suite.encCfg.TxConfig.TxDecoder(),
	)
	suite.baseApp.SetCMS(testCtx.CMS)

	bank.RegisterInterfaces(suite.encCfg.InterfaceRegistry)

	msr := suite.baseApp.MsgServiceRouter()
	msr.SetInterfaceRegistry(suite.encCfg.InterfaceRegistry)

	suite.keeper = keeper.NewKeeper(key, suite.encCfg.Codec, msr, suite.accountKeeper)
}

var (
	granteePub  = secp256k1.GenPrivKey().PubKey()
	granterPub  = secp256k1.GenPrivKey().PubKey()
	granteeAddr = sdk.AccAddress(granteePub.Address())
	granterAddr = sdk.AccAddress(granterPub.Address())
)

func (suite *GenesisTestSuite) TestImportExportGenesis() {
	coins := sdk.NewCoins(sdk.NewCoin("foo", sdk.NewInt(1_000)))

	now := suite.ctx.BlockTime()
	expires := now.Add(time.Hour)
	grant := &bank.SendAuthorization{SpendLimit: coins}
	err := suite.keeper.SaveGrant(suite.ctx, granteeAddr, granterAddr, grant, &expires)
	suite.Require().NoError(err)
	genesis := suite.keeper.ExportGenesis(suite.ctx)

	// TODO, recheck!
	// Clear keeper
	suite.keeper.DeleteGrant(suite.ctx, granteeAddr, granterAddr, grant.MsgTypeURL())

	suite.keeper.InitGenesis(suite.ctx, genesis)
	newGenesis := suite.keeper.ExportGenesis(suite.ctx)
	suite.Require().Equal(genesis, newGenesis)
}

func TestGenesisTestSuite(t *testing.T) {
	suite.Run(t, new(GenesisTestSuite))
}
