package ops

import (
	"gxclient-go/core/types"
	"gxclient-go/core/transaction"
)

func NewProposalCreateOperation(feePayingAccount types.GrapheneID, expirationTime types.Time, proposedOps types.OperationEnvelopeHolders, reviewPeriodSeconds *types.UInt32) *ProposalCreateOperation {
	op := &ProposalCreateOperation{
		FeePayingAccount: 		feePayingAccount,
		ExpirationTime:			expirationTime,
		ProposedOps: 			proposedOps,
		ReviewPeriodSeconds:	reviewPeriodSeconds,
		Extensions: 			types.Extensions{},
	}

	return op
}

type ProposalCreateOperation struct {
	types.OperationFee
	FeePayingAccount    types.GrapheneID               `json:"fee_paying_account"`
	ExpirationTime      types.Time                     `json:"expiration_time"`
	ProposedOps         types.OperationEnvelopeHolders `json:"proposed_ops"`
	ReviewPeriodSeconds *types.UInt32                  `json:"review_period_seconds,omitempty"`
	Extensions          types.Extensions               `json:"extensions"`
}

func (p ProposalCreateOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.Encode(int8(p.Type()))
	enc.Encode(p.Fee)
	enc.Encode(p.FeePayingAccount)
	enc.Encode(p.ExpirationTime)
	enc.Encode(p.ProposedOps)
	enc.Encode(p.ReviewPeriodSeconds != nil)
	if p.ReviewPeriodSeconds != nil {
		enc.Encode(*p.ReviewPeriodSeconds)
	}
	enc.EncodeUVarint(0)

	return enc.Err()
}

func (op *ProposalCreateOperation) Type() types.OpType { return types.ProposalCreateOpType }