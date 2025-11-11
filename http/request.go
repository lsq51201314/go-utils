package http

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

func request(requrl string, method string, params, header map[string]string, data []byte, timeout int) ([]byte, error) {
	baseURL, err := url.Parse(requrl)
	if err != nil {
		return nil, err
	}
	if len(params) > 0 {
		pq := url.Values{}
		for k, v := range params {
			pq.Add(k, v)
		}
		baseURL.RawQuery = pq.Encode()
	}
	client := &http.Client{
		Timeout: time.Duration(timeout) * time.Second,
	}
	req, err := http.NewRequest(method, baseURL.String(), bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/142.0.0.0 Safari/537.36")
	for k, v := range header {
		req.Header.Set(k, v)
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	//处理状态
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("错误状态:%d", resp.StatusCode)
	}
	return io.ReadAll(resp.Body)
}
