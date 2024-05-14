package utils

import (
	"fmt"
	"time"
)

// 日志显示
type Logger struct{}

// 信息显示
func (l Logger) Info(text string) {
	fmt.Printf("[%s] (%s) * %s\n",
		time.Now().In(location).Format("2006-01-02 15:04:05.999999"),
		"msg",
		text)
}

// 错误显示
func (l Logger) Error(text string, err error) {
	fmt.Printf("[%s] (%s) * %s:%v\n",
		time.Now().In(location).Format("2006-01-02 15:04:05.999999"),
		"err",
		text, err)
}
