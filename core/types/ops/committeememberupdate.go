package ops

import (
	"gxclient-go/core/types"
	"gxclient-go/core/transaction"
)

func NewCommitteeMemberUpdateOperation(committeeMember, committeeMemberAccount types.GrapheneID, newUrl *string) *CommitteeMemberUpdateOperation {
	op := &CommitteeMemberUpdateOperation{
		CommitteeMember:			committeeMember,
		CommitteeMemberAccount:		committeeMemberAccount,
		NewURL:						newUrl,
	}

	return op
}

type CommitteeMemberUpdateOperation struct {
	types.OperationFee
	CommitteeMember    			types.GrapheneID `json:"committee_member"`
	CommitteeMemberAccount    	types.GrapheneID `json:"committee_member_account"`
	NewURL 						*string `json:"new_url,omitempty"`
}

func (p CommitteeMemberUpdateOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.Encode(int8(p.Type()))
	enc.Encode(p.Fee)
	enc.Encode(p.CommitteeMember)
	enc.Encode(p.CommitteeMemberAccount)
	enc.Encode(p.NewURL != nil)
	if p.NewURL != nil {
		enc.Encode(*p.NewURL)
	}

	return enc.Err()
}

func (op *CommitteeMemberUpdateOperation) Type() types.OpType { return types.CommitteeMemberUpdateOpType }
