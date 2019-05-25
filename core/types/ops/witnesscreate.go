package ops

import (
	"gxclient-go/core/types"
	"gxclient-go/core/transaction"
)

func NewWitnessCreateOperation(witnessAccount types.GrapheneID, url string, blockSigningKey types.PublicKey) *WitnessCreateOperation {
	op := &WitnessCreateOperation{
		WitnessAccount:	witnessAccount,
		URL: 			url,
		BlockSigningKey:blockSigningKey,
	}

	return op
}

type WitnessCreateOperation struct {
	types.OperationFee
	WitnessAccount  types.GrapheneID `json:"witness_account"`
	URL             string           `json:"url"`
	BlockSigningKey types.PublicKey  `json:"block_signing_key"`
}

func (p WitnessCreateOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.Encode(int8(p.Type()))
	enc.Encode(p.Fee)
	enc.Encode(p.WitnessAccount)
	enc.Encode(p.URL)
	enc.Encode(p.BlockSigningKey)

	return enc.Err()
}

func (op *WitnessCreateOperation) Type() types.OpType { return types.WitnessCreateOpType }