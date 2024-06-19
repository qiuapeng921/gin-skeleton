package cmd

import (
	_ "github.com/go-sql-driver/mysql"
	"micro-base/internal/pkg/core/ctx"

	"micro-base/internal/config"
	"micro-base/internal/pkg/core/db"
	"micro-base/internal/pkg/core/log"
	"micro-base/internal/pkg/core/rdb"
	"micro-base/internal/pkg/helper"
)

// RegisterConfig 加载配置
func RegisterConfig(c ctx.Context, configFile string) {
	helper.Must(config.Init(configFile))
	helper.Must(log.Init(config.CfgData.Log()))

	helper.Must(db.InitConnection(c, config.CfgData.DB))
	helper.Must(rdb.InitConnection(c, config.CfgData.RDB))
}
