package ops

import (
	"gxclient-go/core/types"
	"gxclient-go/core/transaction"
)

func NewDataTransactionCreateOperation(requestId, params, version  string, productId, requester types.GrapheneID,
		leagueId *types.GrapheneID, createDateTime types.Time) *DataTransactionCreateOperation {
	op := &DataTransactionCreateOperation{
		RequestId: requestId,
		ProductId:  productId,
		Version: version,
		Params:  params,
		Requester: requester,
		CreateDateTime:  createDateTime,
		LeagueId: leagueId,
		Extensions: types.Extensions{},
	}

	return op
}

type DataTransactionCreateOperation struct {
	RequestId    		string `json:"request_id"`
	ProductId 			types.GrapheneID  `json:"product_id"`
	Version	   			string	`json:"version"`
	Params 				string	`json:"params"`
	types.OperationFee
	Requester 			types.GrapheneID		`json:"requester"`
	CreateDateTime 		types.Time		`json:"create_date_time"`
	LeagueId 			*types.GrapheneID		`json:"league_id,omitempty"`
	Extensions      	types.Extensions `json:"extensions"`
}

func (p DataTransactionCreateOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.Encode(int8(p.Type()))
	enc.Encode(p.RequestId)
	enc.Encode(p.ProductId)
	enc.Encode(p.Version)
	enc.Encode(p.Params)
	enc.Encode(p.Fee)
	enc.Encode(p.Requester)
	enc.Encode(p.CreateDateTime)
	enc.Encode(p.LeagueId != nil)
	if p.LeagueId != nil {
		enc.Encode(p.LeagueId)
	}
	enc.EncodeUVarint(0)

	return enc.Err()
}

func (op *DataTransactionCreateOperation) Type() types.OpType { return types.DataTransactionCreateOpType }
