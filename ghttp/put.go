package ghttp

//修改请求
func Put(requrl string, params, header map[string]string, data []byte, timeout ...int) (int,[]byte, error) {
	to := 10
	if len(timeout) > 0 && timeout[0] > 0 {
		to = timeout[0]
	}
	return request(requrl, "PUT", params, header, data, to)
}
