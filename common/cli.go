package common

import (
	"context"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/swarm"
	"github.com/docker/docker/client"
)

func getCLI() *client.Client {
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}
	return cli
}

func GetServiceList() []swarm.Service {
	cli := getCLI()
	services, err := cli.ServiceList(context.Background(), types.ServiceListOptions{})
	if err != nil {
		panic(err)
	}
	cli.Close()
	return services
}
func GetNodeList() []swarm.Node {
	cli := getCLI()
	nodes, err := cli.NodeList(context.Background(), types.NodeListOptions{})
	if err != nil {
		panic(err)
	}
	cli.Close()
	return nodes
}
func GetSwarmInfo() types.Info {
	cli := getCLI()
	info, err := cli.Info(context.Background())
	if err != nil {
		panic(err)
	}
	cli.Close()
	return info
}

// func (c cliWrapper) getInfo() types.Info {
// 	info, err := c.Info(context.Background())
// 	if err != nil {
// 		panic(err)
// 	}
// 	return info
// }
