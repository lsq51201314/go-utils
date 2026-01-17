package glog

import (
	"fmt"
	"time"

	"github.com/lsq51201314/go-utils/gtime"
)

// 信息
func Info(text string) {
	fmt.Printf("%s\t%s\t%s\n",
		time.Now().In(gtime.Location()).Format("2006-01-02 15:04:05.000000"),
		"INFO",
		text)
}

// 错误
func Error(text string) {
	fmt.Printf("%s\t%s\t%s\n",
		time.Now().In(gtime.Location()).Format("2006-01-02 15:04:05.000000"),
		"ERROR",
		text)
}
