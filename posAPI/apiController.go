package posAPI

import (
	"github.com/ShowPay/ShowMoneyPosService/model"
	"github.com/ShowPay/ShowMoneyPosService/service"
	"net/http"
)

//Broadcast Payment
func BroadcastPayment(w http.ResponseWriter, r *http.Request)  {
	req := &model.ApiPaymentReq{}
	resp := &model.ApiPaymentResultResp{}
	err := pares(r, req)
	if err != nil {
		ResponseError(w, err.Error())
		return
	}
	if err = new(service.Api_showPOS).BroadcastPayment(req, resp);err != nil{
		ResponseError(w, err.Error())
		return
	}
	ResponseSuccess(w, resp)
}