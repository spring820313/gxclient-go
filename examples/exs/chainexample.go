package exs

import (
	gxc "gxclient-go"
	"gxclient-go/core/types"
	"math/rand"
)

func Vote() error {
	client, err := gxc.NewClient(testNetHttp)
	if err != nil {
		return err
	}

	cali123IDActiveKey := "5K1ravLd96RmdTf8rcGVhtJFS7hqvihwurfV3cdao8QVnkoEkXq"
	accounts := []string{"w1", "w2"}
	err = client.Vote(cali123IDActiveKey, "spring123", accounts, "" , "GXC")

	return err
}

func Transfer() error {
	client, _ := gxc.NewClient(testNetHttp)

	account, err := client.Database.GetAccountByName("spring123")
	println(account, err)

	account456, err := client.Database.GetAccountByName("spring456")
	println(account456)

	cali4888arr, err := client.Database.LookupAccounts("spring123", 2)

	cali123ID := cali4888arr["spring123"]

	cali4889arr, err := client.Database.LookupAccounts("spring456", 2)
	cali456ID := cali4889arr["spring456"]

	objs := make([]types.ObjectID, 1)
	objs[0] = types.MustParseObjectID("1.3.1")
	assets, err := client.Database.LookupAssetSymbols("GXC")

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

	var keys []*types.PublicKey
	for k := range account.Active.KeyAuths {
		keys = append(keys, k)
	}

	memo.From = *keys[0]
	memoKey := account456.Options.MemoKey
	memo.To = memoKey
	memo.Nonce = types.UInt64(rand.Int63())
	if len(message) > 0 {
		memo.Encrypt(fromPrivKey, message)
	} else {
		memo = nil
	}

	return  client.Transfer(cali123IDActiveKey, from, to, amount, fee, memo)
}

func ChainExample()  {
	client, err := gxc.NewClient(testNetHttp)
	println(err)

	chainId, err := client.Database.GetChainID()
	println(chainId)

	dynamicGlobalProperties, err := client.Database.GetDynamicGlobalProperties()
	println(dynamicGlobalProperties)

	block, err := client.Database.GetBlock(12345)
	println(block)

	objs := make([]types.ObjectID, 2)
	objs[0] = types.MustParseObjectID("1.3.1")
	objs[1] = types.MustParseObjectID("1.3.2")
	objects, err := client.Database.GetObjects(objs[0], objs[1])
	println(objects)


}
