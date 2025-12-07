package gpostgresql

import (
	"fmt"
	"testing"
)

func TestGpostgresql(t *testing.T) {
	_, err := New("127.0.0.1", 5432, "postgres", "123456", "test")
	fmt.Println(err)
}
