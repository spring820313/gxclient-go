package ops

import (
	"encoding/hex"
	"encoding/json"
	"gxclient-go/core/types"
	"gxclient-go/core/transaction"
)

func NewCreateContractOperation(from types.GrapheneID,  contractName string, code []byte, abi types.Abi, vmType, vmVersion string) *CreateContractOperation {
	op := &CreateContractOperation{
		Name:			contractName,
		Account:        from,
		VmType:     	vmType,
		VmVersion:		vmVersion,
		Code:     		string(code[:]),
		Abi:   			abi,
		Extensions: 	[]json.RawMessage{},
	}

	return op
}

type CreateContractOperation struct {
	types.OperationFee
	Name        string                 `json:"name"`
	Account    types.GrapheneID              `json:"account"`
	VmType 		string                 `json:"vm_type"`
	VmVersion	string			   	   `json:"vm_version"`
	Code 	    string	       		   `json:"code"`
	Abi		    types.Abi			   		   `json:"abi"`
	Extensions []json.RawMessage	   `json:"extensions"`
}

func (p CreateContractOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.Encode(int8(p.Type()))
	enc.Encode(p.Fee)
	enc.Encode(p.Name)
	enc.Encode(p.Account)
	enc.Encode(p.VmType)
	enc.Encode(p.VmVersion)
	var codeHex []byte
	var err error
	if codeHex, err = hex.DecodeString(p.Code); err != nil {
		return err
	}
	codeSize := len(codeHex)
	enc.EncodeUVarint(uint64(codeSize))
	enc.Encode(codeHex)
	enc.Encode(p.Abi)

	enc.EncodeUVarint(0)

	return enc.Err()
}

func (op *CreateContractOperation) Type() types.OpType { return types.CreateContractOpType }
