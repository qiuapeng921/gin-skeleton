package main

import (
	"flag"
	"log"
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
)

func main() {
	// 定义命令行参数
	target := flag.String("target", "https://hub.docker.com", "The target URL to proxy to")
	port := flag.String("port", "8081", "The port to listen on")
	flag.Parse()

	// 解析目标服务器的地址
	targetURL, err := url.Parse(*target)
	if err != nil {
		log.Fatal("Failed to parse target URL:", err)
	}

	// 创建一个反向代理
	proxy := httputil.NewSingleHostReverseProxy(targetURL)

	// 使用Gin创建一个路由
	r := gin.Default()
	gin.SetMode(gin.ReleaseMode)

	// 设置代理路由
	r.Any("/*proxyPath", func(c *gin.Context) {
		// 修改请求的Host头为目标服务器的Host
		c.Request.Host = targetURL.Host

		// 转发请求到目标服务器
		proxy.ServeHTTP(c.Writer, c.Request)
	})

	// 启动Gin服务器
	log.Printf("Starting %s proxy server on :%s\n", *target, *port)
	if err := r.Run(":" + *port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
