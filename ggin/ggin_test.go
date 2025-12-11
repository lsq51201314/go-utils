package ggin

import (
	"context"
	"fmt"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/lsq51201314/go-utils/gexit"
	"github.com/lsq51201314/go-utils/ghttp"
	"github.com/lsq51201314/go-utils/gjwt"
)

func TestGgin(t *testing.T) {
	g := New("/api", 22345, true, func(r *gin.RouterGroup) {
		r.GET("/a", nil)
		r.POST("/b", nil)
		r.PUT("/c", nil)
		r.DELETE("/d", nil)
		r.GET("/ws", wstest)
		j := gjwt.New("123456")
		r.Use(Validate(j))
		{
			r.GET("/e", nil)
			r.POST("/f", nil)
			r.PUT("/g", nil)
			r.DELETE("/h", nil)
		}
	})
	g.Run()
	gexit.WaitExit(func(ctx context.Context) {
		g.Stop(ctx)
	})
}

func wstest(c *gin.Context) {
	var wss ghttp.WSServer
	wss.SetOnOpen(onopen)
	wss.SetOnClose(onclose)
	wss.SetOnMessage(onmessage)
	wss.Run(c.Writer, c.Request)
}

func onopen(conn *websocket.Conn) {
	fmt.Println("客户进入：" + conn.RemoteAddr().String())
}

func onclose(conn *websocket.Conn, err error) {
	fmt.Println("客户离开："+conn.RemoteAddr().String(), err.Error())
}

func onmessage(conn *websocket.Conn, msg string) {
	fmt.Println("消息到达(" + conn.RemoteAddr().String() + ")：" + msg)
}
