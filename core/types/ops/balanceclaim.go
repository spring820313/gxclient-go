package ops

import (
	"gxclient-go/core/types"
	"gxclient-go/core/transaction"
)

func NewBalanceClaimOperation(depositToAccount, balanceToClaim types.GrapheneID, balanceOwnerKey types.PublicKey, totalClaimed types.AssetAmount) *BalanceClaimOperation {
	op := &BalanceClaimOperation{
		DepositToAccount:	depositToAccount,
		BalanceToClaim:		balanceToClaim,
		BalanceOwnerKey:	balanceOwnerKey,
		TotalClaimed:		totalClaimed,
	}

	return op
}

type BalanceClaimOperation struct {
	types.OperationFee
	DepositToAccount types.GrapheneID  `json:"deposit_to_account"`
	BalanceToClaim   types.GrapheneID  `json:"balance_to_claim"`
	BalanceOwnerKey  types.PublicKey   `json:"balance_owner_key"`
	TotalClaimed     types.AssetAmount `json:"total_claimed"`
}

func (p BalanceClaimOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.Encode(int8(p.Type()))
	enc.Encode(p.Fee)
	enc.Encode(p.DepositToAccount)
	enc.Encode(p.BalanceToClaim)
	enc.Encode(p.BalanceOwnerKey)
	enc.Encode(p.TotalClaimed)

	return enc.Err()
}

func (op *BalanceClaimOperation) Type() types.OpType { return types.BalanceClaimOpType }