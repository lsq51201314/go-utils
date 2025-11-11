package monitor

type CallBack func(file string, err error)

// 设置回调
func (m *Monitor) SetCallback(cfunc CallBack) {
	m.callback = cfunc
}
