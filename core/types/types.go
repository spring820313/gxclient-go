package types


import (
	"fmt"
	"gxclient-go/core/transaction"
)

var (
	ErrRPCClientNotInitialized      = fmt.Errorf("RPC client is not initialized")
	ErrNotImplemented               = fmt.Errorf("not implemented")
	ErrInvalidInputType             = fmt.Errorf("invalid input type")
	ErrInvalidInputLength           = fmt.Errorf("invalid input length")
	ErrInvalidPublicKey             = fmt.Errorf("invalid PublicKey")
	ErrInvalidAddress               = fmt.Errorf("invalid Address")
	ErrPublicKeyChainPrefixMismatch = fmt.Errorf("PublicKey chain prefix mismatch")
	ErrAddressChainPrefixMismatch   = fmt.Errorf("Address chain prefix mismatch")
	ErrInvalidChecksum              = fmt.Errorf("invalid checksum")
	ErrNoSigningKeyFound            = fmt.Errorf("no signing key found")
	ErrNoVerifyingKeyFound          = fmt.Errorf("no verifying key found")
	ErrInvalidDigestLength          = fmt.Errorf("invalid digest length")
	ErrInvalidPrivateKeyCurve       = fmt.Errorf("invalid PrivateKey curve")
	ErrCurrentChainConfigIsNotSet   = fmt.Errorf("current chain config is not set")
)

type Int8 int8
type UInt8 uint8
type UInt16 uint16
type UInt32 uint32
type UInt64 uint64
type Int64 int64

type AssetType Int8

func (num UInt64) MarshalTransaction(enc *transaction.Encoder) error {
	return enc.EncodeNumber(uint64(num))
}

const (
	AssetTypeUndefined AssetType = -1
	AssetTypeCoreAsset AssetType = iota
	AssetTypeUIA
	AssetTypeSmartCoin
	AssetTypePredictionMarket
)

type SpaceType Int8

const (
	SpaceTypeUndefined SpaceType = -1
	SpaceTypeProtocol  SpaceType = iota
	SpaceTypeImplementation
)

type ObjectType Int8

const (
	ObjectTypeUndefined ObjectType = -1
)

//for SpaceTypeProtocol
const (
	ObjectTypeBase ObjectType = iota + 1
	ObjectTypeAccount
	ObjectTypeAsset
	ObjectTypeForceSettlement
	ObjectTypeCommiteeMember
	ObjectTypeWitness
	ObjectTypeLimitOrder
	ObjectTypeCallOrder
	ObjectTypeCustom
	ObjectTypeProposal
	ObjectTypeOperationHistory
	ObjectTypeWithdrawPermission
	ObjectTypeVestingBalance
	ObjectTypeWorker
	ObjectTypeBalance
)

// for SpaceTypeImplementation
const (
	ObjectTypeGlobalProperty ObjectType = iota + 1
	ObjectTypeDynamicGlobalProperty
	ObjectTypeAssetDynamicData
	ObjectTypeAssetBitAssetData
	ObjectTypeAccountBalance
	ObjectTypeAccountStatistics
	ObjectTypeTransaction
	ObjectTypeBlockSummary
	ObjectTypeAccountTransactionHistory
	ObjectTypeBlindedBalance
	ObjectTypeChainProperty
	ObjectTypeWitnessSchedule
	ObjectTypeBudgetRecord
	ObjectTypeSpecialAuthority
)

type AccountCreateExtensionsType UInt8

const (
	AccountCreateExtensionsNullExt AccountCreateExtensionsType = iota
	AccountCreateExtensionsOwnerSpecial
	AccountCreateExtensionsActiveSpecial
	AccountCreateExtensionsBuyback
)

type SpecialAuthorityType UInt8

const (
	SpecialAuthorityTypeNoSpecial SpecialAuthorityType = iota
	SpecialAuthorityTypeTopHolders
)

type VestingPolicyType UInt8

const (
	VestingPolicyTypeLinear VestingPolicyType = iota
	VestingPolicyTypeCCD
)




