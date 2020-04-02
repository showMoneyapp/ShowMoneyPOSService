package posWS

import (
	"fmt"
	"github.com/ShowPay/ShowMoneyPosService/model"
	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"net/http"
)

var (
	wsItemMap *WsItemMap //pos端map
)

func init() {
	wsItemMap = NewWsItemMap()
}

func WsHandler(w http.ResponseWriter, r *http.Request) {
	conn, _, _, err := ws.UpgradeHTTP(r, w)
	if err != nil {
		fmt.Println("showClient连接失败...", err)
	}

	//localAdd := conn.LocalAddr()
	remoteAdd := conn.RemoteAddr()
	//fmt.Println("localAdd：", localAdd.String())
	fmt.Println("remoteAdd：", remoteAdd.String())

	//判断是否已连接该client，有则不用重复加入
	if _, ok := wsItemMap.Get(conn); !ok {
		wsClient := NewWsClient()
		wsItemMap.Set(conn, wsClient)
	}

	go func() {
		defer func() {
			//先清除内存再close
			if _, ok := wsItemMap.Get(conn); ok {
				wsItemMap.Deleted(conn)
				fmt.Println("ShowServer-" + remoteAdd.String() + "-主动断开")
			}
			if err = conn.Close(); err != nil{
				fmt.Println("ShowClient-" + remoteAdd.String() + "-close err:"  + err.Error())
			}
		}()
		for {
			msg, _, err := wsutil.ReadClientData(conn)
			if err != nil {
				//发生err后马上断开
				if _, ok := wsItemMap.Get(conn); ok {
					wsItemMap.Deleted(conn)
					fmt.Println("ShowClient-" + remoteAdd.String() + "-断开：" + err.Error())
					break
				}
			}else {
				fmt.Println("ShowClient-" + remoteAdd.String() + "ReadData :" + string(msg))
				wsData := model.WsDataFromStringMsg(string(msg))
				if wsData == nil {
					fmt.Println("wsData is nil")
					SendWSResponseForErr(conn, wsItemMap, "wsData is nil")//返回业务错误信息
					continue
				}

				if wsData.M == "" {
					fmt.Println("wsData.M is nil")
					SendWSResponseForErr(conn, wsItemMap, "wsData.M is nil")//返回业务错误信息
					continue
				}
				if wsData.D == "" {
					fmt.Println("wsData.D is nil")
					SendWSResponseForErr(conn, wsItemMap, "wsData.D is nil")//返回业务错误信息
					continue
				}

				//选择方法
				switch wsData.M {
				case model.HEART_BEAT: //返回心跳
					SendHeartBeat(conn)
					break
				case model.WS_PAYMENTREQUEST:
					err := GetPaymentRequest(conn, wsItemMap, wsData.D.(string))
					if err != nil {
						SendWSResponseForErr(conn, wsItemMap, err.Error())//返回业务错误信息
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
					SendWSResponseForErr(conn, wsItemMap, "WsData.M can not find")//返回业务错误信息
				}
			}
		}
		Back :{
			fmt.Println("跳出For")
		}
	}()
}

