package ghttp

import (
	"bufio"
	"bytes"
	"net/http"
	"net/url"
	"time"
)

// 回调函数
type CallBack func(line string)

// 流式请求 (POST)
func Stream(requrl string, params, header map[string]string, data []byte, cfunc CallBack, timeout ...int) (int, []byte, error) {
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
	client := &http.Client{}
	if len(timeout) > 0 && timeout[0] > 0 {
		client.Timeout = time.Duration(timeout[0]) * time.Second
	}
	req, err := http.NewRequest("POST", baseURL.String(), bytes.NewBuffer(data))
	if err != nil {
		return 0, nil, err
	}
	//req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/142.0.0.0 Safari/537.36")
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
	//处理消息
	scanner := bufio.NewScanner(resp.Body)
	str := ""
	for scanner.Scan() {
		line := scanner.Text()
		str += line
		if cfunc != nil {
			cfunc(line)
		}
	}
	return resp.StatusCode, []byte(str), nil
}
