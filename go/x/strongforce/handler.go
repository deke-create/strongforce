package strongforce

import (
	"fmt"

	"github.com/comrade-coop/strongforce/go/x/strongforce/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// NewHandler returns a handler for "strongforce" messages.
func NewHandler(keeper Keeper, serverAddress string) sdk.Handler {
	connection := NewConnection(serverAddress, keeper)
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		switch msg := msg.(type) {
		case types.MsgExecuteAction:
			return handleMsgExecuteAction(ctx, keeper, connection, msg)
		default:
			errMsg := fmt.Sprintf("Unrecognized strongforce Msg type: %v", msg.Type())
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}
	}
}

func handleMsgExecuteAction(ctx sdk.Context, keeper Keeper, connection Connection, msg types.MsgExecuteAction) (*sdk.Result, error) {
	actionResult := connection.SendAction(ctx, msg.Doer, msg.Action)
	if !actionResult.IsOK() {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInternal, "Couldn't execute action!")
	}
	return &sdk.Result{Events: actionResult.Events}, nil
}
