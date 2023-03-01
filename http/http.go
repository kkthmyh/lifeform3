package http

import (
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"lifeform/response"
	"log"
)

var client *resty.Client

func InitHttpClient() {
	client = resty.New()
	log.Println("HttpClient初始化成功......")
}

func WalletLogin(url, address, sign string, chainId uint) (*response.WalletLogin, error) {
	var login *response.WalletLogin
	m := map[string]interface{}{
		"address":  address,
		"chain_id": chainId,
		"sign":     sign,
	}
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(m).
		Post(url)

	if err != nil {
		return nil, err
	}
	err1 := json.Unmarshal(resp.Body(), &login)
	if err1 != nil {
		return nil, err1
	}

	return login, nil

}

func EasyMintAvatar(url, address, affAddress, gender string) (*response.EasyMintAvatar, error) {
	var login *response.EasyMintAvatar
	m := map[string]interface{}{
		"gender":     gender,
		"address":    address,
		"affAddress": affAddress,
	}
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(m).
		Post(url)

	if err != nil {
		return nil, err
	}
	err1 := json.Unmarshal(resp.Body(), &login)
	if err1 != nil {
		return nil, err1
	}

	return login, nil

}
