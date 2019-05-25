package exs

import (
	gxc "gxclient-go"
	"gxclient-go/core/types"
)

const (
	testNetHttp = "https://testnet.gxchain.org"
	testNetWss = "wss://testnet.gxchain.org"
)

func AccountExample()  {
	client, err := gxc.NewClient(testNetHttp)
	println(err)

	//getAccount(account_name)
	account, err := client.Database.GetAccountByName("spring123")
	println(account)

	wif := "5K1ravLd96RmdTf8rcGVhtJFS7hqvihwurfV3cdao8QVnkoEkXq"
	priv,_ := types.NewPrivateKeyFromWif(wif)
	pubk := priv.PublicKey()
	strpub := pubk.String()
	println(strpub)

	//getAccountByPublicKey(publicKey)
	accounts, err := client.Database.GetAccountsByPublicKeys("GXC7vr8Wre4UJgJz7H7GmYYGW7NEe6sxdmGhZPDUnHwmKnATrEBu9",
		"GXC7HnJw47kyv7hmHaEwQr1eiuhTHhG7LqySSVdExkmFCHbfgjn2w")
	println(accounts)

	//getAccountBalances(account_name)
	objs := make([]types.ObjectID, 1)
	objs[0] = types.MustParseObjectID("1.3.1")
	o2, err := client.Database.GetNamedAccountBalances("spring123", objs[0])
	println(o2)

}
