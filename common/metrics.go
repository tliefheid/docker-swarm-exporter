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
var (
	nodeInfo         = promauto.NewGaugeVec(prometheus.GaugeOpts{Name: swarmNodeInfo}, []string{lblNodeID, lblHost, lblRole, lblOS, lblArch, lblEngVersion, lblNodeState, lblManagerReachable, lblLeader})
	nodeResourcesCPU = promauto.NewGaugeVec(prometheus.GaugeOpts{Name: swarmNodeResourceCPU}, []string{lblNodeID})
	nodeResourcesMem = promauto.NewGaugeVec(prometheus.GaugeOpts{Name: swarmNodeResourceMemory}, []string{lblNodeID})
)
var (
	swarmClusterInfo = promauto.NewGaugeVec(prometheus.GaugeOpts{Name: swarmInfo}, []string{lblID})
	swarmContainers  = promauto.NewGaugeVec(prometheus.GaugeOpts{Name: swarmClusterContainers}, []string{lblState})
)
