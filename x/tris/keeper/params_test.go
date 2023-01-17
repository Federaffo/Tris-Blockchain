package keeper_test

import (
	"testing"

	testkeeper "github.com/raffo/tris/testutil/keeper"
	"github.com/raffo/tris/x/tris/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.TrisKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
