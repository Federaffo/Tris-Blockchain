package keeper_test

import (
	"testing"

	"github.com/raffo/tris/x/tris/types"
	"github.com/stretchr/testify/require"
)

func TestCreateGame(t *testing.T) {
	msgServer, context := setupMsgServer(t)
	response, err := msgServer.CreateGame(context, &types.MsgCreateGame{
		Creator: alice,
		X:       bob,
		O:       carol,
	})
	require.Nil(t, err)
	require.EqualValues(t, types.MsgCreateGameResponse{
		GameIndex: "",
	}, *response)
}
