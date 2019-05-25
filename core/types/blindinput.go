package types

import (
	"encoding/hex"
	"gxclient-go/core/transaction"
)

type BlindInput struct {
	Commitment 		string  `json:"commitment"`
	Owner         	Authority  `json:"owner"`
}

func (p BlindInput) MarshalTransaction(encoder *transaction.Encoder) error {
	var codeHex []byte
	var err error
	if codeHex, err = hex.DecodeString(p.Commitment); err != nil {
		return err
	}
	size := len([]byte(codeHex))
	encoder.EncodeUVarint(uint64(size))
	encoder.Encode([]byte(codeHex))

	encoder.Encode(p.Owner)

	return nil
}
