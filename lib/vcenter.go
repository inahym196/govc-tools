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
			Name:   "user",
			EnvVar: "GOVCTOOLS_USER",
		},
		cli.StringFlag{
			Name:   "pass",
			EnvVar: "GOVCTOOLS_PASS",
		},
		cli.StringFlag{
			Name:   "host",
			EnvVar: "GOVCTOOLS_HOST",
		},
		cli.IntFlag{
			Name:   "port, p",
			EnvVar: "GOVCTOOLS_PORT",
		},
		cli.BoolFlag{
			Name:   "disable-tls",
			EnvVar: "GOVCTOOLS_DISABLE_TLS",
		},
	}
)

type connectArgs struct {
	UserName   string
	Password   string
	HostName   string
	Port       string
	DisableTLS bool
}

func NewConnectArgs(c *cli.Context) (*connectArgs, error) {
	return &connectArgs{
		UserName:   c.String("user"),
		Password:   c.String("pass"),
		HostName:   c.String("host"),
		Port:       strconv.Itoa(c.Int("port")),
		DisableTLS: c.Bool("disable-tls"),
	}, nil
}

func ConnectVcenter(ctx *context.Context, args *connectArgs) (*govmomi.Client, error) {
	u := &url.URL{
		Scheme: "https",
		Host:   args.HostName + ":" + args.Port,
		User:   url.UserPassword(args.UserName, args.Password),
		Path:   "/sdk",
	}
	if args.DisableTLS {
		u.Scheme = "http"
	}
	client, err := govmomi.NewClient(*ctx, u, true)
	if err != nil {
		return nil, err
	}

	return client, nil
}
