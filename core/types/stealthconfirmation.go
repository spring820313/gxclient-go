package types

import (
	"gxclient-go/core/transaction"
	"encoding/hex"
)

type StealthConfirmation struct {
	OneTimeKey      PublicKey `json:"one_time_key"`
	To   			*PublicKey `json:"to,omitempty"`
	EncryptedMemo 	string  `json:"encrypted_memo"`
}

func (p StealthConfirmation) MarshalTransaction(encoder *transaction.Encoder) error {
	encoder.Encode(p.OneTimeKey)
	encoder.Encode(p.To != nil)
	if p.To != nil {
		encoder.Encode(p.To)
	}

	var codeHex []byte
	var err error
	if codeHex, err = hex.DecodeString(p.EncryptedMemo); err != nil {
		return err
	}
	size := len([]byte(codeHex))
	encoder.EncodeUVarint(uint64(size))
	encoder.Encode([]byte(codeHex))

	return nil
}