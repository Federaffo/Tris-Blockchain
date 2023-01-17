package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/tris module sentinel errors
var (
	ErrSample = sdkerrors.Register(ModuleName, 1100, "sample error")
)
var (
	ErrInvalidXAddress  = sdkerrors.Register(ModuleName, 1101, "black address is invalid: %s")
	ErrInvalidOAddress  = sdkerrors.Register(ModuleName, 1102, "red address is invalid: %s")
	ErrGameNotParseable = sdkerrors.Register(ModuleName, 1103, "game cannot be parsed")
)
