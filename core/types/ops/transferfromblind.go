package ops

import (
	"gxclient-go/core/types"
	"encoding/hex"
	"gxclient-go/core/transaction"
)

func NewTransferFromBlindOperation(to types.GrapheneID, amount types.AssetAmount, blindingFactor string, inputs []types.BlindInput) *TransferFromBlindOperation {
	op := &TransferFromBlindOperation{
		Amount: 		amount,
		To:				to,
		BlindingFactor:	blindingFactor,
		Inputs:			inputs,
	}

	return op
}

type TransferFromBlindOperation struct {
	types.OperationFee
	Amount     		types.AssetAmount `json:"amount"`
	To       		types.GrapheneID  `json:"to"`
	BlindingFactor  string       `json:"blinding_factor"`
	Inputs 		[]types.BlindInput  `json:"inputs"`
}

func (p TransferFromBlindOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.Encode(int8(p.Type()))
	enc.Encode(p.Fee)
	enc.Encode(p.Amount)
	enc.Encode(p.To)

	var codeHex []byte
	var err error
	if codeHex, err = hex.DecodeString(p.BlindingFactor); err != nil {
		return err
	}
	size := len([]byte(codeHex))
	encoder.EncodeUVarint(uint64(size))
	encoder.Encode([]byte(codeHex))

	enc.EncodeUVarint(uint64(len(p.Inputs)))
	for _, i := range p.Inputs {
		enc.Encode(i)
	}

	return enc.Err()
}

func (op *TransferFromBlindOperation) Type() types.OpType { return types.TransferFromBlindOpType }