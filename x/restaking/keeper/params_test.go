package keeper_test

import (
	testkeeper "github.com/lightmos/lightmos-sdk/x/restaking/ibc"
	"testing"

	"github.com/lightmos/lightmos-sdk/x/restaking/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.RestakingKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
