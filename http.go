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
func (h Http) Get(url string, header ...map[string][]string) (code int, body []byte, rspheader map[string][]string, err error) {
	client := http.Client{
		Timeout: time.Second * time.Duration(10), //10秒
	}
	var req *http.Request
	if req, err = http.NewRequestWithContext(context.Background(), http.MethodGet, url, nil); err != nil {
		return
	}
	if len(header) > 0 {
		for k, v := range header[0] {
			val := ""
			for _, str := range v {
				val += str + ";"
			}
			req.Header.Set(k, val)
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
	rspheader = rep.Header
	return
}

// post请求
func (h Http) Post(url string, data interface{}, header ...map[string][]string) (code int, body []byte, rspheader map[string][]string, err error) {
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
		for k, v := range header[0] {
			val := ""
			for _, str := range v {
				val += str + ";"
			}
			req.Header.Set(k, val)
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
	rspheader = rep.Header
	return
}
