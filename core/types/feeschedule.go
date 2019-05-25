package types

import "gxclient-go/core/transaction"

type FeeSchedule struct {
	Parameters      FeeParameterses `json:"parameters"`
	Scale			uint32 `json:"scale"`
}

func (p FeeSchedule) MarshalTransaction(enc *transaction.Encoder) error {
	enc.Encode(uint8(FeeParametersTypeTransfer))
	enc.Encode(p.Parameters)
	enc.Encode(p.Scale)
	return nil
}