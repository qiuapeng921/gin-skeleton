package rdb

import (
	"github.com/go-redis/redis/v8"
	"micro-base/internal/config"
	"micro-base/internal/pkg/core/ctx"
	"micro-base/internal/pkg/core/log"
	"runtime"
	"time"
)

var client *redis.Client

func InitConnection(c ctx.Context, cfg config.RedisConfig) error {
	if cfg.PoolSize < runtime.NumCPU() {
		cfg.PoolSize = runtime.NumCPU() * 10
	}
	if cfg.MinIdle < runtime.NumCPU() {
		cfg.MinIdle = runtime.NumCPU()
	}

	client = redis.NewClient(&redis.Options{
		Network:         "tcp",
		Addr:            cfg.Addr,
		Username:        cfg.User,
		Password:        cfg.Pwd,
		DB:              cfg.DB,
		MaxRetries:      0,
		MinRetryBackoff: 8,
		MaxRetryBackoff: 512 * time.Millisecond,
		DialTimeout:     cfg.DialTimeout * time.Second,
		ReadTimeout:     cfg.CmdTimeout * time.Second,
		WriteTimeout:    cfg.CmdTimeout * time.Second,
		PoolSize:        cfg.PoolSize,
		MinIdleConns:    cfg.MinIdle,
	})

	_, err := client.Ping(ctx.New()).Result()
	if err != nil {
		return err
	}

	log.Debug(c).Msgf("redis连接成功")
	return nil
}

// Client 返回 redis 客户端
func Client() *redis.Client {
	return client
}
