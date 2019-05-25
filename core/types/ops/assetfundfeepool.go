package ops

import (
	"gxclient-go/core/types"
	"gxclient-go/core/transaction"
)

func NewAssetFundFeePoolOperation(fromAccount, assetID types.GrapheneID, amount types.UInt64) *AssetFundFeePoolOperation {
	op := &AssetFundFeePoolOperation{
		FromAccount:	fromAccount,
		AssetID: 		assetID,
		Amount: 		amount,
		Extensions: 	types.Extensions{},
	}

	return op
}

type AssetFundFeePoolOperation struct {
	types.OperationFee
	FromAccount types.GrapheneID `json:"from_account"`
	AssetID     types.GrapheneID `json:"asset_id"`
	Amount      types.UInt64     `json:"amount"`
	Extensions  types.Extensions `json:"extensions"`

}

func (p AssetFundFeePoolOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.Encode(int8(p.Type()))
	enc.Encode(p.Fee)
	enc.Encode(p.FromAccount)
	enc.Encode(p.AssetID)
	enc.Encode(p.Amount)
	enc.EncodeUVarint(0)

	return enc.Err()
}

func (op *AssetFundFeePoolOperation) Type() types.OpType { return types.AssetFundFeePoolOpType }