package ghttp

import (
	"fmt"
	"log"
	"testing"
)

func TestGhttp(t *testing.T) {
	if code, buf, err := Get("http://127.0.0.1:22345/api/verify", nil, nil); err != nil {
		log.Fatalln(err)
	} else {
		fmt.Println(code, string(buf))
	}

	if code, buf, err := Post("http://127.0.0.1:22345/api/login", nil, nil, []byte("hello")); err != nil {
		log.Fatalln(err)
	} else {
		fmt.Println(code, string(buf))
	}

}
