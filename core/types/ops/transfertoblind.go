package ops

import (
	"gxclient-go/core/types"
	"gxclient-go/core/transaction"
	"encoding/hex"
)

func NewTransferToBlindOperation(from types.GrapheneID, amount types.AssetAmount, blindingFactor string, outputs []types.BlindOutput) *TransferToBlindOperation {
	op := &TransferToBlindOperation{
		Amount: 		amount,
		From:			from,
		BlindingFactor:	blindingFactor,
		Outputs:		outputs,
	}

	return op
}

type TransferToBlindOperation struct {
	types.OperationFee
	Amount     		types.AssetAmount `json:"amount"`
	From       		types.GrapheneID  `json:"from"`
	BlindingFactor  string       `json:"blinding_factor"`
	Outputs 		[]types.BlindOutput  `json:"outputs"`
}

func (p TransferToBlindOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.Encode(int8(p.Type()))
	enc.Encode(p.Fee)
	enc.Encode(p.Amount)
	enc.Encode(p.From)

	var codeHex []byte
	var err error
	if codeHex, err = hex.DecodeString(p.BlindingFactor); err != nil {
		return err
	}
	size := len([]byte(codeHex))
	encoder.EncodeUVarint(uint64(size))
	encoder.Encode([]byte(codeHex))

	enc.EncodeUVarint(uint64(len(p.Outputs)))
	for _, o := range p.Outputs {
		enc.Encode(o)
	}

	return enc.Err()
}

func (op *TransferToBlindOperation) Type() types.OpType { return types.TransferToBlindOpType }