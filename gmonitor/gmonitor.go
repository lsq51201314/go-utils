package gmonitor

import (
	"errors"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/lsq51201314/go-utils/gfile"
	"github.com/lsq51201314/go-utils/gtime"
)

type GMonitor struct {
	path     string
	watch    *fsnotify.Watcher
	lastTime int64
	callback CallBack
}

// 新建监视 延迟3秒触发
func New(path string, cfun CallBack) (*GMonitor, error) {
	if !gfile.IsExist(path) {
		return nil, errors.New("文件不存在:" + path)
	}
	obj := &GMonitor{
		path:     path,
		lastTime: time.Now().In(gtime.Location()).Unix(),
		callback: cfun,
	}
	var err error
	//监视文件
	if obj.watch, err = fsnotify.NewWatcher(); err != nil {
		return nil, err
	}
	//加入文件
	if err = obj.watch.Add(path); err != nil {
		return nil, err
	}
	return obj, nil
}
