package main

import (
	"flag"
	"github.com/gin-gonic/gin/binding"
	"micro-base/cmd"
	"micro-base/cmd/http/api"
	"micro-base/internal/config"
	"micro-base/internal/pkg/core/ctx"
	"micro-base/internal/pkg/core/ginplus"
	"micro-base/internal/pkg/core/log"
	"micro-base/internal/pkg/core/servers"
	"net/http"
)

var flagConf string

func init() {
	flag.StringVar(&flagConf, "conf", config.DefaultConfigFile, "config path, eg: -conf app.yaml")
	// 汉化参数验证器
	binding.Validator = ginplus.NewValidator()
}

func main() {
	c := ctx.New()

	// 注册配置
	cmd.RegisterConfig(c, flagConf)

	serverWrapper := &servers.HTTPServerWrapper{
		Server: &http.Server{
			Addr:    config.CfgData.Restful.Addr,
			Handler: api.Api(config.CfgData.Mode),
		},
		Named: config.CfgData.Restful.Addr,
	}

	serverGroup := servers.Group(serverWrapper)

	serverGroup.ListenAndServe(c, func(context ctx.Context) {
		log.Info(c).Msg("服务已经停止")
	})
}
