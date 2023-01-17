package types

import (
	"errors"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/raffo/tris/x/tris/rules"
)

func (storedGame StoredGame) GetXAddress() (x sdk.AccAddress, err error) {
	xAddress, err := sdk.AccAddressFromBech32(storedGame.X)
	return xAddress, sdkerrors.Wrapf(err, ErrInvalidXAddress.Error(), storedGame.X)
}

func (storedGame StoredGame) GetOAddress() (o sdk.AccAddress, err error) {
	oAddress, err := sdk.AccAddressFromBech32(storedGame.X)
	return oAddress, sdkerrors.Wrapf(err, ErrInvalidOAddress.Error(), storedGame.O)
}

func (storedGame StoredGame) ParseGame() (game *rules.Game, err error) {
	board, errBoard := rules.Parse(storedGame.Board)
	if errBoard != nil {
		return nil, sdkerrors.Wrapf(errBoard, ErrGameNotParseable.Error())
	}
	board.Turn = rules.StringPieces[storedGame.Turn].Player
	if board.Turn.Symbol == "" {
		return nil, sdkerrors.Wrapf(errors.New(fmt.Sprintf("Turn: %s", storedGame.Turn)), ErrGameNotParseable.Error())
	}
	return board, nil
}

func (storedGame StoredGame) Validate() (err error) {
	_, err = storedGame.GetXAddress()
	if err != nil {
		return err
	}
	_, err = storedGame.GetOAddress()
	if err != nil {
		return err
	}
	_, err = storedGame.ParseGame()
	return err
}
