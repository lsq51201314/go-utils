package utils

import (
	"github.com/mojocn/base64Captcha"
)

var auth = base64Captcha.DefaultMemStore

type VerifyCode struct{}

func (v VerifyCode) GetCode() (id, code, image string, err error) {
	d := base64Captcha.DefaultDriverDigit
	b := base64Captcha.NewCaptcha(d, auth)
	id, image, code, err = b.Generate()
	return
}

func (v VerifyCode) Verify(id, code string) (ok bool) {
	ok = auth.Verify(id, code, true)
	return
}
