package ops

import (
	"gxclient-go/core/types"
	"gxclient-go/core/transaction"
)

func NewAccountTransferOperation(accountId, newOwner types.GrapheneID) *AccountTransferOperation {
	op := &AccountTransferOperation{
		AccountId:		accountId,
		NewOwner: 		newOwner,
		Extensions: 	types.Extensions{},
	}

	return op
}

type AccountTransferOperation struct {
	types.OperationFee
	AccountId       types.GrapheneID `json:"account_id"`
	NewOwner 		types.GrapheneID `json:"new_owner"`
	Extensions      types.Extensions `json:"extensions"`
}

func (p AccountTransferOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.Encode(int8(p.Type()))
	enc.Encode(p.Fee)
	enc.Encode(p.AccountId)
	enc.Encode(p.NewOwner)
	enc.EncodeUVarint(0)

	return enc.Err()
}

func (op *AccountTransferOperation) Type() types.OpType { return types.AccountTransferOpType }