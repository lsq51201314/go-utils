package utils

import (
	"errors"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/zh"
	uTranslations "github.com/go-playground/universal-translator"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
)

// 参数绑定
type GinBind struct {
	trans ut.Translator
}

// 新建实例
func NewGinBind() (g GinBind, err error) {
	if t, ok := binding.Validator.Engine().(*validator.Validate); !ok {
		err = errors.New("translator error")
		return
	} else {
		t.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("zh"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})
		zhT := zh.New()
		uni := uTranslations.New(zhT, zhT)
		if g.trans, ok = uni.GetTranslator("zh"); ok {
			err = zhTranslations.RegisterDefaultTranslations(t, g.trans)
		}
		return
	}
}

func (g *GinBind) translate(errs validator.ValidationErrors) (res string) {
	fields := errs.Translate(g.trans)
	for _, err := range fields {
		res += err + "。 "
	}
	return
}

// 获取参数
func (g *GinBind) Get(c *gin.Context, params interface{}, data interface{}) (err error) {
	//绑定参数
	if params != nil {
		err = c.ShouldBindQuery(params)
		if err != nil {
			if errs, ok := err.(validator.ValidationErrors); !ok {
				return
			} else {
				return errors.New(g.translate(errs))
			}
		}
	}
	//绑定数据
	if data != nil {
		err = c.ShouldBindJSON(data)
		if err != nil {
			if errs, ok := err.(validator.ValidationErrors); !ok {
				return
			} else {
				return errors.New(g.translate(errs))
			}
		}
	}
	return
}
