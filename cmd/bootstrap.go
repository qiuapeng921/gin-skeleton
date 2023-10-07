package cmd

import (
	_ "github.com/go-sql-driver/mysql"
	"micro-base/internal/pkg/core/ctx"

	"micro-base/internal/config"
	"micro-base/internal/pkg/core/db"
	"micro-base/internal/pkg/core/log"
	"micro-base/internal/pkg/core/rdb"
	"micro-base/internal/pkg/core/utils"
	"micro-base/internal/pkg/helper"
)

// RegisterConfig 加载配置
func RegisterConfig(c ctx.Context) {
	config.Init()
	helper.Must(log.Init(config.CfgData.Log()))

	utils.Must(db.InitConnection(c, config.CfgData.DB))
	utils.Must(rdb.InitConnection(c, config.CfgData.RDB))
}

// ExitHandle 处理系统退出信号
func ExitHandle() {

}
