package ops

import (
	"gxclient-go/core/types"
	"gxclient-go/core/transaction"
)

func NewVestingBalanceWithdrawOperation(vestingBalance, owner types.GrapheneID, amount types.AssetAmount, policy types.VestingPolicy) *VestingBalanceWithdrawOperation {
	op := &VestingBalanceWithdrawOperation{
		VestingBalance:	vestingBalance,
		Owner: 			owner,
		Amount: 		amount,
	}

	return op
}

type VestingBalanceWithdrawOperation struct {
	types.OperationFee
	VestingBalance types.GrapheneID  `json:"vesting_balance"`
	Owner          types.GrapheneID  `json:"owner"`
	Amount         types.AssetAmount `json:"amount"`
}

func (p VestingBalanceWithdrawOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.Encode(int8(p.Type()))
	enc.Encode(p.Fee)
	enc.Encode(p.VestingBalance)
	enc.Encode(p.Owner)
	enc.Encode(p.Amount)

	return enc.Err()
}

func (op *VestingBalanceWithdrawOperation) Type() types.OpType { return types.VestingBalanceWithdrawOpType }