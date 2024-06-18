package types

import (
	"cosmossdk.io/errors"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	_ sdk.Msg = &MsgRegisterCronContract{}
	_ sdk.Msg = &MsgUnregisterCronContract{}
	_ sdk.Msg = &MsgActivateCronContract{}
	_ sdk.Msg = &MsgUpdateParams{}
)

func (msg MsgRegisterCronContract) ValidateBasic() error {
	return ValidateAddresses(msg.SenderAddress, msg.ContractAddress)
}

func (msg MsgUnregisterCronContract) ValidateBasic() error {
	return ValidateAddresses(msg.SenderAddress, msg.ContractAddress)
}

func (msg MsgActivateCronContract) ValidateBasic() error {
	return ValidateAddresses(msg.ContractAddress)
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

func (msg MsgUpdateParams) ValidateBasic() error {
	ValidateAddresses(msg.Authority, msg.Params.AdminAddress)

	return msg.Params.Validate()
}

func ValidateAddresses(addresses ...string) error {
	for _, address := range addresses {
		if _, err := sdk.AccAddressFromBech32(address); err != nil {
			return errors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid address :%s", address)
		}
	}
	return nil
}
