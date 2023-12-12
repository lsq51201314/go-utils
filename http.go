package utils

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"time"
)

type Http struct{}

func (h Http) Get(url string) (code int, body []byte, err error) {
	client := http.Client{
		Timeout: time.Second * 5,
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

func (h Http) Post(url string, data interface{}) (code int, body []byte, err error) {
	client := http.Client{
		Timeout: time.Second * 5,
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
