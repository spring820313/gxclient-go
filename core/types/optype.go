package types

type OpType uint16

const (
	TransferOpType OpType = iota
	LimitOrderCreateOpType
	LimitOrderCancelOpType
	CallOrderUpdateOpType
	FillOrderOpType
	AccountCreateOpType
	AccountUpdateOpType
	AccountWhitelistOpType
	AccountUpgradeOpType
	AccountTransferOpType
	AssetCreateOpType
	AssetUpdateOpType
	AssetUpdateBitassetOpType
	AssetUpdateFeedProducersOpType
	AssetIssueOpType
	AssetReserveOpType
	AssetFundFeePoolOpType
	AssetSettleOpType
	AssetGlobalSettleOpType
	AssetPublishFeedOpType
	WitnessCreateOpType
	WitnessUpdateOpType
	ProposalCreateOpType
	ProposalUpdateOpType
	ProposalDeleteOpType
	WithdrawPermissionCreateOpType
	WithdrawPermissionUpdateOpType
	WithdrawPermissionClaimOpType
	WithdrawPermissionDeleteOpType
	CommitteeMemberCreateOpType
	CommitteeMemberUpdateOpType
	CommitteeMemberUpdateGlobalParametersOpType
	VestingBalanceCreateOpType
	VestingBalanceWithdrawOpType
	WorkerCreateOpType
	CustomOpType
	AssertOpType
	BalanceClaimOpType
	OverrideTransferOpType
	TransferToBlindOpType
	BlindTransferOpType
	TransferFromBlindOpType
	AssetSettleCancelOpType
	AssetClaimFeesOpType
	FbaDistributeOperationOpType
	AccountUpgradeMerchantOpType
	AccountUpgradeDatasourceOpType
	StaleDataMarketCategoryCreateOpType
	StaleDataMarketCategoryUpdateOpType
	StaleFreeDataProductCreateOpType
	StaleFreeDataProductUpdateOpType
	StaleLeagueDataProductCreateOpType
	StaleLeagueDataProductUpdateOpType
	StaleLeagueCreateOpType
	StaleLeagueUpdateOpType
	DataTransactionCreateOpType
	DataTransactionUpdateOpType
	DataTransactionPayOpType
	AccountUpgradeDataTransactionMemberOpType
	DataTransactionDatasourceUploadOpType
	DataTransactionDatasourceValidateErrorOpType
	DataMarketCategoryCreateOpType
	DataMarketCategoryUpdateOpType
	FreeDataProductCreateOpType
	FreeDataProductUpdateOpType
	LeagueDataProductCreateOpType
	LeagueDataProductUpdateOpType
	LeagueCreateOpType
	LeagueUpdateOpType
	DatasourceCopyrightClearOpType
	DataTransactionComplainOpType
	BalanceLockOpType
	BalanceUnlockOpType
	ProxyTransferOpType = 73
	CreateContractOpType = 74
	CallContractOpType = 75
	UpdateContractOpType = 76
	TrustNodePledgeWithdrawOpType
	InlineTransferOpType
	InterContractCallOpType
)
