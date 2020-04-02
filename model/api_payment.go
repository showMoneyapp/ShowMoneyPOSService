package model

type ApiPaymentReq struct {
	WalletId       string `json:"walletId"`       //string. required. wallet-device
	DeviceId       string `json:"deviceId"`       //string. required. pos-device
	MerchantData   string `json:"merchantData"`   //string. optional.
	TransactionHex string `json:"transactionHex"` //a hex-formatted (and fully-signed and valid) transaction. required.
	RefundTo       string `json:"refundTo"`       //string. paymail to send a refund to. optional.
	Memo           string `json:"memo"`           //string. optional.
}

type ApiPaymentResultResp struct {
	Transaction string `json:"transaction"` //broadcast txId.
	Code        int    `json:"code"`        //int.
	Error       string `json:"error"`       //broadcast error message
}