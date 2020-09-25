package nation

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/disperze/nation-chain/x/nation/types"
)

// NewHandler creates an sdk.Handler for all the nation type messages
func NewHandler(k Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())
		switch msg := msg.(type) {
		case MsgRegisterDni:
			return handleMsgRegisterDni(ctx, k, msg)
		default:
			errMsg := fmt.Sprintf("unrecognized %s message type: %T", ModuleName, msg)
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}
	}
}

// handleRegisterDni does x
func handleMsgRegisterDni(ctx sdk.Context, k Keeper, msg MsgRegisterDni) (*sdk.Result, error) {
	person := types.NewPerson()
	person.Name = msg.Name
	err := k.RegisterDni(ctx, msg.Dni, person)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeyAction, types.EventTypeRegisterDni),
			sdk.NewAttribute(types.AttributeDni, msg.Dni),
		),
	)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
