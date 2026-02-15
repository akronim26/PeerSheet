package utils

import (
	"encoding/base64"
	"errors"
	"os"

	"github.com/libp2p/go-libp2p/core/crypto"
	"github.com/libp2p/go-libp2p/core/peer"
)

func LoadIdentityFromEnv() (crypto.PrivKey, peer.ID, error) {
	privKeyGen := os.Getenv("PRIVATE_KEY_GEN")
	if privKeyGen == "" {
		return nil, "", errors.New("PRIVATE_KEY_GEN environment variable is not set")
	}

	data, err := base64.StdEncoding.DecodeString(privKeyGen)
	if err != nil {
		return nil, "", errors.New("failed to decode private key generator string: " + err.Error())
	}

	privKey, err := crypto.UnmarshalPrivateKey(data)
	if err != nil {
		return nil, "", errors.New("failed to generate private key: " + err.Error())
	}

	pubKey := privKey.GetPublic()

	peerID, err := peer.IDFromPublicKey(pubKey)
	if err != nil {
		return nil, "", errors.New("failed to generate peer id: " + err.Error())
	}
	return privKey, peerID, nil
}