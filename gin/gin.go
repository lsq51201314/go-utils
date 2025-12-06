package gin

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	ggin "github.com/gin-gonic/gin"
	"github.com/lsq51201314/go-utils/log"
)

type RouterSetupFunc func(*ggin.RouterGroup)

type Service struct {
	port int
	r    *ggin.Engine
	v    *ggin.RouterGroup
}

// 新建
func New(group string, port int, allowCors bool, setupRoutes RouterSetupFunc, mode ...string) *Service {
	//模式
	if len(mode) > 0 && mode[0] == ggin.ReleaseMode {
		ggin.SetMode(ggin.ReleaseMode)
	}
	obj := Service{
		port: port,
		r:    ggin.New(),
	}
	//跨域
	if allowCors {
		obj.r.Use(cors())
	}
	obj.r.Use(recovery())                             //错误处理
	if len(mode) > 0 && mode[0] != ggin.ReleaseMode { //测试用
		obj.r.Use(console())
	}
	obj.v = obj.r.Group(group)
	//路由设置
	setupRoutes(obj.v)
	obj.r.NoRoute(func(c *ggin.Context) {
		ServerError(c, errors.New("无效接口"))
	})
	return &obj
}

// 运行
func (t *Service) Run() {
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", t.port),
		Handler: t.r,
	}
	go func() {
		log.Info(fmt.Sprintf("服务开始运行(%d)", t.port))
		if err := srv.ListenAndServe(); err != nil {
			log.Error(err.Error())
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Info("服务正在关闭")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Error(err.Error())
	}
	log.Info("服务关闭成功")
}
