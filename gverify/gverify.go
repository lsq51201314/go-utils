package gverify

import (
	"github.com/mojocn/base64Captcha"
)

type Gverify struct {
	auth base64Captcha.Store
	d    *base64Captcha.DriverDigit
	b    *base64Captcha.Captcha
}

type CodeData struct {
	ID    string `json:"id"`
	Image string `json:"image"`
}

// 新建
func New() *Gverify {
	obj := Gverify{
		auth: base64Captcha.DefaultMemStore,
		d:    base64Captcha.DefaultDriverDigit,
	}
	obj.b = base64Captcha.NewCaptcha(obj.d, obj.auth)
	return &obj
}

// 获取
func (t *Gverify) Get() (CodeData, error) {
	id, b64, _, err := t.b.Generate()
	if err != nil {
		return CodeData{}, err
	}
	return CodeData{
		ID:    id,
		Image: b64,
	}, nil
}

// 验证
func (t *Gverify) Verify(id, code string) bool {
	return t.auth.Verify(id, code, true)
}
