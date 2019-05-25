package ops

import (
	"gxclient-go/core/types"
	"gxclient-go/core/transaction"
)

func NewAssetClaimFeesOperation(issuer types.GrapheneID, amountToClaim types.AssetAmount) *AssetClaimFeesOperation {
	op := &AssetClaimFeesOperation{
		Issuer:			issuer,
		AmountToClaim: 	amountToClaim,
		Extensions: 	types.Extensions{},
	}

	return op
}

type AssetClaimFeesOperation struct {
	types.OperationFee
	Issuer      	types.GrapheneID  `json:"issuer"`
	AmountToClaim   types.AssetAmount `json:"amount_to_claim"`
	Extensions      types.Extensions `json:"extensions"`
}

func (p AssetClaimFeesOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.Encode(int8(p.Type()))
	enc.Encode(p.Fee)
	enc.Encode(p.Issuer)
	enc.Encode(p.AmountToClaim)
	enc.EncodeUVarint(0)

	return enc.Err()
}

func (op *AssetClaimFeesOperation) Type() types.OpType { return types.AssetClaimFeesOpType }