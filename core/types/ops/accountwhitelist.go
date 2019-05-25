package ops

import (
	"gxclient-go/core/transaction"
	"gxclient-go/core/types"
)

func NewAccountWhitelistOperation(accountToList, authorizingAccount types.GrapheneID,  newListing types.UInt8) *AccountWhitelistOperation {
	op := &AccountWhitelistOperation{
		AccountToList:		accountToList,
		AuthorizingAccount: authorizingAccount,
		NewListing:			newListing,
		Extensions: 		types.Extensions{},
	}

	return op
}

type AccountWhitelistOperation struct {
	types.OperationFee
	AccountToList      types.GrapheneID `json:"account_to_list"`
	AuthorizingAccount types.GrapheneID `json:"authorizing_account"`
	NewListing         types.UInt8      `json:"new_listing"`
	Extensions         types.Extensions `json:"extensions"`
}

func (p AccountWhitelistOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.Encode(int8(p.Type()))
	enc.Encode(p.Fee)
	enc.Encode(p.AccountToList)
	enc.Encode(p.AuthorizingAccount)
	enc.Encode(p.NewListing)
	enc.EncodeUVarint(0)

	return enc.Err()
}

func (op *AccountWhitelistOperation) Type() types.OpType { return types.AccountWhitelistOpType }