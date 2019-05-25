package exs

import (
	gxc "gxclient-go"
	"gxclient-go/core/types"
)

func FaucetExample()  {
	kp, _ := types.NewKeyPair("")
	println(kp)

	client, err := gxc.NewClient(testNetHttp)
	println(err)

	res, err := client.Register("gxc789", kp.PublicKey, "", "", "https://testnet.faucet.gxchain.org/account/register")
	println(res)
}
