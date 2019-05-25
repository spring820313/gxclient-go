package ops

import (
	"gxclient-go/core/types"
	"gxclient-go/core/transaction"
)

func NewBlindTransferOperation(inputs []types.BlindInput, outputs []types.BlindOutput) *BlindTransferOperation {
	op := &BlindTransferOperation{
		Inputs: 		inputs,
		Outputs:		outputs,
	}

	return op
}

type BlindTransferOperation struct {
	types.OperationFee
	Inputs 		[]types.BlindInput  `json:"inputs"`
	Outputs 	[]types.BlindOutput  `json:"outputs"`
}

func (p BlindTransferOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.Encode(int8(p.Type()))
	enc.Encode(p.Fee)

	enc.EncodeUVarint(uint64(len(p.Inputs)))
	for _, i := range p.Inputs {
		enc.Encode(i)
	}

	enc.EncodeUVarint(uint64(len(p.Outputs)))
	for _, o := range p.Outputs {
		enc.Encode(o)
	}

	return enc.Err()
}

func (op *BlindTransferOperation) Type() types.OpType { return types.BlindTransferOpType }