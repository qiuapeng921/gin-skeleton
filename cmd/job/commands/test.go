package commands

import (
	"github.com/urfave/cli/v2"
	"micro-base/internal/app"
	"micro-base/internal/pkg/core/ctx"
	"micro-base/internal/pkg/core/log"
)

func Test(c *cli.Context) (err error) {
	context := ctx.Wrap(c.Context)

	vName := c.Args().First()
	if vName == "" {
		return cli.ShowSubcommandHelp(c)
	}

	var databases []string
	err = app.DB().Debug().Raw("SHOW DATABASES").Scan(&databases).Error
	if err != nil || len(databases) == 0 {
		return err
	}

	log.Debug(context).Msgf("%v", databases)

	return nil
}
