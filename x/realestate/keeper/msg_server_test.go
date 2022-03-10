package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/pratikasr/real-estate/testutil/keeper"
	"github.com/pratikasr/real-estate/x/realestate/keeper"
	"github.com/pratikasr/real-estate/x/realestate/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.RealestateKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
