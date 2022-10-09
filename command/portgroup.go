package command

import (
	"fmt"

	"github.com/urfave/cli"
)

var (
	portgroupFlags = []cli.Flag{
		cli.StringFlag{
			Name: "message",
		},
		cli.BoolFlag{
			Name: "verbose, v",
		},
	}
	PortgroupCommands = []cli.Command{
		cli.Command{
			Name:   "list",
			Usage:  "show portgroup list",
			Action: print_hoge,
			Flags:  append(connectFlags, portgroupFlags...),
		},
		cli.Command{
			Name:   "create",
			Usage:  "create portgroup",
			Action: print_hoge,
		},
	}
)

func print_hoge(c *cli.Context) error {
	fmt.Println(c.String("message"))
	return nil
}
