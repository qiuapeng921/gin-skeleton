package commands

import "github.com/urfave/cli/v2"

func GetCommands() []*cli.Command {
	var commandsArr []*cli.Command

	commandsArr = append(commandsArr, &cli.Command{
		Name:        "test",
		Usage:       "job test",
		UsageText:   "job test <arg>",
		Description: "test",
		Action:      Test,
	})

	return commandsArr
}
