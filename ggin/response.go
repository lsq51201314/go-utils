package ggin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 必须登录
func MustLogin(c *gin.Context) {
	c.AbortWithStatus(http.StatusUnauthorized)
}

// 服务器错误
func ServerError(c *gin.Context, err error) {
	str := "服务繁忙"
	c.JSON(http.StatusInternalServerError, gin.H{
		"code": http.StatusInternalServerError,
		"msg":  str,
	})
}

// 操作成功
func SendSuccess(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "success",
	})
}

// 发送文本
func SendText(c *gin.Context, code int, msg string) {
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg,
	})
}

// 发送数据
func SendData(c *gin.Context, data any) {
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "success",
		"data": data,
	})
}

// 发送表格
func SendRows(c *gin.Context, total int64, data any) {
	type Data struct {
		Total int64 `json:"total,string"`
		Rows  any   `json:"rows"`
	}
	SendData(c, &Data{
		Total: total,
		Rows:  data,
	})
}

// 发送对象
func SendObject(c *gin.Context, data any) {
	c.JSON(http.StatusOK, data)
}
