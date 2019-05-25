package ops

import (
	"encoding/json"
	"gxclient-go/core/types"
	"gxclient-go/core/transaction"
)

func NewLimitOrderCreateOperation(seller types.ObjectID, fee, amountToSell, minToReceive types.AssetAmount,
	expiration types.Time, fillOrKill bool) *LimitOrderCreateOperation {
	op := &LimitOrderCreateOperation{
		Fee:        	fee,
		Seller:       	seller,
		AmountToSell:   amountToSell,
		MinToReceive:   minToReceive,
		Expiration:		expiration,
		FillOrKill:		fillOrKill,
		Extensions: 	[]json.RawMessage{},
	}

	return op
}

// LimitOrderCreateOperation
type LimitOrderCreateOperation struct {
	Fee          types.AssetAmount       `json:"fee"`
	Seller       types.ObjectID          `json:"seller"`
	AmountToSell types.AssetAmount       `json:"amount_to_sell"`
	MinToReceive types.AssetAmount       `json:"min_to_receive"`
	Expiration   types.Time              `json:"expiration"`
	FillOrKill   bool              `json:"fill_or_kill"`
	Extensions   []json.RawMessage `json:"extensions"`
}

func (op *LimitOrderCreateOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)

	enc.EncodeUVarint(uint64(op.Type()))
	enc.Encode(op.Fee)
	enc.Encode(op.Seller)
	enc.Encode(op.AmountToSell)
	enc.Encode(op.MinToReceive)
	enc.Encode(op.Expiration)
	enc.EncodeBool(op.FillOrKill)

	//extensions
	enc.EncodeUVarint(0)
	return enc.Err()
}

func (op *LimitOrderCreateOperation) Type() types.OpType { return types.LimitOrderCreateOpType }
