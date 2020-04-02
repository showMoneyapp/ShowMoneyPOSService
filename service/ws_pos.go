package service

import (
	"encoding/json"
	"errors"
	"github.com/ShowPay/ShowMoneyPosService/model"
)

type Ws_showPOS struct {}

//Parsing WSPaymentRequest
func (sh *Ws_showPOS) MsgToWSPaymentRequest(msg string, req *model.WsPaymentRequestReq) error  {
	err := json.Unmarshal([]byte(msg), req)
	if err != nil {
		return errors.New("Json data parsed failed.")
	}
	return nil
}

//Parsing WSPaymentACK
func (sh *Ws_showPOS) MsgToWSPaymentACK(msg string, req *model.WSPaymentACKReq) error  {
	err := json.Unmarshal([]byte(msg), req)
	if err != nil {
		return errors.New("Json data parsed failed.")
	}
	return nil
}

//get PaymentRequest
func (sh *Ws_showPOS) GetPaymentReq(req *model.WsPaymentRequestReq, resp *model.WsResponse) error {
	if req.Outputs == nil || len(req.Outputs) == 0 {
		return errors.New("outputs is empty.")
	}
	if req.Network == "" || len(req.Network) == 0 {
		return errors.New("Network is empty.")
	}
	if req.WalletId == "" || len(req.WalletId) == 0 {
		return errors.New("WalletId is empty.")
	}
	if req.DeviceId == "" || len(req.DeviceId) == 0 {
		return errors.New("DeviceId is empty.")
	}
	if req.PaymentUrl == "" || len(req.PaymentUrl) == 0 {
		return errors.New("PaymentUrl is empty.")
	}
	if req.CreationTimestamp  == 0 {
		return errors.New("CreationTimestamp is empty.")
	}
	if req.ExpirationTimestamp  == 0 {
		return errors.New("ExpirationTimestamp is empty.")
	}
	for _, v := range req.Outputs {
		if v.Amount < 0 || v.Script == "" || len(v.Script) == 0 {
			return errors.New("Output.Amount or Script is empty.")
		}
	}

	//**************************************************************************
	//TODO Temporarily send by pos machine
	//Forward based on wallet server url
	switch req.PaymentUrl {
	default:
		break
	}
	//**************************************************************************

	resp.Code = 200
	resp.Message = "success"
	return nil
}
