package restaking_test

import (
	keepertest "github.com/lightmos/lightmos-sdk/x/restaking/ibc"
	"testing"

	"github.com/lightmos/lightmos-sdk/testutil/nullify"
	"github.com/lightmos/lightmos-sdk/x/restaking"
	"github.com/lightmos/lightmos-sdk/x/restaking/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),
		PortId: types.PortID,
		SellOrderBookList: []types.SellOrderBook{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		BuyOrderBookList: []types.BuyOrderBook{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		DenomTraceList: []types.DenomTrace{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.RestakingKeeper(t)
	restaking.InitGenesis(ctx, *k, genesisState)
	got := restaking.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.Equal(t, genesisState.PortId, got.PortId)

	require.ElementsMatch(t, genesisState.SellOrderBookList, got.SellOrderBookList)
	require.ElementsMatch(t, genesisState.BuyOrderBookList, got.BuyOrderBookList)
	require.ElementsMatch(t, genesisState.DenomTraceList, got.DenomTraceList)
	require.ElementsMatch(t, genesisState.ValidatorTokenList, got.ValidatorTokenList)
	require.Equal(t, genesisState.ValidatorTokenCount, got.ValidatorTokenCount)
	// this line is used by starport scaffolding # genesis/test/assert
}
