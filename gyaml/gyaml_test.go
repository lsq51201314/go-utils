package gyaml

import (
	"fmt"
	"testing"
)

func TestGyaml(t *testing.T) {
	type sub struct {
		AAA int64   `json:"a"`
		BBB float64 `json:"b"`
		CCC string  `json:"c"`
	}
	type obj struct {
		AAA int64   `json:"a"`
		BBB float64 `json:"b"`
		CCC string  `json:"c"`
		DDD []sub   `json:"d"`
	}
	type config struct {
		AAA int64   `json:"a"`
		BBB float64 `json:"b"`
		CCC string  `json:"c"`
		DDD obj     `json:"d"`
	}

	var p config
	Load(&p, "./config.yaml")
	Save(p, "./config2.yaml")
	fmt.Println(p)
}
