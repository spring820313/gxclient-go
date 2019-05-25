package ops

import (
	"gxclient-go/core/types"
	"gxclient-go/core/transaction"
)

func NewAssetSettleCancelOperation(settlement, account types.GrapheneID, amount types.AssetAmount) *AssetSettleCancelOperation {
	op := &AssetSettleCancelOperation{
		Settlement:		settlement,
		Account: 		account,
		Amount:			amount,
		Extensions: 	types.Extensions{},
	}

	return op
}

type AssetSettleCancelOperation struct {
	types.OperationFee
	Settlement      types.GrapheneID  `json:"settlement"`
	Account       	types.GrapheneID  `json:"account"`
	Amount     		types.AssetAmount `json:"amount"`
	Extensions      types.Extensions `json:"extensions"`
}

func (p AssetSettleCancelOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.Encode(int8(p.Type()))
	enc.Encode(p.Fee)
	enc.Encode(p.Settlement)
	enc.Encode(p.Account)
	enc.Encode(p.Amount)
	enc.EncodeUVarint(0)

	return enc.Err()
}

func (op *AssetSettleCancelOperation) Type() types.OpType { return types.AssetSettleCancelOpType }
