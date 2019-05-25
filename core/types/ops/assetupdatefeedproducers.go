package ops

import (
	"gxclient-go/core/types"
	"gxclient-go/core/transaction"
)

func NewAssetUpdateFeedProducersOperation(issuer, assetToUpdate types.GrapheneID, newFeedProducers types.GrapheneIDs) *AssetUpdateFeedProducersOperation {
	op := &AssetUpdateFeedProducersOperation{
		Issuer:				issuer,
		AssetToUpdate: 		assetToUpdate,
		NewFeedProducers: 	newFeedProducers,
		Extensions: 		types.Extensions{},
	}

	return op
}

type AssetUpdateFeedProducersOperation struct {
	types.OperationFee
	Issuer           types.GrapheneID  `json:"issuer"`
	AssetToUpdate    types.GrapheneID  `json:"asset_to_update"`
	NewFeedProducers types.GrapheneIDs `json:"new_feed_producers"`
	Extensions       types.Extensions  `json:"extensions"`
}

func (p AssetUpdateFeedProducersOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.Encode(int8(p.Type()))
	enc.Encode(p.Fee)
	enc.Encode(p.Issuer)
	enc.Encode(p.AssetToUpdate)
	fpSize := len(p.NewFeedProducers)
	encoder.EncodeUVarint(uint64(fpSize))
	for _, f := range p.NewFeedProducers {
		encoder.Encode(f)
	}
	enc.EncodeUVarint(0)

	return enc.Err()
}

func (op *AssetUpdateFeedProducersOperation) Type() types.OpType { return types.AssetUpdateFeedProducersOpType }