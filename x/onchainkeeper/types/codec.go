package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	// this line is used by starport scaffolding # 1
)

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)

func init() {
	RegisterCodec(amino)
	sdk.RegisterLegacyAminoCodec(amino)
}

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgRegisterCronContract{}, "onchainkeeper/MsgRegisterCronContract", nil)
	cdc.RegisterConcrete(&MsgUnregisterCronContract{}, "onchainkeeper/MsgUnregisterCronContract", nil)
	cdc.RegisterConcrete(&MsgAcceptPendingRegisteredCronContract{}, "onchainkeeper/MsgAcceptPendingRegisteredCronContract", nil)
	cdc.RegisterConcrete(&MsgReactivateDeactivatedCronContract{}, "onchainkeeper/MsgReactivateDeactivatedCronContract", nil)
	cdc.RegisterConcrete(&MsgUpdateParams{}, "onchainkeeper/MsgUpdateParams", nil)
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	// this line is used by starport scaffolding # 3

	registry.RegisterImplementations(
		(*sdk.Msg)(nil),
		&MsgRegisterCronContract{},
		&MsgUnregisterCronContract{},
		&MsgAcceptPendingRegisteredCronContract{},
		&MsgReactivateDeactivatedCronContract{},
		&MsgUpdateParams{},
	)
	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}
