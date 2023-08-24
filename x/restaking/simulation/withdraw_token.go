package simulation

import (
	"math/rand"

	"github.com/lightmos/lightmos-sdk/baseapp"
	sdk "github.com/lightmos/lightmos-sdk/types"
	simtypes "github.com/lightmos/lightmos-sdk/types/simulation"
	"lightmos/x/restaking/keeper"
	"lightmos/x/restaking/types"
)

func SimulateMsgWithdrawToken(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgWithdrawToken{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the WithdrawToken simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "WithdrawToken simulation not implemented"), nil, nil
	}
}
