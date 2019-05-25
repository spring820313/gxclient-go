package ops

import (
	"gxclient-go/core/types"
	"gxclient-go/core/transaction"
)

func NewFbaDistributeOperation() *FbaDistributeOperation {
	op := &FbaDistributeOperation{
		Extensions: types.Extensions{},
	}

	return op
}

type FbaDistributeOperation struct {
	types.OperationFee
	Extensions  types.Extensions  `json:"extensions"`
}

func (op *FbaDistributeOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)

	enc.EncodeUVarint(uint64(op.Type()))
	enc.Encode(op.Fee)

	//extensions
	enc.EncodeUVarint(0)
	return enc.Err()
}

func (op *FbaDistributeOperation) Type() types.OpType { return types.FbaDistributeOperationOpType }