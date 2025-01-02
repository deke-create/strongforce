package types

import (
	"encoding/json"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// MsgExecuteAction represents a request to execute an action
type MsgExecuteAction struct {
	Doer   sdk.AccAddress `json:"doer" yaml:"doer"`
	Action []byte         `json:"action" yaml:"action"`
}

// NewMsgExecuteAction is the constructor for MsgExecuteAction
func NewMsgExecuteAction(doer sdk.AccAddress, action []byte) MsgExecuteAction {
	return MsgExecuteAction{
		Doer:   doer,
		Action: action,
	}
}

// Route - Implements sdk.Msg
func (msg MsgExecuteAction) Route() string { return RouterKey }

// Type - Implements sdk.Msg
func (msg MsgExecuteAction) Type() string { return "execute_action" }

// ValidateBasic  - Implements sdk.Msg
func (msg MsgExecuteAction) ValidateBasic() error {
	if msg.Doer.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Doer.String())
	}
	return nil
}

// GetSignBytes  - Implements sdk.Msg
func (msg MsgExecuteAction) GetSignBytes() []byte {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(b)
}

// GetSigners  - Implements sdk.Msg
func (msg MsgExecuteAction) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Doer}
}
