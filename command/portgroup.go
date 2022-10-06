package command

import (
	"fmt"

	"github.com/urfave/cli"
)

var (
	PortgroupFlags = []cli.Flag{
		cli.StringFlag{
			Name: "message",
		},
	}
	PortgroupCommands = []cli.Command{
		cli.Command{
			Name:   "list",
			Usage:  "show portgroup list",
			Action: print_hoge,
			Flags:  append(connectFlags, PortgroupFlags...),
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
