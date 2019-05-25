package ops

import (
	"gxclient-go/core/types"
	"gxclient-go/core/transaction"
)

func NewDataTransactionDatasourceValidateErrorOperation(datasource types.GrapheneID, requestId string) *DataTransactionDatasourceValidateErrorOperation {
	op := &DataTransactionDatasourceValidateErrorOperation{
		RequestId: 	requestId,
		Datasource: datasource,
		Extensions: types.Extensions{},
	}

	return op
}

type DataTransactionDatasourceValidateErrorOperation struct {
	RequestId   		string `json:"request_id"`
	Datasource 			types.GrapheneID `json:"datasource"`
	types.OperationFee
	Extensions  		types.Extensions `json:"extensions"`
}

func (p DataTransactionDatasourceValidateErrorOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.Encode(int8(p.Type()))
	enc.Encode(p.RequestId)
	enc.Encode(p.Datasource)
	enc.Encode(p.Fee)
	enc.EncodeUVarint(0)

	return enc.Err()
}

func (op *DataTransactionDatasourceValidateErrorOperation) Type() types.OpType { return types.DataTransactionDatasourceValidateErrorOpType }
