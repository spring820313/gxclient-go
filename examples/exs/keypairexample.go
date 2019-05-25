package exs

import "gxclient-go/core/types"

func KeypairExample() error  {
	kp, _ := types.NewKeyPair("")
	println(kp)

	priv,_ := types.NewPrivateKeyFromWif(kp.PrivateKey)
	pubk := priv.PublicKey()
	pub := pubk.String()
	println(pub)

	var err error
	var pr *types.PrivateKey
	if pr, err = types.NewPrivateKeyFromWif("5J34qxxxx..."); err != nil {
		return err
	}
	println(pr)

	var pu *types.PublicKey
	if pu, err = types.NewPublicKeyFromString("GXC6K35Bajw29N4fjP4XADHtJ7bEj2xHJ8CoY2P2s1igXTB5oMBhR"); err != nil {
		return err
	}
	println(pu)

	return nil
}
