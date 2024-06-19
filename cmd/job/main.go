package main

import (
	"flag"
	"github.com/urfave/cli/v2"
	"micro-base/cmd"
	"micro-base/cmd/job/commands"
	"micro-base/internal/config"
	"micro-base/internal/pkg/core/ctx"
	"os"
)

var flagConf string

func init() {
	flag.StringVar(&flagConf, "conf", config.DefaultConfigFile, "config path, eg: -conf app.yaml")
}

func main() {
	c := ctx.New()

	// 注册配置
	cmd.RegisterConfig(c, flagConf)

	app := cli.NewApp()
	app.Name = "Job"
	app.Usage = "Job Manager"

	app.Commands = commands.GetCommands()

	if err := app.Run(os.Args); err != nil {
		os.Exit(1)
	}
}
