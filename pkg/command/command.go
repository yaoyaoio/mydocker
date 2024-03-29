//__author__ = "YaoYao"
//Date: 2019-07-27
package command

import (
	"fmt"
	"github.com/urfave/cli"
	"mydocker/pkg/container"
	"mydocker/pkg/run"
)

var Commands = []cli.Command{
	runCommand,
	initCommand,
}

var runCommand = cli.Command{
	Name:  "run",
	Usage: "Create a container with namespace and cgrops limit mydocker run -it [command]",
	Flags: []cli.Flag{
		cli.BoolFlag{
			Name:  "ti",
			Usage: "enable tty",
		},
	},
	Action: func(ctx *cli.Context) error {
		if len(ctx.Args()) < 1 {
			return fmt.Errorf("missing container command")
		}
		cmd := ctx.Args().Get(0)
		tty := ctx.Bool("ti")
		run.Run(tty, cmd)
		return nil
	},
}

var initCommand = cli.Command{
	Name:  "init",
	Usage: "Init container process run user's process in container Do not call it outside",
	Action: func(ctx *cli.Context) error {
		cmd := ctx.Args().Get(0)
		err := container.RunContainerInitProcess(cmd, nil)
		return err
	},
}
