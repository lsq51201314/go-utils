package log

import (
	"fmt"
	"time"
)

var location, _ = time.LoadLocation("Asia/Shanghai")

// 信息
func Info(text string) {
	fmt.Printf("[%s] (%s) * %s\n",
		time.Now().In(location).Format("2006-01-02 15:04:05.000000"),
		"msg",
		text)
}

// 错误
func Error(text string) {
	fmt.Printf("[%s] (%s) * %s\n",
		time.Now().In(location).Format("2006-01-02 15:04:05.000000"),
		"err",
		text)
}
