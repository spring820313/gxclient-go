package types

import (
	"github.com/juju/errors"
	"gxclient-go/core/transaction"
)

type OperationEnvelopeHolder struct {
	Op operationTuple `json:"op"`
}

type OperationEnvelopeHolders []OperationEnvelopeHolder

func (p OperationEnvelopeHolders) MarshalTransaction(enc *transaction.Encoder) error {
	if err := enc.EncodeUVarint(uint64(len(p))); err != nil {
		return errors.Annotate(err, "encode length")
	}

	for _, op := range p {
		if err := enc.Encode(op.Op); err != nil {
			return errors.Annotate(err, "encode Op")
		}
	}

	return nil
}
