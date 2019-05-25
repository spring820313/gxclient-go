package ops

import (
	"gxclient-go/core/types"
	"gxclient-go/core/transaction"
)

func NewAssetCreateOperation(issuer types.GrapheneID,  symbol string, precision  types.UInt8, commonOptions types.AssetOptions, bitassetOptions *types.BitassetOptions, isPredictionMarket bool) *AssetCreateOperation {
	op := &AssetCreateOperation{
		Issuer:				issuer,
		Symbol: 			symbol,
		Precision:   		precision,
		CommonOptions: 		commonOptions,
		BitassetOptions: 	bitassetOptions,
		IsPredictionMarket:	isPredictionMarket,
		Extensions: 		types.Extensions{},
	}

	return op
}

type AssetCreateOperation struct {
	types.OperationFee
	Issuer             types.GrapheneID       `json:"issuer"`
	Symbol             string                 `json:"symbol"`
	Precision          types.UInt8            `json:"precision"`
	CommonOptions      types.AssetOptions     `json:"common_options"`
	BitassetOptions    *types.BitassetOptions `json:"bitasset_opts,omitempty"`
	IsPredictionMarket bool                   `json:"is_prediction_market"`
	Extensions         types.Extensions       `json:"extensions"`
}

func (p AssetCreateOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.Encode(int8(p.Type()))
	enc.Encode(p.Fee)
	enc.Encode(p.Issuer)
	enc.Encode(p.Symbol)
	enc.Encode(p.Precision)
	enc.Encode(p.CommonOptions)
	enc.Encode(p.BitassetOptions != nil)
	if(p.BitassetOptions != nil) {
		enc.Encode(p.BitassetOptions)
	}
	enc.Encode(p.IsPredictionMarket)
	enc.EncodeUVarint(0)

	return enc.Err()
}

func (op *AssetCreateOperation) Type() types.OpType { return types.AssetCreateOpType }