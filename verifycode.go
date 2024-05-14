package utils

import (
	"github.com/mojocn/base64Captcha"
)

var auth = base64Captcha.DefaultMemStore

// 数字验证
type VerifyCode struct{}

// 获取验证码
func (v VerifyCode) GetCode() (id, code, image string, err error) {
	d := base64Captcha.DefaultDriverDigit
	b := base64Captcha.NewCaptcha(d, auth)
	id, image, code, err = b.Generate()
	return
}

// 校验验证码
func (v VerifyCode) Verify(id, code string) (ok bool) {
	ok = auth.Verify(id, code, true)
	return
}
