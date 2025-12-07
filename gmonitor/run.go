package gmonitor

import (
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/lsq51201314/go-utils/gtime"
)

// 运行
func (m *GMonitor) Run() {
	for {
		select {
		case ev := <-m.watch.Events:
			{
				if ev.Op&fsnotify.Write == fsnotify.Write {
					if ev.Name == m.path {
						nt := time.Now().In(gtime.Location).Unix()
						if nt-m.lastTime > 3 { //大于3秒
							if m.callback != nil {
								//延迟3秒(让文件彻底写入释放)
								time.Sleep(time.Duration(3) * time.Second)
								m.callback(m.path, nil)
							}
						}
						m.lastTime = nt
					}
				}
			}
		case err := <-m.watch.Errors:
			{
				if m.callback != nil {
					m.callback("", err)
				}
			}
		}
	}
}
