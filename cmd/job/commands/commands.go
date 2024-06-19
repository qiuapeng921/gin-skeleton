package commands

import "github.com/urfave/cli/v2"

func GetCommands() []*cli.Command {
	var commandsArr []*cli.Command

	commandsArr = append(commandsArr, &cli.Command{
		Name:        "test",
		Aliases:     []string{"t"},
		Usage:       "测试",
		UsageText:   "job test <arg>",
		Description: "测试",
		Action:      Test,
	})

	return commandsArr
}
