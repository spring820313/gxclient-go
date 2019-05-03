package database

import (
	"encoding/json"
	"gxclient-go/core/types"
)

type Asset struct {
	ID                 types.ObjectID `json:"id"`
	Symbol             string         `json:"symbol"`
	Precision          uint8          `json:"precision"`
	Issuer             string         `json:"issuer"`
	DynamicAssetDataID string         `json:"dynamic_asset_data_id"`
}

type BlockHeader struct {
	TransactionMerkleRoot string            `json:"transaction_merkle_root"`
	Previous              string            `json:"previous"`
	Timestamp             types.Time        `json:"timestamp"`
	Witness               string            `json:"witness"`
	Extensions            []json.RawMessage `json:"extensions"`
}

type Block struct {
	TransactionMerkleRoot string              `json:"transaction_merkle_root"`
	Previous              string              `json:"previous"`
	Timestamp             types.Time          `json:"timestamp"`
	Witness               string              `json:"witness"`
	Extensions            []json.RawMessage   `json:"extensions"`
	WitnessSignature      string              `json:"witness_signature"`
	Transactions          []types.Transaction `json:"transactions"`
}

type MarketTicker struct {
	Time          types.Time     `json:"time"`
	Base          types.ObjectID `json:"base"`
	Quote         types.ObjectID `json:"quote"`
	Latest        string         `json:"latest"`
	LowestAsk     string         `json:"lowest_ask"`
	HighestBid    string         `json:"highest_bid"`
	PercentChange string         `json:"percent_change"`
	BaseVolume    string         `json:"base_volume"`
	QuoteVolume   string         `json:"quote_volume"`
}

type LimitOrder struct {
	ID          types.ObjectID `json:"id"`
	Expiration  types.Time     `json:"expiration"`
	Seller      types.ObjectID `json:"seller"`
	ForSale     types.Suint64  `json:"for_sale"`
	DeferredFee uint64         `json:"deferred_fee"`
	SellPrice   types.Price    `json:"sell_price"`
}

type DynamicGlobalProperties struct {
	ID                             types.ObjectID `json:"id"`
	HeadBlockNumber                uint32         `json:"head_block_number"`
	HeadBlockID                    string         `json:"head_block_id"`
	Time                           types.Time     `json:"time"`
	CurrentWitness                 types.ObjectID `json:"current_witness"`
	NextMaintenanceTime            types.Time     `json:"next_maintenance_time"`
	LastBudgetTime                 types.Time     `json:"last_budget_time"`
	AccountsRegisteredThisInterval int            `json:"accounts_registered_this_interval"`
	DynamicFlags                   int            `json:"dynamic_flags"`
	RecentSlotsFilled              string         `json:"recent_slots_filled"`
	LastIrreversibleBlockNum       uint32         `json:"last_irreversible_block_num"`
	CurrentAslot                   int64          `json:"current_aslot"`
	WitnessBudget                  int64          `json:"witness_budget"`
	RecentlyMissedCount            int64          `json:"recently_missed_count"`
}

type Config struct {
	GrapheneSymbol               string `json:"GRAPHENE_SYMBOL"`
	GrapheneAddressPrefix        string `json:"GRAPHENE_ADDRESS_PREFIX"`
	GrapheneMinAccountNameLength uint8  `json:"GRAPHENE_MIN_ACCOUNT_NAME_LENGTH"`
	GrapheneMaxAccountNameLength uint8  `json:"GRAPHENE_MAX_ACCOUNT_NAME_LENGTH"`
	GrapheneMinAssetSymbolLength uint8  `json:"GRAPHENE_MIN_ASSET_SYMBOL_LENGTH"`
	GrapheneMaxAssetSymbolLength uint8  `json:"GRAPHENE_MAX_ASSET_SYMBOL_LENGTH"`
	GrapheneMaxShareSupply       string `json:"GRAPHENE_MAX_SHARE_SUPPLY"`
}

type AccountsMap map[string]types.ObjectID

func (o *AccountsMap) UnmarshalJSON(b []byte) error {
	out := make(map[string]types.ObjectID)

	// unmarshal array
	var arr []json.RawMessage
	if err := json.Unmarshal(b, &arr); err != nil {
		return err
	}

	var (
		key string
		obj types.ObjectID
	)

	for _, item := range arr {
		account := []interface{}{&key, &obj}
		if err := json.Unmarshal(item, &account); err != nil {
			return err
		}

		out[key] = obj
	}

	*o = out
	return nil
}

type Struct struct {
	Name	   						string `json:"name"`
	Base							string `json:"base"`
	Fields							[]Field `json:"fields"`
}

type Table struct {
	Name	   						string `json:"name"`
	IndexType						string `json:"index_type"`
	KeyNames						[]string `json:"key_names"`
	KeyTypes						[]string `json:"key_types"`
	Type							string `json:"type"`
}

type Field struct {
	Name	   						string `json:"name"`
	Type							string `json:"type"`
}

type Action struct {
	Name	   						string `json:"name"`
	Type							string `json:"type"`
	Payable							bool `json:"payable"`
}

type Abi struct {
	Version	   						string `json:"version"`
	Types							[]interface{} `json:"types"`
	Structs							[]Struct `json:"structs"`
	Actions							[]Action `json:"actions"`
	Tables							[]Table `json:"tables"`
	ErrorMessages					[]interface{} `json:"error_messages"`
	AbiExtensions					[]interface{} `json:"abi_extensions"`
}

type ContractAccountProperties struct {
	ID                             	types.ObjectID `json:"id"`
	MembershipExpirationDate	   	string `json:"membership_expiration_date"`
	Registrar						string `json:"registrar"`
	Referrer						string `json:"referrer"`
	LifetimeReferrer				string `json:"lifetimeReferrer"`
	NetworkFeePercentage			int64 `json:"networkFeePercentage"`
	LifetimeReferrerFeePercentage	int64 `json:"lifetime_referrer_fee_percentage"`
	ReferrerRewardsPercentage		int64 `json:"referrer_rewards_percentage"`
	Name 							string `json:"name"`
	Statistics 						string `json:"statistics"`
	WhitelistingAccounts			[]string `json:"whitelisting_accounts"`
	BlacklistingAccounts			[]string `json:"blacklisting_accounts"`
	WhitelistedAccounts				[]string `json:"whitelisted_accounts"`
	BlacklistedAccounts				[]string `json:"blacklisted_accounts"`
	XAbi							Abi `json:"abi"`
	VmType							string `json:"vm_type"`
	VmVersion						string `json:"vm_version"`
	Code							string `json:"code"`
	CodeVersion						string `json:"code_version"`
}


