package ops

import (
	"gxclient-go/core/types"
	"gxclient-go/core/transaction"
)

func NewCommitteeMemberCreateOperation(committeeMemberAccount types.GrapheneID, url string) *CommitteeMemberCreateOperation {
	op := &CommitteeMemberCreateOperation{
		CommitteeMemberAccount:		committeeMemberAccount,
		URL:						url,
	}

	return op
}

type CommitteeMemberCreateOperation struct {
	types.OperationFee
	CommitteeMemberAccount    	types.GrapheneID `json:"committee_member_account"`
	URL 						string `json:"url"`
}

func (p CommitteeMemberCreateOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.Encode(int8(p.Type()))
	enc.Encode(p.Fee)
	enc.Encode(p.CommitteeMemberAccount)
	enc.Encode(p.URL)

	return enc.Err()
}

func (op *CommitteeMemberCreateOperation) Type() types.OpType { return types.CommitteeMemberCreateOpType }