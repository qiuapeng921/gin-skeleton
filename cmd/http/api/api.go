package api

import (
	v1 "micro-base/internal/app/router/v1"
	"micro-base/internal/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
	"micro-base/internal/config"
)

func Api(mode string) http.Handler {
	gin.SetMode(mode)
	router := gin.New()
	router.Use(middleware.Recovery())
	router.Use(middleware.Access())
	router.Use(middleware.Logger())
	if config.CfgData.Restful.Cors.Enable {
		router.Use(middleware.Cors())
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
