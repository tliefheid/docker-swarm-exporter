package common

import (
	"github.com/toml-dev/docker-swarm-exporter/metrics"
)

// ExportServiceMetrics set correct gauges and counters for service metrics
func ExportServiceMetrics(metrics []metrics.ServiceMetrics) {
	for _, m := range metrics {
		serviceInfo.WithLabelValues(m.Name, m.Container, m.ServiceMode).Set(1)
		serviceCPULimit.WithLabelValues(m.Name).Set(toNormalCPU(m.CPULimit))
		serviceCPUReservation.WithLabelValues(m.Name).Set(toNormalCPU(m.CPUReservation))
		serviceMemLimit.WithLabelValues(m.Name).Set(float64(m.MemLimit))
		serviceMemReservation.WithLabelValues(m.Name).Set(float64(m.MemReservation))
		serviceTimeCreated.WithLabelValues(m.Name).Set(float64(m.TimeCreated))
		serviceTimeUpdate.WithLabelValues(m.Name).Set(float64(m.TimeUpdated))
		if m.Replicas != -1 {

		}
	}
}
