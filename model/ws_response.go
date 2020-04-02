package model

type WsResponse struct {
	Code    int64       `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

//type WsBroadcastResultResp struct {
//	TxId         string `json:"txId"`
//	Message      string `json:"message"`
//	MerchantData string `json:"merchantData"`
//}
//
//type WsPosConnectInfoResp struct {
//	GetPaymentRequestApi string `json:"getPaymentRequestApi"`
//	BroadcastApi         string `json:"broadcastApi"`
//	WalletPubKey         string `json:"walletPubKey"`
//	//Token新增
//	GetTokenUnSignTxMsgApi string `json:"getTokenUnSignTxMsgApi"`
//}