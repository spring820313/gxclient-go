package ops

import (
	"gxclient-go/core/types"
	"gxclient-go/core/transaction"
)

func NewAssertOperation(feePayingAccount types.GrapheneID, predicates types.Predicates, requiredAuths types.GrapheneIDs) *AssertOperation {
	op := &AssertOperation{
		FeePayingAccount:	feePayingAccount,
		Predicates:			predicates,
		RequiredAuths:		requiredAuths,
		Extensions:			[]types.AccountCreateExtensions{},
	}

	return op
}

type AssertOperation struct {
	types.OperationFee
	FeePayingAccount    	types.GrapheneID `json:"fee_paying_account"`
	Predicates    			types.Predicates `json:"predicates"`
	RequiredAuths 			types.GrapheneIDs `json:"required_auths"`
	Extensions      		[]types.AccountCreateExtensions `json:"extensions"`
}

func (p AssertOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.Encode(int8(p.Type()))
	enc.Encode(p.Fee)
	enc.Encode(p.FeePayingAccount)
	enc.Encode(p.Predicates)
	enc.Encode(p.RequiredAuths)
	enc.EncodeUVarint(0)
	return enc.Err()
}

func (op *AssertOperation) Type() types.OpType { return types.AssertOpType }
