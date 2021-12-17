package keeper

import (
	"github.com/gosuri/grumpy/x/grumpy/types"
)

var _ types.QueryServer = Keeper{}
