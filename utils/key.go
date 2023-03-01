package utils

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
)

func GenPrivateKey(priKey string) (*ecdsa.PrivateKey, error) {
	privateKey, err := crypto.HexToECDSA(priKey)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return privateKey, err
}

func GenPublicKeyECDSA(priKey string) (*ecdsa.PublicKey, error) {
	privateKey, err := GenPrivateKey(priKey)
	if err != nil {
		log.Printf("Gen private key fail")
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Printf("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	return publicKeyECDSA, nil

}
