package ghttp

import (
	"github.com/gorilla/websocket"
)

type WSClient struct {
	conn      *websocket.Conn
	onopen    OnOpen
	onclose   OnClose
	onmessage OnMessage
}

// 连接成功
func (t *WSClient) SetOnOpen(cfunc OnOpen) {
	t.onopen = cfunc
}

// 连接断开
func (t *WSClient) SetOnClose(cfunc OnClose) {
	t.onclose = cfunc
}

// 消息到达
func (t *WSClient) SetOnMessage(cfunc OnMessage) {
	t.onmessage = cfunc
}

// 发送消息
func (t *WSClient) Send(data []byte) error {
	if t.conn != nil {
		return t.conn.WriteMessage(websocket.TextMessage, data)
	}
	return nil
}

// 关闭连接
func (t *WSClient) Close() {
	if t.conn != nil {
		t.conn.Close()
		t.conn = nil
	}
}

// 阻塞
func (t *WSClient) Run(url string, headers map[string][]string) {
	t.Close()
	conn, _, err := websocket.DefaultDialer.Dial(url, headers)
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
