package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	// this line is used by starport scaffolding # 2
	cdc.RegisterConcrete(&MsgCreateSweetner{}, "stevia/CreateSweetner", nil)
	cdc.RegisterConcrete(&MsgUpdateSweetner{}, "stevia/UpdateSweetner", nil)
	cdc.RegisterConcrete(&MsgDeleteSweetner{}, "stevia/DeleteSweetner", nil)

}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	// this line is used by starport scaffolding # 3
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateSweetner{},
		&MsgUpdateSweetner{},
		&MsgDeleteSweetner{},
	)
}

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewAminoCodec(amino)
)
