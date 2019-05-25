package types

import (
	"gxclient-go/core/transaction"
	"github.com/pquerna/ffjson/ffjson"
	"encoding/json"
)

type TransferOperationFeeParameters struct {
	Fee   			uint64 `json:"fee"`
	PricePerKbyte   uint32 `json:"price_per_kbyte"`
}

func (p TransferOperationFeeParameters) MarshalTransaction(enc *transaction.Encoder) error {
	enc.Encode(uint8(FeeParametersTypeTransfer))
	enc.Encode(p.Fee)
	enc.Encode(p.PricePerKbyte)
	return nil
}

type LimitOrderCreateOperationFeeParameters struct {
	Fee   			uint64 `json:"fee"`
}

func (p LimitOrderCreateOperationFeeParameters) MarshalTransaction(enc *transaction.Encoder) error {
	enc.Encode(uint8(FeeParametersTypeLimitOrderCreate))
	enc.Encode(p.Fee)
	return nil
}

type LimitOrderCancelOperationFeeParameters struct {
	Fee   			uint64 `json:"fee"`
}

func (p LimitOrderCancelOperationFeeParameters) MarshalTransaction(enc *transaction.Encoder) error {
	enc.Encode(uint8(FeeParametersTypeLimitOrderCancel))
	enc.Encode(p.Fee)
	return nil
}

type CallOrderUpdateOperationFeeParameters struct {
	Fee   			uint64 `json:"fee"`
}

func (p CallOrderUpdateOperationFeeParameters) MarshalTransaction(enc *transaction.Encoder) error {
	enc.Encode(uint8(FeeParametersTypeCallOrderUpdate))
	enc.Encode(p.Fee)
	return nil
}

type FillOrderOperationFeeParameters struct {
}

func (p FillOrderOperationFeeParameters) MarshalTransaction(enc *transaction.Encoder) error {
	enc.Encode(uint8(FeeParametersTypeFillOrder))
	return nil
}

type AccountCreateOperationFeeParameters struct {
	BasicFee   			uint64 `json:"basic_fee"`
	PremiumFee   		uint64 `json:"premium_fee"`
	PricePerKbyte   	uint32 `json:"price_per_kbyte"`
}

func (p AccountCreateOperationFeeParameters) MarshalTransaction(enc *transaction.Encoder) error {
	enc.Encode(uint8(FeeParametersTypeAccountCreate))
	enc.Encode(p.BasicFee)
	enc.Encode(p.PremiumFee)
	enc.Encode(p.PricePerKbyte)
	return nil
}

type AccountUpdateOperationFeeParameters struct {
	Fee   				uint64 `json:"fee"`
	PricePerKbyte   	uint32 `json:"price_per_kbyte"`
}

func (p AccountUpdateOperationFeeParameters) MarshalTransaction(enc *transaction.Encoder) error {
	enc.Encode(uint8(FeeParametersTypeAccountUpdate))
	enc.Encode(p.Fee)
	enc.Encode(p.PricePerKbyte)
	return nil
}

type AccountWhitelistOperationFeeParameters struct {
	Fee   				int64 `json:"fee"`
}

func (p AccountWhitelistOperationFeeParameters) MarshalTransaction(enc *transaction.Encoder) error {
	enc.Encode(uint8(FeeParametersTypeAccountWhitelist))
	enc.Encode(p.Fee)
	return nil
}

type AccountUpgradeOperationFeeParameters struct {
	MembershipAnnualFee   				uint64 `json:"membership_annual_fee"`
	MembershipLifetimeFee   			uint64 `json:"membership_lifetime_fee"`
}

func (p AccountUpgradeOperationFeeParameters) MarshalTransaction(enc *transaction.Encoder) error {
	enc.Encode(uint8(FeeParametersTypeAccountUpgrade))
	enc.Encode(p.MembershipAnnualFee)
	enc.Encode(p.MembershipLifetimeFee)
	return nil
}

type AccountTransferOperationFeeParameters struct {
	Fee   				uint64 `json:"fee"`
}

