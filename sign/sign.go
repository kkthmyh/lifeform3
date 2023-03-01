package sign

import (
	sign "github.com/etaaa/Golang-Ethereum-Personal-Sign"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
)

func Sign(privateKey, msg string) (string, error) {
	//fmt.Println(msg)
	// 加载私钥
	priKey, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	// personal_sign
	signature, err := sign.PersonalSign(msg, priKey)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	return signature, nil
}
