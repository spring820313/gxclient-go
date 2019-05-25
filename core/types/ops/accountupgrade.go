package ops

import (
	"gxclient-go/core/types"
	"gxclient-go/core/transaction"
)

func NewAccountUpgradeOperation(accountToUpgrade types.GrapheneID,  upgradeToLifetimeMember bool) *AccountUpgradeOperation {
	op := &AccountUpgradeOperation{
		AccountToUpgrade:			accountToUpgrade,
		UpgradeToLifetimeMember: 	upgradeToLifetimeMember,
		Extensions: 				types.Extensions{},
	}

	return op
}

type AccountUpgradeOperation struct {
	types.OperationFee
	AccountToUpgrade        types.GrapheneID `json:"account_to_upgrade"`
	UpgradeToLifetimeMember bool             `json:"upgrade_to_lifetime_member"`
	Extensions              types.Extensions `json:"extensions"`
}

func (p AccountUpgradeOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.Encode(int8(p.Type()))
	enc.Encode(p.Fee)
	enc.Encode(p.AccountToUpgrade)
	enc.Encode(p.UpgradeToLifetimeMember)
	enc.EncodeUVarint(0)

	return enc.Err()
}

func (op *AccountUpgradeOperation) Type() types.OpType { return types.AccountUpgradeOpType }