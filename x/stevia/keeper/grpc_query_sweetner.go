package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/evan-forbes/stevia/x/stevia/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) SweetnerAll(c context.Context, req *types.QueryAllSweetnerRequest) (*types.QueryAllSweetnerResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var sweetners []*types.Sweetner
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	sweetnerStore := prefix.NewStore(store, types.KeyPrefix(types.SweetnerKey))

	pageRes, err := query.Paginate(sweetnerStore, req.Pagination, func(key []byte, value []byte) error {
		var sweetner types.Sweetner
		if err := k.cdc.UnmarshalBinaryBare(value, &sweetner); err != nil {
			return err
		}

		sweetners = append(sweetners, &sweetner)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllSweetnerResponse{Sweetner: sweetners, Pagination: pageRes}, nil
}

func (k Keeper) Sweetner(c context.Context, req *types.QueryGetSweetnerRequest) (*types.QueryGetSweetnerResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var sweetner types.Sweetner
	ctx := sdk.UnwrapSDKContext(c)

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SweetnerKey))
	k.cdc.MustUnmarshalBinaryBare(store.Get(types.KeyPrefix(types.SweetnerKey+req.Id)), &sweetner)

	return &types.QueryGetSweetnerResponse{Sweetner: &sweetner}, nil
}
