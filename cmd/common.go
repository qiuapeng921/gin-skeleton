package cmd

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"gitlab-ce.k8s.tools.vchangyi.com/common/go-toolbox/args"
	"gitlab-ce.k8s.tools.vchangyi.com/common/go-toolbox/helper"
	"gitlab-ce.k8s.tools.vchangyi.com/common/go-toolbox/log"
	"micro-base/internal/config"
	"os"
)

func ParseConfigFromCmd(arguments []string) {
	appArgs := args.New("micro-base",
		args.Store(config.CfgData),
		args.FileConfigEnabled("config", config.DefaultConfigFile, false, ""),
		args.HelpExit(0),
	)
	err := appArgs.Run(arguments)
	// 解析命令行参数
	if err != nil {
		fmt.Println("args error:", err.Error())
		os.Exit(1)
	}
}

// RegisterConfig 加载配置
func RegisterConfig(logToStderr bool) {
	helper.Must(config.Init())
	helper.Must(log.Init(config.CfgData.Log(logToStderr)))
}

// ExitHandle 处理系统退出信号
func ExitHandle() {

}
