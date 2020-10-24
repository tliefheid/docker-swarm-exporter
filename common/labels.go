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

	cpuLimit       = "_cpu_limit"
	cpuReservation = "_cpu_reservation"
	memLimit       = "_memory_limit"
	memReservation = "_memory_reservation"

	resourcesCPU = "_resources_cpu"
	resourcesMem = "_resources_memory"
)

const (
	SwarmNodeInfo = prefixMetric + prefixSwarmNode + info

	ServiceSpecInfo           = prefixMetric + prefixServiceSpec + info
	ServiceSpecMemLimit       = prefixMetric + prefixServiceSpec + memLimit
	ServiceSpecCPULimit       = prefixMetric + prefixServiceSpec + cpuLimit
	ServiceSpecMemReservation = prefixMetric + prefixServiceSpec + memReservation
	ServiceSpecCPUReservation = prefixMetric + prefixServiceSpec + cpuReservation
)
