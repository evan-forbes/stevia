package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/evan-forbes/stevia/x/stevia/types"
	"strconv"
)

// GetSweetnerCount get the total number of sweetner
func (k Keeper) GetSweetnerCount(ctx sdk.Context) int64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SweetnerCountKey))
	byteKey := types.KeyPrefix(types.SweetnerCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	count, err := strconv.ParseInt(string(bz), 10, 64)
	if err != nil {
		// Panic because the count should be always formattable to int64
		panic("cannot decode count")
	}

	return count
}

// SetSweetnerCount set the total number of sweetner
func (k Keeper) SetSweetnerCount(ctx sdk.Context, count int64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SweetnerCountKey))
	byteKey := types.KeyPrefix(types.SweetnerCountKey)
	bz := []byte(strconv.FormatInt(count, 10))
	store.Set(byteKey, bz)
}

// CreateSweetner creates a sweetner with a new id and update the count
func (k Keeper) CreateSweetner(ctx sdk.Context, msg types.MsgCreateSweetner) {
	// Create the sweetner
	count := k.GetSweetnerCount(ctx)
	var sweetner = types.Sweetner{
		Creator:  msg.Creator,
		Id:       strconv.FormatInt(count, 10),
		Calories: msg.Calories,
		Organic:  msg.Organic,
		Source:   msg.Source,
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SweetnerKey))
	key := types.KeyPrefix(types.SweetnerKey + sweetner.Id)
	value := k.cdc.MustMarshalBinaryBare(&sweetner)
	store.Set(key, value)

	// Update sweetner count
	k.SetSweetnerCount(ctx, count+1)
}

// SetSweetner set a specific sweetner in the store
func (k Keeper) SetSweetner(ctx sdk.Context, sweetner types.Sweetner) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SweetnerKey))
	b := k.cdc.MustMarshalBinaryBare(&sweetner)
	store.Set(types.KeyPrefix(types.SweetnerKey+sweetner.Id), b)
}

// GetSweetner returns a sweetner from its id
func (k Keeper) GetSweetner(ctx sdk.Context, key string) types.Sweetner {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SweetnerKey))
	var sweetner types.Sweetner
	k.cdc.MustUnmarshalBinaryBare(store.Get(types.KeyPrefix(types.SweetnerKey+key)), &sweetner)
	return sweetner
}

// HasSweetner checks if the sweetner exists
func (k Keeper) HasSweetner(ctx sdk.Context, id string) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SweetnerKey))
	return store.Has(types.KeyPrefix(types.SweetnerKey + id))
}

// GetSweetnerOwner returns the creator of the sweetner
func (k Keeper) GetSweetnerOwner(ctx sdk.Context, key string) string {
	return k.GetSweetner(ctx, key).Creator
}

// DeleteSweetner deletes a sweetner
func (k Keeper) DeleteSweetner(ctx sdk.Context, key string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SweetnerKey))
	store.Delete(types.KeyPrefix(types.SweetnerKey + key))
}

// GetAllSweetner returns all sweetner
func (k Keeper) GetAllSweetner(ctx sdk.Context) (msgs []types.Sweetner) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SweetnerKey))
	iterator := sdk.KVStorePrefixIterator(store, types.KeyPrefix(types.SweetnerKey))

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var msg types.Sweetner
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &msg)
		msgs = append(msgs, msg)
	}

	return
}
