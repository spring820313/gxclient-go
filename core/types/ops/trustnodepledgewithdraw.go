package ops

import (
	"gxclient-go/core/types"
	"gxclient-go/core/transaction"
)

func NewTrustNodePledgeWithdrawOperation(witnessAccount types.GrapheneID) *TrustNodePledgeWithdrawOperation {
	op := &TrustNodePledgeWithdrawOperation{
		WitnessAccount:	witnessAccount,
	}

	return op
}

type TrustNodePledgeWithdrawOperation struct {
	types.OperationFee
	WitnessAccount      	types.GrapheneID  `json:"witness_account"`
}

func (p TrustNodePledgeWithdrawOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.Encode(int8(p.Type()))
	enc.Encode(p.Fee)
	enc.Encode(p.WitnessAccount)

	return enc.Err()
}

func (op *TrustNodePledgeWithdrawOperation) Type() types.OpType { return types.TrustNodePledgeWithdrawOpType }