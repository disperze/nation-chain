package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// Custom errors
var (
	ErrExistValue = sdkerrors.Register(ModuleName, 1, "value already exist")
)
