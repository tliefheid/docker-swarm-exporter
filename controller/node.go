package controller

import (
	"context"

	"github.com/toml-dev/docker-swarm-exporter/model"

	"github.com/docker/docker/api/types"
	"github.com/toml-dev/docker-swarm-exporter/common"
)

// UpdateNodeMetrics gathers swarm node metrics and exposes them to prometheus
func UpdateNodeMetrics() {
	cli := common.GetCLI()
	nodes, err := cli.NodeList(context.Background(), types.NodeListOptions{})
	if err != nil {
		panic(err)
	}
	var nodeMetrics []model.NodeMetrics

	for _, n := range nodes {
		role := string(n.Spec.Role)
		var nm = model.NodeMetrics{
			ID:            n.ID,
			Host:          n.Description.Hostname,
			Role:          role,
			Os:            n.Description.Platform.OS,
			Architecture:  n.Description.Platform.Architecture,
			Availability:  string(n.Spec.Availability),
			EngineVersion: n.Description.Engine.EngineVersion,
			NodeStatus:    string(n.Status.State),
		}
		if role == "manager" {
			nm.ManagerInfo.Leader = n.ManagerStatus.Leader
			nm.ManagerInfo.Reachable = string(n.ManagerStatus.Reachability)
		}
		nodeMetrics = append(nodeMetrics, nm)
	}
	cli.Close()
	common.ExportNodeMetrics(nodeMetrics)
}
