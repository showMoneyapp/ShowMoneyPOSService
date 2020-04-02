package main

import (
	"fmt"
	"github.com/ShowPay/ShowMoneyPosService/model"
	"github.com/ShowPay/ShowMoneyPosService/posAPI"
	"github.com/ShowPay/ShowMoneyPosService/posWS"
	"github.com/ShowPay/ShowMoneyPosService/util"
	"github.com/gookit/ini/v2"
	"net/http"
	"path/filepath"
)


func init() {
	absFile := filepath.Join("conf", "config.ini")
	err := ini.LoadExists(absFile, "not-exist.ini")
	if err != nil {
		panic(err)
	}
	model.POS_WS_Port = ini.String("pos_ws_port")
}

func startWS() {
	http.HandleFunc("/ws", posWS.WsHandler)
	fmt.Println("Start showPOS-WS-service...")
	err := http.ListenAndServe(util.AddStr(":", model.POS_WS_Port), nil)
	if err != nil {
		panic(err)
	}

}

func main() {
	go posAPI.StartAPI()
	startWS()
}
