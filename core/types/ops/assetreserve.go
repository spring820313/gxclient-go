package ops

import (
	"gxclient-go/core/types"
	"gxclient-go/core/transaction"
)

func NewAssetReserveOperation(payer types.GrapheneID, amountToReserve types.AssetAmount) *AssetReserveOperation {
	op := &AssetReserveOperation{
		Payer:				payer,
		AmountToReserve: 	amountToReserve,
		Extensions: 		types.Extensions{},
	}

	return op
}

type AssetReserveOperation struct {
	types.OperationFee
	Payer           types.GrapheneID  `json:"payer"`
	AmountToReserve types.AssetAmount `json:"amount_to_reserve"`
	Extensions      types.Extensions  `json:"extensions"`
}

func (p AssetReserveOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.Encode(int8(p.Type()))
	enc.Encode(p.Fee)
	enc.Encode(p.Payer)
	enc.Encode(p.AmountToReserve)
	enc.EncodeUVarint(0)

	return enc.Err()
}

func (op *AssetReserveOperation) Type() types.OpType { return types.AssetReserveOpType }