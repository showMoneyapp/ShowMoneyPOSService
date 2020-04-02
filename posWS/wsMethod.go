package posWS

import (
	"fmt"
	"github.com/ShowPay/ShowMoneyPosService/model"
	"github.com/ShowPay/ShowMoneyPosService/service"
	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"net"
)

//Push message heartbeat back
func SendHeartBeat(c net.Conn) {
	wsData := &model.WsData{
		M: model.HEART_BEAT,
		C: model.WS_CODE_HEART_BEAT,
		D: "Heart beat back",
	}
	SendMsgToConn(c, wsData)
}

//get paymentRequest from ws
func GetPaymentRequest(c net.Conn, wsItemMap *WsItemMap, msg string) error {
	req := &model.WsPaymentRequestReq{}
	resp := &model.WsResponse{}
	err := new(service.Ws_showPOS).MsgToWSPaymentRequest(msg, req)
	if err != nil {
		fmt.Println("MsgToWSPaymentRequest err:", err)
		return err
	}
	//Caching conn
	wsItemMap.SetConn(req.DeviceId, c)

	err = new(service.Ws_showPOS).GetPaymentReq(req, resp)
	if err != nil {
		fmt.Println("GetAndCreatePaymentReq err:", err)
		return err
	}
	fmt.Println(resp)
	wsData := &model.WsData{
		M: model.WS_RESPONSE_SUCCESS,
		C: model.WS_CODE_SEND_SUCCESS,
		D: "send success",
	}
	SendMsgToConn(c, wsData)
	return nil
}

//send PaymentACK to POS
func NotifyPaymentACKToPos(c net.Conn, wsItemMap *WsItemMap, msg string) {
	req := &model.WSPaymentACKReq{}
	err := new(service.Ws_showPOS).MsgToWSPaymentACK(msg, req)
	if err != nil {
		fmt.Println("MsgToWSPaymentACK err:", err)
		return
	}

	//get POS conn
	if req.Payment == nil {
		fmt.Println("Payment is empty.")
		return
	}
	if req.Payment.DeviceId == "" || len(req.Payment.DeviceId) == 0 {
		fmt.Println("DeviceId is empty.")
		return
	}
	if c, ok := wsItemMap.GetConn(req.Payment.DeviceId); ok {
		wsData := &model.WsData{
			M: model.WS_POS_NOTIFY,
			C: model.WS_CODE_SERVER,
			D: req,
		}
		SendMsgToConn(c, wsData)
		return
	}
	fmt.Println("pos is not exist | id: [", req.Payment.DeviceId, "]")
}

//WS Reply
func SendWSResponseForErr(c net.Conn, wsItemMap *WsItemMap, msg string) {
	if _, ok := wsItemMap.Get(c); !ok {
		fmt.Println("conn is not exist")
		return
	} else {
		wsData := &model.WsData{
			M: model.WS_RESPONSE_ERROR,
			C: model.WS_CODE_SEND_ERROR,
			D: msg,
		}
		SendMsgToConn(c, wsData)
		return
	}
}

func SendMsgToConn(c net.Conn, wsData *model.WsData)  {
	wsDataStr,err := wsData.ToString()
	if err != nil {
		fmt.Println(err.Error())
	}
	err = wsutil.WriteServerMessage(c, ws.OpText, []byte(wsDataStr))
	if err != nil {
		// handle error
		fmt.Println("WriteServerMessage Err:" + err.Error())
	}
	return
}

