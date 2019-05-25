package ops

import (
	"gxclient-go/core/types"
	"gxclient-go/core/transaction"
)

func NewFillOrderOperation(orderId, accountId types.ObjectID, fee, pays, recives types.AssetAmount,) *FillOrderOperation {
	op := &FillOrderOperation{
		Fee:        fee,
		OrderId:   	orderId,
		AccountId:  accountId,
		Pays: 		pays,
		Recives:	recives,
	}

	return op
}

// FillOrderOpType
type FillOrderOperation struct {
	Fee     types.AssetAmount `json:"fee"`
	OrderId   types.ObjectID `json:"order_id"`
	AccountId types.ObjectID `json:"account_id"`
	Pays   types.AssetAmount `json:"pays"`
	Recives types.AssetAmount `json:"receives"`
}

func (op *FillOrderOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)

	enc.EncodeUVarint(uint64(op.Type()))
	enc.Encode(op.Fee)
	enc.Encode(op.OrderId)
	enc.Encode(op.AccountId)
	enc.Encode(op.Pays)
	enc.Encode(op.Recives)
	return enc.Err()
}

func (op *FillOrderOperation) Type() types.OpType { return types.FillOrderOpType }
