package grumpy_test

import (
	"testing"

	keepertest "github.com/gosuri/grumpy/testutil/keeper"
	"github.com/gosuri/grumpy/x/grumpy"
	"github.com/gosuri/grumpy/x/grumpy/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.GrumpyKeeper(t)
	grumpy.InitGenesis(ctx, *k, genesisState)
	got := grumpy.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	// this line is used by starport scaffolding # genesis/test/assert
}
