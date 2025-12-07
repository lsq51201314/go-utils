package gexit

import (
	"context"
	"testing"
	"time"

	"github.com/lsq51201314/go-utils/glog"
)

func TestGexit(t *testing.T) {
	glog.Info("运行中...")
	WaitExit(func(ctx context.Context) {
		glog.Info("正在退出")
		glog.Info("3秒...")
		time.Sleep(1 * time.Second)
		glog.Info("2秒...")
		time.Sleep(1 * time.Second)
		glog.Info("1秒...")
		time.Sleep(1 * time.Second)
		glog.Info("成功退出")
	})
}
