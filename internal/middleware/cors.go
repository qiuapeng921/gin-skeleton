package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"micro-base/internal/config"
)

// CorsOption cors 配置方法
type CorsOption func(*cors.Config)

// Cors gin cors config
func Cors(config config.Cors, options ...CorsOption) gin.HandlerFunc {
	cfg := cors.DefaultConfig()
	cfg.AllowAllOrigins = config.AllowAllOrigins
	cfg.AllowOrigins = config.AllowOrigins
	cfg.AllowHeaders = config.AllowHeaders
	cfg.AllowMethods = config.AllowMethods
	cfg.AllowCredentials = config.AllowCredentials
	cfg.ExposeHeaders = config.ExposeHeaders
	cfg.AllowWebSockets = config.AllowWebSockets
	cfg.AllowFiles = config.AllowFiles
	cfg.MaxAge = config.MaxAge
	for _, option := range options {
		option(&cfg)
	}
	return func(context *gin.Context) {
		cors.New(cfg)
	}
}

// AllowAllOrigins 是否允许跨域
func AllowAllOrigins(allow bool) CorsOption {
	return func(config *cors.Config) {
		config.AllowAllOrigins = allow
	}
}

// AllowOrigins 跨域域名
func AllowOrigins(origins []string) CorsOption {
	return func(config *cors.Config) {
		config.AllowOrigins = origins
	}
}
