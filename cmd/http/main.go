package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"micro-base/cmd"
	"micro-base/cmd/http/api"
	"micro-base/internal/config"
	"micro-base/internal/pkg/core/ctx"
	"micro-base/internal/pkg/core/ginplus"
	"micro-base/internal/pkg/core/log"
	"micro-base/internal/pkg/core/servers"
	"net/http"
	"runtime"
	"time"
)

var flagConf string

func init() {
	flag.StringVar(&flagConf, "conf", config.DefaultConfigFile, "config path, eg: -conf app.yaml")
	// 汉化参数验证器
	binding.Validator = ginplus.NewValidator()
}

func main() {
	c := ctx.New()

	start := time.Now() // 获取当前时间

	// 注册配置
	cmd.RegisterConfig(c)

	serverGroup := servers.Group(&http.Server{
		Addr:    config.CfgData.Restful.Addr,
		Handler: api.Api(config.CfgData.Mode),
	})

	elapsed := time.Since(start)

	log.Info(c).Msgf("服务启动用时：%v", elapsed)
	fmt.Println(fmt.Sprintf("Server      Name:     %s", config.CfgData.App))
	fmt.Println(fmt.Sprintf("System      Name:     %s", runtime.GOOS))
	fmt.Println(fmt.Sprintf("Go          Version:  %s", runtime.Version()))
	fmt.Println(fmt.Sprintf("Gin         Version:  %s", gin.Version))
	fmt.Println(fmt.Sprintf("Listen      Address:  %s", "http://"+config.CfgData.Restful.Addr))

	serverGroup.ListenAndServe(c, func(context ctx.Context) {

	})
}
