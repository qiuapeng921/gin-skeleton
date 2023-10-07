package servers

import (
	"context"
	"errors"
	"micro-base/internal/pkg/core/ctx"
	"micro-base/internal/pkg/core/log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// ServerGroup 服务组实现接口
type ServerGroup interface {
	Start(c ctx.Context)
	ListenAndServe(c ctx.Context, cleans ...func(ctx.Context))
	Shutdown(ctx ctx.Context) error
	Add(servers ...Server) ServerGroup
}

// Server 服务接口
type Server interface {
	ListenAndServe() error
	Shutdown(ctx context.Context) error
}

type Named interface {
	Name() string
}

// NamedServer 命名服务
type NamedServer interface {
	Server
	Named
}

// serverGroup 服务组
type serverGroup []Server

// Start 启动服务
func (sg serverGroup) Start(c ctx.Context) {
	for _, server := range sg {
		go func(server Server) {
			name := ""
			if ns, ok := server.(NamedServer); ok {
				name = ns.Name()
			}
			log.Info(c).Msgf("server running at %v", name)

			if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
				log.Err(c, err).Msgf("server running error: %v", err)
			}
		}(server)
	}
}

// ListenAndServe 启动服务，同时监听退出
func (sg serverGroup) ListenAndServe(c ctx.Context, cleans ...func(ctx.Context)) {
	sg.Start(c)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)

	<-quit
	log.Info(c).Msg("Shutdown Server ...")

	cc, cancel := context.WithTimeout(c, 15*time.Second)
	defer cancel()
	if err := sg.Shutdown(ctx.Wrap(cc)); err != nil {
		log.Err(c, err).Msgf("Server Shutdown Error: %v", err)
	}

	for _, clean := range cleans {
		clean(c)
	}
	log.Info(c).Msg("Server exiting")
}

// Shutdown 停止服务
func (sg serverGroup) Shutdown(c ctx.Context) error {
	for _, server := range sg {
		name := ""
		if ns, ok := server.(NamedServer); ok {
			name = ns.Name()
		}
		log.Info(c).Msgf("%v Server shutdown", name)

		if err := server.Shutdown(c); err != nil {
			log.Err(c, err).Msg("Server shutdown error: %v")
		}
	}

	return nil
}

// Add 添加服务。需要接受返回值，可能返回一个新服务组对象
func (sg serverGroup) Add(servers ...Server) ServerGroup {
	return append(sg, servers...)
}

// Group 打包一个服务组
func Group(servers ...Server) ServerGroup {
	return serverGroup(servers)
}
