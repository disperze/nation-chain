package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// TODO: Describe your actions, these will implment the interface of `sdk.Msg`

// verify interface at compile time
var _ sdk.Msg = &MsgRegisterDni{}

// MsgRegisterDni - struct for unjailing jailed validator
type MsgRegisterDni struct {
	Dni        string         `json:"dni" yaml:"dni"`
	Name       string         `json:"name" yaml:"name"`
	MiddleName string         `json:"middle_name" yaml:"middle_name"`
	Surname1   string         `json:"surname1" yaml:"surname1"`
	Surname2   string         `json:"surname2" yaml:"surname2"`
	Validator  sdk.ValAddress `json:"validator" yaml:"validator"` // address of the validator operator
}

// NewMsgRegisterDni creates a new MsgRegisterDni instance
func NewMsgRegisterDni(validatorAddr sdk.ValAddress) MsgRegisterDni {
	return MsgRegisterDni{
		Validator: validatorAddr,
	}
}

// RegisterDniConst value
const RegisterDniConst = "RegisterDni"

// nolint
func (msg MsgRegisterDni) Route() string { return RouterKey }
func (msg MsgRegisterDni) Type() string  { return RegisterDniConst }
func (msg MsgRegisterDni) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Validator)}
}

// GetSignBytes gets the bytes for the message signer to sign on
func (msg MsgRegisterDni) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic validity check for the AnteHandler
func (msg MsgRegisterDni) ValidateBasic() error {
	if len(msg.Dni) != 8 {
		return sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "Dni invalid length")
	}
	if len(msg.Name) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "Name cannot be empty")
	}
	if len(msg.Surname1) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "First surname cannot be empty")
	}
	if len(msg.Surname2) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "Second surname cannot be empty")
	}
	if msg.Validator.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "missing validator address")
	}
	return nil
}