func (p AccountTransferOperationFeeParameters) MarshalTransaction(enc *transaction.Encoder) error {
	enc.Encode(uint8(FeeParametersTypeAccountTransfer))
	enc.Encode(p.Fee)
	return nil
}

type AssetCreateOperationFeeParameters struct {
	Symbol3   			uint64 `json:"symbol3"`
	Symbol4   			uint64 `json:"symbol4"`
	LongSymbol   		uint64 `json:"long_symbol"`
	PricePerKbyte   	uint32 `json:"price_per_kbyte"`
}

func (p AssetCreateOperationFeeParameters) MarshalTransaction(enc *transaction.Encoder) error {
	enc.Encode(uint8(FeeParametersTypeAssetCreate))
	enc.Encode(p.Symbol3)
	enc.Encode(p.Symbol4)
	enc.Encode(p.LongSymbol)
	enc.Encode(p.PricePerKbyte)
	return nil
}

type AssetUpdateOperationFeeParameters struct {
	Fee   				uint64 `json:"fee"`
	PricePerKbyte   	uint32 `json:"price_per_kbyte"`
}

func (p AssetUpdateOperationFeeParameters) MarshalTransaction(enc *transaction.Encoder) error {
	enc.Encode(uint8(FeeParametersTypeAssetUpdate))
	enc.Encode(p.Fee)
	enc.Encode(p.PricePerKbyte)
	return nil
}

type AssetUpdateBitassetOperationFeeParameters struct {
	Fee   				uint64 `json:"fee"`
}

func (p AssetUpdateBitassetOperationFeeParameters) MarshalTransaction(enc *transaction.Encoder) error {
	enc.Encode(uint8(FeeParametersTypeAssetUpdateBitasset))
	enc.Encode(p.Fee)
	return nil
}

type AssetUpdateFeedProducersOperationFeeParameters struct {
	Fee   				uint64 `json:"fee"`
}

func (p AssetUpdateFeedProducersOperationFeeParameters) MarshalTransaction(enc *transaction.Encoder) error {
	enc.Encode(uint8(FeeParametersTypeAssetUpdateFeedProducers))
	enc.Encode(p.Fee)
	return nil
}

type AssetIssueOperationFeeParameters struct {
	Fee   				uint64 `json:"fee"`
	PricePerKbyte   	uint32 `json:"price_per_kbyte"`
}

func (p AssetIssueOperationFeeParameters) MarshalTransaction(enc *transaction.Encoder) error {
	enc.Encode(uint8(FeeParametersTypeAssetIssue))
	enc.Encode(p.Fee)
	enc.Encode(p.PricePerKbyte)
	return nil
}

type AssetReserveOperationFeeParameters struct {
	Fee   				uint64 `json:"fee"`
}

func (p AssetReserveOperationFeeParameters) MarshalTransaction(enc *transaction.Encoder) error {
	enc.Encode(uint8(FeeParametersTypeAssetReserve))
	enc.Encode(p.Fee)
	return nil
}

type AssetFundFeePoolOperationFeeParameters struct {
	Fee   				uint64 `json:"fee"`
}

func (p AssetFundFeePoolOperationFeeParameters) MarshalTransaction(enc *transaction.Encoder) error {
	enc.Encode(uint8(FeeParametersTypeAssetFundFeePool))
	enc.Encode(p.Fee)
	return nil
}

type AssetSettleOperationFeeParameters struct {
	Fee   				uint64 `json:"fee"`
}

func (p AssetSettleOperationFeeParameters) MarshalTransaction(enc *transaction.Encoder) error {
	enc.Encode(uint8(FeeParametersTypeAssetSettle))
	enc.Encode(p.Fee)
	return nil
}

type AssetGlobalSettleOperationFeeParameters struct {
	Fee   				uint64 `json:"fee"`
}

func (p AssetGlobalSettleOperationFeeParameters) MarshalTransaction(enc *transaction.Encoder) error {
	enc.Encode(uint8(FeeParametersTypeAssetGlobalSettle))
	enc.Encode(p.Fee)
	return nil
}

type AssetPublishFeedOperationFeeParameters struct {
	Fee   				uint64 `json:"fee"`
}

