package impl

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

func (sp *Api_showPOS) BroadcastPayment(req *model.ApiPaymentReq, resp *model.ApiPaymentACKResp) error {
	if req.DeviceId == "" || len(req.DeviceId) == 0 {
		return errors.New("DeviceId为空")
	}
	if req.WalletId == "" || len(req.WalletId) == 0 {
		return errors.New("WalletId为空")
	}
	if req.TransactionHex == "" || len(req.TransactionHex) == 0 {
		return errors.New("TransactionHex为空")
	}
	if req.RefundTo == "" || len(req.RefundTo) == 0 {
		return errors.New("RefundTo为空")
	}

	//广播
	txId, err := node.NewClientController().BroadcastTx(req.TransactionHex)
	if err != nil {
		return errors.New("Broadcast Tx Fail:" + err.Error())
	}
	resp.Message = "Broadcast Success"
	resp.Code = 200
	resp.Memo = req.Memo
	resp.TxId = txId
	resp.Payment = req



	u := url.URL{Scheme:"ws", Host:model.Domain_pos_ws, Path:"/ws"}
	fmt.Println("connecting to :", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		return errors.New("dial:" + err.Error())
	}
	//转发PaymentACK给pos
	wsData := &model.WsData{
		M: model.WS_POS_NOTIFY,
		C: 0,
		D: resp,
	}
	wsDataStr, err := wsData.ToString()
	if err != nil {
		return errors.New("wsData转str失败:" + err.Error())
	}
	err = c.WriteMessage(websocket.TextMessage, []byte(wsDataStr))
	if err != nil {
		log.Println("write | err :", err)
		return errors.New("write | err :" + err.Error())
	}
	//主动断开
	wsDataDis := &model.WsData{
		M: model.WS_DISCONNECT,
		C: 0,
		D: "disconnect",
	}
	wsDataDisStr, err := wsDataDis.ToString()
	if err != nil {
		return errors.New("wsDataDis转str失败:" + err.Error())
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