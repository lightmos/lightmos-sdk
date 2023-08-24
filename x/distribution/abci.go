package distribution

import (
	"time"

	abci "github.com/cometbft/cometbft/abci/types"

	"github.com/lightmos/lightmos-sdk/telemetry"
	sdk "github.com/lightmos/lightmos-sdk/types"
	"github.com/lightmos/lightmos-sdk/x/distribution/keeper"
	"github.com/lightmos/lightmos-sdk/x/distribution/types"
)

// BeginBlocker sets the proposer for determining distribution during endblock
// and distribute rewards for the previous block.
func BeginBlocker(ctx sdk.Context, req abci.RequestBeginBlock, k keeper.Keeper) {
	defer telemetry.ModuleMeasureSince(types.ModuleName, time.Now(), telemetry.MetricKeyBeginBlocker)

	// determine the total power signing the block
	var previousTotalPower int64
	for _, voteInfo := range req.LastCommitInfo.GetVotes() {
		previousTotalPower += voteInfo.Validator.Power
	}

	// TODO this is Tendermint-dependent
	// ref https://github.com/lightmos/lightmos-sdk/issues/3095
	if ctx.BlockHeight() > 1 {
		k.AllocateTokens(ctx, previousTotalPower, req.LastCommitInfo.GetVotes())
	}

	// record the proposer for when we payout on the next block
	consAddr := sdk.ConsAddress(req.Header.ProposerAddress)
	k.SetPreviousProposerConsAddr(ctx, consAddr)
}
