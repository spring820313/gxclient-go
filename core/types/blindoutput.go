package types

import (
	"encoding/hex"
	"gxclient-go/core/transaction"
)

type BlindOutput struct {
	Commitment 		string  `json:"commitment"`
	RangeProof      string  `json:"range_proof"`
	Owner         	Authority  `json:"owner"`
	StealthMemo     *StealthConfirmation `json:"stealth_memo,omitempty"`
}

func (p BlindOutput) MarshalTransaction(encoder *transaction.Encoder) error {
	var codeHex []byte
	var err error
	if codeHex, err = hex.DecodeString(p.Commitment); err != nil {
		return err
	}
	size := len([]byte(codeHex))
	encoder.EncodeUVarint(uint64(size))
	encoder.Encode([]byte(codeHex))

	if codeHex, err = hex.DecodeString(p.RangeProof); err != nil {
		return err
	}
	size = len([]byte(codeHex))
	encoder.EncodeUVarint(uint64(size))
	encoder.Encode([]byte(codeHex))

	encoder.Encode(p.Owner)
	encoder.Encode(p.StealthMemo != nil)
	if p.StealthMemo != nil {
		encoder.Encode(p.StealthMemo)
	}

	return nil
}
