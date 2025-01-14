package types

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
)

var _ = (*callRequestMarshaling)(nil)

// MarshalJSON marshals as JSON.
func (c CallRequest) MarshalJSON() ([]byte, error) {
	// handle input and data follow go-ethereum server side
	if c.Input != nil && c.Data != nil && !bytes.Equal(c.Input, c.Data) {
		return nil, fmt.Errorf("both 'input' and 'data' provided but not same")
	}
	if c.Input != nil {
		c.Data = c.Input
		c.Input = nil
	}

	type CallRequest struct {
		From                 *common.Address `json:"from"`
		To                   *common.Address `json:"to"`
		Gas                  *hexutil.Uint64 `json:"gas"`
		GasPrice             *hexutil.Big    `json:"gasPrice"`
		MaxFeePerGas         *hexutil.Big    `json:"maxFeePerGas"`
		MaxPriorityFeePerGas *hexutil.Big    `json:"maxPriorityFeePerGas"`
		Value                *hexutil.Big    `json:"value"`
		Nonce                *hexutil.Uint64 `json:"nonce"`
		Data                 *hexutil.Bytes  `json:"data"`
		// Input                *hexutil.Bytes    `json:"input"`
		AccessList *types.AccessList `json:"accessList,omitempty"`
		ChainID    *hexutil.Big      `json:"chainId,omitempty"`
	}
	var enc CallRequest
	enc.From = c.From
	enc.To = c.To
	enc.Gas = (*hexutil.Uint64)(c.Gas)
	enc.GasPrice = (*hexutil.Big)(c.GasPrice)
	enc.MaxFeePerGas = (*hexutil.Big)(c.MaxFeePerGas)
	enc.MaxPriorityFeePerGas = (*hexutil.Big)(c.MaxPriorityFeePerGas)
	enc.Value = (*hexutil.Big)(c.Value)
	enc.Nonce = (*hexutil.Uint64)(c.Nonce)
	if c.Data != nil {
		enc.Data = (*hexutil.Bytes)(&c.Data)
	}
	enc.AccessList = c.AccessList
	enc.ChainID = (*hexutil.Big)(c.ChainID)

	return json.Marshal(&enc)
}

// UnmarshalJSON unmarshals from JSON.
func (c *CallRequest) UnmarshalJSON(input []byte) error {
	type CallRequest struct {
		From                 *common.Address   `json:"from"`
		To                   *common.Address   `json:"to"`
		Gas                  *hexutil.Uint64   `json:"gas"`
		GasPrice             *hexutil.Big      `json:"gasPrice"`
		MaxFeePerGas         *hexutil.Big      `json:"maxFeePerGas"`
		MaxPriorityFeePerGas *hexutil.Big      `json:"maxPriorityFeePerGas"`
		Value                *hexutil.Big      `json:"value"`
		Nonce                *hexutil.Uint64   `json:"nonce"`
		Data                 *hexutil.Bytes    `json:"data"`
		Input                *hexutil.Bytes    `json:"input"`
		AccessList           *types.AccessList `json:"accessList,omitempty"`
		ChainID              *hexutil.Big      `json:"chainId,omitempty"`
	}
	var dec CallRequest
	if err := json.Unmarshal(input, &dec); err != nil {
		return err
	}
	if dec.From != nil {
		c.From = dec.From
	}
	if dec.To != nil {
		c.To = dec.To
	}
	if dec.Gas != nil {
		c.Gas = (*uint64)(dec.Gas)
	}
	if dec.GasPrice != nil {
		c.GasPrice = (*big.Int)(dec.GasPrice)
	}
	if dec.MaxFeePerGas != nil {
		c.MaxFeePerGas = (*big.Int)(dec.MaxFeePerGas)
	}
	if dec.MaxPriorityFeePerGas != nil {
		c.MaxPriorityFeePerGas = (*big.Int)(dec.MaxPriorityFeePerGas)
	}
	if dec.Value != nil {
		c.Value = (*big.Int)(dec.Value)
	}
	if dec.Nonce != nil {
		c.Nonce = (*uint64)(dec.Nonce)
	}
	if dec.Data != nil {
		c.Data = *dec.Data
	}
	if dec.Input != nil {
		c.Input = *dec.Input
	}
	if dec.AccessList != nil {
		c.AccessList = dec.AccessList
	}
	if dec.ChainID != nil {
		c.ChainID = (*big.Int)(dec.ChainID)
	}
	return nil
}
