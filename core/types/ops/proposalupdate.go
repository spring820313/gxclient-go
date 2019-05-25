package ops

import (
	"gxclient-go/core/types"
	"gxclient-go/core/transaction"
)

func NewProposalUpdateOperation(feePayingAccount, proposal types.GrapheneID, activeApprovalsToAdd, activeApprovalsToRemove, ownerApprovalsToAdd, ownerApprovalsToRemove, keyApprovalsToAdd,  keyApprovalsToRemove types.GrapheneIDs) *ProposalUpdateOperation {
	op := &ProposalUpdateOperation{
		FeePayingAccount: 			feePayingAccount,
		Proposal:					proposal,
		ActiveApprovalsToAdd: 		activeApprovalsToAdd,
		ActiveApprovalsToRemove:	activeApprovalsToRemove,
		OwnerApprovalsToAdd: 		ownerApprovalsToAdd,
		OwnerApprovalsToRemove:		ownerApprovalsToRemove,
		KeyApprovalsToAdd: 			keyApprovalsToAdd,
		KeyApprovalsToRemove:		keyApprovalsToRemove,
		Extensions: 				types.Extensions{},
	}

	return op
}

type ProposalUpdateOperation struct {
	types.OperationFee
	FeePayingAccount        types.GrapheneID  `json:"fee_paying_account"`
	Proposal                types.GrapheneID  `json:"proposal"`
	ActiveApprovalsToAdd    types.GrapheneIDs `json:"active_approvals_to_add"`
	ActiveApprovalsToRemove types.GrapheneIDs `json:"active_approvals_to_remove"`
	OwnerApprovalsToAdd     types.GrapheneIDs `json:"owner_approvals_to_add"`
	OwnerApprovalsToRemove  types.GrapheneIDs `json:"owner_approvals_to_remove"`
	KeyApprovalsToAdd       types.GrapheneIDs `json:"key_approvals_to_add"`
	KeyApprovalsToRemove    types.GrapheneIDs `json:"key_approvals_to_remove"`
	Extensions              types.Extensions  `json:"extensions"`
}

func (p ProposalUpdateOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.Encode(int8(p.Type()))
	enc.Encode(p.Fee)
	enc.Encode(p.FeePayingAccount)
	enc.Encode(p.Proposal)
	enc.Encode(p.ActiveApprovalsToAdd)
	enc.Encode(p.ActiveApprovalsToRemove)
	enc.Encode(p.OwnerApprovalsToAdd)
	enc.Encode(p.OwnerApprovalsToRemove)
	enc.Encode(p.KeyApprovalsToAdd)
	enc.Encode(p.KeyApprovalsToRemove)
	enc.EncodeUVarint(0)

	return enc.Err()
}

func (op *ProposalUpdateOperation) Type() types.OpType { return types.ProposalUpdateOpType }
