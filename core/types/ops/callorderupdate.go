package ops

import (
	"gxclient-go/core/transaction"
	"gxclient-go/core/types"
	"encoding/json"
)

func NewCallOrderUpdateOperation(fundingAccount types.ObjectID, fee, deltaCollateral, deltaDebt types.AssetAmount,) *CallOrderUpdateOperation {
	op := &CallOrderUpdateOperation{
		Fee:        		fee,
		FundingAccount:   	fundingAccount,
		DeltaCollateral:  	deltaCollateral,
		DeltaDebt: 			deltaDebt,
		Extensions: 		[]json.RawMessage{},
	}

	return op
}

// CallOrderUpdateOpType
type CallOrderUpdateOperation struct {
	Fee     		types.AssetAmount `json:"fee"`
	FundingAccount  types.ObjectID `json:"funding_account"`
	DeltaCollateral types.AssetAmount `json:"delta_collateral"`
	DeltaDebt    	types.AssetAmount `json:"delta_debt"`
	Extensions 		[]json.RawMessage `json:"extensions"`
}

func (op *CallOrderUpdateOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)

	enc.EncodeUVarint(uint64(op.Type()))
	enc.Encode(op.Fee)
	enc.Encode(op.FundingAccount)
	enc.Encode(op.DeltaCollateral)
	enc.Encode(op.DeltaDebt)
	enc.EncodeUVarint(0)

	return enc.Err()
}

func (op *CallOrderUpdateOperation) Type() types.OpType { return types.CallOrderUpdateOpType }
