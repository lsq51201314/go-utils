package http

//删除请求
func Delete(requrl string, params, header map[string]string, data []byte, timeout ...int) ([]byte, error) {
	to := 10
	if len(timeout) > 0 && timeout[0] > 0 {
		to = timeout[0]
	}
	return request(requrl, "DELETE", params, header, data, to)
}
