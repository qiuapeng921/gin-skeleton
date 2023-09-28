package servers

import (
	"context"
	"micro-base/internal/pkg/core/ctx"
	"micro-base/internal/pkg/core/log"
	"micro-base/internal/pkg/core/naming"
	"micro-base/internal/pkg/helper"
	"os"
	"time"
)

// ServerGroup 服务组实现接口
type ServerGroup interface {
	Start()
	ListenAndServe(cleans ...func(ctx.Context))
	Shutdown(ctx ctx.Context) error
	Add(servers ...Server) ServerGroup
}

// Server 服务接口
type Server interface {
	// ListenAndServe 要求同步启动服务，服务需要阻塞等待
	ListenAndServe() error
	Shutdown(ctx context.Context) error
}

// NamedServer 命名服务
type NamedServer interface {
	Server
	naming.Namer
}

// serverGroup 服务组
type serverGroup []Server

// Start 启动服务
func (sg serverGroup) Start() {
	c := ctx.New()
	for i := 0; i < len(sg); i++ {
		go func(server Server) {
			name := ""
			if ns, ok := server.(NamedServer); ok {
				name = ns.Name()
			}
			log.Info(c).Msgf("server running at %v", name)
			if err := server.ListenAndServe(); err != nil {
				log.Err(c, err).Msgf("server running error", err)
			}
		}(sg[i])
	}
}

// ListenAndServe 启动服务，同时监听退出
func (sg serverGroup) ListenAndServe(cleans ...func(ctx.Context)) {
	sg.Start()

	// 处理系统退出信号
	helper.ObserveExitSignal(func(signal os.Signal) {
		c := ctx.New()
		log.Info(c).Msg("Shutdown Server ...")

		cc, cancel := context.WithTimeout(c, 15*time.Second)
		defer cancel()
		if err := sg.Shutdown(ctx.Wrap(cc)); err != nil {
			log.Err(c, err).Msgf("Server Shutdown Error: %v", err)
		}

		for _, cf := range cleans {
			cf(c)
		}
		log.Info(c).Msg("Server exiting")
	})
}

// Shutdown 停止服务
func (sg serverGroup) Shutdown(c ctx.Context) error {
	for i := 0; i < len(sg); i++ {
		s := sg[i]
		name := ""
		if ns, ok := s.(NamedServer); ok {
			name = ns.Name()
		}
		log.Info(c).Msgf("%v Server shutdown", name)
		if err := s.Shutdown(c); err != nil {
			log.Err(c, err).Msg("Server shutdown: %v")
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
