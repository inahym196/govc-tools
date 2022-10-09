package vcenter

import (
	"context"
	"net/url"
	"strconv"

	"github.com/urfave/cli"
	"github.com/vmware/govmomi"
)

var (
	ConnectFlags = []cli.Flag{
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
		cli.BoolFlag{
			Name: "disable-tls",
		},
	}
)

type ConnectVars struct {
	UserName   string
	Password   string
	HostName   string
	Port       int
	DisableTLS bool
}

func ConnectVcenter(ctx *context.Context, c *cli.Context) (*govmomi.Client, error) {
	vars := ConnectVars{
		UserName:   c.String("user"),
		Password:   c.String("pass"),
		HostName:   c.String("host"),
		Port:       c.Int("port"),
		DisableTLS: c.Bool("disable-tls"),
	}
	u := &url.URL{
		Host: vars.HostName + ":" + strconv.Itoa(vars.Port),
		User: url.UserPassword(vars.UserName, vars.Password),
		Path: "/sdk",
	}
	if vars.DisableTLS {
		u.Scheme = "http"
	} else {
		u.Scheme = "https"
	}
	client, err := govmomi.NewClient(*ctx, u, true)
	if err != nil {
		return nil, err
	}
	return client, nil
}
