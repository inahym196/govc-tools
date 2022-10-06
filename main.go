package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

var (
	portgroupCommand = cli.Command{
		Name:  "portgroup",
		Usage: "portgroupをhogehogeする",
		Action: func(c *cli.Context) error {
			fmt.Println("portgroupをhogehoge")
			return nil
		},
	}
	permissionCommand = cli.Command{
		Name:  "permission",
		Usage: "permissionをhogehogeする",
		Action: func(c *cli.Context) error {
			fmt.Println("permissionをhogehoge")
			return nil
		},
	}
)

func main() {
	app := cli.NewApp()
	app.Name = "govc-tools"
	app.Usage = "vcenter cli tools"
	app.Version = "0.0.0"
	app.Commands = []cli.Command{
		portgroupCommand,
		permissionCommand,
	}
	app.Run(os.Args)
}
