package main

import (
	"fmt"
	"log"
	"os"

	"github.com/inahym196/govc-tools/command"
	"github.com/urfave/cli"
)

var (
	portgroupCommand = cli.Command{
		Name:        "portgroup",
		Usage:       "portgroupをhogehogeする",
		Subcommands: command.PortgroupCommands,
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
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
