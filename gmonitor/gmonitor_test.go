package gmonitor

import (
	"fmt"
	"os"
	"testing"
)

func TestGmonitor(t *testing.T) {
	path, _ := os.Getwd()
	//必须是完整的路径，不可用 ./
	gm, _ := New(path+"\\test.txt", func(file string, err error) {
		fmt.Println(file, err)
	})
	gm.Run()
}
