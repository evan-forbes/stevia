package stevia

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/evan-forbes/stevia/x/stevia/keeper"
	"github.com/evan-forbes/stevia/x/stevia/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// this line is used by starport scaffolding # genesis/module/init
	// Set all the sweetner
	for _, elem := range genState.SweetnerList {
		k.SetSweetner(ctx, *elem)
	}

	// Set sweetner count
	k.SetSweetnerCount(ctx, int64(len(genState.SweetnerList)))

}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()

	// this line is used by starport scaffolding # genesis/module/export
	// Get all sweetner
	sweetnerList := k.GetAllSweetner(ctx)
	for _, elem := range sweetnerList {
		elem := elem
		genesis.SweetnerList = append(genesis.SweetnerList, &elem)
	}

	return genesis
}
