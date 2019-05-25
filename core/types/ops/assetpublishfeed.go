package ops

import (
	"gxclient-go/core/types"
	"gxclient-go/core/transaction"
)


func NewAssetPublishFeedOperation(publisher, assetID types.GrapheneID, feed types.PriceFeed) *AssetPublishFeedOperation {
	op := &AssetPublishFeedOperation{
		Publisher:		publisher,
		AssetID: 		assetID,
		Feed: 			feed,
		Extensions: 	types.Extensions{},
	}

	return op
}

type AssetPublishFeedOperation struct {
	types.OperationFee
	Publisher  types.GrapheneID `json:"publisher"`
	AssetID    types.GrapheneID `json:"asset_id"`
	Feed       types.PriceFeed  `json:"feed"`
	Extensions types.Extensions `json:"extensions"`
}


func (p AssetPublishFeedOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.Encode(int8(p.Type()))
	enc.Encode(p.Fee)
	enc.Encode(p.Publisher)
	enc.Encode(p.AssetID)
	enc.Encode(p.Feed)
	enc.EncodeUVarint(0)

	return enc.Err()
}

func (op *AssetPublishFeedOperation) Type() types.OpType { return types.AssetPublishFeedOpType }