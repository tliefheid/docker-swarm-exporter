package common

import (
	"github.com/toml-dev/docker-swarm-exporter/metrics"
)

// ExportServiceMetrics set correct gauges and counters for service metrics
func ExportServiceMetrics(metrics []metrics.ServiceMetrics) {
	for _, m := range metrics {
		serviceInfo.WithLabelValues(m.Name, m.Container, m.ServiceMode).Set(1)
	}
}
