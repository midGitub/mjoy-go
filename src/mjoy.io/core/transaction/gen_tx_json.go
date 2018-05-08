// Code generated by github.com/fjl/gencodec. DO NOT EDIT.

package transaction

import (
	"encoding/json"
	"errors"
	"math/big"

	"mjoy.io/common/types/util/hex"
	"mjoy.io/common/types"
)


func (t Txdata) MarshalJSON() ([]byte, error) {
	type txdata struct {
		AccountNonce hex.Uint64     `json:"nonce"    gencodec:"required"`
		Recipient    *types.Address `json:"to"       msgp:"nil"`
		Amount       *hex.Big       `json:"value"    gencodec:"required"`
		Payload      hex.Bytes      `json:"input"    gencodec:"required"`
		V            *hex.Big       `json:"v" gencodec:"required"`
		R            *hex.Big       `json:"r" gencodec:"required"`
		S            *hex.Big       `json:"s" gencodec:"required"`
		Hash         *types.Hash    `json:"hash" msgp:"-"`
	}
	var enc txdata
	enc.AccountNonce = hex.Uint64(t.AccountNonce)
	enc.Recipient = t.Recipient
	enc.Amount = (*hex.Big)(&t.Amount.IntVal)
	enc.Payload = t.Payload
	enc.V = (*hex.Big)(&t.V.IntVal)
	enc.R = (*hex.Big)(&t.R.IntVal)
	enc.S = (*hex.Big)(&t.S.IntVal)
	enc.Hash = t.Hash
	return json.Marshal(&enc)
}

func (t *Txdata) UnmarshalJSON(input []byte) error {
	type txdata struct {
		AccountNonce *hex.Uint64    `json:"nonce"    gencodec:"required"`
		Recipient    *types.Address `json:"to"       msgp:"nil"`
		Amount       *hex.Big       `json:"value"    gencodec:"required"`
		Payload      *hex.Bytes     `json:"input"    gencodec:"required"`
		V            *hex.Big       `json:"v" gencodec:"required"`
		R            *hex.Big       `json:"r" gencodec:"required"`
		S            *hex.Big       `json:"s" gencodec:"required"`
		Hash         *types.Hash    `json:"hash" msgp:"-"`
	}
	var dec txdata
	if err := json.Unmarshal(input, &dec); err != nil {
		return err
	}
	if dec.AccountNonce == nil {
		return errors.New("missing required field 'nonce' for txdata")
	}
	t.AccountNonce = uint64(*dec.AccountNonce)
	if dec.Recipient != nil {
		t.Recipient = dec.Recipient
	}
	if dec.Amount == nil {
		return errors.New("missing required field 'value' for txdata")
	}
	t.Amount = &types.BigInt{(big.Int)(*dec.Amount)}
	if dec.Payload == nil {
		return errors.New("missing required field 'input' for txdata")
	}
	t.Payload = *dec.Payload
	if dec.V == nil {
		return errors.New("missing required field 'v' for txdata")
	}
	t.V = &types.BigInt{(big.Int)(*dec.V)}
	if dec.R == nil {
		return errors.New("missing required field 'r' for txdata")
	}
	t.R = &types.BigInt{(big.Int)(*dec.R)}
	if dec.S == nil {
		return errors.New("missing required field 's' for txdata")
	}
	t.S = &types.BigInt{(big.Int)(*dec.S)}
	if dec.Hash != nil {
		t.Hash = dec.Hash
	}
	return nil
}
