package keeper

import (
	"github.com/evan-forbes/stevia/x/stevia/types"
)

var _ types.QueryServer = Keeper{}
