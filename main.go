package main

import (
	"fmt"
	"github.com/ShowPay/ShowMoneyPosService/posAPI"
	"github.com/ShowPay/ShowMoneyPosService/posWS"
	"github.com/ShowPay/ShowMoneyPosService/util"
	"github.com/gookit/ini/v2"
	"net/http"
	"path/filepath"
)

var pos_ws_port = "1234"

func init() {
	absFile := filepath.Join("conf", "config.ini")
	err := ini.LoadExists(absFile, "not-exist.ini")
	if err != nil {
		panic(err)
	}
	pos_ws_port = ini.String("pos_ws_port")
}

func startWS() {
	http.HandleFunc("/ws", posWS.WsHandler)
	fmt.Println("Start showPOS-WS-service...")
	err := http.ListenAndServe(util.AddStr(":", pos_ws_port), nil)
	if err != nil {
		panic(err)
	}

}

func main() {
	go posAPI.StartAPI()
	startWS()
}
