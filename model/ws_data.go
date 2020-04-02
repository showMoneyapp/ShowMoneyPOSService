package model

import (
	"github.com/godaddy-x/jorm/util"
	"github.com/tidwall/gjson"
)

type WsData struct {
	M string      `json:"M"` //method
	C int64       `json:"C"` //code
	D interface{} `json:"D"` //data
}

func WsDataFromStringMsg(msg string) *WsData  {
	if !gjson.Valid(msg) {
		return nil
	}
	ws := &WsData{}
	ws.M = gjson.Get(msg, "M").String()
	ws.C = gjson.Get(msg, "C").Int()
	ws.D = gjson.Get(msg, "D").String()
	return ws
}

func (w *WsData) ToString() (string, error)  {
	return util.ObjectToJson(w)
}