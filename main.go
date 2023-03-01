package main

import (
	"fmt"
	"lifeform/client"
	"lifeform/config"
	"lifeform/http"
	"lifeform/mint"
	"lifeform/sign"
	"lifeform/utils"
	"log"
	"time"
)

func main() {
	//初始化配置文件
	config.InitConfig()
	//初始化Eth client
	client.InitWeb3EtherClient()
	http.InitHttpClient()
	// 加载私钥
	privateKeyList := utils.ReadCSV("./data/privateKey.csv")
	// 加载需要mint的地址
	addressList := utils.ReadCSV("./data/address.csv")

	for i := 0; i < len(addressList); i++ {
		// mint一次后休眠一会防止一个ip多次请求被查女巫
		time.Sleep(time.Second * 3)
		var privateKey = privateKeyList[i]
		var address = addressList[i]
		var affAddress = config.ConfigInfo.MintConfig.AffAddress
		var gender = config.ConfigInfo.MintConfig.Gender

		signStr, err := sign.Sign(privateKey, fmt.Sprintf("address=%s,chain_id=56", address))
		//log.Println(signStr)
		if err != nil {
			log.Printf("地址:%s失败,签名失败,开始下个地址mint......", addressList[i])
			continue
		}
		// 模拟网页端签名登陆，其实没卵用，但是不进行次步操作，无法mint
		_, err = http.WalletLogin("https://api-v2.lifeform.cc/api/v1/wallet_login", address, signStr, 56)
		if err != nil {
			log.Printf("地址:%s失败,登陆失败,开始下个地址mint......", addressList[i])
			continue
		}
		// 获取mint需要的相关参数
		res, err := http.EasyMintAvatar("https://pandora.lifeform.cc/lifeform_bsc_prod/api/avatarCartoon/easyMintAvatar", address, affAddress, gender)
		if err != nil {
			log.Printf("地址:%s失败,获取mint参数失败或该地址已经免费mint过,开始下个地址mint......", addressList[i])
			continue
		}
		rule := res.Result.Condition.MintRule
		index := res.Result.Condition.UdIndex
		stakeErc20 := res.Result.Condition.StakeErc20
		costErc20 := res.Result.Condition.CostErc20
		signCode := res.Result.Condition.SignCode
		wlSignature := res.Result.Condition.WlSignature
		signature := res.Result.DataSignature.Signature

		if res.Result.Condition.LimitTimes == 0 {
			log.Printf("地址:%s失败,剩余免费mint次数为0,开始下个地址mint......", addressList[i])
			continue
		}

		// free mint
		tx, err := mint.MintAvatar721(privateKey, rule, index, stakeErc20, costErc20, signCode, wlSignature, signature)
		if err != nil {
			log.Printf("地址:%s失败,调用合约失败,开始下个地址mint......", addressList[i])
			continue
		}
		log.Printf("地址:%s成功,tx%s", addressList[i], tx)
	}
}
