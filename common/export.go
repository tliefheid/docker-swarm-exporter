package common

import (
	"github.com/toml-dev/docker-swarm-exporter/model"
)

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
			// set counter
		}
	}
}

// ExportNodeMetrics expose the node metrics
func ExportNodeMetrics(metrics []model.NodeMetrics) {
	for _, m := range metrics {
		nodeInfo.WithLabelValues(
			m.ID,
			m.Host,
			m.Role,
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
