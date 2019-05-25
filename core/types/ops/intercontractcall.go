package ops

import (
	"gxclient-go/core/types"
	"gxclient-go/core/transaction"
)

func NewInterContractCallOperation(senderContract, contractId types.GrapheneID, methodName, data string, amount types.AssetAmount) *InterContractCallOperation {
	op := &InterContractCallOperation{
		SenderContract: senderContract,
		ContractId:     contractId,
		Amount:			amount,
		MethodName:     methodName,
		Data:			data,
		Extensions: types.Extensions{},
	}

	return op
}

type InterContractCallOperation struct {
	types.OperationFee
	SenderContract    	types.GrapheneID `json:"sender_contract"`
	ContractId 			types.GrapheneID  `json:"contract_id"`
	Amount	   			types.AssetAmount	`json:"amount"`
	MethodName 			string		`json:"method_name"`
	Data 				string		`json:"data"`
	Extensions      	types.Extensions `json:"extensions"`
}

func (p InterContractCallOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.Encode(int8(p.Type()))
	enc.Encode(p.Fee)
	enc.Encode(p.SenderContract)
	enc.Encode(p.ContractId)
	enc.Encode(p.Amount)

	enc.Encode(p.MethodName)
	enc.Encode(p.Data)

	enc.EncodeUVarint(0)

	return enc.Err()
}

func (op *InterContractCallOperation) Type() types.OpType { return types.InterContractCallOpType }