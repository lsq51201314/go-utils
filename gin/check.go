package gin

import (
	ggin "github.com/gin-gonic/gin"
	"github.com/lsq51201314/go-utils/jwt"
	"github.com/spf13/cast"
)

// 校验凭证
func CheckToken(token *jwt.Token) func(c *ggin.Context) {
	return func(c *ggin.Context) {
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

// 获取用户
func GetUserID(c *ggin.Context, userId *int64) bool {
	val, ok := c.Get("uid")
	if !ok {
		return false
	}
	*userId = cast.ToInt64(val)
	return true
}