func (p AssetPublishFeedOperationFeeParameters) MarshalTransaction(enc *transaction.Encoder) error {
	enc.Encode(uint8(FeeParametersTypeAssetPublishFeed))
	enc.Encode(p.Fee)
	return nil
}

type WitnessCreateOperationFeeParameters struct {
	Fee   				uint64 `json:"fee"`
}

func (p WitnessCreateOperationFeeParameters) MarshalTransaction(enc *transaction.Encoder) error {
	enc.Encode(uint8(FeeParametersTypeWitnessCreate))
	enc.Encode(p.Fee)
	return nil
}

type WitnessUpdateOperationFeeParameters struct {
	Fee   				uint64 `json:"fee"`
}

func (p WitnessUpdateOperationFeeParameters) MarshalTransaction(enc *transaction.Encoder) error {
	enc.Encode(uint8(FeeParametersTypeWitnessUpdate))
	enc.Encode(p.Fee)
	return nil
}

type ProposalCreateOperationFeeParameters struct {
	Fee   				uint64 `json:"fee"`
	PricePerKbyte   	uint32 `json:"price_per_kbyte"`
}

func (p ProposalCreateOperationFeeParameters) MarshalTransaction(enc *transaction.Encoder) error {
	enc.Encode(uint8(FeeParametersTypeProposalCreate))
	enc.Encode(p.Fee)
	enc.Encode(p.PricePerKbyte)
	return nil
}

type ProposalUpdateOperationFeeParameters struct {
	Fee   				uint64 `json:"fee"`
	PricePerKbyte   	uint32 `json:"price_per_kbyte"`
}

func (p ProposalUpdateOperationFeeParameters) MarshalTransaction(enc *transaction.Encoder) error {
	enc.Encode(uint8(FeeParametersTypeProposalUpdate))
	enc.Encode(p.Fee)
	enc.Encode(p.PricePerKbyte)
	return nil
}

type ProposalDeleteOperationFeeParameters struct {
	Fee   				uint64 `json:"fee"`
}

func (p ProposalDeleteOperationFeeParameters) MarshalTransaction(enc *transaction.Encoder) error {
	enc.Encode(uint8(FeeParametersTypeProposalDelete))
	enc.Encode(p.Fee)
	return nil
}

type WithdrawPermissionCreateOperationFeeParameters struct {
	Fee   				uint64 `json:"fee"`
}

func (p WithdrawPermissionCreateOperationFeeParameters) MarshalTransaction(enc *transaction.Encoder) error {
	enc.Encode(uint8(FeeParametersTypeWithdrawPermissionCreate))
	enc.Encode(p.Fee)
	return nil
}

type WithdrawPermissionUpdateOperationFeeParameters struct {
	Fee   				uint64 `json:"fee"`
}

func (p WithdrawPermissionUpdateOperationFeeParameters) MarshalTransaction(enc *transaction.Encoder) error {
	enc.Encode(uint8(FeeParametersTypeWithdrawPermissionUpdate))
	enc.Encode(p.Fee)
	return nil
}

type WithdrawPermissionClaimOperationFeeParameters struct {
	Fee   				uint64 `json:"fee"`
	PricePerKbyte   	uint32 `json:"price_per_kbyte"`
}

func (p WithdrawPermissionClaimOperationFeeParameters) MarshalTransaction(enc *transaction.Encoder) error {
	enc.Encode(uint8(FeeParametersTypeWithdrawPermissionClaim))
	enc.Encode(p.Fee)
	enc.Encode(p.PricePerKbyte)
	return nil
}

type WithdrawPermissionDeleteOperationFeeParameters struct {
	Fee   				uint64 `json:"fee"`
}

func (p WithdrawPermissionDeleteOperationFeeParameters) MarshalTransaction(enc *transaction.Encoder) error {
	enc.Encode(uint8(FeeParametersTypeWithdrawPermissionDelete))
	enc.Encode(p.Fee)
	return nil
}

type CommitteeMemberCreateOperationFeeParameters struct {
	Fee   				uint64 `json:"fee"`
}

