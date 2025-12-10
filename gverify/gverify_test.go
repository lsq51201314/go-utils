package gverify

import (
	"fmt"
	"testing"
)

func TestGverify(t *testing.T) {
	v := New()
	data, _ := v.Get()
	fmt.Println(data)
	fmt.Println(v.Verify("X49Z1ctBsF7x3hTDEAbF", "24051"))
}
