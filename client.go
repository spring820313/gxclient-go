package gxclient_go

import (
	"encoding/json"
	oldHttp "net/http"
	"github.com/pkg/errors"
	"gxclient-go/api/database"
	"gxclient-go/api/history"
	"gxclient-go/api/login"
	"gxclient-go/api/networkbroadcast"
	"gxclient-go/core/caller"
	"gxclient-go/core/sign"
	"gxclient-go/ws/transport/websocket"
	"gxclient-go/core/types"
	"log"
	"time"
	"strings"
	"gxclient-go/http"
	"io/ioutil"
	"github.com/pquerna/ffjson/ffjson"
	"bytes"
	"sort"
	_ "encoding/hex"
	_ "gxclient-go/core/transaction"
	_ "fmt"
	_ "encoding/hex"
	"encoding/hex"
	"fmt"
	"gxclient-go/core/transaction"
	"gxclient-go/core/types/ops"
)

type Client struct {
	cc caller.CallCloser

	// Database represents database_api
	Database *database.API

	// NetworkBroadcast represents network_broadcast_api
	NetworkBroadcast *networkbroadcast.API

	// History represents history_api
	History *history.API

	// Login represents login_api
	Login *login.API

	chainID string
}

// NewClient creates a new RPC client
func NewClient(url string) (*Client, error) {
	// transport
	var cc caller.CallCloser
	var err error
	if strings.HasPrefix(url,"http") || strings.HasPrefix(url,"https") {
		cc, err = http.NewHttpTransport(url)
	} else {
		cc, err = websocket.NewTransport(url)
	}
	if err != nil {
		return nil, err
	}

	client := &Client{cc: cc}
	if strings.HasPrefix(url,"http") || strings.HasPrefix(url,"https") {
		client.Database = database.NewAPI(0, cc)
		chainID, err := client.Database.GetChainID()
		if err != nil {
			return nil, errors.Wrap(err, "failed to get chain ID")
		}
		client.chainID = *chainID
		client.History = history.NewAPI(1, cc)
		client.NetworkBroadcast = networkbroadcast.NewAPI(1, cc)
		return client, nil
	}


	// login
	loginAPI := login.NewAPI(cc)
	client.Login = loginAPI

	// database
	databaseAPIID, err := loginAPI.Database()
	if err != nil {
		return nil, err
	}
	client.Database = database.NewAPI(databaseAPIID, client.cc)

	// chain ID
	chainID, err := client.Database.GetChainID()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get chain ID")
	}
	client.chainID = *chainID

	// history
	historyAPIID, err := loginAPI.History()
	if err != nil {
		return nil, err
	}
	client.History = history.NewAPI(historyAPIID, client.cc)

	// network broadcast
	networkBroadcastAPIID, err := loginAPI.NetworkBroadcast()
	if err != nil {
		return nil, err
	}
	client.NetworkBroadcast = networkbroadcast.NewAPI(networkBroadcastAPIID, client.cc)

	return client, nil
}

// Close should be used to close the client when no longer needed.
// It simply calls Close() on the underlying CallCloser.
func (client *Client) Close() error {
	return client.cc.Close()
}


type Parameters struct{
	MaximumCommitteeCount int `json:"maximum_committee_count"`
	MaximumWitnessCount int `json:"maximum_witness_count"`
}

func Distinct(a types.Votes) (ret types.Votes) {
	for i := 0; i < len(a); i++ {
		if i > 0 && types.VoteIDComparator(a[i - 1], a[i]) == 0 {
			continue
		}
		ret = append(ret, a[i])
	}
	return ret
}

