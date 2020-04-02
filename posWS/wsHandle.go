package posWS

import (
	"fmt"
	"github.com/ShowPay/ShowMoneyPosService/model"
	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"net/http"
)

var (
	wsItemMap *WsItemMap //POS-Device Map
)

func init() {
	wsItemMap = NewWsItemMap()
}

func WsHandler(w http.ResponseWriter, r *http.Request) {
	conn, _, _, err := ws.UpgradeHTTP(r, w)
	if err != nil {
		fmt.Println("showClient connect fail...", err)
	}

	//localAdd := conn.LocalAddr()
	remoteAdd := conn.RemoteAddr()
	//fmt.Println("localAdd：", localAdd.String())
	fmt.Println("remoteAdd：", remoteAdd.String())

	//check client if is existed
	if _, ok := wsItemMap.Get(conn); !ok {
		wsClient := NewWsClient()
		wsItemMap.Set(conn, wsClient)
	}

	go func() {
		defer func() {
			//clean and close client
			if _, ok := wsItemMap.Get(conn); ok {
				wsItemMap.Deleted(conn)
				fmt.Println("ShowServer-" + remoteAdd.String() + "-proactive close")
			}
			if err = conn.Close(); err != nil{
				fmt.Println("ShowClient-" + remoteAdd.String() + "-close err:"  + err.Error())
			}
		}()
		for {
			msg, _, err := wsutil.ReadClientData(conn)
			if err != nil {
				//Disconnect immediately if has err
				if _, ok := wsItemMap.Get(conn); ok {
					wsItemMap.Deleted(conn)
					fmt.Println("ShowClient-" + remoteAdd.String() + "-close：" + err.Error())
					break
				}
			}else {
				fmt.Println("ShowClient-" + remoteAdd.String() + "ReadData :" + string(msg))
				wsData := model.WsDataFromStringMsg(string(msg))
				if wsData == nil {
					fmt.Println("wsData is nil")
					SendWSResponseForErr(conn, wsItemMap, "wsData is nil")//Return error message
					continue
				}

				if wsData.M == "" {
					fmt.Println("wsData.M is nil")
					SendWSResponseForErr(conn, wsItemMap, "wsData.M is nil")//Return error message
					continue
				}
				if wsData.D == "" {
					fmt.Println("wsData.D is nil")
					SendWSResponseForErr(conn, wsItemMap, "wsData.D is nil")//Return error message
					continue
				}

				//select method
				switch wsData.M {
				case model.HEART_BEAT:
					SendHeartBeat(conn)
					break
				case model.WS_PAYMENTREQUEST:
					err := GetPaymentRequest(conn, wsItemMap, wsData.D.(string))
					if err != nil {
						SendWSResponseForErr(conn, wsItemMap, err.Error())//Return error message
						continue
					}
					break
				case model.WS_POS_NOTIFY:
					NotifyPaymentACKToPos(conn, wsItemMap, wsData.D.(string))
					break
				case model.WS_DISCONNECT:
					goto Back
					break
				default:
					fmt.Println("WsData.M can not find")
					SendWSResponseForErr(conn, wsItemMap, "WsData.M can not find")//Return error message
				}
			}
		}
		Back :{
			fmt.Println("Jump out the For")
		}
	}()
}

