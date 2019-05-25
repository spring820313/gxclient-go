package ops

import (
	"encoding/json"
	"gxclient-go/core/types"
	"gxclient-go/core/transaction"
)

// NewTransferOperation returns a new instance of TransferOperation
func NewTransferOperation(from, to types.ObjectID, amount, fee types.AssetAmount, memo *types.Memo) *TransferOperation {
	op := &TransferOperation{
		From:       from,
		To:         to,
		Amount:     amount,
		Fee:        fee,
		Memo:		memo,
		Extensions: []json.RawMessage{},
	}

	return op
}

// TransferOperation
type TransferOperation struct {
	From       types.ObjectID          `json:"from"`
	To         types.ObjectID          `json:"to"`
	Amount     types.AssetAmount       `json:"amount"`
	Fee        types.AssetAmount       `json:"fee"`
	Memo       *types.Memo             `json:"memo,omitempty"`
	Extensions []json.RawMessage `json:"extensions"`
}

func (op *TransferOperation) Type() types.OpType { return types.TransferOpType }

func (op *TransferOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(op.Type()))
	enc.Encode(op.Fee)
	enc.Encode(op.From)
	enc.Encode(op.To)
	enc.Encode(op.Amount)

	if op.Memo != nil && op.Memo.Message.Length() > 0 {
		enc.EncodeUVarint(1)
		enc.Encode(op.Memo)
	} else {
		//Memo?
		enc.EncodeUVarint(0)
	}
	//Extensions
	enc.EncodeUVarint(0)
	return enc.Err()
}
