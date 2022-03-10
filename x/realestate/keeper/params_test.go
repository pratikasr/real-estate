package keeper_test

import (
	"testing"

	testkeeper "github.com/pratikasr/real-estate/testutil/keeper"
	"github.com/pratikasr/real-estate/x/realestate/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.RealestateKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