func (p CommitteeMemberCreateOperationFeeParameters) MarshalTransaction(enc *transaction.Encoder) error {
	enc.Encode(uint8(FeeParametersTypeCommitteeMemberCreate))
	enc.Encode(p.Fee)
	return nil
}

type CommitteeMemberUpdateOperationFeeParameters struct {
	Fee   				uint64 `json:"fee"`
}

func (p CommitteeMemberUpdateOperationFeeParameters) MarshalTransaction(enc *transaction.Encoder) error {
	enc.Encode(uint8(FeeParametersTypeCommitteeMemberUpdate))
	enc.Encode(p.Fee)
	return nil
}

type CommitteeMemberUpdateGlobalParametersOperationFeeParameters struct {
	Fee   				uint64 `json:"fee"`
}

func (p CommitteeMemberUpdateGlobalParametersOperationFeeParameters) MarshalTransaction(enc *transaction.Encoder) error {
	enc.Encode(uint8(FeeParametersTypeCommitteeMemberUpdateGlobalParameters))
	enc.Encode(p.Fee)
	return nil
}

type VestingBalanceCreateOperationFeeParameters struct {
	Fee   				uint64 `json:"fee"`
}

func (p VestingBalanceCreateOperationFeeParameters) MarshalTransaction(enc *transaction.Encoder) error {
	enc.Encode(uint8(FeeParametersTypeVestingBalanceCreate))
	enc.Encode(p.Fee)
	return nil
}

type VestingBalanceWithdrawOperationFeeParameters struct {
	Fee   				uint64 `json:"fee"`
}

func (p VestingBalanceWithdrawOperationFeeParameters) MarshalTransaction(enc *transaction.Encoder) error {
	enc.Encode(uint8(FeeParametersTypeVestingBalanceWithdraw))
	enc.Encode(p.Fee)
	return nil
}

type WorkerCreateOperationFeeParameters struct {
	Fee   				uint64 `json:"fee"`
}

func (p WorkerCreateOperationFeeParameters) MarshalTransaction(enc *transaction.Encoder) error {
	enc.Encode(uint8(FeeParametersTypeWorkerCreate))
	enc.Encode(p.Fee)
	return nil
}

type CustomOperationFeeParameters struct {
	Fee   				uint64 `json:"fee"`
	PricePerKbyte   	uint32 `json:"price_per_kbyte"`
}

func (p CustomOperationFeeParameters) MarshalTransaction(enc *transaction.Encoder) error {
	enc.Encode(uint8(FeeParametersTypeCustom))
	enc.Encode(p.Fee)
	enc.Encode(p.PricePerKbyte)
	return nil
}

type AssertOperationFeeParameters struct {
	Fee   				uint64 `json:"fee"`
}

func (p AssertOperationFeeParameters) MarshalTransaction(enc *transaction.Encoder) error {
	enc.Encode(uint8(FeeParametersTypeAssert))
	enc.Encode(p.Fee)
	return nil
}

type BalanceClaimOperationFeeParameters struct {

}

func (p BalanceClaimOperationFeeParameters) MarshalTransaction(enc *transaction.Encoder) error {
	enc.Encode(uint8(FeeParametersTypeBalanceClaim))
	return nil
}

type OverrideTransferOperationFeeParameters struct {
	Fee   				uint64 `json:"fee"`
	PricePerKbyte   	uint32 `json:"price_per_kbyte"`
}

func (p OverrideTransferOperationFeeParameters) MarshalTransaction(enc *transaction.Encoder) error {
	enc.Encode(uint8(FeeParametersTypeOverrideTransfer))
	enc.Encode(p.Fee)
	enc.Encode(p.PricePerKbyte)
	return nil
}

type TransferToBlindOperationFeeParameters struct {
	Fee   				uint64 `json:"fee"`
	PricePerKbyte   	uint32 `json:"price_per_kbyte"`
}

func (p TransferToBlindOperationFeeParameters) MarshalTransaction(enc *transaction.Encoder) error {
	enc.Encode(uint8(FeeParametersTypeTransferToBlind))
	enc.Encode(p.Fee)
	enc.Encode(p.PricePerKbyte)
	return nil
}

