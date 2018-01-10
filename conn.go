package naivechain_demo

import (
	"golang.org/x/net/websocket"
	"time"
)

// 用于操作网络连接的
type Conn struct {
	*websocket.Conn
	t int64
}

func newConn(ws *websocket.Conn) *Conn {
	return &Conn{
		Conn: ws,
		t:    time.Now().Unix(),
	}
}
