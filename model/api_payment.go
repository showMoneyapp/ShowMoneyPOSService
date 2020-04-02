package model

type ApiPaymentReq struct {
	WalletId       string `json:"walletId"`       //钱包标识
	DeviceId       string `json:"deviceId"`       //Pos标识
	MerchantData   string `json:"merchantData"`   //商户自定义信息
	TransactionHex string `json:"transactionHex"` //txHex
	RefundTo       string `json:"refundTo"`       //退款地址
	Memo           string `json:"memo"`           //memo
}

type ApiPaymentACKResp struct {
	Payment *ApiPaymentReq `json:"payment"`
	TxId    string         `json:"txId"`
	Memo    string         `json:"memo"`
	Code    int64          `json:"code"`
	Message string         `json:"message"`
}