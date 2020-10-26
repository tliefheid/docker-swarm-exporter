package common

import (
	"github.com/toml-dev/docker-swarm-exporter/model"
)

// ExportSwarmClusterInfo set gauges regarding swarm cluster information
func ExportSwarmClusterInfo(m model.SwarmMetrics) {
	swarmClusterInfo.WithLabelValues(m.GetSanitizedID()).Set(1)
	swarmClusterInfoCPU.Set(float64(m.NCPU))
	swarmClusterInfoMem.Set(float64(m.Memory))
	swarmClusterInfoImages.Set(float64(m.Images))
	swarmClusterInfoContainers.WithLabelValues("running").Set(float64(m.Container.Running))
	swarmClusterInfoContainers.WithLabelValues("paused").Set(float64(m.Container.Paused))
	swarmClusterInfoContainers.WithLabelValues("stopped").Set(float64(m.Container.Stopped))
}

// ExportServiceMetrics set correct gauges and counters for service metrics
func ExportServiceMetrics(metrics []model.ServiceMetrics) {
	for _, m := range metrics {
		serviceInfo.WithLabelValues(m.Name, m.Container, m.ServiceMode).Set(1)
		serviceCPULimit.WithLabelValues(m.Name).Set(m.Limits.ToNormalCPU())
		serviceCPUReservation.WithLabelValues(m.Name).Set((m.Reservation.ToNormalCPU()))
		serviceMemLimit.WithLabelValues(m.Name).Set(float64(m.Limits.MemoryBytes))
		serviceMemReservation.WithLabelValues(m.Name).Set(float64(m.Reservation.MemoryBytes))
		serviceTimeCreated.WithLabelValues(m.Name).Set(float64(m.TimeCreated))
		serviceTimeUpdate.WithLabelValues(m.Name).Set(float64(m.TimeUpdated))
		if m.Replicas != -1 {
			// set gauge
		}
	}
}

// ExportNodeMetrics expose the node metrics
func ExportNodeMetrics(metrics []model.NodeMetrics) {
	for _, m := range metrics {
		nodeInfo.WithLabelValues(
			m.ID,
			m.Host,
			string(m.Role),
			m.Os,
			m.Architecture,
			m.EngineVersion,
			m.NodeStatus,
			m.ManagerReachability(),
			m.IsLeader(),
		).Set(1)
		nodeResourcesCPU.WithLabelValues(m.ID).Set(m.Resources.ToNormalCPU())
		nodeResourcesMem.WithLabelValues(m.ID).Set(float64(m.Resources.MemoryBytes))
	}
}
