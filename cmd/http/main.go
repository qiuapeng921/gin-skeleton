package main

import (
	"github.com/gin-gonic/gin/binding"
	"gitlab-ce.k8s.tools.vchangyi.com/common/go-toolbox/ginplus"
	"gitlab-ce.k8s.tools.vchangyi.com/common/go-toolbox/log"
	"gitlab-ce.k8s.tools.vchangyi.com/common/go-toolbox/monitor"
	"net/http"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"gitlab-ce.k8s.tools.vchangyi.com/common/go-toolbox/ctx"
	"gitlab-ce.k8s.tools.vchangyi.com/common/go-toolbox/servers"
	"micro-base/cmd"
	"micro-base/cmd/http/api"
	"micro-base/internal/config"
)

// Version ...
var Version = "dev"

// Build ...
var Build = "now"

func init() {
	// 汉化参数验证器
	binding.Validator = ginplus.NewValidator()
}

func main() {
	start := time.Now() // 获取当前时间

	// 解析命令行参数
	cmd.ParseConfigFromCmd(os.Args)
	// 注册配置
	cmd.RegisterConfig(false)

	serverGroup := servers.Group(&http.Server{
		Addr:    config.CfgData.Restful.Addr,
		Handler: api.Api(config.CfgData.Mode),
	})
	if config.CfgData.Monitor.Enable {
		serverGroup = serverGroup.Add(monitor.HTTP(config.CfgData.Restful.BasePath, config.CfgData.Monitor.Addr))
	}

	elapsed := time.Since(start)

	log.Info(ctx.New()).Msgf("服务启动用时：%+v", elapsed)

	serverGroup.ListenAndServe(func(context ctx.Context) {
		cmd.ExitHandle()
	})
}
