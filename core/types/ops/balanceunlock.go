package ops

import (
	"gxclient-go/core/types"
	"gxclient-go/core/transaction"
)

func NewBalanceUnlockOperation(account, lockId types.GrapheneID) *BalanceUnlockOperation {
	op := &BalanceUnlockOperation{
		Account: account,
		LockId:  lockId,
		Extensions: types.Extensions{},
	}

	return op
}

type BalanceUnlockOperation struct {
	types.OperationFee
	Account    		types.GrapheneID `json:"account"`
	LockId 			types.GrapheneID `json:"lock_id"`
	Extensions      types.Extensions `json:"extensions"`
}

func (p BalanceUnlockOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.Encode(int8(p.Type()))
	enc.Encode(p.Fee)
	enc.Encode(p.Account)
	enc.Encode(p.LockId)
	enc.EncodeUVarint(0)

	return enc.Err()
}

func (op *BalanceUnlockOperation) Type() types.OpType { return types.BalanceUnlockOpType }