package common

const (
	lblServiceName = "service_name"
	lblContainer   = "container_image"
	lblDeployMode  = "deploy_mode"
)

const (
	ServiceModeUnknown    = "unknown"
	ServiceModeReplicated = "replicated"
	ServiceModeGlobal     = "global"
)

const (
	prefixMetric      = ""
	prefixSwarmNode   = "swarm_node_"
	prefixServiceSpec = "service_spec_"

	info = "info"

	cpuLimit       = "cpu_limit"
	cpuReservation = "cpu_reservation"
	memLimit       = "memory_limit"
	memReservation = "memory_reservation"

	resourcesCPU = "resources_cpu"
	resourcesMem = "resources_memory"

	timeCreated = "time_created"
	timeUpdated = "time_updated"
	replicas    = "replicas"
)

const (
	swarmNodeInfo = prefixMetric + prefixSwarmNode + info

	serviceSpecInfo           = prefixMetric + prefixServiceSpec + info
	serviceSpecMemLimit       = prefixMetric + prefixServiceSpec + memLimit
	serviceSpecCPULimit       = prefixMetric + prefixServiceSpec + cpuLimit
	serviceSpecMemReservation = prefixMetric + prefixServiceSpec + memReservation
	serviceSpecCPUReservation = prefixMetric + prefixServiceSpec + cpuReservation
	serviceSpecTimeCreated    = prefixMetric + prefixServiceSpec + timeCreated
	serviceSpecTimeUpdated    = prefixMetric + prefixServiceSpec + timeUpdated
	serviceSpecReplicas       = prefixMetric + prefixServiceSpec + replicas
)
