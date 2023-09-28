package cmd

import (
	_ "github.com/go-sql-driver/mysql"
	"micro-base/internal/config"
	"micro-base/internal/pkg/core/log"
	"micro-base/internal/pkg/helper"
)

// RegisterConfig 加载配置
func RegisterConfig(logToStderr bool) {
	config.Init()
	helper.Must(log.Init(config.CfgData.Log(logToStderr)))
}

// ExitHandle 处理系统退出信号
func ExitHandle() {

}
