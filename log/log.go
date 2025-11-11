package log

import (
	"fmt"
	"time"

	"github.com/lsq51201314/go-utils/location"
)

// 信息
func Info(text string) {
	fmt.Printf("[%s] (%s) * %s\n",
		time.Now().In(location.Location).Format("2006-01-02 15:04:05.000000"),
		"msg",
		text)
}

// 错误
func Error(text string) {
	fmt.Printf("[%s] (%s) * %s\n",
		time.Now().In(location.Location).Format("2006-01-02 15:04:05.000000"),
		"err",
		text)
}
