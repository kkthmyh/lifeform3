package response

import "math/big"

type WalletLogin struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		Username    string   `json:"username"`
		AccessToken string   `json:"access_token"`
		TokenType   string   `json:"token_type"`
		ExpiresIn   *big.Int `json:"expires_in"`
		Email       string   `json:"email"`
		UserID      *big.Int `json:"user_id"`
		Header      string   `json:"header"`
		Salt        string   `json:"salt"`
	} `json:"data"`
}

type EasyMintAvatar struct {
	Status int `json:"status"`
	Result struct {
		Condition struct {
			MintRule         string   `json:"mintRule"`
			UdIndex          *big.Int `json:"udIndex"`
			StakeErc20       string   `json:"stakeErc20"`
			StakeErc20Amount string   `json:"stakeErc20Amount"`
			CostErc20        string   `json:"costErc20"`
			CostErc20Amount  int      `json:"costErc20Amount"`
			LimitTimes       int      `json:"limitTimes"`
			MintType         int      `json:"mintType"`
			SignCode         string   `json:"signCode"`
			WlSignature      string   `json:"wlSignature"`
		} `json:"condition"`
		DataSignature struct {
			R         string `json:"r"`
			S         string `json:"s"`
			V         int    `json:"v"`
			Signature string `json:"signature"`
		} `json:"dataSignature"`
		Domain struct {
			Name              string `json:"name"`
			Version           string `json:"version"`
			ChainID           int    `json:"chainId"`
			VerifyingContract string `json:"verifyingContract"`
		} `json:"domain"`
	} `json:"result"`
}
