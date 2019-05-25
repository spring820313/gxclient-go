package ops

import (
	"gxclient-go/core/types"
	"gxclient-go/core/transaction"
)

func NewProposalDeleteOperation(feePayingAccount, proposal types.GrapheneID, usingOwnerAuthority bool) *ProposalDeleteOperation {
	op := &ProposalDeleteOperation{
		FeePayingAccount:		feePayingAccount,
		UsingOwnerAuthority: 	usingOwnerAuthority,
		Proposal:				proposal,
		Extensions: 			types.Extensions{},
	}

	return op
}

type ProposalDeleteOperation struct {
	types.OperationFee
	FeePayingAccount    types.GrapheneID `json:"fee_paying_account"`
	UsingOwnerAuthority bool             `json:"using_owner_authority"`
	Proposal            types.GrapheneID `json:"proposal"`
	Extensions          types.Extensions `json:"extensions"`
}

func (p ProposalDeleteOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.Encode(int8(p.Type()))
	enc.Encode(p.Fee)
	enc.Encode(p.FeePayingAccount)
	enc.Encode(p.UsingOwnerAuthority)
	enc.Encode(p.Proposal)
	enc.EncodeUVarint(0)

	return enc.Err()
}

func (op *ProposalDeleteOperation) Type() types.OpType { return types.ProposalDeleteOpType }