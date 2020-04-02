package impl

import (
	"encoding/json"
	"errors"
	"github.com/ShowPay/ShowMoneyPosService/model"
)

type Ws_showPOS struct {}

//解析
func (sh *Ws_showPOS) MsgToWSPaymentRequest(msg string, req *model.WsPaymentRequestReq) error  {
	err := json.Unmarshal([]byte(msg), req)
	if err != nil {
		return errors.New("json数据解析失败")
	}
	return nil
}

//解析
func (sh *Ws_showPOS) MsgToWSPaymentACK(msg string, req *model.WSPaymentACKReq) error  {
	err := json.Unmarshal([]byte(msg), req)
	if err != nil {
		return errors.New("json数据解析失败")
	}
	return nil
}

//获取并保存PaymentRequest
func (sh *Ws_showPOS) GetAndCreatePaymentReq(req *model.WsPaymentRequestReq, resp *model.WsResponse) error {
	if req.Outputs == nil || len(req.Outputs) == 0 {
		return errors.New("outputs为空")
	}
	if req.Network == "" || len(req.Network) == 0 {
		return errors.New("Network为空")
	}
	if req.WalletId == "" || len(req.WalletId) == 0 {
		return errors.New("WalletId为空")
	}
	if req.DeviceId == "" || len(req.DeviceId) == 0 {
		return errors.New("DeviceId为空")
	}
	if req.PaymentUrl == "" || len(req.PaymentUrl) == 0 {
		return errors.New("PaymentUrl为空")
	}
	if req.CreationTimestamp  == 0 {
		return errors.New("CreationTimestamp为空")
	}
	if req.ExpirationTimestamp  == 0 {
		return errors.New("ExpirationTimestamp为空")
	}
	for _, v := range req.Outputs {
		if v.Amount < 0 || v.Script == "" || len(v.Script) == 0 {
			return errors.New("output.Amount或Script内容为空")
		}
	}

	//**************************************************************************
	//TODO 暂定有pos机做发送
	//根据钱包发送服务器地址，转发
	switch req.PaymentUrl {
	default:
		break
	}
	//**************************************************************************

	resp.Code = 200
	resp.Message = "success"
	return nil
}
