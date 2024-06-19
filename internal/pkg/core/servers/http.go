package servers

import (
	"context"
	"net/http"
)

// HTTPServerWrapper 包装http.Server以实现Named接口
type HTTPServerWrapper struct {
	Server *http.Server
	Named  string
}

// Name 实现Named接口的方法
func (s *HTTPServerWrapper) Name() string {
	return s.Named
}

// ListenAndServe 启动HTTP服务器
func (s *HTTPServerWrapper) ListenAndServe() error {
	return s.Server.ListenAndServe()
}

// Shutdown 关闭HTTP服务器
func (s *HTTPServerWrapper) Shutdown(ctx context.Context) error {
	return s.Server.Shutdown(ctx)
}
