package ops

import (
	"gxclient-go/core/types"
	"gxclient-go/core/transaction"
)

func NewWithdrawPermissionClaimOperation(withdrawPermission, withdrawFromAccount, withdrawToAccount types.GrapheneID,
			amountToWithdraw types.AssetAmount, memo *types.Memo ) *WithdrawPermissionClaimOperation {
	op := &WithdrawPermissionClaimOperation{
		WithdrawPermission:		withdrawPermission,
		WithdrawFromAccount:	withdrawFromAccount,
		WithdrawToAccount: 		withdrawToAccount,
		AmountToWithdraw:		amountToWithdraw,
		Memo:					memo,
	}

	return op
}

type WithdrawPermissionClaimOperation struct {
	types.OperationFee
	WithdrawPermission    	types.GrapheneID `json:"withdraw_permission"`
	WithdrawFromAccount 	types.GrapheneID `json:"withdraw_from_account"`
	WithdrawToAccount       types.GrapheneID `json:"withdraw_to_account"`
	AmountToWithdraw     	types.AssetAmount `json:"amount_to_withdraw"`
	Memo  					*types.Memo 	  `json:"memo,omitempty"`
}

func (p WithdrawPermissionClaimOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.Encode(int8(p.Type()))
	enc.Encode(p.Fee)
	enc.Encode(p.WithdrawPermission)
	enc.Encode(p.WithdrawFromAccount)
	enc.Encode(p.WithdrawToAccount)
	enc.Encode(p.AmountToWithdraw)
	enc.Encode(p.Memo != nil)
	if(p.Memo != nil) {
		enc.Encode(p.Memo)
	}

	return enc.Err()
}

func (op *WithdrawPermissionClaimOperation) Type() types.OpType { return types.WithdrawPermissionClaimOpType }