type BlindTransferOperationFeeParameters struct {
	Fee   				uint64 `json:"fee"`
	PricePerKbyte   	uint32 `json:"price_per_kbyte"`
}

func (p BlindTransferOperationFeeParameters) MarshalTransaction(enc *transaction.Encoder) error {
	enc.Encode(uint8(FeeParametersTypeBlindTransfer))
	enc.Encode(p.Fee)
	enc.Encode(p.PricePerKbyte)
	return nil
}

type TransferFromBlindOperationFeeParameters struct {
	Fee   				uint64 `json:"fee"`
}

func (p TransferFromBlindOperationFeeParameters) MarshalTransaction(enc *transaction.Encoder) error {
	enc.Encode(uint8(FeeParametersTypeTransferFromBlind))
	enc.Encode(p.Fee)
	return nil
}

type AssetSettleCancelOperationFeeParameters struct {

}

func (p AssetSettleCancelOperationFeeParameters) MarshalTransaction(enc *transaction.Encoder) error {
	enc.Encode(uint8(FeeParametersTypeAssetSettleCancel))
	return nil
}

type AssetClaimFeesOperationFeeParameters struct {
	Fee   				uint64 `json:"fee"`
}

func (p AssetClaimFeesOperationFeeParameters) MarshalTransaction(enc *transaction.Encoder) error {
	enc.Encode(uint8(FeeParametersTypeAssetClaimFees))
	enc.Encode(p.Fee)
	return nil
}


type FeeParameters struct {
	Type        FeeParametersType
	Data		interface{}
}

type FeeParameterses []FeeParameters

func (p FeeParameterses) MarshalTransaction(enc *transaction.Encoder) error {
	if err := enc.EncodeUVarint(uint64(len(p))); err != nil {
		return err
	}

	for _, ex := range p {
		if err := enc.Encode(ex); err != nil {
			return err
		}
	}

	return nil
}

func (p FeeParameters) MarshalJSON() ([]byte, error) {
	return ffjson.Marshal([]interface{}{
		p.Type,
		p.Data,
	})
}

