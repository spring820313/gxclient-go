package ops

import (
	"gxclient-go/core/types"
	"gxclient-go/core/transaction"
)

func NewCommitteeMemberUpdateGlobalParametersOperation(newParameters types.ChainParameters) *CommitteeMemberUpdateGlobalParametersOperation {
	op := &CommitteeMemberUpdateGlobalParametersOperation{
		NewParameters:	newParameters,
	}

	return op
}

type CommitteeMemberUpdateGlobalParametersOperation struct {
	types.OperationFee
	NewParameters   types.ChainParameters  `json:"new_parameters"`
}

func (p CommitteeMemberUpdateGlobalParametersOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.Encode(int8(p.Type()))
	enc.Encode(p.Fee)
	enc.Encode(p.NewParameters)

	return enc.Err()
}

func (op *CommitteeMemberUpdateGlobalParametersOperation) Type() types.OpType { return types.CommitteeMemberUpdateGlobalParametersOpType }
