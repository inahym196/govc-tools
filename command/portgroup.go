package command

import (
	"context"
	"fmt"

	"os"

	vcenter "github.com/inahym196/govc-tools/lib"
	"github.com/urfave/cli"
	"github.com/vmware/govmomi/find"
	"github.com/vmware/govmomi/vim25/types"
)

var (
	portgroupFlags    = []cli.Flag{}
	PortgroupCommands = []cli.Command{
		cli.Command{
			Name:   "list",
			Usage:  "show portgroup list",
			Action: list_portgroup,
			Flags:  append(vcenter.ConnectFlags, portgroupFlags...),
		},
	}
)

func exit(err error) {
	fmt.Fprintf(os.Stderr, "Error: %s\n", err)
	os.Exit(1)
}

func list_portgroup(c *cli.Context) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	args, err := vcenter.NewConnectArgs(c)
	if err != nil {
		exit(err)
	}

	vcenter, err := vcenter.ConnectVcenter(&ctx, args)
	if err != nil {
		exit(err)
	}

	f := find.NewFinder(vcenter.Client, true)
	nws, err := f.NetworkList(ctx, "/DC0/network/*")
	if err != nil {
		exit(err)
	}

	var refs []types.ManagedObjectReference
	for _, nw := range nws {
		ref := nw.Reference()
		if ref.Type == "DistributedVirtualPortgroup" {
			refs = append(refs, nw.Reference())
		}
	}
	for _, ref := range refs {
		fmt.Printf("%v\n", ref.Value)
	}
	return nil
}
