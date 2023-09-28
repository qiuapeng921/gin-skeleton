package helper

import (
	"micro-base/internal/pkg/core/ctx"
	"micro-base/internal/pkg/core/log"
	"os"
	"os/signal"
	"syscall"
)

// ObserveExitSignal 监听系统退出信号
// 会阻塞当前运行
func ObserveExitSignal(f func(os.Signal)) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		log.Info(ctx.New()).Msgf("get a signal %s", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			f(s)
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
