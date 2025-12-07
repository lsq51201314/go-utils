package gconst

import (
	"fmt"
	"testing"
)

func TestGconst(t *testing.T) {
	type ErrorTypeInfo struct {
		Methods
		SysApiNotFound Option `id:"404" text:"无效接口"`
		CodeVerifyErr  Option `id:"1101" text:"验证码输入有误"`
		CodeLockErr    Option `id:"1102" text:"错误次数上限已被封禁，请等待%s秒后重试"`
		LoginUserErr   Option `id:"1103" text:"用户名或密码错误，剩余登录次数%s次"`
	}
	var ErrorType ErrorTypeInfo
	AutoUnmarshal(&ErrorType)
	fmt.Println(ErrorType.SysApiNotFound.ID)
	fmt.Println(ErrorType.CodeVerifyErr.Text)
	fmt.Println(ErrorType.CodeLockErr.FormatText("10"))
	fmt.Println(ErrorType.LoginUserErr.FormatText("20"))
}
