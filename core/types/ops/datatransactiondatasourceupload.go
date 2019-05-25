package ops

import (
	"gxclient-go/core/types"
	"gxclient-go/core/transaction"
)

func NewDataTransactionDatasourceUploadOperation(requester, datasource types.GrapheneID, requestId string) *DataTransactionDatasourceUploadOperation {
	op := &DataTransactionDatasourceUploadOperation{
		RequestId: 	requestId,
		Requester:  requester,
		Datasource: datasource,
		Extensions: types.Extensions{},
	}

	return op
}

type DataTransactionDatasourceUploadOperation struct {
	RequestId   		string `json:"request_id"`
	Requester 			types.GrapheneID `json:"requester"`
	Datasource 			types.GrapheneID `json:"datasource"`
	types.OperationFee
	Extensions  		types.Extensions `json:"extensions"`
}

func (p DataTransactionDatasourceUploadOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.Encode(int8(p.Type()))
	enc.Encode(p.RequestId)
	enc.Encode(p.Requester)
	enc.Encode(p.Datasource)
	enc.Encode(p.Fee)
	enc.EncodeUVarint(0)

	return enc.Err()
}

func (op *DataTransactionDatasourceUploadOperation) Type() types.OpType { return types.DataTransactionDatasourceUploadOpType }