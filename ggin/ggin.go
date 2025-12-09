package ggin

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lsq51201314/go-utils/glog"
)

type RouterSetupFunc func(*gin.RouterGroup)

type Ggin struct {
	port int
	r    *gin.Engine
	v    *gin.RouterGroup
	srv  *http.Server
}

// 新建
func New(group string, port int, allowCors bool, setupRoutes RouterSetupFunc, mode ...string) *Ggin {
	//模式
	if len(mode) > 0 && mode[0] == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}
	obj := Ggin{
		port: port,
		r:    gin.New(),
	}
	//跨域
	if allowCors {
		obj.r.Use(cors())
	}
	obj.r.Use(recovery())                            //错误处理
	if len(mode) > 0 && mode[0] != gin.ReleaseMode { //测试用
		obj.r.Use(console())
	}
	obj.v = obj.r.Group(group)
	//路由设置
	setupRoutes(obj.v)
	obj.r.NoRoute(func(c *gin.Context) {
		ServerError(c, errors.New("无效接口"))
	})
	return &obj
}

// 运行
func (t *Ggin) Run() {
	t.srv = &http.Server{
		Addr:    fmt.Sprintf(":%d", t.port),
		Handler: t.r,
	}
	go func() {
		glog.Info(fmt.Sprintf("服务开始运行(%d)", t.port))
		if err := t.srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			glog.Error(err.Error())
		}
	}()
}

// 停止
func (t *Ggin) Stop(ctx context.Context) {
	glog.Info("服务正在关闭")
	if err := t.srv.Shutdown(ctx); err != nil {
		glog.Error(err.Error())
	}
	glog.Info("服务关闭成功")
}

