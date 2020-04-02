package posWS

import (
	"fmt"
	"github.com/ShowPay/ShowMoneyPosService/model"
	"github.com/ShowPay/ShowMoneyPosService/service"
	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"net"
)

//推送消息心跳
func SendHeartBeat(c net.Conn) {
	wsData := &model.WsData{
		M: model.HEART_BEAT,
		C: model.WS_CODE_HEART_BEAT,
		D: "Heart beat back",
	}
	SendMsgToConn(c, wsData)
}

//接收响应保存paymentRequest
func GetPaymentRequest(c net.Conn, wsItemMap *WsItemMap, msg string) error {
	req := &model.WsPaymentRequestReq{}
	resp := &model.WsResponse{}
	err := new(impl.Ws_showPos).MsgToWSPaymentRequest(msg, req)
	if err != nil {
		fmt.Println("MsgToWSPaymentRequest err:", err)
		return err
	}
	//缓存conn
	wsItemMap.SetConn(req.DeviceId, c)

	err = new(impl.Ws_showPos).GetAndCreatePaymentReq(req, resp)
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

//API转发PaymentACK
func NotifyPaymentACKToPos(c net.Conn, wsItemMap *WsItemMap, msg string) {
	req := &model.WSPaymentACKReq{}
	err := new(impl.Ws_showPos).MsgToWSPaymentACK(msg, req)
	if err != nil {
		fmt.Println("MsgToWSPaymentACK err:", err)
		return
	}

	//获取pos的conn
	if req.Payment == nil {
		fmt.Println("Payment为空")
		return
	}
	if req.Payment.DeviceId == "" || len(req.Payment.DeviceId) == 0 {
		fmt.Println("DeviceId为空")
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

//WS回复
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