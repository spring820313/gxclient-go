package ops

import (
	"gxclient-go/core/types"
	"gxclient-go/core/transaction"
)

func NewWitnessUpdateOperation(witness, witnessAccount types.GrapheneID, newURL string, newSigningKey *types.PublicKey) *WitnessUpdateOperation {
	op := &WitnessUpdateOperation{
		Witness: 		witness,
		WitnessAccount:	witnessAccount,
		NewURL: 		newURL,
		NewSigningKey:	newSigningKey,
	}

	return op
}

type WitnessUpdateOperation struct {
	types.OperationFee
	Witness        types.GrapheneID `json:"witness"`
	WitnessAccount types.GrapheneID `json:"witness_account"`
	NewURL         string           `json:"new_url,omitempty"`
	NewSigningKey  *types.PublicKey `json:"new_signing_key,omitempty"`
}

func (p WitnessUpdateOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.Encode(int8(p.Type()))
	enc.Encode(p.Fee)
	enc.Encode(p.Witness)
	enc.Encode(p.WitnessAccount)
	enc.Encode(p.WitnessAccount)
	enc.Encode(p.NewURL != "")
	if p.NewURL != "" {
		enc.Encode(p.NewURL)
	}
	enc.Encode(p.NewSigningKey != nil)
	if p.NewSigningKey != nil {
		enc.Encode(p.NewSigningKey)
	}

	return enc.Err()
}

func (op *WitnessUpdateOperation) Type() types.OpType { return types.WitnessUpdateOpType }
