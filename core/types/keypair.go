package types

import (
	"github.com/tyler-smith/go-bip39"
	"gxclient-go/util"
	"bytes"
	"github.com/btcsuite/btcutil/base58"
)

type KeyPair struct {
	PublicKey string
	PrivateKey string
	BrainKey string
}

func NewKeyPair(seed string) (*KeyPair, error) {
	kp := &KeyPair{}
	if len(seed) <= 0 {
		entropy, err := bip39.NewEntropy(256)
		if err != nil {
			return nil, err
		}

		mne, err := bip39.NewMnemonic(entropy)
		if err != nil {
			return nil, err
		}
		kp.BrainKey = mne
	} else {
		kp.BrainKey = seed
	}

	privateKeyBytes := util.Sha256([]byte(kp.BrainKey))
	bufs := bytes.Buffer{}
	bufs.WriteByte(byte(0x80))
	bufs.Write(privateKeyBytes)
	data := bufs.Bytes()

	hash256 := util.Sha256([]byte(data))
	hash256  = util.Sha256(hash256)
	bufs.Write(hash256[:4])
	data = bufs.Bytes()
	wif := base58.Encode(data)
	kp.PrivateKey = wif

	priv,_ := NewPrivateKeyFromWif(wif)
	pubk := priv.PublicKey()
	spub := pubk.String()
	kp.PublicKey = spub

	return kp, nil
}