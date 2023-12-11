package utils

import (
	"fmt"
	"time"
)

type Logger struct{}

func (l Logger) Info(text string) {
	fmt.Printf("[%s] (%s) * %s\n",
		time.Now().In(location).Format("2006-01-02 15:04:05.999999"),
		"msg",
		text)
}

func (l Logger) Error(text string, err error) {
	fmt.Printf("[%s] (%s) * %s:%v\n",
		time.Now().In(location).Format("2006-01-02 15:04:05.999999"),
		"err",
		text, err)
}
