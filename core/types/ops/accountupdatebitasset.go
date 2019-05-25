package ops

import (
	"gxclient-go/core/types"
	"gxclient-go/core/transaction"
)

func NewAssetUpdateBitassetOperation(issuer, assetToUpdate types.GrapheneID, newOptions types.BitassetOptions) *AssetUpdateBitassetOperation {
	op := &AssetUpdateBitassetOperation{
		Issuer:			issuer,
		AssetToUpdate: 	assetToUpdate,
		NewOptions:		newOptions,
		Extensions: 	types.Extensions{},
	}

	return op
}

type AssetUpdateBitassetOperation struct {
	types.OperationFee
	Issuer        types.GrapheneID      `json:"issuer"`
	AssetToUpdate types.GrapheneID      `json:"asset_to_update"`
	NewOptions    types.BitassetOptions `json:"new_options"`
	Extensions    types.Extensions      `json:"extensions"`
}

func (p AssetUpdateBitassetOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.Encode(int8(p.Type()))
	enc.Encode(p.Fee)
	enc.Encode(p.Issuer)
	enc.Encode(p.AssetToUpdate)
	enc.Encode(p.NewOptions)
	enc.EncodeUVarint(0)

	return enc.Err()
}

func (op *AssetUpdateBitassetOperation) Type() types.OpType { return types.AssetUpdateBitassetOpType }
