package common

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	serviceInfo           = promauto.NewGaugeVec(prometheus.GaugeOpts{Name: serviceSpecInfo}, []string{lblServiceName, lblContainer, lblDeployMode})
	serviceMemLimit       = promauto.NewGaugeVec(prometheus.GaugeOpts{Name: serviceSpecMemLimit}, []string{lblServiceName})
	serviceMemReservation = promauto.NewGaugeVec(prometheus.GaugeOpts{Name: serviceSpecMemReservation}, []string{lblServiceName})
	serviceCPULimit       = promauto.NewGaugeVec(prometheus.GaugeOpts{Name: serviceSpecCPULimit}, []string{lblServiceName})
	serviceCPUReservation = promauto.NewGaugeVec(prometheus.GaugeOpts{Name: serviceSpecCPUReservation}, []string{lblServiceName})
	serviceReplicas       = promauto.NewGaugeVec(prometheus.GaugeOpts{Name: serviceSpecReplicas}, []string{lblServiceName})
	serviceTimeUpdate     = promauto.NewGaugeVec(prometheus.GaugeOpts{Name: serviceSpecTimeUpdated}, []string{lblServiceName})
	serviceTimeCreated    = promauto.NewGaugeVec(prometheus.GaugeOpts{Name: serviceSpecTimeCreated}, []string{lblServiceName})
)
