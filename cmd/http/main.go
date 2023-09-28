package main

import (
	"context"
	"flag"
	"github.com/gin-gonic/gin/binding"
	"micro-base/cmd"
	"micro-base/cmd/http/api"
	"micro-base/internal/config"
	"micro-base/internal/pkg/core/ctx"
	"micro-base/internal/pkg/core/ginplus"
	"micro-base/internal/pkg/core/log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var flagConf string

func init() {
	flag.StringVar(&flagConf, "conf", config.DefaultConfigFile, "config path, eg: -conf app.yaml")
	// 汉化参数验证器
	binding.Validator = ginplus.NewValidator()
}

func main() {
	start := time.Now() // 获取当前时间

	// 注册配置
	cmd.RegisterConfig(false)

	serverGroup := http.Server{
		Addr:    config.CfgData.Restful.Addr,
		Handler: api.Api(config.CfgData.Mode),
	}

	elapsed := time.Since(start)

	cc := ctx.New()

	log.Info(cc).Msgf("服务启动用时：%+v", elapsed)

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	log.Info(cc).Msgf("Shutdown Server ...")

	cw, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := serverGroup.Shutdown(ctx.Wrap(cw)); err != nil {
		log.Error(cc).Msgf("Server Shutdown:", err)
	}
	log.Info(cc).Msgf("Server exiting")
}
