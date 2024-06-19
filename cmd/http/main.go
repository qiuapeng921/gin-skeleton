package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"micro-base/cmd"
	v1 "micro-base/internal/app/router/v1"
	"micro-base/internal/config"
	"micro-base/internal/middleware"
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
			Handler: api(config.CfgData.Mode),
		},
		Named: config.CfgData.Restful.Addr,
	}

	serverGroup := servers.Group(serverWrapper)

	serverGroup.ListenAndServe(c, func(context ctx.Context) {
		log.Info(c).Msg("服务已经停止")
	})
}

func api(mode string) http.Handler {
	gin.SetMode(mode)
	router := gin.New()
	router.Use(middleware.Recovery())
	router.Use(middleware.Access())
	router.Use(middleware.Logger())

	if config.CfgData.Restful.Cors.Enable {
		router.Use(middleware.Cors(config.CfgData.Restful.Cors))
	}

	router.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, config.CfgData.App+"-"+config.CfgData.Env)
	})

	// 心跳检测地址
	router.GET("/heart-beat", func(context *gin.Context) {
		context.String(http.StatusOK, "true")
	})

	microCrm := router.Group(config.CfgData.Restful.BasePath)
	v1.InitRouter(microCrm)

	return router
}
