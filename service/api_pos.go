package service

import (
	"errors"
	"fmt"
	"github.com/ShowPay/ShowMoneyPosService/model"
	"github.com/ShowPay/ShowMoneyPosService/node"
	"github.com/gorilla/websocket"
	"log"
	"net/url"
	"time"
)

type Api_showPOS struct {}

func (sp *Api_showPOS) BroadcastPayment(req *model.ApiPaymentReq, resp *model.ApiPaymentResultResp) error {
	if req.DeviceId == "" || len(req.DeviceId) == 0 {
		return errors.New("DeviceId is empty.")
	}
	if req.WalletId == "" || len(req.WalletId) == 0 {
		return errors.New("WalletId is empty.")
	}
	if req.TransactionHex == "" || len(req.TransactionHex) == 0 {
		return errors.New("TransactionHex is empty.")
	}
	if req.RefundTo == "" || len(req.RefundTo) == 0 {
		return errors.New("RefundTo is empty.")
	}

	//broadcast
	txId, err := node.NewClientController().BroadcastTx(req.TransactionHex)
	if err != nil {
		resp.Error = "err"
		resp.Code = 400
		resp.Transaction = ""
		return nil
	}
	resp.Error = "null"
	resp.Code = 200
	resp.Transaction = txId



	u := url.URL{Scheme:"ws", Host:model.Domain_POS_WS, Path:"/ws"}
	fmt.Println("connecting to :", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		return errors.New("dial:" + err.Error())
	}
	//notify PaymentACK to POS
	wsData := &model.WsData{
		M: model.WS_POS_NOTIFY,
		C: 0,
		D: resp,
	}
	wsDataStr, err := wsData.ToString()
	if err != nil {
		return errors.New("wsData conversion failed:" + err.Error())
	}
	err = c.WriteMessage(websocket.TextMessage, []byte(wsDataStr))
	if err != nil {
		log.Println("write | err :", err)
		return errors.New("write | err :" + err.Error())
	}
	//disconnect
	wsDataDis := &model.WsData{
		M: model.WS_DISCONNECT,
		C: 0,
		D: "disconnect",
	}
	wsDataDisStr, err := wsDataDis.ToString()
	if err != nil {
		return errors.New("wsDataDis conversion failed:" + err.Error())
	}
	err = c.WriteMessage(websocket.TextMessage, []byte(wsDataDisStr))
	if err != nil {
		log.Println("write | err :", err)
		return errors.New("write | err :" + err.Error())
	}

	time.Sleep(100*time.Millisecond)
	defer c.Close()

	return nil
}