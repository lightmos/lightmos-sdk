package distribution_test

import (
	"testing"

	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	"github.com/stretchr/testify/require"

	simtestutil "github.com/lightmos/lightmos-sdk/testutil/sims"
	authkeeper "github.com/lightmos/lightmos-sdk/x/auth/keeper"
	authtypes "github.com/lightmos/lightmos-sdk/x/auth/types"
	"github.com/lightmos/lightmos-sdk/x/distribution/testutil"
	"github.com/lightmos/lightmos-sdk/x/distribution/types"
)

func TestItCreatesModuleAccountOnInitBlock(t *testing.T) {
	var accountKeeper authkeeper.AccountKeeper

	app, err := simtestutil.SetupAtGenesis(testutil.AppConfig, &accountKeeper)
	require.NoError(t, err)

	ctx := app.BaseApp.NewContext(false, tmproto.Header{})
	acc := accountKeeper.GetAccount(ctx, authtypes.NewModuleAddress(types.ModuleName))
	require.NotNil(t, acc)
}
