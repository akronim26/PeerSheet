package utils

import (
	"encoding/base64"
	"fmt"
	"os"

	"github.com/libp2p/go-libp2p/core/crypto"
	"github.com/libp2p/go-libp2p/core/peer"
)

func LoadIdentityFromEnv() (crypto.PrivKey, peer.ID, error) {
	privKeyGen := os.Getenv("PRIVATE_KEY_GEN")
	if privKeyGen == "" {
		return nil, "", fmt.Errorf("PRIVATE_KEY_GEN environment variable is not set")
	}

	data, err := base64.StdEncoding.DecodeString(privKeyGen)
	if err != nil {
		return nil, "", fmt.Errorf("failed to decode private key generator string: %w", err)
	}

	privKey, err := crypto.UnmarshalPrivateKey(data)
	if err != nil {
		return nil, "", fmt.Errorf("failed to generate private key: %w", err)
	}

	pubKey := privKey.GetPublic()

	peerID, err := peer.IDFromPublicKey(pubKey)
	if err != nil {
		return nil, "", fmt.Errorf("failed to generate peer id: %w", err)
	}
	return privKey, peerID, nil
}