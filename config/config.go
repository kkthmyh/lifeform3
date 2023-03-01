package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

var ConfigInfo *Config

type Config struct {
	RpcConfig struct {
		BscRpc string
	}

	MintConfig struct {
		ContractAddress string
		AffAddress      string
		Gender          string
	}
}

func InitConfig() {
	viper.SetConfigFile("./config.yaml")
	//配置文件路径
	viper.AddConfigPath(".")
	//配置文件名字，不需要带后缀
	viper.SetConfigName("config")
	//配置文件类型
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("配置文件读取失败: %s", err))
	}

	config := new(Config)
	config.RpcConfig.BscRpc = viper.GetString("rpcConfig.bscRpc")
	config.MintConfig.ContractAddress = viper.GetString("mintConfig.contractAddress")
	config.MintConfig.AffAddress = viper.GetString("mintConfig.affAddress")
	config.MintConfig.Gender = viper.GetString("mintConfig.Gender")
	ConfigInfo = config
	log.Println("配置文件加载成功......")
}
