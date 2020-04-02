package posWS

import (
	"net"
	"sync"
)

type WsItemMap struct {
	WsConns map[string]net.Conn
	WsItems map[net.Conn]*WsClient
	Lock    *sync.RWMutex
}

func NewWsItemMap() *WsItemMap {
	wsItemMap := &WsItemMap{
		WsConns:make(map[string]net.Conn),
		WsItems: make(map[net.Conn]*WsClient),
		Lock: new(sync.RWMutex),
	}

	return wsItemMap
}

func (i WsItemMap) Get(c net.Conn) (*WsClient, bool)  {
	i.Lock.RLock()
	defer i.Lock.RUnlock()
	if v, ok := i.WsItems[c]; ok {
		return v, true
	}
	return nil, false
}

func (i WsItemMap) Set(c net.Conn, w *WsClient)  {
	i.Lock.Lock()
	defer i.Lock.Unlock()
	i.WsItems[c] = w
}

func (i WsItemMap) GetConn(deviceId string) (net.Conn, bool)  {
	i.Lock.RLock()
	defer i.Lock.RUnlock()
	if c, ok := i.WsConns[deviceId]; ok {
		return c, true
	}
	return nil, false
}

func (i WsItemMap) SetConn(deviceId string, c net.Conn)  {
	i.Lock.Lock()
	defer i.Lock.Unlock()
	i.WsConns[deviceId] = c
}

//func (i WsItemMap) Init()  {
//	i.Lock.Lock()
//	defer i.Lock.Unlock()
//	for key,_ := range i.WsItems{
//		delete(i.WsItems, key)
//	}
//}

func (i WsItemMap) Deleted(c net.Conn)  {
	i.Lock.Lock()
	defer i.Lock.Unlock()
	//delete(i.WsItems, c)
	for key, v := range i.WsConns {
		if c == v {
			delete(i.WsConns, key)
			break
		}
	}
}
