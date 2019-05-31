# gxclient-go
A client to interact with gxchain implemented in go
<p>
 <a href='javascript:;'>
   <img width="300px" src='https://raw.githubusercontent.com/gxchain/gxips/master/assets/images/task-gxclient.png'/>
 </a>
 <a href='javascript:;'>
   <img width="300px" src='https://raw.githubusercontent.com/gxchain/gxips/master/assets/images/task-gxclient-en.png'/>
 </a>
</p> 

# Supported Versions
go version go1.9.2

# Install

You can install this library via go get:

```
go get github.com/spring820313/gxclient-go
```
# APIs
- [x] [Keypair API](#keypair-api)
- [x] [Chain API](#chain-api)
- [x] [Faucet API](#faucet-api)
- [x] [Account API](#account-api)
- [x] [Asset API](#asset-api)
- [x] [Contract API](#contract-api)


## Constructors

``` go
//init GXClient
const (
   	testNetHttp = "https://testnet.gxchain.org"
  	testNetWss = "wss://testnet.gxchain.org" )

client, err := gxc.NewClient(testNetHttp)

```

## Keypair API

``` go
//generate key pair locally
func NewKeyPair(seed string) (*KeyPair, error)
//export public key from private key
func (p PrivateKey) PublicKey() *PublicKey
//check if public key is valid
func NewPublicKeyFromString(key string) (*PublicKey, error)
//check if private key is valid
func NewPrivateKeyFromWif(wifPrivateKey string) (*PrivateKey, error)
```

## Chain API

``` go
//get current blockchain id
func (api *API) GetChainID() (*string, error)
//get dynamic global properties 
func (api *API) GetDynamicGlobalProperties() (*DynamicGlobalProperties, error)
//get object
func (api *API) GetObjects(assets ...types.ObjectID) ([]json.RawMessage, error)
//get objects
func (api *API) GetObjects(assets ...types.ObjectID) ([]json.RawMessage, error)
// get block by block height
func (api *API) GetBlock(blockNum uint32) (*Block, error)
//send transfer request to entryPoint node
func (client *Client) Transfer(key string, from, to types.ObjectID, amount, fee types.AssetAmount, memo *types.Memo) error
//vote for accounts
func (client *Client) Vote(key string, from string, accounts []string,  proxyAccount string, feeAsset string) error
//broadcast transaction
func (client *Client) broadcast(stx *sign.SignedTransaction) error
```

## Faucet API

``` go
//register gxchain account
func (client *Client) Register(account, activeKey, ownerKey, memoKey, faucet string) (interface{}, error)
```
## Account API

``` go
// get account info by account name
func (api *API) GetAccountByName(account string) (*types.Account, error)
//get account_ids by public key
func (api *API) GetAccountsByPublicKeys(publicKeys ...string) (*[][]string, error)
//get account balances by account name
func (api *API) GetNamedAccountBalances(account string, assets ...types.ObjectID) ([]*types.AssetAmount, error)
```

## Asset API

``` go
//get asset info by symbol
func (api *API) GetAssets(symbols ...string) (*[]Asset, error)
```

## Contract API

``` go
// call smart contract method
func (client *Client) CallContract(key string, from string, contractName, method string, parameters interface{}, amount *types.AssetAmount, feeAsset string) error
// create smart contract method
func (client *Client) CreateContract(key string, from string, contractName, code, abi, vmType, vmVersion, feeAsset string) error
// update smart contract method
func (client *Client) UpdateContract(key string, from string, contractName, newOwner, code, abi, feeAsset string) error
//get contract table by contract_name
func (api *API) GetContractTable(account string) (*[]types.Table, error)
//get contract abi by contract_name
func (api *API) GetContractABI(account string) (*types.Abi, error)
//get contract table objects
func (api *API) GetTableObjects(contractName, tableName string, params TableRowsParams) (*interface{}, error) 
```

# Usage

```go
client, err := gxc.NewClient(testNetHttp)
println(err)

//getAccount(account_name) account, err := client.Database.GetAccountByName("spring123")
println(account)

wif := "5K1rav***" priv,_ := types.NewPrivateKeyFromWif(wif)
pubk := priv.PublicKey()
strpub := pubk.String()
println(strpub)

//getAccountByPublicKey(publicKey) accounts, err := client.Database.GetAccountsByPublicKeys("GXC7vr8Wre4UJgJz7H7GmYYGW7NEe6sxdmGhZPDUnHwmKnATrEBu9",
   "GXC7HnJw47kyv7hmHaEwQr1eiuhTHhG7LqySSVdExkmFCHbfgjn2w")
println(accounts)
```

For more examples, please refer to the examples directory.

# Other

- It's very welcome for developers to translate this project into different programing languages
- We are looking forward to your pull requests
