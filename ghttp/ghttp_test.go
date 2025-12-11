package ghttp

import (
	"fmt"
	"testing"

	"github.com/gorilla/websocket"
)

func TestGhttp(t *testing.T) {
	// if code, buf, err := Get("http://127.0.0.1:22345/api/verify", nil, nil); err != nil {
	// 	log.Fatalln(err)
	// } else {
	// 	fmt.Println(code, string(buf))
	// }

	// if code, buf, err := Post("http://127.0.0.1:22345/api/login", nil, nil, []byte("hello")); err != nil {
	// 	log.Fatalln(err)
	// } else {
	// 	fmt.Println(code, string(buf))
	// }

	var wsc WSClient
	wsc.SetOnOpen(onopen)
	wsc.SetOnClose(onclose)
	wsc.SetOnMessage(onmessage)
	wsc.Run("ws://127.0.0.1:22345/api/ws", nil)
}

func onopen(conn *websocket.Conn) {
	fmt.Println("连接成功：" + conn.RemoteAddr().String())
}

func onclose(conn *websocket.Conn, err error) {
	fmt.Println("连接断开："+conn.RemoteAddr().String(), err.Error())
}

func onmessage(conn *websocket.Conn, msg string) {
	fmt.Println("消息到达(" + conn.RemoteAddr().String() + ")：" + msg)
}
