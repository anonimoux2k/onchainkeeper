# Automation on-chain keeper for Cosmos chain

**onchainkeeper** is a Cosmos module that acts as an on-chain Chainlink keeper, triggering wasm contract logic at the beginning and the end of every block.

**wasmapp** is a Cosmos chain for testing the `onchainkeeper` module.

## PoC

**Step 1**: Store and instantiate an instance of a [cw-cron-example] contract. This cron contract simply increases b by 1 at the beginning of the block, and increases e by 1 at the end of the block.

```bash
wasmappd tx wasm store artifacts/cw_cron_example.wasm --from bob --chain-id wasmapp --gas auto --gas-adjustment 1.2
wasmappd tx wasm instantiate 1 '{}' --from bob --label "cw-cron-example" --admin cosmos13577qg2j4q2dvgl90yg6ulk6juh5wgtwkd9pa5 --gas auto --gas-adjustment 1.2 --chain-id wasmapp 
```

Try querying state of the cron, of course this time the values ​​are all 0.

```bash
wasmappd q wasm contract-state smart cosmos14hj2tavq8fpesdwxxcu44rty3hh90vhujrvcmstl4zr3txmfvw9s4hmalr '{"get_config": {}}'
```

The output:

```terminal
data:
  b: 0
  e: 0
```

**Step 2**: Admin (bob: cosmos13577qg2j4q2dvgl90yg6ulk6juh5wgtwkd9pa5) of this cron contract submits a tx to register it with the module.

```bash
wasmappd tx onchainkeeper register-cron cosmos14hj2tavq8fpesdwxxcu44rty3hh90vhujrvcmstl4zr3txmfvw9s4hmalr --from bob --gas auto --gas-adjustment 1.2 --chain-id wasmapp
```

**Step 3**: Security admin (alice: cosmos105khsf5w9x08alwcqx055nhag2xm7dcahxegda) of the module accepts the above registration by submitting a tx to activate the cron contract.

```bash
wasmappd tx onchainkeeper activate-cron cosmos14hj2tavq8fpesdwxxcu44rty3hh90vhujrvcmstl4zr3txmfvw9s4hmalr --from alice --gas auto --gas-adjustment 1.2 --chain-id wasmapp
```

**Step 4**: Try querying state of the cron contract after some blocks.

```bash
wasmappd q wasm contract-state smart cosmos14hj2tavq8fpesdwxxcu44rty3hh90vhujrvcmstl4zr3txmfvw9s4hmalr '{"get_config": {}}'
```

Output:

```terminal
data:
  b: 9
  e: 10
```

In the above output, you can see the value of b is one less than e. Why? Because in the first block where cron takes effect, the contract only received the SudoMsg::CronEndBlock message. Only when this block is finalized and move to the next block, the contract starts receiving SudoMsg::CronBeginBlock

Finally, immediately query state of the cron contract in the next block.

```bash
wasmappd q wasm contract-state smart cosmos14hj2tavq8fpesdwxxcu44rty3hh90vhujrvcmstl4zr3txmfvw9s4hmalr '{"get_config": {}}'
```

Output:

```terminal
data:
  b: 10
  e: 11
```

[cw-cron-example]: https://github.com/anonimoux2k/cw-cron-example
