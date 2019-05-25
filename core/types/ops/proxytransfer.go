package ops

import (
	"bytes"
	"encoding/hex"
	"gxclient-go/core/transaction"
	"encoding/json"
	"gxclient-go/core/types"
)

func NewSignedProxyTransferParams(from, to, proxyAccount types.GrapheneID,  amount types.AssetAmount,  percentage uint16, memo string, expiration types.Time) *SignedProxyTransferParams {
	op := &SignedProxyTransferParams{
		From:		    from,
		To:       		to,
		ProxyAccount:   proxyAccount,
		Amount:     	amount,
		Percentage:   	percentage,
		Memo:			memo,
		Expiration:		expiration,
	}

	return op
}

type SignedProxyTransferParams struct {
	From       		types.GrapheneID            `json:"from"`
	To    			types.GrapheneID            `json:"to"`
	ProxyAccount 	types.GrapheneID            `json:"proxy_account"`
	Amount 	    	types.AssetAmount	       	  `json:"amount"`
	Percentage		uint16			   	  `json:"percentage"`
	Memo            string	              `json:"memo"`
	Expiration 	    types.Time	       		  `json:"expiration"`
	Signatures		[]string			  `json:"signatures"`
}

func (p SignedProxyTransferParams) ToUnsignBytes(signed bool) []byte {
	var b bytes.Buffer
	encoder := transaction.NewEncoder(&b)

	encoder.Encode(p.From)
	encoder.Encode(p.To)
	encoder.Encode(p.ProxyAccount)
	encoder.Encode(p.Amount)
	encoder.Encode(p.Percentage)
	encoder.Encode(p.Memo)
	encoder.Encode(p.Expiration)
	if signed == false {
		encoder.EncodeUVarint(0)
	}

	return b.Bytes()
}


func (p SignedProxyTransferParams) MarshalTransaction(encoder *transaction.Encoder) error {
	unsigned := p.ToUnsignBytes(true)
	encoder.Encode(unsigned)
	sigSize := len(p.Signatures)
	encoder.EncodeUVarint(uint64(sigSize))
	for _, s := range p.Signatures {
		sb, _ := hex.DecodeString(s)
		encoder.Encode(sb)
	}

	return nil
}

func NewProxyTransferOperation(proxyMemo string,  requestParams  SignedProxyTransferParams) *ProxyTransferOperation {
	op := &ProxyTransferOperation{
		ProxyMemo:		proxyMemo,
		RequestParams:  requestParams,
		Extensions: 	[]json.RawMessage{},
	}

	return op
}

type ProxyTransferOperation struct {
	ProxyMemo           string             			`json:"proxy_memo"`
	Fee 				types.AssetAmount			   		`json:"fee"`
	RequestParams       SignedProxyTransferParams   `json:"request_params"`
	Extensions 			[]json.RawMessage	   		`json:"extensions"`
}

func (p ProxyTransferOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.Encode(int8(p.Type()))
	enc.Encode(p.ProxyMemo)
	enc.Encode(p.Fee)
	enc.Encode(p.RequestParams)
	enc.EncodeUVarint(0)

	return enc.Err()
}

func (op *ProxyTransferOperation) Type() types.OpType { return types.ProxyTransferOpType }
