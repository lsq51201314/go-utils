package gexit

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// 退出回调
type ExitFunc func(ctx context.Context)

// Ctrl+C 30秒退出
func WaitExit(cfun ...ExitFunc) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	//执行退出程序
	for _, cf := range cfun {
		if cf != nil {
			cf(ctx)
		}
	}
}
