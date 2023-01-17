package types_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/raffo/tris/x/tris/rules"
	"github.com/raffo/tris/x/tris/testutil"
	"github.com/raffo/tris/x/tris/types"
	"github.com/stretchr/testify/require"
)

const (
	alice = testutil.Alice
	bob   = testutil.Bob
)

func GetStoredGame1() types.StoredGame {
	return types.StoredGame{
		X:     alice,
		O:     bob,
		Index: "1",
		Board: rules.New().String(),
		Turn:  "x",
	}
}

func TestCanGetAddressX(t *testing.T) {
	aliceAddress, err1 := sdk.AccAddressFromBech32(alice)
	xAddress, err2 := GetStoredGame1().GetXAddress()
	require.Equal(t, aliceAddress, xAddress)
	require.Nil(t, err1)
	require.Nil(t, err2)
}

func TestWrongAddressX(t *testing.T) {
	storedGame := GetStoredGame1()
	storedGame.X = "wrongAddress"
	xAddress, _ := storedGame.GetXAddress()
	require.Nil(t, xAddress)
}

func TestParseGameCorrect(t *testing.T) {
	game, err := GetStoredGame1().ParseGame()
	require.EqualValues(t, rules.New().Pieces, game.Pieces)
	require.Nil(t, err)
}

func TestParseGameWrongTurnColor(t *testing.T) {
	storedGame := GetStoredGame1()
	storedGame.Turn = "w"
	game, err := storedGame.ParseGame()
	require.Nil(t, game)
	require.EqualError(t, err, "game cannot be parsed: Turn: w")
	require.EqualError(t, storedGame.Validate(), err.Error())
}
