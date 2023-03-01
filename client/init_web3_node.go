package client

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"lifeform/config"
	"log"
)

var EthClient *ethclient.Client

func InitWeb3EtherClient() {
	// bsc main net
	client, err := ethclient.Dial(config.ConfigInfo.RpcConfig.BscRpc)
	if err != nil {
		log.Println(err)
	}
	EthClient = client
	log.Println("RPC初始化成功......")

}
