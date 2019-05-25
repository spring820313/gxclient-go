package ops

import (
	"gxclient-go/core/types"
	"gxclient-go/core/transaction"
)

func NewBalanceLockOperation(account types.GrapheneID, createDateTime types.Time, programId, memo string, lockDays, interestRate uint32, amount types.AssetAmount) *BalanceLockOperation {
	op := &BalanceLockOperation{
		Account: 			account,
		CreateDateTime:    	createDateTime,
		ProgramId: 			programId,
		Amount:				amount,
		LockDays:     		lockDays,
		InterestRate:		interestRate,
		Memo:				memo,
		Extensions: 		types.Extensions{},
	}

	return op
}

type BalanceLockOperation struct {
	types.OperationFee
	Account    		types.GrapheneID `json:"account"`
	CreateDateTime 	types.Time  `json:"create_date_time"`
	ProgramId	   	string	`json:"program_id"`
	Amount 			types.AssetAmount	`json:"amount"`
	LockDays 		uint32		`json:"lock_days"`
	InterestRate 	uint32		`json:"interest_rate"`
	Memo 			string		`json:"memo"`
	Extensions      types.Extensions `json:"extensions"`
}

func (p BalanceLockOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.Encode(int8(p.Type()))
	enc.Encode(p.Fee)
	enc.Encode(p.Account)
	enc.Encode(p.CreateDateTime)
	enc.Encode(p.ProgramId)
	enc.Encode(p.Amount)
	enc.Encode(p.LockDays)
	enc.Encode(p.InterestRate)
	enc.Encode(p.Memo)
	enc.EncodeUVarint(0)

	return enc.Err()
}

func (op *BalanceLockOperation) Type() types.OpType { return types.BalanceLockOpType }
