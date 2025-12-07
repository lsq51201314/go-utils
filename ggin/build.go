package ggin

import (
	"errors"
	"io"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// 解析参数
func Unmarshal(c *gin.Context, params any, data any) bool {
	//绑定参数
	if params != nil {
		if err := c.ShouldBindQuery(params); err != nil {
			if errs, ok := err.(validator.ValidationErrors); !ok {
				ServerError(c, err)
			} else {
				ServerError(c, errors.Join(errs))
			}
			return false
		}
	}
	//绑定数据
	if data != nil {
		if err := c.ShouldBindJSON(data); err != nil {
			if err == io.EOF {
				ServerError(c, errors.New("数据为空"))
			} else if errs, ok := err.(validator.ValidationErrors); !ok {
				ServerError(c, err)
			} else {
				ServerError(c, errors.Join(errs))
			}
			return false
		}
	}
	return true
}
