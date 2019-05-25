package ops

import (
	"gxclient-go/core/types"
	"gxclient-go/core/transaction"
)

func NewAssetIssueOperation(issuer, issueToAccount types.GrapheneID, assetToIssue types.AssetAmount, memo *types.Memo) *AssetIssueOperation {
	op := &AssetIssueOperation{
		Issuer:			issuer,
		AssetToIssue: 	assetToIssue,
		IssueToAccount:	issueToAccount,
		Memo:			memo,
		Extensions: 	types.Extensions{},
	}

	return op
}

type AssetIssueOperation struct {
	types.OperationFee
	Issuer         types.GrapheneID  `json:"issuer"`
	AssetToIssue   types.AssetAmount `json:"asset_to_issue"`
	IssueToAccount types.GrapheneID  `json:"issue_to_account"`
	Memo           *types.Memo       `json:"memo,omitempty"`
	Extensions     types.Extensions  `json:"extensions"`
}

func (p AssetIssueOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.Encode(int8(p.Type()))
	enc.Encode(p.Fee)
	enc.Encode(p.Issuer)
	enc.Encode(p.AssetToIssue)
	enc.Encode(p.IssueToAccount)
	enc.Encode(p.Memo != nil)
	if(p.Memo != nil) {
		enc.Encode(p.Memo)
	}
	enc.EncodeUVarint(0)

	return enc.Err()
}

func (op *AssetIssueOperation) Type() types.OpType { return types.AssetIssueOpType }
