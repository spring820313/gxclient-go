package ops

import (
	"gxclient-go/core/types"
	"gxclient-go/core/transaction"
)

func NewCustomOperation(payer types.GrapheneID, requiredAuths types.GrapheneIDs, id uint16, data []byte) *CustomOperation {
	op := &CustomOperation{
		Payer:			payer,
		RequiredAuths:	requiredAuths,
		ID:				id,
		Data:			data,
	}

	return op
}

type CustomOperation struct {
	types.OperationFee
	Payer    			types.GrapheneID `json:"payer"`
	RequiredAuths    	types.GrapheneIDs `json:"required_auths"`
	ID 					uint16 `json:"id"`
	Data				[]byte `json:"data"`
}

func (p CustomOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.Encode(int8(p.Type()))
	enc.Encode(p.Fee)
	enc.Encode(p.Payer)
	enc.Encode(p.RequiredAuths)
	enc.Encode(p.ID)

	size := len(p.Data)
	enc.EncodeUVarint(uint64(size))
	enc.Encode(p.Data)

	return enc.Err()
}

func (op *CustomOperation) Type() types.OpType { return types.CustomOpType }