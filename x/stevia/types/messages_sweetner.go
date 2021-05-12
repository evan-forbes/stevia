package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateSweetner{}

func NewMsgCreateSweetner(creator string, calories uint64, organic bool, source string) *MsgCreateSweetner {
	return &MsgCreateSweetner{
		Creator:  creator,
		Calories: calories,
		Organic:  organic,
		Source:   source,
	}
}

func (msg *MsgCreateSweetner) Route() string {
	return RouterKey
}

func (msg *MsgCreateSweetner) Type() string {
	return "CreateSweetner"
}

func (msg *MsgCreateSweetner) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateSweetner) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateSweetner) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateSweetner{}

func NewMsgUpdateSweetner(creator string, id string, calories uint64, organic bool, source string) *MsgUpdateSweetner {
	return &MsgUpdateSweetner{
		Id:       id,
		Creator:  creator,
		Calories: calories,
		Organic:  organic,
		Source:   source,
	}
}

func (msg *MsgUpdateSweetner) Route() string {
	return RouterKey
}

func (msg *MsgUpdateSweetner) Type() string {
	return "UpdateSweetner"
}

func (msg *MsgUpdateSweetner) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateSweetner) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateSweetner) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgCreateSweetner{}

func NewMsgDeleteSweetner(creator string, id string) *MsgDeleteSweetner {
	return &MsgDeleteSweetner{
		Id:      id,
		Creator: creator,
	}
}
func (msg *MsgDeleteSweetner) Route() string {
	return RouterKey
}

func (msg *MsgDeleteSweetner) Type() string {
	return "DeleteSweetner"
}

func (msg *MsgDeleteSweetner) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteSweetner) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteSweetner) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
