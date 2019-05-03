package tests

import (
	"gxclient-go/core/types"
	"github.com/stretchr/testify/require"
	"log"
	"testing"
	gxc "gxclient-go"
	"math/rand"
)

const (
	testNetHttp = "https://testnet.gxchain.org"
	testNetWss = "wss://testnet.gxchain.org"
)

func TestClient_Register(t *testing.T) {
	client, err := gxc.NewClient(testNetHttp)
	require.Nil(t, err)

	//contract, err := client.Database.GetContractAccountByName("bank")
	//println(contract)

	//accounts, err := client.Database.GetAccountsByPublicKeys("GXC7vr8Wre4UJgJz7H7GmYYGW7NEe6sxdmGhZPDUnHwmKnATrEBu9", "GXC7HnJw47kyv7hmHaEwQr1eiuhTHhG7LqySSVdExkmFCHbfgjn2w")
	//println(accounts)

	assets, err := client.Database.GetAssets("PPS")
	println(assets)

	account789, err := client.Database.GetAccountByName("gxc789")
	println(account789)

	kp, _ := types.NewKeyPair("")
	println(kp)

	res, err := client.Register("gxc789", kp.PublicKey, "", "", "https://testnet.faucet.gxchain.org/account/register")
	require.Nil(t, err)
	println(res)
}

func TestClient_Transfer(t *testing.T) {
	client, err := gxc.NewClient(testNetHttp)
	require.Nil(t, err)

	account, err := client.Database.GetAccountByName("spring123")
	println(account)

	account456, err := client.Database.GetAccountByName("spring456")
	println(account456)

	cali4888arr, err := client.Database.LookupAccounts("spring123", 2)
	require.Nil(t, err)

	log.Println(cali4888arr["spring123"])

	cali123ID := cali4888arr["spring123"]

	cali4889arr, err := client.Database.LookupAccounts("spring456", 2)
	require.Nil(t, err)
	cali456ID := cali4889arr["spring456"]

	objs := make([]types.ObjectID, 1)
	objs[0] = types.MustParseObjectID("1.3.1")
	o, err := client.Database.GetObjects(objs[0])
	str := string([]byte(o[0])[:])
	log.Println(str)

	o2, err := client.Database.GetNamedAccountBalances("spring123", objs[0])
	log.Println(o2)

	assets, err := client.Database.LookupAssetSymbols("GXC")
	require.Nil(t, err)

	cali123IDActiveKey := "5K1ravLd96RmdTf8rcGVhtJFS7hqvihwurfV3cdao8QVnkoEkXq"
	from := cali123ID
	to := cali456ID
	amount := types.AssetAmount{
		AssetID: assets[0].ID,
		Amount:  10,
	}
	fee := types.AssetAmount{
		AssetID: assets[0].ID,
		Amount:  0,
	}

	var memo = &types.Memo{}
	message := "12345"

	fromPrivKey, err := types.NewPrivateKeyFromWif(cali123IDActiveKey)
	require.Nil(t, err)

	//fromPubKey := fromPrivKey.PublicKey()

	var keys []*types.PublicKey
	for k := range account.Active.KeyAuths {
		keys = append(keys, k)
	}

	memo.From = *keys[0]
	memoKey := account456.Options.MemoKey
	memo.To = memoKey
	memo.Nonce = types.UInt64(rand.Int63())
	if len(message) > 0 {
		err := memo.Encrypt(fromPrivKey, message)
		require.Nil(t, err)
	} else {
		memo = nil
	}

	require.NoError(t, client.Transfer(cali123IDActiveKey, from, to, amount, fee, memo))
}
