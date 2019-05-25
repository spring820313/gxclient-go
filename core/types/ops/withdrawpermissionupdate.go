package ops

import (
	"gxclient-go/core/types"
	"gxclient-go/core/transaction"
)

func NewWithdrawPermissionUpdateOperation(withdrawFromAccount, authorizedAccount, permissionToUpdate types.GrapheneID,
	withdrawalLimit types.AssetAmount, withdrawalPeriodSec,  periodsUntilExpiration types.UInt32, periodStartTime types.Time) *WithdrawPermissionUpdateOperation {
	op := &WithdrawPermissionUpdateOperation{
		WithdrawFromAccount:		withdrawFromAccount,
		AuthorizedAccount: 			authorizedAccount,
		PermissionToUpdate:			permissionToUpdate,
		WithdrawalLimit:			withdrawalLimit,
		WithdrawalPeriodSec: 		withdrawalPeriodSec,
		PeriodsUntilExpiration:		periodsUntilExpiration,
		PeriodStartTime: 			periodStartTime,
	}

	return op
}

type WithdrawPermissionUpdateOperation struct {
	types.OperationFee
	WithdrawFromAccount    	types.GrapheneID `json:"withdraw_from_account"`
	AuthorizedAccount 		types.GrapheneID `json:"authorized_account"`
	PermissionToUpdate      types.GrapheneID `json:"permission_to_update"`
	WithdrawalLimit         types.AssetAmount `json:"withdrawal_limit"`
	WithdrawalPeriodSec     types.UInt32 `json:"withdrawal_period_sec"`
	PeriodStartTime         types.Time `json:"period_start_time"`
	PeriodsUntilExpiration  types.UInt32 `json:"periods_until_expiration"`
}

func (p WithdrawPermissionUpdateOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.Encode(int8(p.Type()))
	enc.Encode(p.Fee)
	enc.Encode(p.WithdrawFromAccount)
	enc.Encode(p.AuthorizedAccount)
	enc.Encode(p.PermissionToUpdate)
	enc.Encode(p.WithdrawalLimit)
	enc.Encode(p.WithdrawalPeriodSec)
	enc.Encode(p.PeriodStartTime)
	enc.Encode(p.PeriodsUntilExpiration)

	return enc.Err()
}

func (op *WithdrawPermissionUpdateOperation) Type() types.OpType { return types.WithdrawPermissionUpdateOpType }
