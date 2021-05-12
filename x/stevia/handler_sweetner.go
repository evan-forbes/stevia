package stevia

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/evan-forbes/stevia/x/stevia/keeper"
	"github.com/evan-forbes/stevia/x/stevia/types"
)

func handleMsgCreateSweetner(ctx sdk.Context, k keeper.Keeper, msg *types.MsgCreateSweetner) (*sdk.Result, error) {
	k.CreateSweetner(ctx, *msg)

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}

func handleMsgUpdateSweetner(ctx sdk.Context, k keeper.Keeper, msg *types.MsgUpdateSweetner) (*sdk.Result, error) {
	var sweetner = types.Sweetner{
		Creator:  msg.Creator,
		Id:       msg.Id,
		Calories: msg.Calories,
		Organic:  msg.Organic,
		Source:   msg.Source,
	}

	// Checks that the element exists
	if !k.HasSweetner(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %s doesn't exist", msg.Id))
	}

	// Checks if the the msg sender is the same as the current owner
	if msg.Creator != k.GetSweetnerOwner(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.SetSweetner(ctx, sweetner)

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}

func handleMsgDeleteSweetner(ctx sdk.Context, k keeper.Keeper, msg *types.MsgDeleteSweetner) (*sdk.Result, error) {
	if !k.HasSweetner(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %s doesn't exist", msg.Id))
	}
	if msg.Creator != k.GetSweetnerOwner(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.DeleteSweetner(ctx, msg.Id)

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}
