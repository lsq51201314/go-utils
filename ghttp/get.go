package ghttp

import (
	"io"
	"net/http"
	"net/url"
	"time"
)

// 获取请求
func Get(requrl string, params, header map[string]string, timeout ...int) (int, []byte, error) {
	baseURL, err := url.Parse(requrl)
	if err != nil {
		return 0, nil, err
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
		return 0, nil, err
	}
	for k, v := range header {
		req.Header.Set(k, v)
	}
	resp, err := client.Do(req)
	if err != nil {
		return 0, nil, err
	}
	defer resp.Body.Close()
	//处理状态
	if resp.StatusCode != 200 {
		return resp.StatusCode, nil, nil
	}
	buf, err := io.ReadAll(resp.Body)
	if err != nil {
		return resp.StatusCode, nil, err
	}
	return resp.StatusCode, buf, nil
}
