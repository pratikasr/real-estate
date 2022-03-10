package keeper

import (
	"github.com/pratikasr/real-estate/x/realestate/types"
)

var _ types.QueryServer = Keeper{}