func (client *Client) GetSignedProxyTransferParams(from, to, proxy string,  amount types.AssetAmount, percentage uint16, memo string) (*ops.SignedProxyTransferParams, error) {
	var fromAccount *types.Account
	var err error
	if fromAccount, err = client.Database.GetAccountByName(from); err != nil {
		return nil, err
	}
	fromAccountId := fromAccount.ID

	var toAccount *types.Account
	if toAccount, err = client.Database.GetAccountByName(to); err != nil {
		return nil, err
	}
	toAccountId := toAccount.ID

	var proxyAccount *types.Account
	if proxyAccount, err = client.Database.GetAccountByName(proxy); err != nil {
		return nil, err
	}
	proxyAccountId := proxyAccount.ID

	props, err := client.Database.GetDynamicGlobalProperties()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get dynamic global properties")
	}
	expiration := props.Time.Add(10 * time.Minute)

	signedProxyTransferParams := ops.NewSignedProxyTransferParams(fromAccountId, toAccountId, proxyAccountId, amount, percentage, memo, types.Time{Time: &expiration})
	return signedProxyTransferParams, nil
}

func (client *Client) ProxyTransfer(key string, proxyMemo string,  requestParams  ops.SignedProxyTransferParams, feeAsset string) error {
	var err error
	var assets *[]database.Asset
	if assets, err = client.Database.GetAssets(feeAsset); err != nil {
		return err
	}

	if len(*assets) <= 0 {
		return errors.New("can't get fees")
	}
	feeAssetId := (*assets)[0].ID

	fee := types.AssetAmount{
		AssetID: feeAssetId,
		Amount:  0,
	}
	op := ops.NewProxyTransferOperation(proxyMemo, requestParams)
	op.Fee = fee

	fees, err := client.Database.GetRequiredFee([]types.Operation{op}, feeAssetId.String())
	if err != nil {
		log.Println(err)
		return errors.Wrap(err, "can't get fees")
	}

	op.Fee.Amount = fees[0].Amount


	stx, err := client.sign([]string{key}, op)
	if err != nil {
		return err
	}
	return client.broadcast(stx)
}

func (client *Client) UpdateContract(key string, from string, contractName, newOwner, code, abi, feeAsset string) error {
	var myAccount *types.Account
	var err error
	if myAccount, err = client.Database.GetAccountByName(from); err != nil {
		return err
	}
	myAccountId := myAccount.ID

	var assets *[]database.Asset
	if assets, err = client.Database.GetAssets(feeAsset); err != nil {
		return err
	}

	if len(*assets) <= 0 {
		return errors.New("can't get fees")
	}
	feeAssetId := (*assets)[0].ID

	var contract *types.ContractAccountProperties
	if contract, err = client.Database.GetContractAccountByName(contractName); err != nil {
		return err
	}
	contractId := types.NewGrapheneID(contract.ID.String())

	var newOwnerId *types.GrapheneID
	if len(newOwner) > 0 {
		var newAccount *types.Account
		var err error
		if newAccount, err = client.Database.GetAccountByName(newOwner); err != nil {
			return err
		}
		newOwnerId = &newAccount.ID
	}

	fee := types.AssetAmount{
		AssetID: feeAssetId,
		Amount:  0,
	}

	var codeHex []byte
	codeHex = []byte(code)

	var abiObj types.Abi
	if err := json.Unmarshal([]byte(abi), &abiObj); err != nil {
		return err
	}

	op := ops.NewUpdateContractOperation(myAccountId, newOwnerId, *contractId, codeHex, abiObj)
	op.Fee = &fee

	fees, err := client.Database.GetRequiredFee([]types.Operation{op}, feeAssetId.String())
	if err != nil {
		log.Println(err)
		return errors.Wrap(err, "can't get fees")
	}

	op.Fee.Amount = fees[0].Amount


	stx, err := client.sign([]string{key}, op)
	if err != nil {
		return err
	}
	return client.broadcast(stx)
}

