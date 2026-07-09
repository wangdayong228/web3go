# debug_blockProperties RPC Design

## Context

`conflux-rust` exposes an eSpace custom RPC method named `debug_blockProperties`.
The method accepts a block number or block hash and returns the block execution
properties used by Ethereum-space transactions in that eSpace block.

`web3go` already has a `RpcDebugClient` for `debug_*` RPC methods and a
`types.BlockNumberOrHash` wrapper used by other RPC clients for parameters that
can be either a block number or hash.

## Goal

Add support for `debug_blockProperties` to `web3go` while keeping the public API
consistent with existing RPC methods.

## Non-Goals

- Do not add separate `ByNumber` or `ByHash` convenience methods.
- Do not change existing trace RPC behavior.
- Do not add network-dependent tests as required CI coverage.

## Public API

Add one method to `client.RpcDebugClient`:

```go
func (c *RpcDebugClient) BlockProperties(block types.BlockNumberOrHash) (val []types.BlockProperties, err error)
```

The method calls:

```text
debug_blockProperties
```

with the provided `types.BlockNumberOrHash` as the sole parameter.

The API intentionally uses a non-pointer `types.BlockNumberOrHash`, matching
existing methods such as `RpcTraceClient.Blocks` and
`RpcTraceClient.ReplayBlockTransactions`. Callers construct the parameter with
`types.BlockNumberOrHashWithNumber` or `types.BlockNumberOrHashWithHash`.

## Data Model

Add `types.BlockProperties` matching the JSON returned by `conflux-rust`:

```go
type BlockProperties struct {
	TxHash         *common.Hash    `json:"txHash,omitempty"`
	InnerBlockHash common.Hash     `json:"innerBlockHash"`
	Coinbase       common.Address  `json:"coinbase"`
	Difficulty     *big.Int        `json:"difficulty"`
	GasLimit       *big.Int        `json:"gasLimit"`
	Timestamp      uint64          `json:"timestamp"`
	BaseFeePerGas  *big.Int        `json:"baseFeePerGas,omitempty"`
}
```

The exact formatting may be adjusted by `gofmt`, but field names and JSON names
must remain aligned with `conflux-rust`.

### JSON Encoding

`conflux-rust` serializes numeric fields as Ethereum hex quantities. `web3go`
should follow its existing model:

- `difficulty`, `gasLimit`, and `baseFeePerGas` are represented as `*big.Int`.
- `timestamp` is represented as `uint64`.
- Generated JSON codec overrides should map those fields to `hexutil.Big` and
  `hexutil.Uint64`.

`txHash` and `baseFeePerGas` are nullable in Rust and should remain pointers in
Go.

## Error Handling

- `CallContext` errors are returned unchanged.
- If the RPC returns `null`, the result remains a nil slice.
- No client-side fallback to `latest` is added because the method accepts a
  concrete non-pointer block identifier.

## Tests

Add focused unit coverage:

- JSON decode/encode test for `types.BlockProperties` using a fixture based on
  the `conflux-rust` documentation example.
- Coverage for nullable `txHash` and `baseFeePerGas`.
- RPC client test asserting that `RpcDebugClient.BlockProperties` calls
  `debug_blockProperties` with the provided `types.BlockNumberOrHash` parameter
  and decodes the response into `[]types.BlockProperties`.

Integration testing can be added as a manual test helper if useful, but the
required verification should not depend on live Conflux RPC access.

## Implementation Notes

- Place the client method in `client/client_debug.go`.
- Place the type near other block-related types in `types/types.go` unless the
  existing code generation pattern requires a separate file.
- If code generation is used, add the corresponding `//go:generate gencodec`
  directive and generated codec file.
- Run `gofmt` and the relevant Go tests before claiming completion.
