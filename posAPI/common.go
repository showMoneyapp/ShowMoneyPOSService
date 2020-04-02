package posAPI

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ShowPay/ShowMoneyPosService/util"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	success_code = 200
	fail_code = 400
)

type securityRequest struct {
	Data  string `json:"data"`
}

type securityResponse struct {
	Result  interface{} `json:"result"`
	Code    int         `json:"code"`
	Message string      `json:"msg"`
	Time    int64       `json:"time"`
	Error   string      `json:"error"`
}

func pares(r *http.Request, request interface{}) error {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return errors.New("get params from request fail:" + err.Error())
	}
	r.Body.Close()
	if body == nil || len(body) == 0 {
		return errors.New("params is null")
	}
	securityReq := &securityRequest{}
	if err := util.JsonToObject(string(body), securityReq); err != nil {
		return errors.New("parsing the params fail：" + err.Error())
	}
	data := securityReq.Data

	if data == "" || len(data) == 0 {
		return errors.New("Data is null")
	}

	if err := util.JsonToObject(data, request); err != nil {
		return errors.New("data change fail" + err.Error())
	}

	fmt.Println("POST Data：", request)//Print to check
	return nil
}

func response(w http.ResponseWriter, response *securityResponse)  {
	resp, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.Write(resp)
}

func ResponseSuccess(w http.ResponseWriter, result interface{})  {
	response(w, &securityResponse{
		Result:  result,
		Code:    success_code,
		Message: "",
		Time:    time.Now().Unix(),
		Error:   "",
	})
}

func ResponseError(w http.ResponseWriter, errorMsg string)  {
	response(w, &securityResponse{
		Result:  nil,
		Code:    fail_code,
		Message: "",
		Time:    time.Now().Unix(),
		Error:   errorMsg,
	})
}
