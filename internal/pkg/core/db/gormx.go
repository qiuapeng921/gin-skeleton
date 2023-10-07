package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"io"
	"log"
	"micro-base/internal/config"
	"os"
	"time"
)

// Config 配置
type Config struct {
	// dsn
	Dsn string
	// 最大空闲连接数
	MaxIdleConnections int
	// 最大打开连接数
	MaxOpenConnections int
	// 最长活跃时间
	MaxLifeTime int
}

func connect(conf *Config) (*gorm.DB, error) {
	var directory gorm.Dialector

	directory = mysql.Open(conf.Dsn)

	conn, err := gorm.Open(directory, &gorm.Config{
		SkipDefaultTransaction: true,
		Logger:                 multiLogger(),
	})
	if err != nil {
		return nil, err
	}

	sqlDB, err := conn.DB()
	sqlDB.SetMaxOpenConns(conf.MaxOpenConnections)
	sqlDB.SetConnMaxLifetime(time.Duration(conf.MaxLifeTime))
	sqlDB.SetMaxIdleConns(conf.MaxIdleConnections)

	return conn, err
}

func multiLogger() logger.Interface {
	var output io.Writer
	if config.CfgData.Logger.TargetType == "file" {
		file, err := os.OpenFile(config.CfgData.Logger.Target, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
		if err != nil {
			fmt.Errorf("failed to open log file: %v", err)
		}
		output = io.MultiWriter(file)
	} else {
		output = os.Stdout
	}

	// 默认为不显示sql日志
	gormLogger := logger.Silent
	if config.CfgData.DB.Log {
		gormLogger = logger.Info
	}

	return logger.New(log.New(output, "\r\n", log.LstdFlags), logger.Config{
		SlowThreshold:             time.Second, // 慢 SQL 阈值
		LogLevel:                  gormLogger,  // 日志级别
		IgnoreRecordNotFoundError: true,        // 忽略ErrRecordNotFound（记录未找到）错误
		Colorful:                  false,       // 禁用彩色打印
	})
}