func (client *Client) CreateContract(key string, from string, contractName, code, abi, vmType, vmVersion, feeAsset string) error {
	var myAccount *types.Account
	var err error
	if myAccount, err = client.Database.GetAccountByName(from); err != nil {
		return err
	}
	myAccountId := myAccount.ID

	var assets *[]database.Asset
	if assets, err = client.Database.GetAssets(feeAsset); err != nil {
		return err
	}

	if len(*assets) <= 0 {
		return errors.New("can't get fees")
	}
	feeAssetId := (*assets)[0].ID

	fee := types.AssetAmount{
		AssetID: feeAssetId,
		Amount:  0,
	}

	var codeHex []byte
	/*
	if codeHex, err = hex.DecodeString(code); err != nil {
		return err
	}
	*/

	codeHex = []byte(code)

	var abiObj types.Abi
	if err := json.Unmarshal([]byte(abi), &abiObj); err != nil {
		return err
	}

	op := ops.NewCreateContractOperation(myAccountId, contractName, codeHex, abiObj, vmType, vmVersion)
	op.Fee = &fee

	fees, err := client.Database.GetRequiredFee([]types.Operation{op}, feeAssetId.String())
	if err != nil {
		log.Println(err)
		return errors.Wrap(err, "can't get fees")
	}

	op.Fee.Amount = fees[0].Amount


	stx, err := client.sign([]string{key}, op)
	if err != nil {
		return err
	}
	return client.broadcast(stx)
}

func (client *Client) CallContract(key string, from string, contractName, method string, parameters interface{}, amount *types.AssetAmount, feeAsset string) error {
	var myAccount *types.Account
	var err error
	if myAccount, err = client.Database.GetAccountByName(from); err != nil {
		return err
	}
	myAccountId := myAccount.ID

	var contract *types.ContractAccountProperties
	if contract, err = client.Database.GetContractAccountByName(contractName); err != nil {
		return err
	}
	contractId := types.NewGrapheneID(contract.ID.String())

	var data string
	if data, err = client.Database.SerializeContractParams(contractName, method, parameters); err != nil {
		return err
	}

	var assets *[]database.Asset
	if assets, err = client.Database.GetAssets(feeAsset); err != nil {
		return err
	}

	if len(*assets) <= 0 {
		return errors.New("can't get fees")
	}
	feeAssetId := (*assets)[0].ID

	fee := types.AssetAmount{
		AssetID: feeAssetId,
		Amount:  0,
	}

	op := ops.NewCallContractOperation(myAccountId, *contractId, method, amount)
	op.Data = data
	op.Fee = &fee

	fees, err := client.Database.GetRequiredFee([]types.Operation{op}, feeAssetId.String())
	if err != nil {
		log.Println(err)
		return errors.Wrap(err, "can't get fees")
	}

	op.Fee.Amount = fees[0].Amount


	stx, err := client.sign([]string{key}, op)
	if err != nil {
		return err
	}
	return client.broadcast(stx)
}

