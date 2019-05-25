package ops

import (
	"gxclient-go/core/types"
	"gxclient-go/core/transaction"
)

func NewWithdrawPermissionCreateOperation(withdrawFromAccount, authorizedAccount types.GrapheneID,
		withdrawalLimit types.AssetAmount, withdrawalPeriodSec,  periodsUntilExpiration types.UInt32, periodStartTime types.Time) *WithdrawPermissionCreateOperation {
	op := &WithdrawPermissionCreateOperation{
		WithdrawFromAccount:		withdrawFromAccount,
		AuthorizedAccount: 			authorizedAccount,
		WithdrawalLimit:			withdrawalLimit,
		WithdrawalPeriodSec: 		withdrawalPeriodSec,
		PeriodsUntilExpiration:		periodsUntilExpiration,
		PeriodStartTime: 			periodStartTime,
	}

	return op
}

type WithdrawPermissionCreateOperation struct {
	types.OperationFee
	WithdrawFromAccount    	types.GrapheneID `json:"withdraw_from_account"`
	AuthorizedAccount 		types.GrapheneID `json:"authorized_account"`
	WithdrawalLimit         types.AssetAmount `json:"withdrawal_limit"`
	WithdrawalPeriodSec     types.UInt32 `json:"withdrawal_period_sec"`
	PeriodsUntilExpiration  types.UInt32 `json:"periods_until_expiration"`
	PeriodStartTime         types.Time `json:"period_start_time"`
}

func (p WithdrawPermissionCreateOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.Encode(int8(p.Type()))
	enc.Encode(p.Fee)
	enc.Encode(p.WithdrawFromAccount)
	enc.Encode(p.AuthorizedAccount)
	enc.Encode(p.WithdrawalLimit)
	enc.Encode(p.WithdrawalPeriodSec)
	enc.Encode(p.PeriodsUntilExpiration)
	enc.Encode(p.PeriodStartTime)

	return enc.Err()
}

func (op *WithdrawPermissionCreateOperation) Type() types.OpType { return types.WithdrawPermissionCreateOpType }
