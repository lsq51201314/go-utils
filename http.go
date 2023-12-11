package utils

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

type Http struct{}

func (h Http) Get(url string) (body []byte, err error) {
	var resp *http.Response
	if resp, err = http.Get(url); err != nil {
		return
	}
	defer resp.Body.Close()
	body, err = io.ReadAll(resp.Body)
	return
}

func (h Http) Post(url string, data interface{}) (body []byte, err error) {
	var b []byte
	if b, err = json.Marshal(data); err != nil {
		return
	}
	var resp *http.Response
	if resp, err = http.Post(url,
		"application/json;charset=UTF-8",
		strings.NewReader(string(b))); err != nil {
		return
	}
	defer resp.Body.Close()
	body, err = io.ReadAll(resp.Body)
	return
}
