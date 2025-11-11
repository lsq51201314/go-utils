package monitor

import (
	"errors"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/lsq51201314/go-utils/file"
	"github.com/lsq51201314/go-utils/location"
)

type Monitor struct {
	path     string
	watch    *fsnotify.Watcher
	lastTime int64
	callback CallBack
}

// 新建监视
func New(path string) (monitor *Monitor, err error) {
	if !file.IsExist(path) {
		err = errors.New("文件不存在:" + path)
		return
	}
	obj := &Monitor{
		path:     path,
		lastTime: time.Now().In(location.Location).Unix(),
	}
	//监视文件
	if obj.watch, err = fsnotify.NewWatcher(); err != nil {
		return
	}
	//加入文件
	if err = obj.watch.Add(path); err != nil {
		return
	}
	return obj, nil
}
