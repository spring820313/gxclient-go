package ops

import (
	"gxclient-go/core/types"
	"gxclient-go/core/transaction"
)


func NewOverrideTransferOperation(issuer, from, to types.GrapheneID, amount types.AssetAmount, memo *types.Memo) *OverrideTransferOperation {
	op := &OverrideTransferOperation{
		Issuer:		issuer,
		From:		from,
		To:			to,
		Amount: 	amount,
		Memo:		memo,
		Extensions: types.Extensions{},
	}

	return op
}

type OverrideTransferOperation struct {
	types.OperationFee
	Issuer     types.GrapheneID  `json:"issuer"`
	From       types.GrapheneID  `json:"from"`
	To         types.GrapheneID  `json:"to"`
	Amount     types.AssetAmount `json:"amount"`
	Memo       *types.Memo       `json:"memo,omitempty"`
	Extensions types.Extensions  `json:"extensions"`
}

func (p OverrideTransferOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.Encode(int8(p.Type()))
	enc.Encode(p.Fee)
	enc.Encode(p.Issuer)
	enc.Encode(p.From)
	enc.Encode(p.To)
	enc.Encode(p.Amount)
	enc.Encode(p.Memo != nil)
	if p.Memo != nil {
		enc.Encode(p.Memo)
	}
	enc.EncodeUVarint(0)

	return enc.Err()
}

func (op *OverrideTransferOperation) Type() types.OpType { return types.OverrideTransferOpType }