package model

type WsPaymentRequestReq struct {
	Network             string      `json:"network"`             //比特币网络标识：默认"bitcoin"，测试为"test"
	Outputs             []*WsOutput `json:"outputs"`             // outputs数组
	WalletId            string      `json:"walletId"`            //支付方钱包唯一标识
	DeviceId            string      `json:"deviceId"`            //收款方钱包唯一标识 pos
	CreationTimestamp   int64       `json:"creationTimestamp"`   //订单生成时间戳
	ExpirationTimestamp int64       `json:"expirationTimestamp"` //订单期限时间戳
	Memo                string      `json:"memo"`                //备忘录
	PaymentUrl          string      `json:"paymentUrl"`          //wallet钱包服务器地址
	MerchantData        string      `json:"merchantData"`        //商家自定义信息
}


type WsOutput struct {
	Amount      int64  `json:"amount"`      //satoshi
	Script      string `json:"script"`      //脚本
	Description string `json:"description"` //描述
	TokenIndex int64 `json:"tokenIndex"` //Token的公认index
}

