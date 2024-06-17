package types

import (
	"cosmossdk.io/errors"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	TypeMsgRegisterCronContract                = "register_cron_contract"
	TypeMsgUnregisterCronContract              = "unregister_cron_contract"
	TypeMsgAcceptPendingRegisteredCronContract = "accept_pending_registered_cron_contract"
	TypeMsgReactivateDeactivatedCronContract   = "reactivate_deactivated_cron_contract"
	TypeMsgUpdateParams                        = "update_onchainkeeper_params"
)

var (
	_ sdk.Msg = &MsgRegisterCronContract{}
	_ sdk.Msg = &MsgUnregisterCronContract{}
	_ sdk.Msg = &MsgAcceptPendingRegisteredCronContract{}
	_ sdk.Msg = &MsgReactivateDeactivatedCronContract{}
	_ sdk.Msg = &MsgUpdateParams{}
)

func (msg MsgRegisterCronContract) Route() string { return RouteKey }

func (msg MsgRegisterCronContract) Type() string { return TypeMsgRegisterCronContract }

func (msg MsgRegisterCronContract) ValidateBasic() error {
	return ValidateAddresses(msg.SenderAddress, msg.ContractAddress)
}

func (msg *MsgRegisterCronContract) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

func (msg MsgRegisterCronContract) GetSigners() []sdk.AccAddress {
	from, _ := sdk.AccAddressFromBech32(msg.SenderAddress)
	return []sdk.AccAddress{from}
}

func (msg MsgUnregisterCronContract) Route() string { return RouteKey }

func (msg MsgUnregisterCronContract) Type() string { return TypeMsgUnregisterCronContract }

func (msg MsgUnregisterCronContract) ValidateBasic() error {
	return ValidateAddresses(msg.SenderAddress, msg.ContractAddress)
}

func (msg *MsgUnregisterCronContract) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

func (msg MsgUnregisterCronContract) GetSigners() []sdk.AccAddress {
	from, _ := sdk.AccAddressFromBech32(msg.SenderAddress)
	return []sdk.AccAddress{from}
}

func (msg MsgAcceptPendingRegisteredCronContract) Route() string { return RouteKey }

func (msg MsgAcceptPendingRegisteredCronContract) Type() string {
	return TypeMsgAcceptPendingRegisteredCronContract
}

func (msg MsgAcceptPendingRegisteredCronContract) ValidateBasic() error {
	return ValidateAddresses(msg.Authority, msg.ContractAddress)
}

func (msg *MsgAcceptPendingRegisteredCronContract) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

func (msg MsgAcceptPendingRegisteredCronContract) GetSigners() []sdk.AccAddress {
	from, _ := sdk.AccAddressFromBech32(msg.Authority)
	return []sdk.AccAddress{from}
}

func (msg MsgReactivateDeactivatedCronContract) Route() string { return RouteKey }

func (msg MsgReactivateDeactivatedCronContract) Type() string {
	return TypeMsgReactivateDeactivatedCronContract
}

func (msg MsgReactivateDeactivatedCronContract) ValidateBasic() error {
	return ValidateAddresses(msg.Authority, msg.ContractAddress)
}

func (msg *MsgReactivateDeactivatedCronContract) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

func (msg MsgReactivateDeactivatedCronContract) GetSigners() []sdk.AccAddress {
	from, _ := sdk.AccAddressFromBech32(msg.Authority)
	return []sdk.AccAddress{from}
}

// NewMsgUpdateParams creates new instance of MsgUpdateParams
func NewMsgUpdateParams(
	sender sdk.Address,
	adminAddress string,
	cronGasLimit uint64,
) *MsgUpdateParams {
	return &MsgUpdateParams{
		Authority: sender.String(),
		Params:    NewParams(adminAddress, cronGasLimit),
	}
}

func (msg MsgUpdateParams) Route() string { return RouteKey }

func (msg MsgUpdateParams) Type() string {
	return TypeMsgUpdateParams
}

func (msg MsgUpdateParams) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Authority); err != nil {
		return errors.Wrap(err, "invalid authority address")
	}

	return msg.Params.Validate()
}

func (msg *MsgUpdateParams) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

func (msg MsgUpdateParams) GetSigners() []sdk.AccAddress {
	from, _ := sdk.AccAddressFromBech32(msg.Authority)
	return []sdk.AccAddress{from}
}

func ValidateAddresses(addresses ...string) error {
	for _, address := range addresses {
		if _, err := sdk.AccAddressFromBech32(address); err != nil {
			return errors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid address :%s", address)
		}
	}
	return nil
}
