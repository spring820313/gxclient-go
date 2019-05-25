package ops

import (
	"gxclient-go/core/types"
	"gxclient-go/core/transaction"
)

func NewDataTransactionUpdateOperation(requestId, memo  string, requester types.GrapheneID, newStatus uint8) *DataTransactionUpdateOperation {
	op := &DataTransactionUpdateOperation{
		RequestId: 		requestId,
		NewStatus: 		newStatus,
		NewRequester: 	requester,
		Memo:  			memo,
		Extensions: types.Extensions{},
	}

	return op
}

type DataTransactionUpdateOperation struct {
	RequestId    		string `json:"request_id"`
	NewStatus 			uint8  `json:"new_status"`
	types.OperationFee
	NewRequester 		types.GrapheneID `json:"new_requester"`
	Memo 				string		`json:"memo"`
	Extensions      	types.Extensions `json:"extensions"`
}

func (p DataTransactionUpdateOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.Encode(int8(p.Type()))
	enc.Encode(p.RequestId)
	enc.Encode(p.NewStatus)
	enc.Encode(p.Fee)
	enc.Encode(p.NewRequester)
	enc.Encode(p.Memo)
	enc.EncodeUVarint(0)

	return enc.Err()
}

func (op *DataTransactionUpdateOperation) Type() types.OpType { return types.DataTransactionUpdateOpType }
