package ops

import (
	"gxclient-go/core/transaction"
	"gxclient-go/core/types"
)

func NewAccountUpdateOperation(from types.GrapheneID, fee types.AssetAmount, newOptions *types.AccountOptions, Owner, Active *types.Authority) *AccountUpdateOperation {
	opFee := types.OperationFee{&fee}
	op := &AccountUpdateOperation{
		OperationFee:   opFee,
		Account:        from,
		NewOptions:     newOptions,
		Owner:			Owner,
		Active:         Active,
		Extensions: 	[]types.AccountUpdateExtensions{},
	}

	return op
}

type AccountUpdateOperation struct {
	types.OperationFee
	Account    types.GrapheneID              `json:"account"`
	Active     *types.Authority              `json:"active,omitempty"`
	NewOptions *types.AccountOptions         `json:"new_options,omitempty"`
	Owner      *types.Authority              `json:"owner,omitempty"`
	Extensions []types.AccountUpdateExtensions `json:"extensions"`
}

func (p AccountUpdateOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.Encode(int8(p.Type()))
	enc.Encode(p.Fee)
	enc.Encode(p.Account)
	enc.Encode(p.Owner != nil)
	if p.Owner != nil {
		enc.Encode(p.Owner)
	}
	enc.Encode(p.Active != nil)
	if p.Active != nil {
		enc.Encode(p.Active)
	}
	enc.Encode(p.NewOptions != nil)
	if p.NewOptions != nil {
		enc.Encode(p.NewOptions)
	}
	enc.EncodeUVarint(0)

	return enc.Err()
}

func (op *AccountUpdateOperation) Type() types.OpType { return types.AccountUpdateOpType }
