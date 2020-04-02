package model

type WSPaymentACKReq struct {
	Payment *Payment `json:"payment"` //Payment. required.
	TxId    string   `json:"txId"`    //
	Memo    string   `json:"memo"`    //string. optional
	Error   int64    `json:"error"`   // number. optional.
}

type Payment struct {
	WalletId       string `json:"walletId"`       //string. required. wallet-device
	DeviceId       string `json:"deviceId"`       //string. required. pos-device
	MerchantData   string `json:"merchantData"`   //string. optional.
	TransactionHex string `json:"transactionHex"` //a hex-formatted (and fully-signed and valid) transaction. required.
	RefundTo       string `json:"refundTo"`       //string. paymail to send a refund to. optional.
	Memo           string `json:"memo"`           //string. optional.
}