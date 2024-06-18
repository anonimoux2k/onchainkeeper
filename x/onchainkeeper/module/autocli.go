package onchainkeeper

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "wasmapp/api/wasmapp/onchainkeeper"
)

// AutoCLIOptions implements the autocli.HasAutoCLIConfig interface.
func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Query: &autocliv1.ServiceCommandDescriptor{
			Service: modulev1.Query_ServiceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "Params",
					Use:       "params",
					Short:     "Shows the parameters of the module",
				},
				{
					RpcMethod:      "CronContract",
					Use:            "cron-contract [contract_address]",
					Short:          "Shows the info of a cron contract given its address",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "contract_address"}},
				},
				{
					RpcMethod: "CronContracts",
					Use:       "list-cron-contract",
					Short:     "List all registered cron contract",
				},
			},
		},
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service:              modulev1.Msg_ServiceDesc.ServiceName,
			EnhanceCustomCommand: true, // only required if you want to use the custom command
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "UpdateParams",
					Skip:      true, // skipped because authority gated
				},
				{
					RpcMethod:      "RegisterCronContract",
					Use:            "register-cron [contract_address]",
					Short:          "Register a cron contract",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "contract_address"}},
				},
				{
					RpcMethod:      "UnregisterCronContract",
					Use:            "unregister-cron [contract_address]",
					Short:          "Unregister a cron contract",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "contract_address"}},
				},
				{
					RpcMethod:      "ActivateCronContract",
					Use:            "activate-cron [contract_address]",
					Short:          "Activate a cron contract. For security admin only",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "contract_address"}},
				},
			},
		},
	}
}
