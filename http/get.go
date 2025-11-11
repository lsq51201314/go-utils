package http

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

//获取请求
func Get(requrl string, params, header map[string]string, timeout ...int) ([]byte, error) {
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
	to := 10
	if len(timeout) > 0 && timeout[0] > 0 {
		to = timeout[0]
	}
	client := &http.Client{
		Timeout: time.Duration(to) * time.Second,
	}
	req, err := http.NewRequest("GET", baseURL.String(), nil)
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
