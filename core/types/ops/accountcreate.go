package ops

import (
	"gxclient-go/core/types"
	"gxclient-go/core/transaction"
)

func NewAccountCreateOperation(registrar, referrer types.GrapheneID,  referrerPercent types.UInt16,  owner, active types.Authority, name string, options types.AccountOptions) *AccountCreateOperation {
	op := &AccountCreateOperation{
		Registrar:		registrar,
		Referrer:  		referrer,
		ReferrerPercent:referrerPercent,
		Owner:			owner,
		Active:			active,
		Name:			name,
		Options: 		options,
		Extensions: 	[]types.AccountCreateExtensions{},
	}

	return op
}

type AccountCreateOperation struct {
	types.OperationFee
	Registrar       types.GrapheneID              `json:"registrar"`
	Referrer        types.GrapheneID              `json:"referrer"`
	ReferrerPercent types.UInt16                  `json:"referrer_percent"`
	Owner           types.Authority               `json:"owner"`
	Active          types.Authority               `json:"active"`
	Name            string                        `json:"name"`
	Options         types.AccountOptions          `json:"options"`
	Extensions      []types.AccountCreateExtensions `json:"extensions"`
}

func (p AccountCreateOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.Encode(int8(p.Type()))
	enc.Encode(p.Fee)
	enc.Encode(p.Registrar)
	enc.Encode(p.Referrer)
	enc.Encode(p.ReferrerPercent)
	enc.Encode(p.Name)
	enc.Encode(p.Owner)
	enc.Encode(p.Active)
	enc.Encode(p.Options)
	enc.EncodeUVarint(0)

	return enc.Err()
}

func (op *AccountCreateOperation) Type() types.OpType { return types.AccountCreateOpType }