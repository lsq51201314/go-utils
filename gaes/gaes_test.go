package gaes

import (
	"encoding/hex"
	"fmt"
	"testing"
)

func TestGaes(t *testing.T) {
	buf := []byte("Hello World!")
	src := Enc(buf)
	fmt.Println(hex.EncodeToString(src))
	fmt.Println(string(Dec(src)))
}
