package ops

import (
	"gxclient-go/core/types"
	"gxclient-go/core/transaction"
)


func NewAssetGlobalSettleOperation(issuer, assetToSettle types.GrapheneID, settlePrice types.Price) *AssetGlobalSettleOperation {
	op := &AssetGlobalSettleOperation{
		Issuer:			issuer,
		AssetToSettle: 	assetToSettle,
		SettlePrice: 	settlePrice,
		Extensions: 	types.Extensions{},
	}

	return op
}

type AssetGlobalSettleOperation struct {
	types.OperationFee
	Issuer    		types.GrapheneID  `json:"issuer"`
	AssetToSettle   types.GrapheneID `json:"asset_to_settle"`
	SettlePrice    	types.Price  `json:"settle_price"`
	Extensions 		types.Extensions  `json:"extensions"`
}

func (p AssetGlobalSettleOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.Encode(int8(p.Type()))
	enc.Encode(p.Fee)
	enc.Encode(p.Issuer)
	enc.Encode(p.AssetToSettle)
	enc.Encode(p.SettlePrice)
	enc.EncodeUVarint(0)

	return enc.Err()
}

func (op *AssetGlobalSettleOperation) Type() types.OpType { return types.AssetGlobalSettleOpType }
