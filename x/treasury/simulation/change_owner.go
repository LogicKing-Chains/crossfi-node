package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/mineplexio/mineplex-2-node/x/treasury/keeper"
	"github.com/mineplexio/mineplex-2-node/x/treasury/types"
)

func SimulateMsgChangeOwner(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgChangeOwner{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the ChangeOwner simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "ChangeOwner simulation not implemented"), nil, nil
	}
}
