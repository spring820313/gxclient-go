package ops

import (
	"gxclient-go/core/types"
	"gxclient-go/core/transaction"
)

func NewWithdrawPermissionDeleteOperation(withdrawPermission, withdrawFromAccount, authorizedAccount types.GrapheneID) *WithdrawPermissionDeleteOperation {
	op := &WithdrawPermissionDeleteOperation{
		WithdrawPermission:		withdrawPermission,
		WithdrawFromAccount:	withdrawFromAccount,
		AuthorizedAccount: 		authorizedAccount,
	}

	return op
}

type WithdrawPermissionDeleteOperation struct {
	types.OperationFee
	WithdrawFromAccount 	types.GrapheneID `json:"withdraw_from_account"`
	AuthorizedAccount 		types.GrapheneID `json:"authorized_account"`
	WithdrawPermission    	types.GrapheneID `json:"withdraw_permission"`
}

func (p WithdrawPermissionDeleteOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.Encode(int8(p.Type()))
	enc.Encode(p.Fee)
	enc.Encode(p.WithdrawFromAccount)
	enc.Encode(p.AuthorizedAccount)
	enc.Encode(p.WithdrawPermission)

	return enc.Err()
}

func (op *WithdrawPermissionDeleteOperation) Type() types.OpType { return types.WithdrawPermissionDeleteOpType }