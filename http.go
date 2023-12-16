package utils

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"time"
)

// http实例
type Http struct{}

// get请求
func (h Http) Get(url string, timeout ...int) (code int, body []byte, err error) {
	to := 5
	if len(timeout) > 0 {
		to = timeout[0]
	}
	client := http.Client{
		Timeout: time.Second * time.Duration(to),
	}
	var req *http.Request
	if req, err = http.NewRequestWithContext(context.Background(), http.MethodGet, url, nil); err != nil {
		return
	}
	var rep *http.Response
	if rep, err = client.Do(req); err != nil {
		return
	}
	if body, err = io.ReadAll(rep.Body); err != nil {
		return
	}
	defer rep.Body.Close()
	code = rep.StatusCode
	return
}

// post请求
func (h Http) Post(url string, data interface{}, timeout ...int) (code int, body []byte, err error) {
	to := 5
	if len(timeout) > 0 {
		to = timeout[0]
	}
	client := http.Client{
		Timeout: time.Second * time.Duration(to),
	}
	var dataByte []byte
	if dataByte, err = json.Marshal(data); err != nil {
		return
	}
	bodyReader := bytes.NewReader(dataByte)

	var req *http.Request
	if req, err = http.NewRequestWithContext(context.Background(), http.MethodPost, url, bodyReader); err != nil {
		return
	}
	req.Header.Set("Content-Type", "application/json;charset=UTF-8")
	var rep *http.Response
	if rep, err = client.Do(req); err != nil {
		return
	}
	if body, err = io.ReadAll(rep.Body); err != nil {
		return
	}
	defer rep.Body.Close()
	code = rep.StatusCode
	return
}
