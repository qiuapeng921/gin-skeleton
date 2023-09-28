package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// CorsOption cors 配置方法
type CorsOption func(*cors.Config)

// Cors gin cors config
func Cors(options ...CorsOption) gin.HandlerFunc {
	config := cors.DefaultConfig()
	for _, option := range options {
		option(&config)
	}
	return func(context *gin.Context) {
		cors.New(config)
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
