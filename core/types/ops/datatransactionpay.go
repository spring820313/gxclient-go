package ops

import (
	"gxclient-go/core/types"
	"gxclient-go/core/transaction"
)

func NewDataTransactionPayOperation(from, to types.GrapheneID, amount types.AssetAmount, requestId string) *DataTransactionPayOperation {
	op := &DataTransactionPayOperation{
		From: 		from,
		To:  		to,
		Amount: 	amount,
		RequestId:  requestId,
		Extensions: types.Extensions{},
	}

	return op
}

type DataTransactionPayOperation struct {
	types.OperationFee
	From 		types.GrapheneID `json:"from"`
	To 			types.GrapheneID `json:"to"`
	Amount 		types.AssetAmount	`json:"amount"`
	RequestId   string `json:"request_id"`
	Extensions  types.Extensions `json:"extensions"`
}

func (p DataTransactionPayOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.Encode(int8(p.Type()))
	enc.Encode(p.Fee)
	enc.Encode(p.From)
	enc.Encode(p.To)
	enc.Encode(p.Amount)
	enc.Encode(p.RequestId)
	enc.EncodeUVarint(0)

	return enc.Err()
}

func (op *DataTransactionPayOperation) Type() types.OpType { return types.DataTransactionPayOpType }