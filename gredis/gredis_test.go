package gredis

import (
	"fmt"
	"testing"
)

func TestGredis(t *testing.T) {
	rdb, _ := New("127.0.0.1", 6379, "123456")
	fmt.Println(rdb.Info())
}
