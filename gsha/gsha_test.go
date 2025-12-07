package gsha

import (
	"encoding/hex"
	"fmt"
	"testing"
)

func TestGsha(t *testing.T) {
	buf := []byte("Hello World!")
	fmt.Println(hex.EncodeToString(Sha1(buf)))
	fmt.Println(hex.EncodeToString(Sha256(buf)))
	fmt.Println(Sha1Str(buf))
	fmt.Println(Sha256Str(buf))
}
