package posAPI

import (
	"fmt"
	"github.com/ShowPay/ShowMoneyPosService/model"
	"github.com/ShowPay/ShowMoneyPosService/util"
	"github.com/gookit/ini/v2"
	"github.com/gorilla/mux"
	"net/http"
	"path/filepath"
)

func init() {
	absFile := filepath.Join("conf", "config.ini")
	err := ini.LoadExists(absFile, "not-exist.ini")
	if err != nil {
		panic(err)
	}
	model.POS_API_Port = ini.String("POS_API_Port")
}

func StartAPI() {
	r := mux.NewRouter()
	r.HandleFunc("/api/v1/pos/broadcastPayment", BroadcastPayment).Methods("POST")
	fmt.Println("Start showPOS-API-Service...")
	fmt.Println(util.AddStr("Listen: http://:", model.POS_API_Port))
	err := http.ListenAndServe(util.AddStr(":", model.POS_API_Port), r)
	if err != nil {
		panic(err)
	}
}