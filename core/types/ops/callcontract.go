package ops

import (
	"gxclient-go/util"
	"gxclient-go/core/transaction"
	"encoding/json"
	"gxclient-go/core/types"
)

func NewCallContractOperation(from, contractId types.GrapheneID, methodName string, amount *types.AssetAmount) *CallContractOperation {
	op := &CallContractOperation{
		Account:        from,
		ContractId:     contractId,
		Amount:			amount,
		MethodName:     methodName,
		Extensions: 	[]json.RawMessage{},
	}

	return op
}


type CallContractOperation struct {
	types.OperationFee
	Account    types.GrapheneID              `json:"account"`
	ContractId types.GrapheneID              `json:"contract_id"`
	Amount	   *types.AssetAmount			   `json:"amount,omitempty"`
	MethodName string				   `json:"method_name"`
	Data string			   			   `json:"data"`
	Extensions []json.RawMessage	   `json:"extensions"`
}

func (p CallContractOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.Encode(int8(p.Type()))
	enc.Encode(p.Fee)
	enc.Encode(p.Account)
	enc.Encode(p.ContractId)
	enc.Encode(p.Amount != nil)
	if p.Amount != nil {
		enc.Encode(p.Amount)
	}
	encodeName := util.StringToName(p.MethodName)
	enc.Encode(encodeName)
	enc.Encode(p.Data)

	enc.EncodeUVarint(0)

	return enc.Err()
}

func (op *CallContractOperation) Type() types.OpType { return types.CallContractOpType }
