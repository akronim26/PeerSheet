package utils

import (
	"encoding/base64"
	"errors"
	"os"

	"github.com/libp2p/go-libp2p/core/crypto"
	"github.com/libp2p/go-libp2p/core/peer"
)

var (
	PublicKey crypto.PubKey
	PeerID    peer.ID
)

func initKeysFromEnv() error {
	privKeyGen := os.Getenv("PRIVATE_KEY_GEN")
	if privKeyGen == "" {
		return errors.New("PRIVATE_KEY_GEN environment variable is not set")
	}

	data, err := base64.StdEncoding.DecodeString(privKeyGen)
	if err != nil {
		return errors.New("failed to decode private key generator string: " + err.Error())
	}

	privKey, err := crypto.UnmarshalPrivateKey(data)
	if err != nil {
		return errors.New("failed to generate private key: " + err.Error())
	}

	PublicKey = privKey.GetPublic()

	peerID, err := peer.IDFromPublicKey(PublicKey)
	if err != nil {
		return errors.New("failed to generate peer id: " + err.Error())
	}
	PeerID = peerID
	return nil
}

func GetPublicKey() (crypto.PubKey, error) {
	err := initKeysFromEnv()
	if err != nil {
		return nil, err
	}
	return PublicKey, nil
}

func GetPeerID() (peer.ID, error) {
	err := initKeysFromEnv() 
	if err != nil {
		return "", err
	}
	return PeerID, nil
}