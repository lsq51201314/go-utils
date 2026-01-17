package gsnowflake

import (
	"fmt"
	"testing"
)

func TestSnowflake(t *testing.T) {
	sf, _ := New("2026-01-01", 1)
	fmt.Println(sf.Get())
	for range 100 {
		fmt.Println(sf.Get())
	}
}
