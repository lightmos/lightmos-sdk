package simulation

import (
	"math/rand"

	"github.com/lightmos/lightmos-sdk/baseapp"
	sdk "github.com/lightmos/lightmos-sdk/types"
	simtypes "github.com/lightmos/lightmos-sdk/types/simulation"
	"github.com/lightmos/lightmos-sdk/x/restaking/keeper"
	"github.com/lightmos/lightmos-sdk/x/restaking/types"
)

func SimulateMsgCancelSellOrder(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgCancelSellOrder{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the CancelSellOrder simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "CancelSellOrder simulation not implemented"), nil, nil
	}
}
