package ops

import (
	"gxclient-go/core/types"
	"gxclient-go/core/transaction"
)

func NewVestingBalanceCreateOperation(creator, owner types.GrapheneID, amount types.AssetAmount, policy types.VestingPolicy) *VestingBalanceCreateOperation {
	op := &VestingBalanceCreateOperation{
		Creator:	creator,
		Owner: 		owner,
		Amount: 	amount,
		Policy: 	policy,
	}

	return op
}

type VestingBalanceCreateOperation struct {
	types.OperationFee
	Creator types.GrapheneID    `json:"creator"`
	Owner   types.GrapheneID    `json:"owner"`
	Amount  types.AssetAmount   `json:"amount"`
	Policy  types.VestingPolicy `json:"policy"`
}

func (p VestingBalanceCreateOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.Encode(int8(p.Type()))
	enc.Encode(p.Fee)
	enc.Encode(p.Creator)
	enc.Encode(p.Owner)
	enc.Encode(p.Amount)
	enc.Encode(p.Policy)

	return enc.Err()
}

func (op *VestingBalanceCreateOperation) Type() types.OpType { return types.VestingBalanceCreateOpType }