package ggin

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lsq51201314/go-utils/glog"
)

// 跨域设置
func cors() gin.HandlerFunc {
	return func(context *gin.Context) {
		method := context.Request.Method
		context.Header("Access-Control-Allow-Origin", "*")
		context.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
		context.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		context.Header("Access-Control-Expose-Headers", "Content-Length, Content-Type, Access-Control-Allow-Origin, Access-Control-Allow-Headers")
		context.Header("Access-Control-Allow-Credentials", "true")
		if method == "OPTIONS" {
			context.AbortWithStatus(http.StatusNoContent)
			return
		}
		context.Next()
	}
}

// 处理错误
func recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				c.AbortWithStatus(http.StatusInternalServerError)
				ServerError(c, fmt.Errorf("%v", err))
				return
			}
		}()
		c.Next()
	}
}

// 控制台输出
func console() func(c *gin.Context) {
	return func(c *gin.Context) {
		glog.Info(fmt.Sprintf("method:%s\t status:%d\t path:%s",
			c.Request.Method,
			c.Writer.Status(),
			c.Request.URL.Path,
		))
		c.Next()
	}
}
