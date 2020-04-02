package posAPI

import (
	"fmt"
	"github.com/ShowPay/ShowMoneyPosService/util"
	"github.com/gookit/ini/v2"
	"github.com/gorilla/mux"
	"net/http"
	"path/filepath"
)

var pos_api_port = "5467"

func init() {
	absFile := filepath.Join("conf", "config.ini")
	err := ini.LoadExists(absFile, "not-exist.ini")
	if err != nil {
		panic(err)
	}
	pos_api_port = ini.String("pos_api_port")
}

func StartAPI() {
	r := mux.NewRouter()
	r.HandleFunc("/api/v1/pos/broadcastPayment", BroadcastPayment).Methods("POST")
	fmt.Println("Start showPOS-API-Service...")
	err := http.ListenAndServe(util.AddStr(":", pos_api_port), r)
	if err != nil {
		panic(err)
	}
}