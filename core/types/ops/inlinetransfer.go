package ops

import (
	"gxclient-go/core/types"
	"gxclient-go/core/transaction"
)

func NewInlineTransferOperation(from, to types.GrapheneID, amount types.AssetAmount, memo string) *InlineTransferOperation {
	op := &InlineTransferOperation{
		From:		from,
		To: 		to,
		Amount: 	amount,
		Memo: 		memo,
		Extensions: types.Extensions{},
	}

	return op
}

type InlineTransferOperation struct {
	types.OperationFee
	From      		types.GrapheneID  `json:"from"`
	To       		types.GrapheneID  `json:"to"`
	Amount     		types.AssetAmount `json:"amount"`
	Memo 			string `json:"memo"`
	Extensions      types.Extensions `json:"extensions"`
}

func (p InlineTransferOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.Encode(int8(p.Type()))
	enc.Encode(p.Fee)
	enc.Encode(p.From)
	enc.Encode(p.To)
	enc.Encode(p.Amount)
	enc.Encode(p.Memo)
	enc.EncodeUVarint(0)

	return enc.Err()
}

func (op *InlineTransferOperation) Type() types.OpType { return types.InlineTransferOpType }
