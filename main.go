package main

import (
	"fmt"
	"os"

	"github.com/inahym196/govc-tools/command"
	"github.com/joho/godotenv"
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

func exit(err error) {
	fmt.Fprintf(os.Stderr, "Error: %s\n", err)
	os.Exit(1)
}

func load_env() {
	err := godotenv.Load(".cred.env")
	if err != nil {
		exit(err)
	}
}

func main() {
	load_env()
	app := cli.NewApp()
	app.Name = "govc-tools"
	app.Usage = "vcenter cli tools"
	app.Commands = []cli.Command{
		portgroupCommand,
		permissionCommand,
	}
	if err := app.Run(os.Args); err != nil {
		exit(err)
	}
}
