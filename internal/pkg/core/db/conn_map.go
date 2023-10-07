package db

import (
	"gorm.io/gorm"
	"micro-base/internal/config"
	"micro-base/internal/pkg/core/ctx"
	"micro-base/internal/pkg/core/log"
	"sync"
)

var databasesConn sync.Map

func InitConnection(c ctx.Context, conf config.MysqlConfig) error {
	for alias, dns := range conf.DnsList {
		microDatabase := &Config{
			Dsn:                dns,
			MaxOpenConnections: conf.Pool.MaxOpen,
			MaxIdleConnections: conf.Pool.MaxIdle,
			MaxLifeTime:        int(conf.Pool.MaxLifeTime),
		}
		conn, err := connect(microDatabase)
		if err != nil {
			return err
		}
		databasesConn.Store(alias, conn)
		log.Debug(c).Msgf("db %s 连接成功", alias)
	}

	return nil
}

// GetConnection 获取数据库连接
func GetConnection(name string) *gorm.DB {
	// 判断map中是否存在数据库链接对象key
	conn, ok := databasesConn.Load(name)
	if !ok {
		return nil
	}

	return conn.(*gorm.DB)
}
