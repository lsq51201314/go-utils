package ggin

import (
	"github.com/gin-gonic/gin"
	"github.com/lsq51201314/go-utils/gjwt"
	"github.com/spf13/cast"
)

// 校验凭证 Authorization
func Validate(token *gjwt.Token) func(c *gin.Context) {
	return func(c *gin.Context) {
		//获取token
		h := c.Request.Header.Get("Authorization")
		if h == "" {
			MustLogin(c)
			c.Abort()
			return
		}
		//解析token
		userId, err := token.Validate(h)
		if err != nil {
			MustLogin(c)
			c.Abort()
			return
		}
		//存到请求
		c.Set("uid", userId)
		c.Next()
	}
}

// 获取用户 uid
func GetUserID(c *gin.Context, userId *string) bool {
	val, ok := c.Get("uid")
	if !ok {
		return false
	}
	*userId = cast.ToString(val)
	return true
}
