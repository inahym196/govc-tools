package command

import (
	"context"

	"github.com/vmware/govmomi/property"

	"github.com/k0kubun/pp/v3"
	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/types"

	vcenter "github.com/inahym196/govc-tools/lib"
	"github.com/urfave/cli"
	"github.com/vmware/govmomi/find"
)

var (
	portgroupFilterCommonFlags = []cli.Flag{
		cli.StringFlag{
			Name: "portgroup",
		},
	}
	PortgroupFilterCommands = []cli.Command{
		cli.Command{
			Name:   "export",
			Usage:  "export portgroup's filter to file",
			Action: export_portgroup_filter,
			Flags:  append(vcenter.ConnectFlags, portgroupFilterCommonFlags...),
		},
	}
)

func export_portgroup_filter(c *cli.Context) error {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	args, err := vcenter.NewConnectArgs(c)
	if err != nil {
		return err
	}

	vcenter, err := vcenter.ConnectVcenter(&ctx, args)
	if err != nil {
		return err
	}

	finder := find.NewFinder(vcenter.Client, true)
	dc, _ := finder.DatacenterList(ctx, "*")
	finder.SetDatacenter(dc[0])
	net, err := finder.Network(ctx, "/DC0/network/DVS0")
	if err != nil {
		return err
	}

	dvs, _ := net.(*object.DistributedVirtualSwitch)
	var s mo.DistributedVirtualSwitch
	err = dvs.Properties(ctx, dvs.Reference(), []string{"config"}, &s)
	if err != nil {
		return err
	}
	spec := &types.VMwareDVSConfigSpec{
		DVSConfigSpec: types.DVSConfigSpec{
			Name: "hoge",
		},
		MaxMtu: 100,
	}
	task, err := dvs.Reconfigure(ctx, spec)
	if err != nil {
		return err
	}
	result, err := task.WaitForResult(ctx, nil)
	if err != nil {
		return err
	}
	pp.Print(result)
	netRef, err := finder.Network(ctx, "/DC0/network/DVS0")
	if err != nil {
		return err
	}
	var data mo.DistributedVirtualSwitch
	pc := property.DefaultCollector(vcenter.Client)
	pc.RetrieveOne(ctx, netRef.Reference(), nil, &data)
	pp.Print(data)
	return nil
	// rule1Quali1 := types.DvsIpNetworkRuleQualifier{
	// 	DestinationAddress: &types.IpAddress{},
	// }
	// dvsTrafficFilterConfig := types.DvsTrafficFilterConfig{
	// 	TrafficRuleset: &types.DvsTrafficRuleset{
	// 		Enabled: &[]bool{true}[0],
	// 		Rules: []types.DvsTrafficRule{
	// 			types.DvsTrafficRule{
	// 				Action:      &types.DvsAcceptNetworkRuleAction{},
	// 				Description: "rule1",
	// 				Key:         "rule1",
	// 				Qualifier: []types.BaseDvsNetworkRuleQualifier{
	// 					rule1Quali1.GetDvsNetworkRuleQualifier(),
	// 				},
	// 			},
	// 		},
	// 	},
	// }
	// dvPortSetting := &types.DVPortSetting{
	// 	FilterPolicy: &types.DvsFilterPolicy{
	// 		FilterConfig: []types.BaseDvsFilterConfig{
	// 			dvsTrafficFilterConfig.GetDvsTrafficFilterConfig(),
	// 		},
	// 	},
	// }
	// dvPortConfigSpec := types.DVPortConfigSpec{
	// 	Name:    "hoge",
	// 	Setting: dvPortSetting,
	// }
	// pp.Print(dvPortConfigSpec.Setting)
}
