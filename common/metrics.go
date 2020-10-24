package common

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	serviceInfo           = promauto.NewGaugeVec(prometheus.GaugeOpts{Name: ServiceSpecInfo}, []string{lblServiceName, lblContainer, lblDeployMode})
	serviceMemLimit       = promauto.NewGaugeVec(prometheus.GaugeOpts{Name: ServiceSpecMemLimit}, []string{lblServiceName})
	serviceMemReservation = promauto.NewGaugeVec(prometheus.GaugeOpts{Name: ServiceSpecMemReservation}, []string{lblServiceName})
	// serviceTimeUpdate     = promauto.NewGaugeVec(prometheus.GaugeOpts{Name: "service_spec_time_last_update"}, []string{lblServiceName})
	serviceCPULimit       = promauto.NewGaugeVec(prometheus.GaugeOpts{Name: ServiceSpecCPULimit}, []string{lblServiceName})
	serviceCPUReservation = promauto.NewGaugeVec(prometheus.GaugeOpts{Name: ServiceSpecCPUReservation}, []string{lblServiceName})
	serviceTimeUpdate     = promauto.NewGaugeVec(prometheus.GaugeOpts{Name: "service_spec_time_last_update"}, []string{lblServiceName})
	serviceTimeCreated    = promauto.NewGaugeVec(prometheus.GaugeOpts{Name: "service_spec_time_created"}, []string{lblServiceName})
)
