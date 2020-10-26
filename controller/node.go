package controller

import (
	"github.com/toml-dev/docker-swarm-exporter/model"

	"github.com/docker/docker/api/types/swarm"
	"github.com/toml-dev/docker-swarm-exporter/common"
)

// UpdateNodeMetrics gathers swarm node metrics and exposes them to prometheus
func UpdateNodeMetrics() {
	nodes := common.GetNodeList()
	var nodeMetrics []model.NodeMetrics

	for _, n := range nodes {
		role := n.Spec.Role
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
		nm.Resources.MemoryBytes = n.Description.Resources.MemoryBytes
		nm.Resources.NanoCPUs = n.Description.Resources.NanoCPUs
		if role == swarm.NodeRoleManager {
			nm.ManagerInfo.Leader = n.ManagerStatus.Leader
			nm.ManagerInfo.Reachability = string(n.ManagerStatus.Reachability)
		}
		// fmt.Printf("\nNode: %s\nInfo:\n%#v\nMetrics\n%#v\n", n.ID, n, nm)
		nodeMetrics = append(nodeMetrics, nm)
	}
	common.ExportNodeMetrics(nodeMetrics)
}
