package model

//type WsPosOutput struct {
//	Outputs       []*WsOutput      `json:"outputs"`
//	Devices *WsDevices `json:"devices"`
//	PaymentUrl  string           `json:"paymentUrl"` //walletServer的地址
//	Action        string           `json:"action"`
//	OrderDetail   *WsOrderDetail   `json:"orderDetail"`//其他金额交易时的详细情况
//}
//
//type WsOutput struct {
//	Amount      int64  `json:"amount"`      //satoshi
//	Script      string `json:"script"`      //脚本
//	Description string `json:"description"` //描述
//	TokenIndex int64 `json:"tokenIndex"` //Token的公认index
//}
//
//type WsDevices struct {
//	WalletMAC string `json:"walletMAC"` //
//	PosMAC    string `json:"posMAC"`    //
//}
//
//type WsOrderDetail struct {
//	Amount     string `json:"amount"`
//	AmountUnit string `json:"amountUnit"`
//	PosName    string `json:"posName"`
//	////Token新增
//	//AssetType string `json:"assetType"` //资产类型：BSV、CUR
//	//Symbol    string `json:"symbol"`    //Token币种类型：BTE、BTC
//	//SymbolId  string  `json:"symbolId"`  //Token币种ID
//	//FeeAndCost uint64 `json:"feeAndCost"`//使用token时合约的费用
//}
