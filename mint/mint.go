package mint

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"lifeform/client"
	"lifeform/config"
	"lifeform/contract"
	utils "lifeform/utils"
	"log"
	"math/big"
	"unsafe"
)

func MintAvatar721(priKey string, rule string, index *big.Int, stakeErc20, costErc20 string, signCode, wlSignature, signature string) (string, error) {

	ecdsa, err1 := utils.GenPublicKeyECDSA(priKey)
	privateKey, err2 := utils.GenPrivateKey(priKey)
	if err1 != nil || err2 != nil {
		log.Printf("Gen public key ecdsa or private key error")
		return "", err1
	}
	fromAddress := crypto.PubkeyToAddress(*ecdsa)
	// 获取nonce
	nonce, err3 := client.EthClient.PendingNonceAt(context.Background(), fromAddress)
	if err3 != nil {
		log.Println(err3)
		return "", err3
	}
	//log.Printf("=====================当前nonce为: [%v]=====================", nonce)

	auth, _ := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(56))
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(350000)
	// 注意这里单位是Wei
	auth.GasPrice = big.NewInt(5000000000)
	address := common.HexToAddress(config.ConfigInfo.MintConfig.ContractAddress)
	instance, err5 := contract.NewContract(address, client.EthClient)
	if err5 != nil {
		log.Println(err5)
		return "", err5
	}

	// 组装参数
	parameters := contract.ICartoonMintRuleMintRule{}
	parameters.MintRule = common.HexToAddress(rule)
	parameters.UdIndex = index
	parameters.StakeErc20 = common.HexToAddress(stakeErc20)
	parameters.StakeErc20Amount = big.NewInt(0)
	parameters.CostErc20 = common.HexToAddress(costErc20)
	parameters.CostErc20Amount = big.NewInt(0)
	parameters.LimitTimes = big.NewInt(1)
	parameters.MintType = big.NewInt(1)
	parameters.SignCode = *byte32(common.FromHex(signCode))
	parameters.WlSignature = common.FromHex(wlSignature)
	//log.Println(parameters)
	dataSignature := common.FromHex(signature)
	tx, err6 := instance.MintAvatar721(auth, parameters, dataSignature)
	if err6 != nil {
		//log.Printf("=====================调用合约方法出错===================")
		return "", err6
	}
	return tx.Hash().Hex(), nil
}

func byte32(s []byte) (a *[32]byte) {
	if len(a) <= len(s) {
		a = (*[len(a)]byte)(unsafe.Pointer(&s[0]))
	}
	return a
}
