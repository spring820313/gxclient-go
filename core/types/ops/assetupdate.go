package ops

import (
	"gxclient-go/core/types"
	"gxclient-go/core/transaction"
)

func NewAssetUpdateOperation(issuer, assetToUpdate types.GrapheneID, newIssuer *types.GrapheneID, newOptions types.AssetOptions) *AssetUpdateOperation {
	op := &AssetUpdateOperation{
		Issuer:			issuer,
		AssetToUpdate: 	assetToUpdate,
		NewIssuer:		newIssuer,
		NewOptions:		newOptions,
		Extensions: 	types.Extensions{},
	}

	return op
}

type AssetUpdateOperation struct {
	types.OperationFee
	Issuer        types.GrapheneID   `json:"issuer"`
	AssetToUpdate types.GrapheneID   `json:"asset_to_update"`
	NewIssuer     *types.GrapheneID  `json:"new_issuer,omitempty"`
	NewOptions    types.AssetOptions `json:"new_options"`
	Extensions    types.Extensions   `json:"extensions"`
}

func (p AssetUpdateOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.Encode(int8(p.Type()))
	enc.Encode(p.Fee)
	enc.Encode(p.Issuer)
	enc.Encode(p.AssetToUpdate)
	enc.Encode(p.NewIssuer != nil)
	if(p.NewIssuer != nil) {
		enc.Encode(p.NewIssuer)
	}
	enc.Encode(p.NewOptions)
	enc.EncodeUVarint(0)

	return enc.Err()
}

func (op *AssetUpdateOperation) Type() types.OpType { return types.AssetUpdateOpType }
