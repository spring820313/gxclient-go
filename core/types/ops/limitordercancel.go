package ops

import (
	"encoding/json"
	"gxclient-go/core/types"
	"gxclient-go/core/transaction"
)

func NewLimitOrderCancelOperation(feePayingAccount, order types.ObjectID, fee types.AssetAmount,) *LimitOrderCancelOperation {
	op := &LimitOrderCancelOperation{
		Fee:        		fee,
		FeePayingAccount:   feePayingAccount,
		Order:   			order,
		Extensions: 		[]json.RawMessage{},
	}

	return op
}


// LimitOrderCancelOpType
type LimitOrderCancelOperation struct {
	Fee              types.AssetAmount       `json:"fee"`
	FeePayingAccount types.ObjectID          `json:"fee_paying_account"`
	Order            types.ObjectID          `json:"order"`
	Extensions       []json.RawMessage `json:"extensions"`
}

func (op *LimitOrderCancelOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)

	enc.EncodeUVarint(uint64(op.Type()))
	enc.Encode(op.Fee)
	enc.Encode(op.FeePayingAccount)
	enc.Encode(op.Order)

	// extensions
	enc.EncodeUVarint(0)
	return enc.Err()
}

func (op *LimitOrderCancelOperation) Type() types.OpType { return types.LimitOrderCancelOpType }
