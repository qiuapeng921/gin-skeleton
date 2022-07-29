package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab-ce.k8s.tools.vchangyi.com/common/go-toolbox/docs"
	"gitlab-ce.k8s.tools.vchangyi.com/common/go-toolbox/docs/swagger"
	middleware "gitlab-ce.k8s.tools.vchangyi.com/common/go-toolbox/ginplus/mw"
	"micro-scrm/internal/config"
	interMiddleware "micro-scrm/internal/pkg/middleware"
)

// Api api 描述
func Api(mode string) http.Handler {
	api := docs.New(
		docs.Host(config.CfgData.Restful.Host),
		docs.BasePath(config.CfgData.Restful.BasePath),
		docs.Endpoints(endpoints()...),
		docs.Title(config.CfgData.Restful.Doc.Title),
		docs.Description(config.CfgData.Restful.Doc.Description),
		docs.ContactEmail(config.CfgData.Restful.Doc.Contact),
		docs.TermsOfService(config.CfgData.Restful.Doc.TermOfService),
	)

	return handler(api, mode)
}

// handler swagger.API 生成 http.Handler 服务
func handler(api *swagger.API, mode string) http.Handler {
	gin.SetMode(mode)
	router := gin.New()
	router.Use(middleware.Recovery())
	router.Use(interMiddleware.Access())
	router.Use(middleware.Logger())
	if config.CfgData.Restful.Cors.Enable {
		router.Use(middleware.Cors())
	}

	router.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, config.CfgData.App+"-"+config.CfgData.Env)
	})

	// 心跳检测地址
	router.GET(api.BasePath+"/heart-beat", func(context *gin.Context) {
		context.String(http.StatusOK, "true")
	})

	router.GET(api.DocEndpointPath(), gin.WrapH(api.Handler(true)))

	api.Walk(func(path string, endpoint *swagger.Endpoint) {
		h := endpoint.Handler.(func(c *gin.Context))
		path = docs.ColonPath(path)
		router.Handle(endpoint.Method, path, h)
	})

	return router
}
