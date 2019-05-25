package exs

import (
	gxc "gxclient-go"
)

func AssetExample()  {
	client, err := gxc.NewClient(testNetHttp)
	println(err)

	assets, err := client.Database.GetAssets("PPS")
	println(assets)
}
