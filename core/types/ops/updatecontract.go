package ops

import (
	"encoding/hex"
	"encoding/json"
	"gxclient-go/core/types"
	"gxclient-go/core/transaction"
)

func NewUpdateContractOperation(owner types.GrapheneID,  newOwner *types.GrapheneID,  contract types.GrapheneID, code []byte, abi types.Abi) *UpdateContractOperation {
	op := &UpdateContractOperation{
		Owner:		    owner,
		NewOwner:       newOwner,
		Contract:     	contract,
		Code:     		string(code[:]),
		Abi:   			abi,
		Extensions: 	[]json.RawMessage{},
	}

	return op
}

type UpdateContractOperation struct {
	types.OperationFee
	Owner       types.GrapheneID             `json:"owner"`
	NewOwner    *types.GrapheneID            `json:"new_owner,omitempty"`
	Contract 	types.GrapheneID             `json:"contract"`
	Code 	    string	       		   `json:"code"`
	Abi		   types.Abi			   		   `json:"abi"`
	Extensions []json.RawMessage	   `json:"extensions"`
}

func (p UpdateContractOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.Encode(int8(p.Type()))
	enc.Encode(p.Fee)
	enc.Encode(p.Owner)
	enc.Encode(p.NewOwner != nil)
	if p.NewOwner != nil {
		enc.Encode(p.NewOwner)
	}
	enc.Encode(p.Contract)
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

func (op *UpdateContractOperation) Type() types.OpType { return types.UpdateContractOpType }