func (client *Client) Vote(key string, from string, accounts []string,  proxyAccount string, feeAsset string) error {
	var myAccount *types.Account
	var err error
	if myAccount, err = client.Database.GetAccountByName(from); err != nil {
		return err
	}
	myAccountId := myAccount.ID
	println(myAccount)

	votingAccountId := types.NewGrapheneID("1.2.5")
	if len(proxyAccount) > 0 {
		var votingAccount *types.Account
		if votingAccount, err = client.Database.GetAccountByName(proxyAccount); err != nil {
			return err
		}
		votingAccountId = types.NewGrapheneID(votingAccount.ID.String())
	}
	println(votingAccountId)

	var voteIds []string
	if voteIds, err = client.Database.GetVoteIdsByAccounts(accounts...); err != nil {
		return err
	}
	println(voteIds)

	var assets *[]database.Asset
	if assets, err = client.Database.GetAssets(feeAsset); err != nil {
		return err
	}

	if len(*assets) <= 0 {
		return errors.New("can't get fees")
	}
	feeAssetId := (*assets)[0].ID

	objId := types.MustParseObjectID("2.0.0")
	var rawObjs []json.RawMessage
	if rawObjs, err = client.Database.GetObjects(objId); err != nil {
		return err
	}
	rawObj := rawObjs[0]
	println(rawObj)

	globalParams := struct {
		Parameters `json:"parameters"`
	}{}

	if err := json.Unmarshal(rawObj, &globalParams); err != nil {
		return err
	}

	newOptions := myAccount.Options
	for _, voteId := range voteIds {
		vote := types.NewVoteIDV2(voteId)
		newOptions.Votes = append(newOptions.Votes, *vote)
	}

	sort.Slice(newOptions.Votes, func(i, j int) bool {
		return types.VoteIDComparator(newOptions.Votes[i], newOptions.Votes[j]) == -1

	})
	newOptions.Votes = Distinct(newOptions.Votes)

	maximumCommitteeCount := globalParams.Parameters.MaximumCommitteeCount
	maximumWitnessCount := globalParams.Parameters.MaximumWitnessCount

	numCommitee := 0
	numWitness := 0

	for _, voteId := range newOptions.Votes {
		typ := voteId.GetType()
		if typ == 0 {
			numCommitee += 1
		}
		if typ == 1 {
			numWitness += 1
		}
	}

	if numWitness < maximumWitnessCount {
		newOptions.NumWitness = types.UInt16(numWitness)
	} else {
		newOptions.NumWitness = types.UInt16(maximumWitnessCount)
	}

	if numCommitee < maximumCommitteeCount {
		newOptions.NumCommittee = types.UInt16(numCommitee)
	} else {
		newOptions.NumCommittee = types.UInt16(maximumCommitteeCount)
	}
	newOptions.VotingAccount = *votingAccountId

	fee := types.AssetAmount{
		Amount: 0,
		AssetID: feeAssetId,
	}
	op := ops.NewAccountUpdateOperation(myAccountId, fee, &newOptions, nil, nil)
	fees, err := client.Database.GetRequiredFee([]types.Operation{op}, fee.AssetID.String())
	if err != nil {
		log.Println(err)
		return errors.Wrap(err, "can't get fees")
	}
	op.Fee.Amount = fees[0].Amount

	stx, err := client.sign([]string{key}, op)
	if err != nil {
		return err
	}
	return client.broadcast(stx)
}

func (client *Client) Register(account, activeKey, ownerKey, memoKey, faucet string) (interface{}, error) {
	accountInfo := types.RegisterAccountInfo{}
	accountInfo.ActiveKey = activeKey
	if len(ownerKey) > 0 {
		accountInfo.OwnerKey = ownerKey
	} else {
		accountInfo.OwnerKey = activeKey
	}
	if len(memoKey) > 0 {
		accountInfo.MemoKey = memoKey
	} else {
		accountInfo.MemoKey = activeKey
	}
	accountInfo.Name = account

	accountReg := types.RegisterAccount{}
	accountReg.Account = accountInfo

	decBuf := new(bytes.Buffer)
	enc := ffjson.NewEncoder(decBuf)

	if err := enc.Encode(&accountReg); err != nil {
		return  nil, errors.Wrap(err, "Encode")
	}

	c := &oldHttp.Client{
		Timeout: 10 * time.Second,
	}
	req, err := oldHttp.NewRequest("POST", faucet, decBuf)
	if err != nil {
		return nil, errors.Wrap(err, "NewRequest")
	}

	req.Close = true
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	resp, err := c.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "do request")
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	s1 := string(body[:])
	println(s1)

	var res interface{}
	if err := json.Unmarshal(body, &res); err != nil {
		return nil, err
	}

	return res, nil
}

// Transfer a certain amount of the given asset
func (client *Client) Transfer(key string, from, to types.ObjectID, amount, fee types.AssetAmount, memo *types.Memo) error {
	op := ops.NewTransferOperation(from, to, amount, fee, memo)

	fees, err := client.Database.GetRequiredFee([]types.Operation{op}, fee.AssetID.String())
	if err != nil {
		log.Println(err)
		return errors.Wrap(err, "can't get fees")
	}
	op.Fee.Amount = fees[0].Amount

	stx, err := client.sign([]string{key}, op)
	if err != nil {
		return err
	}
	return client.broadcast(stx)
}

