package ghttp

import (
	"net/http"

	"github.com/gorilla/websocket"
)

type OnOpen func(conn *websocket.Conn)
type OnClose func(conn *websocket.Conn, err error)
type OnMessage func(conn *websocket.Conn, msg string)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type WSServer struct {
	conn      *websocket.Conn
	onopen    OnOpen
	onclose   OnClose
	onmessage OnMessage
}

// 客户进入
func (t *WSServer) SetOnOpen(cfunc OnOpen) {
	t.onopen = cfunc
}

// 客户离开
func (t *WSServer) SetOnClose(cfunc OnClose) {
	t.onclose = cfunc
}

// 消息到达
func (t *WSServer) SetOnMessage(cfunc OnMessage) {
	t.onmessage = cfunc
}

// 发送消息
func (t *WSServer) Send(data []byte) error {
	if t.conn != nil {
		return t.conn.WriteMessage(websocket.TextMessage, data)
	}
	return nil
}

// 关闭连接
func (t *WSServer) Close() {
	if t.conn != nil {
		t.conn.Close()
		t.conn = nil
	}
}

// 阻塞
func (t *WSServer) Run(w http.ResponseWriter, r *http.Request) {
	t.Close()
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		//连接断开
		if t.onclose != nil {
			t.onclose(t.conn, err)
		}
		return
	}
	t.conn = conn
	//连接成功
	if t.onopen != nil {
		t.onopen(t.conn)
	}
	//监听消息
	for {
		mt, msg, err := t.conn.ReadMessage()
		if err != nil {
			//连接断开
			if t.onclose != nil {
				t.onclose(t.conn, err)
			}
			break
		}
		if mt == websocket.TextMessage && len(msg) > 0 {
			if t.onmessage != nil { //消息到达
				t.onmessage(t.conn, string(msg))
			}
		}
	}
}
