package node

import (
	"fmt"
	"github.com/gookit/ini/v2"
	"path/filepath"
)

type ClientController struct {
	Client *Client
}

var (
	RPC_url      string
	RPC_username string
	RPC_password string
)

func NewClientController() *ClientController {
	// config, err := ini.LoadFiles("testdata/tesdt.ini")
	// LoadExists will ignore not exists file
	absFile := filepath.Join("conf", "config.ini")
	err := ini.LoadExists(absFile, "not-exist.ini")
	if err != nil {
		panic(err)
	}
	RPC_url = ini.String("RPC_url")
	RPC_username = ini.String("RPC_username")
	RPC_password = ini.String("RPC_password")

	fmt.Println( "************************Print Config***************************")
	fmt.Println( "*******RPC_url : [ ", RPC_url, " ]")
	fmt.Println( "*******RPC_username : [ ", "*********", " ]")
	fmt.Println( "*******RPC_password : [ ", "*********", " ]")
	fmt.Println( "************************Print Config***************************")

	accessToken := BasicAuth(RPC_username, RPC_password)
	clientController := &ClientController{
		Client:NewClientNode(RPC_url, accessToken, true),
	}
	fmt.Println( "****** Build new Client completed ******")

	return clientController
}

func (c *ClientController) BroadcastTx(txHexStr string) (string, error) {
	request := []interface{}{
		txHexStr,
		false,
	}

	result, err := c.Client.Call("sendrawtransaction", request)
	if err != nil {
		return "", err
	}
	return result.String(), nil
}
