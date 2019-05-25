package types


type VestingBalances []VestingBalance

type VestingBalance struct {
	ID      GrapheneID    `json:"id"`
	Balance AssetAmount   `json:"balance"`
	Owner   GrapheneID    `json:"owner"`
	Policy  VestingPolicy `json:"policy"`
}
