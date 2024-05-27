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

type HttpHeader struct {
	Key  string
	Name string
}

// get请求
func (h Http) Get(url string, header ...HttpHeader) (code int, body []byte, err error) {
	client := http.Client{
		Timeout: time.Second * time.Duration(10), //10秒
	}
	var req *http.Request
	if req, err = http.NewRequestWithContext(context.Background(), http.MethodGet, url, nil); err != nil {
		return
	}
	if len(header) > 0 {
		for _, v := range header {
			req.Header.Set(v.Key, v.Name)
		}
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
func (h Http) Post(url string, data interface{}, header ...HttpHeader) (code int, body []byte, err error) {
	client := http.Client{
		Timeout: time.Second * time.Duration(10), //10秒
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
	if len(header) > 0 {
		for _, v := range header {
			req.Header.Set(v.Key, v.Name)
		}
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
