package gzlib

import (
	"encoding/hex"
	"fmt"
	"testing"
)

func TestGzlib(t *testing.T) {
	buf := []byte("Hello World!")
	src := Compress(buf)
	fmt.Println(hex.EncodeToString(src))
	fmt.Println(string(UnCompress(src)))
}
