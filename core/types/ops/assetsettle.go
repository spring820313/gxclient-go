package ops

import (
	"gxclient-go/core/types"
	"gxclient-go/core/transaction"
)

func NewAssetSettleOperation(account types.GrapheneID, amount types.AssetAmount) *AssetSettleOperation {
	op := &AssetSettleOperation{
		Account:		account,
		Amount: 		amount,
		Extensions: 	types.Extensions{},
	}

	return op
}

type AssetSettleOperation struct {
	types.OperationFee
	Account    types.GrapheneID  `json:"account"`
	Amount     types.AssetAmount `json:"amount"`
	Extensions types.Extensions  `json:"extensions"`
}

func (p AssetSettleOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.Encode(int8(p.Type()))
	enc.Encode(p.Fee)
	enc.Encode(p.Account)
	enc.Encode(p.Amount)
	enc.EncodeUVarint(0)

	return enc.Err()
}

func (op *AssetSettleOperation) Type() types.OpType { return types.AssetSettleOpType }