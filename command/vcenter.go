package command

import "github.com/urfave/cli"

var (
	connectFlags = []cli.Flag{
		cli.StringFlag{
			Name:  "user",
			Value: "user",
		},
		cli.StringFlag{
			Name:  "pass",
			Value: "pass",
		},
		cli.StringFlag{
			Name:  "host",
			Value: "localhost",
		},
		cli.IntFlag{
			Name:  "port, p",
			Value: 8989,
		},
	}
)