func (client *Client) LimitOrderCreate(key string, seller types.ObjectID, fee, amToSell, minToRecive types.AssetAmount, expiration time.Duration, fillOrKill bool) (string, error) {
	props, err := client.Database.GetDynamicGlobalProperties()
	if err != nil {
		return "", errors.Wrap(err, "failed to get dynamic global properties")
	}

	op := &ops.LimitOrderCreateOperation{
		Fee:          fee,
		Seller:       seller,
		AmountToSell: amToSell,
		MinToReceive: minToRecive,
		Expiration:   types.NewTime(props.Time.Add(expiration)),
		FillOrKill:   fillOrKill,
		Extensions:   []json.RawMessage{},
	}

	fees, err := client.Database.GetRequiredFee([]types.Operation{op}, fee.AssetID.String())
	if err != nil {
		log.Println(err)
		return "", errors.Wrap(err, "can't get fees")
	}
	op.Fee.Amount = fees[0].Amount

	stx, err := client.sign([]string{key}, op)
	if err != nil {
		return "", err
	}
	result, err := client.broadcastSync(stx)
	if err != nil {
		return "", err
	}

	res := result.Trx["operation_results"]
	ops, ok := res.([]interface{})
	if !ok {
		return "", errors.New("invalid result format")
	}
	create_op, ok := ops[0].([]interface{})
	if !ok {
		return "", errors.New("invalid result format")
	}
	id, ok := create_op[1].(string)
	if !ok {
		return "", errors.New("invalid result format")
	}

	return id, err
}

func (client *Client) LimitOrderCancel(key string, feePayingAccount, order types.ObjectID, fee types.AssetAmount) error {
	op := &ops.LimitOrderCancelOperation{
		Fee:              fee,
		FeePayingAccount: feePayingAccount,
		Order:            order,
		Extensions:       []json.RawMessage{},
	}

	fees, err := client.Database.GetRequiredFee([]types.Operation{op}, fee.AssetID.String())
	if err != nil {
		log.Println(err)
		return errors.Wrap(err, "can't get fees")
	}
	op.Fee.Amount = fees[0].Amount

	stx, err := client.sign([]string{key}, op)
	if err != nil {
		return err
	}
	return client.broadcast(stx)
}

func (client *Client) sign(wifs []string, operations ...types.Operation) (*sign.SignedTransaction, error) {
	props, err := client.Database.GetDynamicGlobalProperties()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get dynamic global properties")
	}

	block, err := client.Database.GetBlock(props.LastIrreversibleBlockNum)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get block")
	}

	refBlockPrefix, err := sign.RefBlockPrefix(block.Previous)
	if err != nil {
		return nil, errors.Wrap(err, "failed to sign block prefix")
	}

	expiration := props.Time.Add(10 * time.Minute)
	stx := sign.NewSignedTransaction(&types.Transaction{
		RefBlockNum:    sign.RefBlockNum(props.LastIrreversibleBlockNum - 1&0xffff),
		RefBlockPrefix: refBlockPrefix,
		Expiration:     types.Time{Time: &expiration},
	})

	for _, op := range operations {
		stx.PushOperation(op)
	}


	var b bytes.Buffer
	x := transaction.NewEncoder(&b)

	if err := x.Encode(stx.Transaction); err != nil {
		return nil, nil
	}
	s := hex.EncodeToString(b.Bytes())
	fmt.Println(s)


	if err = stx.Sign(wifs, client.chainID); err != nil {
		return nil, errors.Wrap(err, "failed to sign the transaction")
	}

	return stx, nil
}

func (client *Client) broadcast(stx *sign.SignedTransaction) error {
	return client.NetworkBroadcast.BroadcastTransaction(stx.Transaction)
}

func (client *Client) broadcastSync(stx *sign.SignedTransaction) (*networkbroadcast.BroadcastResponse, error) {
	return client.NetworkBroadcast.BroadcastTransactionSynchronous(stx.Transaction)
}
