package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/gosuri/grumpy/testutil/keeper"
	"github.com/gosuri/grumpy/x/grumpy/keeper"
	"github.com/gosuri/grumpy/x/grumpy/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.GrumpyKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
