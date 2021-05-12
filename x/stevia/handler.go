package stevia

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/evan-forbes/stevia/x/stevia/keeper"
	"github.com/evan-forbes/stevia/x/stevia/types"
)

// NewHandler ...
func NewHandler(k keeper.Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())

		switch msg := msg.(type) {
		// this line is used by starport scaffolding # 1
		case *types.MsgCreateSweetner:
			return handleMsgCreateSweetner(ctx, k, msg)

		case *types.MsgUpdateSweetner:
			return handleMsgUpdateSweetner(ctx, k, msg)

		case *types.MsgDeleteSweetner:
			return handleMsgDeleteSweetner(ctx, k, msg)

		default:
			errMsg := fmt.Sprintf("unrecognized %s message type: %T", types.ModuleName, msg)
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}
	}
}