func (p *FeeParameters) UnmarshalJSON(data []byte) error {
	raw := make([]json.RawMessage, 2)
	if err := ffjson.Unmarshal(data, &raw); err != nil {
		return err
	}

	if len(raw) != 2 {
		return ErrInvalidInputLength
	}

	if err := ffjson.Unmarshal(raw[0], &p.Type); err != nil {
		return err
	}

	switch p.Type {
	case FeeParametersTypeTransfer:
		p.Data = &TransferOperationFeeParameters{}
	case FeeParametersTypeLimitOrderCreate:
		p.Data = &LimitOrderCreateOperationFeeParameters{}
	case FeeParametersTypeLimitOrderCancel:
		p.Data = &LimitOrderCancelOperationFeeParameters{}
	case FeeParametersTypeCallOrderUpdate:
		p.Data = &CallOrderUpdateOperationFeeParameters{}
	case FeeParametersTypeFillOrder:
		p.Data = &FillOrderOperationFeeParameters{}
	case FeeParametersTypeAccountCreate:
		p.Data = &AccountCreateOperationFeeParameters{}
	case FeeParametersTypeAccountUpdate:
		p.Data = &AccountUpdateOperationFeeParameters{}
	case FeeParametersTypeAccountWhitelist:
		p.Data = &AccountWhitelistOperationFeeParameters{}
	case FeeParametersTypeAccountUpgrade:
		p.Data = &AccountUpgradeOperationFeeParameters{}
	case FeeParametersTypeAccountTransfer:
		p.Data = &AccountTransferOperationFeeParameters{}
	case FeeParametersTypeAssetCreate:
		p.Data = &AssetCreateOperationFeeParameters{}
	case FeeParametersTypeAssetUpdate:
		p.Data = &AssetUpdateOperationFeeParameters{}
	case FeeParametersTypeAssetUpdateBitasset:
		p.Data = &AssetUpdateBitassetOperationFeeParameters{}
	case FeeParametersTypeAssetUpdateFeedProducers:
		p.Data = &AssetUpdateFeedProducersOperationFeeParameters{}
	case FeeParametersTypeAssetIssue:
		p.Data = &AssetIssueOperationFeeParameters{}
	case FeeParametersTypeAssetReserve:
		p.Data = &AssetReserveOperationFeeParameters{}
	case FeeParametersTypeAssetFundFeePool:
		p.Data = &AssetFundFeePoolOperationFeeParameters{}
	case FeeParametersTypeAssetSettle:
		p.Data = &AssetSettleOperationFeeParameters{}
	case FeeParametersTypeAssetGlobalSettle:
		p.Data = &AssetGlobalSettleOperationFeeParameters{}
	case FeeParametersTypeAssetPublishFeed:
		p.Data = &AssetPublishFeedOperationFeeParameters{}
	case FeeParametersTypeWitnessCreate:
		p.Data = &WitnessCreateOperationFeeParameters{}
	case FeeParametersTypeWitnessUpdate:
		p.Data = &WitnessUpdateOperationFeeParameters{}
	case FeeParametersTypeProposalCreate:
		p.Data = &ProposalCreateOperationFeeParameters{}
	case FeeParametersTypeProposalUpdate:
		p.Data = &ProposalUpdateOperationFeeParameters{}
	case FeeParametersTypeProposalDelete:
		p.Data = &ProposalDeleteOperationFeeParameters{}
	case FeeParametersTypeWithdrawPermissionCreate:
		p.Data = &WithdrawPermissionCreateOperationFeeParameters{}
	case FeeParametersTypeWithdrawPermissionUpdate:
		p.Data = &WithdrawPermissionUpdateOperationFeeParameters{}
	case FeeParametersTypeWithdrawPermissionClaim:
		p.Data = &WithdrawPermissionClaimOperationFeeParameters{}
	case FeeParametersTypeWithdrawPermissionDelete:
		p.Data = &WithdrawPermissionDeleteOperationFeeParameters{}
	case FeeParametersTypeCommitteeMemberCreate:
		p.Data = &CommitteeMemberCreateOperationFeeParameters{}
	case FeeParametersTypeCommitteeMemberUpdate:
		p.Data = &CommitteeMemberUpdateOperationFeeParameters{}
	case FeeParametersTypeCommitteeMemberUpdateGlobalParameters:
		p.Data = &CommitteeMemberUpdateGlobalParametersOperationFeeParameters{}
	case FeeParametersTypeVestingBalanceCreate:
		p.Data = &VestingBalanceCreateOperationFeeParameters{}
	case FeeParametersTypeVestingBalanceWithdraw:
		p.Data = &VestingBalanceWithdrawOperationFeeParameters{}
	case FeeParametersTypeWorkerCreate:
		p.Data = &WorkerCreateOperationFeeParameters{}
	case FeeParametersTypeCustom:
		p.Data = &CustomOperationFeeParameters{}
	case FeeParametersTypeAssert:
		p.Data = &AssertOperationFeeParameters{}
	case FeeParametersTypeBalanceClaim:
		p.Data = &BalanceClaimOperationFeeParameters{}
	case FeeParametersTypeOverrideTransfer:
		p.Data = &OverrideTransferOperationFeeParameters{}
	case FeeParametersTypeTransferToBlind:
		p.Data = &TransferToBlindOperationFeeParameters{}
	case FeeParametersTypeBlindTransfer:
		p.Data = &BlindTransferOperationFeeParameters{}
	case FeeParametersTypeTransferFromBlind:
		p.Data = &TransferFromBlindOperationFeeParameters{}
	case FeeParametersTypeAssetSettleCancel:
		p.Data = &AssetSettleCancelOperationFeeParameters{}
	case FeeParametersTypeAssetClaimFees:
		p.Data = &AssetClaimFeesOperationFeeParameters{}
	}

	if err := ffjson.Unmarshal(raw[1], p.Data); err != nil {
		return err
	}

	return nil
}