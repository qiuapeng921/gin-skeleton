package app

import (
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
	"micro-base/internal/config"
	"micro-base/internal/pkg/core/db"
	"micro-base/internal/pkg/core/rdb"
)

func DB() *gorm.DB {
	return db.GetConnection("default")
}

// Redis redis 实例
func Redis() *redis.Client {
	return rdb.Client()
}

// Cfg 全局配置
func Cfg() *config.Config {
	return config.CfgData
}
